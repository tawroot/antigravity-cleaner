package quota

import (
	"database/sql"
	"path/filepath"
	"strings"
	"testing"
	"time"

	_ "modernc.org/sqlite"
)

func TestExtractRefreshToken(t *testing.T) {
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

	// Insert mock oauthToken
	fakeToken := "dummy_prefix_1//04test_mock_refresh_token_xyz123_suffix"
	_, err = db.Exec("INSERT INTO ItemTable (key, value) VALUES ('antigravityUnifiedStateSync.oauthToken', ?)", fakeToken)
	if err != nil {
		t.Fatalf("failed to insert token: %v", err)
	}
	db.Close()

	token, err := ExtractRefreshToken(dbPath)
	if err != nil {
		t.Fatalf("ExtractRefreshToken failed: %v", err)
	}

	if !strings.HasPrefix(token, "1//04test_mock_refresh_token") {
		t.Errorf("unexpected extracted token: %s", token)
	}
}

func TestRenderQuota(t *testing.T) {
	rem := 0.75
	groups := []QuotaGroup{
		{
			DisplayName: "Claude and GPT models",
			Buckets: []QuotaBucket{
				{
					Window:            "5h",
					RemainingFraction: &rem,
					ResetTime:         time.Now().Add(2 * time.Hour),
				},
			},
		},
	}
	detailed := []DetailedQuotaBucket{
		{
			ModelID:           "gemini-2.5-pro",
			RemainingFraction: &rem,
			ResetTime:         time.Now().Add(1 * time.Hour),
		},
		{
			ModelID:           "gemini-3.1-flash-image",
			RemainingFraction: &rem,
			ResetTime:         time.Now().Add(1 * time.Hour),
		},
	}

	output := RenderQuota(groups, detailed)
	if !strings.Contains(output, "Gemini Models") {
		t.Errorf("expected Gemini Models in output, got: %s", output)
	}
	if !strings.Contains(output, "Claude and GPT models") {
		t.Errorf("expected Claude and GPT models in output, got: %s", output)
	}
	if !strings.Contains(output, "Gemini Image Generation") || !strings.Contains(output, "gemini-3.1-flash-image") {
		t.Errorf("expected Gemini Image Generation in output, got: %s", output)
	}
	if !strings.Contains(output, "75%") {
		t.Errorf("expected 75%% in output, got: %s", output)
	}

	rawOutput := RenderQuotaRaw(groups, detailed)
	if !strings.Contains(rawOutput, "gemini-2.5-pro") {
		t.Errorf("expected gemini-2.5-pro in raw output, got: %s", rawOutput)
	}
}

func BenchmarkRenderQuota(b *testing.B) {
	rem := 0.75
	groups := []QuotaGroup{
		{
			DisplayName: "Claude and GPT models",
			Buckets: []QuotaBucket{
				{
					Window:            "5h",
					RemainingFraction: &rem,
					ResetTime:         time.Now().Add(2 * time.Hour),
				},
			},
		},
	}
	detailed := []DetailedQuotaBucket{
		{
			ModelID:           "gemini-2.5-pro",
			RemainingFraction: &rem,
			ResetTime:         time.Now().Add(1 * time.Hour),
		},
		{
			ModelID:           "gemini-3.1-flash-image",
			RemainingFraction: &rem,
			ResetTime:         time.Now().Add(1 * time.Hour),
		},
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = RenderQuota(groups, detailed)
	}
}
