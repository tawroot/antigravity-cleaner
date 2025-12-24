"""
Antigravity Cleaner - Material Design 3 GUI
============================================

Premium Google-style interface with:
- Material Design 3 theme
- Smooth rounded corners
- Elegant animations
- Modern color palette
- Full feature set

License: MIT
"""

import sys
import os
import platform
import threading
import webbrowser
from datetime import datetime

# CustomTkinter for modern UI
try:
    import customtkinter as ctk
    from tkinter import messagebox
    import tkinter as tk
except ImportError:
    print("Missing customtkinter. Install: pip install customtkinter")
    sys.exit(1)

# Add src to path
sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))

try:
    from browser_helper import BrowserHelper
    from network_optimizer import NetworkOptimizer
    from session_manager import SessionManager
    import logging
except ImportError as e:
    print(f"Warning: Some modules not available: {e}")
    BrowserHelper = None
    NetworkOptimizer = None
    SessionManager = None


# ==================== Material Design 3 Colors ====================

class MaterialColors:
    """Google Material Design 3 Color Palette"""
    
    # Primary colors (Google Blue)
    PRIMARY = "#1a73e8"
    PRIMARY_CONTAINER = "#d3e3fd"
    ON_PRIMARY = "#ffffff"
    ON_PRIMARY_CONTAINER = "#041e49"
    
    # Secondary colors (Teal)
    SECONDARY = "#00897b"
    SECONDARY_CONTAINER = "#b2dfdb"
    ON_SECONDARY = "#ffffff"
    
    # Tertiary (Purple)
    TERTIARY = "#7c4dff"
    TERTIARY_CONTAINER = "#e8def8"
    
    # Error
    ERROR = "#dc3545"
    ERROR_CONTAINER = "#f8d7da"
    
    # Success
    SUCCESS = "#34a853"
    SUCCESS_CONTAINER = "#ceead6"
    
    # Warning
    WARNING = "#fbbc04"
    WARNING_CONTAINER = "#fef7e0"
    
    # Surface colors (Light theme)
    SURFACE = "#ffffff"
    SURFACE_VARIANT = "#f1f3f4"
    SURFACE_CONTAINER = "#f8f9fa"
    SURFACE_CONTAINER_HIGH = "#e8eaed"
    
    # Dark theme
    DARK_SURFACE = "#1f1f1f"
    DARK_SURFACE_VARIANT = "#2d2d2d"
    DARK_SURFACE_CONTAINER = "#252525"
    
    # Text
    ON_SURFACE = "#1f1f1f"
    ON_SURFACE_VARIANT = "#5f6368"
    ON_SURFACE_DARK = "#e3e3e3"
    
    # Outline
    OUTLINE = "#dadce0"
    OUTLINE_DARK = "#444444"


# ==================== Translations ====================

TRANSLATIONS = {
    'en': {
        'title': 'Antigravity Cleaner',
        'subtitle': 'Professional Browser & IDE Cleaning Tool',
        'tab_cleaner': 'Cleaner',
        'tab_session': 'Sessions',
        'tab_browser': 'Browser',
        'tab_network': 'Network',
        'quick_clean': 'Quick Clean',
        'deep_clean': 'Deep Clean',
        'network_reset': 'Network Reset',
        'full_repair': 'Full Repair',
        'dry_run': 'Preview Mode (No changes)',
        'ready': 'Ready',
        'backup': 'Backup Session',
        'restore': 'Restore Session',
        'detect': 'Detect Browsers',
        'clean_traces': 'Clean Traces',
        'flush_dns': 'Flush DNS',
        'diagnostics': 'Run Diagnostics',
        'reset_stack': 'Reset Network',
        'select_browser': 'Select Browser',
        'select_profile': 'Select Profile',
        'saved_sessions': 'Saved Sessions',
        'no_sessions': 'No saved sessions',
        'success': 'Success',
        'error': 'Error',
        'warning': 'Warning',
        'confirm': 'Confirm',
        'close_browser': 'Browser is running. Close it first?',
    },
    'fa': {
        'title': 'پاک‌کننده آنتی‌گرویتی',
        'subtitle': 'ابزار حرفه‌ای پاکسازی مرورگر',
        'tab_cleaner': 'پاکسازی',
        'tab_session': 'نشست‌ها',
        'tab_browser': 'مرورگر',
        'tab_network': 'شبکه',
        'quick_clean': 'پاکسازی سریع',
        'deep_clean': 'پاکسازی عمیق',
        'network_reset': 'ریست شبکه',
        'full_repair': 'تعمیر کامل',
        'dry_run': 'حالت پیش‌نمایش',
        'ready': 'آماده',
        'backup': 'پشتیبان‌گیری',
        'restore': 'بازیابی',
        'detect': 'شناسایی مرورگرها',
        'clean_traces': 'پاکسازی ردها',
        'flush_dns': 'پاکسازی DNS',
        'diagnostics': 'تشخیص مشکلات',
        'reset_stack': 'ریست شبکه',
        'select_browser': 'انتخاب مرورگر',
        'select_profile': 'انتخاب پروفایل',
        'saved_sessions': 'نشست‌های ذخیره‌شده',
        'no_sessions': 'نشستی ذخیره نشده',
        'success': 'موفقیت',
        'error': 'خطا',
        'warning': 'هشدار',
        'confirm': 'تأیید',
        'close_browser': 'مرورگر باز است. ببندم؟',
    }
}


