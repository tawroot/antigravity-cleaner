# 🔬 Forensic Troubleshooting Clinic & Diagnostics Guide

> Comprehensive diagnostic manual for identifying and resolving all known errors, network timeouts, region blocks, and proxy connection faults in **Google Antigravity IDE**, the `agy` CLI, and related services.

---

## 📑 Diagnostic Matrix

| Error Code / Symptom | Root Cause | 1-Click Solution |
| :--- | :--- | :--- |
| **`Agent terminated due to error` & `write EIO`** | Broken pipe exception in Electron logging / middlebox idle drops | `antigravity-cleaner patch all` (ASAR Shield + 15s KeepAlive) |
| **`Error Unknown: There was a network issue connecting to the server`** | DPI idle timeout on silent SSE/gRPC streaming inference | `antigravity-cleaner launch` (Direct SOCKS5 loopback) |
| **`403 Forbidden` (`/auth-success`)** | Regional IP restriction during OAuth callback | `antigravity-cleaner patch all` (`hasValidAuth = 1`) |
| **`proxyconnect tcp: dial tcp 127.0.0.1:10808: connection refused`** | Proxy client closed or port mismatch | `antigravity-cleaner doctor` (Auto-scans ports) |
| **`Failed to fetch` (Setting up account)** | Stale DIPS records & corrupted token cache | `antigravity-cleaner clean` (Surgical quota reset) |
| **`macOS Operation Not Permitted (syscall.EPERM)`** | macOS App Management permission boundary | Enable Terminal in **Privacy & Security ➔ App Management** |
| **`HTTP 429: Resource has been exhausted`** | Temporary rate-limit bucket corruption | `antigravity-cleaner clean` (Preserves 100% chats) |
| **`Your current account is not eligible for Antigravity`** | MultiGate client gatekeeper | `antigravity-cleaner patch all` |
| **`SingletonLock: process already running`** | Orphaned PID file after sudden shutdown | `antigravity-cleaner launch` (Auto-cleans stale locks) |

---

## 1. Agent Terminated Due to Error (`write EIO` / Broken Pipe)

![Agent Terminated Error](images/errors/google-antigravity-agent-terminated-due-to-error.png)

### Error Signature
```text
Agent terminated due to error
You can prompt the model to try again or start a new conversation if the error persists.
See our troubleshooting guide for more help.

Dismiss  Copy debug info  Retry
```

### Forensic Analysis
* **Root Cause:** Electron's internal logging library (`electron-log`) writes excessively verbose debug traces to standard error (`stderr`). When middleboxes or firewalls abruptly drop long-lived streaming connections during inference, writing to the closed file descriptor triggers an uncaught POSIX `EIO` (Input/output error), terminating the agent loop.
* **Resolution:**
  Execute:
  ```bash
  antigravity-cleaner patch all
  ```
  The **ASAR Shield** surgically patches `app.asar`, setting `level: false` on verbose loggers and wrapping stream write functions in robust exception guards. Additionally, a 15-second TCP KeepAlive probe is injected.

---

## 2. Network Issue Connecting to Server (`Unknown Error ID`)

![Network Issue Error](images/errors/google-antigravity-network-issue-connecting-to-server.png)

### Error Signature
```text
Worked for 1s >

Error Unknown: There was a network issue connecting to the server, please try again. ˅
Error ID: 95ea05dd-6419-4bf0-a8ce-8faf25687606-108
Unknown: There was a network issue connecting to the server, please try again.
```

### Forensic Analysis
* **Root Cause:** State-sponsored Deep Packet Inspection (DPI) middleboxes inject silent RST packets into idle gRPC streams. Furthermore, background telemetry calls originating from the IDE can leak raw non-proxy IPs midway through sessions.
* **Resolution:**
  Launch the IDE via process proxy injection:
  ```bash
  antigravity-cleaner launch
  ```
  This forces all network operations through a hardened local SOCKS5h loopback with zero telemetry exfiltration.

---

## 3. Authentication Gating & 403 Forbidden (`/auth-success`)

![403 Forbidden Auth](images/errors/google-antigravity-error-403-forbidden-auth-success.png)

### Error Signature
```text
Google

403. That’s an error.

Your client does not have permission to get URL /auth-success from this server. That’s all we know.
```

### Forensic Analysis
* **Root Cause:** Upstream Google Cloud edge filters detect regional origins during OAuth web redirection to `/auth-success`.
* **Resolution:**
  ```bash
  antigravity-cleaner patch all
  ```
  Applies machine-code patch to the core binary (`language_server`), forcing `hasValidAuth = 1`, and injects `--stubby_auth_mode=business_licensed_enterprise`.

---

## 4. Proxy Connection Refused (`dial tcp 127.0.0.1:10808: connection refused`)

![Proxy Connection Refused](images/errors/google-antigravity-loadcodeassist-proxy-connection-refused.png)

### Error Signature
```text
There was an unexpected issue setting up your account.

failed to get load code assist response: Post "https://cloudcode-pa.googleapis.com/v1internal:loadCodeAssist": proxyconnect tcp: dial tcp 127.0.0.1:10808: connect: connection refused
```

### Forensic Analysis
* **Root Cause:** Antigravity is configured to route via port `10808`, but the local proxy client (v2ray, Xray, Clash, NekoBox) is inactive or running on another port (such as `2080`, `7890`, or `1080`).
* **Resolution:**
  Run the automated diagnostic scanner:
  ```bash
  antigravity-cleaner doctor
  ```
  The doctor inspects all active listening ports on `127.0.0.1` and guides you to connect to the active listener.

---

## 5. Account Setup Failed (`Failed to fetch`)

![Failed to fetch Error](images/errors/google-antigravity-setting-up-account-failed-to-fetch.png)

### Error Signature
```text
There was an unexpected issue setting up your account.

Failed to fetch

Continue with different account
Having trouble? Let us know
```

### Forensic Analysis
* **Root Cause:** Corrupted DIPS database artifacts and orphaned WebSocket sessions in the Chromium storage layer.
* **Resolution:**
  Execute non-destructive quota cleaning:
  ```bash
  antigravity-cleaner clean
  ```
  Flushes 16 transient cache targets in under 2 milliseconds without deleting workspace files or chat history.

---

## 6. macOS App Management Permission (`syscall.EPERM`)

### Forensic Analysis
* **Root Cause:** macOS Ventura, Sonoma, and Sequoia restrict command-line processes from modifying signed application bundles in `/Applications/Antigravity.app`.
* **Resolution:**
  1. Open macOS **System Settings ➔ Privacy & Security ➔ App Management**.
  2. Toggle **ON** the switch for your terminal application (Terminal, iTerm2, Kitty, Ghostty, or VS Code).
  3. Re-run:
     ```bash
     ./antigravity-cleaner patch all
     ```

---

## 7. HTTP 429: Resource Exhausted (Quota Bucket Corruption)

### Forensic Analysis
* **Root Cause:** Local rate-limiting counters store stale quota tokens even after Google server-side quota windows reset.
* **Resolution:**
  ```bash
  antigravity-cleaner clean
  ```
  Purges `~/.config/Antigravity/Cache` rate-limit trackers while preserving `Cookies` and user configurations.
