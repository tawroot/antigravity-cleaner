# 📖 Antigravity Cleaner — Comprehensive User Guide (v5.2.0)

A complete manual for troubleshooting, bypassing regional restrictions, fixing quota locks, and configuring proxies for **Google Antigravity IDE**, `agy` CLI, and Google Antigravity extension.

---

## 📑 Table of Contents
1. [Core Problems Solved](#-core-problems-solved)
2. [Interface Options](#-interface-options)
   - [1-Click Retro KeyGen GUI (Zero-Terminal)](#1-1-click-retro-keygen-gui)
   - [Modern Interactive Terminal (TUI)](#2-modern-interactive-terminal-ui-tui)
   - [Headless CLI Subcommands](#3-headless-cli-subcommands)
3. [CLI Command Reference](#-cli-command-reference)
   - [`doctor` — Health & Diagnostic Probe](#ag-cleaner-doctor)
   - [`clean` — Surgical Quota & Cache Reset](#ag-cleaner-clean)
   - [`patch` — Regional Restriction Neutralizer](#ag-cleaner-patch)
   - [`proxy` — No-TUN Smart Socket Injector](#ag-cleaner-proxy)
   - [`run` — Wrap & Launch IDE](#ag-cleaner-run)
4. [Step-by-Step Troubleshooting Workflows](#-step-by-step-troubleshooting-workflows)
   - [Fixing "User location is not supported" (403 Forbidden)](#workflow-1-bypassing-user-location-is-not-supported-403-forbidden)
   - [Fixing Stuck Rate-Limits & Quota Locks (429 Quota Exhausted)](#workflow-2-fixing-stuck-rate-limits--quota-locks-429)
   - [Fixing Streaming Freezes & Drops Mid-Generation](#workflow-3-fixing-streaming-freezes--drops-mid-generation)
   - [Connecting Without Root TUN Mode (Clash / Xray / Sing-box / Hiddify)](#workflow-4-connecting-via-local-proxy-without-tun-mode)
5. [Backups & Safety Rollback](#-backups--safety-rollback)

---

## 🎯 Core Problems Solved

| Problem / Error | Root Cause | Antigravity Cleaner Solution |
| :--- | :--- | :--- |
| **`User location is not supported` (HTTP 403)** | Google geo-fencing checks country eligibility inside `language_server` bytecode. | **`ag-cleaner patch`**: Scans bytecode at 1.9 GB/s and flips eligibility checks. |
| **`Quota exceeded / Rate limited` (HTTP 429)** | Corrupted local quota counters and orphaned lock files trick the IDE into believing limits are exceeded. | **`ag-cleaner clean`**: Surgically purges stale quota state without touching code or settings. |
| **Stream Freezes Mid-Generation** | Intermediate firewalls drop idle TLS connections during heavy LLM generation. | **`ag-cleaner proxy`**: Injects active TCP KeepAlive heartbeats into outbound connections. |
| **Proxy Works in Browser but Fails in IDE** | Antigravity runs as a decoupled daemon that ignores system-level proxy settings unless TUN is used. | **`ag-cleaner proxy`**: Injects direct SOCKS5/HTTP loopback routing (`127.0.0.1:10808` / `7890`). |

---

## 🖥️ Interface Options

### 1. 1-Click Retro KeyGen GUI
* **Target Audience:** Users who want instant fixes without touching terminal commands.
* **How to run:** Double-click the downloaded executable (`.exe` on Windows, native binary on macOS/Linux).
* **Features:** Windows 95 classic design, real-time demoscene starfield, Web Audio retro chiptune SFX, one-click **"Auto-Fix & Patch"** button.

### 2. Modern Interactive Terminal UI (TUI)
* **Target Audience:** Developers who prefer an in-terminal dashboard.
* **How to run:**
  ```bash
  ag-cleaner tui
  # or
  ./antigravity-cleaner tui
  ```
* **Controls:** `Tab` / `Shift+Tab` to switch panels, `Enter` to run actions, `q` or `Ctrl+C` to quit.

### 3. Headless CLI Subcommands
* **Target Audience:** Power users, CI/CD scripts, and automated developer environments.

---

## 🛠️ CLI Command Reference

### `ag-cleaner doctor`
Performs an exhaustive health check of your Antigravity environment.
```bash
ag-cleaner doctor
```
**Checks Performed:**
* Detects running instances of Antigravity IDE and `language_server`.
* Verifies socket accessibility on internal gRPC ports.
* Probes latency to Google Cloud Code endpoints (`cloudcode.googleapis.com`).
* Inspects system proxy environment variables (`HTTP_PROXY`, `ALL_PROXY`).
* Checks binary patch status (Patched vs. Original).

---

### `ag-cleaner clean`
Performs a surgical reset of volatile caches and stale lock files.
```bash
ag-cleaner clean
```
**What it Cleans:**
* Stale process locks (`.lock`, `.pid`) preventing `language_server` from spawning.
* Corrupted local rate-limit and quota token caches (HTTP 429).
* Dangling IPC sockets.

> [!IMPORTANT]
> `ag-cleaner clean` **NEVER** deletes your project source code, local git repositories, custom keybindings, or user settings. It is strictly non-destructive.

---

### `ag-cleaner patch`
Applies an in-place binary bytecode patch to `language_server` to bypass regional sanctions.
```bash
ag-cleaner patch
```
* Locates the active `language_server` binary across standard installation directories.
* Creates an automatic timestamped backup (`language_server.bak.<timestamp>`).
* Employs Boyer-Moore pattern matching (1.9 GB/s scan rate) to locate the country verification routine.
* Safely neutralizes the geo-fence check.

To revert a patch:
```bash
ag-cleaner patch --restore
```

---

### `ag-cleaner proxy`
Configures and verifies smart proxy routing without requiring root VPN / TUN mode.
```bash
# Auto-detect local proxy (probes 10808, 7890, 1080, 2080)
ag-cleaner proxy

# Specify custom SOCKS5 proxy port
ag-cleaner proxy --socks5 127.0.0.1:10808

# Specify custom HTTP proxy port
ag-cleaner proxy --http 127.0.0.1:7890
```

---

### `ag-cleaner run`
Launches Antigravity IDE with verified proxy environment variables injected into the child process environment.
```bash
ag-cleaner run
```

---

## 🚀 Step-by-Step Troubleshooting Workflows

### Workflow 1: Bypassing "User location is not supported" (403 Forbidden)
1. Close Antigravity IDE completely.
2. Run the patcher:
   ```bash
   ag-cleaner patch
   ```
3. Run diagnostic check:
   ```bash
   ag-cleaner doctor
   ```
4. Start your local VPN or proxy client (e.g., Clash, v2rayA, NekoBox).
5. Launch Antigravity IDE. Regional restriction errors will be gone.

---

### Workflow 2: Fixing Stuck Rate-Limits & Quota Locks (429)
1. Close Antigravity IDE.
2. Run surgical cache clean:
   ```bash
   ag-cleaner clean
   ```
3. Re-open Antigravity IDE. Your AI code completions and assistant queries will resume immediately.

---

### Workflow 3: Fixing Streaming Freezes & Drops Mid-Generation
1. Ensure your proxy client has stable nodes with UDP/TCP support.
2. Launch Antigravity via the proxy wrapper with active TCP KeepAlive:
   ```bash
   ag-cleaner run
   ```
3. Long streaming responses will remain alive without premature connection resets.

---

### Workflow 4: Connecting via Local Proxy Without TUN Mode
If you don't want or cannot enable system-wide TUN (VPN) mode:
1. Open your proxy client (Clash on `7890`, v2rayN / Xray on `10808`, Hiddify on `2334`, etc.).
2. Run:
   ```bash
   ag-cleaner proxy --socks5 127.0.0.1:10808
   ```
3. Antigravity IDE will connect directly through the local loopback port with zero system routing modifications.

---

## 🛡️ Backups & Safety Rollback

Every time `ag-cleaner patch` is executed, a full pristine copy of the target binary is preserved:
- **Location:** Identical directory as the target executable.
- **Naming Pattern:** `<original_binary>.bak.<timestamp>`
- **Restoration Command:**
  ```bash
  ag-cleaner patch --restore
  ```

---

## 📜 License
Antigravity Cleaner is released under the **Apache License 2.0**.
