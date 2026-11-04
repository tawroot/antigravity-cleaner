package patcher

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	ideTargetRe = regexp.MustCompile(`(resetIsTierGCPTos\(\),)(?:this\.)?[A-Za-z_$0-9.]*isGoogleInternal`)
	idePatchedStr = "resetIsTierGCPTos(),true"
)

// FindAllIdeMainJs searches for all main.js occurrences across Antigravity IDE and installations
func FindAllIdeMainJs() []string {
	home, _ := os.UserHomeDir()
	var results []string
	seen := make(map[string]bool)
	add := func(p string) {
		if p == "" {
			return
		}
		if real, err := filepath.EvalSymlinks(p); err == nil && real != "" {
			p = real
		}
		clean := filepath.Clean(p)
		if !seen[clean] {
			if fi, err := os.Stat(clean); err == nil && !fi.IsDir() {
				seen[clean] = true
				results = append(results, clean)
			}
		}
	}

	var appPaths []string
	if running := FindRunningAntigravityApp(nil); running != "" {
		appPaths = append(appPaths, running)
	}
	for _, name := range []string{"antigravity", "antigravity.exe", "antigravity-ide"} {
		if p, err := exec.LookPath(name); err == nil && p != "" {
			appPaths = append(appPaths, p)
		}
	}

	jsSubpaths := []string{
		filepath.Join("resources", "app", "out", "main.js"),
		filepath.Join("resources", "app", "main.js"),
		filepath.Join("resources", "app", "dist", "main.js"),
		filepath.Join("Contents", "Resources", "app", "out", "main.js"),
		filepath.Join("Contents", "Resources", "app", "main.js"),
		filepath.Join("Contents", "Resources", "app", "dist", "main.js"),
		filepath.Join("Resources", "app", "out", "main.js"),
		filepath.Join("Resources", "app", "main.js"),
		filepath.Join("Resources", "app", "dist", "main.js"),
	}

	for _, appPath := range appPaths {
		for _, root := range getCandidateRoots(appPath) {
			for _, sub := range jsSubpaths {
				add(filepath.Join(root, sub))
			}
		}
	}

	candidates := []string{
		"/usr/share/antigravity/resources/app/out/main.js",
		"/usr/share/antigravity/resources/app/main.js",
		"/usr/share/antigravity-ide/resources/app/out/main.js",
		"/usr/share/antigravity-ide/resources/app/main.js",
		"/opt/antigravity/resources/app/out/main.js",
		"/opt/antigravity/resources/app/main.js",
		"/opt/Antigravity/resources/app/out/main.js",
		"/opt/Antigravity/resources/app/main.js",
		"/opt/antigravity-ide/resources/app/out/main.js",
		"/opt/Antigravity IDE/resources/app/out/main.js",
		filepath.Join(home, ".local", "share", "antigravity", "resources", "app", "out", "main.js"),
		filepath.Join(home, ".local", "share", "antigravity", "resources", "app", "main.js"),
		filepath.Join(home, "Applications", "Antigravity", "resources", "app", "out", "main.js"),
		filepath.Join(home, "Applications", "Antigravity", "resources", "app", "main.js"),
		filepath.Join(home, "Applications", "antigravity", "resources", "app", "out", "main.js"),
		filepath.Join(home, "Applications", "antigravity", "resources", "app", "main.js"),
		filepath.Join(home, "Applications", "Antigravity-x64", "resources", "app", "out", "main.js"),
		filepath.Join(home, "antigravity", "resources", "app", "out", "main.js"),
		filepath.Join(home, "antigravity", "resources", "app", "main.js"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Antigravity", "resources", "app", "out", "main.js"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Antigravity", "resources", "app", "main.js"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "antigravity", "resources", "app", "out", "main.js"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "antigravity", "resources", "app", "main.js"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Antigravity IDE", "resources", "app", "out", "main.js"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Antigravity IDE", "resources", "app", "main.js"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Antigravity IDE", "resources", "app", "dist", "main.js"),
		filepath.Join(os.Getenv("PROGRAMFILES"), "Antigravity", "resources", "app", "out", "main.js"),
		filepath.Join(os.Getenv("ProgramFiles(x86)"), "Antigravity", "resources", "app", "out", "main.js"),
		"/Applications/Antigravity.app/Contents/Resources/app/out/main.js",
		"/Applications/Antigravity.app/Contents/Resources/app/main.js",
		"/Applications/Antigravity IDE.app/Contents/Resources/app/out/main.js",
		filepath.Join(home, "Applications", "Antigravity.app", "Contents", "Resources", "app", "out", "main.js"),
	}

	for _, c := range candidates {
		add(c)
	}

	return results
}

// FindIdeMainJs searches for main.js in Antigravity IDE installations
func FindIdeMainJs() string {
	all := FindAllIdeMainJs()
	if len(all) > 0 {
		return all[0]
	}
	return ""
}

// ClearIdeCache removes VS Code compile cache, V8 bytecode caches, and partition caches
// so patched JS reloads cleanly across Linux, macOS, and Windows.
func ClearIdeCache() {
	home, _ := os.UserHomeDir()
	var cacheDirs []string
	seen := make(map[string]bool)

	add := func(p string) {
		if p == "" {
			return
		}
		clean := filepath.Clean(p)
		if !seen[clean] {
			seen[clean] = true
			cacheDirs = append(cacheDirs, clean)
		}
	}

	// 1. Candidate parent config & cache roots across all OSes
	var parentDirs []string
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		parentDirs = append(parentDirs, xdg)
	}
	if home != "" {
		parentDirs = append(parentDirs,
			filepath.Join(home, ".config"),
			filepath.Join(home, "Library", "Application Support"),
			filepath.Join(home, "Library", "Caches"),
		)
	}
	if appData := os.Getenv("APPDATA"); appData != "" {
		parentDirs = append(parentDirs, appData)
	}
	if localAppData := os.Getenv("LOCALAPPDATA"); localAppData != "" {
		parentDirs = append(parentDirs, localAppData)
	}

	// Subfolders holding V8 and Electron caches
	cacheSubdirs := []string{
		"CachedData",
		"Code Cache",
		filepath.Join("Code Cache", "js"),
		filepath.Join("Code Cache", "wasm"),
		filepath.Join("Partitions"),
	}

	for _, parent := range parentDirs {
		entries, err := os.ReadDir(parent)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if entry.IsDir() {
				nameLower := strings.ToLower(entry.Name())
				if strings.Contains(nameLower, "antigravity") {
					targetRoot := filepath.Join(parent, entry.Name())
					for _, sub := range cacheSubdirs {
						add(filepath.Join(targetRoot, sub))
					}
				}
			}
		}
	}

	// Standard macOS cache bundles
	if home != "" {
		add(filepath.Join(home, "Library", "Caches", "com.google.antigravity"))
		add(filepath.Join(home, "Library", "Caches", "com.google.antigravity.helper"))
	}

	for _, d := range cacheDirs {
		if d != "" {
			_ = os.RemoveAll(d)
		}
	}
}