# ==================== Logger ====================

def setup_logger():
    """Setup logging"""
    log_dir = os.path.join(os.path.expanduser('~'), '.antigravity-cleaner', 'logs')
    os.makedirs(log_dir, exist_ok=True)
    
    logger = logging.getLogger('antigravity_material')
    logger.setLevel(logging.DEBUG)
    
    if not logger.handlers:
        handler = logging.FileHandler(
            os.path.join(log_dir, 'material-gui.log'),
            encoding='utf-8'
        )
        handler.setFormatter(logging.Formatter(
            '[%(asctime)s] %(levelname)s: %(message)s',
            datefmt='%Y-%m-%d %H:%M:%S'
        ))
        logger.addHandler(handler)
    
    return logger


# ==================== Custom Widgets ====================

class MaterialCard(ctk.CTkFrame):
    """Material Design Card Component"""
    
    def __init__(self, master, **kwargs):
        super().__init__(
            master,
            corner_radius=16,
            fg_color=MaterialColors.SURFACE,
            border_width=1,
            border_color=MaterialColors.OUTLINE,
            **kwargs
        )


class MaterialButton(ctk.CTkButton):
    """Material Design 3 Button"""
    
    def __init__(self, master, text, color="primary", icon=None, **kwargs):
        colors = {
            "primary": (MaterialColors.PRIMARY, MaterialColors.ON_PRIMARY),
            "secondary": (MaterialColors.SECONDARY, MaterialColors.ON_SECONDARY),
            "tertiary": (MaterialColors.TERTIARY, "#ffffff"),
            "error": (MaterialColors.ERROR, "#ffffff"),
            "success": (MaterialColors.SUCCESS, "#ffffff"),
            "warning": (MaterialColors.WARNING, "#000000"),
            "surface": (MaterialColors.SURFACE_CONTAINER_HIGH, MaterialColors.ON_SURFACE),
        }
        
        bg, fg = colors.get(color, colors["primary"])
        
        display_text = f"{icon} {text}" if icon else text
        
        super().__init__(
            master,
            text=display_text,
            corner_radius=20,
            height=44,
            font=ctk.CTkFont(family="Segoe UI", size=14, weight="bold"),
            fg_color=bg,
            text_color=fg,
            hover_color=self._darken(bg),
            **kwargs
        )
    
    def _darken(self, hex_color, factor=0.85):
        """Darken a hex color"""
        hex_color = hex_color.lstrip('#')
        rgb = tuple(int(hex_color[i:i+2], 16) for i in (0, 2, 4))
        darker = tuple(int(c * factor) for c in rgb)
        return f"#{darker[0]:02x}{darker[1]:02x}{darker[2]:02x}"


class MaterialIconButton(ctk.CTkButton):
    """Icon-only button"""
    
    def __init__(self, master, icon, tooltip=None, **kwargs):
        super().__init__(
            master,
            text=icon,
            width=44,
            height=44,
            corner_radius=22,
            font=ctk.CTkFont(size=18),
            fg_color="transparent",
            hover_color=MaterialColors.SURFACE_CONTAINER_HIGH,
            text_color=MaterialColors.ON_SURFACE_VARIANT,
            **kwargs
        )


# ==================== Main Application ====================

