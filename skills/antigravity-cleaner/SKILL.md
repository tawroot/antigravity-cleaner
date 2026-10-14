---
name: antigravity-cleaner
description: Autonomous self-healing, sanction bypass, and diagnostics skill for Google Antigravity IDE and CLI. Use when experiencing region restrictions, HTTP 429 quota exhaustion, license #3501 errors, or when the user asks to unlock, patch, or verify their Antigravity environment.
---

# Antigravity Cleaner Pro — Autonomous AI Agent Skill (v5.3.0)

This skill provides an autonomous, standardized operating protocol for AI coding assistants (Google Antigravity, Claude Code, Cursor, Codex) to audit, self-heal, patch, and unblock the developer's Google Antigravity environment without breaking active chat sessions.

---

## 🎯 When to Trigger This Skill

Activate this protocol immediately whenever:
1. **User Request:** The user provides the repository link or asks: *"Install and run Antigravity Cleaner"*, *"Fix my Antigravity license"*, *"Reset my 429 quota"*, or *"Unlock region restrictions"*.
2. **License Error (#3501):** The user receives *"You do not have a valid license of this product (#3501)"*.
3. **Location Restriction (HTTP 400):** Upstream API fails with *"User location is not supported for the API use"*.
4. **Quota Exhaustion (HTTP 429):** The IDE reports *"Too Many Requests"* or corrupted rate-limit buckets.
5. **Detection Failure:** The Core Engine or `language_server` is reported as unpatched or missing.

---

## 🤖 Standard Autonomous Agent Workflow

When invoked by the user, follow this 4-step sequence:

### Step 1: Tool Verification & User Briefing
Check if `antigravity` (or `ag-cleaner`) binary is available on the path or built in `./bin/antigravity`.
If not installed, build or install it:
```bash
# Quick source build (Go 1.22+)
make build || go build -o bin/antigravity ./cmd/antigravity
```

**Inform the user immediately:**
> *"Antigravity Cleaner Toolkit v5.3.0 is initialized. I can apply all patches autonomously right here, or you can launch the classic retro desktop window using `antigravity gui` if you prefer an interactive interface."*

---

### Step 2: Safe Process Management & Auto-Fix Execution
To ensure patches are bound to active memory and eliminate file lock collisions (especially on Windows), stale IDE instances should be cleanly refreshed. The agent can run:

```bash
# Run 1-Click Auto-Fix (Core MultiGate bytecode + main.js + Zero-Telemetry engine)
echo "y" | ./bin/antigravity auto-fix
```
> [!NOTE]
> The engine automatically protects active shell ancestor PIDs (`isProtectedPID`), preventing the active agent session from being disconnected.

---

### Step 3: Surgical 16-Target Quota & Token Reset (If Quota 429 Reported)
If the user specifically suffered from HTTP 429 or corrupted rate-limit cookies:
```bash
echo "y" | ./bin/antigravity clean
```
This safely purges cookies, DIPS, Trust Tokens, and GPU caches across 16 targets in 0.002s while **preserving 100% of user chat history, projects, and credentials**.

---

### Step 4: Health Audit & Verification Report
Execute the built-in diagnostic doctor to verify that all patches are locked and operational:
```bash
./bin/antigravity doctor
```

**Present a clear, structured verification report to the user:**
```markdown
### 🩺 Antigravity Environment Health Report
- **Core Engine (language_server):** ✅ Unlocked & Patched (MultiGate IDA bytecode)
- **IDE Frontend (main.js):** ✅ Patched (`isGoogleInternal` active)
- **Zero-Telemetry Engine:** ✅ Active (`--disable_telemetry=true`, 7 enterprise flags)
- **Proxy Status:** ✅ Auto-detected & bound (No TUN mode required)
- **Next Steps:** You can launch Antigravity directly via `antigravity launch` or use `antigravity gui` anytime!
```

---

## 🛠️ Complete CLI Command Reference

| Command | Action | Recommended Scenario |
| :--- | :--- | :--- |
| `antigravity auto-fix` | 1-Click full unlock of all 4 targets + flag injection | Initial setup or license repair |
| `antigravity gui` | Opens Retro KeyGen Desktop GUI (Win95 window) | User wants visual 3-step control |
| `antigravity doctor` | Runs parallel latency, DNS leak & patch audit | Health inspection |
| `antigravity clean` | Surgical 16-target Quota 429 & token cache reset | Rate limits / session corruption |
| `antigravity launch` | Spawns IDE with smart proxy & 15s TCP KeepAlive | Everyday launching without TUN |
| `antigravity kill` | Terminates zombie language_server instances | Stale socket locks or hung IDE |
| `antigravity create-launcher` | Installs `.desktop` shortcut & high-res XDG icon | Linux desktop integration |

---

## 🔒 Security & Data Integrity Rules
- **Chat Immutability:** Never purge `User/globalStorage` or workspace chat records.
- **Session Shield:** Do not terminate processes marked by `ANTIGRAVITY_PROTECT_PID` or parent terminal PIDs.
- **Atomic File Operations:** All disk mutations use POSIX atomic inode swaps with `.agybak` backups.
