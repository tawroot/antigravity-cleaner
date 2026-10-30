package account

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/tawroot/antigravity-cleaner/pkg/quota"
	"github.com/zalando/go-keyring"
	_ "modernc.org/sqlite"
)

var ideKeys = []string{
	"antigravityUnifiedStateSync.oauthToken",
	"antigravityUnifiedStateSync.userStatus",
	"antigravityUnifiedStateSync.profileUrl",
	"antigravityUnifiedStateSync.modelCredits",
}

// SavedAccount represents a saved Google login profile
type SavedAccount struct {
	Name          string            `json:"name"`
	SavedAt       string            `json:"saved_at"`
	Email         string            `json:"email,omitempty"`
	Source        string            `json:"source"` // "hub" or "ide"
	KeyringSecret string            `json:"keyring_secret,omitempty"`
	StateItems    map[string]string `json:"state_items,omitempty"`
}

// AccountSummary holds metadata for listing accounts
type AccountSummary struct {
	Name    string `json:"name"`
	SavedAt string `json:"saved_at"`
	Email   string `json:"email"`
	Source  string `json:"source"`
}

// GetStorageDir returns the directory where account profiles are stored
func GetStorageDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".config", "antigravity-cleaner", "accounts")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", err
	}
	return dir, nil
}

// FindAntigravityHubDir returns the standard Antigravity standalone hub configuration directory
func FindAntigravityHubDir() string {
	home, _ := os.UserHomeDir()
	var candidates []string

	if runtime.GOOS == "windows" {
		appData := os.Getenv("APPDATA")
		candidates = append(candidates, filepath.Join(appData, "Antigravity"))
	} else if runtime.GOOS == "darwin" {
		candidates = append(candidates, filepath.Join(home, "Library", "Application Support", "Antigravity"))
	} else {
		// Linux
		candidates = append(candidates, filepath.Join(home, ".config", "Antigravity"))
	}

	for _, c := range candidates {
		if fi, err := os.Stat(c); err == nil && fi.IsDir() {
			return c
		}
	}
	return ""
}

var emailRe = regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

var hubFiles = []string{
	"app_storage.json",
	"Cookies",
	"Cookies-journal",
	"Preferences",
	"Local State",
	"Network Persistent State",
	"DIPS",
	"DIPS-wal",
	"Trust Tokens",
	"Trust Tokens-journal",
}

var hubDirs = []string{
	"Local Storage",
	"Session Storage",
	"Partitions",
}

var hubCacheExcludes = []string{
	"Cache",
	"Code Cache",
	"DawnGraphiteCache",
	"DawnWebGPUCache",
	"GPUCache",
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0600)
}

func copyDir(src, dst string, excludeNames []string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dst, 0700); err != nil {
		return err
	}
	for _, entry := range entries {
		name := entry.Name()
		shouldExclude := false
		for _, ex := range excludeNames {
			if strings.EqualFold(name, ex) {
				shouldExclude = true
				break
			}
		}
		if shouldExclude {
			continue
		}
		srcPath := filepath.Join(src, name)
		dstPath := filepath.Join(dst, name)
		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath, excludeNames); err != nil {
				return err
			}
		} else {
			_ = copyFile(srcPath, dstPath)
		}
	}
	return nil
}

