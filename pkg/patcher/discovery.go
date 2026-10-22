package patcher

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// LogFunc defines a structured logging function callback
type LogFunc func(format string, args ...interface{})

// defaultLogger provides simple standard library logging with [DISCOVERY] tag
var defaultLogger LogFunc = func(format string, args ...interface{}) {
	log.Printf("[DISCOVERY] "+format, args...)
}

// FindRunningLanguageServer discovers the exact path of an active language_server process
func FindRunningLanguageServer(logger LogFunc) string {
	if logger == nil {
		logger = defaultLogger
	}

	switch runtime.GOOS {
	case "linux":
		// 1. First attempt /proc scanning directly (fastest, zero external process overhead)
		if procPath := scanLinuxProc("language_server"); procPath != "" {
			logger("Active language_server detected via /proc: %s", procPath)
			return procPath
		}
	case "darwin":
		// On macOS, probe with ps / pgrep
		if p := scanUnixPs("language_server"); p != "" {
			logger("Active language_server detected via ps: %s", p)
			return p
		}
	case "windows":
		// On Windows, probe with wmic or tasklist
		if p := scanWindowsProcess("language_server.exe"); p != "" {
			logger("Active language_server detected on Windows: %s", p)
			return p
		}
	}

	return ""
}

// FindRunningAntigravityApp discovers the exact path of an active Antigravity GUI / IDE application
func FindRunningAntigravityApp(logger LogFunc) string {
	if logger == nil {
		logger = defaultLogger
	}

	switch runtime.GOOS {
	case "linux":
		if procPath := scanLinuxProc("antigravity"); procPath != "" {
			logger("Active Antigravity app detected via /proc: %s", procPath)
			return procPath
		}
	case "darwin":
		if p := scanUnixPs("Antigravity"); p != "" {
			logger("Active Antigravity app detected via ps: %s", p)
			return p
		}
	case "windows":
		if p := scanWindowsProcess("Antigravity.exe"); p != "" {
			logger("Active Antigravity app detected on Windows: %s", p)
			return p
		}
	}

	return ""
}

// isSelfOrCleanerProcess determines if a process is the cleaner itself or related tool
func isSelfOrCleanerProcess(pid int, exePath string) bool {
	if pid > 0 {
		if pid == os.Getpid() || (os.Getppid() > 0 && pid == os.Getppid()) {
			return true
		}
	}

	if selfExe, err := os.Executable(); err == nil && selfExe != "" {
		if exePath == selfExe || filepath.Base(exePath) == filepath.Base(selfExe) {
			return true
		}
	}

	base := strings.ToLower(filepath.Base(exePath))
	base = strings.TrimSuffix(base, ".exe")

	cleanerKeywords := []string{"cleaner", "patcher", "ag-cleaner", "gui"}
	for _, kw := range cleanerKeywords {
		if strings.Contains(base, kw) {
			return true
		}
	}

	return false
}

// scanLinuxProc inspects /proc for running processes matching target substring
func scanLinuxProc(targetSubstr string) string {
	entries, err := os.ReadDir("/proc")
	if err != nil {
		return ""
	}

	targetLow := strings.ToLower(targetSubstr)

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		pid, err := strconv.Atoi(entry.Name())
		if err != nil {
			continue // Not a PID folder
		}

		exeLink := fmt.Sprintf("/proc/%d/exe", pid)
		target, err := os.Readlink(exeLink)
		if err == nil {
			if isSelfOrCleanerProcess(pid, target) {
				continue
			}
			base := strings.ToLower(filepath.Base(target))
			if strings.Contains(base, targetLow) {
				// Verify file exists
				if fi, err := os.Stat(target); err == nil && !fi.IsDir() {
					return target
				}
			}
		}

		// Also check cmdline in case exe link points to an interpreter or wrapper
		cmdlineFile := fmt.Sprintf("/proc/%d/cmdline", pid)
		if data, err := os.ReadFile(cmdlineFile); err == nil {
			args := strings.Split(string(data), "\x00")
			if len(args) > 0 {
				arg0 := args[0]
				if isSelfOrCleanerProcess(pid, arg0) {
					continue
				}
				base := strings.ToLower(filepath.Base(arg0))
				if strings.Contains(base, targetLow) {
					if fi, err := os.Stat(arg0); err == nil && !fi.IsDir() {
						return arg0
					}
				}
			}
		}
	}
	return ""
}

