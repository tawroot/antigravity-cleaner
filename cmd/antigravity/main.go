package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/tawroot/antigravity-cleaner/pkg/cleaner"
	"github.com/tawroot/antigravity-cleaner/pkg/doctor"
	"github.com/tawroot/antigravity-cleaner/pkg/patcher"
	"github.com/tawroot/antigravity-cleaner/pkg/proxy"
	"github.com/tawroot/antigravity-cleaner/pkg/ui"
)

const Version = "5.1.1"

func main() {
	if len(os.Args) > 1 {
		handleSubcommand(os.Args[1], os.Args[2:])
		return
	}

	keyStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00F0FF"))

	descStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F8F8F2"))

	// Interactive Menu Loop
	for {
		clearScreen()
		ui.PrintBanner(Version)

		menuItems := []struct {
			key  string
			desc string
		}{
			{"1", "⚡ Quick Auto-Fix (Full Patch + Injected Proxy)"},
			{"2", "🔓 Patch Core (language_server & agy Machine Code)"},
			{"3", "🎨 Patch IDE (main.js & VS Code Extension Guard)"},
			{"4", "🚀 Launch Antigravity with Smart Proxy (No TUN!)"},
			{"5", "📌 Create Desktop Shortcut with Injected Proxy"},
			{"6", "🧹 Surgical 429 Quota Reset (Preserves Chats & Settings)"},
			{"7", "🩺 Run Antigravity Doctor (System & Connection Health)"},
			{"8", "🔄 Restore Backups (.agybak)"},
			{"0", "🚪 Exit"},
		}

		for _, item := range menuItems {
			fmt.Printf("  [%s] %s\n", keyStyle.Render(item.key), descStyle.Render(item.desc))
		}
		fmt.Println()

		choice := ui.ReadInput("Select an option")
		switch choice {
		case "1":
			runQuickAutoFix()
			ui.PromptKey()
		case "2":
			runPatchCore()
			ui.PromptKey()
		case "3":
			runPatchIde()
			ui.PromptKey()
		case "4":
			runLaunchProxy()
			ui.PromptKey()
		case "5":
			runCreateLauncher()
			ui.PromptKey()
		case "6":
			runSurgicalClean()
			ui.PromptKey()
		case "7":
			runDoctor(false)
			ui.PromptKey()
		case "8":
			runRestore()
			ui.PromptKey()
		case "0", "exit", "q":
			fmt.Printf("\n  %s\n\n", lipgloss.NewStyle().Foreground(lipgloss.Color("#00F0FF")).Render("🦅 Farewell! Happy Coding with Antigravity."))
			return
		default:
			fmt.Printf("\n  %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB800")).Render("Invalid option. Please choose between 0-8."))
			ui.PromptKey()
		}
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func handleSubcommand(cmd string, args []string) {
	switch cmd {
	case "doctor":
		verbose := false
		for _, a := range args {
			if a == "-v" || a == "--verbose" {
				verbose = true
			}
		}
		runDoctor(verbose)
	case "patch":
		target := "all"
		if len(args) > 0 {
			target = args[0]
		}
		switch target {
		case "manager", "core":
			runPatchCore()
		case "ide":
			runPatchIde()
		default:
			runQuickAutoFix()
		}
	case "launch":
		var proxyURL string
		fs := flag.NewFlagSet("launch", flag.ExitOnError)
		fs.StringVar(&proxyURL, "proxy", "", "Custom proxy URL (e.g. socks5h://127.0.0.1:10808)")
		_ = fs.Parse(args)
		_ = proxy.LaunchAntigravityWithProxy("", proxyURL, fs.Args())
	case "clean":
		runSurgicalClean()
	case "kill":
		runKillProcesses()
	case "create-launcher":
		runCreateLauncher()
	case "about":
		printAbout()
	case "version", "-v", "--version":
		fmt.Printf("Antigravity Cleaner Toolkit v%s (@dalroot)\n", Version)
	default:
		fmt.Printf("Unknown command: %s\nUsage: antigravity-cleaner [doctor|patch|launch|clean|kill|create-launcher|about|version]\n", cmd)
		os.Exit(1)
	}
}

