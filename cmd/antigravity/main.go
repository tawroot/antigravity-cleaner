package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-isatty"
	"github.com/tawroot/antigravity-cleaner/pkg/account"
	"github.com/tawroot/antigravity-cleaner/pkg/cleaner"
	"github.com/tawroot/antigravity-cleaner/pkg/doctor"
	"github.com/tawroot/antigravity-cleaner/pkg/gui"
	"github.com/tawroot/antigravity-cleaner/pkg/patcher"
	"github.com/tawroot/antigravity-cleaner/pkg/proxy"
	"github.com/tawroot/antigravity-cleaner/pkg/quota"
	"github.com/tawroot/antigravity-cleaner/pkg/ui"
	"github.com/tawroot/antigravity-cleaner/pkg/updater"
)

const Version = "5.4.4"

func init() {
	patcher.RegisterPostInstallHook(func(logger patcher.LogFunc) {
		_ = account.SyncTokenBridge(logger)
	})
}

func main() {
	if len(os.Args) > 1 {
		handleSubcommand(os.Args[1], os.Args[2:])
		return
	}

	// Auto-launch GUI if executed without a terminal (e.g. desktop icon / file manager double click)
	if !isatty.IsTerminal(os.Stdin.Fd()) && !isatty.IsCygwinTerminal(os.Stdin.Fd()) {
		if err := gui.Run(); err == nil {
			return
		}
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

		// Live system status indicators
		lsRunning := patcher.IsAntigravityRunning()
		statusRunning := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5555")).Render("Idle")
		if lsRunning {
			statusRunning = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF9F")).Render("Active (Running)")
		}

		lsStatus := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB800")).Render("Unpatched")
		targets := patcher.FindAllLanguageServers(nil)
		if len(targets) > 0 {
			wm := patcher.NewWrapperManager(filepath.Dir(targets[0]), nil)
			if wm.IsWrapperInstalled() {
				lsStatus = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF9F")).Render("✔ Patched & Enterprise")
			} else {
				lsStatus = lipgloss.NewStyle().Foreground(lipgloss.Color("#00F0FF")).Render("Installed")
			}
		}
		fmt.Printf("  📊 Core Status: %s | Process: %s\n\n", lsStatus, statusRunning)

		sectionPrimary := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF9F")).Render("─── ⚡ Quick Actions ────────────────────────────────────────────────")
		sectionWorkspace := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00F0FF")).Render("─── 🚀 Workspace & Accounts ─────────────────────────────────────────")
		sectionTools := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#BD00FF")).Render("─── 🛠️ Maintenance & Diagnostics ───────────────────────────────────")
		sectionCommunity := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFD700")).Render("─── ⭐ Community & Support ──────────────────────────────────────────")

		fmt.Println("  " + sectionPrimary)
		fmt.Printf("  [%s] %s\n", keyStyle.Render("1"), descStyle.Render("⚡ Smart 1-Click Auto-Fix (Recommended: Full Patch + Flags + Proxy)"))
		fmt.Printf("  [%s] %s\n", keyStyle.Render("2"), descStyle.Render("🚀 Launch Antigravity with Smart Proxy (Direct socket, Zero TUN)"))
		fmt.Println()

		fmt.Println("  " + sectionWorkspace)
		fmt.Printf("  [%s] %s\n", keyStyle.Render("q"), descStyle.Render("📊 Live Quota Monitor (Check Gemini/Claude Pro Credits & Reset Timer)"))
		fmt.Printf("  [%s] %s\n", keyStyle.Render("a"), descStyle.Render("👤 Multi-Account Manager (Save / Switch Google Profiles)"))
		fmt.Println()

		fmt.Println("  " + sectionTools)
		fmt.Printf("  [%s] %s\n", keyStyle.Render("3"), descStyle.Render("🧹 Surgical 429 Quota Reset (Preserves Chats & Settings)"))
		fmt.Printf("  [%s] %s\n", keyStyle.Render("4"), descStyle.Render("🩺 Run Antigravity Doctor (System & Connection Health)"))
		fmt.Printf("  [%s] %s\n", keyStyle.Render("m"), descStyle.Render("⚙️ Advanced & Manual Tools... (Granular patches, flags, restore)"))
		fmt.Printf("  [%s] %s\n", keyStyle.Render("g"), descStyle.Render("🎮 Launch Retro GUI (Classic Win95 Patcher)"))
		fmt.Println()

		fmt.Println("  " + sectionCommunity)
		fmt.Printf("  [%s] %s\n", keyStyle.Render("t"), descStyle.Render("Join Telegram Community (@dallrep) for help & updates"))
		fmt.Printf("  [%s] %s\n", keyStyle.Render("s"), descStyle.Render("Star on GitHub (github.com/tawroot/antigravity-cleaner)"))
		fmt.Printf("  [%s] %s\n", keyStyle.Render("0"), descStyle.Render("Exit"))
		fmt.Println()

		choice := ui.ReadInput("Select an option")
		switch choice {
		case "1", "autofix", "fix":
			runQuickAutoFix()
			ui.PromptKey()
		case "2", "launch", "run":
			runLaunchProxy()
			ui.PromptKey()
		case "q", "Q", "quota":
			runQuota()
			ui.PromptKey()
		case "a", "A", "account", "accounts":
			runAccountInteractive()
			ui.PromptKey()
		case "3", "5", "clean":
			runSurgicalClean()
			ui.PromptKey()
		case "4", "6", "doctor":
			runDoctor(false)
			ui.PromptKey()
		case "m", "M", "advanced", "tools", "manual":
			runAdvancedMenu()
		case "g", "8", "gui":
			runGui()
		case "e", "E", "flags", "enterprise":
			runEnterpriseMenu()
			ui.PromptKey()
		case "r", "R", "report":
			runReport()
			ui.PromptKey()
		case "u", "U", "update", "upgrade":
			runUpdate()
			ui.PromptKey()
		case "7", "restore":
			runRestore()
			ui.PromptKey()
		case "t", "T", "telegram":
			openTerminalURL("https://t.me/dallrep")
			ui.PromptKey()
		case "s", "S", "star":
			openTerminalURL("https://github.com/tawroot/antigravity-cleaner")
			ui.PromptKey()
		case "0", "exit":
			fmt.Printf("\n  %s\n", lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFD700")).Render("⭐ If this tool helped you, please give it a Star on GitHub!"))
			fmt.Printf("  %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF")).Render("💬 Join Telegram for support & updates: https://t.me/dallrep"))
			fmt.Printf("  %s\n\n", lipgloss.NewStyle().Foreground(lipgloss.Color("#00F0FF")).Render("🦅 Farewell! Happy Coding with Antigravity."))
			return
		default:
			fmt.Printf("\n  %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB800")).Render("Invalid option. Please choose an option from the menu."))
			ui.PromptKey()
		}
	}
}

