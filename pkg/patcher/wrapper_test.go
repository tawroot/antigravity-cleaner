package patcher

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWrapperManagerInstallAndRemove(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "wrapper-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	fakeBinaryPath := filepath.Join(tempDir, "language_server")
	if err := os.WriteFile(fakeBinaryPath, []byte("#!/bin/sh\necho real"), 0755); err != nil {
		t.Fatalf("failed to create fake binary: %v", err)
	}

	mgr := NewWrapperManager(tempDir, func(format string, args ...interface{}) {})

	if mgr.IsWrapperInstalled() {
		t.Errorf("expected wrapper not installed initially")
	}

	opts := DefaultEnterpriseFlagsOptions()
	if err := mgr.InstallWrapper(opts); err != nil {
		t.Fatalf("failed to install wrapper: %v", err)
	}

	if !mgr.IsWrapperInstalled() {
		t.Errorf("expected wrapper to be installed")
	}

	// Verify content of installed wrapper
	content, err := os.ReadFile(fakeBinaryPath)
	if err != nil {
		t.Fatalf("failed to read wrapper: %v", err)
	}
	if len(content) == 0 {
		t.Errorf("wrapper file is empty")
	}

	// Remove wrapper
	if err := mgr.RemoveWrapper(); err != nil {
		t.Fatalf("failed to remove wrapper: %v", err)
	}

	if mgr.IsWrapperInstalled() {
		t.Errorf("expected wrapper to be uninstalled")
	}

	// Verify original content restored
	restored, err := os.ReadFile(fakeBinaryPath)
	if err != nil {
		t.Fatalf("failed to read restored file: %v", err)
	}
	if string(restored) != "#!/bin/sh\necho real" {
		t.Errorf("restored content mismatch: %s", string(restored))
	}
}

func TestWrapperManagerCustomBinaryName(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "wrapper-custom-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	customName := "language_server_linux_x64"
	fakeBinaryPath := filepath.Join(tempDir, customName)
	originalContent := "#!/bin/sh\necho custom-real"
	if err := os.WriteFile(fakeBinaryPath, []byte(originalContent), 0755); err != nil {
		t.Fatalf("failed to create fake custom binary: %v", err)
	}

	// 1. Pass full binary file path
	mgr := NewWrapperManager(fakeBinaryPath, func(format string, args ...interface{}) {})

	if mgr.GetBinaryName() != customName {
		t.Errorf("expected binary name %q, got %q", customName, mgr.GetBinaryName())
	}
	if mgr.GetBinaryDir() != tempDir {
		t.Errorf("expected binary dir %q, got %q", tempDir, mgr.GetBinaryDir())
	}

	if mgr.IsWrapperInstalled() {
		t.Errorf("expected wrapper not installed initially")
	}

	opts := DefaultEnterpriseFlagsOptions()
	if err := mgr.InstallWrapper(opts); err != nil {
		t.Fatalf("failed to install wrapper: %v", err)
	}

	if !mgr.IsWrapperInstalled() {
		t.Errorf("expected wrapper to be installed")
	}

	// Verify backup file exists with correct name
	backupFile := filepath.Join(tempDir, BackupFileName(customName))
	backupContent, err := os.ReadFile(backupFile)
	if err != nil {
		t.Fatalf("failed to read backup file %s: %v", backupFile, err)
	}
	if string(backupContent) != originalContent {
		t.Errorf("backup content mismatch: got %q, want %q", string(backupContent), originalContent)
	}

	// Verify wrapper script content references backup file
	wrapperContent, err := os.ReadFile(fakeBinaryPath)
	if err != nil {
		t.Fatalf("failed to read wrapper script: %v", err)
	}
	if !strings.Contains(string(wrapperContent), BackupFileName(customName)) {
		t.Errorf("wrapper script does not reference backup name %s: %s", BackupFileName(customName), string(wrapperContent))
	}

	// 2. Test directory-based initialization when wrapper is already installed
	dirMgr := NewWrapperManager(tempDir, func(format string, args ...interface{}) {})
	if dirMgr.GetBinaryName() != customName {
		t.Errorf("expected dirMgr to detect %q from backup, got %q", customName, dirMgr.GetBinaryName())
	}
	if !dirMgr.IsWrapperInstalled() {
		t.Errorf("expected dirMgr to recognize wrapper is installed")
	}

	// 3. Remove wrapper via dirMgr
	if err := dirMgr.RemoveWrapper(); err != nil {
		t.Fatalf("failed to remove wrapper via dirMgr: %v", err)
	}

	if dirMgr.IsWrapperInstalled() {
		t.Errorf("expected wrapper to be uninstalled")
	}

	// Verify original file restored
	restored, err := os.ReadFile(fakeBinaryPath)
	if err != nil {
		t.Fatalf("failed to read restored binary: %v", err)
	}
	if string(restored) != originalContent {
		t.Errorf("restored content mismatch: got %q, want %q", string(restored), originalContent)
	}
}

func TestWrapperManagerMacArmBinary(t *testing.T) {
	tempDir := t.TempDir()

	// Create language_server_macos_arm
	macArmName := "language_server_macos_arm"
	macArmPath := filepath.Join(tempDir, macArmName)
	originalContent := "#!/bin/sh\necho mac-arm-binary"
	if err := os.WriteFile(macArmPath, []byte(originalContent), 0755); err != nil {
		t.Fatalf("failed to create fake mac arm binary: %v", err)
	}

	// Initialize WrapperManager using directory path only
	mgr := NewWrapperManager(tempDir, func(format string, args ...interface{}) {})

	if mgr.GetBinaryName() != macArmName {
		t.Errorf("expected binary name %q, got %q", macArmName, mgr.GetBinaryName())
	}

	opts := DefaultEnterpriseFlagsOptions()
	if err := mgr.InstallWrapper(opts); err != nil {
		t.Fatalf("failed to install wrapper on mac arm binary: %v", err)
	}

	if !mgr.IsWrapperInstalled() {
		t.Errorf("expected wrapper to be installed")
	}

	if err := mgr.RemoveWrapper(); err != nil {
		t.Fatalf("failed to remove wrapper: %v", err)
	}

	restored, err := os.ReadFile(macArmPath)
	if err != nil {
		t.Fatalf("failed to read restored binary: %v", err)
	}
	if string(restored) != originalContent {
		t.Errorf("restored content mismatch: got %q, want %q", string(restored), originalContent)
	}
}
