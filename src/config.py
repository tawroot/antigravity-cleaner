"""
Antigravity Cleaner - Configuration & Constants
===============================================
Shared configuration, constants, and theme definitions.
"""

# App Metadata
APP_NAME = "Antigravity Cleaner"
VERSION = "4.0.0"
GITHUB_URL = "https://github.com/tawroot/antigravity-cleaner"
LICENSE_TEXT = """
TAWANA NETWORK PROPRIETARY LICENSE
==================================

Copyright (c) 2024-2025 Tawana Mohammadi / Tawana Network
All Rights Reserved.

- Use: Free for everyone.
- Copy/Fork: STRICTLY PROHIBITED.
- Modify: STRICTLY PROHIBITED.
- Support: GitHub Issues Only.

Violation of these terms will result in legal action.
"""

# Logging
LOG_FILENAME_GUI = "antigravity-gui.log"
LOG_FILENAME_AGENT = "browser-helper-operations.log"

# Search Paths for cleanup (Base patterns)
WINDOWS_UNINSTALL_ROOTS = [
    r"SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall",
    r"SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall"
]

class AppleColors:
    """Apple Design System Colors"""
    
    # Minimal "Air" Palette
    BG_PRIMARY = "#FFFFFF"       # Pure White
    BG_SECONDARY = "#F5F5F7"     # Light Gray (Apple style)
    BG_TERTIARY = "#E8E8ED"      # Hover states
    
    # Text
    TEXT_PRIMARY = "#000000"     # Primary Text
    TEXT_SECONDARY = "#86868B"   # Subtitles
    TEXT_TERTIARY = "#D2D2D7"    # Disabled/Placeholder
    
    # BORDERS
    BORDER = "#E5E5EA"

    # Accents
    ACCENT_PRIMARY = "#007AFF"   # Apple Blue
    ACCENT_HOVER = "#0062CC"
    ACCENT_DANGER = "#FF3B30"    # Apple Red
    ACCENT_SUCCESS = "#34C759"   # Apple Green
    
    # Mappings
    BLUE = ACCENT_PRIMARY
    RED = ACCENT_DANGER
    GREEN = ACCENT_SUCCESS
    ORANGE = "#FF9500"
    PURPLE = "#AF52DE"
    PINK = "#FF2D55"
    TEAL = "#5AC8FA"
    YELLOW = "#FFCC00"
    
    LABEL = TEXT_PRIMARY
    SECONDARY_LABEL = TEXT_SECONDARY
    SEPARATOR = BORDER
    SIDEBAR_HOVER = BG_TERTIARY

class ProTheme:
    """Professional Theme Colors (Dark)"""
    BG_DARK = '#0d1117'
    BG_CARD = '#161b22'
    BG_HOVER = '#21262d'
    ACCENT_BLUE = '#58a6ff'
    ACCENT_GREEN = '#3fb950'
    ACCENT_RED = '#f85149'
    ACCENT_YELLOW = '#d29922'
    ACCENT_PURPLE = '#a371f7'
    ACCENT_CYAN = '#39c5cf'
    TEXT_PRIMARY = '#f0f6fc'
    TEXT_SECONDARY = '#8b949e'
    BORDER = '#30363d'
