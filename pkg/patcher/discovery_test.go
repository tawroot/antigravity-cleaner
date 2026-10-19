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
	var foundOpt, foundHomeApp, foundRpmExtension, foundRpmClassic bool
	for _, c := range candidates {
		if c == "/opt/antigravity/resources/bin/language_server" {
			foundOpt = true
		}
		if filepath.Base(c) == "language_server" || filepath.Base(c) == "language_server.exe" {
			foundHomeApp = true
		}
		if c == "/usr/share/antigravity/resources/app/extensions/antigravity/bin/language_server_linux_x64" {
			foundRpmExtension = true
		}
		if c == "/usr/share/antigravity/resources/bin/language_server" {
			foundRpmClassic = true
		}
	}

	if !foundOpt {
		t.Errorf("expected /opt/antigravity/... candidate in standard paths")
	}
	if !foundHomeApp {
		t.Errorf("expected language_server in candidate list")
	}
	if !foundRpmExtension {
		t.Errorf("expected RPM extension path /usr/share/antigravity/resources/app/extensions/antigravity/bin/language_server_linux_x64 in candidate list")
	}
	if !foundRpmClassic {
		t.Errorf("expected classic RPM path /usr/share/antigravity/resources/bin/language_server in candidate list")
	}
}

func TestIsSelfOrCleanerProcess(t *testing.T) {
	// 1. Current PID must be recognized as self
	if !isSelfOrCleanerProcess(os.Getpid(), "/usr/share/antigravity/bin/antigravity") {
		t.Errorf("expected current PID %d to be recognized as self", os.Getpid())
	}

	// 2. Parent PID must be recognized as self/parent
	if ppid := os.Getppid(); ppid > 0 {
		if !isSelfOrCleanerProcess(ppid, "/usr/share/antigravity/bin/antigravity") {
			t.Errorf("expected parent PID %d to be recognized as parent", ppid)
		}
	}

	// 3. Cleaner and patcher binary keywords
	testCases := []struct {
		pid      int
		exePath  string
		expected bool
	}{
		{999999, "/usr/bin/antigravity-cleaner-linux-amd64", true},
		{999999, "/usr/bin/antigravity-cleaner", true},
		{999999, "/usr/local/bin/ag-cleaner", true},
		{999999, "/opt/patcher/patcher", true},
		{999999, "C:\\Tools\\antigravity-cleaner.exe", true},
		{999999, "/usr/bin/antigravity-gui", true},
		{999999, "/usr/share/antigravity/bin/antigravity", false},
		{999999, "/usr/share/antigravity/antigravity", false},
		{999999, "/usr/share/antigravity/resources/app/extensions/antigravity/bin/language_server_linux_x64", false},
		{999999, "/opt/antigravity/resources/bin/language_server", false},
	}

	for _, tc := range testCases {
		res := isSelfOrCleanerProcess(tc.pid, tc.exePath)
		if res != tc.expected {
			t.Errorf("isSelfOrCleanerProcess(%d, %q) = %v; want %v", tc.pid, tc.exePath, res, tc.expected)
		}
	}
}

func TestFindLanguageServerDynamicPath(t *testing.T) {
	tmpDir := t.TempDir()
	binDir := filepath.Join(tmpDir, "bin")
	shareDir := filepath.Join(tmpDir, "share", "antigravity")
	appBinDir := filepath.Join(shareDir, "bin")
	extBinDir := filepath.Join(shareDir, "resources", "app", "extensions", "antigravity", "bin")

	if err := os.MkdirAll(binDir, 0755); err != nil {
		t.Fatalf("failed to create binDir: %v", err)
	}
	if err := os.MkdirAll(appBinDir, 0755); err != nil {
		t.Fatalf("failed to create appBinDir: %v", err)
	}
	if err := os.MkdirAll(extBinDir, 0755); err != nil {
		t.Fatalf("failed to create extBinDir: %v", err)
	}

	// Create real application executable in shareDir/bin/antigravity
	realApp := filepath.Join(appBinDir, "antigravity")
	if err := os.WriteFile(realApp, []byte("#!/bin/sh\necho antigravity"), 0755); err != nil {
		t.Fatalf("failed to create real app: %v", err)
	}

	// Create symlink in binDir/antigravity -> shareDir/bin/antigravity
	symlinkApp := filepath.Join(binDir, "antigravity")
	if err := os.Symlink(realApp, symlinkApp); err != nil {
		// Fallback to copy if symlinks not supported on filesystem
		if err := os.WriteFile(symlinkApp, []byte("#!/bin/sh\necho antigravity"), 0755); err != nil {
			t.Fatalf("failed to create symlinkApp: %v", err)
		}
	}

	// Create target language server binary
	binNames := GetLanguageServerBinaryNames()
	targetName := binNames[0]
	targetLs := filepath.Join(extBinDir, targetName)
	if err := os.WriteFile(targetLs, []byte("fake language server binary"), 0755); err != nil {
		t.Fatalf("failed to create targetLs: %v", err)
	}

	// 1. Verify candidate roots and probe logic with symlink
	roots := getCandidateRoots(symlinkApp)
	found := probeRootsForLanguageServer(roots)
	if found != targetLs {
		t.Errorf("probeRootsForLanguageServer() = %q; want %q", found, targetLs)
	}

	// 2. Also verify FindIdeMainJs dynamic discovery via PATH
	mainJsPath := filepath.Join(shareDir, "resources", "app", "out", "main.js")
	if err := os.MkdirAll(filepath.Dir(mainJsPath), 0755); err != nil {
		t.Fatalf("failed to create main.js dir: %v", err)
	}
	if err := os.WriteFile(mainJsPath, []byte("// ide main.js"), 0644); err != nil {
		t.Fatalf("failed to create main.js: %v", err)
	}

	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", binDir+string(os.PathListSeparator)+oldPath)
	defer os.Setenv("PATH", oldPath)

	foundJs := FindIdeMainJs()
	if foundJs != mainJsPath {
		t.Errorf("FindIdeMainJs() dynamic PATH = %q; want %q", foundJs, mainJsPath)
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
