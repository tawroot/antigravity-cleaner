package patcher

import (
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
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

	cliGateARM64_W4 = func() BinaryGate {
		u, _ := ParsePattern("?? ?? ?? b4 04 20 40 39 ?? ?? ?? 37 04 94 44 a9")
		p, _ := ParsePattern("?? ?? ?? b4 24 00 80 52 ?? ?? ?? 37 04 94 44 a9")
		r, _ := hex.DecodeString("24008052")
		return BinaryGate{
			Name:        "agy eligibility check bypass (arm64, w4)",
			Unpatched:   u,
			Patched:     p,
			Replacement: r,
			Offset:      4,
		}
	}()

	cliGateARM64_W2 = func() BinaryGate {
		u, _ := ParsePattern("?? ?? ?? b5 ?? ?? ?? b4 02 20 40 39 ?? ?? ?? 37")
		p, _ := ParsePattern("?? ?? ?? b5 ?? ?? ?? b4 22 00 80 52 ?? ?? ?? 37")
		r, _ := hex.DecodeString("22008052")
		return BinaryGate{
			Name:        "agy eligibility check bypass (arm64, w2)",
			Unpatched:   u,
			Patched:     p,
			Replacement: r,
			Offset:      8,
		}
	}()

	cliGateAuthX64 = func() BinaryGate {
		u, _ := ParsePattern("80 78 08 00 74 ?? 48 8b ?? 24 ?? 48 89 ?? 60")
		p, _ := ParsePattern("c6 40 08 01 90 90 48 8b ?? 24 ?? 48 89 ?? 60")
		r, _ := hex.DecodeString("c64008019090")
		return BinaryGate{
			Name:        "agy internal hasValidAuth (x64)",
			Unpatched:   u,
			Patched:     p,
			Replacement: r,
			Offset:      0,
		}
	}()

	cliGateAuthARM64 = func() BinaryGate {
		u, _ := ParsePattern("03 20 40 39 ?? ?? ?? 36 ?? ?? ?? ?? 03 10 06 a9")
		p, _ := ParsePattern("23 00 80 52 03 20 00 39 ?? ?? ?? ?? 03 10 06 a9")
		r, _ := hex.DecodeString("2300805203200039")
		return BinaryGate{
			Name:        "agy internal hasValidAuth (arm64)",
			Unpatched:   u,
			Patched:     p,
			Replacement: r,
			Offset:      0,
		}
	}()
)

// FindAgy dynamically searches for the agy binary in PATH, user environments, and dynamic runtime locations
func FindAgy() string {
	names := []string{"agy", "agy.exe", "antigravity", "antigravity.exe"}
	for _, n := range names {
		if p, err := exec.LookPath(n); err == nil && p != "" {
			return p
		}
	}

	home, _ := os.UserHomeDir()
	seen := make(map[string]bool)
	var candidates []string

	add := func(p string) {
		if p == "" {
			return
		}
		clean := filepath.Clean(p)
		if !seen[clean] {
			seen[clean] = true
			if fi, err := os.Stat(clean); err == nil && !fi.IsDir() {
				candidates = append(candidates, clean)
			}
		}
	}

	// 1. Dynamic check across all PATH directories
	for _, dir := range filepath.SplitList(os.Getenv("PATH")) {
		for _, name := range []string{"agy", "agy.exe"} {
			add(filepath.Join(dir, name))
		}
	}

	// 2. User environment & standard toolchains
	if home != "" {
		for _, sub := range []string{
			filepath.Join(".local", "bin"),
			filepath.Join(".gemini", "bin"),
			filepath.Join("bin"),
			filepath.Join(".antigravity", "bin"),
		} {
			for _, name := range []string{"agy", "agy.exe"} {
				add(filepath.Join(home, sub, name))
			}
		}
	}

	// 3. Windows environment locations
	if localAppData := os.Getenv("LOCALAPPDATA"); localAppData != "" {
		add(filepath.Join(localAppData, "agy", "bin", "agy.exe"))
		add(filepath.Join(localAppData, "Programs", "agy", "agy.exe"))
	}
	if userProfile := os.Getenv("USERPROFILE"); userProfile != "" {
		add(filepath.Join(userProfile, "scoop", "apps", "agy", "current", "agy.exe"))
	}
	if progFiles := os.Getenv("ProgramFiles"); progFiles != "" {
		add(filepath.Join(progFiles, "agy", "agy.exe"))
	}

	if len(candidates) > 0 {
		return candidates[0]
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

	gates := []BinaryGate{
		cliGateX64,
		cliGateAuthX64,
		cliGateARM64,
		cliGateARM64_W4,
		cliGateARM64_W2,
		cliGateAuthARM64,
	}
	if runtime.GOARCH == "arm64" {
		gates = []BinaryGate{
			cliGateARM64,
			cliGateARM64_W4,
			cliGateARM64_W2,
			cliGateAuthARM64,
			cliGateX64,
			cliGateAuthX64,
		}
	}

	patchedData := make([]byte, len(data))
	copy(patchedData, data)
	totalPatches := 0
	var appliedNames []string

	for _, g := range gates {
		locs := g.Unpatched.FindAll(patchedData)
		if len(locs) > 0 {
			for _, idx := range locs {
				targetIdx := idx + g.Offset
				copy(patchedData[targetIdx:targetIdx+len(g.Replacement)], g.Replacement)
				totalPatches++
			}
			appliedNames = append(appliedNames, fmt.Sprintf("%s (%d)", g.Name, len(locs)))
		}
	}

	if totalPatches == 0 {
		for _, g := range gates {
			if g.Patched.Find(data) != -1 {
				return fmt.Sprintf("Already patched: %s (%s)", filepath.Base(binPath), g.Name), nil
			}
		}
		return "", fmt.Errorf("eligibility signature not found in %s", filepath.Base(binPath))
	}

	if _, err := BackupFile(binPath); err != nil {
		return "", fmt.Errorf("backup failed: %w", err)
	}

	if err := AtomicWrite(binPath, patchedData); err != nil {
		return "", fmt.Errorf("atomic write failed: %w", err)
	}

	return fmt.Sprintf("Successfully patched %s [%s]", filepath.Base(binPath), strings.Join(appliedNames, ", ")), nil
}
