"""
Antigravity Cleaner - Utilities
===============================
Shared utility functions and dependency management.
"""

import os
import sys

def get_base_path():
    """Get the base path for data storage (Portable friendly)"""
    try:
        if getattr(sys, 'frozen', False):
            # Running as EXE
            base = os.path.dirname(sys.executable)
        else:
            # Running as Script
            base = os.path.dirname(os.path.abspath(__file__))
        
        # Check if we should use local 'data' folder (Portable standard)
        local_data = os.path.join(base, 'data')
        
        # Ensure 'data' folder exists for portable mode
        if os.path.exists(local_data) or getattr(sys, 'frozen', False):
            try:
                os.makedirs(local_data, exist_ok=True)
                return local_data
            except:
                pass
        
        # Fallback to user home
        return os.path.join(os.path.expanduser('~'), '.antigravity-cleaner')
    except Exception:
        return os.path.join(os.path.expanduser('~'), '.antigravity-cleaner')

class DependencyLoader:
    """Helper to load optional modules safely"""
    
    def __init__(self, logger=None):
        self.logger = logger
        self.browser_helper = None
        self.network_optimizer = None
        self.session_manager = None
        self.detector = None
        self.google_checker = None
        self.google_window = None

    def load_all(self):
        """Attempt to load all helper modules"""
        try:
            from browser_helper import BrowserHelper
            self.browser_helper = BrowserHelper
        except ImportError as e:
            self.log_error("BrowserHelper", e)
            
        try:
            from network_optimizer import NetworkOptimizer
            self.network_optimizer = NetworkOptimizer
        except ImportError as e:
            self.log_error("NetworkOptimizer", e)

        try:
            from session_manager import SessionManager
            self.session_manager = SessionManager
        except ImportError as e:
             self.log_error("SessionManager", e)

        try:
            from antigravity_detector import AntigravityDetector
            self.detector = AntigravityDetector
        except ImportError as e:
            self.log_error("AntigravityDetector", e)
            
        try:
            from google_checker import GoogleServicesChecker
            self.google_checker = GoogleServicesChecker
        except ImportError as e:
            self.log_error("GoogleServicesChecker", e)
            
        try:
            from google_test_window import open_google_test_window
            self.google_window = open_google_test_window
        except ImportError as e:
             self.log_error("GoogleTestWindow", e)
             
    def log_error(self, module, error):
        if self.logger:
            self.logger.warning(f"Failed to load {module}: {error}")
