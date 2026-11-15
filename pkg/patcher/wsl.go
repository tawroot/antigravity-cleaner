package patcher

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// WslDistroInfo holds discovered metadata about a WSL distribution
type WslDistroInfo struct {
	Name    string
	RootPath string // e.g. \\wsl.localhost\Ubuntu
}

// WslReport provides structured results of a WSL patch operation
type WslReport struct {
	DistrosFound    []string
	ServersPatched  []string
	WslConfigUpdated bool
	ShellRcsUpdated []string
}

// FindWslDistroRoots dynamically discovers all installed WSL distribution roots
// without relying on hardcoded distribution names like "Ubuntu".
func FindWslDistroRoots(logger LogFunc) []WslDistroInfo {
	if runtime.GOOS != "windows" {
		return nil
	}
	if logger == nil {
		logger = defaultLogger
	}

	var distros []WslDistroInfo
	seen := make(map[string]bool)

	addDistro := func(name, path string) {
		name = strings.TrimSpace(name)
		path = strings.TrimSpace(path)
		if name == "" || path == "" {
			return
		}
		clean := filepath.Clean(path)
		if !seen[strings.ToLower(clean)] {
			if fi, err := os.Stat(clean); err == nil && fi.IsDir() {
				seen[strings.ToLower(clean)] = true
				distros = append(distros, WslDistroInfo{
					Name:     name,
					RootPath: clean,
				})
				logger("Discovered active WSL distribution: %s at %s", name, clean)
			}
		}
	}

	// 1. Enumerate \\wsl.localhost\ (Modern Windows 10/11)
	if entries, err := os.ReadDir(`\\wsl.localhost`); err == nil {
		for _, e := range entries {
			if e.IsDir() {
				addDistro(e.Name(), filepath.Join(`\\wsl.localhost`, e.Name()))
			}
		}
	}

	// 2. Fallback to legacy \\wsl$\ (Windows 10 1903+)
	if entries, err := os.ReadDir(`\\wsl$`); err == nil {
		for _, e := range entries {
			if e.IsDir() {
				addDistro(e.Name(), filepath.Join(`\\wsl$`, e.Name()))
			}
		}
	}

	// 3. Fallback: Query registry via reg.exe for registered distributions
	// HKCU\Software\Microsoft\Windows\CurrentVersion\Lxss
	cmd := exec.Command("reg", "query", `HKCU\Software\Microsoft\Windows\CurrentVersion\Lxss`, "/s")
	if out, err := cmd.Output(); err == nil {
		scanner := bufio.NewScanner(bytes.NewReader(out))
		var currentDistroName string
		var currentBasePath string

		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if strings.HasPrefix(line, "HKEY_CURRENT_USER") {
				if currentDistroName != "" {
					addDistro(currentDistroName, filepath.Join(`\\wsl.localhost`, currentDistroName))
					addDistro(currentDistroName, filepath.Join(`\\wsl$`, currentDistroName))
				}
				currentDistroName = ""
				currentBasePath = ""
			} else if strings.HasPrefix(line, "DistributionName") {
				fields := strings.Fields(line)
				if len(fields) >= 3 {
					currentDistroName = fields[len(fields)-1]
				}
			} else if strings.HasPrefix(line, "BasePath") {
				fields := strings.Fields(line)
				if len(fields) >= 3 {
					currentBasePath = fields[len(fields)-1]
					_ = currentBasePath
				}
			}
		}
		if currentDistroName != "" {
			addDistro(currentDistroName, filepath.Join(`\\wsl.localhost`, currentDistroName))
			addDistro(currentDistroName, filepath.Join(`\\wsl$`, currentDistroName))
		}
	}

	return distros
}

