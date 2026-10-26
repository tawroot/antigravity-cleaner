package doctor

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
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
