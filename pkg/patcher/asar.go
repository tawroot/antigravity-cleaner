package patcher

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

var (
	// asarConsoleTarget: Original console transport configuration in electron-log
	// "format: DEFAULT_FORMAT,\n    level: 'silly'," (39 bytes)
	asarConsoleTarget = []byte("format: DEFAULT_FORMAT,\n    level: 'silly',")
	// asarConsolePatched: Zero-Log configuration with exact same length (39 bytes)
	asarConsolePatched = []byte("format: DEFAULT_FORMAT,\n    level: false,  ")

	// asarWriteFnTarget: Original unshielded writeFn (147 bytes)
	asarWriteFnTarget = []byte("    writeFn({ message }) {\n" +
		"      const consoleLogFn = consoleMethods[message.level] || consoleMethods.info;\n" +
		"      consoleLogFn(...message.data);\n" +
		"    },")

	// asarWriteFnPatched: Zero-Log shielded writeFn wrapped in try/catch with exact same length (151 bytes)
	asarWriteFnPatched = []byte("    writeFn({ message }) {\n" +
		"      try { (consoleMethods[message.level] || consoleMethods.info)(...message.data); } catch (_) {}\n" +
		"                      " + "},")

	// asarUpdaterTarget: electron-updater initialization in dist/updater.js (127 bytes)
	asarUpdaterTarget = []byte("function initAutoUpdater(isHeadless, settingsService) {\n    // In dev mode (npm start), electron-updater skips checks because the app")
	// asarUpdaterPatched: Anti-Update shield neutralizing silent background updates with exact same length (127 bytes)
	asarUpdaterPatched = []byte("function initAutoUpdater(isHeadless, settingsService) {\n    return; // In dev mode (npm start), electron-updater skips checks because")
)

// FindAntigravityAsar searches for app.asar in standard Antigravity installation paths
func FindAntigravityAsar() string {
	home, _ := os.UserHomeDir()

	candidates := []string{
		filepath.Join(home, "Applications", "Antigravity", "resources", "app.asar"),
		filepath.Join(home, "Applications", "antigravity", "resources", "app.asar"),
		filepath.Join(home, "Applications", "Antigravity-x64", "resources", "app.asar"),
		filepath.Join(home, "antigravity", "resources", "app.asar"),
		"/usr/share/antigravity-ide/resources/app.asar",
		"/usr/share/antigravity/resources/app.asar",
		"/opt/antigravity/resources/app.asar",
		"/opt/Antigravity/resources/app.asar",
		"/opt/antigravity-ide/resources/app.asar",
		"/opt/Antigravity IDE/resources/app.asar",
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Antigravity", "resources", "app.asar"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Antigravity IDE", "resources", "app.asar"),
		"/Applications/Antigravity.app/Contents/Resources/app.asar",
		"/Applications/Antigravity IDE.app/Contents/Resources/app.asar",
	}

	for _, c := range candidates {
		if fi, err := os.Stat(c); err == nil && !fi.IsDir() {
			return c
		}
	}
	return ""
}

