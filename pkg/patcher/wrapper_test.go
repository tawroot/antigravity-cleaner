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

func TestHealWindowsBinariesMissingServerBin(t *testing.T) {
	tempDir := t.TempDir()

	// Simulate Bug #40: language_server.exe is missing, but language_server.bin.exe exists
	fakePE := make([]byte, 150*1024)
	copy(fakePE, []byte("MZ-FAKE-PE-BINARY"))
	binBackup := filepath.Join(tempDir, "language_server.bin.exe")
	if err := os.WriteFile(binBackup, fakePE, 0755); err != nil {
		t.Fatalf("failed to write fake backup: %v", err)
	}

	HealWindowsBinaries(tempDir, nil)

	serverBin := filepath.Join(tempDir, "language_server.exe")
	info, err := os.Stat(serverBin)
	if err != nil {
		t.Fatalf("expected language_server.exe to be restored, got error: %v", err)
	}
	if info.Size() != int64(len(fakePE)) {
		t.Errorf("expected restored size %d, got %d", len(fakePE), info.Size())
	}

	// Stale .bin.exe should be cleaned up
	if _, err := os.Stat(binBackup); !os.IsNotExist(err) {
		t.Errorf("expected language_server.bin.exe to be cleaned up")
	}
}

func TestHealWindowsBinariesRecursiveBinLoop(t *testing.T) {
	tempDir := t.TempDir()

	// Simulate Bug #36 / repeated runs: language_server.bin.bin.bin.exe exists
	fakePE := make([]byte, 200*1024)
	copy(fakePE, []byte("MZ-NESTED-BIN-LOOP"))
	nestedBin := filepath.Join(tempDir, "language_server.bin.bin.bin.exe")
	if err := os.WriteFile(nestedBin, fakePE, 0755); err != nil {
		t.Fatalf("failed to write nested bin: %v", err)
	}

	HealWindowsBinaries(tempDir, nil)

	serverBin := filepath.Join(tempDir, "language_server.exe")
	info, err := os.Stat(serverBin)
	if err != nil {
		t.Fatalf("expected language_server.exe to be restored from nested bin, got error: %v", err)
	}
	if info.Size() != int64(len(fakePE)) {
		t.Errorf("expected restored size %d, got %d", len(fakePE), info.Size())
	}

	if _, err := os.Stat(nestedBin); !os.IsNotExist(err) {
		t.Errorf("expected nested bin to be cleaned up")
	}
}

func TestHealWindowsBinariesLegacyScriptWrapper(t *testing.T) {
	tempDir := t.TempDir()

	// Simulate legacy batch script wrapper in language_server.exe and real PE in .bin.exe
	batchContent := []byte("@echo off\r\nstart language_server.bin.exe %*")
	serverBin := filepath.Join(tempDir, "language_server.exe")
	if err := os.WriteFile(serverBin, batchContent, 0755); err != nil {
		t.Fatalf("failed to write batch script: %v", err)
	}

	fakePE := make([]byte, 180*1024)
	copy(fakePE, []byte("MZ-REAL-PE"))
	binBackup := filepath.Join(tempDir, "language_server.bin.exe")
	if err := os.WriteFile(binBackup, fakePE, 0755); err != nil {
		t.Fatalf("failed to write fake backup: %v", err)
	}

	HealWindowsBinaries(tempDir, nil)

	info, err := os.Stat(serverBin)
	if err != nil {
		t.Fatalf("expected serverBin to exist, got: %v", err)
	}
	if info.Size() != int64(len(fakePE)) {
		t.Errorf("expected PE binary restored, got size %d", info.Size())
	}
}

func TestHealDarwinBinariesLegacyScriptWrapper(t *testing.T) {
	tempDir := t.TempDir()

	// Simulate Bug #37: bash script wrapper in language_server and real Mach-O in language_server.bin
	scriptContent := []byte("#!/usr/bin/env bash\nexec \"$DIR/language_server.bin\" \"$@\"\n")
	serverBin := filepath.Join(tempDir, "language_server")
	if err := os.WriteFile(serverBin, scriptContent, 0755); err != nil {
		t.Fatalf("failed to write script wrapper: %v", err)
	}

	fakeMachO := make([]byte, 180*1024)
	// Mach-O 64-bit LE magic: 0xcffaedfe
	fakeMachO[0] = 0xcf
	fakeMachO[1] = 0xfa
	fakeMachO[2] = 0xed
	fakeMachO[3] = 0xfe
	binBackup := filepath.Join(tempDir, "language_server.bin")
	if err := os.WriteFile(binBackup, fakeMachO, 0755); err != nil {
		t.Fatalf("failed to write fake Mach-O backup: %v", err)
	}

	HealDarwinBinaries(tempDir, nil)

	info, err := os.Stat(serverBin)
	if err != nil {
		t.Fatalf("expected serverBin to exist after healing, got: %v", err)
	}
	if info.Size() != int64(len(fakeMachO)) {
		t.Errorf("expected Mach-O binary restored, got size %d, want %d", info.Size(), len(fakeMachO))
	}

	// Backup .bin should be cleaned up
	if _, err := os.Stat(binBackup); !os.IsNotExist(err) {
		t.Errorf("expected language_server.bin to be cleaned up after restore")
	}
}

func TestHealDarwinBinariesMissingServerBin(t *testing.T) {
	tempDir := t.TempDir()

	fakeMachO := make([]byte, 150*1024)
	fakeMachO[0] = 0xcf
	fakeMachO[1] = 0xfa
	fakeMachO[2] = 0xed
	fakeMachO[3] = 0xfe
	binBackup := filepath.Join(tempDir, "language_server.bin")
	if err := os.WriteFile(binBackup, fakeMachO, 0755); err != nil {
		t.Fatalf("failed to write fake backup: %v", err)
	}

	HealDarwinBinaries(tempDir, nil)

	serverBin := filepath.Join(tempDir, "language_server")
	info, err := os.Stat(serverBin)
	if err != nil {
		t.Fatalf("expected language_server to be restored, got error: %v", err)
	}
	if info.Size() != int64(len(fakeMachO)) {
		t.Errorf("expected restored size %d, got %d", len(fakeMachO), info.Size())
	}

	if _, err := os.Stat(binBackup); !os.IsNotExist(err) {
		t.Errorf("expected language_server.bin to be cleaned up")
	}
}
