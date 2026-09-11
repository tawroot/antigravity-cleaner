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
	out, err := exec.Command("pgrep", "-f", "language_server").Output()
	if err == nil {
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		for _, l := range lines {
			l = strings.TrimSpace(l);
			if l != "" && l != selfPID {
				return true
			}
		}
	}

	out2, err2 := exec.Command("pgrep", "-f", "antigravity").Output()
	if err2 == nil {
		lines := strings.Split(strings.TrimSpace(string(out2)), "\n")
		for _, l := range lines {
			l = strings.TrimSpace(l);
			if l != "" && l != selfPID {
				return true
			}
		}
	}

	return false
}

// KillAntigravityProcesses terminates active Antigravity and language_server instances
func KillAntigravityProcesses() error {
	if runtime.GOOS == "windows" {
		_ = exec.Command("taskkill", "/F", "/IM", "language_server.exe").Run()
		_ = exec.Command("taskkill", "/F", "/IM", "Antigravity.exe").Run()
		time.Sleep(1 * time.Second)
		return nil
	}

	_ = exec.Command("pkill", "-9", "-f", "language_server").Run()
	_ = exec.Command("pkill", "-9", "-f", "antigravity").Run()
	time.Sleep(1 * time.Second)
	return nil
}
