package proxy

import (
	_ "embed"
	"os"
	"path/filepath"
)

//go:embed icon.png
var DefaultIconBytes []byte

// EnsureDefaultIcon writes the embedded official icon if no icon exists
func EnsureDefaultIcon(targetPath string) string {
	if fi, err := os.Stat(targetPath); err == nil && fi.Size() > 100 {
		return targetPath
	}
	_ = os.MkdirAll(filepath.Dir(targetPath), 0755)
	if len(DefaultIconBytes) > 0 {
		_ = os.WriteFile(targetPath, DefaultIconBytes, 0644)
	}
	return targetPath
}