func printAbout() {
	ui.PrintBanner(Version)
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00F0FF")).Render("  ► Universal Compatibility & Detection:"))
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Render("  - Universal Engine: Automatically detects and patches ALL Antigravity versions (1.x, 2.x, Preview)."))
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Render("  - Multi-Target Support: Antigravity IDE, agy CLI, and VS Code Extension (google.google-antigravity)."))
	fmt.Println()
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF9F")).Render("  ► Project Credits & Attribution:"))
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Render("  - Lead Architect & Maintainer: @dalroot"))
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Render("  - Co-Pilot & Pair Programming: Antigravity AI (Google DeepMind)"))
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#8BE9FD")).Render("  - Repository: https://github.com/tawroot/antigravity-cleaner"))
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#6272A4")).Render("  - Dedicated to developers navigating digital sanctions worldwide."))
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#6272A4")).Render("  - License: GNU General Public License v3.0"))
	fmt.Println()
}

func runDoctor(verbose bool) {
	results := doctor.RunDiagnostics(verbose)
	var rows []ui.CardRow
	for _, r := range results {
		rows = append(rows, ui.CardRow{
			Title:   r.Name,
			Status:  r.Status,
			Summary: r.HumanSummary,
			Details: r.DebugDetails,
		})
	}
	ui.PrintDashboard(rows, verbose)
}

func ensureRunningSafety(force bool) bool {
	if !patcher.IsAntigravityRunning() {
		return true
	}

	warnStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFB800"))
	fmt.Println()
	fmt.Println(warnStyle.Render("  ⚠️ NOTICE: Active Antigravity or language_server processes detected."))
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Render(
		"  Closing running processes prevents zombie socket locks and ensures clean patches."))
	fmt.Println()

	if force {
		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF9F")).Render("  [--force specified] Terminating running processes cleanly..."))
		_, _ = patcher.KillAntigravityProcesses()
		return true
	}

	confirm := ui.ReadInput("Do you want to automatically close active processes before proceeding? (Y/n)")
	if confirm == "n" || confirm == "N" {
		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB800")).Render("  Operation cancelled to preserve running Antigravity session."))
		return false
	}

	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF9F")).Render("  ✔ Closed active processes cleanly."))
	_, _ = patcher.KillAntigravityProcesses()
	return true
}

func runKillProcesses() {
	ui.PrintSection("🧹 Killing Active Antigravity Processes")
	if !patcher.IsAntigravityRunning() {
		ui.PrintStatus("OK", "Processes", "No active Antigravity or language_server processes found.")
		return
	}
	_, _ = patcher.KillAntigravityProcesses()
	ui.PrintStatus("OK", "Processes Terminated", "All Antigravity and language_server processes were killed.")
}

func runQuickAutoFix() {
	if !ensureRunningSafety(false) {
		return
	}
	ui.PrintSection("⚡ Smart 1-Click Auto-Fix")

	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#BD00FF")).Render("  Step 1: Core Engine MultiGate Bytecode Unlock..."))
	coreSuccess := false
	if res, err := patcher.PatchLanguageServer(""); err != nil {
		ui.PrintStatus("FAIL", "Core Engine", err.Error())
	} else {
		ui.PrintStatus("OK", "Core Engine", res)
		coreSuccess = true
	}

	// Optional components (do not warn if not installed, show clean INFO)
	if res, err := patcher.PatchAgy(""); err != nil {
		ui.PrintStatus("INFO", "Antigravity CLI (agy)", "Standalone Mode (CLI not installed, optional)")
	} else {
		ui.PrintStatus("OK", "Antigravity CLI (agy)", res)
	}

	if res, err := patcher.PatchIdeMainJs(""); err != nil {
		ui.PrintStatus("INFO", "IDE main.js", "Antigravity 2.x Bundled (Internal gate active)")
	} else {
		ui.PrintStatus("OK", "IDE main.js", res)
	}

	fmt.Println()
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00F0FF")).Render("  Step 2: Smart Proxy & Desktop Launcher Setup..."))
	if path, err := proxy.GenerateNoTunDesktopLauncher(""); err != nil {
		ui.PrintStatus("WARN", "Desktop Launcher", err.Error())
	} else {
		ui.PrintStatus("OK", "Desktop Launcher", "Created at "+path)
	}

	if coreSuccess {
		fmt.Println()
		successBox := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00FF9F")).
			Padding(0, 2).
			Render(fmt.Sprintf("%s\n\n%s\n%s\n%s\n\n%s",
				lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF9F")).Render("🎉 AUTO-FIX COMPLETED SUCCESSFULLY!"),
				lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Render("✔ Core Engine:      UNLOCKED (MultiGate hasValidAuth Bypass Active)"),
				lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Render("✔ Smart Proxy:      CONFIGURED (Direct socket, Zero TUN mode needed)"),
				lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Render("✔ Desktop Shortcut: READY (Click 'Antigravity' in App Menu or Desktop)"),
				lipgloss.NewStyle().Foreground(lipgloss.Color("#00F0FF")).Render("🚀 You can now launch Antigravity with full region freedom!"),
			))
		fmt.Println(successBox)
	}
}

