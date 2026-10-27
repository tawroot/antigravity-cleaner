package patcher

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// BinaryGate defines byte patterns for checking and patching binaries
type BinaryGate struct {
	Name        string
	Unpatched   *Pattern
	Patched     *Pattern
	Replacement []byte
	Offset      int
}

var (
	managerGateX64 = func() BinaryGate {
		u, _ := ParsePattern("80 78 08 00 74 ?? 48 8b ?? 24 ?? 48 89 ?? 60")
		p, _ := ParsePattern("c6 40 08 01 90 90 48 8b ?? 24 ?? 48 89 ?? 60")
		r, _ := hex.DecodeString("c64008019090")
		return BinaryGate{
			Name:        "language_server hasValidAuth=true (x64)",
			Unpatched:   u,
			Patched:     p,
			Replacement: r,
			Offset:      0,
		}
	}()

	managerGateARM64 = func() BinaryGate {
		u, _ := ParsePattern("03 20 40 39 ?? ?? ?? 36 ?? ?? ?? ?? 03 10 06 a9")
		p, _ := ParsePattern("23 00 80 52 03 20 00 39 ?? ?? ?? ?? 03 10 06 a9")
		r, _ := hex.DecodeString("2300805203200039")
		return BinaryGate{
			Name:        "language_server hasValidAuth=true (arm64)",
			Unpatched:   u,
			Patched:     p,
			Replacement: r,
			Offset:      0,
		}
	}()
)

// FindLanguageServer searches for language_server binary using process inspection and standard locations
func FindLanguageServer() string {
	return FindLanguageServerModular(nil)
}

// PatchSingleLanguageServer patches a specific language_server binary to unlock auth and region checks
func PatchSingleLanguageServer(binPath string) (string, error) {
	data, err := os.ReadFile(binPath)
	if err != nil {
		return "", fmt.Errorf("failed to read binary %s: %w", binPath, err)
	}

	gates := []BinaryGate{managerGateX64, managerGateARM64}
	if runtime.GOARCH == "arm64" {
		gates = []BinaryGate{managerGateARM64, managerGateX64}
	}

	for _, g := range gates {
		if g.Patched.Find(data) != -1 {
			return fmt.Sprintf("Already patched: %s (%s)", filepath.Base(binPath), g.Name), nil
		}
	}

	var matchedGate *BinaryGate
	var matchIdx = -1

	for _, g := range gates {
		idx := g.Unpatched.Find(data)
		if idx != -1 {
			matchedGate = &g
			matchIdx = idx
			break
		}
	}

	if matchedGate == nil {
		return "", fmt.Errorf("signature not found in %s (unsupported or already modified build)", filepath.Base(binPath))
	}

	// Create backup before modification
	if _, err := BackupFile(binPath); err != nil {
		return "", fmt.Errorf("backup failed: %w", err)
	}

	targetIdx := matchIdx + matchedGate.Offset
	patchedData := make([]byte, len(data))
	copy(patchedData, data)
	copy(patchedData[targetIdx:targetIdx+len(matchedGate.Replacement)], matchedGate.Replacement)

	if err := AtomicWrite(binPath, patchedData); err != nil {
		return "", fmt.Errorf("atomic write failed: %w", err)
	}

	return fmt.Sprintf("Successfully patched %s (%s) at offset 0x%x", filepath.Base(binPath), matchedGate.Name, targetIdx), nil
}

// PatchLanguageServer patches all detected language_server binaries to unlock auth and region checks
func PatchLanguageServer(binPath string) (string, error) {
	if binPath != "" {
		return PatchSingleLanguageServer(binPath)
	}

	targets := FindAllLanguageServers(nil)
	if len(targets) == 0 {
		return "", fmt.Errorf("language_server binary not found in standard paths")
	}

	if len(targets) == 1 {
		return PatchSingleLanguageServer(targets[0])
	}

	var patchedCount int
	var details []string
	var lastErr error

	for _, target := range targets {
		_, err := PatchSingleLanguageServer(target)
		if err == nil {
			patchedCount++
			details = append(details, filepath.Base(filepath.Dir(target))+"/"+filepath.Base(target))
		} else {
			lastErr = err
		}
	}

	if patchedCount == 0 && lastErr != nil {
		return "", lastErr
	}

	if patchedCount > 1 {
		return fmt.Sprintf("Successfully unlocked %d core binaries across installations", patchedCount), nil
	}
	if len(details) > 0 {
		return fmt.Sprintf("Successfully unlocked: %s", details[0]), nil
	}
	return "Core binary unlocked", nil
}

// CheckLanguageServerStatus checks whether the binary is patched, unpatched, or unknown
func CheckLanguageServerStatus(binPath string) string {
	if binPath == "" {
		binPath = FindLanguageServer()
		if binPath == "" {
			return "Not Installed"
		}
	}

	data, err := os.ReadFile(binPath)
	if err != nil {
		return "Unreadable"
	}

	for _, g := range []BinaryGate{managerGateX64, managerGateARM64} {
		if g.Patched.Find(data) != -1 {
			return "Patched (" + g.Name + ")"
		}
		if g.Unpatched.Find(data) != -1 {
			return "Unpatched (Official)"
		}
	}

	return "Unknown Build"
}
