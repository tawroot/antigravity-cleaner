---
name: antigravity-cleaner
description: Self-healing and diagnostic toolkit for Google Antigravity IDE and CLI. Use when experiencing region lock errors, HTTP 429 quota exhaustion, stream dropouts, or proxy connection issues.
---

# Antigravity Cleaner AI Skill

This skill equips AI coding agents with the capability to audit, self-heal, and unblock their own Google Antigravity operating environment.

## When to Use

Activate this skill when:
- The user reports "Region not supported", "Account eligibility check failed", or Gate 1 barriers.
- Antigravity experiences HTTP 429 Quota exhaustion or token corruption.
- API streaming gets interrupted or throttled by middlebox DPI firewalls.
- The environment requires routing through local proxies (Xray, Clash, NekoBox) without root TUN mode.

## Available Diagnostic & Repair Commands

Always prefer invoking the pre-compiled native Go binary `ag-cleaner` (or `antigravity-cleaner` / `agc`):

### 1. Health & Connection Audit
Run concurrent diagnostics to pinpoint connection or patch issues:
```bash
ag-cleaner doctor
```
For verbose network traces:
```bash
ag-cleaner doctor --verbose
```

### 2. Smart 1-Click Repair (Recommended)
Unlocks the core machine-code gates, configures local smart proxy injection, and creates the desktop launcher:
```bash
ag-cleaner patch all
```

### 3. Surgical 429 Quota Reset
Flushes corrupted token cookies, DIPS, and network state without wiping user conversations:
```bash
echo "y" | ag-cleaner clean
```

### 4. Smart Proxy Launch (No TUN)
Spawns Antigravity with auto-detected local proxy and TCP KeepAlive tuning:
```bash
ag-cleaner launch
```

## Security & Data Integrity Rules
- **Preserve Chats:** Never delete `User/globalStorage` or workspace chat records.
- **Atomic Modification:** Patches are applied atomically via POSIX Inode replacement without killing running processes.
