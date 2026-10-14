# 🪲 Upstream Bug Registry & Knowledge Base

This directory maintains an engineering knowledge base of upstream bugs, race conditions, unhandled exceptions, and architectural defects discovered in **Google Antigravity**, **Electron runtime**, and **underlying Node.js libraries** during reverse engineering and system hardening.

Each record documents:
- **Exact Stack Trace & Components Affected**
- **Upstream Root Cause & Runtime Triggers**
- **Security & Privacy Implications**
- **Antigravity Cleaner Pro Neutralization Strategy**

---

## 📋 Registry Index

| Bug ID | Title & Component | Severity | Mitigation Status in Toolkit |
| :--- | :--- | :---: | :--- |
| [**UPSTREAM-001**](UPSTREAM-001-electron-log-write-eio.md) | `Uncaught Exception: Error: write EIO` in `electron-log` Console Transport | Low (UX / Noise) | 🟡 **In Progress** (Launcher workaround applied; ASAR root patch planned in v5.4.0) |

---

## 📝 Contribution & Registration Standards

When logging an upstream bug:
1. Provide the exact, unadulterated call stack from the Electron error dialog or system journal.
2. Differentiate between upstream application crashes vs harmless unhandled console logging events.
3. Detail how Antigravity Cleaner Pro isolates or neutralizes the bug without requiring root privileges or destabilizing user sessions.
