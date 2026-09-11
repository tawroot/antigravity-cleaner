package patcher

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

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
		return fmt.Errorf("failed to create temp file in %s: %w", dir, err)
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
		return fmt.Errorf("failed to atomically rename %s to %s: %w", tmpName, targetPath, err)
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
		return "", fmt.Errorf("cannot create backup file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to copy bytes to backup: %w", err)
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
		return fmt.Errorf("failed to read backup: %w", err)
	}

	return AtomicWrite(filePath, bakData)
}