// scanUnixPs queries ps for unix targets
func scanUnixPs(targetSubstr string) string {
	cmd := exec.Command("ps", "-eo", "pid,comm,args")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}

	targetLow := strings.ToLower(targetSubstr)
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(strings.ToLower(line), targetLow) {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		pid, _ := strconv.Atoi(fields[0])

		// Search fields for an existing executable path
		for _, f := range fields[1:] {
			clean := strings.Trim(f, "\"'")
			if isSelfOrCleanerProcess(pid, clean) {
				continue
			}
			base := strings.ToLower(filepath.Base(clean))
			if strings.Contains(base, targetLow) {
				if fi, err := os.Stat(clean); err == nil && !fi.IsDir() {
					return clean
				}
			}
		}
	}
	return ""
}

// scanWindowsProcess queries wmic or powershell on Windows
func scanWindowsProcess(targetName string) string {
	cmd := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command",
		fmt.Sprintf("(Get-Process -Name '%s' -ErrorAction SilentlyContinue).Path", strings.TrimSuffix(targetName, ".exe")))
	out, err := cmd.Output()
	if err == nil {
		path := strings.TrimSpace(string(out))
		if path != "" {
			if fi, err := os.Stat(path); err == nil && !fi.IsDir() {
				return path
			}
		}
	}
	return ""
}

// isBackupOrTempFile returns true if fileName ends with backup or temporary suffixes
func isBackupOrTempFile(fileName string) bool {
	low := strings.ToLower(fileName)
	suffixes := []string{".agybak", ".bin", ".bak", ".tmp", ".old", ".orig"}
	for _, s := range suffixes {
		if strings.HasSuffix(low, s) {
			return true
		}
	}
	return false
}

// scoreBinaryForArchitecture returns a preference score for a binary file matching runtime.GOARCH
func scoreBinaryForArchitecture(fileName string) int {
	low := strings.ToLower(fileName)
	arch := runtime.GOARCH

	// Exact matches / high affinity
	switch arch {
	case "arm64":
		if strings.Contains(low, "arm64") || strings.Contains(low, "aarch64") {
			return 100
		}
		if strings.Contains(low, "arm") {
			return 95 // Prioritize _arm on Apple Silicon
		}
	case "amd64":
		if strings.Contains(low, "x64") || strings.Contains(low, "amd64") || strings.Contains(low, "x86_64") {
			return 100
		}
	}

	// Generic fallback names
	if low == "language_server" || low == "language_server.exe" {
		return 50
	}

	// Any other language_server binary
	if strings.HasPrefix(low, "language_server") {
		return 10
	}

	return 0
}

// DetectLanguageServerInDir dynamically scans binDir for any language_server binary matching host architecture.
// Completely eliminates hardcoded binary name dependencies across all platforms.
func DetectLanguageServerInDir(binDir string) string {
	entries, err := os.ReadDir(binDir)
	if err != nil {
		return ""
	}

	var bestMatch string
	bestScore := -1

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		lowName := strings.ToLower(name)

		if !strings.HasPrefix(lowName, "language_server") {
			continue
		}
		if isBackupOrTempFile(name) {
			continue
		}

		fullPath := filepath.Join(binDir, name)
		fi, err := os.Stat(fullPath)
		if err != nil || fi.IsDir() {
			continue
		}

		score := scoreBinaryForArchitecture(name)
		if score > bestScore {
			bestScore = score
			bestMatch = fullPath
		}
	}

	return bestMatch
}

// GetLanguageServerBinaryNames returns architecture-aware binary names across Linux, macOS, and Windows
func GetLanguageServerBinaryNames() []string {
	switch runtime.GOOS {
	case "linux":
		if runtime.GOARCH == "arm64" {
			return []string{
				"language_server_linux_arm64",
				"language_server_linux_x64",
				"language_server",
			}
		}
		return []string{
			"language_server_linux_x64",
			"language_server_linux_arm64",
			"language_server",
		}
	case "darwin":
		if runtime.GOARCH == "arm64" {
			return []string{
				"language_server_macos_arm",
				"language_server_macos_arm64",
				"language_server_darwin_arm64",
				"language_server_macos_x64",
				"language_server_darwin_x64",
				"language_server",
			}
		}
		return []string{
			"language_server_macos_x64",
			"language_server_darwin_x64",
			"language_server_macos_arm",
			"language_server_macos_arm64",
			"language_server_darwin_arm64",
			"language_server",
		}
	case "windows":
		if runtime.GOARCH == "arm64" {
			return []string{
				"language_server_windows_arm64.exe",
				"language_server_windows_x64.exe",
				"language_server.exe",
				"language_server",
			}
		}
		return []string{
			"language_server_windows_x64.exe",
			"language_server_windows_arm64.exe",
			"language_server.exe",
			"language_server",
		}
	default:
		return []string{
			"language_server_linux_x64",
			"language_server_linux_arm64",
			"language_server",
		}
	}
}

