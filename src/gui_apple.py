"""
Antigravity Cleaner - macOS Style GUI
=====================================

Apple-inspired design with:
- Compact window size
- Rounded corners
- Glassmorphism effects
- San Francisco-style fonts
- Smooth animations

License: MIT
"""

import sys
import os
import platform
import threading
import webbrowser
from datetime import datetime

try:
    import customtkinter as ctk
    from tkinter import messagebox
    import tkinter as tk
except ImportError:
    print("Missing customtkinter. Install: pip install customtkinter")
    sys.exit(1)

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))

try:
    from browser_helper import BrowserHelper
    from network_optimizer import NetworkOptimizer
    from session_manager import SessionManager
    from antigravity_detector import AntigravityDetector
    from google_checker import GoogleServicesChecker
    from google_test_window import open_google_test_window
    import logging
except ImportError as e:
    print(f"Warning: {e}")
    BrowserHelper = None
    NetworkOptimizer = None
    SessionManager = None
    AntigravityDetector = None
    GoogleServicesChecker = None
    open_google_test_window = None





# ==================== Apple-Style Colors ====================

class AppleColors:
    """Apple Design System Colors"""
    
    # Primary
    BLUE = "#007AFF"
    GREEN = "#34C759"
    RED = "#FF3B30"
    ORANGE = "#FF9500"
    YELLOW = "#FFCC00"
    PURPLE = "#AF52DE"
    PINK = "#FF2D55"
    TEAL = "#5AC8FA"
    
    # Grays (Light mode)
    LABEL = "#000000"
    SECONDARY_LABEL = "#3C3C43"
    TERTIARY_LABEL = "#3C3C4399"
    
    # Backgrounds
    BG_PRIMARY = "#FFFFFF"
    BG_SECONDARY = "#F2F2F7"
    BG_TERTIARY = "#E5E5EA"
    
    # Sidebar
    SIDEBAR_BG = "#F5F5F7"
    SIDEBAR_ACTIVE = "#007AFF"
    SIDEBAR_HOVER = "#E8E8ED"
    
    # Separators
    SEPARATOR = "#C6C6C8"
    
    # Dark mode
    DARK_BG = "#1C1C1E"
    DARK_SECONDARY = "#2C2C2E"
    DARK_TERTIARY = "#3A3A3C"
    DARK_LABEL = "#FFFFFF"


# ==================== Translations ====================

TRANSLATIONS = {
    'en': {
        'title': 'Antigravity Cleaner',
        'cleaner': 'Cleaner',
        'sessions': 'Sessions',
        'browser': 'Browser',
        'network': 'Network',
        'quick_clean': 'Quick Clean',
        'deep_clean': 'Deep Clean',
        'full_repair': 'Full Repair',
        'preview': 'Preview Mode',
        'backup': 'Backup',
        'restore': 'Restore',
        'detect': 'Detect',
        'clean': 'Clean',
        'flush_dns': 'Flush DNS',
        'diagnostics': 'Diagnostics',
        'reset': 'Reset',
        'dark_mode': 'Dark Mode',
        'ready': 'Ready',
    },
    'fa': {
        'title': 'آنتی‌گرویتی کلینر',
        'cleaner': 'پاکسازی',
        'sessions': 'نشست‌ها',
        'browser': 'مرورگر',
        'network': 'شبکه',
        'quick_clean': 'پاکسازی سریع',
        'deep_clean': 'پاکسازی عمیق',
        'full_repair': 'تعمیر کامل',
        'preview': 'حالت پیش‌نمایش',
        'backup': 'پشتیبان',
        'restore': 'بازیابی',
        'detect': 'شناسایی',
        'clean': 'پاکسازی',
        'flush_dns': 'پاکسازی DNS',
        'diagnostics': 'تشخیص',
        'reset': 'ریست',
        'dark_mode': 'حالت تاریک',
        'ready': 'آماده',
    }
}


# ==================== Paths & Portable Mode ====================

def get_base_path():
    """Get the base path for data storage (Portable friendly)"""
    if getattr(sys, 'frozen', False):
        # Running as EXE
        base = os.path.dirname(sys.executable)
    else:
        # Running as Script
        base = os.path.dirname(os.path.abspath(__file__))
    
    # Check if we should use local 'data' folder (Portable standard)
    local_data = os.path.join(base, 'data')
    if os.path.exists(local_data) or getattr(sys, 'frozen', False):
        return local_data
    
    # Fallback to User Profile
    return os.path.join(os.path.expanduser('~'), '.antigravity-cleaner')

def setup_logger():
    data_dir = get_base_path()
    log_dir = os.path.join(data_dir, 'logs')
    os.makedirs(log_dir, exist_ok=True)
    
    logger = logging.getLogger('antigravity_apple')
    logger.setLevel(logging.DEBUG)
    
    if not logger.handlers:
        handler = logging.FileHandler(
            os.path.join(log_dir, 'apple-gui.log'),
            encoding='utf-8'
        )
        handler.setFormatter(logging.Formatter('[%(asctime)s] %(message)s'))
        logger.addHandler(handler)
    
    return logger



# ==================== Main App ====================

