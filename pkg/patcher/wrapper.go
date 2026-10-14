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
	binDir string
	logger LogFunc
}

// NewWrapperManager creates a new WrapperManager for a given binary directory
func NewWrapperManager(binDir string, logger LogFunc) *WrapperManager {
	if logger == nil {
		logger = defaultFlagsLogger
	}
	return &WrapperManager{
		binDir: binDir,
		logger: logger,
	}
}

// IsWrapperInstalled checks whether a wrapper script is active and original binary backed up
func (m *WrapperManager) IsWrapperInstalled() bool {
	serverBin := filepath.Join(m.binDir, "language_server")
	originalBin := filepath.Join(m.binDir, "language_server.bin")

	if runtime.GOOS == "windows" {
		serverBin = filepath.Join(m.binDir, "language_server.exe")
		originalBin = filepath.Join(m.binDir, "language_server.bin.exe")
	}

	origInfo, err := os.Stat(originalBin)
	if err != nil || origInfo.IsDir() {
		return false
	}

	serverInfo, err := os.Stat(serverBin)
	if err != nil || serverInfo.IsDir() {
		return false
	}

	return true
}

// InstallWrapper backs up the original binary and installs the flag-forwarding wrapper
func (m *WrapperManager) InstallWrapper(opts EnterpriseFlagsOptions) error {
	serverBin := filepath.Join(m.binDir, "language_server")
	originalBin := filepath.Join(m.binDir, "language_server.bin")

	if runtime.GOOS == "windows" {
		serverBin = filepath.Join(m.binDir, "language_server.exe")
		originalBin = filepath.Join(m.binDir, "language_server.bin.exe")
	}

	// 1. If original backup doesn't exist, backup the real binary
	if _, err := os.Stat(originalBin); os.IsNotExist(err) {
		if _, err := os.Stat(serverBin); err != nil {
			return fmt.Errorf("language_server not found at %s: %w", serverBin, err)
		}
		m.logger("Backing up original binary to %s", originalBin)
		if err := os.Rename(serverBin, originalBin); err != nil {
			return fmt.Errorf("failed to backup language_server: %w", err)
		}
	}

	// 2. Build flags
	flags := BuildServerFlags(opts, m.logger)
	flagsStr := strings.Join(flags, " ")

	// 3. Create wrapper script
	var scriptContent string
	if runtime.GOOS == "windows" {
		scriptContent = fmt.Sprintf("@echo off\r\n\"%%~dp0language_server.bin.exe\" %s %%*\r\n", flagsStr)
	} else {
		scriptContent = fmt.Sprintf("#!/usr/bin/env bash\nDIR=\"$(cd \"$(dirname \"${BASH_SOURCE[0]}\")\" && pwd)\"\nexec \"$DIR/language_server.bin\" %s \"$@\"\n", flagsStr)
	}

	if err := os.WriteFile(serverBin, []byte(scriptContent), 0755); err != nil {
		return fmt.Errorf("failed to write wrapper script: %w", err)
	}

	m.logger("Successfully installed enterprise wrapper at %s", serverBin)
	return nil
}

// RemoveWrapper restores the original binary and removes the wrapper
func (m *WrapperManager) RemoveWrapper() error {
	serverBin := filepath.Join(m.binDir, "language_server")
	originalBin := filepath.Join(m.binDir, "language_server.bin")

	if runtime.GOOS == "windows" {
		serverBin = filepath.Join(m.binDir, "language_server.exe")
		originalBin = filepath.Join(m.binDir, "language_server.bin.exe")
	}

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
