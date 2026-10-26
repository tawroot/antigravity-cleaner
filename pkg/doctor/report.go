package doctor

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// GenerateDiagnosticReport produces a clean, structured Markdown report
func GenerateDiagnosticReport(version string) string {
	results := RunDiagnostics(true)
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
			sb.WriteString(fmt.Sprintf("  ↳ `%s`\n", item.DebugDetails))
		}
	}

	sb.WriteString("\n💡 *Paste this report directly into GitHub Issues or Telegram support (@dallrep)*\n")
	return sb.String()
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

	if cmd == nil {
		return false
	}

	cmd.Stdin = strings.NewReader(text)
	return cmd.Run() == nil
}
