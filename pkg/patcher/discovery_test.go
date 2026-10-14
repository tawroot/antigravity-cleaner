package patcher

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetStandardLanguageServerCandidates(t *testing.T) {
	candidates := GetStandardLanguageServerCandidates()
	if len(candidates) == 0 {
		t.Fatalf("expected non-empty candidates list")
	}

	// Verify that modern lowercase Linux paths and user Applications are present
	var foundOpt, foundHomeApp bool
	for _, c := range candidates {
		if c == "/opt/antigravity/resources/bin/language_server" {
			foundOpt = true
		}
		if filepath.Base(c) == "language_server" || filepath.Base(c) == "language_server.exe" {
			foundHomeApp = true
		}
	}

	if !foundOpt {
		t.Errorf("expected /opt/antigravity/... candidate in standard paths")
	}
	if !foundHomeApp {
		t.Errorf("expected language_server in candidate list")
	}
}

func TestFindLanguageServerModularEnvOverride(t *testing.T) {
	// Create a dummy temporary file
	tmpDir := t.TempDir()
	dummyBin := filepath.Join(tmpDir, "dummy_language_server")
	if err := os.WriteFile(dummyBin, []byte("dummy binary"), 0755); err != nil {
		t.Fatalf("failed to create dummy binary: %v", err)
	}

	os.Setenv("ANTIGRAVITY_LANGUAGE_SERVER_PATH", dummyBin)
	defer os.Unsetenv("ANTIGRAVITY_LANGUAGE_SERVER_PATH")

	var logMessages []string
	logger := func(format string, args ...interface{}) {
		logMessages = append(logMessages, format)
	}

	found := FindLanguageServerModular(logger)
	if found != dummyBin {
		t.Errorf("expected %s, got %s", dummyBin, found)
	}
	if len(logMessages) == 0 {
		t.Errorf("expected logs to be recorded")
	}
}

func TestFindRunningLanguageServerOnCurrentHost(t *testing.T) {
	var logged []string
	logger := func(format string, args ...interface{}) {
		logged = append(logged, format)
	}

	found := FindRunningLanguageServer(logger)
	t.Logf("Discovery result for running language_server: %s (logged %d messages)", found, len(logged))
}
