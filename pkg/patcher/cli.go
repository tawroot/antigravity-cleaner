package patcher

import (
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var (
	cliGateX64 = func() BinaryGate {
		u, _ := ParsePattern("48 85 c0 0f 84 ?? ?? ?? ?? 80 78 08 00 0f 85 ?? ?? ?? ??")
		p, _ := ParsePattern("48 85 c0 0f 84 ?? ?? ?? ?? 48 85 c0 90 0f 85 ?? ?? ?? ??")
		r, _ := hex.DecodeString("4885c090")
		return BinaryGate{
			Name:        "agy eligibility check bypass (x64)",
			Unpatched:   u,
			Patched:     p,
			Replacement: r,
			Offset:      9,
		}
	}()

	cliGateARM64 = func() BinaryGate {
		u, _ := ParsePattern("?? ?? ?? b5 ?? ?? ?? b4 01 20 40 39 ?? ?? ?? 37")
		p, _ := ParsePattern("?? ?? ?? b5 ?? ?? ?? b4 21 00 80 52 ?? ?? ?? 37")
		r, _ := hex.DecodeString("21008052")
		return BinaryGate{
			Name:        "agy eligibility check bypass (arm64)",
			Unpatched:   u,
			Patched:     p,
			Replacement: r,
			Offset:      8,
		}
	}()
)

// FindAgy searches for agy binary in PATH and standard directories
func FindAgy() string {
	if p, err := exec.LookPath("agy"); err == nil {
		return p
	}
	if p, err := exec.LookPath("agy.exe"); err == nil {
		return p
	}

	home, _ := os.UserHomeDir()
	candidates := []string{
		filepath.Join(home, ".local", "bin", "agy"),
		filepath.Join(home, ".gemini", "bin", "antigravity"),
		filepath.Join(home, ".gemini", "bin", "agy"),
		"/usr/local/bin/agy",
		"/usr/bin/agy",
		filepath.Join(home, "bin", "agy"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "agy", "bin", "agy.exe"),
		filepath.Join(os.Getenv("USERPROFILE"), "scoop", "apps", "agy", "current", "agy.exe"),
	}

	for _, c := range candidates {
		if fi, err := os.Stat(c); err == nil && !fi.IsDir() {
			return c
		}
	}
	return ""
}

// PatchAgy patches the agy binary to bypass the Eligibility Check screen
func PatchAgy(binPath string) (string, error) {
	if binPath == "" {
		binPath = FindAgy()
		if binPath == "" {
			return "", fmt.Errorf("agy CLI binary not found")
		}
	}

	data, err := os.ReadFile(binPath)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", binPath, err)
	}

	gates := []BinaryGate{cliGateX64, cliGateARM64}
	if runtime.GOARCH == "arm64" {
		gates = []BinaryGate{cliGateARM64, cliGateX64}
	}

	for _, g := range gates {
		if g.Patched.Find(data) != -1 {
			return fmt.Sprintf("Already patched: %s (%s)", filepath.Base(binPath), g.Name), nil
		}
	}

	var matchedGate *BinaryGate
	var matches []int

	for _, g := range gates {
		locs := g.Unpatched.FindAll(data)
		if len(locs) > 0 {
			matchedGate = &g
			matches = locs
			break
		}
	}

	if matchedGate == nil {
		return "", fmt.Errorf("eligibility signature not found in %s", filepath.Base(binPath))
	}

	if _, err := BackupFile(binPath); err != nil {
		return "", fmt.Errorf("backup failed: %w", err)
	}

	patchedData := make([]byte, len(data))
	copy(patchedData, data)

	for _, idx := range matches {
		targetIdx := idx + matchedGate.Offset
		copy(patchedData[targetIdx:targetIdx+len(matchedGate.Replacement)], matchedGate.Replacement)
	}

	if err := AtomicWrite(binPath, patchedData); err != nil {
		return "", fmt.Errorf("atomic write failed: %w", err)
	}

	return fmt.Sprintf("Successfully patched %s (%d occurrences)", filepath.Base(binPath), len(matches)), nil
}
