package patcher

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var postInstallHooks []func(logger LogFunc)

// RegisterPostInstallHook registers a hook executed during wrapper installation
func RegisterPostInstallHook(h func(logger LogFunc)) {
	postInstallHooks = append(postInstallHooks, h)
}

// WrapperManager manages transparent flag injection wrappers around language_server
type WrapperManager struct {
	binDir  string
	binName string
	logger  LogFunc
}

// BackupFileName returns the backup file name for a given binary name
func BackupFileName(binName string) string {
	if runtime.GOOS == "windows" {
		if strings.HasSuffix(strings.ToLower(binName), ".exe") {
			return strings.TrimSuffix(binName, filepath.Ext(binName)) + ".bin.exe"
		}
		return binName + ".bin.exe"
	}
	return binName + ".bin"
}

// NewWrapperManager creates a new WrapperManager for a given binary directory or executable path
func NewWrapperManager(pathOrDir string, logger LogFunc) *WrapperManager {
	if logger == nil {
		logger = defaultFlagsLogger
	}

	binDir := pathOrDir
	binName := ""

	if fi, err := os.Stat(pathOrDir); err == nil && !fi.IsDir() {
		// pathOrDir is an existing file
		binDir = filepath.Dir(pathOrDir)
		binName = filepath.Base(pathOrDir)
	} else if filepath.Ext(pathOrDir) != "" && !strings.HasSuffix(pathOrDir, string(filepath.Separator)) {
		// Could be a file path that doesn't exist yet or is currently moved/backed up
		dir := filepath.Dir(pathOrDir)
		base := filepath.Base(pathOrDir)
		if fiDir, errDir := os.Stat(dir); errDir == nil && fiDir.IsDir() {
			binDir = dir
			binName = base
		}
	}

	// Normalize if binName was passed with backup suffixes or polluted with repeated .bin
	for strings.Contains(binName, ".bin") {
		binName = strings.ReplaceAll(binName, ".bin", "")
	}
	if runtime.GOOS == "windows" {
		if !strings.HasSuffix(strings.ToLower(binName), ".exe") {
			binName += ".exe"
		}
		// On Windows, proactively heal any damaged/missing core binaries in binDir
		HealWindowsBinaries(binDir, logger)
	} else if runtime.GOOS == "darwin" {
		// On macOS, proactively heal any damaged/missing core binaries or invalid script wrappers in binDir
		HealDarwinBinaries(binDir, logger)
	}

	// If binName was not explicitly identified from a file path:
	if binName == "" {
		// 1. Dynamic detection for any existing language_server binary in binDir
		if detected := DetectLanguageServerInDir(binDir); detected != "" {
			binName = filepath.Base(detected)
		} else {
			// 2. Check if a backup (.bin) already exists in binDir
			entries, _ := os.ReadDir(binDir)
			for _, entry := range entries {
				if entry.IsDir() {
					continue
				}
				name := entry.Name()
				low := strings.ToLower(name)
				if strings.HasPrefix(low, "language_server") && strings.HasSuffix(low, ".bin") {
					binName = strings.TrimSuffix(name, ".bin")
					break
				}
			}

			// 3. Search for architecture-aware binary names
			if binName == "" {
				for _, name := range GetLanguageServerBinaryNames() {
					target := filepath.Join(binDir, name)
					if fi, err := os.Stat(target); err == nil && !fi.IsDir() {
						binName = name
						break
					}
					backup := filepath.Join(binDir, BackupFileName(name))
					if fi, err := os.Stat(backup); err == nil && !fi.IsDir() {
						binName = name
						break
					}
				}
			}

			// 4. Fallback default
			if binName == "" {
				if runtime.GOOS == "windows" {
					binName = "language_server.exe"
				} else {
					binName = "language_server"
				}
			}
		}
	}

	return &WrapperManager{
		binDir:  binDir,
		binName: binName,
		logger:  logger,
	}
}