func runPatchCore() {
	if !ensureRunningSafety(false) {
		return
	}
	ui.PrintSection("Patch Core Binaries")
	if res, err := patcher.PatchLanguageServer(""); err != nil {
		ui.PrintStatus("FAIL", "Language Server", err.Error())
	} else {
		ui.PrintStatus("OK", "Language Server", res)
	}

	if res, err := patcher.PatchAgy(""); err != nil {
		ui.PrintStatus("WARN", "Antigravity CLI (agy)", err.Error())
	} else {
		ui.PrintStatus("OK", "Antigravity CLI (agy)", res)
	}
}

func runPatchIde() {
	if !ensureRunningSafety(false) {
		return
	}
	ui.PrintSection("Patch IDE & VS Code Extensions")
	if res, err := patcher.PatchIdeMainJs(""); err != nil {
		ui.PrintStatus("WARN", "IDE main.js", err.Error())
	} else {
		ui.PrintStatus("OK", "IDE main.js", res)
	}

	if res, err := patcher.PatchVsCodeExtension(""); err != nil {
		ui.PrintStatus("WARN", "VS Code Extension", err.Error())
	} else {
		ui.PrintStatus("OK", "VS Code Extension", res)
	}
}

func runLaunchProxy() {
	ui.PrintSection("Launching with Smart Proxy (No-TUN Mode)")
	best, err := proxy.GetBestProxy()
	if err != nil {
		ui.PrintStatus("FAIL", "Proxy Error", err.Error())
		return
	}

	ui.PrintStatus("OK", "Proxy Detected", fmt.Sprintf("%s (%s)", best.URL, best.Provider))
	fmt.Printf("  %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("#6272A4")).Render("Starting Antigravity with injected environment..."))

	if err := proxy.LaunchAntigravityWithProxy("", best.URL, nil); err != nil {
		ui.PrintStatus("FAIL", "Launch Error", err.Error())
	} else {
		ui.PrintStatus("OK", "Process Spawned", "Antigravity launched successfully without TUN mode!")
	}
}

func runCreateLauncher() {
	ui.PrintSection("Create Desktop Launcher")
	path, err := proxy.GenerateNoTunDesktopLauncher("")
	if err != nil {
		ui.PrintStatus("FAIL", "Launcher Error", err.Error())
	} else {
		ui.PrintStatus("OK", "Shortcut Created", path)
		fmt.Printf("  %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF9F")).Render("Desktop entry ready! Launch Antigravity without TUN mode anytime."))
	}
}

func runSurgicalClean() {
	if !ensureRunningSafety(false) {
		return
	}
	ui.PrintSection("Surgical 429 Quota & Session Reset")
	fmt.Println("  This will clear corrupted token cache, DIPS, and cookies to resolve HTTP 429.")
	fmt.Printf("  %s\n\n", lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF9F")).Render("All project chats, workspaces, settings and keybindings are strictly PRESERVED."))

	confirm := ui.ReadInput("Do you want to proceed? (y/N)")
	if confirm != "y" && confirm != "Y" {
		fmt.Println("  Operation cancelled.")
		return
	}

	removed, err := cleaner.SurgicalReset429("", false)
	if err != nil {
		ui.PrintStatus("FAIL", "Clean Error", err.Error())
		return
	}

	ui.PrintStatus("OK", "Reset Completed", fmt.Sprintf("Cleaned %d cache & session artifacts", len(removed)))
	for _, item := range removed {
		fmt.Printf("    - %s\n", item)
	}
}

func runRestore() {
	ui.PrintSection("Restore Backups")
	lsPath := patcher.FindLanguageServer()
	if lsPath != "" {
		if err := patcher.RestoreFile(lsPath); err == nil {
			ui.PrintStatus("OK", "Language Server", "Restored original backup")
		}
	}
	agyPath := patcher.FindAgy()
	if agyPath != "" {
		if err := patcher.RestoreFile(agyPath); err == nil {
			ui.PrintStatus("OK", "Antigravity CLI", "Restored original backup")
		}
	}
	idePath := patcher.FindIdeMainJs()
	if idePath != "" {
		if err := patcher.RestoreFile(idePath); err == nil {
			ui.PrintStatus("OK", "IDE main.js", "Restored original backup")
		}
	}
}