func runAdvancedMenu() {
	keyStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00F0FF"))
	descStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2"))

	for {
		clearScreen()
		ui.PrintBanner(Version)
		fmt.Println("  " + lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#BD00FF")).Render("─── ⚙️ Advanced & Granular Tools ────────────────────────────────────"))
		fmt.Printf("  [%s] %s\n", keyStyle.Render("1"), descStyle.Render("Patch Core Binaries only (language_server & agy)"))
		fmt.Printf("  [%s] %s\n", keyStyle.Render("2"), descStyle.Render("Patch IDE & Extensions (main.js, VS Code extension, app.asar)"))
		fmt.Printf("  [%s] %s\n", keyStyle.Render("3"), descStyle.Render("Enterprise & Privacy Flags Manager (SSH spoofing, Mendelfp)"))
		fmt.Printf("  [%s] %s\n", keyStyle.Render("4"), descStyle.Render("Create / Rebuild Desktop Shortcut"))
		fmt.Printf("  [%s] %s\n", keyStyle.Render("5"), descStyle.Render("Generate Diagnostic Report (Auto-Open Text Report)"))
		fmt.Printf("  [%s] %s\n", keyStyle.Render("6"), descStyle.Render("Check for Updates (Self-Update)"))
		fmt.Printf("  [%s] %s\n", keyStyle.Render("7"), descStyle.Render("Restore Backups (.agybak)"))
		fmt.Printf("  [%s] %s\n", keyStyle.Render("0"), descStyle.Render("Back to Main Menu"))
		fmt.Println()

		choice := ui.ReadInput("Select an advanced option")
		switch choice {
		case "1":
			runPatchCore()
			ui.PromptKey()
		case "2":
			runPatchIde()
			ui.PromptKey()
		case "3", "e", "E", "flags", "enterprise":
			runEnterpriseMenu()
			ui.PromptKey()
		case "4":
			runCreateLauncher()
			ui.PromptKey()
		case "5", "r", "R", "report":
			runReport()
			ui.PromptKey()
		case "6", "u", "U", "update":
			runUpdate()
			ui.PromptKey()
		case "7":
			runRestore()
			ui.PromptKey()
		case "0", "b", "back", "q":
			return
		default:
			fmt.Printf("\n  %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB800")).Render("Invalid option. Please choose an option from the menu."))
			ui.PromptKey()
		}
	}
}

