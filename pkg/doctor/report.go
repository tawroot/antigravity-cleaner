package doctor

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// BuildMarkdownReport converts DiagnosticItem slice into clean, structured Markdown
func BuildMarkdownReport(version string, results []DiagnosticItem) string {
	var sb strings.Builder

	sb.WriteString("🩺 **Antigravity Cleaner Diagnostic Report**\n")
	sb.WriteString(fmt.Sprintf("📅 *Generated at %s*\n\n", time.Now().Format("2006-01-02 15:04:05 MST")))
	sb.WriteString(fmt.Sprintf("• **Cleaner Toolkit**: v%s (@dalroot)\n", version))

	for _, item := range results {
		icon := "✔"
		if item.Status == "WARN" {
			icon = "⚠️"
		} else if item.Status == "FAIL" {
			icon = "❌"
		}
		sb.WriteString(fmt.Sprintf("• **%s**: %s %s\n", item.Name, icon, item.HumanSummary))
		if item.DebugDetails != "" && item.Status != "OK" {
			details := strings.TrimSpace(item.DebugDetails)
			if len(details) > 85 {
				details = details[:82] + "..."
			}
			sb.WriteString(fmt.Sprintf("  ↳ `%s`\n", details))
		}
	}

	sb.WriteString("\n💡 *Paste this report directly into GitHub Issues or Telegram support (@dallrep)*\n")
	return sb.String()
}

// GenerateDiagnosticReport produces a clean, structured Markdown report
func GenerateDiagnosticReport(version string) string {
	results := RunDiagnostics(true)
	return BuildMarkdownReport(version, results)
}

// CopyToClipboard attempts to place text onto the system clipboard
func CopyToClipboard(text string) bool {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "windows":
		cmd = exec.Command("powershell", "-NoProfile", "-Command", "Set-Clipboard", "-Value", "$input")
	default: // Linux
		if _, err := exec.LookPath("wl-copy"); err == nil {
			cmd = exec.Command("wl-copy")
		} else if _, err := exec.LookPath("xclip"); err == nil {
			cmd = exec.Command("xclip", "-selection", "clipboard")
		} else if _, err := exec.LookPath("xsel"); err == nil {
			cmd = exec.Command("xsel", "--clipboard", "--input")
		}
	}

	if cmd != nil {
		cmd.Stdin = strings.NewReader(text)
		if err := cmd.Run(); err == nil {
			return true
		}
	}

	// Fallback to OSC 52 ANSI escape sequence (supported natively by modern terminals)
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	osc52 := fmt.Sprintf("\x1b]52;c;%s\x07", encoded)
	if tty, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0); err == nil {
		_, _ = tty.WriteString(osc52)
		_ = tty.Close()
		return true
	}
	_, _ = os.Stdout.WriteString(osc52)
	return true
}

// SaveAndOpenReport saves the report to a user-visible text file and opens it in default text viewer
func SaveAndOpenReport(report string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = os.TempDir()
	}

	targetDir := homeDir
	desktopDir := filepath.Join(homeDir, "Desktop")
	if info, err := os.Stat(desktopDir); err == nil && info.IsDir() {
		targetDir = desktopDir
	}

	filePath := filepath.Join(targetDir, "antigravity-report.txt")

	instructions := "================================================================================\n" +
		"  🩺 ANTIGRAVITY CLEANER — DIAGNOSTIC REPORT\n" +
		"================================================================================\n" +
		"  📋 INSTRUCTIONS (راهنمای ارسال گزارش):\n" +
		"  1. Select and copy all text below (Ctrl+A then Ctrl+C / Cmd+A then Cmd+C).\n" +
		"  2. Paste it into the GitHub Discussions Clinic:\n" +
		"     👉 https://github.com/tawroot/antigravity-cleaner/discussions/34\n" +
		"     Or Telegram Community:\n" +
		"     👉 https://t.me/dallrep\n" +
		"================================================================================\n\n"

	fullContent := instructions + report
	if err := os.WriteFile(filePath, []byte(fullContent), 0644); err != nil {
		filePath = filepath.Join(os.TempDir(), "antigravity-report.txt")
		if err2 := os.WriteFile(filePath, []byte(fullContent), 0644); err2 != nil {
			return "", err2
		}
	}

	openFileInDefaultEditor(filePath)
	return filePath, nil
}

func openFileInDefaultEditor(filePath string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", filePath)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", filePath)
	default: // Linux
		if _, err := exec.LookPath("xdg-open"); err == nil {
			cmd = exec.Command("xdg-open", filePath)
		}
	}
	if cmd != nil {
		_ = cmd.Start()
	}
}
