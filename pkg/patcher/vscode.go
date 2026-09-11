package patcher

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	channelChangedRe = regexp.MustCompile(`const\s+isChannelChanged\s*=\s*manifestFetched\s*&&\s*lastInstalledUrl\s*!==\s*releaseBaseUrl;`)
	channelPatched   = `const isChannelChanged = false;`

	downloadGuardTarget = regexp.MustCompile(`(outputChannel\.appendLine\(\s*['"]\[INSTALL\]\s*Checking Antigravity releases\.\.\.['"]\s*\);)`)
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
	}

	var latestExtJs string
	var latestModTime int64

	for _, dir := range candidateDirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if e.IsDir() && strings.HasPrefix(e.Name(), "google.google-antigravity-") {
				extJs := filepath.Join(dir, e.Name(), "extension.js")
				if fi, err := os.Stat(extJs); err == nil {
					if fi.ModTime().Unix() > latestModTime {
						latestModTime = fi.ModTime().Unix()
						latestExtJs = extJs
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
	if strings.Contains(content, channelPatched) && strings.Contains(content, "Skipping version check and re-download") {
		return fmt.Sprintf("Already patched: %s", filepath.Base(filepath.Dir(extPath))), nil
	}

	if _, err := BackupFile(extPath); err != nil {
		return "", fmt.Errorf("backup failed: %w", err)
	}

	// Apply Part 1: Channel lock
	content = channelChangedRe.ReplaceAllString(content, channelPatched)

	// Apply Part 2: Re-download guard injection
	if downloadGuardTarget.MatchString(content) {
		content = downloadGuardTarget.ReplaceAllString(content, "$1\n"+downloadGuardCode)
	}

	if err := AtomicWrite(extPath, []byte(content)); err != nil {
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
