# 📜 Changelog - Antigravity Cleaner Pro

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [5.4.3] - 2026-10-31

### 🚀 The Freedom & Enterprise Release: Multi-Account Manager, Live Quota Monitor, Universal SSH Spoofing, Token Bridge & 1-Click Surgical Auto-Fix

This landmark release introduces complete process-level remote environment spoofing, seamless multi-account profile management, real-time quota telemetry, bi-directional token bridging, zero-downtime atomic patching, and surgical 429 quota cache wiping integrated directly into the 1-Click Auto-Fix workflow.

#### 👤 Multi-Account Profile Switcher (`pkg/account`)
- **Seamless Account Switching:** Safely save, list, switch, and delete multiple Google accounts across Antigravity IDE (`state.vscdb`) and Standalone Hub (`~/.config/Antigravity`).
- **OS Keyring Security:** Connects natively to the operating system's Secret Service / Keyring (`zalando/go-keyring`) to securely persist `refresh_token` and `id_token` credentials without plain-text storage or accidental session collisions.
- **Ergonomic Quick Navigation:** Numbered interactive profile selector (`[a] Multi-Account Manager`) with instant soft-reloading.

#### 📊 Real-Time Live Quota Monitor (`pkg/quota`)
- **Accurate Quota Telemetry:** Direct Connect RPC query to the running `language_server` process with automated CSRF token and port extraction.
- **Visual Tier & Credit Breakdown:** Matches official UI formatting with precise tracking of Gemini 2.5 Pro (5-hour and weekly buckets), Claude 3.5 Sonnet / GPT-4o allowances, and image generation credits with live reset countdown timers.
- **Offline Fallback:** Automated fallback to CloudCode API endpoints when the IDE is closed.

#### ⚡ 1-Click Smart Auto-Fix Integration (Step 4 Surgical Reset)
- **Direct Quota & Cookie Reset in Auto-Fix:** Option 1 (`runQuickAutoFix`) now automatically executes surgical cleanup of corrupted HTTP 429 quota buckets, stale DIPS states, TransportSecurity cookies, and orphaned WebSockets alongside binary bytecode patching.
- **Strict Data Preservation:** 100% of user chats, conversations, workspaces, custom prompt configurations, and keybindings remain completely untouched.

#### 🛰️ Universal SSH Environment Spoofing & Fraud Detection Bypass
- **Kernel & Process Sensor Reconstruction:** Exploits the core binary's internal remote detection routines at `0x6935840` (`isSSH()`) and `0x6920b72`, injecting `SSH_CONNECTION=127.0.0.1 1234 127.0.0.1 22` and `SSH_CLIENT=127.0.0.1 1234 22`.
- **File-Based Token Storage Activation:** Forces the runtime behavior `Using file-based token storage because SSH session detected` (`~/.gemini/jetski-standalone-oauth-token`), completely eliminating GNOME Keyring, D-Bus Secret Service deadlocks, and `Failed to fetch` account setup crashes.
- **Universal Surface Coverage:** SSH spoofing environment is automatically applied across the CLI launcher (`pkg/proxy/launcher.go`), smart desktop launchers (`antigravity-smart-launcher.sh`), and the binary wrapper (`pkg/patcher/wrapper.go`).

#### 🔑 Bi-Directional OAuth Token Bridge (`pkg/account/token_bridge.go`)
- **Keyring-to-File Synchronization:** Automatically extracts active OAuth tokens from the OS Secret Service / Keyring and exports them to Antigravity's standalone token file with strict `0600` permissions.
- **Zero-Logout Live Restarts:** Functions as an automated Post-Install Hook during binary wrapping and proxy launching, ensuring existing user logins never drop when applying patches or upgrading.

#### 🛡️ Non-Destructive Live Core Restart (Zero Session Disruption)
- **Harnessing `monitorLsCrashInternal`:** Utilizes the editor's built-in supervisor recovery loop in `app.asar` to perform graceful single-PID restarts of `language_server` within 2 seconds.
- **No More Application Restarts:** Eliminates crude `pkill antigravity` actions; active editor windows, open file tabs, and running subagent conversations remain fully active and connected during patching.

#### ⚙️ 114 Internal Core Flags Reverse-Engineered & Enterprise Shield
- **Complete Flag Discovery (`docs/INTERNAL_FLAGS_AND_REVERSE_ENGINEERING.md`):** Extracted and documented all 114 internal flags from registration function `0x7608aa0`.
- **Golden Anti-Sanction & Stability Flags (`DefaultEnterpriseFlagsOptions`):**
  - `--disable_unleash=true`: Shuts down Google Unleash remote geographic gating and dynamic region killswitches.
  - `--use_snapshot_mendelfp=true`: Enforces the embedded feature snapshot, halting dynamic server experiment downgrades.
  - `--disable_log_to_disk=true`: Stops massive multi-gigabyte log accumulation, eliminating `write EIO` pipe crashes.
  - `--suppress_cpp_logger_module_check_and_dont_log_cpp_to_stderr=true`: Completely silences noisy C++ internal logger output from stderr.
  - `--is_google_internal=true`, `--jetbox_release_channel=insiders`, `--stubby_auth_mode=business_licensed_enterprise`: Activates internal employee metadata and unlocks all experimental features.
  - `--allow_all_sidecars=true`, `--bypass_sidecar_user_config=true`: Forces all sidecar processes, companion tools, and MCP servers to stay active.

