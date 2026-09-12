package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Brand and Version metadata
const (
	BrandName    = "Antigravity Cleaner"
	BrandTagline = "The High-Performance AI Freedom Toolkit"
	Author       = "@dalroot & Antigravity AI"
)

// 2026 Modern Aesthetic Color Palette
var (
	ColorPrimary   = lipgloss.Color("#00F0FF") // Electric Cyan
	ColorSecondary = lipgloss.Color("#BD00FF") // Neon Purple
	ColorSuccess   = lipgloss.Color("#00FF9F") // Mint Emerald
	ColorWarning   = lipgloss.Color("#FFB800") // Warm Amber
	ColorDanger    = lipgloss.Color("#FF3366") // Coral Red
	ColorMuted     = lipgloss.Color("#6272A4") // Slate Gray

	StyleTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary)

	StyleSubtitle = lipgloss.NewStyle().
			Foreground(ColorSecondary)

	StyleAuthor = lipgloss.NewStyle().
			Foreground(ColorMuted).
			Italic(true)

	StyleCard = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorSecondary).
			Padding(0, 1).
			MarginBottom(1)

	StyleItemKey = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#F8F8F2")).
			Width(22)

	BadgeOk = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(ColorSuccess).
		Padding(0, 1)

	BadgeWarn = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#000000")).
			Background(ColorWarning).
			Padding(0, 1)

	BadgeFail = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(ColorDanger).
			Padding(0, 1)
)

// PrintBanner prints the Antigravity Cleaner Banner
func PrintBanner(version string) {
	bannerText := `
   ___          __  _                         _ __         
  /   |  ____  / /_(_)___ __________ __   __(_) /___  __  
 / /| | / __ \/ __/ / __ '/ ___/ __ '/ | / / / __/ / / /  
/ ___ |/ / / / /_/ / /_/ / /  / /_/ /| |/ / / /_/ /_/ /   
/_/  |_/_/ /_/\__/_/\__, /_/   \__,_/ |___/_/\__/\__, /    
                  /____/                       /____/     `

	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(ColorPrimary).Render(bannerText))
	fmt.Printf("  %s %s%s v%s%s\n",
		"⚡",
		lipgloss.NewStyle().Bold(true).Foreground(ColorSecondary).Render("Antigravity Cleaner"),
		lipgloss.NewStyle().Foreground(ColorMuted).Render(" — Universal AI Freedom Toolkit"),
		version,
		lipgloss.NewStyle().Foreground(ColorMuted).Render(" | @dalroot"))
	fmt.Printf("  %s\n", lipgloss.NewStyle().Bold(true).Foreground(ColorSuccess).Render("✔ Universal Auto-Detection: Supports ALL Antigravity Versions (1.x, 2.x, IDE & CLI)"))
	fmt.Printf("  %s\n", lipgloss.NewStyle().Foreground(ColorPrimary).Render("🔒 100% Offline & Zero-Track: Zero telemetry, zero analytics, completely local & safe"))
	fmt.Printf("  %s\n\n", StyleAuthor.Render("Crafted in creative pair-programming with Antigravity AI (Google DeepMind)"))
}

// CardRow represents a single diagnostic check row
type CardRow struct {
	Title   string
	Status  string // "OK", "WARN", "FAIL"
	Summary string
	Details string
}

// PrintDashboard displays an ultra-clean, modern 2026 Lipgloss Card
func PrintDashboard(rows []CardRow, verbose bool) {
	fmt.Println()

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorPrimary).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(ColorMuted).
		Width(72).
		Align(lipgloss.Left)

	header := headerStyle.Render(" 🩺 Antigravity Cleaner — Health & Connectivity Audit")
	var content strings.Builder
	content.WriteString(header + "\n\n")

	for _, r := range rows {
		var badge string
		switch r.Status {
		case "OK":
			badge = BadgeOk.Render("READY")
		case "WARN":
			badge = BadgeWarn.Render("CHECK")
		case "FAIL":
			badge = BadgeFail.Render("OFFLINE")
		default:
			badge = BadgeWarn.Render("INFO")
		}

		key := StyleItemKey.Render(r.Title)
		val := lipgloss.NewStyle().Foreground(lipgloss.Color("#E2E8F0")).Render(r.Summary)

		content.WriteString(fmt.Sprintf(" %s %s  %s\n", key, badge, val))

		if verbose && r.Details != "" {
			sub := lipgloss.NewStyle().Foreground(ColorMuted).Render(fmt.Sprintf("   ↳ %s", r.Details))
			content.WriteString(sub + "\n")
		}
	}

	cardBox := StyleCard.Width(76).Render(content.String())
	fmt.Println(cardBox)

	// Summary Footer
	okCount := 0
	for _, r := range rows {
		if r.Status == "OK" {
			okCount++
		}
	}

	footerStyle := lipgloss.NewStyle().PaddingLeft(2)
	if okCount >= len(rows)-1 {
		fmt.Println(footerStyle.Render(
			BadgeOk.Render("ONLINE") +
				lipgloss.NewStyle().Foreground(ColorSuccess).Render(" System is optimized. You can launch Antigravity with Smart Proxy!\n"),
		))
	} else {
		fmt.Println(footerStyle.Render(
			BadgeWarn.Render("ACTION") +
				lipgloss.NewStyle().Foreground(ColorWarning).Render(" Run with --verbose for detailed trace or run Option 1 to Auto-Fix.\n"),
		))
	}
}

// PrintStatus prints formatted status with indicator
func PrintStatus(level, title, message string) {
	var badge string
	switch level {
	case "OK", "SUCCESS":
		badge = BadgeOk.Render(" OK ")
	case "WARN", "WARNING":
		badge = BadgeWarn.Render("WARN")
	case "FAIL", "ERROR":
		badge = BadgeFail.Render("FAIL")
	case "SKIP":
		badge = lipgloss.NewStyle().Bold(true).Background(ColorMuted).Foreground(lipgloss.Color("#FFFFFF")).Padding(0, 1).Render("SKIP")
	default:
		badge = lipgloss.NewStyle().Bold(true).Background(ColorPrimary).Foreground(lipgloss.Color("#000000")).Padding(0, 1).Render("INFO")
	}

	fmt.Printf("  %s %s%-24s%s %s\n",
		badge,
		lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#F8F8F2")).Render(""),
		title,
		lipgloss.NewStyle().Foreground(ColorMuted).Render("→"),
		message)
}

// PrintSection prints a styled section header
func PrintSection(title string) {
	divider := lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorPrimary).
		Render(fmt.Sprintf("\n─── %s ───────────────────────────────────────────────────", title))
	fmt.Println(divider)
}

// PromptKey prompts the user to press Enter to continue
func PromptKey() {
	fmt.Printf("\n  %s", lipgloss.NewStyle().Foreground(ColorMuted).Render("Press Enter to return to menu..."))
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}

// ReadInput reads a single line of input from stdin
func ReadInput(prompt string) string {
	fmt.Printf("  %s: ", lipgloss.NewStyle().Bold(true).Foreground(ColorWarning).Render(prompt))
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}
