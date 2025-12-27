# Antigravity Cleaner Shell (v3.0.0)

[![License](https://img.shields.io/badge/License-TACL-red.svg)](docs/LICENSE.md)
[![Platform](https://img.shields.io/badge/Platform-Windows-blue)](https://microsoft.com/windows)
[![Security](https://img.shields.io/badge/Security-Zero%20Telemetry-green)](docs/SECURITY.md)

> *Dedicated to the people of Iran, Cuba, Syria, and all those navigating the complexities of digital sanctions. We believe access to technology is a fundamental right.*

---

## 📢 Disclaimer
**THIS SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND.**
By using this tool, you acknowledge that you are responsible for your own data. The Session Manager manipulates sensitive browser files. While tested extensively, we recommend keeping independent backups of critical data before performing restore operations.

---

## 🚀 Overview
**Antigravity Cleaner** is an enterprise-grade, PowerShell-based automation tool designed for developers, freelancers, and power users living under digital sanctions. It provides a robust suite of tools to manage browser identities, optimize system performance, and ensure connectivity to essential development services without relying on heavy external dependencies.

**Key Value Proposition:**
*   **Sanction Evasion:** Tools to inspect and manage Google "Country Association".
*   **Identity Management:** Seamlessly backup/restore browser sessions (Cookies/Logins) across system resets.
*   **Zero Dependencies:** Runs natively on Windows. No Python, No Node.js required.
*   **Privacy First:** 100% Offline execution. No telemetry.

---

## 📚 Documentation
For detailed technical and security information, please refer to our Wiki:

*   🔒 **[Security Policy & Privacy](docs/SECURITY.md)** - How we handle your data (Zero-Trust) and encryption details.
*   🏗️ **[Technical Architecture](docs/ARCHITECTURE.md)** - Deep dive into the PowerShell engine, profile parsing, and file locking logic.
*   📜 **[License Agreement](docs/LICENSE.md)** - Review the Tawana Anti-Copy License (TACL).
*   📘 **[User Guide (Persian)](README.fa.md)** - راهنمای کامل فارسی.
*   📕 **[User Guide (Turkish)](README.tr.md)** - Türkçe Tam Kılavuz.

---

## 🔥 Features at a Glance

### 🛡️ Session Manager (v3.0)
Stop fearing Windows re-installs. Keep your logins alive.
*   **Smart Detection:** Auto-maps `Profile Folders` to `Email Addresses` for Chrome, Edge, Brave, and Opera.
*   **Dual-Mode Backup:**
    *   **Light Mode:** Backs up ONLY Logins/Cookies (~20MB). Perfect for quick portability.
    *   **Full Mode:** Clones the entire browser identity (~500MB+).
*   **Antigravity Desktop:** Specialized backup for the Antigravity IDE (VS Code based).

### 🌎 Region Inspector
*   **Pre-Flight Checks:** Automated analysis of IP, DNS, and WebRTC leaks before you interact with Google.
*   **Direct Access:** Deep-links into Google's Country Association settings using specific browser profiles.

### ⚡ System Analysis & Optimization
*   **Connectivity Probe:** Real-time health check for Google Services (Gemini, Cloud, Colab) and Dev Repositories.
*   **Network Reset:** One-click repair for Winsock/DNS/TCP stack issues.
*   **Cleaner:** Intelligent removal of development cache (VS Code, JetBrains) and system temp files.

---

## 💾 Installation

### Option 1: One-Line Installer (Recommended)
Open PowerShell and paste this command. It will download the latest core and create a shortcut on your Desktop.

```powershell
iwr https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/antigravity-cleaner/install.ps1 -useb | iex
```

### Option 2: Manual (Git)
```powershell
git clone https://github.com/tawroot/antigravity-cleaner.git
cd antigravity-cleaner
.\Antigravity.ps1
```

---

## ⚖️ License & Terms
**Copyright (c) 2025 Tawana Network. All Rights Reserved.**

This project is released under the **Tawana Anti-Copy License (TACL)**.
*   ✅ You may view, modify, and use this code for personal/internal purposes.
*   ❌ You may **NOT** copy, distribute, sell, or sublicense this software.
*   ❌ You may **NOT** remove copyright headers or present this as your own work.

For detailed terms, see [LICENSE](docs/LICENSE.md).

---

## 👨‍💻 Author
Developed by **Tawana Network**.
*Building bridges where others build walls.*