class AntigravityMaterialApp(ctk.CTk):
    """
    Antigravity Cleaner with Google Material Design 3
    """
    
    GITHUB_URL = "https://github.com/tawroot/antigravity-cleaner"
    VERSION = "3.0"
    
    def __init__(self):
        super().__init__()
        
        # Settings
        self.current_lang = 'en'
        self.is_busy = False
        self.dark_mode = False
        
        # Logger
        self.logger = setup_logger()
        self.logger.info("=== Material Design GUI Started ===")
        
        # Initialize helpers
        self.init_helpers()
        
        # Configure window
        self.title(f"Antigravity Cleaner v{self.VERSION}")
        self.geometry("1200x800")
        self.minsize(1000, 700)
        
        # Set appearance
        ctk.set_appearance_mode("light")
        ctk.set_default_color_theme("blue")
        
        # Configure grid
        self.grid_columnconfigure(1, weight=1)
        self.grid_rowconfigure(0, weight=1)
        
        # Create UI
        self.create_sidebar()
        self.create_main_content()
        
        # Center window
        self.center_window()
        
        # Initial log
        self.log_message(f"🚀 Antigravity Cleaner v{self.VERSION}")
        self.log_message(f"💻 {platform.system()} {platform.release()}")
        self.log_message(f"🐍 Python {platform.python_version()}")
        self.log_message("")
    
    def init_helpers(self):
        """Initialize helper modules"""
        self.browser_helper = None
        self.network_optimizer = None
        self.session_manager = None
        
        if BrowserHelper:
            try:
                self.browser_helper = BrowserHelper(self.logger, dry_run=False)
            except Exception as e:
                self.logger.error(f"BrowserHelper error: {e}")
        
        if NetworkOptimizer:
            try:
                self.network_optimizer = NetworkOptimizer(self.logger, dry_run=False)
            except Exception as e:
                self.logger.error(f"NetworkOptimizer error: {e}")
        
        if SessionManager:
            try:
                storage = os.path.join(os.path.expanduser('~'), '.antigravity-cleaner', 'sessions')
                self.session_manager = SessionManager(storage, self.logger, dry_run=False)
            except Exception as e:
                self.logger.error(f"SessionManager error: {e}")
    
    def t(self, key):
        """Get translation"""
        return TRANSLATIONS[self.current_lang].get(key, key)
    
    def center_window(self):
        """Center window on screen"""
        self.update_idletasks()
        w = self.winfo_width()
        h = self.winfo_height()
        x = (self.winfo_screenwidth() // 2) - (w // 2)
        y = (self.winfo_screenheight() // 2) - (h // 2)
        self.geometry(f'{w}x{h}+{x}+{y}')
    
    # ==================== Sidebar ====================
    
    def create_sidebar(self):
        """Create Material Design sidebar navigation"""
        self.sidebar = ctk.CTkFrame(
            self,
            width=280,
            corner_radius=0,
            fg_color=MaterialColors.SURFACE
        )
        self.sidebar.grid(row=0, column=0, sticky="nsew")
        self.sidebar.grid_rowconfigure(10, weight=1)
        
        # Logo/Title section
        title_frame = ctk.CTkFrame(self.sidebar, fg_color="transparent")
        title_frame.grid(row=0, column=0, padx=20, pady=(30, 10), sticky="ew")
        
        # App icon (emoji as placeholder)
        icon_label = ctk.CTkLabel(
            title_frame,
            text="🚀",
            font=ctk.CTkFont(size=40)
        )
        icon_label.pack(anchor="w")
        
        # Title
        title = ctk.CTkLabel(
            title_frame,
            text="Antigravity",
            font=ctk.CTkFont(family="Segoe UI", size=24, weight="bold"),
            text_color=MaterialColors.PRIMARY
        )
        title.pack(anchor="w", pady=(5, 0))
        
        subtitle = ctk.CTkLabel(
            title_frame,
            text="Cleaner",
            font=ctk.CTkFont(family="Segoe UI", size=24),
            text_color=MaterialColors.ON_SURFACE_VARIANT
        )
        subtitle.pack(anchor="w")
        
        # Separator
        sep = ctk.CTkFrame(self.sidebar, height=1, fg_color=MaterialColors.OUTLINE)
        sep.grid(row=1, column=0, sticky="ew", padx=20, pady=20)
        
        # Navigation buttons
        self.nav_buttons = {}
        nav_items = [
            ("cleaner", "🧹", self.t('tab_cleaner')),
            ("session", "💾", self.t('tab_session')),
            ("browser", "🌐", self.t('tab_browser')),
            ("network", "🔧", self.t('tab_network')),
        ]
        
        for i, (key, icon, text) in enumerate(nav_items):
            btn = ctk.CTkButton(
                self.sidebar,
                text=f"  {icon}  {text}",
                font=ctk.CTkFont(family="Segoe UI", size=15),
                height=50,
                corner_radius=25,
                fg_color="transparent",
                text_color=MaterialColors.ON_SURFACE,
                hover_color=MaterialColors.PRIMARY_CONTAINER,
                anchor="w",
                command=lambda k=key: self.show_page(k)
            )
            btn.grid(row=i+2, column=0, padx=15, pady=3, sticky="ew")
            self.nav_buttons[key] = btn
        
        # Bottom section
        bottom_frame = ctk.CTkFrame(self.sidebar, fg_color="transparent")
        bottom_frame.grid(row=11, column=0, padx=20, pady=20, sticky="ew")
        
        # Theme toggle
        self.theme_btn = ctk.CTkButton(
            bottom_frame,
            text="🌙 Dark Mode",
            font=ctk.CTkFont(size=13),
            height=40,
            corner_radius=20,
            fg_color=MaterialColors.SURFACE_CONTAINER_HIGH,
            text_color=MaterialColors.ON_SURFACE_VARIANT,
            hover_color=MaterialColors.SURFACE_VARIANT,
            command=self.toggle_theme
        )
        self.theme_btn.pack(fill="x", pady=5)
        
        # Language toggle
        self.lang_btn = ctk.CTkButton(
            bottom_frame,
            text="🌐 فارسی",
            font=ctk.CTkFont(size=13),
            height=40,
            corner_radius=20,
            fg_color=MaterialColors.SURFACE_CONTAINER_HIGH,
            text_color=MaterialColors.ON_SURFACE_VARIANT,
            hover_color=MaterialColors.SURFACE_VARIANT,
            command=self.toggle_language
        )
        self.lang_btn.pack(fill="x", pady=5)
        
        # GitHub button
        github_btn = ctk.CTkButton(
            bottom_frame,
            text="⭐ Star on GitHub",
            font=ctk.CTkFont(size=13, weight="bold"),
            height=40,
            corner_radius=20,
            fg_color=MaterialColors.TERTIARY,
            hover_color="#6a3de8",
            command=lambda: webbrowser.open(self.GITHUB_URL)
        )
        github_btn.pack(fill="x", pady=5)
        
        # Set initial active
        self.active_page = "cleaner"
        self.update_nav_active()
    
    def update_nav_active(self):
        """Update navigation button states"""
        for key, btn in self.nav_buttons.items():
            if key == self.active_page:
                btn.configure(
                    fg_color=MaterialColors.PRIMARY_CONTAINER,
                    text_color=MaterialColors.PRIMARY
                )
            else:
                btn.configure(
                    fg_color="transparent",
                    text_color=MaterialColors.ON_SURFACE
                )
    
    def show_page(self, page):
        """Switch to a page"""
        self.active_page = page
        self.update_nav_active()
        
        # Hide all pages
        for child in self.main_content.winfo_children():
            child.grid_forget()
        
        # Show selected page
        pages = {
            "cleaner": self.page_cleaner,
            "session": self.page_session,
            "browser": self.page_browser,
            "network": self.page_network,
        }
        pages[page].grid(row=0, column=0, sticky="nsew", padx=30, pady=30)
    
    # ==================== Main Content ====================
    
    def create_main_content(self):
        """Create main content area"""
        self.main_content = ctk.CTkFrame(
            self,
            fg_color=MaterialColors.SURFACE_CONTAINER,
            corner_radius=0
        )
        self.main_content.grid(row=0, column=1, sticky="nsew")
        self.main_content.grid_columnconfigure(0, weight=1)
        self.main_content.grid_rowconfigure(0, weight=1)
        
        # Create pages
        self.create_cleaner_page()
        self.create_session_page()
        self.create_browser_page()
        self.create_network_page()
        
        # Show default page
        self.show_page("cleaner")
    
    def create_cleaner_page(self):
        """Create Cleaner page"""
        self.page_cleaner = ctk.CTkFrame(self.main_content, fg_color="transparent")
        self.page_cleaner.grid_columnconfigure(1, weight=1)
        self.page_cleaner.grid_rowconfigure(1, weight=1)
        
        # Header
        header = ctk.CTkLabel(
            self.page_cleaner,
            text="🧹 Cleaner",
            font=ctk.CTkFont(family="Segoe UI", size=28, weight="bold"),
            text_color=MaterialColors.ON_SURFACE
        )
        header.grid(row=0, column=0, columnspan=2, sticky="w", pady=(0, 20))
        
        # Left panel - Actions
        actions_card = MaterialCard(self.page_cleaner)
        actions_card.grid(row=1, column=0, sticky="nsew", padx=(0, 20))
        
        # Preview mode switch
        self.preview_var = ctk.BooleanVar(value=False)
        preview_switch = ctk.CTkSwitch(
            actions_card,
            text=self.t('dry_run'),
            font=ctk.CTkFont(size=14),
            variable=self.preview_var,
            progress_color=MaterialColors.PRIMARY
        )
        preview_switch.pack(padx=25, pady=20, anchor="w")
        
        # Action buttons
        actions = [
            (self.t('quick_clean'), "success", lambda: self.run_clean('quick')),
            (self.t('deep_clean'), "primary", lambda: self.run_clean('deep')),
            (self.t('network_reset'), "warning", lambda: self.run_clean('network')),
            (self.t('full_repair'), "error", lambda: self.run_clean('full')),
        ]
        
        for text, color, cmd in actions:
            btn = MaterialButton(actions_card, text, color, command=cmd)
            btn.pack(padx=25, pady=8, fill="x")
        
        # Right panel - Log
        log_card = MaterialCard(self.page_cleaner)
        log_card.grid(row=1, column=1, sticky="nsew")
        log_card.grid_columnconfigure(0, weight=1)
        log_card.grid_rowconfigure(1, weight=1)
        
        # Log header
        log_header = ctk.CTkFrame(log_card, fg_color="transparent")
        log_header.grid(row=0, column=0, sticky="ew", padx=20, pady=15)
        
        log_title = ctk.CTkLabel(
            log_header,
            text="📋 Operation Log",
            font=ctk.CTkFont(size=16, weight="bold"),
            text_color=MaterialColors.ON_SURFACE
        )
        log_title.pack(side="left")
        
        clear_btn = ctk.CTkButton(
            log_header,
            text="Clear",
            width=70,
            height=32,
            corner_radius=16,
            font=ctk.CTkFont(size=12),
            fg_color=MaterialColors.SURFACE_CONTAINER_HIGH,
            text_color=MaterialColors.ON_SURFACE_VARIANT,
            hover_color=MaterialColors.SURFACE_VARIANT,
            command=self.clear_log
        )
        clear_btn.pack(side="right")
        
        # Log textbox
        self.log_textbox = ctk.CTkTextbox(
            log_card,
            font=ctk.CTkFont(family="Consolas", size=12),
            fg_color=MaterialColors.SURFACE_VARIANT,
            text_color=MaterialColors.ON_SURFACE,
            corner_radius=12,
            state="disabled"
        )
        self.log_textbox.grid(row=1, column=0, sticky="nsew", padx=20, pady=(0, 20))
    
    def create_session_page(self):
        """Create Session Manager page"""
        self.page_session = ctk.CTkFrame(self.main_content, fg_color="transparent")
        self.page_session.grid_columnconfigure(1, weight=1)
        self.page_session.grid_rowconfigure(1, weight=1)
        
        # Header
        header = ctk.CTkLabel(
            self.page_session,
            text="💾 Session Manager",
            font=ctk.CTkFont(family="Segoe UI", size=28, weight="bold"),
            text_color=MaterialColors.ON_SURFACE
        )
        header.grid(row=0, column=0, columnspan=2, sticky="w", pady=(0, 20))
        
        # Left panel
        left_card = MaterialCard(self.page_session, width=320)
        left_card.grid(row=1, column=0, sticky="nsew", padx=(0, 20))
        
        # Browser selection
        browser_label = ctk.CTkLabel(
            left_card,
            text=self.t('select_browser'),
            font=ctk.CTkFont(size=14, weight="bold"),
            text_color=MaterialColors.ON_SURFACE
        )
        browser_label.pack(padx=20, pady=(20, 10), anchor="w")
        
        self.browser_combo = ctk.CTkComboBox(
            left_card,
            width=260,
            height=40,
            corner_radius=12,
            font=ctk.CTkFont(size=13),
            dropdown_font=ctk.CTkFont(size=13),
            command=self.on_browser_select
        )
        self.browser_combo.pack(padx=20, pady=(0, 15))
        
        # Profile selection
        profile_label = ctk.CTkLabel(
            left_card,
            text=self.t('select_profile'),
            font=ctk.CTkFont(size=14, weight="bold"),
            text_color=MaterialColors.ON_SURFACE
        )
        profile_label.pack(padx=20, pady=(10, 10), anchor="w")
        
        # Profile scrollable frame
        self.profile_frame = ctk.CTkScrollableFrame(
            left_card,
            width=240,
            height=200,
            corner_radius=12,
            fg_color=MaterialColors.SURFACE_VARIANT
        )
        self.profile_frame.pack(padx=20, pady=(0, 20), fill="both", expand=True)
        
        # Session buttons
        backup_btn = MaterialButton(
            left_card, self.t('backup'), "success", "💾",
            command=self.backup_session
        )
        backup_btn.pack(padx=20, pady=8, fill="x")
        
        restore_btn = MaterialButton(
            left_card, self.t('restore'), "primary", "📥",
            command=self.restore_session
        )
        restore_btn.pack(padx=20, pady=(8, 20), fill="x")
        
        # Right panel - Sessions list
        right_card = MaterialCard(self.page_session)
        right_card.grid(row=1, column=1, sticky="nsew")
        right_card.grid_columnconfigure(0, weight=1)
        right_card.grid_rowconfigure(1, weight=1)
        
        # Sessions header
        sessions_header = ctk.CTkFrame(right_card, fg_color="transparent")
        sessions_header.grid(row=0, column=0, sticky="ew", padx=20, pady=15)
        
        sessions_title = ctk.CTkLabel(
            sessions_header,
            text=f"📁 {self.t('saved_sessions')}",
            font=ctk.CTkFont(size=16, weight="bold"),
            text_color=MaterialColors.ON_SURFACE
        )
        sessions_title.pack(side="left")
        
        refresh_btn = ctk.CTkButton(
            sessions_header,
            text="🔄 Refresh",
            width=90,
            height=32,
            corner_radius=16,
            font=ctk.CTkFont(size=12),
            fg_color=MaterialColors.PRIMARY,
            command=self.refresh_sessions
        )
        refresh_btn.pack(side="right")
        
        # Sessions scrollable frame
        self.sessions_frame = ctk.CTkScrollableFrame(
            right_card,
            corner_radius=12,
            fg_color=MaterialColors.SURFACE_VARIANT
        )
        self.sessions_frame.grid(row=1, column=0, sticky="nsew", padx=20, pady=(0, 20))
        
        # Load data after UI is ready
        self.after(200, self.load_session_data)
    
    def create_browser_page(self):
        """Create Browser Helper page"""
        self.page_browser = ctk.CTkFrame(self.main_content, fg_color="transparent")
        self.page_browser.grid_columnconfigure(0, weight=1)
        self.page_browser.grid_rowconfigure(1, weight=1)
        
        # Header
        header = ctk.CTkLabel(
            self.page_browser,
            text="🌐 Browser Helper",
            font=ctk.CTkFont(family="Segoe UI", size=28, weight="bold"),
            text_color=MaterialColors.ON_SURFACE
        )
        header.grid(row=0, column=0, sticky="w", pady=(0, 20))
        
        # Action buttons
        btn_frame = ctk.CTkFrame(self.page_browser, fg_color="transparent")
        btn_frame.grid(row=0, column=0, sticky="e", pady=(0, 20))
        
        detect_btn = MaterialButton(
            btn_frame, self.t('detect'), "primary", "🔍",
            command=self.detect_browsers
        )
        detect_btn.pack(side="left", padx=5)
        
        clean_btn = MaterialButton(
            btn_frame, self.t('clean_traces'), "success", "🧹",
            command=self.clean_traces
        )
        clean_btn.pack(side="left", padx=5)
        
        # Browser profiles card
        profiles_card = MaterialCard(self.page_browser)
        profiles_card.grid(row=1, column=0, sticky="nsew")
        profiles_card.grid_columnconfigure(0, weight=1)
        profiles_card.grid_rowconfigure(0, weight=1)
        
        self.browser_profiles_frame = ctk.CTkScrollableFrame(
            profiles_card,
            corner_radius=12,
            fg_color=MaterialColors.SURFACE_VARIANT
        )
        self.browser_profiles_frame.pack(fill="both", expand=True, padx=20, pady=20)
    
    def create_network_page(self):
        """Create Network page"""
        self.page_network = ctk.CTkFrame(self.main_content, fg_color="transparent")
        self.page_network.grid_columnconfigure(0, weight=1)
        self.page_network.grid_rowconfigure(1, weight=1)
        
        # Header
        header = ctk.CTkLabel(
            self.page_network,
            text="🔧 Network Tools",
            font=ctk.CTkFont(family="Segoe UI", size=28, weight="bold"),
            text_color=MaterialColors.ON_SURFACE
        )
        header.grid(row=0, column=0, sticky="w", pady=(0, 20))
        
        # Action buttons
        btn_frame = ctk.CTkFrame(self.page_network, fg_color="transparent")
        btn_frame.grid(row=0, column=0, sticky="e", pady=(0, 20))
        
        dns_btn = MaterialButton(
            btn_frame, self.t('flush_dns'), "success", "🧹",
            command=self.flush_dns
        )
        dns_btn.pack(side="left", padx=5)
        
        diag_btn = MaterialButton(
            btn_frame, self.t('diagnostics'), "primary", "🔍",
            command=self.run_diagnostics
        )
        diag_btn.pack(side="left", padx=5)
        
        reset_btn = MaterialButton(
            btn_frame, self.t('reset_stack'), "error", "⚠️",
            command=self.reset_network
        )
        reset_btn.pack(side="left", padx=5)
        
        # Output card
        output_card = MaterialCard(self.page_network)
        output_card.grid(row=1, column=0, sticky="nsew")
        output_card.grid_columnconfigure(0, weight=1)
        output_card.grid_rowconfigure(0, weight=1)
        
        self.network_textbox = ctk.CTkTextbox(
            output_card,
            font=ctk.CTkFont(family="Consolas", size=12),
            fg_color=MaterialColors.SURFACE_VARIANT,
            text_color=MaterialColors.ON_SURFACE,
            corner_radius=12,
            state="disabled"
        )
        self.network_textbox.pack(fill="both", expand=True, padx=20, pady=20)
    
    # ==================== Actions ====================
    
    def toggle_theme(self):
        """Toggle dark/light theme"""
        self.dark_mode = not self.dark_mode
        if self.dark_mode:
            ctk.set_appearance_mode("dark")
            self.theme_btn.configure(text="☀️ Light Mode")
        else:
            ctk.set_appearance_mode("light")
            self.theme_btn.configure(text="🌙 Dark Mode")
    
    def toggle_language(self):
        """Toggle language"""
        self.current_lang = 'fa' if self.current_lang == 'en' else 'en'
        self.lang_btn.configure(text='🌐 English' if self.current_lang == 'fa' else '🌐 فارسی')
        self.log_message(f"Language: {'Persian' if self.current_lang == 'fa' else 'English'}")
    
    def log_message(self, message, level="info"):
        """Add message to log"""
        self.log_textbox.configure(state="normal")
        timestamp = datetime.now().strftime("%H:%M:%S")
        self.log_textbox.insert("end", f"[{timestamp}] {message}\n")
        self.log_textbox.see("end")
        self.log_textbox.configure(state="disabled")
        self.logger.info(message)
    
    def clear_log(self):
        """Clear log"""
        self.log_textbox.configure(state="normal")
        self.log_textbox.delete("1.0", "end")
        self.log_textbox.configure(state="disabled")
    
    def run_clean(self, mode):
        """Run cleaning operation"""
        if self.is_busy:
            return
        
        mode_names = {
            'quick': self.t('quick_clean'),
            'deep': self.t('deep_clean'),
            'network': self.t('network_reset'),
            'full': self.t('full_repair')
        }
        
        if not messagebox.askyesno(self.t('confirm'), f"Start {mode_names[mode]}?"):
            return
        
        self.is_busy = True
        self.log_message(f"🚀 Starting {mode_names[mode]}...")
        
        def do_clean():
            try:
                dry_run = self.preview_var.get()
                
                if mode in ['quick', 'deep', 'full']:
                    self._clean_files(mode, dry_run)
                
                if mode in ['network', 'full']:
                    self._reset_network_cmd(dry_run)
                
                self.after(0, lambda: self.log_message("✅ Operation completed!"))
                self.after(0, lambda: messagebox.showinfo(self.t('success'), "Operation completed!"))
            except Exception as e:
                self.after(0, lambda: self.log_message(f"❌ Error: {e}"))
            finally:
                self.is_busy = False
        
        threading.Thread(target=do_clean, daemon=True).start()
    
    def _clean_files(self, mode, dry_run):
        """Clean Antigravity files"""
        import shutil
        import glob
        
        home = os.path.expanduser("~")
        paths = []
        
        if platform.system() == "Windows":
            local = os.environ.get("LOCALAPPDATA", "")
            roaming = os.environ.get("APPDATA", "")
            paths = [
                os.path.join(local, "Programs", "Antigravity"),
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
                os.path.join(home, ".local", "share", "Antigravity"),
                os.path.join(home, ".cache", "Antigravity"),
            ]
        
        found = 0
        for path in paths:
            if os.path.exists(path):
                found += 1
                if dry_run:
                    self.after(0, lambda p=path: self.log_message(f"[Preview] Would remove: {p}"))
                else:
                    try:
                        if os.path.isdir(path):
                            shutil.rmtree(path)
                        else:
                            os.remove(path)
                        self.after(0, lambda p=path: self.log_message(f"🗑️ Removed: {p}"))
                    except Exception as e:
                        self.after(0, lambda e=e: self.log_message(f"⚠️ Error: {e}"))
        
        if found == 0:
            self.after(0, lambda: self.log_message("ℹ️ No Antigravity files found"))
    
    def _reset_network_cmd(self, dry_run):
        """Reset network"""
        import subprocess
        
        if platform.system() == "Windows":
            cmds = ["ipconfig /flushdns", "netsh winsock reset"]
        elif platform.system() == "Darwin":
            cmds = ["dscacheutil -flushcache"]
        else:
            cmds = ["resolvectl flush-caches"]
        
        for cmd in cmds:
            if dry_run:
                self.after(0, lambda c=cmd: self.log_message(f"[Preview] Would run: {c}"))
            else:
                try:
                    subprocess.run(cmd, shell=True, capture_output=True)
                    self.after(0, lambda c=cmd: self.log_message(f"✅ {c}"))
                except Exception as e:
                    self.after(0, lambda e=e: self.log_message(f"❌ {e}"))
    
    # ==================== Session Manager ====================
    
    def load_session_data(self):
        """Load browser and session data"""
        if self.browser_helper:
            browsers = self.browser_helper.detect_installed_browsers()
            self.browser_combo.configure(values=browsers)
            if browsers:
                self.browser_combo.set(browsers[0])
                self.on_browser_select(browsers[0])
        
        self.refresh_sessions()
    
    def on_browser_select(self, browser):
        """Handle browser selection"""
        if not self.browser_helper:
            return
        
        # Clear profile frame
        for widget in self.profile_frame.winfo_children():
            widget.destroy()
        
        profiles = self.browser_helper.get_browser_profiles_with_email(browser)
        self.profiles_data = profiles
        self.selected_profile = None
        
        for i, p in enumerate(profiles):
            email = p.get('email') or 'No email'
            
            profile_btn = ctk.CTkRadioButton(
                self.profile_frame,
                text=f"{p['name']}\n{email}",
                font=ctk.CTkFont(size=12),
                variable=ctk.IntVar(),
                value=i,
                command=lambda idx=i: self.select_profile(idx)
            )
            profile_btn.pack(pady=5, padx=10, anchor="w")
    
    def select_profile(self, idx):
        """Select a profile"""
        self.selected_profile = self.profiles_data[idx]
    
    def refresh_sessions(self):
        """Refresh sessions list"""
        # Clear frame
        for widget in self.sessions_frame.winfo_children():
            widget.destroy()
        
        if not self.session_manager:
            return
        
        sessions = self.session_manager.list_saved_sessions()
        
        if not sessions:
            label = ctk.CTkLabel(
                self.sessions_frame,
                text=self.t('no_sessions'),
                font=ctk.CTkFont(size=14),
                text_color=MaterialColors.ON_SURFACE_VARIANT
            )
            label.pack(pady=30)
            return
        
        self.sessions_data = sessions
        
        for i, s in enumerate(sessions):
            session_frame = ctk.CTkFrame(
                self.sessions_frame,
                fg_color=MaterialColors.SURFACE,
                corner_radius=12
            )
            session_frame.pack(fill="x", pady=5, padx=5)
            
            status = "🟢" if not s.get('expired') else "🔴"
            
            info = ctk.CTkLabel(
                session_frame,
                text=f"{status} {s['name']} | {s['browser']} | {s['cookie_count']} cookies",
                font=ctk.CTkFont(size=13),
                text_color=MaterialColors.ON_SURFACE
            )
            info.pack(side="left", padx=15, pady=12)
            
            del_btn = ctk.CTkButton(
                session_frame,
                text="🗑️",
                width=40,
                height=32,
                corner_radius=16,
                fg_color=MaterialColors.ERROR,
                hover_color="#c82333",
                command=lambda idx=i: self.delete_session(idx)
            )
            del_btn.pack(side="right", padx=10, pady=8)
    
    def backup_session(self):
        """Backup browser session"""
        if not self.selected_profile or not self.session_manager:
            messagebox.showwarning(self.t('warning'), "Select a profile first")
            return
        
        browser = self.browser_combo.get()
        
        # Check if browser running
        if self.browser_helper and self.browser_helper.is_browser_running(browser):
            if messagebox.askyesno(self.t('warning'), self.t('close_browser')):
                self.browser_helper.close_browser_gracefully(browser)
            else:
                return
        
        self.log_message(f"💾 Backing up {browser} session...")
        
        def do_backup():
            try:
                result = self.session_manager.backup_session(browser, self.selected_profile['path'])
                if result:
                    self.after(0, lambda: self.log_message(f"✅ Backup saved: {result}"))
                    self.after(0, lambda: messagebox.showinfo(self.t('success'), f"Session backed up!\n\n{result}"))
                    self.after(0, self.refresh_sessions)
            except Exception as e:
                self.after(0, lambda: self.log_message(f"❌ Backup failed: {e}"))
        
        threading.Thread(target=do_backup, daemon=True).start()
    
    def restore_session(self):
        """Restore a session"""
        messagebox.showinfo("Info", "Select a session from the list and use the restore function")
    
    def delete_session(self, idx):
        """Delete a session"""
        session = self.sessions_data[idx]
        if messagebox.askyesno(self.t('confirm'), f"Delete session '{session['name']}'?"):
            self.session_manager.delete_session(session['name'])
            self.log_message(f"🗑️ Deleted: {session['name']}")
            self.refresh_sessions()
    
    # ==================== Browser Helper ====================
    
    def detect_browsers(self):
        """Detect installed browsers"""
        if not self.browser_helper:
            return
        
        # Clear frame
        for widget in self.browser_profiles_frame.winfo_children():
            widget.destroy()
        
        browsers = self.browser_helper.detect_installed_browsers()
        
        for browser in browsers:
            profiles = self.browser_helper.get_browser_profiles_with_email(browser)
            
            # Browser header
            browser_label = ctk.CTkLabel(
                self.browser_profiles_frame,
                text=f"🌐 {browser.capitalize()} ({len(profiles)} profiles)",
                font=ctk.CTkFont(size=16, weight="bold"),
                text_color=MaterialColors.PRIMARY
            )
            browser_label.pack(pady=(15, 5), anchor="w")
            
            for p in profiles:
                profile_frame = ctk.CTkFrame(
                    self.browser_profiles_frame,
                    fg_color=MaterialColors.SURFACE,
                    corner_radius=10
                )
                profile_frame.pack(fill="x", pady=3)
                
                email = p.get('email') or 'No email'
                info = ctk.CTkLabel(
                    profile_frame,
                    text=f"  📁 {p['name']} - {email}",
                    font=ctk.CTkFont(size=13),
                    text_color=MaterialColors.ON_SURFACE
                )
                info.pack(side="left", padx=10, pady=10)
        
        self.log_message(f"🔍 Found {len(browsers)} browsers")
    
    def clean_traces(self):
        """Clean browser traces"""
        if not self.browser_helper:
            return
        
        if not messagebox.askyesno(self.t('confirm'), "Clean Antigravity traces from all browsers?"):
            return
        
        self.log_message("🧹 Cleaning browser traces...")
        
        def do_clean():
            try:
                browsers = self.browser_helper.detect_installed_browsers()
                total = 0
                for browser in browsers:
                    stats = self.browser_helper.clean_browser_completely(browser)
                    total += stats['cookies']
                    self.after(0, lambda b=browser, s=stats: 
                        self.log_message(f"✅ {b}: {s['cookies']} cookies cleaned"))
                
                self.after(0, lambda: messagebox.showinfo(self.t('success'), f"Cleaned {total} items"))
            except Exception as e:
                self.after(0, lambda: self.log_message(f"❌ Error: {e}"))
        
        threading.Thread(target=do_clean, daemon=True).start()
    
    # ==================== Network ====================
    
    def network_output(self, text):
        """Write to network output"""
        self.network_textbox.configure(state="normal")
        self.network_textbox.insert("end", text + "\n")
        self.network_textbox.see("end")
        self.network_textbox.configure(state="disabled")
    
    def flush_dns(self):
        """Flush DNS"""
        import subprocess
        
        self.network_textbox.configure(state="normal")
        self.network_textbox.delete("1.0", "end")
        self.network_textbox.configure(state="disabled")
        
        self.network_output("🧹 Flushing DNS cache...")
        
        if platform.system() == "Windows":
            cmd = "ipconfig /flushdns"
        elif platform.system() == "Darwin":
            cmd = "dscacheutil -flushcache"
        else:
            cmd = "resolvectl flush-caches"
        
        try:
            result = subprocess.run(cmd, shell=True, capture_output=True, text=True)
            self.network_output(result.stdout or "✅ DNS cache flushed!")
        except Exception as e:
            self.network_output(f"❌ Error: {e}")
    
    def run_diagnostics(self):
        """Run network diagnostics"""
        if not self.network_optimizer:
            return
        
        self.network_textbox.configure(state="normal")
        self.network_textbox.delete("1.0", "end")
        self.network_textbox.configure(state="disabled")
        
        self.network_output("🔍 Running diagnostics...")
        
        def do_diag():
            try:
                report = self.network_optimizer.generate_diagnostic_report()
                self.after(0, lambda: self.network_output(report))
            except Exception as e:
                self.after(0, lambda: self.network_output(f"❌ Error: {e}"))
        
        threading.Thread(target=do_diag, daemon=True).start()
    
    def reset_network(self):
        """Reset network stack"""
        if platform.system() != "Windows":
            messagebox.showinfo("Info", "Network reset only available on Windows")
            return
        
        if not messagebox.askyesno(self.t('warning'), "Reset network stack? Restart required after."):
            return
        
        import subprocess
        
        self.network_output("⚠️ Resetting network stack...")
        
        for cmd in ["netsh winsock reset", "netsh int ip reset"]:
            try:
                subprocess.run(cmd, shell=True, capture_output=True)
                self.network_output(f"✅ {cmd}")
            except Exception as e:
                self.network_output(f"❌ {e}")
        
        self.network_output("\n⚠️ Please restart your computer!")


# ==================== Main ====================

def main():
    app = AntigravityMaterialApp()
    app.mainloop()


if __name__ == "__main__":
    main()
