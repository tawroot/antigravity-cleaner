package proxy

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEnsureDefaultIcon(t *testing.T) {
	tmpDir := t.TempDir()
	targetIcon := filepath.Join(tmpDir, "icons", "antigravity.png")

	result := EnsureDefaultIcon(targetIcon)
	if result != targetIcon {
		t.Fatalf("expected result %s, got %s", targetIcon, result)
	}

	fi, err := os.Stat(targetIcon)
	if err != nil {
		t.Fatalf("failed to stat written icon: %v", err)
	}
	if fi.Size() < 1000 {
		t.Fatalf("icon size suspiciously small: %d bytes", fi.Size())
	}
}

func TestFindAllAntigravityExecutables(t *testing.T) {
	// Ensure the function executes safely and returns slice without panicking
	targets := FindAllAntigravityExecutables()
	t.Logf("Discovered targets: %d", len(targets))
	for _, target := range targets {
		t.Logf(" - Target: %s, ExePath: %s, IsIDE: %v", target.Name, target.ExePath, target.IsIDE)
	}
}