// PatchAsarZeroLog applies the Zero-Log shield to app.asar, silencing console logging and broken-pipe EIO popups
func PatchAsarZeroLog(asarPath string) (string, error) {
	if asarPath == "" {
		asarPath = FindAntigravityAsar()
		if asarPath == "" {
			return "", fmt.Errorf("Antigravity app.asar not found")
		}
	}

	data, err := os.ReadFile(asarPath)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", asarPath, err)
	}

	// Check if already patched
	if bytes.Contains(data, asarConsolePatched) {
		return fmt.Sprintf("Already patched: %s (Zero-Log active)", filepath.Base(asarPath)), nil
	}

	idxConsole := bytes.Index(data, asarConsoleTarget)
	if idxConsole == -1 {
		// If level is already false or modified
		if bytes.Contains(data, []byte("level: false")) {
			return fmt.Sprintf("Already patched: %s (Zero-Log active)", filepath.Base(asarPath)), nil
		}
		return "", fmt.Errorf("console transport target signature not found in %s", filepath.Base(asarPath))
	}

	// Create backup
	if _, err := BackupFile(asarPath); err != nil {
		return "", fmt.Errorf("backup failed: %w", err)
	}

	patched := make([]byte, len(data))
	copy(patched, data)

	// 1. Patch console level: 'silly' -> level: false
	copy(patched[idxConsole:idxConsole+len(asarConsolePatched)], asarConsolePatched)

	// 2. Patch writeFn if present
	idxWrite := bytes.Index(patched, asarWriteFnTarget)
	if idxWrite != -1 && len(asarWriteFnTarget) == len(asarWriteFnPatched) {
		copy(patched[idxWrite:idxWrite+len(asarWriteFnPatched)], asarWriteFnPatched)
	}

	// 3. Patch initAutoUpdater if present (Anti-Update Shield)
	idxUpdater := bytes.Index(patched, asarUpdaterTarget)
	if idxUpdater != -1 && len(asarUpdaterTarget) == len(asarUpdaterPatched) {
		copy(patched[idxUpdater:idxUpdater+len(asarUpdaterPatched)], asarUpdaterPatched)
	}

	if err := AtomicWrite(asarPath, patched); err != nil {
		return "", fmt.Errorf("failed to write patched app.asar: %w", err)
	}

	// Also neutralize app-update.yml alongside app.asar if present
	neutralizeAppUpdateYml(asarPath)

	ClearIdeCache()

	return fmt.Sprintf("Successfully applied Zero-Log Shield & Anti-Update Shield to %s", filepath.Base(asarPath)), nil
}

// NeutralizeAppUpdateYml neutralizes app-update.yml to prevent Google background updates
func neutralizeAppUpdateYml(asarPath string) {
	ymlPath := filepath.Join(filepath.Dir(asarPath), "app-update.yml")
	if fi, err := os.Stat(ymlPath); err == nil && !fi.IsDir() {
		data, err := os.ReadFile(ymlPath)
		if err == nil && !bytes.Contains(data, []byte("127.0.0.1:0")) {
			_, _ = BackupFile(ymlPath)
			neutralized := "provider: generic\nurl: http://127.0.0.1:0/\nupdaterCacheDirName: antigravity-updater\n"
			_ = AtomicWrite(ymlPath, []byte(neutralized))
		}
	}
}

// PatchAsarAntiUpdate explicitly freezes background updates by patching app.asar and app-update.yml
func PatchAsarAntiUpdate(asarPath string) (string, error) {
	if asarPath == "" {
		asarPath = FindAntigravityAsar()
		if asarPath == "" {
			return "", fmt.Errorf("Antigravity app.asar not found")
		}
	}

	data, err := os.ReadFile(asarPath)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", asarPath, err)
	}

	idxUpdater := bytes.Index(data, asarUpdaterTarget)
	if idxUpdater == -1 {
		if bytes.Contains(data, asarUpdaterPatched) {
			neutralizeAppUpdateYml(asarPath)
			return fmt.Sprintf("Already patched: %s (Anti-Update Shield active)", filepath.Base(asarPath)), nil
		}
		neutralizeAppUpdateYml(asarPath)
		return fmt.Sprintf("Note: initAutoUpdater target signature not found in %s, neutralized app-update.yml", filepath.Base(asarPath)), nil
	}

	if _, err := BackupFile(asarPath); err != nil {
		return "", fmt.Errorf("backup failed: %w", err)
	}

	patched := make([]byte, len(data))
	copy(patched, data)
	copy(patched[idxUpdater:idxUpdater+len(asarUpdaterPatched)], asarUpdaterPatched)

	if err := AtomicWrite(asarPath, patched); err != nil {
		return "", fmt.Errorf("failed to write patched app.asar: %w", err)
	}

	neutralizeAppUpdateYml(asarPath)
	ClearIdeCache()

	return fmt.Sprintf("Successfully applied Anti-Update Shield to %s (background updates permanently frozen)", filepath.Base(asarPath)), nil
}
