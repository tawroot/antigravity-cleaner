package patcher

import (
	"os"
	"testing"
)

func TestPattern(t *testing.T) {
	sig := "80 78 08 00 74 ?? 48 8b ?? 24 ?? 48 89 ?? 60"
	p, err := ParsePattern(sig)
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	sample := []byte("\x80\x78\x08\x00\x74\x3b\x48\x8b\x54\x24\x70\x48\x89\x50\x60")
	idx := p.Find(sample)
	if idx != 0 {
		t.Fatalf("Expected match at 0, got %d", idx)
	}
	t.Log("Pattern match passed!")
}

func TestLanguageServerRealBinary(t *testing.T) {
	lsPath := FindLanguageServer()
	if lsPath == "" {
		t.Skip("language_server not installed locally")
	}

	data, err := os.ReadFile(lsPath)
	if err != nil {
		t.Fatalf("Failed to read language_server: %v", err)
	}

	sigUnpatched := "80 78 08 00 74 ?? 48 8b ?? 24 ?? 48 89 ?? 60"
	sigPatched := "c6 40 08 01 90 90 48 8b ?? 24 ?? 48 89 ?? 60"

	pUnpatched, err := ParsePattern(sigUnpatched)
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	pPatched, err := ParsePattern(sigPatched)
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	idxU := pUnpatched.Find(data)
	idxP := pPatched.Find(data)

	if idxU != -1 {
		t.Logf("Found official unpatched signature at offset: 0x%x (%d)", idxU, idxU)
	} else if idxP != -1 {
		t.Logf("Found unlocked patched signature at offset: 0x%x (%d)", idxP, idxP)
	} else {
		t.Fatalf("Neither unpatched nor patched signature found in real binary!")
	}
}