// getCandidateRoots extracts possible application root directories from an executable path
func getCandidateRoots(exePath string) []string {
	if realPath, err := filepath.EvalSymlinks(exePath); err == nil && realPath != "" {
		exePath = realPath
	}
	dir := filepath.Dir(exePath)
	roots := []string{dir}

	baseDir := filepath.Base(dir)
	if strings.EqualFold(baseDir, "bin") || strings.EqualFold(baseDir, "MacOS") {
		parent := filepath.Dir(dir)
		roots = append(roots, parent)
		roots = append(roots, filepath.Dir(parent))
	} else {
		roots = append(roots, filepath.Dir(dir))
	}
	return roots
}

// probeRootsForLanguageServer probes candidate root directories for language server binaries
func probeRootsForLanguageServer(roots []string) string {
	subpaths := []string{
		filepath.Join("resources", "app", "extensions", "antigravity", "bin"),
		filepath.Join("resources", "bin"),
		filepath.Join("Contents", "Resources", "app", "extensions", "antigravity", "bin"),
		filepath.Join("Contents", "Resources", "bin"),
		filepath.Join("Resources", "app", "extensions", "antigravity", "bin"),
		filepath.Join("Resources", "bin"),
	}

	for _, root := range roots {
		if root == "" {
			continue
		}
		for _, sub := range subpaths {
			dir := filepath.Join(root, sub)
			// 1. Dynamic Zero-Hardcode discovery inside the directory
			if found := DetectLanguageServerInDir(dir); found != "" {
				return found
			}
			// 2. Direct fallback checks for static names
			for _, binName := range GetLanguageServerBinaryNames() {
				candidate := filepath.Join(dir, binName)
				if fi, err := os.Stat(candidate); err == nil && !fi.IsDir() {
					return candidate
				}
			}
		}
	}
	return ""
}

// GetStandardLanguageServerCandidates returns all standard paths across Linux, Windows, macOS
func GetStandardLanguageServerCandidates() []string {
	home, _ := os.UserHomeDir()

	var candidates []string
	seen := make(map[string]bool)
	addCandidate := func(p string) {
		if p != "" && !seen[p] {
			seen[p] = true
			candidates = append(candidates, p)
		}
	}

	linuxRoots := []string{
		filepath.Join(home, "Applications", "Antigravity"),
		filepath.Join(home, "Applications", "antigravity"),
		filepath.Join(home, "Applications", "Antigravity-x64"),
		"/usr/share/antigravity",
		"/usr/share/antigravity-ide",
		"/opt/antigravity",
		"/opt/Antigravity",
		"/opt/antigravity-ide",
		"/opt/Antigravity IDE",
		filepath.Join(home, ".local", "share", "antigravity"),
		filepath.Join(home, "antigravity"),
	}

	linuxBins := []string{
		"language_server_linux_x64",
		"language_server_linux_arm64",
		"language_server",
	}

	linuxSubdirs := []string{
		filepath.Join("resources", "app", "extensions", "antigravity", "bin"),
		filepath.Join("resources", "bin"),
	}

	// 1. Linux RPM & standard candidates
	for _, root := range linuxRoots {
		for _, sub := range linuxSubdirs {
			for _, b := range linuxBins {
				addCandidate(filepath.Join(root, sub, b))
			}
		}
	}

	// 2. Legacy & home relative paths (ensure classic paths preserved)
	addCandidate(filepath.Join(home, "antigravity", "resources", "bin", "language_server"))
	addCandidate(filepath.Join(home, "antigravity", "resources", "bin", "language_server.exe"))

	// 3. Windows candidate roots
	winRoots := []string{
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Antigravity"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "antigravity"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Antigravity IDE"),
		filepath.Join(os.Getenv("PROGRAMFILES"), "Antigravity"),
		filepath.Join(os.Getenv("ProgramFiles(x86)"), "Antigravity"),
		filepath.Join(home, "antigravity"),
	}

	winBins := []string{
		"language_server_windows_x64.exe",
		"language_server_windows_arm64.exe",
		"language_server.exe",
		"language_server",
	}

	winSubdirs := []string{
		filepath.Join("resources", "app", "extensions", "antigravity", "bin"),
		filepath.Join("resources", "bin"),
	}

	for _, root := range winRoots {
		if root == "" {
			continue
		}
		for _, sub := range winSubdirs {
			for _, b := range winBins {
				addCandidate(filepath.Join(root, sub, b))
			}
		}
	}

	// 4. macOS candidate roots
	macRoots := []string{
		"/Applications/Antigravity.app/Contents",
		"/Applications/Antigravity IDE.app/Contents",
		filepath.Join(home, "Applications", "Antigravity.app", "Contents"),
		filepath.Join(home, "Applications", "Antigravity IDE.app", "Contents"),
	}

	macBins := []string{
		"language_server_macos_arm",
		"language_server_macos_arm64",
		"language_server_macos_x64",
		"language_server_darwin_arm64",
		"language_server_darwin_x64",
		"language_server",
	}

	macSubdirs := []string{
		filepath.Join("Resources", "app", "extensions", "antigravity", "bin"),
		filepath.Join("Resources", "bin"),
	}

	for _, root := range macRoots {
		for _, sub := range macSubdirs {
			for _, b := range macBins {
				addCandidate(filepath.Join(root, sub, b))
			}
		}
	}

	return candidates
}

