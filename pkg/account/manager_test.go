package account

import (
	"database/sql"
	"path/filepath"
	"testing"

	_ "modernc.org/sqlite"
)

func TestSaveAndSwitchAccount(t *testing.T) {
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "state.vscdb")

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		t.Fatalf("failed to create sqlite db: %v", err)
	}

	_, err = db.Exec("CREATE TABLE ItemTable (key TEXT PRIMARY KEY, value TEXT)")
	if err != nil {
		t.Fatalf("failed to create ItemTable: %v", err)
	}

	// 1. Insert Account 1
	token1 := "token_account_one_1//mock1"
	email1 := `{"email": "user1@example.com"}`
	_, _ = db.Exec("INSERT INTO ItemTable (key, value) VALUES ('antigravityUnifiedStateSync.oauthToken', ?)", token1)
	_, _ = db.Exec("INSERT INTO ItemTable (key, value) VALUES ('antigravityUnifiedStateSync.userStatus', ?)", email1)
	db.Close()

	// Save Account 1
	err = SaveAccount("acc1", dbPath)
	if err != nil {
		t.Fatalf("SaveAccount acc1 failed: %v", err)
	}

	// 2. Overwrite DB with Account 2
	db2, _ := sql.Open("sqlite", dbPath)
	token2 := "token_account_two_1//mock2"
	email2 := `{"email": "user2@example.com"}`
	_, _ = db2.Exec("UPDATE ItemTable SET value = ? WHERE key = 'antigravityUnifiedStateSync.oauthToken'", token2)
	_, _ = db2.Exec("UPDATE ItemTable SET value = ? WHERE key = 'antigravityUnifiedStateSync.userStatus'", email2)
	db2.Close()

	// Save Account 2
	err = SaveAccount("acc2", dbPath)
	if err != nil {
		t.Fatalf("SaveAccount acc2 failed: %v", err)
	}

	// Verify listing
	accounts, err := ListAccounts()
	if err != nil {
		t.Fatalf("ListAccounts failed: %v", err)
	}
	if len(accounts) < 2 {
		t.Fatalf("expected at least 2 accounts, got %d", len(accounts))
	}

	// 3. Switch back to Account 1
	err = SwitchAccount("acc1", dbPath)
	if err != nil {
		t.Fatalf("SwitchAccount acc1 failed: %v", err)
	}

	// Check DB values
	dbCheck, _ := sql.Open("sqlite", dbPath)
	defer dbCheck.Close()
	var currentToken string
	_ = dbCheck.QueryRow("SELECT value FROM ItemTable WHERE key = 'antigravityUnifiedStateSync.oauthToken'").Scan(&currentToken)

	if currentToken != token1 {
		t.Errorf("expected token1 (%s), got: %s", token1, currentToken)
	}

	// Clean up
	_ = DeleteAccount("acc1")
	_ = DeleteAccount("acc2")
}
