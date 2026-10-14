# 🤝 Antigravity Cleaner Pro — Engineering Handoff & Session State

> **Last Updated:** October 2026 (v5.3.0 Stable)  
> **Repository:** [`tawroot/antigravity-cleaner`](https://github.com/tawroot/antigravity-cleaner)  
> **Lead Architect & Maintainer:** [@dalroot](https://github.com/dalroot)  
> **Community Channel:** Telegram [@dallrep](https://t.me/dallrep) | [GitHub Discussions](https://github.com/tawroot/antigravity-cleaner/discussions)  
> **Primary Technology Stack:** Pure Go (1.22+), Charmbracelet Lipgloss & Bubbletea (TUI), Embedded Web Audio/HTML5 Starfield Retro KeyGen GUI (Win95 style), POSIX Atomic Inode Patching.

---

## 🧭 Purpose of this Document

This **Handoff Document** is engineered to enable any future AI assistant, subagent, or human developer to **instantly inherit 100% of the project context, technical decisions, active architecture, and upcoming roadmap goals** without losing momentum, asking redundant questions, or violating critical safety constraints.

Whenever starting a new chat or continuing work in a fresh context, **read this file first**.

---

## 🏛️ Project Overview & Value Proposition

**Antigravity Cleaner Pro** is a standalone, enterprise-grade systems utility developed to liberate developer environments from regional blocks, account licensing gates, quota exhaustion (HTTP 429), and connection stability drops across:
1. **Google Antigravity 2.x**
2. **Antigravity IDE**
3. **Antigravity CLI (`agy`)**
4. **Official VS Code Extension (`google.google-antigravity`)**

### Core Philosophy
* **Zero System TUN / VPN Overhead:** Instead of routing the whole OS through heavy TUN/VPN drivers, Antigravity Cleaner uses a **No-TUN Smart Proxy Injector** that routes only Antigravity traffic directly to local proxies (Xray, Clash, NekoBox, Hiddify on `:10808`, `:2080`, `:7890`, etc.).
* **100% Local & Zero-Telemetry:** All operations run locally on `127.0.0.1`. Zero analytics, zero metrics, zero remote tracking.
* **Non-Destructive & Atomic:** Running IDE sessions are never violently killed. Patches are swapped atomically via POSIX Inodes, and chat histories are strictly preserved.
* **Dual UI Formats:**
  - 🖥️ **1-Click Retro Win95 KeyGen GUI:** Nostalgic standalone portable window with live demoscene starfield, 8-bit sound effects, and zero terminal prerequisites.
  - 💻 **2026 TUI Terminal Dashboard:** Interactive keyboard-driven UI powered by Go's Bubbletea and Lipgloss.

---

## ⚡ Core 4-Target Unlocking Architecture

The tool applies surgical modifications across four critical system layers:

| Layer | Target Component | Mechanism & Technical Details |
| :--- | :--- | :--- |
| **1. Core Engine** | `language_server` (Binary) | MultiGate machine-code byte patch (`RAX+8 hasValidAuth = 1`) reversing eligibility gate in binary RAM. |
| **2. IDE Frontend** | `main.js` (JavaScript AST) | Unlocks `isGoogleInternal` flag in Electron frontend & flushes V8 compile caches (`Code Cache`, `CachedData`). |
| **3. Chat UI Gate** | `workbench.desktop.main.js` | Client-side entitlement gate interception for AI sidebar and chat panels. |
| **4. Zero-Telemetry Wrapper** | `language_server` Launcher Wrapper | Transparent wrapper injecting 7 enterprise isolation flags: `--disable_telemetry=true`, `--stubby_auth_mode=business_licensed_enterprise`, etc. |

---

## 📂 Repository Directory Layout

```
.
├── cmd/
│   └── antigravity/            # CLI and GUI entrypoints
│       ├── main.go             # Application bootstrap & subcommand router
│       ├── cmd.go              # CLI flags, doctor, clean, patch, and run commands
│       └── gui.go              # Win95 GUI webview/browser server & bridge
├── pkg/
│   ├── cleaner/                # 16-target surgical quota & cache purger
│   │   ├── cleaner.go          # DIPS, GPU Cache, rate-limit cookie reset
│   │   └── cleaner_test.go     # Unit tests verifying chat history preservation
│   ├── patcher/                # MultiGate IDA binary scanner & atomic patcher
│   │   ├── patcher.go          # Bytecode search & atomic Inode file replacement
│   │   └── patcher_test.go     # 1.9 GB/s pattern matching benchmark tests
│   ├── proxy/                  # Local proxy detector & KeepAlive injector
│   │   ├── detector.go         # Auto-scans ports :10808, :2080, :7890, etc.
│   │   └── keepalive.go        # 15s TCP KeepAlive socket tuner preventing DPI drop
│   ├── doctor/                 # Comprehensive system diagnostics & audit checks
│   │   └── doctor.go           # Parallel goroutine health checks
│   ├── gui/                    # Embedded Retro Win95 GUI assets & HTTP bridge
│   │   ├── assets/             # HTML5, CSS3, Win95 UI, WebAudio sound FX
│   │   └── server.go           # Local lightweight IPC server
│   └── ui/                     # Bubbletea & Lipgloss TUI terminal dashboard
├── skills/
│   └── antigravity-cleaner/    # Standardized autonomous AI Agent Skill
│       └── SKILL.md            # LLM-executable operational specification
├── docs/                       # Architecture diagrams, test logs, error screenshots
│   └── test-reports/           # benchmarks.log, live_system_test.log
├── .github/
│   ├── ISSUE_TEMPLATE/         # Streamlined bug_report.yml and feature_request.yml
│   └── workflows/              # GitHub Actions cross-compile matrix
├── ROADMAP.md / ROADMAP.fa.md  # Horizon milestones (v5.4.0 & v6.0.0)
├── README.md / README.fa.md    # Primary bilingual documentations
└── Makefile                    # Standard build & release targets
```

---

## 🔒 Critical Invariants & Rules for Developers and AI Agents

When modifying this repository, **you MUST strictly adhere to the following constraints**:

1. **Persian Right-to-Left (RTL) Rule:**
   - Any response or documentation generated in Persian (Farsi) MUST be completely enclosed inside `<div dir="rtl">...</div>`.
2. **Preserve User Data & Chat History:**
   - Never write cleaner code that removes `User/globalStorage` or project state databases. Quota cleaning is strictly limited to ephemeral session cookies, DIPS, and network caches.
3. **No Violent Process Kills:**
   - Antigravity Cleaner uses atomic Inode replacement (`os.Rename` on Unix, move-and-replace on Windows). Do **NOT** introduce blind `pkill -9` or taskkill commands during automatic fixing, as this abruptly closes active user IDE sessions.
4. **Zero Telemetry Integrity:**
   - Never add network tracking, third-party analytics, remote pings, or telemetry reporting to the Go code. The tool must remain 100% offline and local on `127.0.0.1`.
5. **International Date Standardization:**
   - All dates in user-facing roadmaps and releases must use international standard format (`Dec 2025`, `Sep 2026`, `Oct 2026`). Avoid Iranian solar calendar dates in roadmap timelines.
6. **Multi-Ecosystem Research Scope:**
   - In documentation discussing the private research laboratory, accurately describe it as a **Multi-Ecosystem Applied AI Security Lab** covering Google DeepMind (Gemini / Antigravity), OpenAI, Anthropic (Claude Code), and Cursor, with Antigravity Cleaner Pro as the foundational flagship release.

---

## 🧪 Verification & Build Commands

Always run the test suite and verify clean builds before pushing commits:

```bash
# 1. Run all unit tests & benchmarks
go test -v ./...

# 2. Build local binary
go build -o bin/ag-cleaner ./cmd/antigravity

# 3. Test CLI Subcommands
./bin/ag-cleaner doctor
./bin/ag-cleaner status

# 4. Cross-compile all platform binaries (Windows, Linux, macOS ARM64 & Intel)
make release
```

---

## 🚀 Active Roadmap Priorities (Upcoming Work)

### v5.4.0 Horizon (In Progress / Next to Implement)
1. **Intelligent Dynamic Egress IP Detector & Error Normalizer:**
   - Detect real egress IP and map upstream cryptic Google codes (like `FAILED_PRECONDITION: User location is not supported #400` or `License token #3501`) into humanized, dignified English diagnostic notifications with 1-click recovery actions.
2. **Electron Broken-Pipe Sanitization (`write EIO` Prevention):**
   - Intercept broken pipe errors in `electron-log` transport when launched detached from terminals, suppressing annoying JS popups.
3. **Exhaustive Patch Whitepaper:**
   - Technical deep-dive documentation detailing exact byte mutations, IDA signatures, and AST node replacements for each of the core 4 patches.

### v6.0.0 Horizon (Future Architecture)
- **Universal BYOM Engine (Bring Your Own Model):** Local proxy bridge routing Antigravity IDE requests to local Ollama, vLLM, DeepSeek, or custom OpenRouter endpoints.
- **Offline Token Vault:** Local encrypted cache for session tokens.

---

## 💡 Quick Tips for New AI Sessions
- If a user asks to run an auto-fix, refer to `skills/antigravity-cleaner/SKILL.md`.
- If a user asks about the GitHub branch warning (`Your main branch isn't protected`), remind them it's a security best-practice suggestion and recommend enabling *Block force pushes* and *Do not allow deletions* without requiring PR approvals for solo workflow.
- All core code resides in pure Go packages under `pkg/`.