// SaveAccount captures current Antigravity session (either Standalone Hub or IDE) and saves under name
func SaveAccount(name string, dbPath string) error {
	if name == "" {
		return fmt.Errorf("account name cannot be empty")
	}

	storageDir, err := GetStorageDir()
	if err != nil {
		return err
	}

	accountProfileDir := filepath.Join(storageDir, name)

	// Mode 1: Check Antigravity Standalone Hub first (if dbPath is not explicitly provided)
	hubDir := FindAntigravityHubDir()
	if dbPath == "" && hubDir != "" {
		appStoragePath := filepath.Join(hubDir, "app_storage.json")
		if fi, err := os.Stat(appStoragePath); err == nil && !fi.IsDir() {
			// Read email from app_storage.json
			var email string
			data, err := os.ReadFile(appStoragePath)
			if err == nil {
				var d map[string]interface{}
				if err := json.Unmarshal(data, &d); err == nil {
					if userVal, ok := d["jetski.onboarding.lastLoginUsername"].(string); ok && userVal != "" {
						email = userVal
					}
				}
				if email == "" {
					if match := emailRe.FindString(string(data)); match != "" {
						email = match
					}
				}
			}

			// Capture OS Keyring token payload (gemini / antigravity)
			var keyringSecret string
			if krVal, err := keyring.Get("gemini", "antigravity"); err == nil && krVal != "" {
				keyringSecret = krVal
				if email == "" || email == "Google User (Antigravity Hub)" {
					if match := emailRe.FindString(krVal); match != "" {
						email = match
					}
				}
			}

			if email == "" {
				email = "Google User (Antigravity Hub)"
			}

			_ = os.MkdirAll(accountProfileDir, 0700)

			// Backup Hub files
			for _, f := range hubFiles {
				src := filepath.Join(hubDir, f)
				dst := filepath.Join(accountProfileDir, f)
				if _, err := os.Stat(src); err == nil {
					_ = copyFile(src, dst)
				}
			}

			// Backup Hub directories (Partitions, Local Storage, Session Storage)
			for _, d := range hubDirs {
				src := filepath.Join(hubDir, d)
				dst := filepath.Join(accountProfileDir, d)
				if fi, err := os.Stat(src); err == nil && fi.IsDir() {
					_ = copyDir(src, dst, hubCacheExcludes)
				}
			}

			if keyringSecret != "" {
				_ = os.WriteFile(filepath.Join(accountProfileDir, "keyring_secret.json"), []byte(keyringSecret), 0600)
			}

			accountMeta := SavedAccount{
				Name:          name,
				SavedAt:       time.Now().Format("2006-01-02 15:04:05"),
				Email:         email,
				Source:        "hub",
				KeyringSecret: keyringSecret,
			}

			metaBytes, _ := json.MarshalIndent(accountMeta, "", "  ")
			_ = os.WriteFile(filepath.Join(accountProfileDir, "profile.json"), metaBytes, 0600)

			// Also save flat json for fast listing
			_ = os.WriteFile(filepath.Join(storageDir, name+".json"), metaBytes, 0600)
			return nil
		}
	}

	// Mode 2: Check Antigravity IDE (VS Code state.vscdb)
	if dbPath == "" {
		dbPath = quota.FindStateVscdb()
	}

	if dbPath != "" {
		db, err := sql.Open("sqlite", fmt.Sprintf("file:%s?mode=ro", dbPath))
		if err == nil {
			defer db.Close()

			items := make(map[string]string)
			var detectedEmail string

			for _, k := range ideKeys {
				var val string
				err := db.QueryRow("SELECT value FROM ItemTable WHERE key = ?", k).Scan(&val)
				if err == nil && val != "" {
					items[k] = val
					if detectedEmail == "" {
						if match := emailRe.FindString(val); match != "" {
							detectedEmail = match
						}
					}
				}
			}

			if len(items) > 0 && items["antigravityUnifiedStateSync.oauthToken"] != "" {
				if detectedEmail == "" {
					detectedEmail = "Google User (IDE)"
				}
				account := SavedAccount{
					Name:       name,
					SavedAt:    time.Now().Format("2006-01-02 15:04:05"),
					Email:      detectedEmail,
					Source:     "ide",
					StateItems: items,
				}

				filePath := filepath.Join(storageDir, name+".json")
				data, err := json.MarshalIndent(account, "", "  ")
				if err != nil {
					return err
				}
				return os.WriteFile(filePath, data, 0600)
			}
		}
	}

	return fmt.Errorf("no active Antigravity session found (checked Hub at %s and IDE state.vscdb) — please launch Antigravity and sign in first", hubDir)
}

