package patcher

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestAsarZeroLogByteLengths(t *testing.T) {
	if len(asarConsoleTarget) != len(asarConsolePatched) {
		t.Fatalf("length mismatch: console target is %d bytes, patched is %d bytes",
			len(asarConsoleTarget), len(asarConsolePatched))
	}

	if len(asarWriteFnTarget) != len(asarWriteFnPatched) {
		t.Fatalf("length mismatch: writeFn target is %d bytes, patched is %d bytes",
			len(asarWriteFnTarget), len(asarWriteFnPatched))
	}

	if len(asarUpdaterTarget) != len(asarUpdaterPatched) {
		t.Fatalf("length mismatch: updater target is %d bytes, patched is %d bytes",
			len(asarUpdaterTarget), len(asarUpdaterPatched))
	}
}

func TestPatchAsarAntiUpdate(t *testing.T) {
	tmpDir := t.TempDir()
	fakeAsar := filepath.Join(tmpDir, "app.asar")
	fakeYml := filepath.Join(tmpDir, "app-update.yml")

	fakeContent := bytes.Join([][]byte{
		[]byte("header_metadata\n"),
		asarUpdaterTarget,
		[]byte("\nend_of_file"),
	}, nil)

	if err := os.WriteFile(fakeAsar, fakeContent, 0644); err != nil {
		t.Fatalf("failed to write fake asar: %v", err)
	}
	if err := os.WriteFile(fakeYml, []byte("provider: generic\nurl: https://antigravity-hub-auto-updater.run.app/\n"), 0644); err != nil {
		t.Fatalf("failed to write fake yml: %v", err)
	}

	res, err := PatchAsarAntiUpdate(fakeAsar)
	if err != nil {
		t.Fatalf("PatchAsarAntiUpdate failed: %v", err)
	}
	if res == "" {
		t.Fatalf("expected non-empty success message")
	}

	patchedData, err := os.ReadFile(fakeAsar)
	if err != nil {
		t.Fatalf("failed to read patched asar: %v", err)
	}
	if !bytes.Contains(patchedData, asarUpdaterPatched) {
		t.Errorf("patched asar missing updater patched bytes")
	}

	ymlData, err := os.ReadFile(fakeYml)
	if err != nil {
		t.Fatalf("failed to read yml: %v", err)
	}
	if !bytes.Contains(ymlData, []byte("127.0.0.1:0")) {
		t.Errorf("expected neutralized yml, got: %s", string(ymlData))
	}
}

func TestPatchAsarZeroLog(t *testing.T) {
	tmpDir := t.TempDir()
	fakeAsar := filepath.Join(tmpDir, "app.asar")

	fakeContent := bytes.Join([][]byte{
		[]byte("header_metadata_test_data\n"),
		asarConsoleTarget,
		[]byte("\nsome_intermediate_code\n"),
		asarWriteFnTarget,
		[]byte("\nend_of_asar"),
	}, nil)

	if err := os.WriteFile(fakeAsar, fakeContent, 0644); err != nil {
		t.Fatalf("failed to write fake asar: %v", err)
	}

	// 1. First patch
	res, err := PatchAsarZeroLog(fakeAsar)
	if err != nil {
		t.Fatalf("PatchAsarZeroLog failed: %v", err)
	}
	if res == "" {
		t.Fatalf("expected non-empty success message")
	}

	// Verify content
	patchedData, err := os.ReadFile(fakeAsar)
	if err != nil {
		t.Fatalf("failed to read patched asar: %v", err)
	}

	if len(patchedData) != len(fakeContent) {
		t.Fatalf("asar length changed! original=%d, patched=%d", len(fakeContent), len(patchedData))
	}

	if !bytes.Contains(patchedData, asarConsolePatched) {
		t.Errorf("patched asar missing console patched string")
	}

	if !bytes.Contains(patchedData, asarWriteFnPatched) {
		t.Errorf("patched asar missing writeFn patched string")
	}

	// 2. Second patch (idempotent)
	res2, err := PatchAsarZeroLog(fakeAsar)
	if err != nil {
		t.Fatalf("second patch call failed: %v", err)
	}
	if !bytes.Contains([]byte(res2), []byte("Already patched")) {
		t.Errorf("expected idempotent message, got: %s", res2)
	}
}
