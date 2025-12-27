# Technical Architecture

## Overview
Antigravity Cleaner (v3.0.0) is a native **PowerShell Automation Engine** designed to interface directly with the Windows File System and Registry to manage application states without external dependencies.

## Core Components

### 1. The Shell Engine (`Antigravity.ps1`)
*   **Runtime:** Windows PowerShell 5.1 / Core 7+
*   **State Management:** Stateless execution. Configuration is runtime-defined constants.
*   **UI Layer:** Host-based text UI with color-coded feedback (`Write-Host` wrappers).

### 2. Browser Hooking Mechanism
Instead of using heavy automation drivers (Selenium/Playwright), we use **File System Analysis**:
*   **Profile Detection:**
    *   Scans `%LOCALAPPDATA%` for Chromium-based browsers (Chrome, Edge, Brave).
    *   Parses `Local State` and `Preferences` JSON files to map Profile Folders (e.g., `Profile 14`) to User Accounts (e.g., `user@gmail.com`).
*   **Lock Handling:** Uses `Stop-Process` to forcefully release file locks before backup/restore operations to ensure data integrity.

### 3. Backup Strategy
*   **Light Mode (Smart Select):**
    *   Filters only critical SQLite databases (`Cookies`, `Login Data`, `Web Data`) and JSON configs.
    *   Ignores cache, temporary files, and heavy binary blobs (Service Workers).
    *   Result: ~95% size reduction compared to full profile copy.
*   **Full Mode (Robocopy Equivalent):**
    *   Recursive copy of the entire User Data directory.

### 4. Network Stack
*   **Connectivity Check:** Uses `.NET` classes (`System.Net.WebRequest`) for rapid HTTP status checks.
*   **Reset Operations:** Invokes native Windows binaries (`netsh.exe`, `ipconfig.exe`) with elevated privileges to flush the networking stack.

## Directory Structure
```
/antigravity-cleaner
│   Antigravity.ps1       # Main Engine
│   install.ps1           # One-line Installer
│
├── /Data                 # (Generated at runtime)
│   ├── /Sessions         # Backup storage
│   │   ├── /Google Chrome
│   │   └── /...
│   └── /Logs             # Application logs
│
└── /docs                 # Technical Documentation
```