// SwitchAccount restores credentials for name into Antigravity Hub or IDE
func SwitchAccount(name string, dbPath string) error {
	storageDir, err := GetStorageDir()
	if err != nil {
		return err
	}

	accountProfileDir := filepath.Join(storageDir, name)
	flatJsonPath := filepath.Join(storageDir, name+".json")

	var account SavedAccount
	var metaBytes []byte

	if data, err := os.ReadFile(filepath.Join(accountProfileDir, "profile.json")); err == nil {
		metaBytes = data
	} else if data, err := os.ReadFile(flatJsonPath); err == nil {
		metaBytes = data
	} else {
		return fmt.Errorf("account profile '%s' not found", name)
	}

	if err := json.Unmarshal(metaBytes, &account); err != nil {
		return fmt.Errorf("failed to parse profile '%s': %w", name, err)
	}

	// Restore Mode 1: Antigravity Standalone Hub
	if account.Source == "hub" {
		hubDir := FindAntigravityHubDir()
		if hubDir == "" {
			return fmt.Errorf("Antigravity Hub directory not found on this system")
		}

		// 0. Restore OS Keyring credentials if present (gemini / antigravity)
		if account.KeyringSecret != "" {
			_ = keyring.Set("gemini", "antigravity", account.KeyringSecret)
		} else if krBytes, err := os.ReadFile(filepath.Join(accountProfileDir, "keyring_secret.json")); err == nil {
			_ = keyring.Set("gemini", "antigravity", string(krBytes))
		}

		// 1. Restore Hub files
		for _, f := range hubFiles {
			src := filepath.Join(accountProfileDir, f)
			dst := filepath.Join(hubDir, f)
			if _, err := os.Stat(src); err == nil {
				_ = copyFile(src, dst)
			}
		}

		// 2. Restore Hub directories (Partitions, Local Storage, Session Storage)
		for _, d := range hubDirs {
			src := filepath.Join(accountProfileDir, d)
			dst := filepath.Join(hubDir, d)
			if fi, err := os.Stat(src); err == nil && fi.IsDir() {
				_ = os.RemoveAll(dst)
				_ = copyDir(src, dst, hubCacheExcludes)
			}
		}

		// 3. Purge temporary cache to prevent quota/session cross-contamination
		for _, cacheDir := range []string{"Cache", "Code Cache", "GPUCache", "DawnGraphiteCache", "DawnWebGPUCache"} {
			_ = os.RemoveAll(filepath.Join(hubDir, cacheDir))
		}

		return nil
	}

	// Restore Mode 2: IDE (VS Code state.vscdb)
	if len(account.StateItems) > 0 {
		if dbPath == "" {
			dbPath = quota.FindStateVscdb()
			if dbPath == "" {
				return fmt.Errorf("state.vscdb not found")
			}
		}

		db, err := sql.Open("sqlite", dbPath)
		if err != nil {
			return fmt.Errorf("failed to open database for write: %w", err)
		}
		defer db.Close()

		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}
		defer tx.Rollback()

		for k, v := range account.StateItems {
			_, err := tx.Exec("INSERT INTO ItemTable(key, value) VALUES(?, ?) ON CONFLICT(key) DO UPDATE SET value=excluded.value", k, v)
			if err != nil {
				return fmt.Errorf("failed to update key %s: %w", k, err)
			}
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit account switch: %w", err)
		}
		return nil
	}

	return fmt.Errorf("profile '%s' has no restorable session data", name)
}

// ListAccounts returns all saved accounts
func ListAccounts() ([]AccountSummary, error) {
	storageDir, err := GetStorageDir()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(storageDir)
	if err != nil {
		return nil, err
	}

	var list []AccountSummary
	seen := make(map[string]bool)

	for _, e := range entries {
		var metaBytes []byte
		var name string

		if e.IsDir() {
			name = e.Name()
			if data, err := os.ReadFile(filepath.Join(storageDir, name, "profile.json")); err == nil {
				metaBytes = data
			}
		} else if filepath.Ext(e.Name()) == ".json" {
			name = e.Name()[:len(e.Name())-5]
			if data, err := os.ReadFile(filepath.Join(storageDir, e.Name())); err == nil {
				metaBytes = data
			}
		}

		if len(metaBytes) > 0 && !seen[name] {
			seen[name] = true
			var acc SavedAccount
			if err := json.Unmarshal(metaBytes, &acc); err == nil {
				src := "IDE"
				if acc.Source == "hub" {
					src = "Standalone"
				}
				list = append(list, AccountSummary{
					Name:    acc.Name,
					SavedAt: acc.SavedAt,
					Email:   acc.Email,
					Source:  src,
				})
			}
		}
	}
	return list, nil
}

// DeleteAccount removes a saved account profile
func DeleteAccount(name string) error {
	storageDir, err := GetStorageDir()
	if err != nil {
		return err
	}
	_ = os.Remove(filepath.Join(storageDir, name+".json"))
	_ = os.RemoveAll(filepath.Join(storageDir, name))
	return nil
}
