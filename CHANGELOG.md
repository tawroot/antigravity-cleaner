# 📜 Changelog - Antigravity Cleaner Pro

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [5.2.0] - 2026-09-13

### 💾 Retro KeyGen Desktop GUI & GitHub Actions Release Pipeline

- **Cross-Platform Retro KeyGen Desktop GUI:** Authentic Windows 95/98 warez-style graphical patcher with beveled 3D controls, raster demoscene starfield, Web Audio 8-bit sound effects, and 1-click execution.
- **Zero-Terminal Standalone Mode:** Automatically launches the GUI when double-clicked from file managers (Explorer / Finder / Desktop) without requiring any terminal interaction.
- **Full 4-Target Auto-Fix:** One-click unlocking covers Core `language_server` MultiGate bytecode, `agy` CLI binary, IDE `main.js`, and VS Code Extension guard, combined with Zero-TUN smart proxy injection.
- **Process Termination Shield:** Rewrote process termination logic to strictly exclude the cleaner process, GUI window, and helper processes, preventing self-termination during Auto-Fix.
- **Automated GitHub Actions Releases (`release.yml`):** Fully automated CI/CD pipeline compiling standalone binaries for Linux (amd64/arm64), macOS (Intel/Apple Silicon), and Windows with automated SHA-256 checksums.

---

## [5.1.1] - 2026-09-11

### 🛡️ Safety Guard: Session Disconnect Prevention

- **Active Process Detection:** Added real-time process monitoring (`IsAntigravityRunning()`) before executing `patch` or `clean` operations.
- **Session Protection Warning:** Prompts user with a prominent warning if Antigravity IDE or `language_server` is running, preventing accidental disconnection of active AI sessions.
- **Graceful Termination / Cancellation:** Allows users to safely cancel patching/cleaning to preserve their ongoing chat session, or close processes safely first if intended.

---

## [5.1.0] - 2026-09-11

### ⚡ Massive Core Architecture Rewrite in Pure Go (Zero Dependencies)

This milestone transforms **Antigravity Cleaner** into a blazing-fast, compiled Go system utility with a 2026-standard terminal interface (**Lipgloss & Bubbletea**), eliminating any need for Python, PowerShell runtimes, or system-wide root TUN modes.

