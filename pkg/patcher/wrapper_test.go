package patcher

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWrapperManagerInstallAndRemove(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "wrapper-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	fakeBinaryPath := filepath.Join(tempDir, "language_server")
	if err := os.WriteFile(fakeBinaryPath, []byte("#!/bin/sh\necho real"), 0755); err != nil {
		t.Fatalf("failed to create fake binary: %v", err)
	}

	mgr := NewWrapperManager(tempDir, func(format string, args ...interface{}) {})

	if mgr.IsWrapperInstalled() {
		t.Errorf("expected wrapper not installed initially")
	}

	opts := DefaultEnterpriseFlagsOptions()
	if err := mgr.InstallWrapper(opts); err != nil {
		t.Fatalf("failed to install wrapper: %v", err)
	}

	if !mgr.IsWrapperInstalled() {
		t.Errorf("expected wrapper to be installed")
	}

	// Verify content of installed wrapper
	content, err := os.ReadFile(fakeBinaryPath)
	if err != nil {
		t.Fatalf("failed to read wrapper: %v", err)
	}
	if len(content) == 0 {
		t.Errorf("wrapper file is empty")
	}

	// Remove wrapper
	if err := mgr.RemoveWrapper(); err != nil {
		t.Fatalf("failed to remove wrapper: %v", err)
	}

	if mgr.IsWrapperInstalled() {
		t.Errorf("expected wrapper to be uninstalled")
	}

	// Verify original content restored
	restored, err := os.ReadFile(fakeBinaryPath)
	if err != nil {
		t.Fatalf("failed to read restored file: %v", err)
	}
	if string(restored) != "#!/bin/sh\necho real" {
		t.Errorf("restored content mismatch: %s", string(restored))
	}
}
