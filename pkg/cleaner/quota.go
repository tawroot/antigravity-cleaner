package cleaner

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// AntigravityConfigPaths returns potential Antigravity user data directories
func AntigravityConfigPaths() []string {
	home, _ := os.UserHomeDir()
	return []string{
		filepath.Join(home, ".config", "Antigravity"),
		filepath.Join(home, ".config", "Antigravity IDE"),
		filepath.Join(os.Getenv("APPDATA"), "Antigravity"),
		filepath.Join(os.Getenv("APPDATA"), "Antigravity IDE"),
		filepath.Join(home, "Library", "Application Support", "Antigravity"),
		filepath.Join(home, "Library", "Application Support", "Antigravity IDE"),
	}
}

// SurgicalReset429 clears ONLY corrupted token, quota and auth state files,
// preserving all user chat histories, settings, and workspace data.
func SurgicalReset429(configDir string, dryRun bool) ([]string, error) {
	if configDir == "" {
		for _, p := range AntigravityConfigPaths() {
			if fi, err := os.Stat(p); err == nil && fi.IsDir() {
				configDir = p
				break
			}
		}
	}

	if configDir == "" {
		return nil, fmt.Errorf("no Antigravity configuration directory found")
	}

	// Targeted list of files responsible for 429 Quota & corrupt tokens
	targetFiles := []string{
		"Cookies",
		"Cookies-journal",
		"DIPS",
		"DIPS-wal",
		"TransportSecurity",
		"Trust Tokens",
		"Trust Tokens-journal",
		"SharedStorage",
		"SharedStorage-wal",
		"Network Persistent State",
		"app_storage.json",
	}

	targetDirs := []string{
		"blob_storage",
		"DawnGraphiteCache",
		"DawnWebGPUCache",
		"GPUCache",
		"Session Storage",
	}

	var removed []string

	// Remove target files
	for _, f := range targetFiles {
		fullPath := filepath.Join(configDir, f)
		if _, err := os.Stat(fullPath); err == nil {
			removed = append(removed, f)
			if !dryRun {
				_ = os.Remove(fullPath)
			}
		}
	}

	// Remove target directories
	for _, d := range targetDirs {
		fullPath := filepath.Join(configDir, d)
		if _, err := os.Stat(fullPath); err == nil {
			removed = append(removed, d+"/")
			if !dryRun {
				_ = os.RemoveAll(fullPath)
			}
		}
	}

	return removed, nil
}

// BackupUserData creates a timestamped zip or folder backup of the User directory
func BackupUserData(configDir string) (string, error) {
	userDir := filepath.Join(configDir, "User")
	if _, err := os.Stat(userDir); err != nil {
		return "", fmt.Errorf("User directory not found in %s", configDir)
	}

	home, _ := os.UserHomeDir()
	backupDir := filepath.Join(home, "Antigravity-Backups", time.Now().Format("2006-01-02_15-04-05"))
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return "", err
	}

	// Copy User directory to backupDir
	err := copyDir(userDir, filepath.Join(backupDir, "User"))
	if err != nil {
		return "", err
	}

	return backupDir, nil
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, info.Mode())
		}
		return copyFile(path, target)
	})
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