// FindAllWslLanguageServers recursively searches all user homes in all WSL distros
// for Linux language_server ELF binaries.
func FindAllWslLanguageServers(logger LogFunc) []string {
	if runtime.GOOS != "windows" {
		return nil
	}
	if logger == nil {
		logger = defaultLogger
	}

	distros := FindWslDistroRoots(logger)
	if len(distros) == 0 {
		return nil
	}

	var results []string
	seen := make(map[string]bool)

	addServer := func(p string) {
		clean := filepath.Clean(p)
		if !seen[strings.ToLower(clean)] {
			if fi, err := os.Stat(clean); err == nil && !fi.IsDir() && fi.Size() > 10*1024*1024 {
				// Verify ELF header
				if isElfBinary(clean) {
					seen[strings.ToLower(clean)] = true
					results = append(results, clean)
					logger("Found WSL Linux language_server ELF binary: %s (%d MB)", clean, fi.Size()/(1024*1024))
				}
			}
		}
	}

	for _, d := range distros {
		// Potential user home roots in Linux: /home/<user> and /root
		var homeDirs []string

		// Scan /home
		linuxHome := filepath.Join(d.RootPath, "home")
		if entries, err := os.ReadDir(linuxHome); err == nil {
			for _, e := range entries {
				if e.IsDir() {
					homeDirs = append(homeDirs, filepath.Join(linuxHome, e.Name()))
				}
			}
		}
		// Scan /root
		rootHome := filepath.Join(d.RootPath, "root")
		if fi, err := os.Stat(rootHome); err == nil && fi.IsDir() {
			homeDirs = append(homeDirs, rootHome)
		}

		// Common Antigravity & server directories within Linux user homes
		serverSubdirs := []string{
			".antigravity",
			".antigravity-server",
			".gemini",
			".vscode-server",
			".local/share/antigravity",
			".local/share/antigravity-ide",
		}

		for _, h := range homeDirs {
			for _, sub := range serverSubdirs {
				targetDir := filepath.Join(h, filepath.FromSlash(sub))
				if fi, err := os.Stat(targetDir); err == nil && fi.IsDir() {
					// Walk targetDir for any language_server binary
					_ = filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
						if err != nil || info == nil {
							return nil
						}
						name := strings.ToLower(info.Name())
						// Avoid backups and intermediate files
						if strings.HasSuffix(name, ".agybak") || strings.HasSuffix(name, ".old") || strings.HasSuffix(name, ".bin") {
							return nil
						}
						if name == "language_server" || strings.HasPrefix(name, "language_server_") {
							addServer(path)
						}
						return nil
					})
				}
			}
		}
	}

	return results
}

// isElfBinary checks if a file begins with the standard 4-byte Linux ELF magic: 0x7F, 'E', 'L', 'F'
func isElfBinary(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	header := make([]byte, 4)
	if _, err := f.Read(header); err != nil {
		return false
	}
	return header[0] == 0x7f && header[1] == 'E' && header[2] == 'L' && header[3] == 'F'
}

// DetectActiveLocalPort dynamically detects the fastest open local proxy port on 127.0.0.1
func DetectActiveLocalPort() (int, string) {
	candidates := []struct {
		port   int
		scheme string
	}{
		{10808, "socks5"},
		{2080, "socks5"},
		{7890, "socks5"},
		{1080, "socks5"},
		{10809, "http"},
		{2081, "http"},
	}

	for _, c := range candidates {
		addr := fmt.Sprintf("127.0.0.1:%d", c.port)
		conn, err := net.DialTimeout("tcp", addr, 60*time.Millisecond)
		if err == nil {
			conn.Close()
			return c.port, c.scheme
		}
	}

	return 10808, "socks5"
}

