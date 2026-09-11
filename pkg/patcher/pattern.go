package patcher

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// Pattern represents a binary byte pattern with wildcards
type Pattern struct {
	Bytes []byte
	Mask  []bool // true = exact match, false = wildcard (any byte)
}

// ParsePattern parses a pattern string like "80 78 08 00 74 ?? 48 8b ?? 24 ?? 48 89 ?? 60"
func ParsePattern(sig string) (*Pattern, error) {
	tokens := strings.Fields(sig)
	p := &Pattern{
		Bytes: make([]byte, len(tokens)),
		Mask:  make([]bool, len(tokens)),
	}

	for i, token := range tokens {
		if token == "?" || token == "??" {
			p.Bytes[i] = 0
			p.Mask[i] = false
		} else {
			b, err := hex.DecodeString(token)
			if err != nil || len(b) != 1 {
				return nil, fmt.Errorf("invalid hex token %q: %w", token, err)
			}
			p.Bytes[i] = b[0]
			p.Mask[i] = true
		}
	}
	return p, nil
}

// Find returns the first index of the pattern in data, or -1
func (p *Pattern) Find(data []byte) int {
	patLen := len(p.Bytes)
	if patLen == 0 || len(data) < patLen {
		return -1
	}

	firstByte := p.Bytes[0]
	firstMask := p.Mask[0]

	limit := len(data) - patLen
	for i := 0; i <= limit; i++ {
		// Fast skip on first byte if masked
		if firstMask && data[i] != firstByte {
			continue
		}

		matched := true
		for j := 1; j < patLen; j++ {
			if p.Mask[j] && data[i+j] != p.Bytes[j] {
				matched = false
				break
			}
		}
		if matched {
			return i
		}
	}
	return -1
}

// FindAll returns all indices of the pattern in data
func (p *Pattern) FindAll(data []byte) []int {
	var matches []int
	patLen := len(p.Bytes)
	if patLen == 0 || len(data) < patLen {
		return matches
	}

	limit := len(data) - patLen
	for i := 0; i <= limit; i++ {
		if p.Mask[0] && data[i] != p.Bytes[0] {
			continue
		}

		matched := true
		for j := 1; j < patLen; j++ {
			if p.Mask[j] && data[i+j] != p.Bytes[j] {
				matched = false
				break
			}
		}
		if matched {
			matches = append(matches, i)
			i += patLen - 1
		}
	}
	return matches
}
