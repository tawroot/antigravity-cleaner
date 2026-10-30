package patcher

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

// IsAppManagementError checks if an error was caused by macOS App Management (EPERM / operation not permitted)
func IsAppManagementError(err error) bool {
	if err == nil {
		return false
	}
	if runtime.GOOS != "darwin" {
		return false
	}
	if errors.Is(err, syscall.EPERM) {
		return true
	}
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "operation not permitted")
}

// FormatAppManagementError wraps an error with actionable guidance for macOS App Management consent
func FormatAppManagementError(operation, targetPath string, err error) error {
	if !IsAppManagementError(err) {
		return err
	}
	return fmt.Errorf("%w\n\n"+
		"🍎 [macOS App Management Protection Blocked This Action]\n"+
		"macOS prevented %s at: %s\n"+
		"To resolve this:\n"+
		"  1. Open System Settings → Privacy & Security → App Management\n"+
		"  2. Allow your Terminal (or the application running Antigravity Cleaner)\n"+
		"  (Quick command: open \"x-apple.systempreferences:com.apple.preference.security?Privacy_AppManagement\")",
		err, operation, targetPath)
}

// CheckAppBundleWriteability verifies that the target application or file path is writeable on macOS
func CheckAppBundleWriteability(targetPath string) (string, error) {
	if targetPath == "" {
		return "", fmt.Errorf("target path is empty")
	}

	testDir := targetPath
	if fi, err := os.Stat(targetPath); err == nil && !fi.IsDir() {
		testDir = filepath.Dir(targetPath)
	}

	tmpFile, err := os.CreateTemp(testDir, ".ag_canary_*")
	if err != nil {
		return "", FormatAppManagementError("bundle write test", testDir, err)
	}
	tmpName := tmpFile.Name()
	_ = tmpFile.Close()
	_ = os.Remove(tmpName)

	return "Write permissions verified", nil
}

// AtomicWrite replaces targetPath with data atomically, preserving file permissions.
func AtomicWrite(targetPath string, data []byte) error {
	dir := filepath.Dir(targetPath)
	info, err := os.Stat(targetPath)
	var perm os.FileMode = 0755
	if err == nil {
		perm = info.Mode().Perm()
	}

	// Create temp file in same directory to ensure it is on the same filesystem
	tmpFile, err := os.CreateTemp(dir, ".ag_patch_*")
	if err != nil {
		return FormatAppManagementError("creating temporary file for atomic write", dir, err)
	}
	tmpName := tmpFile.Name()
	defer os.Remove(tmpName) // Cleanup on failure

	if _, err := tmpFile.Write(data); err != nil {
		tmpFile.Close()
		return fmt.Errorf("failed to write patched bytes: %w", err)
	}

	if err := tmpFile.Sync(); err != nil {
		tmpFile.Close()
		return fmt.Errorf("failed to sync temp file: %w", err)
	}

	if err := tmpFile.Chmod(perm); err != nil {
		// Non-fatal on some OSes
	}

	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}

	// On Windows, if destination exists and is running, rename target to .old first
	if runtime.GOOS == "windows" {
		bakPath := targetPath + ".old"
		_ = os.Remove(bakPath)
		_ = os.Rename(targetPath, bakPath)
	}

	// Atomic rename replaces the target inode
	if err := os.Rename(tmpName, targetPath); err != nil {
		return FormatAppManagementError("atomically replacing target file", targetPath, err)
	}

	// On macOS, if the written file is an executable Mach-O binary, clear quarantine and re-sign ad-hoc
	// to prevent AMFI (Apple Mobile File Integrity) kernel SIGTERM/SIGKILL termination.
	if runtime.GOOS == "darwin" && IsMachOBinary(data) {
		_ = ResignDarwinBinary(targetPath)
	}

	return nil
}

// IsMachOBinary checks if the data slice begins with Mach-O magic numbers (32-bit, 64-bit, or Universal fat binary)
func IsMachOBinary(data []byte) bool {
	if len(data) < 4 {
		return false
	}
	magic := binary.BigEndian.Uint32(data[:4])
	return magic == 0xfeedface || magic == 0xcefaedfe || // 32-bit Mach-O
		magic == 0xfeedfacf || magic == 0xcffaedfe || // 64-bit Mach-O
		magic == 0xcafebabe || magic == 0xbebafeca // Fat / Universal binary
}

// ResignDarwinBinary clears quarantine flags and applies an ad-hoc signature on macOS binaries.
// This prevents AMFI (Apple Mobile File Integrity) kernel crashes (SIGTERM/SIGKILL) on modified Mach-O files.
func ResignDarwinBinary(targetPath string) error {
	if runtime.GOOS != "darwin" {
		return nil
	}

	// 1. Remove Apple quarantine attribute
	_ = exec.Command("xattr", "-dr", "com.apple.quarantine", targetPath).Run()

	// 2. Force ad-hoc re-signature (codesign --force --sign - <target>)
	cmd := exec.Command("codesign", "--force", "--sign", "-", targetPath)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("codesign failed on %s: %s (%w)", targetPath, strings.TrimSpace(string(out)), err)
	}
	return nil
}

// BackupFile creates a .agybak backup if one does not already exist.
func BackupFile(filePath string) (string, error) {
	bakPath := filePath + ".agybak"
	if _, err := os.Stat(bakPath); err == nil {
		// Backup already exists, keep original
		return bakPath, nil
	}

	src, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("cannot open original file for backup: %w", err)
	}
	defer src.Close()

	info, err := src.Stat()
	if err != nil {
		return "", err
	}

	dst, err := os.OpenFile(bakPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode().Perm())
	if err != nil {
		return "", FormatAppManagementError("creating backup file", bakPath, err)
	}

	if _, err := io.Copy(dst, src); err != nil {
		_ = dst.Close()
		return "", fmt.Errorf("failed to copy bytes to backup: %w", err)
	}

	if err := dst.Sync(); err != nil {
		_ = dst.Close()
		return "", fmt.Errorf("failed to flush backup file: %w", err)
	}

	if err := dst.Close(); err != nil {
		return "", fmt.Errorf("failed to close backup file: %w", err)
	}

	return bakPath, nil
}

// RestoreFile restores filePath from filePath.agybak.
func RestoreFile(filePath string) error {
	bakPath := filePath + ".agybak"
	if _, err := os.Stat(bakPath); err != nil {
		return fmt.Errorf("no backup found at %s", bakPath)
	}

	bakData, err := os.ReadFile(bakPath)
	if err != nil {
		return FormatAppManagementError("reading backup file", bakPath, err)
	}

	return AtomicWrite(filePath, bakData)
}
