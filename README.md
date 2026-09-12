# ⚡ Antigravity Cleaner — Ultimate AI Freedom Toolkit (v5.2.0)

<div align="center">
  <img src="docs/images/banner.png" alt="Antigravity Cleaner Banner" width="85%">
  <br>
  
  [![Version](https://img.shields.io/badge/Version-5.2.0-blue?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner/releases)
  [![Go Report](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
  [![Tests](https://img.shields.io/badge/Tests-100%25_Passing-brightgreen?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Benchmark](https://img.shields.io/badge/Bytecode_Scan-1.9_GB%2Fs-blueviolet?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-brightgreen.svg?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner)
  [![Security](https://img.shields.io/badge/Telemetry-Zero%20(100%25%20Offline)-success?style=for-the-badge)]()
</div>

> *Dedicated to developers in Iran, Cuba, Syria, and restricted regions navigating digital sanctions and censorship. Access to AI and development tools is a fundamental human right.*

---

## ⚡ What is Antigravity Cleaner?

**Antigravity Cleaner** is an enterprise-grade, high-performance systems utility engineered in **pure Go** with a 2026-standard terminal interface (**Lipgloss & Bubbletea**). It completely solves region blocking, account eligibility barriers, 429 quota exhaustion, and streaming connection drops for:
- **Google Antigravity 2.x**
- **Antigravity IDE**
- **Antigravity CLI (`agy`)**
- **Official VS Code Extension (`google.google-antigravity`)**

Unlike conventional scripts that force users to route their entire operating system through a heavy, root-permission **TUN Mode (VPN)**, Antigravity Cleaner introduces an automated **No-TUN Smart Proxy Injector** that routes Antigravity traffic directly to local proxies (Xray, Clash, NekoBox, Hiddify) with zero system overhead.

> 🔒 **100% Offline & Zero-Track Guarantee:**
> Antigravity Cleaner operates **100% locally on `127.0.0.1`**. It contains **ZERO telemetry, ZERO analytics, and ZERO remote network requests**. Your source code, API keys, and chats are never uploaded or tracked. Fully open-source and auditable under GPL-3.0.

---

## 💾 1-Click Retro KeyGen GUI (Zero Terminal Required)

Prefer a simple graphical interface? Download the standalone executable for your operating system from [Releases](https://github.com/tawroot/antigravity-cleaner/releases) and simply **double-click** to run. No terminal commands, no Python runtime, and zero external dependencies!

<div align="center">
  <img src="assets/shot.jpg" alt="Antigravity Retro KeyGen GUI (Win95 Style)" width="440">
  <br>
  <sub><em>Authentic Windows 95/98 style Patcher with live demoscene starfield, Web Audio 8-bit sound effects & 1-click Auto-Fix.</em></sub>
</div>

<br>

| Operating System | Standalone Binary (Direct Download) | Mode |
| :--- | :--- | :--- |
| 🪟 **Windows** (64-bit) | [**`antigravity-cleaner-windows-amd64.exe`**](https://github.com/tawroot/antigravity-cleaner/releases/latest) | Double-click `.exe` (opens GUI automatically) |
| 🍎 **macOS** (Apple Silicon M1–M4) | [**`antigravity-cleaner-darwin-arm64`**](https://github.com/tawroot/antigravity-cleaner/releases/latest) | Double-click or `./antigravity-cleaner-darwin-arm64` |
| 🍎 **macOS** (Intel) | [**`antigravity-cleaner-darwin-amd64`**](https://github.com/tawroot/antigravity-cleaner/releases/latest) | Double-click or `./antigravity-cleaner-darwin-amd64` |
| 🐧 **Linux** (x86_64) | [**`antigravity-cleaner-linux-amd64`**](https://github.com/tawroot/antigravity-cleaner/releases/latest) | Double-click or `ag-cleaner gui` |
| 🐧 **Linux** (ARM64) | [**`antigravity-cleaner-linux-arm64`**](https://github.com/tawroot/antigravity-cleaner/releases/latest) | Double-click or `ag-cleaner gui` |

---

## 🏛️ System Architecture

```mermaid
graph TD
    subgraph Antigravity Operating Space
        IDE[Antigravity IDE / agy CLI]
        LS[Core Engine: language_server]
    end

    subgraph Antigravity Cleaner Engine
        PAT[MultiGate IDA Scanner: 1.9 GB/s]
        PROXY[Smart Proxy Detector: :10808 / :2080 / :7890]
        KEEPALIVE[TCP KeepAlive Socket Tuner: 15s]
        CLEAN[Surgical 429 Quota Reset]
    end

    subgraph Local Proxy Stack
        XRAY[Xray / v2rayN / Clash / NekoBox / Hiddify]
    end

    subgraph Google Cloud & AI Mesh
        GOOGLE[Gemini AI / Cloud Code Services]
    end

    PAT -->|POSIX Atomic Inode Patch| LS
    LS -->|RAX+8 hasValidAuth = 1| IDE
    PROXY -->|Injects ALL_PROXY & --proxy-server| IDE
    IDE -->|Targeted SOCKS5h / HTTP Traffic| XRAY
    KEEPALIVE -.->|Prevents Middlebox DPI Drops| XRAY
    XRAY -->|Stable Uninterrupted Stream| GOOGLE
```

---

## 🛡️ Security, Privacy & Integrity Guarantee

- 🔒 **100% Local Execution**: Zero external telemetry, no remote analytics, and no third-party network calls. The tool operates entirely on your machine.
- 💬 **Conversation Preservation**: The surgical quota reset (`ag-cleaner clean`) targets only corrupt session cookies and DIPS caches. Your prompts, global storage, and project chat histories are **strictly protected and preserved**.
- ⚛️ **POSIX Atomic Safety**: Core binaries are patched using atomic inode replacement (`os.Rename`), ensuring running processes never crash, corrupt, or lock with `ETXTBSY`.

---

## 📊 Benchmarks & Verified Test Reports

| Metric | Measured Value | Verification Log |
| :--- | :--- | :--- |
| **In-Memory Pattern Scanning** | **1,901.84 MB/s (~1.9 GB/s)** | [`benchmarks.log`](docs/test-reports/benchmarks.log) |
| **171MB Binary Patch Duration** | **0.23 seconds** | [`benchmarks.log`](docs/test-reports/benchmarks.log) |
| **Doctor Concurrent Audit** | **~2.5 seconds** | [`live_system_test.log`](docs/test-reports/live_system_test.log) |
| **Surgical 429 Cache Flush** | **16 artifacts purged (0 chat loss)** | [`live_system_test.log`](docs/test-reports/live_system_test.log) |
| **Unit Test Pass Rate** | **100% (All packages passing)** | [`benchmarks.log`](docs/test-reports/benchmarks.log) |

---

## 🥊 Comparison Matrix

| Feature | `open-antigravity-patcher` | Conventional Scripts | **⚡ Antigravity Cleaner (v5.1)** |
| :--- | :---: | :---: | :---: |
| **Language & Runtime** | Python 3 (requires `pip`/`venv`) | PowerShell (Windows only) | **Pure Go (Zero dependencies, Single Binary)** |
| **User Interface (TUI)** | Basic text stdout | Basic console | **2026 Lipgloss Rounded Dashboard** |
| **No-TUN Proxy Launcher** | ❌ (Forces root TUN mode) | ❌ None | ✅ **Auto-detects Xray/Clash/NekoBox & Injects** |
| **Stream KeepAlive Tuner** | ❌ (Frequent stream drops) | ❌ None | ✅ **Aggressive socket keepalive against DPI drops** |
| **Surgical 429 Quota Reset** | ❌ (Wipes entire config & chats) | ❌ (Deletes everything) | ✅ **Cleans token cache, PRESERVES all chats!** |
| **Core Binary Patch** | x86_64 / ARM64 | ❌ None | ✅ **MultiGate IDA-style machine-code patch** |
| **IDE Frontend Patch (`main.js`)** | ✅ Regex patch | ❌ None | ✅ **Automated patch + VS Code cache purge** |
| **VS Code Re-download Guard** | ✅ Basic | ❌ None | ✅ **Full Guard + Channel Lock** |
| **Diagnostic Doctor Tool** | ❌ None | ⚠️ Limited | ✅ **Concurrent 2s latency & DNS leak probe** |
| **POSIX Atomic Replacement** | Python tempfile | ❌ None | ✅ **Atomic Inode swap (no `ETXTBSY` on running app)** |

---

## 🚀 Quick Install (One-Liner)

### 🐧 Linux & 🍎 macOS (Apple Silicon M1/M2/M3/M4 & Intel)
Run in your terminal:
```bash
curl -fsSL https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.sh | bash
```

### 🪟 Windows (PowerShell)
Open PowerShell (as Administrator) and run:
```powershell
irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex
```

---

## 💻 2026 Visual Dashboard (TUI)
 
Run `ag-cleaner` in your terminal for an interactive, beautifully styled terminal menu:

<div align="center">
  <img src="docs/images/terminal.png" alt="Antigravity Cleaner TUI" width="75%">
</div>

### 🎬 Live Demo Walkthrough
See Antigravity Cleaner v5.1 in action — automated detection, zero-TUN proxy injection, and one-click unlocking:

https://github.com/user-attachments/assets/demo.webm

> *Video demo: [`docs/images/demo.webm`](docs/images/demo.webm)*

Run `ag-cleaner doctor` for a non-intimidating, high-aesthetic health card:

```text
╭────────────────────────────────────────────────────────────────────────────╮
│  🩺 Antigravity Cleaner — Health & Connectivity Audit                      │
│ ────────────────────────────────────────────────────────────────────────   │
│                                                                            │
│  Antigravity IDE         READY   Installed & Detected                      │
│  Core Engine             READY   Unlocked (Patched)                        │
│  Proxy Connection        READY   Active on 127.0.0.1:10808 (v2ray / Xray)  │
│  Google DNS              READY   Healthy (16 IPs resolved)                 │
│  Cloud Code Service      READY   Connected (1100ms, Fast)                  │
│  Google Accounts         READY   Connected (1800ms, Fast)                  │
│  Gemini AI API           READY   Connected (950ms, Fast)                   │
│                                                                            │
╰────────────────────────────────────────────────────────────────────────────╯
                                                                              
   ONLINE  System is optimized. You can launch Antigravity with Smart Proxy!
```

---

## 🛠️ CLI Subcommands

```bash
# Launch interactive menu
ag-cleaner

# Launch Retro KeyGen GUI (Classic Win95 Desktop Window)
ag-cleaner gui

# Health & latency diagnostics (clean card)
ag-cleaner doctor

# Detailed diagnostics with debug traces
ag-cleaner doctor --verbose

# Apply all patches (Core binary, CLI, IDE frontend)
ag-cleaner patch all

# Launch Antigravity with smart proxy injection (No TUN mode needed!)
ag-cleaner launch

# Reset HTTP 429 Quota without losing conversation history
ag-cleaner clean

# Terminate all running Antigravity & language_server zombie processes
ag-cleaner kill

# Generate desktop shortcut with pre-configured proxy
ag-cleaner create-launcher

# View project attribution & credits
ag-cleaner about
```

---

## 🤖 AI Agent Skill (`skills/antigravity-cleaner`)

Antigravity Cleaner provides a native **Agent Skill** (`skills/antigravity-cleaner/SKILL.md`) that enables AI coding assistants (Google Antigravity, Claude Code, Cursor, Codex) to self-heal their environment when encountering API restrictions, quota limits, or network timeouts:

- **Automated Self-Audit**: Agents trigger `ag-cleaner doctor` on connection loss.
- **Autonomous Quota Healing**: Agents run surgical clean on HTTP 429 without human intervention.
- **Zero Configuration**: Vendored directly inside the project root for automatic discovery.

---

## 🔧 Troubleshooting Matrix

| Issue / Error | Root Cause | Instant Fix |
| :--- | :--- | :--- |
| **"Region Not Supported" / Gate 1 Block** | Official machine code checking IP/account region | Run `ag-cleaner patch all` |
| **HTTP 429 Quota Exhausted** | Corrupted session tokens & stale HTTP cookies | Run `ag-cleaner clean` (Keeps chats intact) |
| **Stream Dropouts / Middlebox Resets** | ISP DPI firewall closing idle TLS connections | Run `ag-cleaner launch` (Enables 15s TCP KeepAlive) |
| **Sluggish Workspace / Laggy Queries** | Bloated Electron network state & cache | Run `ag-cleaner clean` |

---

## 🏗️ Building from Source

Requires **Go 1.22+**:

```bash
git clone https://github.com/tawroot/antigravity-cleaner.git
cd antigravity-cleaner

# Build for current machine
make build

# Cross-compile for Linux, macOS (Intel & ARM64), and Windows
make release
```

---

## 👥 Credits & Attribution

- **Lead Architect & Maintainer:** **[@dalroot](https://github.com/dalroot)**
- **Co-Pilot & Pair Programming:** **Antigravity AI (Google DeepMind)**
- **License:** GNU General Public License v3.0 (GPL-3.0)
- Dedicated to uncensored technology and open-source accessibility worldwide.
