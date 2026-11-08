package patcher

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPatchVsCodeExtensionClassic(t *testing.T) {
	tempDir := t.TempDir()
	extJs := filepath.Join(tempDir, "extension.js")

	fakeContent := `
const isChannelChanged = manifestFetched && lastInstalledUrl !== releaseBaseUrl;
outputChannel.appendLine('[INSTALL] Checking Antigravity releases...');
const target = options.targetPathOverride;
`
	if err := os.WriteFile(extJs, []byte(fakeContent), 0644); err != nil {
		t.Fatalf("failed to write fake extension.js: %v", err)
	}

	res, err := PatchVsCodeExtension(extJs)
	if err != nil {
		t.Fatalf("PatchVsCodeExtension failed: %v", err)
	}
	if !strings.Contains(res, "Successfully patched") {
		t.Errorf("unexpected result: %s", res)
	}

	// Test already patched
	res2, err2 := PatchVsCodeExtension(extJs)
	if err2 != nil || !strings.Contains(res2, "Already patched") {
		t.Errorf("expected already patched, got res: %s, err: %v", res2, err2)
	}
}

func TestPatchVsCodeExtensionV110Minified(t *testing.T) {
	tempDir := t.TempDir()
	extJs := filepath.Join(tempDir, "extension.js")

	fakeContent := `
let Sc=Qa&&Md&&a&&!ed;Sc&&g.appendLine(` + "`" + `[INSTALL] Release channel change detected (last installed URL: '${sc}', target URL: '${Na}'). Forcing backend binary re-download.` + "`" + `);
if(!Sc&&Qa) return B;
`
	if err := os.WriteFile(extJs, []byte(fakeContent), 0644); err != nil {
		t.Fatalf("failed to write fake extension.js: %v", err)
	}

	res, err := PatchVsCodeExtension(extJs)
	if err != nil {
		t.Fatalf("PatchVsCodeExtension failed on 1.10.0 minified: %v", err)
	}
	if !strings.Contains(res, "Successfully patched") {
		t.Errorf("unexpected result: %s", res)
	}

	// Test already patched
	res2, err2 := PatchVsCodeExtension(extJs)
	if err2 != nil || !strings.Contains(res2, "Already patched") {
		t.Errorf("expected already patched, got res: %s, err: %v", res2, err2)
	}
}

func TestPatchVsCodeExtensionReal110File(t *testing.T) {
	realPath := "/tmp/ext-1.10.0/extension/extension.js"
	if _, err := os.Stat(realPath); os.IsNotExist(err) {
		t.Skip("real 1.10.0 file not found, skipping")
	}

	tempDir := t.TempDir()
	testCopy := filepath.Join(tempDir, "extension.js")
	data, err := os.ReadFile(realPath)
	if err != nil {
		t.Fatalf("failed to read real file: %v", err)
	}
	if err := os.WriteFile(testCopy, data, 0644); err != nil {
		t.Fatalf("failed to copy real file: %v", err)
	}

	res, err := PatchVsCodeExtension(testCopy)
	if err != nil {
		t.Fatalf("PatchVsCodeExtension failed on real 1.10.0: %v", err)
	}
	t.Logf("Result on real 1.10.0: %s", res)

	// Verify idempotency
	res2, err2 := PatchVsCodeExtension(testCopy)
	if err2 != nil || !strings.Contains(res2, "Already patched") {
		t.Errorf("expected already patched on real 1.10.0, got: %s (err: %v)", res2, err2)
	}
}
