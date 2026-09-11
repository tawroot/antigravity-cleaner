package patcher

import (
	"bytes"
	"crypto/rand"
	"testing"
)

func TestPatternWildcardMatching(t *testing.T) {
	patternStr := "80 78 08 00 74 ?? 48 8b ?? 24 ?? 48 89 ?? 60"
	p, err := ParsePattern(patternStr)
	if err != nil {
		t.Fatalf("Failed to parse pattern: %v", err)
	}

	if len(p.Bytes) != 15 {
		t.Fatalf("Expected pattern length 15, got %d", len(p.Bytes))
	}

	// Test exact match with wildcards
	data := []byte{
		0x90, 0x90, // NOPs
		0x80, 0x78, 0x08, 0x00, 0x74, 0xAA, 0x48, 0x8b, 0xBB, 0x24, 0xCC, 0x48, 0x89, 0xDD, 0x60,
		0x90,
	}

	idx := p.Find(data)
	if idx != 2 {
		t.Fatalf("Expected match at offset 2, got %d", idx)
	}

	// Test no match
	badData := []byte{0x80, 0x78, 0x08, 0x01, 0x74, 0xAA}
	if p.Find(badData) != -1 {
		t.Fatalf("Expected no match, but found one")
	}
}

func BenchmarkPatternSearch(b *testing.B) {
	patternStr := "80 78 08 00 74 ?? 48 8b ?? 24 ?? 48 89 ?? 60"
	p, err := ParsePattern(patternStr)
	if err != nil {
		b.Fatalf("Failed to parse pattern: %v", err)
	}

	// Create 10MB test buffer with signature at the end
	size := 10 * 1024 * 1024
	buf := make([]byte, size)
	_, _ = rand.Read(buf)

	target := []byte{0x80, 0x78, 0x08, 0x00, 0x74, 0x12, 0x48, 0x8b, 0x34, 0x24, 0x56, 0x48, 0x89, 0x78, 0x60}
	copy(buf[size-100:], target)

	b.ResetTimer()
	b.SetBytes(int64(size))

	for i := 0; i < b.N; i++ {
		idx := p.Find(buf)
		if idx == -1 {
			b.Fatalf("Pattern not found in benchmark buffer")
		}
	}
}

func TestArm64Pattern(t *testing.T) {
	sig := "?? ?? ?? 39 ?? ?? ?? 35 ?? ?? ?? 39 ?? ?? ?? 14"
	p, err := ParsePattern(sig)
	if err != nil {
		t.Fatalf("Failed to parse ARM64 pattern: %v", err)
	}

	dummy := bytes.Repeat([]byte{0x00, 0x00, 0x00, 0x39, 0x00, 0x00, 0x00, 0x35, 0x00, 0x00, 0x00, 0x39, 0x00, 0x00, 0x00, 0x14}, 2)
	idx := p.Find(dummy)
	if idx != 0 {
		t.Fatalf("Expected match at offset 0, got %d", idx)
	}
}