// GetBinaryName returns the resolved binary name (e.g. language_server_linux_x64)
func (m *WrapperManager) GetBinaryName() string {
	return m.binName
}

// GetBinaryDir returns the directory containing the binary
func (m *WrapperManager) GetBinaryDir() string {
	return m.binDir
}

// ServerBin returns the full path of the server binary or wrapper
func (m *WrapperManager) ServerBin() string {
	return filepath.Join(m.binDir, m.binName)
}

// OriginalBin returns the full path of the backed-up original binary
func (m *WrapperManager) OriginalBin() string {
	return filepath.Join(m.binDir, BackupFileName(m.binName))
}

// IsWrapperInstalled checks whether a wrapper script is active and original binary backed up
func (m *WrapperManager) IsWrapperInstalled() bool {
	if runtime.GOOS == "windows" {
		return false
	}

	origInfo, err := os.Stat(m.OriginalBin())
	if err != nil || origInfo.IsDir() {
		return false
	}

	serverInfo, err := os.Stat(m.ServerBin())
	if err != nil || serverInfo.IsDir() {
		return false
	}

	return true
}

// InstallWrapper backs up the original binary and installs the flag-forwarding wrapper
func (m *WrapperManager) InstallWrapper(opts EnterpriseFlagsOptions) error {
	serverBin := m.ServerBin()
	originalBin := m.OriginalBin()

	if runtime.GOOS == "windows" {
		// On Windows, replacing language_server.exe with a shell/batch script causes CreateProcessW
		// to fail with ERROR_BAD_EXE_FORMAT (193), resulting in blank Electron windows or missing binary popups.
		// Enterprise flags & SSH spoofing are securely forwarded via desktop launcher environment variables.
		// NEVER rename language_server.exe away on Windows!
		HealWindowsBinaries(m.binDir, m.logger)
		m.logger("Windows environment: Enterprise flags & SSH spoofing are securely forwarded via desktop launcher; keeping native PE binary intact at %s", serverBin)
		return nil
	}

	if runtime.GOOS == "darwin" {
		// On macOS, replacing language_server with a shell script causes Apple Mobile File Integrity (AMFI)
		// and Gatekeeper to abort the process on launch with signal=SIGTERM (Bug #37).
		// Enterprise flags & SSH spoofing are securely forwarded via Desktop launcher (Antigravity-Proxy.command);
		// NEVER rename language_server away on macOS! Keep the native Mach-O binary intact and ad-hoc codesigned.
		HealDarwinBinaries(m.binDir, m.logger)
		m.logger("macOS environment: Verified pristine native Mach-O binary at %s; keeping native binary intact with ad-hoc signature to prevent AMFI SIGTERM", serverBin)
		return nil
	}

	// 1. If original backup doesn't exist, backup the real binary
	if _, err := os.Stat(originalBin); os.IsNotExist(err) {
		if _, err := os.Stat(serverBin); err != nil {
			return fmt.Errorf("%s not found at %s: %w", m.binName, serverBin, err)
		}
		m.logger("Backing up original binary to %s", originalBin)
		if err := os.Rename(serverBin, originalBin); err != nil {
			return FormatAppManagementError("backing up server binary for wrapper", originalBin, err)
		}
	}

	if runtime.GOOS == "darwin" {
		_ = ResignDarwinBinary(originalBin)
	}

	// 2. Build flags
	flags := BuildServerFlags(opts, m.logger)
	flagsStr := strings.Join(flags, " ")

	// Execute post-install hooks (e.g. bi-directional token bridge synchronization)
	for _, hook := range postInstallHooks {
		hook(m.logger)
	}

	// 3. Create wrapper script with remote environment spoofing (switches language_server to file-token mode & reduces cloud fraud friction)
	backupName := BackupFileName(m.binName)
	scriptContent := fmt.Sprintf("#!/usr/bin/env bash\nDIR=\"$(cd \"$(dirname \"${BASH_SOURCE[0]}\")\" && pwd)\"\nexport SSH_CONNECTION=\"127.0.0.1 1234 127.0.0.1 22\"\nexport SSH_CLIENT=\"127.0.0.1 1234 22\"\nexec \"$DIR/%s\" %s \"$@\"\n", backupName, flagsStr)

	if err := os.WriteFile(serverBin, []byte(scriptContent), 0755); err != nil {
		return FormatAppManagementError("writing wrapper script", serverBin, err)
	}

	m.logger("Successfully installed enterprise wrapper at %s", serverBin)
	return nil
}

