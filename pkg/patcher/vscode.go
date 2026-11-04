package patcher

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var (
	channelChangedRe         = regexp.MustCompile(`const\s+isChannelChanged\s*=\s*manifestFetched\s*&&\s*lastInstalledUrl\s*!==\s*releaseBaseUrl;`)
	channelChangedMinifiedRe = regexp.MustCompile(`let\s+([a-zA-Z0-9_$]+)\s*=\s*[^;]+;\s*([a-zA-Z0-9_$]+\s*&&\s*[a-zA-Z0-9_$]+\.appendLine\(\s*[` + "`" + `'"]\[INSTALL\]\s*Release channel change detected)`)
	channelPatched           = `const isChannelChanged = false; /*patched:channel_lock*/`

	downloadGuardTarget = regexp.MustCompile(`([a-zA-Z0-9_$]+\.appendLine\(\s*['"]\[INSTALL\]\s*Checking Antigravity releases\.\.\.['"]\s*\);)`)
	downloadGuardCode   = `{const __primary=options.targetPathOverride||getInstalledTargetPath(); const __candidates=[__primary, (0,path_1.join)((0,path_1.dirname)(__primary),'antigravity'+((0,path_1.extname)(__primary)||''))]; for(const __p of __candidates){if(__p&&await pathExists(__p)){ outputChannel.appendLine('[INSTALL] Existing binary found at '+__p+'. Skipping version check and re-download.');return __p;}}}`
)

// FindVsCodeExtension finds the extension.js of google.google-antigravity
func FindVsCodeExtension() string {
	home, _ := os.UserHomeDir()
	candidateDirs := []string{
		filepath.Join(home, ".vscode", "extensions"),
		filepath.Join(home, ".vscode-insiders", "extensions"),
		filepath.Join(home, ".vscode-oss", "extensions"),
		filepath.Join(home, ".vscode-server", "extensions"),
		filepath.Join(home, ".antigravity-ide", "extensions"),
		filepath.Join(home, ".antigravity", "extensions"),
	}

	if runtime.GOOS == "windows" {
		localApp := os.Getenv("LOCALAPPDATA")
		candidateDirs = append(candidateDirs,
			filepath.Join(localApp, "Programs", "Antigravity IDE", "resources", "app", "extensions"),
			filepath.Join(localApp, "Programs", "Antigravity", "resources", "app", "extensions"),
			filepath.Join(localApp, "Programs", "antigravity", "resources", "app", "extensions"),
		)
	} else if runtime.GOOS == "linux" {
		candidateDirs = append(candidateDirs,
			"/usr/share/antigravity/resources/app/extensions",
			"/usr/share/antigravity-ide/resources/app/extensions",
			"/opt/antigravity/resources/app/extensions",
			"/opt/Antigravity IDE/resources/app/extensions",
		)
	}

	var latestExtJs string
	var latestModTime int64

	for _, dir := range candidateDirs {
		// 1. Direct subfolder check for 'agy' extension in bundled builds
		directAgy := filepath.Join(dir, "agy")
		if fi, err := os.Stat(directAgy); err == nil && fi.IsDir() {
			for _, sub := range []string{filepath.Join("dist", "extension.js"), "extension.js"} {
				target := filepath.Join(directAgy, sub)
				if fiTarget, errTarget := os.Stat(target); errTarget == nil && !fiTarget.IsDir() {
					if fiTarget.ModTime().Unix() > latestModTime {
						latestModTime = fiTarget.ModTime().Unix()
						latestExtJs = target
					}
				}
			}
		}

		// 2. Scan directory entries for google.google-antigravity-* or agy-*
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if e.IsDir() && (strings.HasPrefix(e.Name(), "google.google-antigravity-") || strings.EqualFold(e.Name(), "agy") || strings.HasPrefix(e.Name(), "agy-")) {
				for _, sub := range []string{"extension.js", filepath.Join("dist", "extension.js")} {
					extJs := filepath.Join(dir, e.Name(), sub)
					if fi, err := os.Stat(extJs); err == nil && !fi.IsDir() {
						if fi.ModTime().Unix() > latestModTime {
							latestModTime = fi.ModTime().Unix()
							latestExtJs = extJs
						}
					}
				}
			}
		}
	}
	return latestExtJs
}

// PatchVsCodeExtension patches extension.js to disable re-download and lock channel
func PatchVsCodeExtension(extPath string) (string, error) {
	if extPath == "" {
		extPath = FindVsCodeExtension()
		if extPath == "" {
			return "", fmt.Errorf("VS Code extension google.google-antigravity not found")
		}
	}

	data, err := os.ReadFile(extPath)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", extPath, err)
	}

	content := string(data)
	alreadyPatched := (strings.Contains(content, channelPatched) && strings.Contains(content, "Skipping version check and re-download")) ||
		strings.Contains(content, "/*patched:channel_lock*/")
	if alreadyPatched {
		return fmt.Sprintf("Already patched: %s", filepath.Base(filepath.Dir(extPath))), nil
	}

	// Apply Part 1: Channel lock (Classic unminified and v1.10.0+ minified)
	patched := content
	channelApplied := false

	if channelChangedRe.MatchString(patched) {
		patched = channelChangedRe.ReplaceAllString(patched, channelPatched)
		channelApplied = true
	} else if channelChangedMinifiedRe.MatchString(patched) {
		patched = channelChangedMinifiedRe.ReplaceAllString(patched, "let $1 = false; /*patched:channel_lock*/ $2")
		channelApplied = true
	} else if strings.Contains(content, channelPatched) || strings.Contains(content, "/*patched:channel_lock*/") {
		channelApplied = true
	}

	// Apply Part 2: Re-download guard injection (if unminified AST present)
	guardApplied := strings.Contains(patched, "Skipping version check and re-download")
	if !guardApplied && downloadGuardTarget.MatchString(patched) && strings.Contains(patched, "options.targetPathOverride") {
		patched = downloadGuardTarget.ReplaceAllString(patched, "$1\n"+downloadGuardCode)
		guardApplied = true
	}

	if !channelApplied && !guardApplied {
		return "", fmt.Errorf("patch signatures not found in %s (unsupported or already modified build) — extension was NOT patched", filepath.Base(filepath.Dir(extPath)))
	}

	if _, err := BackupFile(extPath); err != nil {
		return "", fmt.Errorf("backup failed: %w", err)
	}

	if err := AtomicWrite(extPath, []byte(patched)); err != nil {
		return "", fmt.Errorf("failed to write patched extension.js: %w", err)
	}

	// Also patch ~/.gemini/bin/antigravity if it exists
	home, _ := os.UserHomeDir()
	downloadedBin := filepath.Join(home, ".gemini", "bin", "antigravity")
	if fi, err := os.Stat(downloadedBin); err == nil && !fi.IsDir() {
		_, _ = PatchAgy(downloadedBin)
	}

	return fmt.Sprintf("Successfully patched VS Code extension %s (re-download guard + channel lock)", filepath.Base(filepath.Dir(extPath))), nil
}
