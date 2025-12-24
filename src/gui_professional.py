"""
Antigravity Cleaner - Professional GUI Version 3.0
===================================================

Full-featured graphical interface with:
- Real cleaning operations
- Session Manager (Backup/Restore)
- Browser Helper
- Network Optimizer
- Modern dark theme UI

License: MIT
"""

import sys
import os
import platform
import threading
import webbrowser
import json
from pathlib import Path
from datetime import datetime

try:
    import tkinter as tk
    from tkinter import ttk, messagebox, scrolledtext, filedialog
except ImportError:
    print("Error: tkinter is not installed!")
    sys.exit(1)

# Add src to path for imports
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

try:
    import psutil
except ImportError:
    psutil = None


# ==================== Translations ====================

TRANSLATIONS = {
    'en': {
        'title': 'Antigravity Cleaner',
        'subtitle': 'Professional Browser & IDE Cleaning Tool',
        'tab_cleaner': '🧹 Cleaner',
        'tab_session': '💾 Session Manager',
        'tab_browser': '🌐 Browser Helper',
        'tab_network': '🔧 Network',
        'tab_settings': '⚙️ Settings',
        'os': 'Operating System',
        'python': 'Python Version',
        'options': 'Cleaning Options',
        'dry_run': 'Dry Run (Preview only, no changes)',
        'quick_clean': 'Quick Clean',
        'deep_clean': 'Deep Clean',
        'network_reset': 'Network Reset',
        'full_repair': 'Full Repair',
        'ready': 'Ready...',
        'log_title': 'Operation Log',
        'footer': '© 2025 Tawroot | Open Source under MIT License',
        'github_btn': '⭐ Star on GitHub',
        'language': 'Language',
        'confirm_title': 'Confirm Action',
        'confirm_msg': 'Start {mode}?',
        'warning': 'Warning',
        'in_progress': 'Operation in progress!',
        'success': 'Success',
        'success_msg': 'Operation completed successfully!',
        'error': 'Error',
        'running': 'Running {mode}...',
        'scanning': 'Scanning...',
        'completed': 'Completed!',
        # Session Manager
        'session_backup': 'Backup Session',
        'session_restore': 'Restore Session',
        'session_list': 'Saved Sessions',
        'session_delete': 'Delete Selected',
        'select_browser': 'Select Browser',
        'select_profile': 'Select Profile',
        'no_sessions': 'No saved sessions',
        'backup_success': 'Session backed up successfully!',
        'restore_success': 'Session restored successfully!',
        # Browser Helper
        'clean_traces': 'Clean Browser Traces',
        'detect_browsers': 'Detect Browsers',
        'profiles_found': 'Profiles Found',
        # Network
        'dns_flush': 'Flush DNS Cache',
        'network_diag': 'Run Diagnostics',
        'reset_stack': 'Reset Network Stack',
    },
    'fa': {
        'title': 'پاک‌کننده آنتی‌گرویتی',
        'subtitle': 'ابزار حرفه‌ای پاکسازی مرورگر و IDE',
        'tab_cleaner': '🧹 پاکسازی',
        'tab_session': '💾 مدیریت نشست',
        'tab_browser': '🌐 کمک‌کننده مرورگر',
        'tab_network': '🔧 شبکه',
        'tab_settings': '⚙️ تنظیمات',
        'os': 'سیستم عامل',
        'python': 'نسخه پایتون',
        'options': 'گزینه‌ها',
        'dry_run': 'حالت آزمایشی (بدون تغییر)',
        'quick_clean': 'پاکسازی سریع',
        'deep_clean': 'پاکسازی عمیق',
        'network_reset': 'ریست شبکه',
        'full_repair': 'تعمیر کامل',
        'ready': 'آماده...',
        'log_title': 'گزارش عملیات',
        'footer': '© ۱۴۰۴ تاوروت | متن‌باز تحت مجوز MIT',
        'github_btn': '⭐ ستاره در گیت‌هاب',
        'language': 'زبان',
        'confirm_title': 'تأیید',
        'confirm_msg': 'شروع {mode}؟',
        'warning': 'هشدار',
        'in_progress': 'عملیات در حال انجام!',
        'success': 'موفقیت',
        'success_msg': 'عملیات با موفقیت انجام شد!',
        'error': 'خطا',
        'running': 'در حال اجرای {mode}...',
        'scanning': 'در حال اسکن...',
        'completed': 'تکمیل شد!',
        # Session Manager
        'session_backup': 'پشتیبان‌گیری نشست',
        'session_restore': 'بازیابی نشست',
        'session_list': 'نشست‌های ذخیره‌شده',
        'session_delete': 'حذف انتخاب‌شده',
        'select_browser': 'انتخاب مرورگر',
        'select_profile': 'انتخاب پروفایل',
        'no_sessions': 'نشستی ذخیره نشده',
        'backup_success': 'نشست با موفقیت ذخیره شد!',
        'restore_success': 'نشست با موفقیت بازیابی شد!',
        # Browser Helper
        'clean_traces': 'پاکسازی ردهای مرورگر',
        'detect_browsers': 'شناسایی مرورگرها',
        'profiles_found': 'پروفایل‌های یافت‌شده',
        # Network
        'dns_flush': 'پاکسازی DNS',
        'network_diag': 'تشخیص مشکلات',
        'reset_stack': 'ریست پشته شبکه',
    }
}


# ==================== Theme Colors ====================

THEME = {
    'bg_dark': '#0d1117',
    'bg_card': '#161b22',
    'bg_hover': '#21262d',
    'accent_blue': '#58a6ff',
    'accent_green': '#3fb950',
    'accent_red': '#f85149',
    'accent_yellow': '#d29922',
    'accent_purple': '#a371f7',
    'accent_cyan': '#39c5cf',
    'text_primary': '#f0f6fc',
    'text_secondary': '#8b949e',
    'border': '#30363d',
}


