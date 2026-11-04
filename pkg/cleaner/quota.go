package cleaner

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

// AntigravityConfigPaths dynamically discovers all Antigravity user data directories
// across Linux, macOS, and Windows without relying on single hardcoded paths.
func AntigravityConfigPaths() []string {
	home, _ := os.UserHomeDir()
	var candidates []string
	seen := make(map[string]bool)

	add := func(p string) {
		if p == "" {
			return
		}
		clean := filepath.Clean(p)
		if !seen[clean] {
			if fi, err := os.Stat(clean); err == nil && fi.IsDir() {
				seen[clean] = true
				candidates = append(candidates, clean)
			}
		}
	}

	// 1. Root candidate parents across OS environments
	var parentDirs []string
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		parentDirs = append(parentDirs, xdg)
	}
	if home != "" {
		parentDirs = append(parentDirs,
			filepath.Join(home, ".config"),
			filepath.Join(home, ".gemini"),
			filepath.Join(home, "Library", "Application Support"),
		)
	}
	if appData := os.Getenv("APPDATA"); appData != "" {
		parentDirs = append(parentDirs, appData)
	}
	if localAppData := os.Getenv("LOCALAPPDATA"); localAppData != "" {
		parentDirs = append(parentDirs, localAppData)
	}

	// 2. Scan candidate parents dynamically for any Antigravity-related directories
	for _, parent := range parentDirs {
		entries, err := os.ReadDir(parent)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if entry.IsDir() {
				nameLower := strings.ToLower(entry.Name())
				if strings.Contains(nameLower, "antigravity") {
					add(filepath.Join(parent, entry.Name()))
				}
			}
		}
	}

	// 3. Fallback standard folder names in case directory enumeration is restricted
	for _, parent := range parentDirs {
		for _, name := range []string{"Antigravity", "Antigravity IDE", "antigravity", "antigravity-ide"} {
			add(filepath.Join(parent, name))
		}
	}

	return candidates
}

// SurgicalReset429 clears ONLY corrupted token, quota and auth state files,
// including modern Chromium Network/ directories and tainted OAuth tokens in state.vscdb,
// strictly preserving all user chat histories, settings, and workspace data.
func SurgicalReset429(configDir string, dryRun bool) ([]string, error) {
	var targetDirsToClean []string
	if configDir != "" {
		targetDirsToClean = append(targetDirsToClean, configDir)
	} else {
		targetDirsToClean = AntigravityConfigPaths()
	}

	if len(targetDirsToClean) == 0 {
		return nil, fmt.Errorf("no Antigravity configuration directory found")
	}

	// Targeted list of files responsible for 429 Quota & corrupt tokens
	// Includes both root files and modern Chromium Network/ subfolder paths.
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
		"Local State",
		"app_storage.json",
		filepath.Join("Network", "Cookies"),
		filepath.Join("Network", "Cookies-journal"),
		filepath.Join("Network", "Network Persistent State"),
		filepath.Join("Network", "NetworkDataMigrated"),
		filepath.Join("Network", "Trust Tokens"),
		filepath.Join("Network", "Trust Tokens-journal"),
		filepath.Join("Network", "TransportSecurity"),
	}

	targetDirs := []string{
		"blob_storage",
		"DawnGraphiteCache",
		"DawnWebGPUCache",
		"GPUCache",
		"Session Storage",
		"WebStorage",
		"Partitions",
	}

	var totalRemoved []string
	seenRemoved := make(map[string]bool)

	for _, cDir := range targetDirsToClean {
		// 1. Remove target files
		for _, f := range targetFiles {
			fullPath := filepath.Join(cDir, f)
			if _, err := os.Stat(fullPath); err == nil {
				display := filepath.Base(cDir) + "/" + filepath.ToSlash(f)
				if !seenRemoved[display] {
					seenRemoved[display] = true
					totalRemoved = append(totalRemoved, display)
				}
				if !dryRun {
					_ = os.Remove(fullPath)
				}
			}
		}

		// 2. Remove target directories
		for _, d := range targetDirs {
			fullPath := filepath.Join(cDir, d)
			if _, err := os.Stat(fullPath); err == nil {
				display := filepath.Base(cDir) + "/" + filepath.ToSlash(d) + "/"
				if !seenRemoved[display] {
					seenRemoved[display] = true
					totalRemoved = append(totalRemoved, display)
				}
				if !dryRun {
					_ = os.RemoveAll(fullPath)
				}
			}
		}

		// 3. Reset tainted OAuth tokens in state.vscdb while preserving 100% of chats and workspaces
		stateDbPath := filepath.Join(cDir, "User", "globalStorage", "state.vscdb")
		if _, err := os.Stat(stateDbPath); err == nil {
			if dryRun {
				display := filepath.Base(cDir) + "/User/globalStorage/state.vscdb (uss-oauth)"
				if !seenRemoved[display] {
					seenRemoved[display] = true
					totalRemoved = append(totalRemoved, display)
				}
			} else {
				if resetTaintedTokensInStateDb(stateDbPath) {
					display := filepath.Base(cDir) + "/User/globalStorage/state.vscdb (uss-oauth)"
					if !seenRemoved[display] {
						seenRemoved[display] = true
						totalRemoved = append(totalRemoved, display)
					}
				}
			}
		}
	}

	return totalRemoved, nil
}

// resetTaintedTokensInStateDb safely purges only the corrupted OAuth session keys from SQLite state.vscdb
// leaving all user chats, history, keybindings, and extensions completely intact.
func resetTaintedTokensInStateDb(dbPath string) bool {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return false
	}
	defer db.Close()

	keysToPurge := []string{
		"antigravityUnifiedStateSync.oauthToken",
		"antigravityUnifiedStateSync.userStatus",
	}

	purgedAny := false
	for _, k := range keysToPurge {
		res, err := db.Exec("DELETE FROM ItemTable WHERE key = ?", k)
		if err == nil {
			if affected, _ := res.RowsAffected(); affected > 0 {
				purgedAny = true
			}
		}
	}

	return purgedAny
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
