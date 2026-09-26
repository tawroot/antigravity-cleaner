# Security Policy

## Overview & Philosophy

**Antigravity Cleaner** is an open-source diagnostics and environment-restoration toolkit built for developers navigating digital sanctions and network edge restrictions. Because Antigravity executes low-level diagnostics and bytecode modifications on local binaries, security, transparency, and data integrity are fundamental design requirements.

---

## Supported Versions

Only the latest release versions receive active security updates and patch updates.

| Version | Supported | Status |
| :--- | :--- | :--- |
| `v5.2.x` | :white_check_mark: | Current Stable Release |
| `v5.1.x` | :white_check_mark: | Maintenance Only |
| `< v5.0.0` | :x: | End of Life (Deprecated) |

---

## Core Security & Privacy Commitments

When using `antigravity-cleaner`, the following security boundaries are strictly enforced:

### 1. Zero Telemetry & Complete Privacy
- **No Remote Phone-Home:** The toolkit does not collect, log, or transmit personal data, code, or workspace paths to any external server.
- **Account & Session Integrity:** Surgical cleanup functions (`ag-cleaner clean`) strictly preserve user login credentials, Google Auth tokens, and active workspace configurations.
- **Chat History Protection:** User conversations and session trajectories located in `User/globalStorage` and `.gemini` are immutable to the cleaner and are never purged.

### 2. POSIX Atomic Inode Operations
- All file system modifications (such as unlocking `language_server` or updating configuration settings) are performed **atomically**. 
- A `.agybak` backup file is automatically generated before any binary modification, ensuring that reverting to the pristine official binary is always a one-command operation (`ag-cleaner restore` or option `8` in menu).

### 3. Transparent Process-Level Proxy Routing (No TUN Mode)
- `antigravity-cleaner` does not install kernel network drivers, virtual TUN adapters, or system-wide root network hooks.
- Proxy routing is strictly isolated to the Antigravity child process using process-level environment variables (`ALL_PROXY`, `HTTPS_PROXY`) and `--proxy-server` CLI arguments.

---

## Reporting a Security Vulnerability

We welcome responsible security research and disclosure. If you identify a security vulnerability, unintended token exposure, or unsafe file operation:

1. **Do NOT open a public GitHub issue.** Public issues should be reserved for functional bugs and feature requests.
2. **Submit via GitHub Private Vulnerability Reporting:**
   - Navigate to the repository's [Security Tab](https://github.com/tawroot/antigravity-cleaner/security).
   - Click **"Report a vulnerability"** to open an encrypted private draft advisory.
3. **Alternative Direct Disclosure:**
   - Contact the lead maintainer directly via GitHub ([@tawroot](https://github.com/tawroot)).
   - Include a detailed technical proof of concept (PoC), steps to reproduce, and impact assessment.

### Response Timelines
- **Initial Response:** Within 24 hours of receipt.
- **Triage & Status Assessment:** Within 48 hours.
- **Fix & Advisory Release:** Critical security patches will be published with an expedited patch release within 3 to 5 business days.

---

## Audit & Verification

All binaries released under GitHub Releases are built from the public repository using automated GitHub Actions CI/CD workflows with verifiable SHA256 checksums. Users are encouraged to inspect and compile from source using Go 1.22+:

```bash
git clone https://github.com/tawroot/antigravity-cleaner.git
cd antigravity-cleaner
go build -o ag-cleaner ./cmd/antigravity
```
