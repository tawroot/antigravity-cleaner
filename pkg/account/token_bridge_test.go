package account

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStandaloneTokenFilePath(t *testing.T) {
	path := StandaloneTokenFilePath()
	if path == "" {
		t.Fatalf("expected non-empty StandaloneTokenFilePath")
	}
	if filepath.Base(path) != "jetski-standalone-oauth-token" {
		t.Errorf("expected filename 'jetski-standalone-oauth-token', got %s", filepath.Base(path))
	}
}

func TestSyncTokenBridgeExecution(t *testing.T) {
	var logs []string
	logger := func(format string, args ...interface{}) {
		logs = append(logs, format)
	}

	err := SyncTokenBridge(logger)
	if err != nil {
		t.Fatalf("SyncTokenBridge failed: %v", err)
	}

	// Verify token file exists if OS Keyring has credentials
	tokenFile := StandaloneTokenFilePath()
	if fi, err := os.Stat(tokenFile); err == nil && !fi.IsDir() {
		// Verify file permissions are 0600
		perm := fi.Mode().Perm()
		if perm != 0600 {
			t.Errorf("expected file mode 0600, got %o", perm)
		}
	}
}