class AntigravityApp(ctk.CTk):
    """
    Antigravity Cleaner with Apple-style design
    """
    
    VERSION = "3.1"
    GITHUB_URL = "https://github.com/tawroot/antigravity-cleaner"
    
    def __init__(self):
        super().__init__()
        
        # Settings
        self.lang = 'en'
        self.dark_mode = False
        self.is_busy = False
        
        # Logger
        self.logger = setup_logger()
        
        # Initialize helpers
        self.init_helpers()
        
        # Window setup - COMPACT SIZE
        self.title(f"Antigravity Cleaner v{self.VERSION}")
        
        # Calculate center position
        screen_w = self.winfo_screenwidth()
        screen_h = self.winfo_screenheight()
        window_w = 950
        window_h = 600
        x = (screen_w - window_w) // 2
        y = (screen_h - window_h) // 2
        
        self.geometry(f"{window_w}x{window_h}+{x}+{y}")
        self.minsize(800, 500)
        self.maxsize(1200, 800)
        
        # Appearance
        ctk.set_appearance_mode("light")
        ctk.set_default_color_theme("blue")
        
        # Grid
        self.grid_columnconfigure(1, weight=1)
        self.grid_rowconfigure(0, weight=0) # Top Header fixed height
        self.grid_rowconfigure(1, weight=1) # Rest of the content

        
        # Build UI
        self.create_sidebar()
        self.create_content()
        
        # Show default page
        self.show_page("dashboard")
        
        # Log
        self.log("🚀 Antigravity Cleaner started")
        self.log(f"💻 {platform.system()} {platform.release()}")
    
    def init_helpers(self):
        """Initialize helper modules"""
        self.browser_helper = None
        self.network_optimizer = None
        self.session_manager = None
        self.detector = None
        
        if BrowserHelper:
            try:
                self.browser_helper = BrowserHelper(self.logger, dry_run=False)
            except: pass
        
        if NetworkOptimizer:
            try:
                self.network_optimizer = NetworkOptimizer(self.logger, dry_run=False)
            except: pass
        
        if SessionManager:
            try:
                data_dir = get_base_path()
                storage = os.path.join(data_dir, 'sessions')
                self.session_manager = SessionManager(storage, self.logger, dry_run=False)
            except: pass

        
        if AntigravityDetector:
            try:
                self.detector = AntigravityDetector(self.logger)
            except: pass
        
        if GoogleServicesChecker:
            try:
                self.google_checker = GoogleServicesChecker(self.logger)
            except: pass
        else:
            self.google_checker = None

    
    def t(self, key):
        return TRANSLATIONS[self.lang].get(key, key)
    
    # ==================== Sidebar ====================
    
    def create_sidebar(self):
        """Apple-style sidebar"""
        # Custom Top Header (Title)
        self.top_header = ctk.CTkFrame(self, height=40, fg_color=AppleColors.BG_PRIMARY, corner_radius=0)
        self.top_header.grid(row=0, column=0, columnspan=2, sticky="ew")
        self.top_header.grid_propagate(False)
        
        self.top_label = ctk.CTkLabel(
            self.top_header, 
            text="TAWANA NETWORK", 
            font=ctk.CTkFont(family="Inter", size=14, weight="bold"),
            text_color=AppleColors.BLUE
        )
        self.top_label.pack(side="left", padx=20)
        
        self.app_title = ctk.CTkLabel(
            self.top_header,
            text=f"Antigravity Cleaner Pro v{self.VERSION}",
            font=ctk.CTkFont(family="Inter", size=12),
            text_color=AppleColors.SECONDARY_LABEL
        )
        self.app_title.pack(side="right", padx=20)

        # Sidebar setup (now in row 1)
        self.sidebar = ctk.CTkFrame(
            self,
            width=220,
            corner_radius=0,
            fg_color=AppleColors.SIDEBAR_BG
        )
        self.sidebar.grid(row=1, column=0, sticky="nsew")
        self.sidebar.grid_propagate(False)
        
        # Logo
        logo_frame = ctk.CTkFrame(self.sidebar, fg_color="transparent")
        logo_frame.pack(pady=(30, 10), padx=20, fill="x")
        
        logo = ctk.CTkLabel(
            logo_frame,
            text="🚀",
            font=ctk.CTkFont(size=40)
        )
        logo.pack(side="left")
        
        title_frame = ctk.CTkFrame(logo_frame, fg_color="transparent")
        title_frame.pack(side="left", padx=12)
        
        title = ctk.CTkLabel(
            title_frame,
            text="Antigravity",
            font=ctk.CTkFont(family="SF Pro Display", size=18, weight="bold"),
            text_color=AppleColors.LABEL
        )
        title.pack(anchor="w")
        
        subtitle = ctk.CTkLabel(
            title_frame,
            text="Cleaner Pro",
            font=ctk.CTkFont(family="SF Pro Display", size=14),
            text_color=AppleColors.SECONDARY_LABEL
        )
        subtitle.pack(anchor="w")

        
        # Separator
        sep = ctk.CTkFrame(self.sidebar, height=1, fg_color=AppleColors.SEPARATOR)
        sep.pack(fill="x", padx=15, pady=15)
        
        # Navigation
        self.nav_btns = {}
        nav_items = [
            ("dashboard", "📊", "Dashboard"),
            ("cleaner", "🧹", self.t('cleaner')),
            ("sessions", "💾", self.t('sessions')),
            ("browser", "🌐", self.t('browser')),
            ("network", "🔧", self.t('network')),
        ]

        
        for key, icon, text in nav_items:
            btn = ctk.CTkButton(
                self.sidebar,
                text=f"  {icon}  {text}",
                font=ctk.CTkFont(size=14),
                height=40,
                corner_radius=10,
                fg_color="transparent",
                text_color=AppleColors.LABEL,
                hover_color=AppleColors.SIDEBAR_HOVER,
                anchor="w",
                command=lambda k=key: self.show_page(k)
            )
            btn.pack(fill="x", padx=10, pady=2)
            self.nav_btns[key] = btn
        
        # Spacer
        spacer = ctk.CTkFrame(self.sidebar, fg_color="transparent")
        spacer.pack(fill="both", expand=True)
        
        # Bottom branding
        branding_frame = ctk.CTkFrame(self.sidebar, fg_color="transparent")
        branding_frame.pack(side="bottom", fill="x", pady=20, padx=15)
        
        credit = ctk.CTkLabel(
            branding_frame,
            text="Powered by",
            font=ctk.CTkFont(size=10),
            text_color=AppleColors.SECONDARY_LABEL
        )
        credit.pack()
        
        tawana = ctk.CTkLabel(
            branding_frame,
            text="TAWANA NETWORK",
            font=ctk.CTkFont(size=12, weight="bold"),
            text_color=AppleColors.BLUE
        )
        tawana.pack()
        
        license_btn = ctk.CTkLabel(
            branding_frame,
            text="Proprietary License",
            font=ctk.CTkFont(size=10, underline=True),
            text_color=AppleColors.SECONDARY_LABEL,
            cursor="hand2"
        )
        license_btn.pack(pady=(5, 0))
        license_btn.bind("<Button-1>", lambda e: self.show_license())
        
        # Bottom controls (Theme/Lang)
        controls = ctk.CTkFrame(self.sidebar, fg_color="transparent")
        controls.pack(side="bottom", fill="x", pady=0, padx=15)
        
        self.theme_btn = ctk.CTkButton(
            controls,
            text="🌙",
            width=40,
            height=30,
            corner_radius=8,
            fg_color=AppleColors.BG_TERTIARY,
            text_color=AppleColors.LABEL,
            hover_color=AppleColors.SEPARATOR,
            command=self.toggle_theme
        )
        self.theme_btn.pack(side="left", padx=2)
        
        self.lang_btn = ctk.CTkButton(
            controls,
            text="FA",
            width=40,
            height=30,
            corner_radius=8,
            fg_color=AppleColors.BG_TERTIARY,
            text_color=AppleColors.LABEL,
            hover_color=AppleColors.SEPARATOR,
            command=self.toggle_lang
        )
        self.lang_btn.pack(side="left", padx=2)

        
        # GitHub
        github_btn = ctk.CTkButton(
            controls,
            text="⭐",
            width=35,
            height=30,
            corner_radius=15,
            font=ctk.CTkFont(size=14),
            fg_color=AppleColors.PURPLE,
            hover_color="#9B3DD1",
            command=lambda: webbrowser.open(self.GITHUB_URL)
        )
        github_btn.pack(side="right", padx=5)
    
    def show_page(self, page):
        """Switch pages"""
        self.current_page = page
        
        # Update nav buttons
        for key, btn in self.nav_btns.items():
            if key == page:
                btn.configure(fg_color=AppleColors.BLUE, text_color="#FFFFFF")
            else:
                btn.configure(fg_color="transparent", text_color=AppleColors.LABEL)
        
        # Hide all pages
        for child in self.content.winfo_children():
            child.grid_forget()
        
        # Show selected
        pages = {
            "dashboard": self.page_dashboard,
            "cleaner": self.page_cleaner,
            "sessions": self.page_sessions,
            "browser": self.page_browser,
            "network": self.page_network,
        }
        pages[page].grid(row=0, column=0, sticky="nsew", padx=20, pady=20)

    
    def show_license(self):
        """Show proprietary license information"""
        license_text = """
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
        messagebox.showinfo("License & Copyright", license_text)
    
    def toggle_theme(self):

        """Toggle dark/light mode"""
        self.dark_mode = not self.dark_mode
        ctk.set_appearance_mode("dark" if self.dark_mode else "light")
    
    def toggle_lang(self):
        """Toggle language"""
        self.lang = 'fa' if self.lang == 'en' else 'en'
        self.lang_btn.configure(text="🌐 EN" if self.lang == 'fa' else "🌐 FA")
    
    # ==================== Content ====================
    
    def create_content(self):
        """Create main content area"""
        self.content = ctk.CTkFrame(self, fg_color=AppleColors.BG_SECONDARY, corner_radius=0)
        self.content.grid(row=1, column=1, sticky="nsew")
        self.content.grid_columnconfigure(0, weight=1)
        self.content.grid_rowconfigure(0, weight=1)
        
        # Create pages
        self.create_dashboard_page()
        self.create_cleaner_page()
        self.create_sessions_page()
        self.create_browser_page()
        self.create_network_page()

    
    def create_card(self, parent, **kwargs):
        """Create Apple-style card"""
        return ctk.CTkFrame(
            parent,
            corner_radius=12,
            fg_color=AppleColors.BG_PRIMARY,
            border_width=1,
            border_color=AppleColors.BG_TERTIARY,
            **kwargs
        )
    
    def create_button(self, parent, text, color="blue", icon=None, **kwargs):
        """Create Apple-style button"""
        colors = {
            "blue": AppleColors.BLUE,
            "green": AppleColors.GREEN,
            "red": AppleColors.RED,
            "orange": AppleColors.ORANGE,
            "gray": AppleColors.BG_TERTIARY,
        }
        
        bg = colors.get(color, AppleColors.BLUE)
        fg = "#FFFFFF" if color != "gray" else AppleColors.LABEL
        
        display = f"{icon} {text}" if icon else text
        
        return ctk.CTkButton(
            parent,
            text=display,
            corner_radius=10,
            height=36,
            font=ctk.CTkFont(size=13, weight="bold"),
            fg_color=bg,
            text_color=fg,
            hover_color=self._lighten(bg) if color != "gray" else AppleColors.SEPARATOR,
            **kwargs
        )
    
    def _lighten(self, hex_color, factor=1.15):
        """Lighten a color"""
        hex_color = hex_color.lstrip('#')
        rgb = tuple(int(hex_color[i:i+2], 16) for i in (0, 2, 4))
        lighter = tuple(min(255, int(c * factor)) for c in rgb)
        return f"#{lighter[0]:02x}{lighter[1]:02x}{lighter[2]:02x}"
    
    # ==================== Dashboard Page ====================
    
    def create_dashboard_page(self):
        """Dashboard with health score and system status"""
        self.page_dashboard = ctk.CTkFrame(self.content, fg_color="transparent")
        self.page_dashboard.grid_columnconfigure(0, weight=1)
        self.page_dashboard.grid_columnconfigure(1, weight=1)
        self.page_dashboard.grid_rowconfigure(1, weight=1)
        
        # Header
        header = ctk.CTkLabel(
            self.page_dashboard,
            text="📊 Dashboard",
            font=ctk.CTkFont(size=24, weight="bold"),
            text_color=AppleColors.LABEL
        )
        header.grid(row=0, column=0, columnspan=2, sticky="w", pady=(0, 20))
        
        # Health Score Card (Left)
        health_card = self.create_card(self.page_dashboard)
        health_card.grid(row=1, column=0, sticky="nsew", padx=(0, 10), pady=(0, 10))
        
        health_title = ctk.CTkLabel(
            health_card,
            text="🏥 System Health",
            font=ctk.CTkFont(size=16, weight="bold"),
            text_color=AppleColors.LABEL
        )
        health_title.pack(pady=(20, 15), padx=20, anchor="w")
        
        # Health score display
        self.health_score_label = ctk.CTkLabel(
            health_card,
            text="--",
            font=ctk.CTkFont(size=72, weight="bold"),
            text_color=AppleColors.GREEN
        )
        self.health_score_label.pack(pady=10)
        
        self.health_status_label = ctk.CTkLabel(
            health_card,
            text="Scanning...",
            font=ctk.CTkFont(size=14),
            text_color=AppleColors.SECONDARY_LABEL
        )
        self.health_status_label.pack(pady=(0, 10))
        
        # Scan button
        scan_btn = self.create_button(
            health_card, "🔍 Scan Now", "blue",
            command=self.run_scan
        )
        scan_btn.pack(pady=20, padx=20, fill="x")
        
        # Status Card (Right)
        status_card = self.create_card(self.page_dashboard)
        status_card.grid(row=1, column=1, sticky="nsew", padx=(10, 0), pady=(0, 10))
        
        status_title = ctk.CTkLabel(
            status_card,
            text="📋 Status",
            font=ctk.CTkFont(size=16, weight="bold"),
            text_color=AppleColors.LABEL
        )
        status_title.pack(pady=(20, 15), padx=20, anchor="w")
        
        # Status items frame
        self.status_frame = ctk.CTkFrame(status_card, fg_color="transparent")
        self.status_frame.pack(fill="both", expand=True, padx=20, pady=10)
        
        # Status items (placeholders)
        status_items = [
            ("installed", "🔧 Antigravity Installed", "Checking..."),
            ("running", "▶️ Antigravity Running", "Checking..."),
            ("leftovers", "📁 Leftover Files", "Checking..."),
            ("browsers", "🌐 Browsers Detected", "Checking..."),
            ("sessions", "💾 Saved Sessions", "Checking..."),
        ]
        
        self.status_labels = {}
        for key, title, value in status_items:
            row = ctk.CTkFrame(self.status_frame, fg_color="transparent")
            row.pack(fill="x", pady=5)
            
            title_label = ctk.CTkLabel(
                row,
                text=title,
                font=ctk.CTkFont(size=13),
                text_color=AppleColors.LABEL
            )
            title_label.pack(side="left")
            
            value_label = ctk.CTkLabel(
                row,
                text=value,
                font=ctk.CTkFont(size=13, weight="bold"),
                text_color=AppleColors.SECONDARY_LABEL
            )
            value_label.pack(side="right")
            self.status_labels[key] = value_label
        
        # Google Services Card (row 2, left)
        google_card = self.create_card(self.page_dashboard)
        google_card.grid(row=2, column=0, sticky="nsew", padx=(0, 10), pady=(10, 0))
        
        google_title = ctk.CTkLabel(
            google_card,
            text="🌐 Google Services",
            font=ctk.CTkFont(size=14, weight="bold"),
            text_color=AppleColors.LABEL
        )
        google_title.pack(pady=(15, 10), padx=20, anchor="w")
        
        # Google services status
        self.google_status_frame = ctk.CTkScrollableFrame(
            google_card,
            height=100,
            fg_color=AppleColors.BG_SECONDARY,
            corner_radius=8
        )
        self.google_status_frame.pack(fill="both", expand=True, padx=15, pady=(0, 10))
        
        # Placeholder
        self.google_status_placeholder = ctk.CTkLabel(
            self.google_status_frame,
            text="Click 'Test Google' to check",
            font=ctk.CTkFont(size=12),
            text_color=AppleColors.SECONDARY_LABEL
        )
        self.google_status_placeholder.pack(pady=20)
        
        # Test button
        test_google_btn = self.create_button(
            google_card, "🔍 Test Google", "blue",
            command=self.test_google_services
        )
        test_google_btn.pack(pady=(0, 15), padx=20, fill="x")
        
        # Quick Actions Card (row 2, right)
        actions_card = self.create_card(self.page_dashboard)
        actions_card.grid(row=2, column=1, sticky="nsew", padx=(10, 0), pady=(10, 0))
        
        actions_title = ctk.CTkLabel(
            actions_card,
            text="⚡ Quick Actions",
            font=ctk.CTkFont(size=14, weight="bold"),
            text_color=AppleColors.LABEL
        )
        actions_title.pack(pady=(15, 10), padx=20, anchor="w")
        
        btn_frame = ctk.CTkFrame(actions_card, fg_color="transparent")
        btn_frame.pack(fill="both", expand=True, padx=20, pady=(0, 15))
        
        quick_btns = [
            ("🧹 Quick Clean", "green", lambda: self.run_clean('quick')),
            ("💾 Backup All", "blue", self.backup_all_sessions),
            ("🔧 Full Repair", "red", lambda: self.run_clean('full')),
            ("🔄 Refresh", "gray", self.run_scan),
        ]
        
        for text, color, cmd in quick_btns:
            btn = self.create_button(btn_frame, text, color, command=cmd)
            btn.pack(pady=3, fill="x")
        
        # Auto-scan on load
        self.after(500, self.run_scan)
    
    def test_google_services(self):
        """Open Google Services Test window"""
        if not open_google_test_window:
            messagebox.showwarning("Warning", "Google Test Window not available")
            return
        
        self.log("🌐 Opening Google Services Test...")
        
        # Open popup window
        open_google_test_window(self, self.google_checker, self.logger)

    
    def run_scan(self):
        """Run Antigravity detection scan"""
        self.health_status_label.configure(text="Scanning...")
        
        def do_scan():
            # Update health score
            if self.detector:
                results = self.detector.detect_all()
                score = results['health_score']
                status = results['status']
                
                # Determine color based on score
                if score >= 90:
                    color = AppleColors.GREEN
                elif score >= 70:
                    color = AppleColors.ORANGE
                else:
                    color = AppleColors.RED
                
                self.after(0, lambda: self.health_score_label.configure(text=str(score), text_color=color))
                self.after(0, lambda: self.health_status_label.configure(text=status))
                
                # Update status items
                installed = "Yes" if results['is_installed'] else "No"
                running = "Yes" if results['is_running'] else "No"
                leftovers = f"{len(results['leftover_paths'])} files ({results['leftover_size_human']})"
                
                self.after(0, lambda: self.status_labels['installed'].configure(
                    text=installed, 
                    text_color=AppleColors.RED if results['is_installed'] else AppleColors.GREEN
                ))
                self.after(0, lambda: self.status_labels['running'].configure(
                    text=running,
                    text_color=AppleColors.RED if results['is_running'] else AppleColors.GREEN
                ))
                self.after(0, lambda: self.status_labels['leftovers'].configure(text=leftovers))
            else:
                self.after(0, lambda: self.health_score_label.configure(text="100", text_color=AppleColors.GREEN))
                self.after(0, lambda: self.health_status_label.configure(text="✅ Clean"))
            
            # Update browser count
            if self.browser_helper:
                browsers = self.browser_helper.detect_installed_browsers()
                self.after(0, lambda: self.status_labels['browsers'].configure(text=str(len(browsers))))
            
            # Update sessions count
            if self.session_manager:
                sessions = self.session_manager.list_saved_sessions()
                self.after(0, lambda: self.status_labels['sessions'].configure(text=str(len(sessions))))
        
        threading.Thread(target=do_scan, daemon=True).start()
    
    def backup_all_sessions(self):
        """Backup all browser profiles"""
        if not self.browser_helper or not self.session_manager:
            messagebox.showwarning("Warning", "Session Manager not available")
            return
        
        if not messagebox.askyesno("Confirm", "Backup all browser profiles?"):
            return
        
        self.log("💾 Starting backup of all profiles...")
        
        def do_backup():
            count = 0
            browsers = self.browser_helper.detect_installed_browsers()
            
            for browser in browsers:
                if self.browser_helper.is_browser_running(browser):
                    self.after(0, lambda b=browser: self.log(f"⚠️ {b} is running, skipping..."))
                    continue
                
                profiles = self.browser_helper.get_browser_profiles_with_email(browser)
                for profile in profiles:
                    try:
                        result = self.session_manager.backup_session(browser, profile['path'])
                        if result:
                            count += 1
                            self.after(0, lambda p=profile: self.log(f"✅ Backed up: {p['name']}"))
                    except Exception as e:
                        self.after(0, lambda e=e: self.log(f"❌ Error: {e}"))
            
            self.after(0, lambda: self.log(f"💾 Completed! {count} profiles backed up"))
            self.after(0, lambda: messagebox.showinfo("Success", f"Backed up {count} profiles!"))
            self.after(0, self.run_scan)
        
        threading.Thread(target=do_backup, daemon=True).start()
    
    # ==================== Cleaner Page ====================

    
    def create_cleaner_page(self):
        """Cleaner page"""
        self.page_cleaner = ctk.CTkFrame(self.content, fg_color="transparent")
        self.page_cleaner.grid_columnconfigure(1, weight=1)
        self.page_cleaner.grid_rowconfigure(0, weight=1)
        
        # Left: Actions
        actions_card = self.create_card(self.page_cleaner, width=220)
        actions_card.grid(row=0, column=0, sticky="ns", padx=(0, 15))
        actions_card.grid_propagate(False)
        
        # Header
        header = ctk.CTkLabel(
            actions_card,
            text="🧹 " + self.t('cleaner'),
            font=ctk.CTkFont(size=18, weight="bold"),
            text_color=AppleColors.LABEL
        )
        header.pack(pady=(20, 15), padx=15, anchor="w")
        
        # Preview switch
        self.preview_var = ctk.BooleanVar(value=False)
        preview = ctk.CTkSwitch(
            actions_card,
            text=self.t('preview'),
            font=ctk.CTkFont(size=12),
            variable=self.preview_var,
            progress_color=AppleColors.BLUE
        )
        preview.pack(padx=15, pady=10, anchor="w")
        
        # Separator
        sep = ctk.CTkFrame(actions_card, height=1, fg_color=AppleColors.SEPARATOR)
        sep.pack(fill="x", padx=15, pady=15)
        
        # Buttons
        buttons = [
            (self.t('quick_clean'), "green", lambda: self.run_clean('quick')),
            (self.t('deep_clean'), "blue", lambda: self.run_clean('deep')),
            (self.t('full_repair'), "red", lambda: self.run_clean('full')),
        ]
        
        for text, color, cmd in buttons:
            btn = self.create_button(actions_card, text, color, command=cmd)
            btn.pack(padx=15, pady=5, fill="x")
        
        # Right: Log
        log_card = self.create_card(self.page_cleaner)
        log_card.grid(row=0, column=1, sticky="nsew")
        log_card.grid_columnconfigure(0, weight=1)
        log_card.grid_rowconfigure(1, weight=1)
        
        # Log header
        log_header = ctk.CTkFrame(log_card, fg_color="transparent")
        log_header.grid(row=0, column=0, sticky="ew", padx=15, pady=10)
        
        log_title = ctk.CTkLabel(
            log_header,
            text="📋 Log",
            font=ctk.CTkFont(size=14, weight="bold"),
            text_color=AppleColors.LABEL
        )
        log_title.pack(side="left")
        
        clear_btn = ctk.CTkButton(
            log_header,
            text="Clear",
            width=60,
            height=28,
            corner_radius=14,
            font=ctk.CTkFont(size=11),
            fg_color=AppleColors.BG_TERTIARY,
            text_color=AppleColors.SECONDARY_LABEL,
            hover_color=AppleColors.SEPARATOR,
            command=self.clear_log
        )
        clear_btn.pack(side="right")
        
        # Log text
        self.log_text = ctk.CTkTextbox(
            log_card,
            font=ctk.CTkFont(family="SF Mono", size=11),
            fg_color=AppleColors.BG_SECONDARY,
            text_color=AppleColors.LABEL,
            corner_radius=8,
            state="disabled"
        )
        self.log_text.grid(row=1, column=0, sticky="nsew", padx=15, pady=(0, 15))
    
    # ==================== Sessions Page ====================
    
    def create_sessions_page(self):
        """Sessions page"""
        self.page_sessions = ctk.CTkFrame(self.content, fg_color="transparent")
        self.page_sessions.grid_columnconfigure(1, weight=1)
        self.page_sessions.grid_rowconfigure(0, weight=1)
        
        # Left: Controls
        ctrl_card = self.create_card(self.page_sessions, width=240)
        ctrl_card.grid(row=0, column=0, sticky="ns", padx=(0, 15))
        ctrl_card.grid_propagate(False)
        
        header = ctk.CTkLabel(
            ctrl_card,
            text="💾 " + self.t('sessions'),
            font=ctk.CTkFont(size=18, weight="bold"),
            text_color=AppleColors.LABEL
        )
        header.pack(pady=(20, 15), padx=15, anchor="w")
        
        # Browser dropdown
        browser_label = ctk.CTkLabel(
            ctrl_card,
            text="Browser",
            font=ctk.CTkFont(size=12, weight="bold"),
            text_color=AppleColors.SECONDARY_LABEL
        )
        browser_label.pack(padx=15, anchor="w")
        
        self.browser_combo = ctk.CTkComboBox(
            ctrl_card,
            width=200,
            height=32,
            corner_radius=8,
            font=ctk.CTkFont(size=12),
            command=self.on_browser_change
        )
        self.browser_combo.pack(padx=15, pady=(5, 15))
        
        # Profile list
        profile_label = ctk.CTkLabel(
            ctrl_card,
            text="Profile",
            font=ctk.CTkFont(size=12, weight="bold"),
            text_color=AppleColors.SECONDARY_LABEL
        )
        profile_label.pack(padx=15, anchor="w")
        
        self.profile_list = ctk.CTkScrollableFrame(
            ctrl_card,
            width=200,
            height=150,
            corner_radius=8,
            fg_color=AppleColors.BG_SECONDARY
        )
        self.profile_list.pack(padx=15, pady=5, fill="both", expand=True)
        
        # Buttons
        btn_frame = ctk.CTkFrame(ctrl_card, fg_color="transparent")
        btn_frame.pack(fill="x", padx=15, pady=15)
        
        backup_btn = self.create_button(
            btn_frame, self.t('backup'), "green",
            command=self.backup_session
        )
        backup_btn.pack(fill="x", pady=3)
        
        restore_btn = self.create_button(
            btn_frame, self.t('restore'), "blue",
            command=self.restore_session
        )
        restore_btn.pack(fill="x", pady=3)
        
        # Right: Sessions list
        sessions_card = self.create_card(self.page_sessions)
        sessions_card.grid(row=0, column=1, sticky="nsew")
        sessions_card.grid_columnconfigure(0, weight=1)
        sessions_card.grid_rowconfigure(1, weight=1)
        
        sessions_header = ctk.CTkFrame(sessions_card, fg_color="transparent")
        sessions_header.grid(row=0, column=0, sticky="ew", padx=15, pady=10)
        
        sessions_title = ctk.CTkLabel(
            sessions_header,
            text="📁 Saved Sessions",
            font=ctk.CTkFont(size=14, weight="bold"),
            text_color=AppleColors.LABEL
        )
        sessions_title.pack(side="left")
        
        refresh_btn = ctk.CTkButton(
            sessions_header,
            text="🔄",
            width=35,
            height=28,
            corner_radius=14,
            fg_color=AppleColors.BLUE,
            command=self.refresh_sessions
        )
        refresh_btn.pack(side="right")
        
        self.sessions_list = ctk.CTkScrollableFrame(
            sessions_card,
            corner_radius=8,
            fg_color=AppleColors.BG_SECONDARY
        )
        self.sessions_list.grid(row=1, column=0, sticky="nsew", padx=15, pady=(0, 15))
        
        # Load data after UI ready
        self.after(100, self.load_session_data)
    
    # ==================== Browser Page ====================
    
    def create_browser_page(self):
        """Browser helper page"""
        self.page_browser = ctk.CTkFrame(self.content, fg_color="transparent")
        self.page_browser.grid_columnconfigure(0, weight=1)
        self.page_browser.grid_rowconfigure(1, weight=1)
        
        # Header
        header_frame = ctk.CTkFrame(self.page_browser, fg_color="transparent")
        header_frame.grid(row=0, column=0, sticky="ew", pady=(0, 15))
        
        header = ctk.CTkLabel(
            header_frame,
            text="🌐 Browser Helper",
            font=ctk.CTkFont(size=20, weight="bold"),
            text_color=AppleColors.LABEL
        )
        header.pack(side="left")
        
        # Buttons
        btn_frame = ctk.CTkFrame(header_frame, fg_color="transparent")
        btn_frame.pack(side="right")
        
        detect_btn = self.create_button(
            btn_frame, self.t('detect'), "blue", "🔍",
            command=self.detect_browsers
        )
        detect_btn.pack(side="left", padx=5)
        
        clean_btn = self.create_button(
            btn_frame, self.t('clean'), "green", "🧹",
            command=self.clean_traces
        )
        clean_btn.pack(side="left", padx=5)
        
        # Browser list card
        browser_card = self.create_card(self.page_browser)
        browser_card.grid(row=1, column=0, sticky="nsew")
        browser_card.grid_columnconfigure(0, weight=1)
        browser_card.grid_rowconfigure(0, weight=1)
        
        self.browser_list = ctk.CTkScrollableFrame(
            browser_card,
            corner_radius=8,
            fg_color=AppleColors.BG_SECONDARY
        )
        self.browser_list.pack(fill="both", expand=True, padx=15, pady=15)
    
    # ==================== Network Page ====================
    
    def create_network_page(self):
        """Network tools page"""
        self.page_network = ctk.CTkFrame(self.content, fg_color="transparent")
        self.page_network.grid_columnconfigure(0, weight=1)
        self.page_network.grid_rowconfigure(1, weight=1)
        
        # Header
        header_frame = ctk.CTkFrame(self.page_network, fg_color="transparent")
        header_frame.grid(row=0, column=0, sticky="ew", pady=(0, 15))
        
        header = ctk.CTkLabel(
            header_frame,
            text="🔧 Network Tools",
            font=ctk.CTkFont(size=20, weight="bold"),
            text_color=AppleColors.LABEL
        )
        header.pack(side="left")
        
        # Buttons
        btn_frame = ctk.CTkFrame(header_frame, fg_color="transparent")
        btn_frame.pack(side="right")
        
        dns_btn = self.create_button(
            btn_frame, self.t('flush_dns'), "green",
            command=self.flush_dns
        )
        dns_btn.pack(side="left", padx=5)
        
        diag_btn = self.create_button(
            btn_frame, self.t('diagnostics'), "blue",
            command=self.run_diagnostics
        )
        diag_btn.pack(side="left", padx=5)
        
        reset_btn = self.create_button(
            btn_frame, self.t('reset'), "red",
            command=self.reset_network
        )
        reset_btn.pack(side="left", padx=5)
        
        # Output card
        output_card = self.create_card(self.page_network)
        output_card.grid(row=1, column=0, sticky="nsew")
        output_card.grid_columnconfigure(0, weight=1)
        output_card.grid_rowconfigure(0, weight=1)
        
        self.network_text = ctk.CTkTextbox(
            output_card,
            font=ctk.CTkFont(family="SF Mono", size=11),
            fg_color=AppleColors.BG_SECONDARY,
            text_color=AppleColors.LABEL,
            corner_radius=8,
            state="disabled"
        )
        self.network_text.pack(fill="both", expand=True, padx=15, pady=15)
    
    # ==================== Actions ====================
    
    def log(self, msg):
        """Add to log"""
        self.log_text.configure(state="normal")
        time = datetime.now().strftime("%H:%M:%S")
        self.log_text.insert("end", f"[{time}] {msg}\n")
        self.log_text.see("end")
        self.log_text.configure(state="disabled")
        self.logger.info(msg)
    
    def clear_log(self):
        """Clear log"""
        self.log_text.configure(state="normal")
        self.log_text.delete("1.0", "end")
        self.log_text.configure(state="disabled")
    
    def run_clean(self, mode):
        """Run cleaning"""
        if self.is_busy:
            return
        
        names = {'quick': 'Quick Clean', 'deep': 'Deep Clean', 'full': 'Full Repair'}
        
        if not messagebox.askyesno("Confirm", f"Start {names[mode]}?"):
            return
        
        self.is_busy = True
        self.log(f"🚀 Starting {names[mode]}...")
        
        def do_clean():
            try:
                dry = self.preview_var.get()
                self._clean_files(mode, dry)
                if mode == 'full':
                    self._reset_network_cmd(dry)
                self.after(0, lambda: self.log("✅ Completed!"))
                self.after(0, lambda: messagebox.showinfo("Success", "Operation completed!"))
            except Exception as e:
                self.after(0, lambda: self.log(f"❌ Error: {e}"))
            finally:
                self.is_busy = False
        
        threading.Thread(target=do_clean, daemon=True).start()
    
    def _clean_files(self, mode, dry_run):
        """Clean files"""
        import shutil
        
        home = os.path.expanduser("~")
        paths = []
        
        if platform.system() == "Windows":
            local = os.environ.get("LOCALAPPDATA", "")
            roaming = os.environ.get("APPDATA", "")
            paths = [
                os.path.join(local, "Antigravity"),
                os.path.join(roaming, "Antigravity"),
            ]
        elif platform.system() == "Darwin":
            paths = [
                os.path.join(home, "Library", "Application Support", "Antigravity"),
                os.path.join(home, "Library", "Caches", "Antigravity"),
            ]
        else:
            paths = [
                os.path.join(home, ".config", "Antigravity"),
                os.path.join(home, ".cache", "Antigravity"),
            ]
        
        for path in paths:
            if os.path.exists(path):
                if dry_run:
                    self.after(0, lambda p=path: self.log(f"[Preview] {p}"))
                else:
                    try:
                        shutil.rmtree(path) if os.path.isdir(path) else os.remove(path)
                        self.after(0, lambda p=path: self.log(f"🗑️ Removed: {p}"))
                    except Exception as e:
                        self.after(0, lambda e=e: self.log(f"⚠️ {e}"))
    
    def _reset_network_cmd(self, dry_run):
        """Reset network"""
        import subprocess
        
        cmds = ["ipconfig /flushdns"] if platform.system() == "Windows" else ["dscacheutil -flushcache"]
        
        for cmd in cmds:
            if dry_run:
                self.after(0, lambda c=cmd: self.log(f"[Preview] {c}"))
            else:
                subprocess.run(cmd, shell=True, capture_output=True)
                self.after(0, lambda c=cmd: self.log(f"✅ {c}"))
    
    # Session Manager
    def load_session_data(self):
        if self.browser_helper:
            browsers = self.browser_helper.detect_installed_browsers()
            self.browser_combo.configure(values=browsers)
            if browsers:
                self.browser_combo.set(browsers[0])
                self.on_browser_change(browsers[0])
        self.refresh_sessions()
    
    def on_browser_change(self, browser):
        if not self.browser_helper:
            return
        
        for w in self.profile_list.winfo_children():
            w.destroy()
        
        profiles = self.browser_helper.get_browser_profiles_with_email(browser)
        self.profiles_data = profiles
        self.selected_profile_idx = None
        
        for i, p in enumerate(profiles):
            email = p.get('email') or 'No email'
            btn = ctk.CTkRadioButton(
                self.profile_list,
                text=f"{p['name']}\n{email}",
                font=ctk.CTkFont(size=11),
                command=lambda idx=i: self.select_profile(idx)
            )
            btn.pack(anchor="w", pady=3)
    
    def select_profile(self, idx):
        self.selected_profile_idx = idx
    
    def refresh_sessions(self):
        for w in self.sessions_list.winfo_children():
            w.destroy()
        
        if not self.session_manager:
            return
        
        sessions = self.session_manager.list_saved_sessions()
        self.sessions_data = sessions
        
        if not sessions:
            label = ctk.CTkLabel(
                self.sessions_list,
                text="No saved sessions",
                text_color=AppleColors.SECONDARY_LABEL
            )
            label.pack(pady=20)
            return
        
        for i, s in enumerate(sessions):
            frame = ctk.CTkFrame(
                self.sessions_list,
                fg_color=AppleColors.BG_PRIMARY,
                corner_radius=8
            )
            frame.pack(fill="x", pady=3)
            
            status = "🟢" if not s.get('expired') else "🔴"
            info = ctk.CTkLabel(
                frame,
                text=f"{status} {s['name']} | {s['browser']} | {s['cookie_count']} cookies",
                font=ctk.CTkFont(size=12)
            )
            info.pack(side="left", padx=10, pady=8)
            
            del_btn = ctk.CTkButton(
                frame,
                text="🗑️",
                width=30,
                height=26,
                corner_radius=13,
                fg_color=AppleColors.RED,
                command=lambda idx=i: self.delete_session(idx)
            )
            del_btn.pack(side="right", padx=8, pady=5)
    
    def backup_session(self):
        if self.selected_profile_idx is None:
            messagebox.showwarning("Warning", "Select a profile first")
            return
        
        browser = self.browser_combo.get()
        profile = self.profiles_data[self.selected_profile_idx]
        
        if self.browser_helper and self.browser_helper.is_browser_running(browser):
            if messagebox.askyesno("Warning", f"Close {browser}?"):
                self.browser_helper.close_browser_gracefully(browser)
            else:
                return
        
        self.log(f"💾 Backing up...")
        
        def do_backup():
            try:
                result = self.session_manager.backup_session(browser, profile['path'])
                if result:
                    self.after(0, lambda: self.log(f"✅ Saved: {result}"))
                    self.after(0, lambda: messagebox.showinfo("Success", f"Backup saved!\n\n{result}"))
                    self.after(0, self.refresh_sessions)
            except Exception as e:
                self.after(0, lambda: self.log(f"❌ {e}"))
        
        threading.Thread(target=do_backup, daemon=True).start()
    
    def restore_session(self):
        messagebox.showinfo("Info", "Select a session from the list")
    
    def delete_session(self, idx):
        s = self.sessions_data[idx]
        if messagebox.askyesno("Confirm", f"Delete '{s['name']}'?"):
            self.session_manager.delete_session(s['name'])
            self.log(f"🗑️ Deleted: {s['name']}")
            self.refresh_sessions()
    
    # Browser Helper
    def detect_browsers(self):
        if not self.browser_helper:
            return
        
        for w in self.browser_list.winfo_children():
            w.destroy()
        
        browsers = self.browser_helper.detect_installed_browsers()
        
        for browser in browsers:
            profiles = self.browser_helper.get_browser_profiles_with_email(browser)
            
            label = ctk.CTkLabel(
                self.browser_list,
                text=f"🌐 {browser.capitalize()} ({len(profiles)} profiles)",
                font=ctk.CTkFont(size=14, weight="bold"),
                text_color=AppleColors.BLUE
            )
            label.pack(pady=(10, 5), anchor="w")
            
            for p in profiles:
                email = p.get('email') or 'No email'
                item = ctk.CTkLabel(
                    self.browser_list,
                    text=f"  📁 {p['name']} - {email}",
                    font=ctk.CTkFont(size=12)
                )
                item.pack(anchor="w", pady=2)
        
        self.log(f"🔍 Found {len(browsers)} browsers")
    
    def clean_traces(self):
        if not self.browser_helper:
            return
        
        if not messagebox.askyesno("Confirm", "Clean all browser traces?"):
            return
        
        def do_clean():
            try:
                for b in self.browser_helper.detect_installed_browsers():
                    stats = self.browser_helper.clean_browser_completely(b)
                    self.after(0, lambda b=b, s=stats: self.log(f"✅ {b}: {s['cookies']} cleaned"))
            except Exception as e:
                self.after(0, lambda: self.log(f"❌ {e}"))
        
        threading.Thread(target=do_clean, daemon=True).start()
    
    # Network
    def net_output(self, text):
        self.network_text.configure(state="normal")
        self.network_text.insert("end", text + "\n")
        self.network_text.see("end")
        self.network_text.configure(state="disabled")
    
    def flush_dns(self):
        import subprocess
        
        self.network_text.configure(state="normal")
        self.network_text.delete("1.0", "end")
        self.network_text.configure(state="disabled")
        
        cmd = "ipconfig /flushdns" if platform.system() == "Windows" else "dscacheutil -flushcache"
        
        try:
            result = subprocess.run(cmd, shell=True, capture_output=True, text=True)
            self.net_output("✅ DNS cache flushed!")
        except Exception as e:
            self.net_output(f"❌ {e}")
    
    def run_diagnostics(self):
        if not self.network_optimizer:
            return
        
        self.network_text.configure(state="normal")
        self.network_text.delete("1.0", "end")
        self.network_text.configure(state="disabled")
        
        def do_diag():
            try:
                report = self.network_optimizer.generate_diagnostic_report()
                self.after(0, lambda: self.net_output(report))
            except Exception as e:
                self.after(0, lambda: self.net_output(f"❌ {e}"))
        
        threading.Thread(target=do_diag, daemon=True).start()
    
    def reset_network(self):
        if platform.system() != "Windows":
            messagebox.showinfo("Info", "Windows only")
            return
        
        if not messagebox.askyesno("Warning", "Reset network? Restart required."):
            return
        
        import subprocess
        for cmd in ["netsh winsock reset", "netsh int ip reset"]:
            subprocess.run(cmd, shell=True, capture_output=True)
            self.net_output(f"✅ {cmd}")
        
        self.net_output("\n⚠️ Restart your computer!")


# ==================== Main ====================

def main():
    app = AntigravityApp()
    app.mainloop()


if __name__ == "__main__":
    main()