// PatchSingleIdeMainJs patches a single main.js file
func PatchSingleIdeMainJs(mainJsPath string) (string, error) {
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

// PatchIdeMainJs replaces the isGoogleInternal check in main.js across all installations
func PatchIdeMainJs(mainJsPath string) (string, error) {
	if mainJsPath != "" {
		return PatchSingleIdeMainJs(mainJsPath)
	}

	all := FindAllIdeMainJs()
	if len(all) == 0 {
		return "", fmt.Errorf("Antigravity IDE main.js not found")
	}

	if len(all) == 1 {
		return PatchSingleIdeMainJs(all[0])
	}

	var patchedCount int
	var lastRes string
	var lastErr error

	for _, target := range all {
		res, err := PatchSingleIdeMainJs(target)
		if err == nil {
			patchedCount++
			lastRes = res
		} else {
			lastErr = err
		}
	}

	if patchedCount == 0 && lastErr != nil {
		return "", lastErr
	}

	if patchedCount > 1 {
		return fmt.Sprintf("Successfully patched %d IDE main.js targets & cleared compile cache", patchedCount), nil
	}
	return lastRes, nil
}

func stringContains(s, substr string) bool {
	return len(s) >= len(substr) && regexp.MustCompile(regexp.QuoteMeta(substr)).MatchString(s)
}