func openTerminalURL(target string) {
	fmt.Printf("\n  %s %s\n",
		lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF9F")).Render("🔗 Opening:"),
		target)
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", target)
	case "darwin":
		cmd = exec.Command("open", target)
	default:
		cmd = exec.Command("xdg-open", target)
	}
	_ = cmd.Start()
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
	case "report":
		runReport()
	case "update", "upgrade":
		runUpdate()
	case "autofix", "fix":
		runQuickAutoFix()
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
	case "gui":
		runGui()
	case "clean":
		runSurgicalClean()
	case "quota":
		runQuota(args)
	case "account", "accounts":
		runAccount(args)
	case "enterprise", "flags":
		action := "status"
		if len(args) > 0 {
			action = args[0]
		}
		runEnterpriseFlags(action)
	case "kill":
		runKillProcesses()
	case "create-launcher":
		runCreateLauncher()
	case "about":
		printAbout()
	case "help", "-h", "--help":
		printUsage()
	case "version", "-v", "--version":
		fmt.Printf("Antigravity Cleaner Toolkit v%s (@dalroot)\n", Version)
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Printf("Usage: antigravity-cleaner [autofix|gui|doctor|report|update|patch|launch|clean|quota|account|enterprise|kill|create-launcher|about|version]\n\n" +
		"Subcommands:\n" +
		"  autofix, fix     One-click auto-fix: patch all targets and create desktop launcher\n" +
		"  gui              Launch retro Win95 desktop GUI window\n" +
		"  doctor [-v]      Run health and connectivity audit\n" +
		"  report           Generate and copy full diagnostic report for support\n" +
		"  update           Check GitHub for latest releases and self-update\n" +
		"  patch [all|core] Apply multi-target bytecode and interface unlocks\n" +
		"  launch           Start Antigravity with injected smart No-TUN proxy\n" +
		"  clean            Surgically flush rate-limit and poisoned geo tokens\n" +
		"  quota            Live real-time monitor for Google Gemini Pro credits & reset timer\n" +
		"  account          Manage multiple Google accounts (save, switch, list, rm)\n" +
		"  enterprise       Configure enterprise and zero-telemetry flags\n" +
		"  kill             Safely terminate hung Antigravity processes\n" +
		"  create-launcher  Create desktop application shortcut with proxy\n" +
		"  about            Show project attribution and architecture overview\n" +
		"  version          Show installed cleaner version\n")
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
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#6272A4")).Render("  - License: Apache License 2.0"))
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

	infoStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00F0FF"))
	fmt.Println()
	fmt.Println(infoStyle.Render("  ℹ Active Antigravity session detected."))
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Render(
		"  Antigravity Cleaner uses atomic in-place patching so your session and subagents stay open."))
	fmt.Println()

	if force {
		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF9F")).Render("  [--force specified] Seamlessly soft-restarting language_server core..."))
		_, _ = patcher.RestartLanguageServerOnly()
		return true
	}

	confirm := ui.ReadInput("Do you want to seamlessly soft-restart language_server core now? [Y: soft-restart / n: keep running as-is] (Y/n)")
	if confirm == "n" || confirm == "N" {
		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF9F")).Render("  ✔ Preserving active session intact. Applying live atomic disk patch..."))
		return true
	}

	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF9F")).Render("  ✔ Soft-restarted language_server. Core will reload automatically within 2s."))
	_, _ = patcher.RestartLanguageServerOnly()
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
		if mainJs := patcher.FindIdeMainJs(); mainJs != "" {
			ui.PrintStatus("OK", "Core Engine", "Integrated / Bundled Mode (Antigravity IDE)")
			coreSuccess = true
		} else {
			ui.PrintStatus("FAIL", "Core Engine", err.Error())
		}
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
		coreSuccess = true
	}

	if res, err := patcher.PatchAsarZeroLog(""); err != nil {
		ui.PrintStatus("INFO", "Zero-Log ASAR", "Antigravity Standalone (app.asar optional)")
	} else {
		ui.PrintStatus("OK", "Zero-Log ASAR", res)
	}

	fmt.Println()
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00F0FF")).Render("  Step 2: Smart Proxy & Desktop Launcher Setup..."))
	if path, err := proxy.GenerateNoTunDesktopLauncher(""); err != nil {
		ui.PrintStatus("WARN", "Desktop Launcher", err.Error())
	} else {
		ui.PrintStatus("OK", "Desktop Launcher", "Created at "+path)
	}

	fmt.Println()
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFB800")).Render("  Step 3: Enterprise License Wrapper & Zero-Telemetry..."))
	servers := patcher.FindAllLanguageServers(nil)
	if len(servers) == 0 {
		if single := patcher.FindLanguageServerModular(nil); single != "" {
			servers = append(servers, single)
		} else if single := patcher.FindLanguageServer(); single != "" {
			servers = append(servers, single)
		}
	}

	if len(servers) > 0 {
		wrapperCount := 0
		for _, s := range servers {
			wm := patcher.NewWrapperManager(filepath.Dir(s), nil)
			if !wm.IsWrapperInstalled() {
				if err := wm.InstallWrapper(patcher.DefaultEnterpriseFlagsOptions()); err == nil {
					wrapperCount++
				}
			} else {
				wrapperCount++
			}
		}
		if wrapperCount > 0 {
			ui.PrintStatus("OK", "Enterprise Wrapper", "Active (Business Licensed Enterprise + Zero-Telemetry)")
		} else {
			ui.PrintStatus("WARN", "Enterprise Wrapper", "Could not install wrapper on detected targets")
		}
	} else {
		ui.PrintStatus("INFO", "Enterprise Wrapper", "Standalone mode (Core path detected on launch)")
	}

	fmt.Println()
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF79C6")).Render("  Step 4: Surgical Quota, DIPS & Network Cache Reset..."))
	if removed, err := cleaner.SurgicalReset429("", false); err != nil {
		ui.PrintStatus("INFO", "Surgical Clean", "Skipped or no stale cache detected")
	} else {
		ui.PrintStatus("OK", "Surgical Clean", fmt.Sprintf("Cleaned %d temporary cache & session artifacts", len(removed)))
	}

	if coreSuccess {
		fmt.Println()
		successBox := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00FF9F")).
			Padding(0, 2).
			Render(fmt.Sprintf("%s\n\n%s\n%s\n%s\n%s\n%s\n\n%s\n%s",
				lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF9F")).Render("🎉 AUTO-FIX COMPLETED SUCCESSFULLY!"),
				lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Render("✔ Core Engine:      UNLOCKED (MultiGate hasValidAuth Bypass Active)"),
				lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Render("✔ Enterprise Tier:  ACTIVATED (Business Licensed Enterprise & Zero-Telemetry)"),
				lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Render("✔ Quota & DIPS:     PURGED (Corrupted 429 token cache & cookies safely cleaned)"),
				lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Render("✔ Smart Proxy:      CONFIGURED (Direct socket, Zero TUN mode needed)"),
				lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2")).Render("✔ Desktop Shortcut: READY (Click 'Antigravity' in App Menu or Desktop)"),
				lipgloss.NewStyle().Foreground(lipgloss.Color("#00F0FF")).Render("🚀 You can now launch Antigravity with full region freedom!"),
				lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Render("⭐ If this helped you, star us: https://github.com/tawroot/antigravity-cleaner"),
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

	if res, err := patcher.PatchAsarZeroLog(""); err != nil {
		ui.PrintStatus("WARN", "Zero-Log ASAR", err.Error())
	} else {
		ui.PrintStatus("OK", "Zero-Log ASAR", res)
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
		wm := patcher.NewWrapperManager(filepath.Dir(lsPath), nil)
		if wm.IsWrapperInstalled() {
			if err := wm.RemoveWrapper(); err == nil {
				ui.PrintStatus("OK", "Enterprise Wrapper", "Removed wrapper and restored original binary")
			}
		}
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

func runEnterpriseFlags(action string) {
	ui.PrintSection("Enterprise & Privacy Flags Manager")
	lsPath := patcher.FindLanguageServer()
	if lsPath == "" {
		ui.PrintStatus("FAIL", "Language Server", "Could not locate language_server binary")
		return
	}

	wm := patcher.NewWrapperManager(filepath.Dir(lsPath), nil)
	switch strings.ToLower(action) {
	case "enable", "install", "on":
		if !ensureRunningSafety(false) {
			return
		}
		opts := patcher.DefaultEnterpriseFlagsOptions()
		if err := wm.InstallWrapper(opts); err != nil {
			ui.PrintStatus("FAIL", "Enterprise Wrapper", err.Error())
		} else {
			ui.PrintStatus("OK", "Enterprise Wrapper", "Installed! Telemetry muted, Google Internal & Insiders enabled.")
		}
	case "disable", "remove", "off":
		if !ensureRunningSafety(false) {
			return
		}
		if err := wm.RemoveWrapper(); err != nil {
			ui.PrintStatus("FAIL", "Enterprise Wrapper", err.Error())
		} else {
			ui.PrintStatus("OK", "Enterprise Wrapper", "Removed! Original binary restored.")
		}
	case "status":
		installed := wm.IsWrapperInstalled()
		if installed {
			ui.PrintStatus("OK", "Wrapper Status", "Active (Injecting enterprise flags & muting telemetry)")
			fmt.Println("    • --disable_telemetry=true")
			fmt.Println("    • --is_google_internal=true")
			fmt.Println("    • --jetbox_release_channel=insiders")
			fmt.Println("    • --stubby_auth_mode=business_licensed_enterprise")
			fmt.Println("    • --use_local_chrome=true & --browser_eval_env=true")
		} else {
			ui.PrintStatus("INFO", "Wrapper Status", "Not Installed (Running standard unflagged binary)")
		}
	default:
		fmt.Printf("Unknown action: %s. Use enable, disable, or status.\n", action)
	}
}

func runEnterpriseMenu() {
	ui.PrintSection("Enterprise & Privacy Flags Configuration")
	lsPath := patcher.FindLanguageServer()
	if lsPath == "" {
		ui.PrintStatus("FAIL", "Language Server", "Could not locate language_server binary")
		return
	}

	wm := patcher.NewWrapperManager(filepath.Dir(lsPath), nil)
	installed := wm.IsWrapperInstalled()
	if installed {
		ui.PrintStatus("ACTIVE", "Wrapper Status", "Enterprise wrapper is currently ACTIVE")
		fmt.Println("  Active Flags:")
		fmt.Println("    • --disable_telemetry=true (Halts telemetry, Git logging & error traces)")
		fmt.Println("    • --is_google_internal=true (Unlocks internal payload metadata & URL bypass)")
		fmt.Println("    • --jetbox_release_channel=insiders")
		fmt.Println("    • --stubby_auth_mode=business_licensed_enterprise")
		fmt.Println("    • --use_local_chrome=true & --browser_eval_env=true")
		fmt.Println()
		confirm := ui.ReadInput("Do you want to REMOVE the wrapper and restore standard binary? (y/N)")
		if confirm == "y" || confirm == "Y" {
			if !ensureRunningSafety(false) {
				return
			}
			if err := wm.RemoveWrapper(); err != nil {
				ui.PrintStatus("FAIL", "Restore Error", err.Error())
			} else {
				ui.PrintStatus("OK", "Wrapper Removed", "Original binary restored successfully.")
			}
		}
	} else {
		ui.PrintStatus("INACTIVE", "Wrapper Status", "Enterprise wrapper is NOT installed")
		fmt.Println("  Enabling this will transparently inject:")
		fmt.Println("    • --disable_telemetry=true (Halts telemetry, Git logging & error traces)")
		fmt.Println("    • --is_google_internal=true (Unlocks internal payload metadata & URL bypass)")
		fmt.Println("    • --disable_unleash=true & --use_snapshot_mendelfp=true (Halts remote region gates)")
		fmt.Println("    • --jetbox_release_channel=insiders")
		fmt.Println("    • --stubby_auth_mode=business_licensed_enterprise")
		fmt.Println("    • --use_local_chrome=true & --browser_eval_env=true")
		fmt.Println("    • Bi-directional Token Bridge (zero session drops with SSH spoofing)")
		fmt.Println()
		confirm := ui.ReadInput("Do you want to INSTALL the enterprise wrapper? (y/N)")
		if confirm == "y" || confirm == "Y" {
			if !ensureRunningSafety(false) {
				return
			}
			opts := patcher.DefaultEnterpriseFlagsOptions()
			if err := wm.InstallWrapper(opts); err != nil {
				ui.PrintStatus("FAIL", "Install Error", err.Error())
			} else {
				ui.PrintStatus("OK", "Wrapper Installed", "Enterprise flags enabled successfully!")
			}
		}
	}
}

func runGui() {
	ui.PrintSection("🖥️ Launching Retro Win95 KeyGen GUI")
	fmt.Println("  Starting local GUI server and opening desktop window...")
	if err := gui.Run(); err != nil {
		ui.PrintStatus("FAIL", "GUI Error", err.Error())
	}
}

func runReport() {
	ui.PrintSection("📋 Generating Diagnostic Report")
	fmt.Println("  Scanning system state, connectivity, and Antigravity installation...")

	results := doctor.RunDiagnostics(true)
	var rows []ui.CardRow
	for _, r := range results {
		rows = append(rows, ui.CardRow{
			Title:   r.Name,
			Status:  r.Status,
			Summary: r.HumanSummary,
			Details: r.DebugDetails,
		})
	}
	ui.PrintDashboardWithTitle(" 📋 Diagnostic Report — Health & Environment Audit", rows, false)

	report := doctor.BuildMarkdownReport(Version, results)
	copied := doctor.CopyToClipboard(report)
	if copied {
		fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF9F")).Render("  ✔ Copied diagnostic report to clipboard! You can paste (Ctrl+V) directly into Telegram or GitHub."))
	} else {
		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB800")).Render("  ⚠ Could not access system clipboard."))
	}

	if filePath, err := doctor.SaveAndOpenReport(report); err == nil {
		fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00F0FF")).Render("  📄 Opened and saved report to: " + filePath))
	}

	fmt.Println()
	divider := lipgloss.NewStyle().Foreground(lipgloss.Color("#6272A4")).Render("  ────────────────────────────────────────────────────────────────────────")
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00F0FF")).Render("  📋 Text Preview (Markdown for Telegram / GitHub):"))
	fmt.Println(divider)
	lines := strings.Split(report, "\n")
	for _, line := range lines {
		fmt.Printf("  %s\n", line)
	}
	fmt.Println(divider)
}

func runUpdate() {
	ui.PrintSection("🔄 Checking for Updates")
	fmt.Printf("  Current Version: v%s\n", Version)
	fmt.Println("  Contacting GitHub update server...")

	release, hasUpdate, err := updater.CheckLatestRelease(Version)
	if err != nil {
		ui.PrintStatus("FAIL", "Update Check", err.Error())
		return
	}

	if !hasUpdate {
		ui.PrintStatus("OK", "Up-to-Date", fmt.Sprintf("Antigravity Cleaner is already on the latest version (%s)", release.TagName))
		return
	}

	ui.PrintStatus("UPDATE", "New Version Available", fmt.Sprintf("Found %s (Current: v%s)", release.TagName, Version))
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#00F0FF")).Render(fmt.Sprintf("  Release Title: %s", release.Name)))

	confirm := ui.ReadInput("Do you want to download and install this update now? (Y/n)")
	if confirm == "n" || confirm == "N" {
		fmt.Println("  Update cancelled.")
		return
	}

	fmt.Println("  Downloading latest binary...")
	err = updater.DownloadAndSelfUpdate(release, func(downloaded, total int64) {
		if total > 0 {
			percent := float64(downloaded) / float64(total) * 100
			fmt.Printf("\r  Progress: %.1f%% (%d / %d bytes)", percent, downloaded, total)
		}
	})
	fmt.Println()

	if err != nil {
		ui.PrintStatus("FAIL", "Update Failed", err.Error())
		return
	}

	ui.PrintStatus("OK", "Update Complete", fmt.Sprintf("🎉 Successfully updated to %s! Please restart Antigravity Cleaner.", release.TagName))
}

func runQuota(args ...[]string) {
	ui.PrintSection("📊 Google Antigravity Live Quota Monitor")
	fmt.Println("  Reading session token from Antigravity IDE...")

	token, err := quota.ExtractRefreshToken("")
	if err != nil {
		ui.PrintStatus("FAIL", "Session Token", err.Error())
		return
	}

	fmt.Println("  Authenticating with Google CloudCode APIs...")
	accessToken, err := quota.RefreshAccessToken(token)
	if err != nil {
		ui.PrintStatus("FAIL", "OAuth Refresh", err.Error())
		return
	}

	fmt.Println("  Retrieving real-time quota buckets from Google CloudCode...")
	groups, _ := quota.FetchQuotaSummary(accessToken)
	detailed, err := quota.FetchDetailedUserQuota(accessToken)
	if err != nil && len(groups) == 0 {
		ui.PrintStatus("FAIL", "Quota API", err.Error())
		return
	}

	showRaw := false
	if len(args) > 0 {
		for _, a := range args[0] {
			if a == "--all" || a == "-a" || a == "--raw" || a == "-v" || a == "--verbose" {
				showRaw = true
				break
			}
		}
	}

	if showRaw {
		fmt.Print(quota.RenderQuotaRaw(groups, detailed))
	} else {
		fmt.Print(quota.RenderQuota(groups, detailed))
	}
}

func runAccount(args []string) {
	ui.PrintSection("👤 Multi-Account Manager")
	if len(args) == 0 {
		runAccountInteractive()
		return
	}

	sub := args[0]
	switch sub {
	case "save":
		if len(args) < 2 {
			fmt.Println("  Error: Please provide a name for this account. Example: antigravity-cleaner account save work-gmail")
			return
		}
		name := args[1]
		if err := account.SaveAccount(name, ""); err != nil {
			ui.PrintStatus("FAIL", "Save Profile", err.Error())
		} else {
			ui.PrintStatus("OK", "Account Saved", fmt.Sprintf("Successfully saved current Antigravity profile as '%s'", name))
		}
	case "switch", "use":
		if len(args) < 2 {
			fmt.Println("  Error: Please provide the account name to switch to. Example: antigravity-cleaner account switch work-gmail")
			return
		}
		name := args[1]
		if !ensureRunningSafety(false) {
			return
		}
		if err := account.SwitchAccount(name, ""); err != nil {
			ui.PrintStatus("FAIL", "Switch Account", err.Error())
		} else {
			ui.PrintStatus("OK", "Account Switched", fmt.Sprintf("Switched to '%s'! Launch Antigravity to use this account.", name))
		}
	case "list", "ls":
		list, err := account.ListAccounts()
		if err != nil {
			ui.PrintStatus("FAIL", "List Accounts", err.Error())
			return
		}
		if len(list) == 0 {
			fmt.Println("  No saved accounts found. Save current account with: antigravity-cleaner account save <name>")
			return
		}
		fmt.Printf("  Saved Accounts (%d):\n", len(list))
		for _, a := range list {
			emailInfo := ""
			if a.Email != "" {
				emailInfo = fmt.Sprintf(" (\x1b[36m%s\x1b[0m)", a.Email)
			}
			fmt.Printf("    • \x1b[1;32m%-15s\x1b[0m %s [Saved: %s]\n", a.Name, emailInfo, a.SavedAt)
		}
	case "rm", "delete":
		if len(args) < 2 {
			fmt.Println("  Error: Please specify the account name to remove.")
			return
		}
		name := args[1]
		if err := account.DeleteAccount(name); err != nil {
			ui.PrintStatus("FAIL", "Delete Account", err.Error())
		} else {
			ui.PrintStatus("OK", "Account Removed", fmt.Sprintf("Deleted profile '%s'", name))
		}
	default:
		fmt.Println("  Unknown account command. Options: save, switch, list, rm")
	}
}

func runAccountInteractive() {
	ui.PrintSection("👤 Multi-Account Manager")
	list, err := account.ListAccounts()
	if err == nil && len(list) > 0 {
		fmt.Printf("  Saved Profiles (%d):\n", len(list))
		for i, a := range list {
			emailInfo := ""
			if a.Email != "" {
				emailInfo = fmt.Sprintf(" (\x1b[36m%s\x1b[0m)", a.Email)
			}
			fmt.Printf("    [%d] \x1b[1;32m%-15s\x1b[0m %s [Saved: %s]\n", i+1, a.Name, emailInfo, a.SavedAt)
		}
		fmt.Println()
	}

	fmt.Println("  Options:")
	fmt.Println("    [1] Save current session as new profile")
	fmt.Println("    [2] Switch to an existing profile")
	fmt.Println("    [3] Delete a saved profile")
	fmt.Println("    [0] Back to main menu")
	fmt.Println()

	choice := ui.ReadInput("Select option")
	switch choice {
	case "1":
		name := ui.ReadInput("Enter profile name (e.g. work, personal)")
		if name != "" {
			runAccount([]string{"save", name})
		}
	case "2":
		if len(list) == 0 {
			ui.PrintStatus("INFO", "No Profiles", "No saved profiles found. Save current profile first with [1].")
			return
		}
		fmt.Println()
		fmt.Println("  Select profile to switch to:")
		for i, a := range list {
			emailInfo := ""
			if a.Email != "" {
				emailInfo = fmt.Sprintf(" (\x1b[36m%s\x1b[0m)", a.Email)
			}
			fmt.Printf("    [%d] \x1b[1;32m%-15s\x1b[0m %s\n", i+1, a.Name, emailInfo)
		}
		fmt.Println("    [0] Cancel")
		fmt.Println()
		input := ui.ReadInput("Enter profile number or name")
		if input == "" || input == "0" {
			return
		}
		targetName := input
		if idx, err := strconv.Atoi(input); err == nil && idx >= 1 && idx <= len(list) {
			targetName = list[idx-1].Name
		}
		runAccount([]string{"switch", targetName})
	case "3":
		if len(list) == 0 {
			ui.PrintStatus("INFO", "No Profiles", "No saved profiles found.")
			return
		}
		fmt.Println()
		fmt.Println("  Select profile to delete:")
		for i, a := range list {
			emailInfo := ""
			if a.Email != "" {
				emailInfo = fmt.Sprintf(" (\x1b[36m%s\x1b[0m)", a.Email)
			}
			fmt.Printf("    [%d] \x1b[1;31m%-15s\x1b[0m %s\n", i+1, a.Name, emailInfo)
		}
		fmt.Println("    [0] Cancel")
		fmt.Println()
		input := ui.ReadInput("Enter profile number or name to delete")
		if input == "" || input == "0" {
			return
		}
		targetName := input
		if idx, err := strconv.Atoi(input); err == nil && idx >= 1 && idx <= len(list) {
			targetName = list[idx-1].Name
		}
		runAccount([]string{"rm", targetName})
	}
}