#### 🍏 macOS Apple Silicon AMFI Resigning & Quarantine Stripping
- **Mach-O Binary Auto-Detection (`pkg/patcher/atomic.go`):** Automatically recognizes Mach-O headers on macOS and applies `xattr -dr com.apple.quarantine`.
- **Ad-hoc Codesigning:** Executes `codesign --force --sign -` immediately after bytecode patching, permanently fixing AMFI code-signature violations (SIGTERM/SIGKILL crashes) on M1/M2/M3/M4 chips (Closes #37).

#### 🔍 Offline Application Discovery (`pkg/patcher/discovery.go`)
- **Proactive Path Scanning:** Dynamically discovers Antigravity installations even when closed by scanning system `PATH`, `~/Downloads`, `~/Desktop`, and `.desktop` entry specs.

#### 🎨 Cyberpunk Terminal User Interface (TUI)
- **Live Status Header:** Real-time visual status chips for Core, Wrapper, CLI, and ASAR patch state.
- **Clean Structure:** Ergonomic 1-key quick actions with an isolated `[m] Advanced Tools` sub-menu.

### 🐛 Resolved Issues & Bug Fixes
- **macOS Apple Silicon Startup Crash / AMFI Signature Invalidation (Closes #37):** Automatically strips quarantine extended attributes (`com.apple.quarantine`) and applies ad-hoc codesigning (`codesign --force --sign -`) to patched Mach-O binaries, resolving `Language server exited unexpectedly (code=null, signal=SIGTERM / SIGKILL)`.
- **Google Cloud Fraud Detection & Account Setup Failures (Closes #32, #34):** Injected universal SSH remote spoofing (`SSH_CONNECTION`, `SSH_CLIENT`), forcing file-based token storage and resolving `Failed to fetch` and GNOME Keyring / D-Bus secret service deadlocks.
- **Accidental Session Drop / Logout During Core Restart (Fixes #35):** Implemented bi-directional OAuth token bridging between system Keyring and Standalone storage (`jetski-standalone-oauth-token`) to ensure 100% session preservation.
- **IDE Window & Subagent Disconnects on Patching (Fixes #28):** Replaced destructive `pkill antigravity` with non-destructive single-PID soft restart leveraging Electron supervisor loop (`monitorLsCrashInternal`).
- **HTTP 429 Quota Exhaustion & Poisoned Cookies (Fixes #29, #33):** Integrated surgical 16-target quota and network persistent state reset directly into the 1-Click Auto-Fix flow.
- **Offline Target Discovery When Closed (Fixes #31):** Added proactive scanning across user Downloads, Desktop, and system `.desktop` files for systems where Antigravity is not running in background.

---

## [5.3.2] - 2026-10-19

### 🌐 Universal Multi-Architecture Discovery & Official Google RPM/DEB Hierarchy Support

This release eliminates all path-matching edge cases across modern Linux distributions (Fedora, RHEL, openSUSE, Ubuntu, Debian), adds multi-architecture binary resolution, and enables offline target discovery without requiring the IDE to be running.

- **Dynamic Multi-Architecture Binary Resolution:** Replaced static binary naming with dynamic architecture-aware detection across Linux (`language_server_linux_x64`, `language_server_linux_arm64`), macOS (`language_server_darwin_arm64`, `language_server_darwin_x64`), and Windows (`language_server_windows_x64.exe`, `language_server_windows_arm64.exe`).
- **Google Official RPM & DEB Layout Support:** Fully maps the modern Electron extension hierarchy (`resources/app/extensions/antigravity/bin/`) used by official Google Linux packages, ensuring seamless zero-config patching.
- **Offline Dynamic PATH Discovery:** Antigravity no longer needs to be actively running in RAM to be patched. The cleaner dynamically resolves the application root via system `PATH` (`exec.LookPath`) and follows symlinks (`filepath.EvalSymlinks`) directly to the core binary.
- **Self-Process Collision Guard:** Implemented `isSelfOrCleanerProcess` to strictly filter out current PIDs (`os.Getpid()`), ancestor PIDs, and cleaner/patcher binaries from process discovery tables, preventing false-positive self-detection.
- **Dynamic Architecture-Aware Wrapper:** Refactored `WrapperManager` to dynamically detect, back up (`<binName>.bin`), and generate custom shell/batch wrappers matching the exact detected architecture binary.
- **Frontend IDE JS & Launcher Resolution:** Added missing `/usr/share/antigravity/resources/app/out/main.js` candidates and resolved symlinks in desktop launcher creation for pristine icon and working directory setup.

### 🐛 Resolved Issues
- **Fedora 44 & Google RPM Discovery Failure (Closes #27):** Resolved the critical bug where `language_server` was reported missing on Fedora 44 with Google's official RPM repository due to self-process false positives and architecture suffix naming.

---

## [5.3.1] - 2026-10-15

### 🔇 Zero-Log ASAR Shield & EIO Broken-Pipe Hotfix

This release introduces an in-place bytecode patch for `electron-log` inside `app.asar` to eliminate annoying OS error popups on terminal disconnects.

- **Zero-Log ASAR Shield (`pkg/patcher/asar.go`):** Silences console logging in `electron-log` transport by patching `level: 'silly'` to `level: false` and shielding `writeFn` within a `try/catch` block.
- **EIO Exception Suppression:** Prevents annoying broken-pipe `EIO` crash dialogs and stdout exceptions when closing terminal sessions or launching through background wrappers.
- **Automatic ASAR Backup:** Generates timestamped `.agybak` backups before applying in-place ASAR modifications.

### 🐛 Resolved Issues
- **Broken Pipe Write EIO Popup (Closes #21, PR #22):** Fixed annoying OS modal error dialogs when Antigravity closes background stdout pipes.

---

## [5.3.0] - 2026-10-14

### 🛡️ 7-Flag Zero-Telemetry Engine, Staged Retro GUI & 16-Target Surgical Quota Reset

This release introduces major enterprise-level privacy protections, intelligent dynamic process discovery, and a completely streamlined, staged GUI workflow based on deep reverse-engineering of Antigravity binaries.

- **7-Flag Zero-Telemetry & Enterprise Engine:** Transparently injects internal runtime flags (`--disable_telemetry=true`, `--is_google_internal=true`, `--jetbox_release_channel=insiders`, `--stubby_auth_mode=business_licensed_enterprise`, `--use_local_chrome=true`, `--browser_eval_env=true`, `--local_chrome_headless=false`) via an atomic wrapper script while keeping original binary backups safe. Completely halts IP and git trace leakage to Google.
- **16-Target Surgical Quota 429 Reset:** Enhanced quota flusher surgically purges corrupted rate-limit cookies, DIPS databases, TransportSecurity, Trust Tokens, and session storage tokens across 16 targets in 0.002s without deleting chat history, projects, or credentials.
- **Streamlined Staged Retro KeyGen GUI (Win95):** Redesigned the retro interface into a guided 3-step workflow (`① Auto-Fix ➔ ② Launch IDE ➔ ③ Shortcut`) with compact dimensions (480x530) ensuring zero window scrollbars or clipping. Converted the telemetry toggle into an authentic Win95 interactive checkbox in the info bar.
- **Dynamic Multi-Platform Process Discovery:** Replaced static path assumptions with live `/proc` (Linux), `ps` (macOS), and PowerShell (Windows) process table inspection to dynamically discover custom Antigravity and `language_server` installations.
- **Session Immunity Shield:** Protected ancestor and active terminal PIDs from being terminated during `Auto-Fix`, ensuring interactive chat sessions are never disrupted.
- **XDG Standard Linux Desktop Launcher & Icon:** Automatically installs `antigravity.desktop` and a 256x256 high-resolution icon in `~/.local/share/icons/hicolor/256x256/apps/` for native desktop integration.
- **Official Telegram Support & Community Channel:** Integrated direct links to the official support channel ([@dallrep](https://t.me/dallrep)) both in terminal CLI and desktop GUI.

### 🐛 Resolved Issues & Community Disconnect Fixes
- **Dynamic Binary Discovery (Closes #11):** Resolved critical issue where Linux, macOS, and custom install directories reported `Core Engine CHECK Not Found` in Doctor, leaving `language_server` unpatched.
- **Persistent License Error #3501 (Closes #13, Discussion #12):** Eliminated the `#3501` license validation rejection by combining atomic 4-target unlocking (`main.js` + `workbench.desktop.main.js` + `language_server`) with `--stubby_auth_mode=business_licensed_enterprise` and `--is_google_internal=true`.
- **Location Error 400 & Disconnects (Discussion #14):** Fixed `User location is not supported` disconnects caused by Google telemetry beacons by deploying the 7-Flag Zero-Telemetry wrapper and 16-target token/DIPS surgical cache wipe.
- **macOS Startup & Environment Sanitation (Discussion #15):** Resolved blank screen on launch by providing clean process spawning and unified desktop launchers.

---

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

- 🎨 **Professional CLI & Retro UI** throughout the application
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

**Developed by:** Antigravity Cleaner Team  
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

**Apache License 2.0**  
**Copyright © 2024-2026 Antigravity Cleaner Authors. All Rights Reserved.**

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
