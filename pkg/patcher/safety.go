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

// KillAntigravityProcesses terminates active Antigravity and language_server instances
func KillAntigravityProcesses() (int, error) {
	killedCount := 0

	if runtime.GOOS == "windows" {
		_ = exec.Command("taskkill", "/F", "/IM", "language_server.exe").Run()
		_ = exec.Command("taskkill", "/F", "/IM", "Antigravity.exe").Run()
		_ = exec.Command("taskkill", "/F", "/IM", "antigravity.exe").Run()
		time.Sleep(1 * time.Second)
		return 1, nil
	}

	// Linux / macOS: kill language_server and antigravity
	_ = exec.Command("pkill", "-9", "-f", "language_server").Run()
	_ = exec.Command("pkill", "-9", "-f", "antigravity").Run()
	_ = exec.Command("pkill", "-9", "-f", "chrome-sandbox").Run()
	time.Sleep(1 * time.Second)

	return killedCount, nil
}
