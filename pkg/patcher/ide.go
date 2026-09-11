package patcher

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var (
	ideTargetRe = regexp.MustCompile(`(resetIsTierGCPTos\(\),)this\.[A-Za-z_$0-9]+\.isGoogleInternal`)
	idePatchedStr = "resetIsTierGCPTos(),true"
)

// FindIdeMainJs searches for main.js in Antigravity IDE installations
func FindIdeMainJs() string {
	home, _ := os.UserHomeDir()

	candidates := []string{
		filepath.Join(home, "Applications", "Antigravity", "resources", "app", "out", "main.js"),
		filepath.Join(home, "Applications", "Antigravity", "resources", "app", "main.js"),
		filepath.Join(home, "Applications", "Antigravity-x64", "resources", "app", "out", "main.js"),
		filepath.Join(home, "antigravity", "resources", "app", "out", "main.js"),
		filepath.Join(home, "antigravity", "resources", "app", "main.js"),
		"/usr/share/antigravity-ide/resources/app/out/main.js",
		"/usr/share/antigravity-ide/resources/app/main.js",
		"/opt/Antigravity IDE/resources/app/out/main.js",
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Antigravity IDE", "resources", "app", "out", "main.js"),
		"/Applications/Antigravity IDE.app/Contents/Resources/app/out/main.js",
	}

	for _, c := range candidates {
		if fi, err := os.Stat(c); err == nil && !fi.IsDir() {
			return c
		}
	}
	return ""
}

// ClearIdeCache removes VS Code compile cache so patched JS reloads cleanly
func ClearIdeCache() {
	home, _ := os.UserHomeDir()
	cacheDirs := []string{
		filepath.Join(home, ".config", "Antigravity IDE", "CachedData"),
		filepath.Join(home, ".config", "Antigravity IDE", "Code Cache"),
		filepath.Join(home, ".config", "Antigravity", "CachedData"),
		filepath.Join(home, ".config", "Antigravity", "Code Cache"),
		filepath.Join(os.Getenv("APPDATA"), "Antigravity IDE", "CachedData"),
		filepath.Join(os.Getenv("APPDATA"), "Antigravity IDE", "Code Cache"),
	}

	for _, d := range cacheDirs {
		_ = os.RemoveAll(d)
	}
}

// PatchIdeMainJs replaces the isGoogleInternal check in main.js
func PatchIdeMainJs(mainJsPath string) (string, error) {
	if mainJsPath == "" {
		mainJsPath = FindIdeMainJs()
		if mainJsPath == "" {
			return "", fmt.Errorf("Antigravity IDE main.js not found")
		}
	}

	data, err := os.ReadFile(mainJsPath)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", mainJsPath, err)
	}

	content := string(data)
	if stringContains(content, idePatchedStr) {
		return fmt.Sprintf("Already patched: %s", filepath.Base(mainJsPath)), nil
	}

	if !ideTargetRe.MatchString(content) {
		return "", fmt.Errorf("isGoogleInternal pattern not found in %s", filepath.Base(mainJsPath))
	}

	if _, err := BackupFile(mainJsPath); err != nil {
		return "", fmt.Errorf("backup failed: %w", err)
	}

	patched := ideTargetRe.ReplaceAllString(content, idePatchedStr)

	if err := AtomicWrite(mainJsPath, []byte(patched)); err != nil {
		return "", fmt.Errorf("failed to write patched main.js: %w", err)
	}

	ClearIdeCache()

	return fmt.Sprintf("Successfully patched %s (isGoogleInternal -> true) & cleared compile cache", filepath.Base(mainJsPath)), nil
}

func stringContains(s, substr string) bool {
	return len(s) >= len(substr) && regexp.MustCompile(regexp.QuoteMeta(substr)).MatchString(s)
}