- **Pure Go Architecture:** Single compiled binary with zero external dependencies for Linux (amd64, arm64), macOS (Intel, Apple Silicon M1-M4), and Windows.
- **Smart No-TUN Proxy Injector:** Automatically discovers local proxy ports (`10808`, `2080`, `7890`) and launches Antigravity with injected proxy environment and flags (`ALL_PROXY`, `HTTP_PROXY`, `--proxy-server`).
- **Aggressive TCP KeepAlive Tuning:** Keeps socket connections persistent every 15s to eliminate deep-packet inspection (DPI) stream resets.
- **Surgical 429 Quota Reset:** Purges corrupt tokens, HTTP cookies, and transport security caches without touching user chats or workspace data.
- **IDA-Style Byte Pattern Scanner:** In-memory bytecode scanning (<0.22s on 170MB binaries) with wildcards for MultiGate binary patching on x86-64 and ARM64.
- **Concurrent Diagnostic Doctor (`ag-cleaner doctor`):** Parallelized latency and DNS leak tests completing in under 3 seconds.
- **POSIX Atomic Inode Replacement:** Modifies binaries safely on disk without `ETXTBSY` or corruption if the application is running.
- **Maintained & Architected by:** **[@dalroot](https://github.com/dalroot)** on [tawroot/antigravity-cleaner](https://github.com/tawroot/antigravity-cleaner).

---

## [4.0.0] - 2025-12-26

### 🎉 Major Release - Complete Redesign & Feature Overhaul

This is a **major milestone** release with a complete UI/UX redesign, new features, and significant improvements across all platforms.

---

### ✨ New Features

#### 🎨 **Premium Apple-Style User Interface**
- **Complete UI Redesign** with macOS Big Sur / Apple Music-inspired aesthetics
  - Glassmorphism effects with frosted glass backgrounds
  - Smooth animations and transitions throughout the app
  - SF Pro-inspired typography (Inter, Outfit fonts)
  - Responsive grid layouts that adapt to window size
  - Beautiful gradient accents and hover effects
  
- **Dark/Light Mode Support**
  - Seamless theme switching with smooth transitions
  - System-aware theme detection
  - Persistent theme preferences
  - Optimized color schemes for both modes

- **Bilingual Interface (EN/FA)**
  - Full English and Persian (فارسی) language support
  - RTL (Right-to-Left) layout support for Persian
  - Language switcher in header
  - Persistent language preferences
  - All UI elements fully translated

#### 🌐 **Google Services Diagnostics**
- **Real-time Connectivity Testing** for Google services
  - Google Account access verification
  - Gemini AI service connectivity
  - Google Cloud Platform status
  - AI Studio accessibility check
  
- **Detailed Diagnostics Window**
  - Latency measurements for each service
  - Visual status indicators (✅ Success, ❌ Failed, ⏳ Testing)
  - Connection speed analysis
  - Troubleshooting recommendations
  - Export diagnostic reports

#### 💾 **Enhanced Session Manager**
- **Multi-Browser Support** (8+ browsers)
  - Google Chrome
  - Microsoft Edge
  - Mozilla Firefox
  - Brave Browser
  - Opera / Opera GX
  - Vivaldi
  - Arc Browser
  - Safari (macOS)

- **Advanced Session Management**
  - **Email-based Profile Detection** - Automatically extract and display email addresses from browser profiles
  - **Search by Email** - Find browser sessions by associated email address
  - **Encrypted Backups** - AES-256-GCM encryption for maximum security
  - **Compressed Archives** - ZIP compression for efficient storage
  - **Metadata Tracking** - Timestamp, browser type, profile name
  - **Quick Restore** - One-click session restoration
  - **Backup Verification** - Integrity checks for backup files

#### 🔧 **Network Optimization Tools**
- **DNS Flush** - Clear DNS cache with one click
- **Network Reset** - Complete Windows network stack reset
- **Network Diagnostics** - Detailed network configuration analysis
- **Connectivity Tests** - Ping tests to major services
- **Integrated Google Test** - Quick access from Network page

#### 🧹 **Enhanced System Cleaning**
- **Antigravity IDE Cleanup**
  - Remove corrupted installation files
  - Clear application caches
  - Clean temporary directories
  - Registry cleanup (Windows)
  - Process termination for locked files

- **Browser Cache Cleaning**
  - Clear browser caches without losing sessions
  - Selective cleaning by browser
  - Safe cleanup with session preservation

- **System Optimization**
  - Temp file removal
  - Log file cleanup
  - Disk space recovery
  - Performance improvements

#### 📊 **Dashboard & Health Score**
- **Real-time System Health Monitoring**
  - Overall health score (0-100)
  - Disk space analysis
  - Memory usage tracking
  - Process monitoring
  
- **Quick Actions Panel**
  - One-click access to common tasks
  - Status indicators for each module
  - Recent activity log

#### 🌍 **Multi-Language Documentation**
- Full README available in 8 languages:
  - 🇺🇸 English (README.md)
  - 🇮🇷 فارسی (README.fa.md)
  - 🇨🇳 中文 (README.zh.md)
  - 🇷🇺 Русский (README.ru.md)
  - 🇹🇷 Türkçe (README.tr.md)
  - 🇸🇦 العربية (README.ar.md)
  - 🇪🇸 Español (README.es.md)
  - 🇵🇰 اردو (README.ur.md)

#### 🌐 **Professional Website**
- **Apple Music-Inspired Design**
  - Premium dark theme with vibrant accents
  - Glassmorphism UI components
  - Smooth scroll animations
  - Responsive mobile design

- **Multi-Page Structure**
  - Homepage with Hero section
  - Features showcase
  - Download page with platform detection
  - Pricing & Premium services
  - Blog & News
  - FAQ & Documentation
  - Regional guides

- **SEO Optimization**
  - Meta tags for all pages
  - Structured data (Schema.org)
  - Multi-language support
  - Dynamic OG images
  - Sitemap generation

- **Community Integration**
  - Telegram channel widget
  - GitHub stats display
  - Funding campaign section
  - Newsletter signup

---

### 🔧 Improvements

#### **Performance**
- ⚡ **Faster Startup** - Optimized initialization sequence
- ⚡ **Reduced Memory Usage** - Efficient resource management
- ⚡ **Smoother Animations** - Hardware-accelerated rendering
- ⚡ **Faster Scans** - Parallel processing for browser detection

#### **Reliability**
- 🛡️ **Error Handling** - Comprehensive exception handling
- 🛡️ **Crash Recovery** - Auto-save and recovery mechanisms
- 🛡️ **Data Validation** - Input validation for all user actions
- 🛡️ **Backup Integrity** - Checksum verification for backups

#### **Usability**
- 🎯 **Intuitive Navigation** - Clear menu structure
- 🎯 **Contextual Help** - Tooltips and help text throughout
- 🎯 **Keyboard Shortcuts** - Quick access to common actions
- 🎯 **Progress Indicators** - Visual feedback for long operations

#### **Portability**
- 📦 **True Portable Mode** - Data folder created next to executable
- 📦 **No Registry Writes** - All settings stored in local files
- 📦 **USB-Friendly** - Run from any location without installation
- 📦 **Self-Contained** - All dependencies bundled

---

### 🐛 Bug Fixes

#### **UI/UX Fixes**
- ✅ Fixed UI collapse issues when resizing window
- ✅ Resolved missing color attributes in AppleColors class
- ✅ Fixed text overflow in small windows
- ✅ Corrected RTL layout issues in Persian mode
- ✅ Fixed animation glitches on slower systems
- ✅ Resolved theme switching bugs

#### **Functionality Fixes**
- ✅ Fixed portable path detection and data folder creation
- ✅ Corrected startup crashes on restricted systems
- ✅ Fixed browser profile detection on non-standard installations
- ✅ Resolved session backup failures for locked files
- ✅ Fixed DNS flush errors on Windows 11
- ✅ Corrected network reset permission issues

#### **Platform-Specific Fixes**
- **Windows:**
  - ✅ Fixed SmartScreen warnings with proper manifest
  - ✅ Resolved UAC elevation prompts
  - ✅ Fixed registry access errors on restricted accounts
  
- **macOS:**
  - ✅ Fixed Gatekeeper warnings
  - ✅ Resolved permission issues with browser directories
  - ✅ Fixed Safari session detection
  
- **Linux:**
  - ✅ Fixed missing tkinter dependencies
  - ✅ Resolved file permission issues
  - ✅ Fixed desktop integration

---

### 📦 Build & Release

#### **Automated Multi-Platform Builds**
- 🤖 **GitHub Actions Workflow** for automated builds
  - Windows (x64) - Portable EXE
  - macOS (Universal) - Portable Binary
  - Linux (x64) - Portable Binary

- 🤖 **Automated Packaging**
  - ZIP archives for Windows/macOS
  - TAR.GZ archives for Linux
  - Includes README, LICENSE, and data folder
  - Proper version information embedded

- 🤖 **Release Automation**
  - Automatic GitHub Releases on tag push
  - Professional release notes
  - Multi-platform asset uploads
  - Changelog integration

#### **Build Optimizations**
- 🔧 **PyInstaller Configuration**
  - `--noupx` flag to reduce antivirus false positives
  - Proper manifest for Windows
  - Version information embedding
  - Icon integration
  - Hidden imports for all dependencies

- 🔧 **Package Structure**
  ```
  AntigravityCleaner-Portable/
  ├── AntigravityCleaner.exe (or binary)
  ├── README.md
  ├── LICENSE
  ├── PORTABLE.txt
  └── data/
  ```

---

### 🔒 Security

- 🔐 **AES-256-GCM Encryption** for session backups
- 🔐 **Secure Password Handling** - Never stored in plain text
- 🔐 **Data Privacy** - All data stored locally, no telemetry
- 🔐 **Code Signing Ready** - Prepared for future code signing
- 🔐 **Permission Checks** - Proper UAC/sudo handling

---

### 📚 Documentation

- 📖 **Comprehensive README** - Installation, usage, troubleshooting
- 📖 **BUILDING.md** - Build instructions for all platforms
- 📖 **RELEASE_GUIDE.md** - Release process documentation
- 📖 **CHANGELOG.md** - This file, detailed change history
- 📖 **GITHUB_TOPICS.md** - SEO and discoverability optimization
- 📖 **Regional Guides** - Platform-specific troubleshooting

---

### 🌟 Branding & Marketing

- 🎨 **Tawana Network Branding** throughout the application
- 🎨 **Professional Logo** and icon design
- 🎨 **Consistent Visual Identity** across all platforms
- 🎨 **Premium Services Integration**
  - Panbeh VPN promotion
  - Antigravity Pro Account services
  - Business Account offerings
  - Region change services

---

### 🎯 Platform Support

| Platform | Version | Architecture | Status |
|----------|---------|--------------|--------|
| **Windows** | 10/11 | x64 | ✅ Fully Supported |
| **macOS** | 10.15+ | Universal | ✅ Fully Supported |
| **Linux** | Modern distros | x64 | ✅ Fully Supported |

---

### 📊 Statistics

- **Total Lines of Code:** ~15,000+
- **Supported Browsers:** 8+
- **Supported Languages:** 8
- **GitHub Stars:** Growing 🌟
- **Downloads:** 10,000+
- **Community:** Active Telegram channel

---

### 🙏 Credits

**Developed by:** Tawana Mohammadi / Tawana Network  
**Contributors:** Community feedback and testing  
**Special Thanks:** All users who reported issues and suggested features

---

### 📞 Support & Community

- 📢 **Telegram Channel:** [t.me/panbehnet](https://t.me/panbehnet)
- 🐛 **Report Issues:** [GitHub Issues](https://github.com/tawroot/antigravity-cleaner/issues)
- 🌐 **Website:** [tawroot.github.io/antigravity-cleaner](https://tawroot.github.io/antigravity-cleaner/)
- 💬 **Contact:** [@RAHBARUSD](https://t.me/RAHBARUSD)

---

## [2.1.0] - 2024-12-15

### Initial Public Release

#### Features
- Basic GUI with cleaning functionality
- Windows Registry scanning and cleanup
- Process detection and termination
- Network reset capabilities
- Multi-language support (EN/FA)
- Browser session detection (Chrome, Edge)

#### Known Issues
- Limited browser support
- Basic UI design
- Windows-only registry features
- No encryption for backups

---

## [1.0.0] - 2024-12-01

### Internal Beta Release

- Initial development version
- Core cleaning functionality
- Basic browser detection
- Command-line interface
- Windows-only support

---

**Powered by TAWANA NETWORK**  
**Copyright © 2024-2025 Tawana Mohammadi. All Rights Reserved.**

---

## Legend

- ✨ New Feature
- 🔧 Improvement
- 🐛 Bug Fix
- 🔒 Security
- 📦 Build/Release
- 📚 Documentation
- ⚠️ Breaking Change
- 🗑️ Deprecated
