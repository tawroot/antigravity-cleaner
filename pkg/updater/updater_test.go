package updater

import (
	"runtime"
	"testing"
)

func TestIsNewerVersion(t *testing.T) {
	tests := []struct {
		latest   string
		current  string
		expected bool
	}{
		{"5.4.1", "5.4.0", true},
		{"5.5.0", "5.4.0", true},
		{"6.0.0", "5.4.0", true},
		{"5.4.0", "5.4.0", false},
		{"5.3.2", "5.4.0", false},
		{"5.4.0.1", "5.4.0", true},
	}

	for _, tt := range tests {
		got := isNewerVersion(tt.latest, tt.current)
		if got != tt.expected {
			t.Errorf("isNewerVersion(%q, %q) = %v; want %v", tt.latest, tt.current, got, tt.expected)
		}
	}
}

func TestGetMatchingAssetName(t *testing.T) {
	name := GetMatchingAssetName()
	if name == "" {
		t.Fatalf("expected non-empty asset name for %s/%s", runtime.GOOS, runtime.GOARCH)
	}
	t.Logf("Platform asset name: %s", name)
}
