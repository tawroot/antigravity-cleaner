# Antigravity Cleaner Shell (v4.1.0)

<div align="center">
  <img src="docs/images/banner.png" alt="Antigravity Cleaner Banner" width="100%">
  <br>
  
  [![Version](https://img.shields.io/badge/Version-4.1.0-blue?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner/releases)
  [![License](https://img.shields.io/badge/License-TACL-red.svg?style=for-the-badge)](docs/LICENSE.md)
  [![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-blue.svg?style=for-the-badge&logo=platform.sh)](https://github.com/tawroot/antigravity-cleaner)
  [![Language](https://img.shields.io/badge/Language-PowerShell-yellow.svg?style=for-the-badge&logo=powershell)]()
  [![Security](https://img.shields.io/badge/Security-Zero%20Telemetry-green.svg?style=for-the-badge&logo=shields)](docs/SECURITY.md)
</div>

> *Dedicated to the people of Iran, Cuba, Syria, and all those navigating the complexities of digital sanctions. We believe access to technology is a fundamental right.*

---

## 📢 Disclaimer
**THIS SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND.**
By using this tool, you acknowledge that you are responsible for your own data. The Session Manager manipulates sensitive browser files. While tested extensively, we recommend keeping independent backups of critical data before performing restore operations.

---

## 🎯 Who Needs This?

This tool is specifically designed for:

### 🌍 **Developers in Sanctioned Regions**
If you're in **Iran, China, Russia, Cuba, Syria, North Korea, Turkmenistan, or Turkey**, you've likely encountered:
*   `HTTP Error 403: Forbidden` when accessing Google services
*   `ModuleNotFoundError` when installing Antigravity IDE
*   Region-locked errors preventing access to Gemini AI, Colab, or Cloud Platform
*   Great Firewall (China) or government censorship blocking developer tools

### 💻 **Programmers Facing Installation Errors**
Common errors this tool fixes:
*   `ERROR: Antigravity installation failed`
*   `Pip install error: Could not find a version that satisfies the requirement`
*   `Dependency conflict detected`
*   Corrupted browser sessions after system reinstall

### 🔧 **Anyone Needing Google Service Access**
*   Developers who need Gemini AI, Google Colab, or Cloud Platform
*   Users experiencing DNS/network interference
*   Those who need to preserve browser sessions across system changes

**If you've ever Googled "how to fix antigravity install error" or "bypass region lock" — this tool is for you.**

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
iwr https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 -useb | iex
```

<div align="center">
  <img src="docs/images/screen_collage.png" alt="Antigravity Shell Interface Collage" width="80%">
  <p><i>Command Center & Modules Overview</i></p>
</div>

### Option 2: Manual (Git)
```powershell
git clone https://github.com/tawroot/antigravity-cleaner.git
cd antigravity-cleaner
.\Antigravity.ps1
```

---

## 📖 Documentation & Verification
For a detailed step-by-step walkthrough of every module (Session Manager, Region Inspector, Network Tools), please check our Wiki-style guides:

👉 **[READ: Professional Usage Guide](docs/GUIDE.md)**
👉 **[SECURITY: Zero-Trust Policy](docs/SECURITY.md)**
👉 **[ARCHITECTURE: How it Works](docs/ARCHITECTURE.md)**

---

## 📈 Star History
We are growing fast! Thank you for your support.
<br>
<a href="https://star-history.com/#tawroot/antigravity-cleaner&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=tawroot/antigravity-cleaner&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=tawroot/antigravity-cleaner&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=tawroot/antigravity-cleaner&type=Date" />
 </picture>
</a>

---

## 💖 Support the Mission (Campaign)
**We are in a race against time.**
**We are in a race against time.**
Sanctions and restrictions are constantly evolving. Antigravity Cleaner v4.1.0 is now **cross-platform**, but maintaining it for Windows, macOS, and Linux requires dedicated effort.

**Your support is the fuel for this engine.**
If this tool saved you hours of debugging or fixed your region-locked account, please consider donating. It helps us stay awake, code faster, and fight for a free internet.

<div align="center">

| **Crypto** | **Wallet / Link** |
| :--- | :--- |
| **NOWPayments** | [👉 **Click to Donate (BTC/ETH/USDT)**](https://nowpayments.io/donation/tawroot) |
| **USDT (TRC20)** | `TN8GzU2X3x... (Select in Link)` |
| **Bitcoin** | `bc1q... (Select in Link)` |

</div>

> *"Independent development is the only way to guarantee zero censorship."*

---

## 🤝 Contributing
We welcome feature suggestions and bug reports!

**How to Contribute:**
1.  **Feature Requests:** Open a [Discussion](https://github.com/tawroot/antigravity-cleaner/discussions) on GitHub.
2.  **Bug Reports:** Open an [Issue](https://github.com/tawroot/antigravity-cleaner/issues) with detailed steps to reproduce.
3.  **Code Contributions:** While we appreciate the interest, please note that direct code contributions (PRs) are currently not accepted due to the proprietary nature of the license. However, your ideas and feedback are highly valued!

---

## 📞 Contact & Community
*   **Telegram Channel:** [t.me/panbehnet](https://t.me/panbehnet) - Updates, tips, and support.
*   **GitHub Issues:** [Report bugs or request features](https://github.com/tawroot/antigravity-cleaner/issues).
*   **GitHub Discussions:** [Join the conversation](https://github.com/tawroot/antigravity-cleaner/discussions).

---

## 👨‍💻 Author
Developed by **Tawana Network**.
*Building bridges where others build walls.*