// findMacAppsWithMdfind searches for Antigravity installations across macOS using Spotlight
func findMacAppsWithMdfind() []string {
	if runtime.GOOS != "darwin" {
		return nil
	}
	cmd := exec.Command("mdfind", "kMDItemFSName == 'Antigravity*.app'")
	out, err := cmd.Output()
	if err != nil {
		return nil
	}
	var apps []string
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			apps = append(apps, line)
		}
	}
	return apps
}

// FindLanguageServerModular searches for language_server with live process probing and full logging
func FindLanguageServerModular(logger LogFunc) string {
	if logger == nil {
		logger = defaultLogger
	}

	// 1. Check custom environment override
	if envPath := os.Getenv("ANTIGRAVITY_LANGUAGE_SERVER_PATH"); envPath != "" {
		if fi, err := os.Stat(envPath); err == nil && !fi.IsDir() {
			logger("Using custom override from ANTIGRAVITY_LANGUAGE_SERVER_PATH: %s", envPath)
			return envPath
		}
	}

	// 2. Probe active running process (Highest accuracy, works with any custom directory)
	if runningPath := FindRunningLanguageServer(logger); runningPath != "" {
		logger("Selected active language_server: %s", runningPath)
		return runningPath
	}

	// 3. If language_server is not running directly, check if Antigravity app is running
	if runningApp := FindRunningAntigravityApp(logger); runningApp != "" {
		roots := getCandidateRoots(runningApp)
		if found := probeRootsForLanguageServer(roots); found != "" {
			logger("Deduced language_server path from running app: %s", found)
			return found
		}
	}

	// 3.5. Dynamic PATH discovery via exec.LookPath (Works even when Antigravity is NOT running!)
	lookNames := []string{"antigravity", "antigravity.exe", "antigravity-ide"}
	for _, name := range lookNames {
		if p, err := exec.LookPath(name); err == nil && p != "" {
			roots := getCandidateRoots(p)
			if found := probeRootsForLanguageServer(roots); found != "" {
				logger("Discovered language_server via PATH (%s): %s", p, found)
				return found
			}
		}
	}

	// 3.6. Dynamic Spotlight discovery on macOS (finds app even if installed in non-standard location)
	if runtime.GOOS == "darwin" {
		for _, app := range findMacAppsWithMdfind() {
			roots := getCandidateRoots(app)
			if found := probeRootsForLanguageServer(roots); found != "" {
				logger("Discovered language_server via Spotlight (%s): %s", app, found)
				return found
			}
		}
	}

	// 4. Scan all standard candidates
	candidates := GetStandardLanguageServerCandidates()
	for _, c := range candidates {
		if fi, err := os.Stat(c); err == nil && !fi.IsDir() {
			logger("Found language_server at standard path: %s", c)
			return c
		}
	}

	logger("language_server binary not found in running processes or standard paths")
	return ""
}

// FindRunningLanguageServerPID returns the PID of an active language_server, or 0 if none is active
func FindRunningLanguageServerPID() int {
	if runtime.GOOS == "linux" {
		entries, err := os.ReadDir("/proc")
		if err == nil {
			for _, entry := range entries {
				if !entry.IsDir() {
					continue
				}
				pid, err := strconv.Atoi(entry.Name())
				if err != nil {
					continue
				}
				exeLink := fmt.Sprintf("/proc/%d/exe", pid)
				target, err := os.Readlink(exeLink)
				if err == nil {
					if isSelfOrCleanerProcess(pid, target) {
						continue
					}
					base := strings.ToLower(filepath.Base(target))
					if strings.Contains(base, "language_server") {
						return pid
					}
				}
			}
		}
	}

	// Cross-platform fallback using pgrep / ps
	out, err := exec.Command("pgrep", "-f", "language_server").Output()
	if err == nil {
		selfPid := os.Getpid()
		ppid := os.Getppid()
		for _, pStr := range strings.Fields(string(out)) {
			var pid int
			if _, err := fmt.Sscanf(pStr, "%d", &pid); err == nil && pid > 0 {
				if pid == selfPid || (ppid > 0 && pid == ppid) {
					continue
				}
				return pid
			}
		}
	}
	return 0
}

