package patcher

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var postInstallHooks []func(logger LogFunc)

// RegisterPostInstallHook registers a hook executed during wrapper installation
func RegisterPostInstallHook(h func(logger LogFunc)) {
	postInstallHooks = append(postInstallHooks, h)
}

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
		// 1. Dynamic detection for any existing language_server binary in binDir
		if detected := DetectLanguageServerInDir(binDir); detected != "" {
			binName = filepath.Base(detected)
		} else {
			// 2. Check if a backup (.bin) already exists in binDir
			entries, _ := os.ReadDir(binDir)
			for _, entry := range entries {
				if entry.IsDir() {
					continue
				}
				name := entry.Name()
				low := strings.ToLower(name)
				if strings.HasPrefix(low, "language_server") && strings.HasSuffix(low, ".bin") {
					binName = strings.TrimSuffix(name, ".bin")
					break
				}
			}

			// 3. Search for architecture-aware binary names
			if binName == "" {
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
			}

			// 4. Fallback default
			if binName == "" {
				if runtime.GOOS == "windows" {
					binName = "language_server.exe"
				} else {
					binName = "language_server"
				}
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
			return FormatAppManagementError("backing up server binary for wrapper", originalBin, err)
		}
	}

	if runtime.GOOS == "darwin" {
		_ = ResignDarwinBinary(originalBin)
	}

	// 2. Build flags
	flags := BuildServerFlags(opts, m.logger)
	flagsStr := strings.Join(flags, " ")

	// Execute post-install hooks (e.g. bi-directional token bridge synchronization)
	for _, hook := range postInstallHooks {
		hook(m.logger)
	}

	// 3. Create wrapper script with remote environment spoofing (switches language_server to file-token mode & reduces cloud fraud friction)
	backupName := BackupFileName(m.binName)
	var scriptContent string
	if runtime.GOOS == "windows" {
		scriptContent = fmt.Sprintf("@echo off\r\nset \"SSH_CONNECTION=127.0.0.1 1234 127.0.0.1 22\"\r\nset \"SSH_CLIENT=127.0.0.1 1234 22\"\r\n\"%%~dp0%s\" %s %%*\r\n", backupName, flagsStr)
	} else {
		scriptContent = fmt.Sprintf("#!/usr/bin/env bash\nDIR=\"$(cd \"$(dirname \"${BASH_SOURCE[0]}\")\" && pwd)\"\nexport SSH_CONNECTION=\"127.0.0.1 1234 127.0.0.1 22\"\nexport SSH_CLIENT=\"127.0.0.1 1234 22\"\nexec \"$DIR/%s\" %s \"$@\"\n", backupName, flagsStr)
	}

	if err := os.WriteFile(serverBin, []byte(scriptContent), 0755); err != nil {
		return FormatAppManagementError("writing wrapper script", serverBin, err)
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

	if err := os.Remove(serverBin); err != nil && !os.IsNotExist(err) {
		return FormatAppManagementError("removing wrapper script", serverBin, err)
	}
	if err := os.Rename(originalBin, serverBin); err != nil {
		return FormatAppManagementError("restoring original binary from wrapper", serverBin, err)
	}

	m.logger("Successfully removed wrapper and restored original binary at %s", serverBin)
	return nil
}
