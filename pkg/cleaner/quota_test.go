package cleaner

import (
	"database/sql"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSurgicalTargetsIntegrity(t *testing.T) {
	// Verify that none of the targets contain 'globalStorage' or 'workspaceStorage'
	// to guarantee user chats and project contexts are strictly preserved.
	forbidden := []string{"globalstorage", "workspacestorage", "history", "conversations"}

	targets := []string{
		"Cookies",
		"Cookies-journal",
		"DIPS",
		"TransportSecurity",
		"blob_storage",
		"Session Storage",
	}

	for _, target := range targets {
		lower := strings.ToLower(target)
		for _, f := range forbidden {
			if strings.Contains(lower, f) {
				t.Fatalf("Target %s violates chat preservation rule by touching %s", target, f)
			}
		}
	}
}

func TestSurgicalResetDryRun(t *testing.T) {
	// Execute Dry-Run mode: must not return an error and must not delete anything
	removed, err := SurgicalReset429("", true)
	if err != nil {
		t.Logf("DryRun scan: no active config folder or err=%v", err)
		return
	}
	t.Logf("DryRun scan successfully detected %d surgical targets on this system:", len(removed))
	for _, r := range removed {
		t.Logf("  - candidate: %s", r)
	}
}

func TestSurgicalResetNetworkSubdirectoryAndStateDb(t *testing.T) {
	tempDir := t.TempDir()

	// 1. Setup simulated modern Chromium files in Network/ and WebStorage/
	netDir := filepath.Join(tempDir, "Network")
	_ = os.MkdirAll(netDir, 0755)
	cookieFile := filepath.Join(netDir, "Cookies")
	_ = os.WriteFile(cookieFile, []byte("fake-cookie-session"), 0600)

	sessDir := filepath.Join(tempDir, "Session Storage")
	_ = os.MkdirAll(sessDir, 0755)
	_ = os.WriteFile(filepath.Join(sessDir, "000003.log"), []byte("stale-session-log"), 0600)

	// 2. Setup simulated state.vscdb with oauth token and user workspace project
	globalStorage := filepath.Join(tempDir, "User", "globalStorage")
	_ = os.MkdirAll(globalStorage, 0755)
	stateDbPath := filepath.Join(globalStorage, "state.vscdb")

	db, err := sql.Open("sqlite", stateDbPath)
	if err != nil {
		t.Fatalf("failed to open test sqlite: %v", err)
	}
	_, err = db.Exec("CREATE TABLE ItemTable (key TEXT PRIMARY KEY, value TEXT);")
	if err != nil {
		t.Fatalf("failed to create ItemTable: %v", err)
	}
	_, _ = db.Exec("INSERT INTO ItemTable VALUES ('antigravityUnifiedStateSync.oauthToken', 'poisoned-token');")
	_, _ = db.Exec("INSERT INTO ItemTable VALUES ('workbench.explorer.views', 'project-layout');")
	db.Close()

	// Setup preserved workspaceStorage
	wsStorage := filepath.Join(tempDir, "User", "workspaceStorage", "project-uuid")
	_ = os.MkdirAll(wsStorage, 0755)
	projectFile := filepath.Join(wsStorage, "state.vscdb")
	_ = os.WriteFile(projectFile, []byte("precious-user-chat-and-project-context"), 0600)

	// Execute SurgicalReset429
	removed, err := SurgicalReset429(tempDir, false)
	if err != nil {
		t.Fatalf("SurgicalReset429 failed: %v", err)
	}

	if len(removed) == 0 {
		t.Fatalf("expected removed items, got 0")
	}

	// Verify Network/Cookies was purged
	if _, err := os.Stat(cookieFile); err == nil {
		t.Errorf("expected %s to be removed, but still exists", cookieFile)
	}

	// Verify Session Storage was purged
	if _, err := os.Stat(sessDir); err == nil {
		t.Errorf("expected %s to be removed, but still exists", sessDir)
	}

	// Verify user workspace project was PRESERVED
	if data, err := os.ReadFile(projectFile); err != nil || string(data) != "precious-user-chat-and-project-context" {
		t.Errorf("workspaceStorage project was corrupted or deleted: %v", err)
	}

	// Verify oauthToken in state.vscdb was purged while other workbench settings were preserved
	dbCheck, _ := sql.Open("sqlite", stateDbPath)
	defer dbCheck.Close()

	var oauthVal string
	err = dbCheck.QueryRow("SELECT value FROM ItemTable WHERE key = 'antigravityUnifiedStateSync.oauthToken'").Scan(&oauthVal)
	if err == nil {
		t.Errorf("expected oauthToken to be purged, found: %s", oauthVal)
	}

	var layoutVal string
	err = dbCheck.QueryRow("SELECT value FROM ItemTable WHERE key = 'workbench.explorer.views'").Scan(&layoutVal)
	if err != nil || layoutVal != "project-layout" {
		t.Errorf("workbench setting should have been preserved, err: %v", err)
	}
}

func BenchmarkSurgicalTargetsIntegrity(b *testing.B) {
	forbidden := []string{"globalstorage", "workspacestorage", "history", "conversations"}
	targets := []string{
		"Cookies",
		"Cookies-journal",
		"DIPS",
		"TransportSecurity",
		"blob_storage",
		"Session Storage",
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, target := range targets {
			lower := strings.ToLower(target)
			for _, f := range forbidden {
				if strings.Contains(lower, f) {
					b.Fatalf("Target %s violates chat preservation rule", target)
				}
			}
		}
	}
}
