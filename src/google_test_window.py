"""
Google Services Test Window
============================

A dedicated popup window for testing Google services connectivity.
Shows categorized results with clear status indicators.

Copyright (c) 2024-2025 Tawana Mohammadi / Tawana Network
All Rights Reserved - See LICENSE file for details.
"""

import customtkinter as ctk
from tkinter import messagebox
import threading
from datetime import datetime


class AppleColors:
    """Apple Design Colors"""
    BLUE = "#007AFF"
    GREEN = "#34C759"
    RED = "#FF3B30"
    ORANGE = "#FF9500"
    YELLOW = "#FFCC00"
    PURPLE = "#AF52DE"
    
    LABEL = "#000000"
    SECONDARY_LABEL = "#3C3C43"
    BG_PRIMARY = "#FFFFFF"
    BG_SECONDARY = "#F2F2F7"
    BG_TERTIARY = "#E5E5EA"
    SEPARATOR = "#C6C6C8"


class GoogleTestWindow(ctk.CTkToplevel):
    """
    Popup window for testing Google services connectivity.
    
    Features:
    - Categorized service list
    - Real-time status updates
    - Clear visual indicators
    - User-friendly explanations
    """
    
    # Categorized Google Services
    SERVICE_CATEGORIES = {
        "🔐 Authentication (احراز هویت)": {
            "description": "Services required for Google login",
            "description_fa": "سرویس‌های لازم برای ورود به حساب گوگل",
            "services": [
                ("accounts.google.com", "Google Accounts", "Main login page"),
                ("oauth2.googleapis.com", "OAuth Server", "Authentication tokens"),
                ("www.googleapis.com", "Google APIs", "API authentication"),
            ]
        },
        "🤖 AI Services (سرویس‌های هوش مصنوعی)": {
            "description": "Gemini and AI Studio services",
            "description_fa": "سرویس‌های جمینی و استودیو هوش مصنوعی",
            "services": [
                ("gemini.google.com", "Gemini AI", "Gemini chatbot"),
                ("aistudio.google.com", "AI Studio", "AI development platform"),
                ("generativelanguage.googleapis.com", "Generative AI API", "AI API endpoint"),
            ]
        },
        "☁️ Cloud & Developer (کلود و توسعه‌دهنده)": {
            "description": "Cloud and developer services",
            "description_fa": "سرویس‌های ابری و توسعه‌دهندگان",
            "services": [
                ("cloud.google.com", "Google Cloud", "Cloud console"),
                ("console.cloud.google.com", "Cloud Console", "Project management"),
            ]
        },
        "📦 Google Products (محصولات گوگل)": {
            "description": "Common Google products",
            "description_fa": "محصولات رایج گوگل",
            "services": [
                ("www.google.com", "Google Search", "Main search"),
                ("drive.google.com", "Google Drive", "Cloud storage"),
                ("www.youtube.com", "YouTube", "Video platform"),
            ]
        },
    }
    
    def __init__(self, parent, google_checker, logger=None):
        super().__init__(parent)
        
        self.checker = google_checker
        self.logger = logger
        self.results = {}
        self.is_testing = False
        
        # Window setup
        self.title("🌐 Google Services Test")
        self.geometry("700x600")
        self.minsize(600, 500)
        
        # Center on parent
        self.transient(parent)
        self.grab_set()
        
        # Build UI
        self.create_ui()
        
        # Start test automatically
        self.after(500, self.start_test)
    
    def create_ui(self):
        """Create the UI"""
        # Header
        header_frame = ctk.CTkFrame(self, fg_color=AppleColors.BLUE, corner_radius=0)
        header_frame.pack(fill="x")
        
        title = ctk.CTkLabel(
            header_frame,
            text="🌐 Google Services Connectivity Test",
            font=ctk.CTkFont(size=20, weight="bold"),
            text_color="#FFFFFF"
        )
        title.pack(pady=20)
        
        subtitle = ctk.CTkLabel(
            header_frame,
            text="تست اتصال به سرویس‌های گوگل",
            font=ctk.CTkFont(size=14),
            text_color="#FFFFFF"
        )
        subtitle.pack(pady=(0, 15))
        
        # Summary bar
        self.summary_frame = ctk.CTkFrame(self, fg_color=AppleColors.BG_TERTIARY, height=60)
        self.summary_frame.pack(fill="x", padx=20, pady=15)
        self.summary_frame.pack_propagate(False)
        
        self.summary_label = ctk.CTkLabel(
            self.summary_frame,
            text="⏳ Starting test...",
            font=ctk.CTkFont(size=16, weight="bold"),
            text_color=AppleColors.LABEL
        )
        self.summary_label.pack(expand=True)
        
        # Main content - scrollable
        self.content = ctk.CTkScrollableFrame(
            self,
            fg_color=AppleColors.BG_SECONDARY,
            corner_radius=0
        )
        self.content.pack(fill="both", expand=True, padx=20, pady=(0, 15))
        
        # Create category sections
        self.service_labels = {}
        
        for category, data in self.SERVICE_CATEGORIES.items():
            self.create_category_section(category, data)
        
        # Bottom buttons
        btn_frame = ctk.CTkFrame(self, fg_color="transparent")
        btn_frame.pack(fill="x", padx=20, pady=(0, 20))
        
        self.test_btn = ctk.CTkButton(
            btn_frame,
            text="🔄 Retest",
            font=ctk.CTkFont(size=14, weight="bold"),
            height=40,
            corner_radius=20,
            fg_color=AppleColors.BLUE,
            command=self.start_test
        )
        self.test_btn.pack(side="left", padx=5, expand=True, fill="x")
        
        close_btn = ctk.CTkButton(
            btn_frame,
            text="✓ Close",
            font=ctk.CTkFont(size=14, weight="bold"),
            height=40,
            corner_radius=20,
            fg_color=AppleColors.GREEN,
            command=self.destroy
        )
        close_btn.pack(side="right", padx=5, expand=True, fill="x")
    
    def create_category_section(self, category, data):
        """Create a category section with services"""
        # Category header
        cat_frame = ctk.CTkFrame(self.content, fg_color=AppleColors.BG_PRIMARY, corner_radius=12)
        cat_frame.pack(fill="x", pady=8, padx=5)
        
        # Category title
        cat_title = ctk.CTkLabel(
            cat_frame,
            text=category,
            font=ctk.CTkFont(size=15, weight="bold"),
            text_color=AppleColors.LABEL
        )
        cat_title.pack(pady=(15, 5), padx=15, anchor="w")
        
        # Description
        desc_text = f"{data['description']} | {data['description_fa']}"
        desc = ctk.CTkLabel(
            cat_frame,
            text=desc_text,
            font=ctk.CTkFont(size=11),
            text_color=AppleColors.SECONDARY_LABEL
        )
        desc.pack(pady=(0, 10), padx=15, anchor="w")
        
        # Services in this category
        for domain, name, desc_short in data['services']:
            self.create_service_row(cat_frame, domain, name, desc_short)
        
        # Bottom padding
        spacer = ctk.CTkFrame(cat_frame, fg_color="transparent", height=10)
        spacer.pack()
    
    def create_service_row(self, parent, domain, name, description):
        """Create a service status row"""
        row = ctk.CTkFrame(parent, fg_color=AppleColors.BG_SECONDARY, corner_radius=8)
        row.pack(fill="x", padx=15, pady=3)
        
        # Status icon (will be updated)
        status_label = ctk.CTkLabel(
            row,
            text="⏳",
            font=ctk.CTkFont(size=14)
        )
        status_label.pack(side="left", padx=(10, 5), pady=8)
        
        # Service info
        info_frame = ctk.CTkFrame(row, fg_color="transparent")
        info_frame.pack(side="left", fill="x", expand=True, pady=5)
        
        name_label = ctk.CTkLabel(
            info_frame,
            text=name,
            font=ctk.CTkFont(size=13, weight="bold"),
            text_color=AppleColors.LABEL
        )
        name_label.pack(anchor="w")
        
        domain_label = ctk.CTkLabel(
            info_frame,
            text=domain,
            font=ctk.CTkFont(size=10),
            text_color=AppleColors.SECONDARY_LABEL
        )
        domain_label.pack(anchor="w")
        
        # Response time / Status
        time_label = ctk.CTkLabel(
            row,
            text="Testing...",
            font=ctk.CTkFont(size=12),
            text_color=AppleColors.SECONDARY_LABEL
        )
        time_label.pack(side="right", padx=10, pady=8)
        
        # HTTP Status
        http_label = ctk.CTkLabel(
            row,
            text="---",
            font=ctk.CTkFont(size=11),
            text_color=AppleColors.SECONDARY_LABEL
        )
        http_label.pack(side="right", padx=5)
        
        # Store references for updating
        self.service_labels[domain] = {
            'status': status_label,
            'time': time_label,
            'http': http_label,
            'name': name
        }
    
    def start_test(self):
        """Start connectivity test"""
        if self.is_testing:
            return
        
        self.is_testing = True
        self.test_btn.configure(state="disabled", text="⏳ Testing...")
        self.summary_label.configure(text="⏳ Testing all services...", text_color=AppleColors.BLUE)
        
        # Reset all status icons
        for domain, labels in self.service_labels.items():
            labels['status'].configure(text="⏳")
            labels['time'].configure(text="Testing...", text_color=AppleColors.SECONDARY_LABEL)
            labels['http'].configure(text="---")
        
        # Run test in thread
        threading.Thread(target=self._run_test, daemon=True).start()
    
    def _run_test(self):
        """Run the test (in background thread)"""
        import socket
        import ssl
        import time
        import urllib.request
        import urllib.error
        
        success_count = 0
        fail_count = 0
        total = 0
        
        for category, data in self.SERVICE_CATEGORIES.items():
            for domain, name, desc in data['services']:
                total += 1
                result = self._test_service(domain)
                
                if result['success']:
                    success_count += 1
                    status_emoji = "✅"
                    time_text = f"{result['time']}ms"
                    time_color = AppleColors.GREEN
                else:
                    fail_count += 1
                    status_emoji = "❌"
                    time_text = result.get('error', 'Failed')[:15]
                    time_color = AppleColors.RED
                
                http_text = f"HTTP {result['status']}" if result['status'] else "N/A"
                http_color = AppleColors.GREEN if result['status'] and result['status'] < 400 else AppleColors.RED
                
                # Update UI
                self.after(0, lambda d=domain, s=status_emoji, t=time_text, tc=time_color, h=http_text, hc=http_color: 
                    self._update_service_ui(d, s, t, tc, h, hc))
        
        # Update summary
        score = int((success_count / total) * 100) if total > 0 else 0
        
        if score >= 80:
            summary_text = f"✅ Excellent! {success_count}/{total} services OK ({score}%)"
            summary_color = AppleColors.GREEN
        elif score >= 50:
            summary_text = f"⚠️ Partial: {success_count}/{total} services OK ({score}%)"
            summary_color = AppleColors.ORANGE
        else:
            summary_text = f"❌ Issues: {success_count}/{total} services OK ({score}%)"
            summary_color = AppleColors.RED
        
        self.after(0, lambda: self.summary_label.configure(text=summary_text, text_color=summary_color))
        self.after(0, lambda: self.test_btn.configure(state="normal", text="🔄 Retest"))
        self.is_testing = False
    
    def _test_service(self, domain):
        """Test a single service - returns dict with success, time, status, error"""
        import socket
        import ssl
        import time
        import urllib.request
        import urllib.error
        
        result = {
            'domain': domain,
            'success': False,
            'time': None,
            'status': None,
            'error': None
        }
        
        try:
            # 1. DNS Resolution first
            socket.gethostbyname(domain)
            
            # 2. HTTPS Request with status code
            start = time.time()
            url = f"https://{domain}/"
            
            context = ssl.create_default_context()
            req = urllib.request.Request(url, headers={
                'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36'
            })
            
            response = urllib.request.urlopen(req, timeout=10, context=context)
            elapsed = int((time.time() - start) * 1000)
            
            result['success'] = True
            result['time'] = elapsed
            result['status'] = response.getcode()
            
        except urllib.error.HTTPError as e:
            # HTTP errors still mean connection worked
            elapsed = int((time.time() - start) * 1000) if 'start' in dir() else 0
            result['time'] = elapsed
            result['status'] = e.code
            result['success'] = e.code < 500  # 4xx is OK (redirect, auth required, etc)
            result['error'] = f"HTTP {e.code}"
            
        except urllib.error.URLError as e:
            result['error'] = str(e.reason)[:30]
            
        except socket.timeout:
            result['error'] = "Timeout"
            
        except socket.gaierror:
            result['error'] = "DNS Failed"
            
        except Exception as e:
            result['error'] = str(e)[:30]
        
        return result
    
    def _update_service_ui(self, domain, status, time_text, time_color, http_text, http_color):
        """Update service UI elements"""
        if domain in self.service_labels:
            labels = self.service_labels[domain]
            labels['status'].configure(text=status)
            labels['time'].configure(text=time_text, text_color=time_color)
            labels['http'].configure(text=http_text, text_color=http_color)


def open_google_test_window(parent, google_checker, logger=None):
    """Helper function to open the Google test window"""
    return GoogleTestWindow(parent, google_checker, logger)
