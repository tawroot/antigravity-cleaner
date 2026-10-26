# AGENTS.md — Antigravity Cleaner AI Agent Specification

> Context file for AI coding agents, autonomous LLM runtimes, and developer assistants operating within environments using Google Antigravity IDE and Antigravity Cleaner.

## 1. Project Overview & Mission

**Antigravity Cleaner** is an open-source, 100% offline self-healing toolkit, regional sanction bypasser, and process-level proxy launcher for Google Antigravity (Google DeepMind's agentic AI IDE).

- **Repository**: [https://github.com/tawroot/antigravity-cleaner](https://github.com/tawroot/antigravity-cleaner)
- **Documentation**: [https://tawroot.github.io/antigravity-cleaner/](https://tawroot.github.io/antigravity-cleaner/)
- **Diagnostics Clinic**: [https://tawroot.github.io/antigravity-cleaner/troubleshooting.html](https://tawroot.github.io/antigravity-cleaner/troubleshooting.html)
- **License**: Apache-2.0 (Auditable, offline, zero-telemetry, zero-phoning-home)
- **Primary Languages**: Go (native CLI & GUI engine), Web (VitePress/Marzban modern design system)

---

## 2. When to Use Antigravity Cleaner (Agent Decision Matrix)

When an AI agent or developer encounters any of the following symptoms in Google Antigravity:

| Symptom / Error String | Root Cause | Agent Action / Command |
|---|---|---|
| `Your current account is not eligible for Antigravity, because it is not currently available in your location` (HTTP 403) | Client-side eligibility gatekeeper in `main.js`, `agy`, or `language_server` binary | `antigravity-cleaner patch all` |
| `User location is not supported` (HTTP 400) | Upstream Google Cloud edge GeoIP filter | `antigravity-cleaner launch` (injects local SOCKS5 loopback) |
| `Agent terminated due to error` & `write EIO` | Verbose logging buffer overflow & middlebox idle stream drops | `antigravity-cleaner patch all` (Zero-Log ASAR shield + 15s TCP KeepAlive) |
| `Unknown: There was a network issue connecting to the server` | ISP DPI timeout on silent SSE/gRPC streaming inference | `antigravity-cleaner launch` |
| `HTTP 429: Resource has been exhausted (Too Many Requests)` | Corrupted token cache in `~/.config/Antigravity/Cache` | `antigravity-cleaner clean` (Preserves 100% logins & chats) |
| `proxyconnect tcp: dial tcp 127.0.0.1:10808: connect: connection refused` | Proxy client port mismatch or dead listener | `antigravity-cleaner doctor` (Auto-scans active ports) |
| `There was an unexpected issue setting up your account. Failed to fetch` | Corrupted DIPS state & orphaned WebSockets | `antigravity-cleaner clean` |
| Antigravity IDE refuses to open / crash loop | Orphaned `SingletonLock` pointing to dead PID | `antigravity-cleaner launch` (Auto-purges stale locks) |

---

## 3. CLI Non-Interactive Automation for Agents

In headless container, CI/CD, or automated subagent environments, execute commands non-interactively:

```bash
# 1. Run complete non-interactive system diagnosis
antigravity-cleaner doctor

# 2. Generate structured diagnostic report for telemetry/bug tickets
antigravity-cleaner report

# 3. Check for updates and self-upgrade
antigravity-cleaner update

# 4. Surgically patch all 4 targets (CLI, Electron main.js, VS Code extension, core binary)
antigravity-cleaner patch all

# 5. Clean temporary 429 quota cache without destroying user sessions
antigravity-cleaner clean

# 6. Generate standalone 1-click desktop launcher
antigravity-cleaner create-launcher

# 7. Restore original pristine binaries from automatic .agybak snapshots
antigravity-cleaner restore
```

---

## 4. Architectural Rules & Safety Constraints

1. **Zero Session Destruction**:
   - **DO NOT** execute crude commands like `rm -rf ~/.config/Antigravity` unless performing a total intentional factory reset.
   - `antigravity-cleaner clean` surgically targets only 16 specific quota bucket and GPU cache files while preserving `Cookies`, `Preferences`, and workspace history.
2. **Process-Level Routing vs TUN Mode**:
   - Never force root TUN VPN adapters for Antigravity. TUN mode causes packet fragmentation and triggers middlebox drops.
   - Use process injection: `--proxy-server=socks5://127.0.0.1:<port>`.
3. **Atomic Backup Integrity**:
   - Whenever `antigravity-cleaner patch` runs, it produces atomic `.agybak` backup files alongside modified assets, ensuring reversible patching.

---

## 5. File System Key Locations

- **Linux Config & Cache**: `~/.config/Antigravity/`
- **Linux Application Root**: `/usr/share/antigravity/` or `/opt/antigravity/`
- **macOS Application Support**: `~/Library/Application Support/Antigravity/`
- **Windows AppData**: `%APPDATA%\Antigravity\`
- **Local Proxy Loopbacks Tested**:
  - `127.0.0.1:10808` (v2ray / Xray SOCKS5)
  - `127.0.0.1:2080` (NekoBox / Hiddify SOCKS5)
  - `127.0.0.1:7890` (Clash / Mihomo Mixed)
  - `127.0.0.1:1080` (Standard SOCKS5)