# ==================== Logger Setup ====================

def setup_gui_logger():
    """Setup logging for GUI operations"""
    log_dir = os.path.join(os.path.expanduser('~'), '.antigravity-cleaner', 'logs')
    os.makedirs(log_dir, exist_ok=True)
    
    log_file = os.path.join(log_dir, 'gui-operations.log')
    
    logger = logging.getLogger('antigravity_gui')
    logger.setLevel(logging.DEBUG)
    
    if not logger.handlers:
        handler = logging.FileHandler(log_file, encoding='utf-8')
        handler.setFormatter(logging.Formatter(
            '[%(asctime)s] [%(levelname)s] %(message)s',
            datefmt='%Y-%m-%d %H:%M:%S'
        ))
        logger.addHandler(handler)
    
    return logger


# ==================== Main Application ====================

class AntigravityCleanerGUI:
    """
    Professional GUI for Antigravity Cleaner with all features.
    """
    
    GITHUB_URL = "https://github.com/tawroot/antigravity-cleaner"
    VERSION = "3.0"
    
    def __init__(self, root):
        self.root = root
        self.current_lang = 'en'
        self.is_busy = False
        self.dry_run = tk.BooleanVar(value=False)
        
        # Setup logger
        self.logger = setup_gui_logger()
        self.logger.info("=== Antigravity Cleaner GUI Started ===")
        
        # Initialize helpers
        self.init_helpers()
        
        # Setup UI
        self.setup_window()
        self.setup_styles()
        self.create_ui()
        
        # Center window
        self.center_window()
        
        # Initial log
        self.log(f"Antigravity Cleaner v{self.VERSION} started")
        self.log(f"OS: {platform.system()} {platform.release()}")
        self.log(f"Python: {platform.python_version()}")
        self.log("")
    
    def init_helpers(self):
        """Initialize helper modules"""
        self.browser_helper = None
        self.network_optimizer = None
        self.session_manager = None
        
        if BrowserHelper:
            try:
                self.browser_helper = BrowserHelper(self.logger, dry_run=False)
                self.logger.info("BrowserHelper initialized")
            except Exception as e:
                self.logger.error(f"BrowserHelper init failed: {e}")
        
        if NetworkOptimizer:
            try:
                self.network_optimizer = NetworkOptimizer(self.logger, dry_run=False)
                self.logger.info("NetworkOptimizer initialized")
            except Exception as e:
                self.logger.error(f"NetworkOptimizer init failed: {e}")
        
        if SessionManager:
            try:
                storage = os.path.join(os.path.expanduser('~'), '.antigravity-cleaner', 'sessions')
                self.session_manager = SessionManager(storage, self.logger, dry_run=False)
                self.logger.info("SessionManager initialized")
            except Exception as e:
                self.logger.error(f"SessionManager init failed: {e}")
    
    def t(self, key):
        """Get translation"""
        return TRANSLATIONS[self.current_lang].get(key, key)
    
    def setup_window(self):
        """Configure main window"""
        self.root.title(f"Antigravity Cleaner v{self.VERSION}")
        self.root.geometry("1100x750")
        self.root.minsize(900, 600)
        self.root.configure(bg=THEME['bg_dark'])
        
        # Set icon if exists
        icon_path = os.path.join(os.path.dirname(__file__), '..', 'icon.ico')
        if os.path.exists(icon_path):
            try:
                self.root.iconbitmap(icon_path)
            except:
                pass
    
    def setup_styles(self):
        """Configure ttk styles"""
        style = ttk.Style()
        style.theme_use('clam')
        
        # Notebook (Tabs)
        style.configure('TNotebook', background=THEME['bg_dark'], borderwidth=0)
        style.configure('TNotebook.Tab', 
                       background=THEME['bg_card'],
                       foreground=THEME['text_secondary'],
                       padding=[20, 10],
                       font=('Segoe UI', 10, 'bold'))
        style.map('TNotebook.Tab',
                 background=[('selected', THEME['bg_hover'])],
                 foreground=[('selected', THEME['accent_blue'])])
        
        # Progress bar
        style.configure('Accent.Horizontal.TProgressbar',
                       troughcolor=THEME['bg_card'],
                       background=THEME['accent_blue'],
                       bordercolor=THEME['border'])
        
        # Treeview
        style.configure('Treeview',
                       background=THEME['bg_card'],
                       foreground=THEME['text_primary'],
                       fieldbackground=THEME['bg_card'],
                       borderwidth=0,
                       font=('Segoe UI', 10))
        style.configure('Treeview.Heading',
                       background=THEME['bg_hover'],
                       foreground=THEME['text_secondary'],
                       font=('Segoe UI', 10, 'bold'))
        style.map('Treeview',
                 background=[('selected', THEME['accent_blue'])],
                 foreground=[('selected', THEME['bg_dark'])])
    
    def center_window(self):
        """Center window on screen"""
        self.root.update_idletasks()
        w = self.root.winfo_width()
        h = self.root.winfo_height()
        x = (self.root.winfo_screenwidth() // 2) - (w // 2)
        y = (self.root.winfo_screenheight() // 2) - (h // 2)
        self.root.geometry(f'{w}x{h}+{x}+{y}')
    
    def create_button(self, parent, text, command, color='blue', width=15):
        """Create a styled button"""
        colors = {
            'blue': (THEME['accent_blue'], '#4090e0'),
            'green': (THEME['accent_green'], '#35a045'),
            'red': (THEME['accent_red'], '#e04545'),
            'yellow': (THEME['accent_yellow'], '#c08820'),
            'purple': (THEME['accent_purple'], '#9060e0'),
            'cyan': (THEME['accent_cyan'], '#30b0c0'),
        }
        bg, active = colors.get(color, colors['blue'])
        
        btn = tk.Button(
            parent,
            text=text,
            command=command,
            font=('Segoe UI', 10, 'bold'),
            bg=bg,
            fg='#ffffff',
            activebackground=active,
            activeforeground='#ffffff',
            bd=0,
            padx=15,
            pady=10,
            cursor='hand2',
            relief=tk.FLAT,
            width=width
        )
        return btn
    
    def create_ui(self):
        """Create main UI"""
        # Header
        self.create_header()
        
        # Main content with tabs
        self.create_tabs()
        
        # Status bar
        self.create_statusbar()
    
    def create_header(self):
        """Create header section"""
        header = tk.Frame(self.root, bg=THEME['bg_card'], height=80)
        header.pack(fill=tk.X)
        header.pack_propagate(False)
        
        # Left side - Title
        left = tk.Frame(header, bg=THEME['bg_card'])
        left.pack(side=tk.LEFT, padx=20, pady=15)
        
        title = tk.Label(
            left,
            text=self.t('title'),
            font=('Segoe UI', 22, 'bold'),
            bg=THEME['bg_card'],
            fg=THEME['accent_blue']
        )
        title.pack(anchor='w')
        
        subtitle = tk.Label(
            left,
            text=self.t('subtitle'),
            font=('Segoe UI', 10),
            bg=THEME['bg_card'],
            fg=THEME['text_secondary']
        )
        subtitle.pack(anchor='w')
        
        # Right side - Buttons
        right = tk.Frame(header, bg=THEME['bg_card'])
        right.pack(side=tk.RIGHT, padx=20, pady=20)
        
        # Language toggle
        self.lang_btn = tk.Button(
            right,
            text='🌐 فارسی',
            command=self.toggle_language,
            font=('Segoe UI', 9),
            bg=THEME['bg_hover'],
            fg=THEME['text_primary'],
            bd=0,
            padx=10,
            pady=5,
            cursor='hand2'
        )
        self.lang_btn.pack(side=tk.LEFT, padx=5)
        
        # GitHub button
        github_btn = tk.Button(
            right,
            text='⭐ GitHub',
            command=lambda: webbrowser.open(self.GITHUB_URL),
            font=('Segoe UI', 9, 'bold'),
            bg=THEME['accent_purple'],
            fg='#ffffff',
            bd=0,
            padx=15,
            pady=5,
            cursor='hand2'
        )
        github_btn.pack(side=tk.LEFT, padx=5)
    
    def create_tabs(self):
        """Create tabbed interface"""
        # Notebook container
        container = tk.Frame(self.root, bg=THEME['bg_dark'])
        container.pack(fill=tk.BOTH, expand=True, padx=15, pady=10)
        
        self.notebook = ttk.Notebook(container)
        self.notebook.pack(fill=tk.BOTH, expand=True)
        
        # Tab 1: Cleaner
        self.tab_cleaner = tk.Frame(self.notebook, bg=THEME['bg_dark'])
        self.notebook.add(self.tab_cleaner, text=self.t('tab_cleaner'))
        self.create_cleaner_tab()
        
        # Tab 2: Session Manager
        self.tab_session = tk.Frame(self.notebook, bg=THEME['bg_dark'])
        self.notebook.add(self.tab_session, text=self.t('tab_session'))
        self.create_session_tab()
        
        # Tab 3: Browser Helper
        self.tab_browser = tk.Frame(self.notebook, bg=THEME['bg_dark'])
        self.notebook.add(self.tab_browser, text=self.t('tab_browser'))
        self.create_browser_tab()
        
        # Tab 4: Network
        self.tab_network = tk.Frame(self.notebook, bg=THEME['bg_dark'])
        self.notebook.add(self.tab_network, text=self.t('tab_network'))
        self.create_network_tab()
    
    def create_cleaner_tab(self):
        """Create Cleaner tab content"""
        # Left panel - Actions
        left = tk.Frame(self.tab_cleaner, bg=THEME['bg_card'], width=300)
        left.pack(side=tk.LEFT, fill=tk.Y, padx=(0, 10), pady=5)
        left.pack_propagate(False)
        
        # Options
        options_label = tk.Label(
            left,
            text=self.t('options'),
            font=('Segoe UI', 12, 'bold'),
            bg=THEME['bg_card'],
            fg=THEME['accent_cyan']
        )
        options_label.pack(pady=(15, 10), padx=15, anchor='w')
        
        # Dry Run checkbox
        self.dry_run_check = tk.Checkbutton(
            left,
            text=self.t('dry_run'),
            variable=self.dry_run,
            font=('Segoe UI', 10),
            bg=THEME['bg_card'],
            fg=THEME['text_primary'],
            selectcolor=THEME['bg_hover'],
            activebackground=THEME['bg_card'],
            activeforeground=THEME['accent_blue']
        )
        self.dry_run_check.pack(padx=15, anchor='w')
        
        # Separator
        sep = tk.Frame(left, bg=THEME['border'], height=1)
        sep.pack(fill=tk.X, padx=15, pady=15)
        
        # Action buttons
        actions_label = tk.Label(
            left,
            text="Actions",
            font=('Segoe UI', 12, 'bold'),
            bg=THEME['bg_card'],
            fg=THEME['accent_cyan']
        )
        actions_label.pack(pady=(5, 10), padx=15, anchor='w')
        
        # Quick Clean
        btn1 = self.create_button(left, self.t('quick_clean'), 
                                  lambda: self.run_clean('quick'), 'green', 20)
        btn1.pack(pady=5, padx=15, fill=tk.X)
        
        # Deep Clean
        btn2 = self.create_button(left, self.t('deep_clean'), 
                                  lambda: self.run_clean('deep'), 'blue', 20)
        btn2.pack(pady=5, padx=15, fill=tk.X)
        
        # Network Reset
        btn3 = self.create_button(left, self.t('network_reset'), 
                                  lambda: self.run_clean('network'), 'yellow', 20)
        btn3.pack(pady=5, padx=15, fill=tk.X)
        
        # Full Repair
        btn4 = self.create_button(left, self.t('full_repair'), 
                                  lambda: self.run_clean('full'), 'red', 20)
        btn4.pack(pady=5, padx=15, fill=tk.X)
        
        # Right panel - Log
        right = tk.Frame(self.tab_cleaner, bg=THEME['bg_dark'])
        right.pack(side=tk.RIGHT, fill=tk.BOTH, expand=True, pady=5)
        
        # Log header
        log_header = tk.Frame(right, bg=THEME['bg_card'])
        log_header.pack(fill=tk.X)
        
        log_label = tk.Label(
            log_header,
            text=self.t('log_title'),
            font=('Segoe UI', 11, 'bold'),
            bg=THEME['bg_card'],
            fg=THEME['accent_cyan'],
            pady=8,
            padx=10
        )
        log_label.pack(side=tk.LEFT)
        
        # Clear log button
        clear_btn = tk.Button(
            log_header,
            text='🗑️ Clear',
            command=self.clear_log,
            font=('Segoe UI', 9),
            bg=THEME['bg_hover'],
            fg=THEME['text_secondary'],
            bd=0,
            padx=10,
            pady=3,
            cursor='hand2'
        )
        clear_btn.pack(side=tk.RIGHT, padx=10, pady=5)
        
        # Log text
        self.log_text = scrolledtext.ScrolledText(
            right,
            font=('Consolas', 10),
            bg='#0d1117',
            fg='#58a6ff',
            insertbackground='#58a6ff',
            wrap=tk.WORD,
            state=tk.DISABLED,
            bd=0,
            padx=10,
            pady=10
        )
        self.log_text.pack(fill=tk.BOTH, expand=True)
        
        # Configure log tags for colors
        self.log_text.tag_configure('success', foreground=THEME['accent_green'])
        self.log_text.tag_configure('error', foreground=THEME['accent_red'])
        self.log_text.tag_configure('warning', foreground=THEME['accent_yellow'])
        self.log_text.tag_configure('info', foreground=THEME['accent_blue'])
    
    def create_session_tab(self):
        """Create Session Manager tab"""
        # Left panel - Actions
        left = tk.Frame(self.tab_session, bg=THEME['bg_card'], width=320)
        left.pack(side=tk.LEFT, fill=tk.Y, padx=(0, 10), pady=5)
        left.pack_propagate(False)
        
        # Browser selection
        browser_label = tk.Label(
            left,
            text=self.t('select_browser'),
            font=('Segoe UI', 11, 'bold'),
            bg=THEME['bg_card'],
            fg=THEME['accent_cyan']
        )
        browser_label.pack(pady=(15, 5), padx=15, anchor='w')
        
        self.browser_var = tk.StringVar()
        self.browser_combo = ttk.Combobox(
            left,
            textvariable=self.browser_var,
            state='readonly',
            font=('Segoe UI', 10)
        )
        self.browser_combo.pack(fill=tk.X, padx=15, pady=5)
        self.browser_combo.bind('<<ComboboxSelected>>', self.on_browser_selected)
        
        # Profile selection
        profile_label = tk.Label(
            left,
            text=self.t('select_profile'),
            font=('Segoe UI', 11, 'bold'),
            bg=THEME['bg_card'],
            fg=THEME['accent_cyan']
        )
        profile_label.pack(pady=(15, 5), padx=15, anchor='w')
        
        # Profile listbox
        self.profile_listbox = tk.Listbox(
            left,
            font=('Segoe UI', 10),
            bg=THEME['bg_hover'],
            fg=THEME['text_primary'],
            selectbackground=THEME['accent_blue'],
            selectforeground='#ffffff',
            bd=0,
            highlightthickness=0,
            height=8
        )
        self.profile_listbox.pack(fill=tk.X, padx=15, pady=5)
        
        # Separator
        sep = tk.Frame(left, bg=THEME['border'], height=1)
        sep.pack(fill=tk.X, padx=15, pady=15)
        
        # Backup button
        backup_btn = self.create_button(
            left, self.t('session_backup'),
            self.backup_session, 'green', 22
        )
        backup_btn.pack(pady=5, padx=15, fill=tk.X)
        
        # Restore button
        restore_btn = self.create_button(
            left, self.t('session_restore'),
            self.restore_session, 'blue', 22
        )
        restore_btn.pack(pady=5, padx=15, fill=tk.X)
        
        # Right panel - Sessions list
        right = tk.Frame(self.tab_session, bg=THEME['bg_dark'])
        right.pack(side=tk.RIGHT, fill=tk.BOTH, expand=True, pady=5)
        
        # Sessions header
        sessions_label = tk.Label(
            right,
            text=self.t('session_list'),
            font=('Segoe UI', 12, 'bold'),
            bg=THEME['bg_dark'],
            fg=THEME['accent_cyan']
        )
        sessions_label.pack(pady=(5, 10), anchor='w')
        
        # Sessions treeview
        columns = ('name', 'browser', 'date', 'cookies', 'status')
        self.sessions_tree = ttk.Treeview(right, columns=columns, show='headings', height=12)
        
        self.sessions_tree.heading('name', text='Session Name')
        self.sessions_tree.heading('browser', text='Browser')
        self.sessions_tree.heading('date', text='Backup Date')
        self.sessions_tree.heading('cookies', text='Cookies')
        self.sessions_tree.heading('status', text='Status')
        
        self.sessions_tree.column('name', width=180)
        self.sessions_tree.column('browser', width=100)
        self.sessions_tree.column('date', width=150)
        self.sessions_tree.column('cookies', width=80)
        self.sessions_tree.column('status', width=80)
        
        self.sessions_tree.pack(fill=tk.BOTH, expand=True)
        
        # Delete button
        delete_frame = tk.Frame(right, bg=THEME['bg_dark'])
        delete_frame.pack(fill=tk.X, pady=10)
        
        delete_btn = self.create_button(
            delete_frame, self.t('session_delete'),
            self.delete_session, 'red', 18
        )
        delete_btn.pack(side=tk.RIGHT)
        
        refresh_btn = self.create_button(
            delete_frame, '🔄 Refresh',
            self.refresh_sessions, 'cyan', 12
        )
        refresh_btn.pack(side=tk.RIGHT, padx=10)
        
        # Load initial data
        self.root.after(100, self.load_browsers)
        self.root.after(200, self.refresh_sessions)
    
    def create_browser_tab(self):
        """Create Browser Helper tab"""
        # Main content
        content = tk.Frame(self.tab_browser, bg=THEME['bg_dark'])
        content.pack(fill=tk.BOTH, expand=True, padx=20, pady=20)
        
        # Title
        title = tk.Label(
            content,
            text="🌐 Browser Helper",
            font=('Segoe UI', 16, 'bold'),
            bg=THEME['bg_dark'],
            fg=THEME['accent_cyan']
        )
        title.pack(pady=(0, 20))
        
        # Buttons row
        btn_frame = tk.Frame(content, bg=THEME['bg_dark'])
        btn_frame.pack(pady=10)
        
        detect_btn = self.create_button(
            btn_frame, self.t('detect_browsers'),
            self.detect_browsers, 'blue', 20
        )
        detect_btn.pack(side=tk.LEFT, padx=10)
        
        clean_btn = self.create_button(
            btn_frame, self.t('clean_traces'),
            self.clean_browser_traces, 'green', 20
        )
        clean_btn.pack(side=tk.LEFT, padx=10)
        
        # Results area
        results_label = tk.Label(
            content,
            text=self.t('profiles_found'),
            font=('Segoe UI', 12, 'bold'),
            bg=THEME['bg_dark'],
            fg=THEME['text_secondary']
        )
        results_label.pack(pady=(30, 10), anchor='w')
        
        # Browser profiles treeview
        columns = ('browser', 'profile', 'email', 'path')
        self.browser_tree = ttk.Treeview(content, columns=columns, show='headings', height=15)
        
        self.browser_tree.heading('browser', text='Browser')
        self.browser_tree.heading('profile', text='Profile')
        self.browser_tree.heading('email', text='Email')
        self.browser_tree.heading('path', text='Path')
        
        self.browser_tree.column('browser', width=120)
        self.browser_tree.column('profile', width=120)
        self.browser_tree.column('email', width=200)
        self.browser_tree.column('path', width=400)
        
        self.browser_tree.pack(fill=tk.BOTH, expand=True)
    
    def create_network_tab(self):
        """Create Network tab"""
        content = tk.Frame(self.tab_network, bg=THEME['bg_dark'])
        content.pack(fill=tk.BOTH, expand=True, padx=20, pady=20)
        
        # Title
        title = tk.Label(
            content,
            text="🔧 Network Tools",
            font=('Segoe UI', 16, 'bold'),
            bg=THEME['bg_dark'],
            fg=THEME['accent_cyan']
        )
        title.pack(pady=(0, 30))
        
        # Buttons
        btn_frame = tk.Frame(content, bg=THEME['bg_dark'])
        btn_frame.pack(pady=10)
        
        dns_btn = self.create_button(
            btn_frame, self.t('dns_flush'),
            self.flush_dns, 'green', 20
        )
        dns_btn.pack(side=tk.LEFT, padx=10)
        
        diag_btn = self.create_button(
            btn_frame, self.t('network_diag'),
            self.run_network_diagnostics, 'blue', 20
        )
        diag_btn.pack(side=tk.LEFT, padx=10)
        
        reset_btn = self.create_button(
            btn_frame, self.t('reset_stack'),
            self.reset_network_stack, 'red', 20
        )
        reset_btn.pack(side=tk.LEFT, padx=10)
        
        # Diagnostics output
        diag_label = tk.Label(
            content,
            text="Diagnostics Output",
            font=('Segoe UI', 12, 'bold'),
            bg=THEME['bg_dark'],
            fg=THEME['text_secondary']
        )
        diag_label.pack(pady=(30, 10), anchor='w')
        
        self.network_output = scrolledtext.ScrolledText(
            content,
            font=('Consolas', 10),
            bg=THEME['bg_card'],
            fg=THEME['text_primary'],
            wrap=tk.WORD,
            state=tk.DISABLED,
            height=15
        )
        self.network_output.pack(fill=tk.BOTH, expand=True)
    
    def create_statusbar(self):
        """Create status bar"""
        statusbar = tk.Frame(self.root, bg=THEME['bg_card'], height=35)
        statusbar.pack(fill=tk.X, side=tk.BOTTOM)
        statusbar.pack_propagate(False)
        
        self.status_label = tk.Label(
            statusbar,
            text=self.t('ready'),
            font=('Segoe UI', 9),
            bg=THEME['bg_card'],
            fg=THEME['text_secondary']
        )
        self.status_label.pack(side=tk.LEFT, padx=15, pady=8)
        
        # Progress bar
        self.progress = ttk.Progressbar(
            statusbar,
            mode='indeterminate',
            length=150,
            style='Accent.Horizontal.TProgressbar'
        )
        self.progress.pack(side=tk.RIGHT, padx=15, pady=8)
    
    # ==================== Logging ====================
    
    def log(self, message, tag=None):
        """Write message to log"""
        self.log_text.configure(state=tk.NORMAL)
        timestamp = datetime.now().strftime("%H:%M:%S")
        line = f"[{timestamp}] {message}\n"
        self.log_text.insert(tk.END, line, tag)
        self.log_text.see(tk.END)
        self.log_text.configure(state=tk.DISABLED)
        self.root.update_idletasks()
        
        # Also log to file
        self.logger.info(message)
    
    def clear_log(self):
        """Clear log text"""
        self.log_text.configure(state=tk.NORMAL)
        self.log_text.delete(1.0, tk.END)
        self.log_text.configure(state=tk.DISABLED)
    
    def set_status(self, text):
        """Update status bar"""
        self.status_label.config(text=text)
        self.root.update_idletasks()
    
    def start_progress(self):
        """Start progress animation"""
        self.progress.start(10)
        self.is_busy = True
    
    def stop_progress(self):
        """Stop progress animation"""
        self.progress.stop()
        self.is_busy = False
        self.set_status(self.t('ready'))
    
    # ==================== Language ====================
    
    def toggle_language(self):
        """Toggle between English and Persian"""
        self.current_lang = 'fa' if self.current_lang == 'en' else 'en'
        self.lang_btn.config(text='🌐 English' if self.current_lang == 'fa' else '🌐 فارسی')
        self.log(f"Language changed to {'Persian' if self.current_lang == 'fa' else 'English'}")
    
    # ==================== Cleaner Actions ====================
    
    def run_clean(self, mode):
        """Run cleaning operation"""
        if self.is_busy:
            messagebox.showwarning(self.t('warning'), self.t('in_progress'))
            return
        
        mode_names = {
            'quick': self.t('quick_clean'),
            'deep': self.t('deep_clean'),
            'network': self.t('network_reset'),
            'full': self.t('full_repair')
        }
        mode_name = mode_names.get(mode, mode)
        
        if not messagebox.askyesno(self.t('confirm_title'), self.t('confirm_msg').format(mode=mode_name)):
            return
        
        # Run in thread
        thread = threading.Thread(target=self._run_clean_thread, args=(mode, mode_name))
        thread.daemon = True
        thread.start()
    
    def _run_clean_thread(self, mode, mode_name):
        """Clean thread execution"""
        self.root.after(0, self.start_progress)
        self.root.after(0, lambda: self.set_status(self.t('running').format(mode=mode_name)))
        
        try:
            self.root.after(0, lambda: self.log(f"\n{'='*50}", 'info'))
            self.root.after(0, lambda: self.log(f">>> {mode_name.upper()}", 'info'))
            self.root.after(0, lambda: self.log(f"{'='*50}\n", 'info'))
            
            dry_run = self.dry_run.get()
            
            if mode in ['quick', 'deep', 'full']:
                self._clean_files(mode, dry_run)
            
            if mode in ['network', 'full']:
                self._reset_network(dry_run)
            
            self.root.after(0, lambda: self.log(f"\n✓ {self.t('completed')}", 'success'))
            self.root.after(0, lambda: messagebox.showinfo(self.t('success'), self.t('success_msg')))
            
        except Exception as e:
            self.root.after(0, lambda: self.log(f"✗ Error: {str(e)}", 'error'))
            self.root.after(0, lambda: messagebox.showerror(self.t('error'), str(e)))
        
        finally:
            self.root.after(0, self.stop_progress)
    
    def _clean_files(self, mode, dry_run):
        """Clean Antigravity files"""
        import shutil
        import glob
        
        home = os.path.expanduser("~")
        paths = []
        
        if platform.system() == "Windows":
            local = os.environ.get("LOCALAPPDATA", os.path.join(home, "AppData", "Local"))
            roaming = os.environ.get("APPDATA", os.path.join(home, "AppData", "Roaming"))
            
            paths = [
                os.path.join(local, "Programs", "Antigravity"),
                os.path.join(local, "Antigravity"),
                os.path.join(roaming, "Antigravity"),
            ]
            
            if mode in ['deep', 'full']:
                temp = os.environ.get("TEMP", os.path.join(local, "Temp"))
                paths.append(os.path.join(temp, "antigravity-*"))
        
        elif platform.system() == "Darwin":
            paths = [
                os.path.join(home, "Library", "Application Support", "Antigravity"),
                os.path.join(home, "Library", "Caches", "Antigravity"),
            ]
        
        else:  # Linux
            paths = [
                os.path.join(home, ".config", "Antigravity"),
                os.path.join(home, ".local", "share", "Antigravity"),
                os.path.join(home, ".cache", "Antigravity"),
            ]
        
        # Expand globs
        expanded = []
        for p in paths:
            if "*" in p:
                expanded.extend(glob.glob(p))
            else:
                expanded.append(p)
        
        # Clean
        found = 0
        for path in expanded:
            if os.path.exists(path):
                found += 1
                if dry_run:
                    self.root.after(0, lambda p=path: self.log(f"[DRY RUN] Would remove: {p}", 'warning'))
                else:
                    try:
                        if os.path.isdir(path):
                            shutil.rmtree(path)
                        else:
                            os.remove(path)
                        self.root.after(0, lambda p=path: self.log(f"Removed: {p}", 'success'))
                    except Exception as e:
                        self.root.after(0, lambda p=path, e=e: self.log(f"Failed: {p} - {e}", 'error'))
        
        if found == 0:
            self.root.after(0, lambda: self.log("No Antigravity files found", 'info'))
    
    def _reset_network(self, dry_run):
        """Reset network settings"""
        import subprocess
        
        commands = []
        if platform.system() == "Windows":
            commands = [
                "ipconfig /flushdns",
                "netsh winsock reset",
            ]
        elif platform.system() == "Darwin":
            commands = [
                "dscacheutil -flushcache",
                "killall -HUP mDNSResponder",
            ]
        else:
            commands = ["resolvectl flush-caches"]
        
        for cmd in commands:
            if dry_run:
                self.root.after(0, lambda c=cmd: self.log(f"[DRY RUN] Would run: {c}", 'warning'))
            else:
                self.root.after(0, lambda c=cmd: self.log(f"Running: {c}", 'info'))
                try:
                    subprocess.run(cmd, shell=True, capture_output=True)
                    self.root.after(0, lambda c=cmd: self.log(f"✓ {c} completed", 'success'))
                except Exception as e:
                    self.root.after(0, lambda e=e: self.log(f"Error: {e}", 'error'))
    
    # ==================== Session Manager ====================
    
    def load_browsers(self):
        """Load available browsers"""
        if not self.browser_helper:
            return
        
        browsers = self.browser_helper.detect_installed_browsers()
        self.browser_combo['values'] = browsers
        if browsers:
            self.browser_combo.current(0)
            self.on_browser_selected(None)
    
    def on_browser_selected(self, event):
        """Handle browser selection"""
        if not self.browser_helper:
            return
        
        browser = self.browser_var.get()
        if not browser:
            return
        
        profiles = self.browser_helper.get_browser_profiles_with_email(browser)
        
        self.profile_listbox.delete(0, tk.END)
        self.profiles_data = profiles
        
        for p in profiles:
            email = p.get('email') or 'No email'
            self.profile_listbox.insert(tk.END, f"{p['name']} ({email})")
    
    def refresh_sessions(self):
        """Refresh sessions list"""
        if not self.session_manager:
            return
        
        # Clear tree
        for item in self.sessions_tree.get_children():
            self.sessions_tree.delete(item)
        
        sessions = self.session_manager.list_saved_sessions()
        
        for s in sessions:
            status = '❌ Expired' if s.get('expired') else '✅ Valid'
            self.sessions_tree.insert('', tk.END, values=(
                s['name'],
                s['browser'],
                s['backup_time'][:19] if s['backup_time'] != 'unknown' else 'unknown',
                s['cookie_count'],
                status
            ))
    
    def backup_session(self):
        """Backup browser session"""
        if self.is_busy:
            return
        
        if not self.session_manager or not self.browser_helper:
            messagebox.showerror(self.t('error'), "Session Manager not available")
            return
        
        browser = self.browser_var.get()
        selection = self.profile_listbox.curselection()
        
        if not browser or not selection:
            messagebox.showwarning(self.t('warning'), "Please select browser and profile")
            return
        
        profile = self.profiles_data[selection[0]]
        
        # Check if browser running
        if self.browser_helper.is_browser_running(browser):
            if messagebox.askyesno(self.t('warning'), 
                f"{browser} is running. Close it to backup?\n\nمرورگر باز است. ببندم؟"):
                self.browser_helper.close_browser_gracefully(browser)
                import time
                time.sleep(1)
                if self.browser_helper.is_browser_running(browser):
                    self.browser_helper.kill_browser_processes(browser)
            else:
                return
        
        # Backup
        self.start_progress()
        self.set_status("Backing up session...")
        
        def do_backup():
            try:
                result = self.session_manager.backup_session(browser, profile['path'])
                self.root.after(0, lambda: self.log(f"Backup created: {result}", 'success'))
                self.root.after(0, lambda: messagebox.showinfo(
                    self.t('success'), 
                    f"{self.t('backup_success')}\n\nPath: {result}"
                ))
                self.root.after(0, self.refresh_sessions)
            except Exception as e:
                self.root.after(0, lambda: self.log(f"Backup failed: {e}", 'error'))
                self.root.after(0, lambda: messagebox.showerror(self.t('error'), str(e)))
            finally:
                self.root.after(0, self.stop_progress)
        
        thread = threading.Thread(target=do_backup)
        thread.daemon = True
        thread.start()
    
    def restore_session(self):
        """Restore browser session"""
        if self.is_busy:
            return
        
        if not self.session_manager or not self.browser_helper:
            messagebox.showerror(self.t('error'), "Session Manager not available")
            return
        
        # Get selected session
        selection = self.sessions_tree.selection()
        if not selection:
            messagebox.showwarning(self.t('warning'), "Please select a session to restore")
            return
        
        item = self.sessions_tree.item(selection[0])
        session_name = item['values'][0]
        session_browser = item['values'][1]
        
        # Get target profile
        browser = self.browser_var.get()
        profile_selection = self.profile_listbox.curselection()
        
        if not browser or not profile_selection:
            messagebox.showwarning(self.t('warning'), "Please select target browser and profile")
            return
        
        profile = self.profiles_data[profile_selection[0]]
        
        # Check if browser running
        if self.browser_helper.is_browser_running(browser):
            if messagebox.askyesno(self.t('warning'), 
                f"{browser} is running. Close it to restore?\n\nمرورگر باز است. ببندم؟"):
                self.browser_helper.close_browser_gracefully(browser)
                import time
                time.sleep(1)
                if self.browser_helper.is_browser_running(browser):
                    self.browser_helper.kill_browser_processes(browser)
            else:
                return
        
        # Restore
        self.start_progress()
        self.set_status("Restoring session...")
        
        def do_restore():
            try:
                result = self.session_manager.restore_session(session_name, browser, profile['path'])
                if result:
                    self.root.after(0, lambda: self.log(f"Session restored: {session_name}", 'success'))
                    self.root.after(0, lambda: messagebox.showinfo(
                        self.t('success'), self.t('restore_success')
                    ))
                else:
                    self.root.after(0, lambda: self.log("Restore failed", 'error'))
            except Exception as e:
                self.root.after(0, lambda: self.log(f"Restore failed: {e}", 'error'))
                self.root.after(0, lambda: messagebox.showerror(self.t('error'), str(e)))
            finally:
                self.root.after(0, self.stop_progress)
        
        thread = threading.Thread(target=do_restore)
        thread.daemon = True
        thread.start()
    
    def delete_session(self):
        """Delete selected session"""
        if not self.session_manager:
            return
        
        selection = self.sessions_tree.selection()
        if not selection:
            messagebox.showwarning(self.t('warning'), "Please select a session to delete")
            return
        
        item = self.sessions_tree.item(selection[0])
        session_name = item['values'][0]
        
        if messagebox.askyesno(self.t('confirm_title'), f"Delete session '{session_name}'?"):
            if self.session_manager.delete_session(session_name):
                self.log(f"Deleted session: {session_name}", 'success')
                self.refresh_sessions()
            else:
                messagebox.showerror(self.t('error'), "Failed to delete session")
    
    # ==================== Browser Helper ====================
    
    def detect_browsers(self):
        """Detect installed browsers and profiles"""
        if not self.browser_helper:
            messagebox.showerror(self.t('error'), "Browser Helper not available")
            return
        
        # Clear tree
        for item in self.browser_tree.get_children():
            self.browser_tree.delete(item)
        
        browsers = self.browser_helper.detect_installed_browsers()
        self.log(f"Found {len(browsers)} browsers: {', '.join(browsers)}", 'info')
        
        for browser in browsers:
            profiles = self.browser_helper.get_browser_profiles_with_email(browser)
            for p in profiles:
                self.browser_tree.insert('', tk.END, values=(
                    browser.capitalize(),
                    p['name'],
                    p.get('email') or 'No email',
                    p['path']
                ))
    
    def clean_browser_traces(self):
        """Clean browser traces"""
        if not self.browser_helper:
            messagebox.showerror(self.t('error'), "Browser Helper not available")
            return
        
        if not messagebox.askyesno(self.t('confirm_title'), 
            "Clean Antigravity traces from all browsers?\n\nپاکسازی ردهای Antigravity از همه مرورگرها؟"):
            return
        
        self.start_progress()
        
        def do_clean():
            try:
                browsers = self.browser_helper.detect_installed_browsers()
                total_cookies = 0
                total_cache = 0
                
                for browser in browsers:
                    stats = self.browser_helper.clean_browser_completely(browser)
                    total_cookies += stats['cookies']
                    total_cache += stats['cache']
                    self.root.after(0, lambda b=browser, s=stats: 
                        self.log(f"{b}: {s['cookies']} cookies, {s['cache']} cache items", 'success'))
                
                self.root.after(0, lambda: messagebox.showinfo(
                    self.t('success'),
                    f"Cleaned {total_cookies} cookies and {total_cache} cache items"
                ))
            except Exception as e:
                self.root.after(0, lambda: self.log(f"Error: {e}", 'error'))
            finally:
                self.root.after(0, self.stop_progress)
        
        thread = threading.Thread(target=do_clean)
        thread.daemon = True
        thread.start()
    
    # ==================== Network ====================
    
    def network_output_write(self, text):
        """Write to network output"""
        self.network_output.configure(state=tk.NORMAL)
        self.network_output.insert(tk.END, text + "\n")
        self.network_output.see(tk.END)
        self.network_output.configure(state=tk.DISABLED)
    
    def flush_dns(self):
        """Flush DNS cache"""
        import subprocess
        
        self.network_output.configure(state=tk.NORMAL)
        self.network_output.delete(1.0, tk.END)
        self.network_output.configure(state=tk.DISABLED)
        
        self.network_output_write("Flushing DNS cache...")
        
        if platform.system() == "Windows":
            cmd = "ipconfig /flushdns"
        elif platform.system() == "Darwin":
            cmd = "dscacheutil -flushcache"
        else:
            cmd = "resolvectl flush-caches"
        
        try:
            result = subprocess.run(cmd, shell=True, capture_output=True, text=True)
            self.network_output_write(result.stdout or "DNS cache flushed successfully!")
            self.log("DNS cache flushed", 'success')
        except Exception as e:
            self.network_output_write(f"Error: {e}")
    
    def run_network_diagnostics(self):
        """Run network diagnostics"""
        if not self.network_optimizer:
            messagebox.showerror(self.t('error'), "Network Optimizer not available")
            return
        
        self.network_output.configure(state=tk.NORMAL)
        self.network_output.delete(1.0, tk.END)
        self.network_output.configure(state=tk.DISABLED)
        
        self.start_progress()
        
        def do_diag():
            try:
                report = self.network_optimizer.generate_diagnostic_report()
                self.root.after(0, lambda: self.network_output_write(report))
            except Exception as e:
                self.root.after(0, lambda: self.network_output_write(f"Error: {e}"))
            finally:
                self.root.after(0, self.stop_progress)
        
        thread = threading.Thread(target=do_diag)
        thread.daemon = True
        thread.start()
    
    def reset_network_stack(self):
        """Reset network stack (Windows only, requires admin)"""
        if platform.system() != "Windows":
            messagebox.showinfo("Info", "Network stack reset is only available on Windows")
            return
        
        if not messagebox.askyesno(self.t('warning'), 
            "This will reset Windows network stack.\nA restart is recommended after.\n\nContinue?"):
            return
        
        import subprocess
        
        self.network_output.configure(state=tk.NORMAL)
        self.network_output.delete(1.0, tk.END)
        self.network_output.configure(state=tk.DISABLED)
        
        commands = [
            ("netsh winsock reset", "Resetting Winsock..."),
            ("netsh int ip reset", "Resetting IP stack..."),
        ]
        
        for cmd, msg in commands:
            self.network_output_write(msg)
            try:
                result = subprocess.run(cmd, shell=True, capture_output=True, text=True)
                self.network_output_write(result.stdout or "Done")
            except Exception as e:
                self.network_output_write(f"Error: {e}")
        
        self.network_output_write("\n⚠️ Please restart your computer for changes to take effect.")
        self.log("Network stack reset - restart recommended", 'warning')


# ==================== Main Entry ====================

def main():
    """Main entry point"""
    root = tk.Tk()
    app = AntigravityCleanerGUI(root)
    root.mainloop()


if __name__ == "__main__":
    main()
