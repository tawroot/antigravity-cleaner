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
import ctypes
from datetime import datetime
import shutil

try:
    import customtkinter as ctk
    from tkinter import messagebox
    import tkinter as tk
except ImportError:
    print("Missing customtkinter. Install: pip install customtkinter")
    sys.exit(1)

# Add src to path for imports
sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))

# Import shared modules
from config import APP_NAME, VERSION, GITHUB_URL, AppleColors
from i18n import get_text
from logger import setup_logger
from utils import get_base_path, DependencyLoader

# Helper modules (loaded dynamically)
loader = DependencyLoader()
loader.load_all()

BrowserHelper = loader.browser_helper
NetworkOptimizer = loader.network_optimizer
SessionManager = loader.session_manager
AntigravityDetector = loader.detector
GoogleServicesChecker = loader.google_checker
open_google_test_window = loader.google_window

# ==================== Main App ====================

class AntigravityApp(ctk.CTk):
    """
    Antigravity Cleaner with Apple-style design
    """
    
    def __init__(self):
        super().__init__()
        
        # Settings
        self.lang = 'en'
        self.dark_mode = False
        self.is_busy = False
        
        # Logger
        self.logger = setup_logger('antigravity_apple', 'apple-gui.log')
        
        # Initialize helpers
        self.init_helpers()
        
        # Window setup - COMPACT SIZE
        self.title(f"{APP_NAME} v{VERSION}")
        
        # Enable DPI awareness for Windows to fix 'out of frame' issues
        if platform.system() == "Windows":
            try:
                ctypes.windll.shcore.SetProcessDpiAwareness(1)
            except:
                pass
        
        # Calculate center position
        window_w = 980
        window_h = 680
        screen_w = self.winfo_screenwidth()
        screen_h = self.winfo_screenheight()
        x = (screen_w - window_w) // 2
        y = (screen_h - window_h) // 2
        
        self.geometry(f"{window_w}x{window_h}+{x}+{y}")
        self.minsize(900, 600)
        # Removed maxsize to allow full resizing
        
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
        self.log(f"🚀 {APP_NAME} started")
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
        return get_text(key, self.lang)
    
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
            text=f"{APP_NAME} Pro v{VERSION}",
            font=ctk.CTkFont(family="Inter", size=12),
            text_color=AppleColors.SECONDARY_LABEL
        )
        self.app_title.pack(side="right", padx=20)

        # Sidebar setup (Floating look)
        self.sidebar = ctk.CTkFrame(
            self,
            width=200,
            corner_radius=0,
            fg_color=AppleColors.BG_PRIMARY # White sidebar
        )
        self.sidebar.grid(row=1, column=0, sticky="nsew")
        self.sidebar.grid_propagate(False)
        
        # Logo Area
        logo_frame = ctk.CTkFrame(self.sidebar, fg_color="transparent")
        logo_frame.pack(pady=(40, 30), padx=20, fill="x")
        
        logo = ctk.CTkLabel(
            logo_frame,
            text="🚀",
            font=ctk.CTkFont(size=32)
        )
        logo.pack(side="left")
        
        title_label = ctk.CTkLabel(
            logo_frame,
            text="Antigravity",
            font=ctk.CTkFont(family="SF Pro Display", size=18, weight="bold"),
            text_color=AppleColors.TEXT_PRIMARY
        )
        title_label.pack(side="left", padx=10)

        # Navigation
        self.nav_btns = {}
        nav_items = [
            ("dashboard", "📊", "Dashboard"),
            ("google", "🔍", "Google Test"),
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
                height=45,
                corner_radius=8,
                fg_color="transparent",
                text_color=AppleColors.TEXT_SECONDARY,
                hover_color=AppleColors.BG_SECONDARY,
                anchor="w",
                command=lambda k=key: self.show_page(k)
            )
            btn.pack(fill="x", padx=15, pady=4)
            self.nav_btns[key] = btn
        
        # Spacer
        spacer = ctk.CTkFrame(self.sidebar, fg_color="transparent")
        spacer.pack(fill="both", expand=True)
        
        # Bottom Branding and Version
        branding_frame = ctk.CTkFrame(self.sidebar, fg_color="transparent")
        branding_frame.pack(side="bottom", fill="x", pady=(0, 20), padx=15)
        
        tawana = ctk.CTkLabel(
            branding_frame,
            text="TAWANA NETWORK",
            font=ctk.CTkFont(size=10, weight="bold"),
            text_color=AppleColors.ACCENT_PRIMARY
        )
        tawana.pack()
        
        ver = ctk.CTkLabel(
            branding_frame,
            text=f"v{VERSION}",
            font=ctk.CTkFont(size=10),
            text_color=AppleColors.TEXT_TERTIARY
        )
        ver.pack()

        license_btn = ctk.CTkLabel(
            branding_frame,
            text="License",
            font=ctk.CTkFont(size=9),
            text_color=AppleColors.TEXT_TERTIARY,
            cursor="hand2"
        )
        license_btn.pack(pady=(2, 0))
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
            command=lambda: webbrowser.open(GITHUB_URL)
        )
        github_btn.pack(side="right", padx=5)
    
    def show_page(self, page):
        """Switch pages"""
        self.current_page = page
        
        # Update nav buttons
        for key, btn in self.nav_btns.items():
            # Refresh text language
            icon = btn.cget("text").split()[0]
            # Simple text mapping based on translation not perfect but works for refresh
            # In a real scenario we'd rebuild the button or update text variable
            
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
            "google": self.page_google,
        }
        pages[page].grid(row=0, column=0, sticky="nsew", padx=20, pady=20)

    
    def show_license(self):
        """Show proprietary license information"""
        from config import LICENSE_TEXT
        messagebox.showinfo("License & Copyright", LICENSE_TEXT)
    
    def toggle_theme(self):

        """Toggle dark/light mode"""
        self.dark_mode = not self.dark_mode
        ctk.set_appearance_mode("dark" if self.dark_mode else "light")
    
    def toggle_lang(self):
        """Toggle language"""
        self.lang = 'fa' if self.lang == 'en' else 'en'
        self.lang_btn.configure(text="🌐 EN" if self.lang == 'fa' else "🌐 FA")
        # Note: A full UI refresh would be needed to update all texts immediately
        # For now, this just toggles state for subsequent renders or specific updates
        messagebox.showinfo("Language", "Language changed! Some elements will update on refresh.")
    
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
        self.create_google_page()

    
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
        """Minimal Dashboard Hub"""
        self.page_dashboard = ctk.CTkFrame(self.content, fg_color="transparent")
        self.page_dashboard.grid_columnconfigure(0, weight=1)
        self.page_dashboard.grid_rowconfigure(0, weight=1)
        
        # Center Hub
        hub_frame = ctk.CTkFrame(self.page_dashboard, fg_color="transparent")
        hub_frame.grid(row=0, column=0, sticky="nsew", padx=40, pady=40)
        
        # 1. Health Display (Big & Clean)
        self.health_score_label = ctk.CTkLabel(
            hub_frame,
            text="--",
            font=ctk.CTkFont(family="SF Pro Display", size=96, weight="bold"),
            text_color=AppleColors.TEXT_PRIMARY
        )
        self.health_score_label.pack(pady=(20, 0))
        
        self.health_status_label = ctk.CTkLabel(
            hub_frame,
            text="System Scanning...",
            font=ctk.CTkFont(size=18),
            text_color=AppleColors.TEXT_SECONDARY
        )
        self.health_status_label.pack(pady=(5, 30))
        
        # 2. Primary Actions (Horizontal, Large)
        action_frame = ctk.CTkFrame(hub_frame, fg_color="transparent")
        action_frame.pack(pady=20)
        
        # Scan Report Button (Safe)
        scan_btn = ctk.CTkButton(
            action_frame,
            text="Analyze Now",
            font=ctk.CTkFont(size=16, weight="bold"),
            height=50,
            width=200,
            corner_radius=25,
            fg_color=AppleColors.ACCENT_PRIMARY,
            hover_color=AppleColors.ACCENT_HOVER,
            command=self.run_scan
        )
        scan_btn.pack(side="left", padx=10)

        # Google Test Button
        google_btn = ctk.CTkButton(
            action_frame,
            text="Test Connectivity",
            font=ctk.CTkFont(size=16, weight="bold"),
            height=50,
            width=200,
            corner_radius=25,
            fg_color="transparent",
            border_width=2,
            border_color=AppleColors.ACCENT_PRIMARY,
            text_color=AppleColors.ACCENT_PRIMARY,
            hover_color=AppleColors.BG_TERTIARY,
            command=lambda: self.show_page('google')
        )
        google_btn.pack(side="left", padx=10)
        
        # 3. Quick Stats (Minimal Text)
        stats_frame = ctk.CTkFrame(hub_frame, fg_color="transparent")
        stats_frame.pack(pady=(40, 0), fill="x")
        
        self.status_labels = {}
        
        # Define stats to track
        stats_keys = ['browsers', 'sessions', 'leftovers']
        
        # Create 3 columns for stats
        stats_frame.grid_columnconfigure((0, 1, 2), weight=1)
        
        for i, key in enumerate(stats_keys):
            frame = ctk.CTkFrame(stats_frame, fg_color=AppleColors.BG_PRIMARY, corner_radius=12)
            frame.grid(row=0, column=i, padx=10, sticky="ew")
            
            val = ctk.CTkLabel(
                frame, 
                text="...", 
                font=ctk.CTkFont(size=24, weight="bold"),
                text_color=AppleColors.TEXT_PRIMARY
            )
            val.pack(pady=(15, 5))
            self.status_labels[key] = val
            
            lbl = ctk.CTkLabel(
                frame,
                text=key.capitalize(),
                font=ctk.CTkFont(size=12),
                text_color=AppleColors.TEXT_SECONDARY
            )
            lbl.pack(pady=(0, 15))

        # Add dummy labels for missing keys to prevent errors
        self.status_labels['installed'] = ctk.CTkLabel(self, text="")
        self.status_labels['running'] = ctk.CTkLabel(self, text="")

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
                try:
                    results = self.detector.detect_all()
                    score = results['health_score']
                    status = results['status']
                    
                    # Determine color based on score
                    if score >= 90:
                        color = AppleColors.ACCENT_SUCCESS
                    elif score >= 70:
                        color = AppleColors.ORANGE
                    else:
                        color = AppleColors.ACCENT_DANGER
                    
                    self.after(0, lambda: self.health_score_label.configure(text=str(score), text_color=color))
                    self.after(0, lambda: self.health_status_label.configure(text=status))
                    
                    # Update status items (Minimal)
                    leftovers_txt = f"{len(results['leftover_paths'])} Files"
                    if results['leftover_size'] > 0:
                         leftovers_txt += f"\n({results['leftover_size_human']})"
                    
                    self.after(0, lambda: self.status_labels['leftovers'].configure(text=leftovers_txt))
                except Exception as e:
                    self.after(0, lambda: self.health_status_label.configure(text=f"Scan Error"))

            else:
                self.after(0, lambda: self.health_score_label.configure(text="100", text_color=AppleColors.ACCENT_SUCCESS))
                self.after(0, lambda: self.health_status_label.configure(text="✅ Clean"))
            
            # Update browser count
            if self.browser_helper:
                try:
                    browsers = self.browser_helper.detect_installed_browsers()
                    self.after(0, lambda: self.status_labels['browsers'].configure(text=str(len(browsers))))
                except: pass
            
            # Update sessions count
            if self.session_manager:
                try:
                    sessions = self.session_manager.list_saved_sessions()
                    self.after(0, lambda: self.status_labels['sessions'].configure(text=str(len(sessions))))
                except: pass
        
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
        
        google_btn = self.create_button(
            btn_frame, self.t('test_google'), "blue",
            command=self.test_google_services
        )
        google_btn.pack(side="left", padx=5)
        
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
    
    # ==================== Google Page ====================
    
    def create_google_page(self):
        """Dedicated Google Services Test page"""
        self.page_google = ctk.CTkFrame(self.content, fg_color="transparent")
        self.page_google.grid_columnconfigure(0, weight=1)
        self.page_google.grid_rowconfigure(1, weight=1)
        
        # Header
        header = ctk.CTkLabel(
            self.page_google,
            text="🔍 Google Services Diagnostics",
            font=ctk.CTkFont(size=24, weight="bold"),
            text_color=AppleColors.LABEL
        )
        header.grid(row=0, column=0, sticky="w", pady=(0, 20))
        
        # Main card
        main_card = self.create_card(self.page_google)
        main_card.grid(row=1, column=0, sticky="nsew")
        main_card.grid_columnconfigure(0, weight=1)
        main_card.grid_rowconfigure(2, weight=1)
        
        # Description
        desc = ctk.CTkLabel(
            main_card,
            text="Test connectivity to Google Account, Gemini AI, Cloud Services, and AI Studio.\nUseful for diagnosing access issues in restricted networks.",
            font=ctk.CTkFont(size=13),
            text_color=AppleColors.SECONDARY_LABEL,
            justify="left"
        )
        desc.pack(pady=(20, 15), padx=25, anchor="w")
        
        # Large Test Button
        test_btn = ctk.CTkButton(
            main_card,
            text="🚀 Run Google Services Test",
            command=self.test_google_services,
            height=60,
            font=ctk.CTkFont(size=18, weight="bold"),
            fg_color=AppleColors.BLUE,
            hover_color=self._lighten(AppleColors.BLUE),
            corner_radius=12
        )
        test_btn.pack(pady=(0, 20), padx=25, fill="x")
        
        # Results area
        results_label = ctk.CTkLabel(
            main_card,
            text="📋 Test Results",
            font=ctk.CTkFont(size=14, weight="bold"),
            text_color=AppleColors.LABEL
        )
        results_label.pack(pady=(10, 5), padx=25, anchor="w")
        
        self.google_results_text = ctk.CTkTextbox(
            main_card,
            font=ctk.CTkFont(family="SF Mono", size=11),
            fg_color=AppleColors.BG_SECONDARY,
            text_color=AppleColors.LABEL,
            corner_radius=8,
            state="disabled"
        )
        self.google_results_text.pack(fill="both", expand=True, padx=25, pady=(0, 25))
        
        # Initial message
        self.google_results_text.configure(state="normal")
        self.google_results_text.insert("end", "Click 'Run Google Services Test' to start diagnostics...\n\n")
        self.google_results_text.insert("end", "Services to be tested:\n")
        self.google_results_text.insert("end", "  • Google Account (accounts.google.com)\n")
        self.google_results_text.insert("end", "  • Gemini AI (gemini.google.com)\n")
        self.google_results_text.insert("end", "  • Google Cloud (cloud.google.com)\n")
        self.google_results_text.insert("end", "  • AI Studio (aistudio.google.com)\n")
        self.google_results_text.configure(state="disabled")
    
    # ==================== Actions ====================

    
    def log(self, msg):
        """Add to log"""
        self.log_text.configure(state="normal")
        time = datetime.now().strftime("%H:%M:%S")
        self.log_text.insert("end", f"[{time}] {msg}\n")
        self.log_text.see("end")
        self.log_text.configure(state="disabled")
        try:
            self.logger.info(msg)
        except: pass
    
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
