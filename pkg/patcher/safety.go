package patcher

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// IsAntigravityRunning checks if language_server or Antigravity IDE is currently active
func IsAntigravityRunning() bool {
	selfPID := fmt.Sprintf("%d", os.Getpid())

	if runtime.GOOS == "windows" {
		out, err := exec.Command("tasklist", "/FI", "IMAGENAME eq language_server.exe").Output()
		if err == nil && strings.Contains(string(out), "language_server") {
			return true
		}
		out2, err2 := exec.Command("tasklist", "/FI", "IMAGENAME eq Antigravity.exe").Output()
		if err2 == nil && strings.Contains(string(out2), "Antigravity") {
			return true
		}
		return false
	}

	// Linux / macOS
	out, err := exec.Command("pgrep", "-a", "-f", "language_server").Output()
	if err == nil {
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		for _, l := range lines {
			parts := strings.Fields(l)
			if len(parts) > 0 && parts[0] != selfPID {
				return true
			}
		}
	}

	out2, err2 := exec.Command("pgrep", "-a", "-f", "antigravity").Output()
	if err2 == nil {
		lines := strings.Split(strings.TrimSpace(string(out2)), "\n")
		for _, l := range lines {
			parts := strings.Fields(l)
			if len(parts) > 1 {
				pid := parts[0]
				cmd := strings.Join(parts[1:], " ")
				if pid != selfPID && !strings.Contains(cmd, "antigravity-cleaner") && !strings.Contains(cmd, "ag-cleaner") {
					return true
				}
			}
		}
	}

	return false
}

// KillAntigravityProcesses terminates active Antigravity and language_server instances safely
func KillAntigravityProcesses() (int, error) {
	killedCount := 0
	selfPID := os.Getpid()

	if runtime.GOOS == "windows" {
		_ = exec.Command("taskkill", "/F", "/IM", "language_server.exe").Run()
		_ = exec.Command("taskkill", "/F", "/IM", "Antigravity.exe").Run()
		time.Sleep(500 * time.Millisecond)
		return 1, nil
	}

	// 1. Terminate language_server processes
	outLS, _ := exec.Command("pgrep", "-f", "language_server").Output()
	for _, pStr := range strings.Fields(string(outLS)) {
		var pid int
		if _, err := fmt.Sscanf(pStr, "%d", &pid); err == nil && pid != selfPID {
			if p, err := os.FindProcess(pid); err == nil {
				_ = p.Kill()
				killedCount++
			}
		}
	}

	// 2. Terminate Antigravity IDE processes, strictly EXCLUDING our own cleaner toolkit and GUI browser
	outAnti, _ := exec.Command("pgrep", "-a", "-f", "antigravity").Output()
	lines := strings.Split(strings.TrimSpace(string(outAnti)), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) > 1 {
			var pid int
			if _, err := fmt.Sscanf(parts[0], "%d", &pid); err == nil && pid != selfPID {
				cmd := strings.Join(parts[1:], " ")
				// Strictly ignore our own cleaner process and the patcher GUI window
				if strings.Contains(cmd, "antigravity-cleaner") ||
					strings.Contains(cmd, "ag-cleaner") ||
					strings.Contains(cmd, "antigravity-patcher") ||
					strings.Contains(cmd, "antigravity-gui") {
					continue
				}
				if p, err := os.FindProcess(pid); err == nil {
					_ = p.Kill()
					killedCount++
				}
			}
		}
	}

	time.Sleep(500 * time.Millisecond)
	return killedCount, nil
}
