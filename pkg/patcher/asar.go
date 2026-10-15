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

	if err := AtomicWrite(asarPath, patched); err != nil {
		return "", fmt.Errorf("failed to write patched app.asar: %w", err)
	}

	ClearIdeCache()

	return fmt.Sprintf("Successfully applied Zero-Log Shield to %s (silenced broken-pipe write EIO)", filepath.Base(asarPath)), nil
}
