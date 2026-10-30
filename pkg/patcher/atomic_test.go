package patcher

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestIsMachOBinary(t *testing.T) {
	tests := []struct {
		name     string
		header   []byte
		expected bool
	}{
		{"Short slice", []byte{0xfe, 0xed}, false},
		{"Mach-O 32-bit BE", []byte{0xfe, 0xed, 0xfa, 0xce, 0x00, 0x00}, true},
		{"Mach-O 32-bit LE", []byte{0xce, 0xfa, 0xed, 0xfe, 0x00, 0x00}, true},
		{"Mach-O 64-bit BE", []byte{0xfe, 0xed, 0xfa, 0xcf, 0x00, 0x00}, true},
		{"Mach-O 64-bit LE", []byte{0xcf, 0xfa, 0xed, 0xfe, 0x00, 0x00}, true},
		{"Fat Mach-O BE", []byte{0xca, 0xfe, 0xba, 0xbe, 0x00, 0x00}, true},
		{"Fat Mach-O LE", []byte{0xbe, 0xba, 0xfe, 0xca, 0x00, 0x00}, true},
		{"ELF binary", []byte{0x7f, 0x45, 0x4c, 0x46, 0x00, 0x00}, false},
		{"Windows PE", []byte{0x4d, 0x5a, 0x90, 0x00, 0x00, 0x00}, false},
		{"JavaScript text", []byte("const x = 1;"), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if res := IsMachOBinary(tc.header); res != tc.expected {
				t.Fatalf("expected IsMachOBinary(%v) = %v, got %v", tc.header, tc.expected, res)
			}
		})
	}
}

func TestResignDarwinBinaryNonDarwin(t *testing.T) {
	if runtime.GOOS != "darwin" {
		// On non-Darwin, ResignDarwinBinary should safely return nil immediately
		if err := ResignDarwinBinary("/tmp/non-existent-binary"); err != nil {
			t.Fatalf("expected nil error on non-darwin, got %v", err)
		}
	}
}

func TestAtomicWriteAndBackup(t *testing.T) {
	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "test.bin")

	initial := []byte("initial binary data")
	if err := os.WriteFile(target, initial, 0755); err != nil {
		t.Fatalf("failed to write initial file: %v", err)
	}

	bak, err := BackupFile(target)
	if err != nil {
		t.Fatalf("BackupFile failed: %v", err)
	}
	if bak != target+".agybak" {
		t.Fatalf("unexpected backup path: %s", bak)
	}

	updated := []byte("updated patched binary data")
	if err := AtomicWrite(target, updated); err != nil {
		t.Fatalf("AtomicWrite failed: %v", err)
	}

	got, err := os.ReadFile(target)
	if err != nil || string(got) != string(updated) {
		t.Fatalf("AtomicWrite data mismatch")
	}

	if err := RestoreFile(target); err != nil {
		t.Fatalf("RestoreFile failed: %v", err)
	}

	restored, err := os.ReadFile(target)
	if err != nil || string(restored) != string(initial) {
		t.Fatalf("RestoreFile data mismatch")
	}
}
