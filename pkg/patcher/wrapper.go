package patcher

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// WrapperManager manages transparent flag injection wrappers around language_server
type WrapperManager struct {
	binDir  string
	binName string
	logger  LogFunc
}

// BackupFileName returns the backup file name for a given binary name
func BackupFileName(binName string) string {
	if runtime.GOOS == "windows" {
		if strings.HasSuffix(strings.ToLower(binName), ".exe") {
			return strings.TrimSuffix(binName, filepath.Ext(binName)) + ".bin.exe"
		}
		return binName + ".bin.exe"
	}
	return binName + ".bin"
}

// NewWrapperManager creates a new WrapperManager for a given binary directory or executable path
func NewWrapperManager(pathOrDir string, logger LogFunc) *WrapperManager {
	if logger == nil {
		logger = defaultFlagsLogger
	}

	binDir := pathOrDir
	binName := ""

	if fi, err := os.Stat(pathOrDir); err == nil && !fi.IsDir() {
		// pathOrDir is an existing file
		binDir = filepath.Dir(pathOrDir)
		binName = filepath.Base(pathOrDir)
	} else if filepath.Ext(pathOrDir) != "" && !strings.HasSuffix(pathOrDir, string(filepath.Separator)) {
		// Could be a file path that doesn't exist yet or is currently moved/backed up
		dir := filepath.Dir(pathOrDir)
		base := filepath.Base(pathOrDir)
		if fiDir, errDir := os.Stat(dir); errDir == nil && fiDir.IsDir() {
			binDir = dir
			binName = base
		}
	}

	// Normalize if binName was passed as a .bin backup
	if strings.HasSuffix(binName, ".bin.exe") {
		binName = strings.TrimSuffix(binName, ".bin.exe") + ".exe"
	} else if strings.HasSuffix(binName, ".bin") {
		binName = strings.TrimSuffix(binName, ".bin")
	}

	// If binName was not explicitly identified from a file path:
	if binName == "" {
		// Search for GetLanguageServerBinaryNames() in binDir
		for _, name := range GetLanguageServerBinaryNames() {
			target := filepath.Join(binDir, name)
			if fi, err := os.Stat(target); err == nil && !fi.IsDir() {
				binName = name
				break
			}
			backup := filepath.Join(binDir, BackupFileName(name))
			if fi, err := os.Stat(backup); err == nil && !fi.IsDir() {
				binName = name
				break
			}
		}

		// Fallback default
		if binName == "" {
			if runtime.GOOS == "windows" {
				binName = "language_server.exe"
			} else {
				binName = "language_server"
			}
		}
	}

	return &WrapperManager{
		binDir:  binDir,
		binName: binName,
		logger:  logger,
	}
}

// GetBinaryName returns the resolved binary name (e.g. language_server_linux_x64)
func (m *WrapperManager) GetBinaryName() string {
	return m.binName
}

// GetBinaryDir returns the directory containing the binary
func (m *WrapperManager) GetBinaryDir() string {
	return m.binDir
}

// ServerBin returns the full path of the server binary or wrapper
func (m *WrapperManager) ServerBin() string {
	return filepath.Join(m.binDir, m.binName)
}

// OriginalBin returns the full path of the backed-up original binary
func (m *WrapperManager) OriginalBin() string {
	return filepath.Join(m.binDir, BackupFileName(m.binName))
}

// IsWrapperInstalled checks whether a wrapper script is active and original binary backed up
func (m *WrapperManager) IsWrapperInstalled() bool {
	origInfo, err := os.Stat(m.OriginalBin())
	if err != nil || origInfo.IsDir() {
		return false
	}

	serverInfo, err := os.Stat(m.ServerBin())
	if err != nil || serverInfo.IsDir() {
		return false
	}

	return true
}

// InstallWrapper backs up the original binary and installs the flag-forwarding wrapper
func (m *WrapperManager) InstallWrapper(opts EnterpriseFlagsOptions) error {
	serverBin := m.ServerBin()
	originalBin := m.OriginalBin()

	// 1. If original backup doesn't exist, backup the real binary
	if _, err := os.Stat(originalBin); os.IsNotExist(err) {
		if _, err := os.Stat(serverBin); err != nil {
			return fmt.Errorf("%s not found at %s: %w", m.binName, serverBin, err)
		}
		m.logger("Backing up original binary to %s", originalBin)
		if err := os.Rename(serverBin, originalBin); err != nil {
			return fmt.Errorf("failed to backup %s: %w", m.binName, err)
		}
	}

	// 2. Build flags
	flags := BuildServerFlags(opts, m.logger)
	flagsStr := strings.Join(flags, " ")

	// 3. Create wrapper script
	backupName := BackupFileName(m.binName)
	var scriptContent string
	if runtime.GOOS == "windows" {
		scriptContent = fmt.Sprintf("@echo off\r\n\"%%~dp0%s\" %s %%*\r\n", backupName, flagsStr)
	} else {
		scriptContent = fmt.Sprintf("#!/usr/bin/env bash\nDIR=\"$(cd \"$(dirname \"${BASH_SOURCE[0]}\")\" && pwd)\"\nexec \"$DIR/%s\" %s \"$@\"\n", backupName, flagsStr)
	}

	if err := os.WriteFile(serverBin, []byte(scriptContent), 0755); err != nil {
		return fmt.Errorf("failed to write wrapper script: %w", err)
	}

	m.logger("Successfully installed enterprise wrapper at %s", serverBin)
	return nil
}

// RemoveWrapper restores the original binary and removes the wrapper
func (m *WrapperManager) RemoveWrapper() error {
	serverBin := m.ServerBin()
	originalBin := m.OriginalBin()

	if _, err := os.Stat(originalBin); os.IsNotExist(err) {
		m.logger("Original binary not found; wrapper is not currently installed")
		return nil
	}

	_ = os.Remove(serverBin)
	if err := os.Rename(originalBin, serverBin); err != nil {
		return fmt.Errorf("failed to restore original binary: %w", err)
	}

	m.logger("Successfully removed wrapper and restored original binary at %s", serverBin)
	return nil
}