// ConfigureWslMirroredNetworking configures %USERPROFILE%\.wslconfig with
// networkingMode=mirrored, dnsTunneling=true, and autoProxy=true,
// preserving all existing user configurations without data destruction.
func ConfigureWslMirroredNetworking(logger LogFunc) (bool, error) {
	if runtime.GOOS != "windows" {
		return false, nil
	}
	if logger == nil {
		logger = defaultLogger
	}

	userProfile := os.Getenv("USERPROFILE")
	if userProfile == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return false, err
		}
		userProfile = home
	}

	wslConfigPath := filepath.Join(userProfile, ".wslconfig")

	var lines []string
	if data, err := os.ReadFile(wslConfigPath); err == nil {
		scanner := bufio.NewScanner(bytes.NewReader(data))
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}

	// Key/value pairs to enforce in [wsl2]
	targetKeys := map[string]string{
		"networkingmode": "mirrored",
		"dnstunneling":   "true",
		"autoproxy":      "true",
	}

	wsl2SectionIdx := -1
	keysFound := make(map[string]bool)

	var outputLines []string
	inWsl2Section := false

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]") {
			sectionName := strings.ToLower(trimmed[1 : len(trimmed)-1])
			if sectionName == "wsl2" {
				inWsl2Section = true
				wsl2SectionIdx = i
			} else {
				inWsl2Section = false
			}
			outputLines = append(outputLines, line)
			continue
		}

		if inWsl2Section && strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			k := strings.ToLower(strings.TrimSpace(parts[0]))
			if newVal, isTarget := targetKeys[k]; isTarget {
				keysFound[k] = true
				outputLines = append(outputLines, fmt.Sprintf("%s=%s", parts[0], newVal))
				continue
			}
		}

		outputLines = append(outputLines, line)
	}

	// If [wsl2] section was absent, add it
	if wsl2SectionIdx == -1 {
		if len(outputLines) > 0 && strings.TrimSpace(outputLines[len(outputLines)-1]) != "" {
			outputLines = append(outputLines, "")
		}
		outputLines = append(outputLines, "[wsl2]")
		outputLines = append(outputLines, "networkingMode=mirrored")
		outputLines = append(outputLines, "dnsTunneling=true")
		outputLines = append(outputLines, "autoProxy=true")
	} else {
		// If [wsl2] existed but some target keys were missing, append them into [wsl2]
		var finalLines []string
		inSection := false
		for _, line := range outputLines {
			finalLines = append(finalLines, line)
			trimmed := strings.TrimSpace(line)
			if strings.EqualFold(trimmed, "[wsl2]") {
				inSection = true
				for k, v := range targetKeys {
					if !keysFound[k] {
						// Format original casing
						formattedKey := k
						if k == "networkingmode" {
							formattedKey = "networkingMode"
						} else if k == "dnstunneling" {
							formattedKey = "dnsTunneling"
						} else if k == "autoproxy" {
							formattedKey = "autoProxy"
						}
						finalLines = append(finalLines, fmt.Sprintf("%s=%s", formattedKey, v))
					}
				}
			} else if strings.HasPrefix(trimmed, "[") {
				inSection = false
			}
			_ = inSection
		}
		outputLines = finalLines
	}

	content := strings.Join(outputLines, "\r\n") + "\r\n"
	if err := os.WriteFile(wslConfigPath, []byte(content), 0644); err != nil {
		return false, fmt.Errorf("failed to write %s: %w", wslConfigPath, err)
	}

	logger("Configured WSL Mirrored Networking in %s", wslConfigPath)
	return true, nil
}

