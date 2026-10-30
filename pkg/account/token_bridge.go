package account

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zalando/go-keyring"
)

// StandaloneTokenFilePath returns the canonical path for jetski-standalone-oauth-token
func StandaloneTokenFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".gemini", "jetski-standalone-oauth-token")
}

// SyncTokenBridge synchronizes OAuth credentials bi-directionally between OS Keyring and file storage
// to ensure zero session drops when running in SSH/remote spoofing mode.
func SyncTokenBridge(logger func(format string, args ...interface{})) error {
	if logger == nil {
		logger = func(format string, args ...interface{}) {}
	}

	tokenFilePath := StandaloneTokenFilePath()
	if tokenFilePath == "" {
		return fmt.Errorf("unable to determine user home directory")
	}

	// 1. Check OS Keyring
	krVal, krErr := keyring.Get("gemini", "antigravity")
	krHasValidToken := false
	if krErr == nil && strings.TrimSpace(krVal) != "" {
		var d map[string]interface{}
		if err := json.Unmarshal([]byte(krVal), &d); err == nil {
			if _, ok := d["token"]; ok {
				krHasValidToken = true
			}
		}
	}

	// 2. Check File Storage (~/.gemini/jetski-standalone-oauth-token)
	var fileBytes []byte
	fileHasValidToken := false
	if fi, err := os.Stat(tokenFilePath); err == nil && !fi.IsDir() {
		if data, err := os.ReadFile(tokenFilePath); err == nil && len(data) > 0 {
			fileBytes = data
			var d map[string]interface{}
			if err := json.Unmarshal(data, &d); err == nil {
				if _, ok := d["token"]; ok {
					fileHasValidToken = true
				}
			}
		}
	}

	// Bridge Direction A: Keyring -> File
	if krHasValidToken && !fileHasValidToken {
		geminiDir := filepath.Dir(tokenFilePath)
		if err := os.MkdirAll(geminiDir, 0700); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", geminiDir, err)
		}
		if err := os.WriteFile(tokenFilePath, []byte(krVal), 0600); err != nil {
			return fmt.Errorf("failed to write token file %s: %w", tokenFilePath, err)
		}
		logger("Synced active credentials from OS Keyring to file token storage (%s)", tokenFilePath)
		return nil
	}

	// Bridge Direction B: File -> Keyring
	if fileHasValidToken && !krHasValidToken {
		if err := keyring.Set("gemini", "antigravity", string(fileBytes)); err != nil {
			logger("Warning: failed to sync file token to OS Keyring: %v", err)
		} else {
			logger("Synced active credentials from file token storage to OS Keyring")
		}
		return nil
	}

	// If both have valid tokens but are different, update file to match active keyring
	if krHasValidToken && fileHasValidToken {
		krHash := sha256.Sum256([]byte(krVal))
		fileHash := sha256.Sum256(fileBytes)
		if krHash != fileHash {
			_ = os.WriteFile(tokenFilePath, []byte(krVal), 0600)
			logger("Updated file token storage with latest active OS Keyring credentials")
		}
	}

	return nil
}