// RemoveWrapper restores the original binary and removes the wrapper
func (m *WrapperManager) RemoveWrapper() error {
	if runtime.GOOS == "windows" {
		HealWindowsBinaries(m.binDir, m.logger)
		m.logger("Windows environment: verified pristine native PE binary at %s", m.ServerBin())
		return nil
	}
	if runtime.GOOS == "darwin" {
		HealDarwinBinaries(m.binDir, m.logger)
		m.logger("macOS environment: verified pristine native Mach-O binary at %s", m.ServerBin())
		return nil
	}

	serverBin := m.ServerBin()
	originalBin := m.OriginalBin()

	if _, err := os.Stat(originalBin); os.IsNotExist(err) {
		m.logger("Original binary not found; wrapper is not currently installed")
		return nil
	}

	if err := os.Remove(serverBin); err != nil && !os.IsNotExist(err) {
		return FormatAppManagementError("removing wrapper script", serverBin, err)
	}
	if err := os.Rename(originalBin, serverBin); err != nil {
		return FormatAppManagementError("restoring original binary from wrapper", serverBin, err)
	}

	m.logger("Successfully removed wrapper and restored original binary at %s", serverBin)
	return nil
}

// HealWindowsBinaries detects and repairs any broken, missing, or renamed language_server.exe on Windows.
// Resolves Bug #40, Bug #43, and Bug #36 caused by legacy .bin.exe renaming loops.
func HealWindowsBinaries(binDir string, logger LogFunc) {
	if binDir == "" {
		return
	}
	if logger == nil {
		logger = defaultFlagsLogger
	}

	serverBin := filepath.Join(binDir, "language_server.exe")

	// 1. If language_server.exe exists but is a script/batch file (<4KB and text), remove it
	if data, err := os.ReadFile(serverBin); err == nil && len(data) < 4096 {
		content := strings.ToLower(string(data))
		if strings.HasPrefix(content, "@echo") || strings.HasPrefix(content, "#!") {
			_ = os.Remove(serverBin)
			logger("Removed invalid script wrapper at %s", serverBin)
		}
	}

	// 2. Scan binDir for candidate language_server binaries (including .bin.exe, .bin.bin.exe, etc.)
	entries, err := os.ReadDir(binDir)
	if err != nil {
		return
	}

	var bestCandidate string
	var bestSize int64

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		low := strings.ToLower(name)
		if strings.HasPrefix(low, "language_server") && strings.HasSuffix(low, ".exe") {
			fullPath := filepath.Join(binDir, name)
			fi, errStat := entry.Info()
			if errStat != nil {
				continue
			}

			// Real executable binary candidate
			if fi.Size() > bestSize {
				bestSize = fi.Size()
				bestCandidate = fullPath
			}
		}
	}

	// 3. If serverBin is missing or abnormally small (<100KB) and we found a valid PE binary candidate, restore it
	fiServer, errServer := os.Stat(serverBin)
	serverNeedsRestore := errServer != nil || fiServer.Size() < 100*1024

	if serverNeedsRestore && bestCandidate != "" && bestCandidate != serverBin {
		logger("Auto-healing Windows binary: restoring %s -> %s (size: %d bytes)", filepath.Base(bestCandidate), serverBin, bestSize)
		if errRename := os.Rename(bestCandidate, serverBin); errRename != nil {
			if data, errRead := os.ReadFile(bestCandidate); errRead == nil {
				_ = os.WriteFile(serverBin, data, 0755)
			}
		}
	}

	// 4. Clean up any leftover duplicate .bin*.exe files to prevent recursive rename pollution
	if fi, err := os.Stat(serverBin); err == nil && fi.Size() > 100*1024 {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			name := entry.Name()
			low := strings.ToLower(name)
			if strings.HasPrefix(low, "language_server") && strings.Contains(low, ".bin") && strings.HasSuffix(low, ".exe") {
				target := filepath.Join(binDir, name)
				if target != serverBin {
					_ = os.Remove(target)
				}
			}
		}
	}
}

