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

// scanLinuxProc inspects /proc for running processes matching target substring
func scanLinuxProc(targetSubstr string) string {
	entries, err := os.ReadDir("/proc")
	if err != nil {
		return ""
	}

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
		if err == nil && strings.Contains(strings.ToLower(target), strings.ToLower(targetSubstr)) {
			// Verify file exists
			if fi, err := os.Stat(target); err == nil && !fi.IsDir() {
				return target
			}
		}

		// Also check cmdline in case exe link points to an interpreter or wrapper
		cmdlineFile := fmt.Sprintf("/proc/%d/cmdline", pid)
		if data, err := os.ReadFile(cmdlineFile); err == nil {
			args := strings.Split(string(data), "\x00")
			if len(args) > 0 {
				arg0 := args[0]
				if strings.Contains(strings.ToLower(arg0), strings.ToLower(targetSubstr)) {
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

	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(targetSubstr)) {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				// Search fields for an existing executable path
				for _, f := range fields[1:] {
					if strings.Contains(strings.ToLower(f), strings.ToLower(targetSubstr)) {
						clean := strings.Trim(f, "\"'")
						if fi, err := os.Stat(clean); err == nil && !fi.IsDir() {
							return clean
						}
					}
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

// GetStandardLanguageServerCandidates returns all standard paths across Linux, Windows, macOS
func GetStandardLanguageServerCandidates() []string {
	home, _ := os.UserHomeDir()

	return []string{
		// Linux & User Applications directory (Discovered via REA on modern Linux)
		filepath.Join(home, "Applications", "Antigravity", "resources", "bin", "language_server"),
		filepath.Join(home, "Applications", "antigravity", "resources", "bin", "language_server"),
		filepath.Join(home, "Applications", "Antigravity-x64", "resources", "bin", "language_server"),
		"/opt/antigravity/resources/bin/language_server",
		"/opt/Antigravity/resources/bin/language_server",
		"/opt/antigravity-ide/resources/bin/language_server",
		"/opt/Antigravity IDE/resources/bin/language_server",
		"/usr/share/antigravity/resources/bin/language_server",
		"/usr/share/antigravity-ide/resources/bin/language_server",
		filepath.Join(home, ".local", "share", "antigravity", "resources", "bin", "language_server"),

		// Legacy & home relative paths
		filepath.Join(home, "antigravity", "resources", "bin", "language_server"),
		filepath.Join(home, "antigravity", "resources", "bin", "language_server.exe"),

		// Windows paths
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Antigravity", "resources", "bin", "language_server.exe"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "antigravity", "resources", "bin", "language_server.exe"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Antigravity IDE", "resources", "bin", "language_server.exe"),
		filepath.Join(os.Getenv("PROGRAMFILES"), "Antigravity", "resources", "bin", "language_server.exe"),
		filepath.Join(os.Getenv("ProgramFiles(x86)"), "Antigravity", "resources", "bin", "language_server.exe"),

		// macOS paths
		"/Applications/Antigravity.app/Contents/Resources/bin/language_server",
		"/Applications/Antigravity IDE.app/Contents/Resources/bin/language_server",
		filepath.Join(home, "Applications", "Antigravity.app", "Contents", "Resources", "bin", "language_server"),
		filepath.Join(home, "Applications", "Antigravity IDE.app", "Contents", "Resources", "bin", "language_server"),
	}
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
		binName := "language_server"
		if runtime.GOOS == "windows" {
			binName = "language_server.exe"
		}
		deducedPath := filepath.Join(filepath.Dir(runningApp), "resources", "bin", binName)
		if fi, err := os.Stat(deducedPath); err == nil && !fi.IsDir() {
			logger("Deduced language_server path from running app: %s", deducedPath)
			return deducedPath
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
				if err == nil && strings.Contains(strings.ToLower(target), "language_server") {
					return pid
				}
			}
		}
	}

	// Cross-platform fallback using pgrep / ps
	out, err := exec.Command("pgrep", "-f", "language_server").Output()
	if err == nil {
		for _, pStr := range strings.Fields(string(out)) {
			var pid int
			if _, err := fmt.Sscanf(pStr, "%d", &pid); err == nil && pid > 0 {
				return pid
			}
		}
	}
	return 0
}