// ConfigureWslShellRc idempotently injects dynamic proxy bridge variables into
// all user .bashrc and .zshrc files across all detected WSL distributions.
func ConfigureWslShellRc(distros []WslDistroInfo, port int, scheme string, logger LogFunc) []string {
	if logger == nil {
		logger = defaultLogger
	}

	const markerStart = "# >>> Antigravity Cleaner WSL Proxy Bridge >>>"
	const markerEnd = "# <<< Antigravity Cleaner WSL Proxy Bridge <<<"

	block := fmt.Sprintf(`%s
export ALL_PROXY="%s://127.0.0.1:%d"
export all_proxy="%s://127.0.0.1:%d"
export HTTP_PROXY="http://127.0.0.1:%d"
export http_proxy="http://127.0.0.1:%d"
export HTTPS_PROXY="http://127.0.0.1:%d"
export https_proxy="http://127.0.0.1:%d"
export NO_PROXY="localhost,127.0.0.1"
export no_proxy="localhost,127.0.0.1"
%s`, markerStart, scheme, port, scheme, port, port, port, port, port, markerEnd)

	var updatedFiles []string

	for _, d := range distros {
		var homeDirs []string
		linuxHome := filepath.Join(d.RootPath, "home")
		if entries, err := os.ReadDir(linuxHome); err == nil {
			for _, e := range entries {
				if e.IsDir() {
					homeDirs = append(homeDirs, filepath.Join(linuxHome, e.Name()))
				}
			}
		}
		rootHome := filepath.Join(d.RootPath, "root")
		if fi, err := os.Stat(rootHome); err == nil && fi.IsDir() {
			homeDirs = append(homeDirs, rootHome)
		}

		for _, h := range homeDirs {
			for _, rcName := range []string{".bashrc", ".zshrc"} {
				rcPath := filepath.Join(h, rcName)
				if fi, err := os.Stat(rcPath); err == nil && !fi.IsDir() {
					contentBytes, err := os.ReadFile(rcPath)
					if err != nil {
						continue
					}
					content := string(contentBytes)
					if strings.Contains(content, markerStart) && strings.Contains(content, markerEnd) {
						// Replace existing block cleanly
						startIdx := strings.Index(content, markerStart)
						endIdx := strings.Index(content, markerEnd) + len(markerEnd)
						newContent := content[:startIdx] + block + content[endIdx:]
						_ = os.WriteFile(rcPath, []byte(newContent), 0644)
						updatedFiles = append(updatedFiles, rcPath)
						logger("Updated existing proxy bridge block in %s", rcPath)
					} else {
						// Append block
						newContent := strings.TrimRight(content, "\r\n") + "\n\n" + block + "\n"
						_ = os.WriteFile(rcPath, []byte(newContent), 0644)
						updatedFiles = append(updatedFiles, rcPath)
						logger("Injected new proxy bridge block into %s", rcPath)
					}
				}
			}
		}
	}

	return updatedFiles
}

// PatchAllWsl executes the complete dynamic WSL self-healing and anti-sanction suite:
// 1. Discovers all installed distributions dynamically without hardcoding.
// 2. Discovers and patches all Linux ELF language_server binaries in all distributions.
// 3. Deploys Mirrored Networking in %USERPROFILE%\.wslconfig to connect loopback seamlessly.
// 4. Injects dynamic proxy environment into all Linux shell configuration files.
func PatchAllWsl(logger LogFunc) (WslReport, error) {
	if logger == nil {
		logger = defaultLogger
	}

	var report WslReport

	if runtime.GOOS != "windows" {
		return report, nil
	}

	distros := FindWslDistroRoots(logger)
	for _, d := range distros {
		report.DistrosFound = append(report.DistrosFound, fmt.Sprintf("%s (%s)", d.Name, d.RootPath))
	}

	// 1. Dynamic Port Discovery
	port, scheme := DetectActiveLocalPort()
	logger("Active local proxy detected for WSL bridge: %s://127.0.0.1:%d", scheme, port)

	// 2. Patch all Linux ELF language_server binaries
	servers := FindAllWslLanguageServers(logger)
	for _, s := range servers {
		if msg, err := PatchSingleLanguageServer(s); err == nil {
			report.ServersPatched = append(report.ServersPatched, fmt.Sprintf("%s [%s]", filepath.Base(s), msg))
		} else {
			logger("Warning: failed to patch WSL language_server %s: %v", s, err)
		}
	}

	// 3. Configure Mirrored Networking
	if updated, err := ConfigureWslMirroredNetworking(logger); err == nil && updated {
		report.WslConfigUpdated = true
	}

	// 4. Configure Shell RCs if any distros exist
	if len(distros) > 0 {
		report.ShellRcsUpdated = ConfigureWslShellRc(distros, port, scheme, logger)
	}

	return report, nil
}