// HealDarwinBinaries detects and repairs any broken, missing, or script-wrapped language_server binaries on macOS.
// Resolves Bug #37 where legacy shell script wrappers cause Apple AMFI kernel SIGTERM aborts and "signature not found" failures.
func HealDarwinBinaries(binDir string, logger LogFunc) {
	if binDir == "" {
		return
	}
	if logger == nil {
		logger = defaultFlagsLogger
	}

	serverBin := filepath.Join(binDir, "language_server")

	// 1. If language_server exists but is a script (<4KB and text starting with #!), remove it
	if data, err := os.ReadFile(serverBin); err == nil && len(data) < 4096 {
		content := strings.TrimSpace(string(data))
		if strings.HasPrefix(content, "#!") {
			_ = os.Remove(serverBin)
			logger("Removed invalid script wrapper at %s", serverBin)
		}
	}

	// 2. Scan binDir for candidate Mach-O binaries (including language_server.bin, architecture binaries, etc.)
	entries, err := os.ReadDir(binDir)
	if err != nil {
		return
	}

	var bestCandidate string
	var bestSize int64

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		low := strings.ToLower(name)
		if strings.HasPrefix(low, "language_server") {
			fullPath := filepath.Join(binDir, name)
			fi, errStat := entry.Info()
			if errStat != nil {
				continue
			}

			// Exclude tiny text scripts
			if fi.Size() < 4096 {
				continue
			}

			// Real executable binary candidate
			if fi.Size() > bestSize {
				bestSize = fi.Size()
				bestCandidate = fullPath
			}
		}
	}

	// 3. If serverBin is missing or abnormally small (<100KB) and we found a valid binary candidate, restore it
	fiServer, errServer := os.Stat(serverBin)
	serverNeedsRestore := errServer != nil || fiServer.Size() < 100*1024

	if serverNeedsRestore && bestCandidate != "" && bestCandidate != serverBin {
		logger("Auto-healing macOS binary: restoring %s -> %s (size: %d bytes)", filepath.Base(bestCandidate), serverBin, bestSize)
		if errRename := os.Rename(bestCandidate, serverBin); errRename != nil {
			if data, errRead := os.ReadFile(bestCandidate); errRead == nil {
				_ = os.WriteFile(serverBin, data, 0755)
			}
		}
	}

	// 4. Clean up any leftover duplicate .bin files (e.g. language_server.bin, language_server.bin.bin)
	if fi, err := os.Stat(serverBin); err == nil && fi.Size() > 100*1024 {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			name := entry.Name()
			low := strings.ToLower(name)
			if strings.HasPrefix(low, "language_server") && strings.Contains(low, ".bin") {
				target := filepath.Join(binDir, name)
				if target != serverBin {
					_ = os.Remove(target)
				}
			}
		}

		// 5. Ensure valid ad-hoc code signature and quarantine removal on macOS
		if runtime.GOOS == "darwin" {
			_ = ResignDarwinBinary(serverBin)
		}
	}
}

// HealPlatformBinaries detects and repairs damaged, missing, or script-wrapped language_server binaries
// across Windows and macOS before any patching or inspection operations.
func HealPlatformBinaries(binDir string, logger LogFunc) {
	if binDir == "" {
		return
	}
	if runtime.GOOS == "windows" {
		HealWindowsBinaries(binDir, logger)
	} else if runtime.GOOS == "darwin" {
		HealDarwinBinaries(binDir, logger)
	}
}
