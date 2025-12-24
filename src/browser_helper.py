"""
Browser Helper Module
=====================

Provides safe, targeted browser cleaning for Antigravity-related traces.
Supports Chrome, Edge, Brave, and Firefox across multiple profiles.

License: MIT
"""

import os
import sys
import platform
import shutil
import sqlite3
import json
import glob
from datetime import datetime
from typing import List, Dict, Optional, Tuple
import logging

try:
    import psutil
except ImportError:
    print("Missing psutil. Install: pip install psutil")
    sys.exit(1)


class BrowserHelper:
    """
    Handles browser-specific operations for Antigravity cleaning.
    
    Features:
    - Detect installed browsers and profiles
    - Selectively clean Antigravity-related data only
    - Manage browser processes (close/kill)
    - Create backups before deletion
    """
    
    SUPPORTED_BROWSERS = {
        'chrome': {
            'name': 'Google Chrome',
            'process_names': ['chrome.exe', 'chrome'],
            'data_paths': {
                'windows': os.path.join(os.environ.get('LOCALAPPDATA', ''), 'Google', 'Chrome', 'User Data'),
                'darwin': os.path.join(os.path.expanduser('~'), 'Library', 'Application Support', 'Google', 'Chrome'),
                'linux': os.path.join(os.path.expanduser('~'), '.config', 'google-chrome')
            }
        },
        'edge': {
            'name': 'Microsoft Edge',
            'process_names': ['msedge.exe', 'msedge'],
            'data_paths': {
                'windows': os.path.join(os.environ.get('LOCALAPPDATA', ''), 'Microsoft', 'Edge', 'User Data'),
                'darwin': os.path.join(os.path.expanduser('~'), 'Library', 'Application Support', 'Microsoft Edge'),
                'linux': os.path.join(os.path.expanduser('~'), '.config', 'microsoft-edge')
            }
        },
        'brave': {
            'name': 'Brave Browser',
            'process_names': ['brave.exe', 'brave'],
            'data_paths': {
                'windows': os.path.join(os.environ.get('LOCALAPPDATA', ''), 'BraveSoftware', 'Brave-Browser', 'User Data'),
                'darwin': os.path.join(os.path.expanduser('~'), 'Library', 'Application Support', 'BraveSoftware', 'Brave-Browser'),
                'linux': os.path.join(os.path.expanduser('~'), '.config', 'BraveSoftware', 'Brave-Browser')
            }
        },
        'chromium': {
            'name': 'Chromium',
            'process_names': ['chromium-browser', 'chromium'],
            'data_paths': {
                'windows': os.path.join(os.environ.get('LOCALAPPDATA', ''), 'Chromium', 'User Data'),
                'darwin': os.path.join(os.path.expanduser('~'), 'Library', 'Application Support', 'Chromium'),
                'linux': os.path.join(os.path.expanduser('~'), '.config', 'chromium')
            }
        },
        'firefox': {
            'name': 'Mozilla Firefox',
            'process_names': ['firefox.exe', 'firefox'],
            'data_paths': {
                'windows': os.path.join(os.environ.get('APPDATA', ''), 'Mozilla', 'Firefox', 'Profiles'),
                'darwin': os.path.join(os.path.expanduser('~'), 'Library', 'Application Support', 'Firefox', 'Profiles'),
                'linux': os.path.join(os.path.expanduser('~'), '.mozilla', 'firefox')
            }
        },
        'opera': {
            'name': 'Opera',
            'process_names': ['opera.exe', 'opera'],
            'data_paths': {
                'windows': os.path.join(os.environ.get('APPDATA', ''), 'Opera Software', 'Opera Stable'),
                'darwin': os.path.join(os.path.expanduser('~'), 'Library', 'Application Support', 'com.operasoftware.Opera'),
                'linux': os.path.join(os.path.expanduser('~'), '.config', 'opera')
            }
        },
        'opera_gx': {
            'name': 'Opera GX',
            'process_names': ['opera.exe', 'opera'],
            'data_paths': {
                'windows': os.path.join(os.environ.get('APPDATA', ''), 'Opera Software', 'Opera GX Stable'),
                'darwin': os.path.join(os.path.expanduser('~'), 'Library', 'Application Support', 'com.operasoftware.OperaGX'),
                'linux': os.path.join(os.path.expanduser('~'), '.config', 'opera-gx')
            }
        },
        'vivaldi': {
            'name': 'Vivaldi',
            'process_names': ['vivaldi.exe', 'vivaldi'],
            'data_paths': {
                'windows': os.path.join(os.environ.get('LOCALAPPDATA', ''), 'Vivaldi', 'User Data'),
                'darwin': os.path.join(os.path.expanduser('~'), 'Library', 'Application Support', 'Vivaldi'),
                'linux': os.path.join(os.path.expanduser('~'), '.config', 'vivaldi')
            }
        },
        'arc': {
            'name': 'Arc Browser',
            'process_names': ['Arc', 'arc'],
            'data_paths': {
                'windows': '',  # Arc is Mac-only currently
                'darwin': os.path.join(os.path.expanduser('~'), 'Library', 'Application Support', 'Arc', 'User Data'),
                'linux': ''
            }
        },
        'safari': {
            'name': 'Safari',
            'process_names': ['Safari'],
            'data_paths': {
                'windows': '',  # Safari is Mac-only
                'darwin': os.path.join(os.path.expanduser('~'), 'Library', 'Safari'),
                'linux': ''
            }
        },
        'waterfox': {
            'name': 'Waterfox',
            'process_names': ['waterfox.exe', 'waterfox'],
            'data_paths': {
                'windows': os.path.join(os.environ.get('APPDATA', ''), 'Waterfox', 'Profiles'),
                'darwin': os.path.join(os.path.expanduser('~'), 'Library', 'Application Support', 'Waterfox', 'Profiles'),
                'linux': os.path.join(os.path.expanduser('~'), '.waterfox')
            }
        },
        'floorp': {
            'name': 'Floorp',
            'process_names': ['floorp.exe', 'floorp'],
            'data_paths': {
                'windows': os.path.join(os.environ.get('APPDATA', ''), 'Floorp', 'Profiles'),
                'darwin': os.path.join(os.path.expanduser('~'), 'Library', 'Application Support', 'Floorp', 'Profiles'),
                'linux': os.path.join(os.path.expanduser('~'), '.floorp')
            }
        },
        'zen': {
            'name': 'Zen Browser',
            'process_names': ['zen.exe', 'zen'],
            'data_paths': {
                'windows': os.path.join(os.environ.get('APPDATA', ''), 'zen', 'Profiles'),
                'darwin': os.path.join(os.path.expanduser('~'), 'Library', 'Application Support', 'zen', 'Profiles'),
                'linux': os.path.join(os.path.expanduser('~'), '.zen')
            }
        },
    }

    
    # Keywords to identify Antigravity-related data
    ANTIGRAVITY_KEYWORDS = [
        'antigravity',
        'anti-gravity',
        'anti_gravity',
        'deepmind',
        'gemini-code',
        'google.com/antigravity',
        'accounts.google.com/antigravity'
    ]
    
    def __init__(self, logger: logging.Logger, dry_run: bool = False):
        """
        Initialize BrowserHelper.
        
        Args:
            logger: Logger instance for detailed logging
            dry_run: If True, only simulate operations without actual changes
        """
        self.logger = logger
        self.dry_run = dry_run
        self.current_os = platform.system().lower()
        self.backup_dir = os.path.join(os.path.expanduser('~'), '.antigravity-cleaner', 'backups')
        
        # Create backup directory
        os.makedirs(self.backup_dir, exist_ok=True)
        
        self.logger.info(f"BrowserHelper initialized (OS: {self.current_os}, Dry-run: {dry_run})")
    
    # ==================== Browser Detection ====================
    
    def detect_installed_browsers(self) -> List[str]:
        """
        Detect which supported browsers are installed.
        
        Returns:
            List of browser keys (e.g., ['chrome', 'edge'])
        """
        self.logger.info("Detecting installed browsers...")
        installed = []
        
        for browser_key, browser_info in self.SUPPORTED_BROWSERS.items():
            data_path = browser_info['data_paths'].get(self.current_os)
            
            if data_path and os.path.exists(data_path):
                installed.append(browser_key)
                self.logger.debug(f"Found {browser_info['name']} at {data_path}")
            else:
                self.logger.debug(f"{browser_info['name']} not found")
        
        self.logger.info(f"Detected browsers: {', '.join(installed) if installed else 'None'}")
        return installed
    
    def get_browser_profiles(self, browser: str) -> List[Tuple[str, str]]:
        """
        Get all profiles for a specific browser.
        
        Args:
            browser: Browser key (e.g., 'chrome')
        
        Returns:
            List of tuples: (profile_name, profile_path)
        """
        if browser not in self.SUPPORTED_BROWSERS:
            self.logger.warning(f"Unsupported browser: {browser}")
            return []
        
        data_path = self.SUPPORTED_BROWSERS[browser]['data_paths'].get(self.current_os)
        if not data_path or not os.path.exists(data_path):
            self.logger.warning(f"Browser data path not found: {data_path}")
            return []
        
        profiles = []
        
        # Firefox uses different structure
        if browser == 'firefox':
            # Firefox profiles are in format: xxxxxxxx.profile_name
            for item in os.listdir(data_path):
                profile_path = os.path.join(data_path, item)
                if os.path.isdir(profile_path) and '.' in item:
                    profiles.append((item, profile_path))
        else:
            # Chromium-based browsers
            # Check for Default profile
            default_path = os.path.join(data_path, 'Default')
            if os.path.exists(default_path):
                profiles.append(('Default', default_path))
            
            # Scan ALL Profile folders (not just 1-19)
            # This handles users with 50+ profiles
            for item in os.listdir(data_path):
                item_path = os.path.join(data_path, item)
                # Check if it's a Profile folder (Profile 1, Profile 2, etc.)
                if os.path.isdir(item_path) and item.startswith('Profile '):
                    profiles.append((item, item_path))
            
            # Sort profiles by number
            def get_profile_num(profile_tuple):
                name = profile_tuple[0]
                if name == 'Default':
                    return 0
                try:
                    return int(name.replace('Profile ', ''))
                except:
                    return 999
            
            profiles.sort(key=get_profile_num)
        
        self.logger.debug(f"Found {len(profiles)} profiles for {browser}: {[p[0] for p in profiles]}")
        return profiles
    
    def get_profile_email(self, browser: str, profile_path: str) -> Optional[str]:
        """
        Extract email address from browser profile.
        
        Args:
            browser: Browser key
            profile_path: Path to browser profile
        
        Returns:
            Email address if found, None otherwise
        """
        if browser == 'firefox':
            # Firefox doesn't store email in the same way
            return None
        
        # Chromium-based browsers store profile info in Preferences file
        prefs_file = os.path.join(profile_path, 'Preferences')
        if not os.path.exists(prefs_file):
            return None
        
        try:
            with open(prefs_file, 'r', encoding='utf-8') as f:
                prefs = json.load(f)
            
            # Try different paths where email might be stored
            # Path 1: account_info
            account_info = prefs.get('account_info', [])
            if account_info and isinstance(account_info, list):
                for account in account_info:
                    if 'email' in account:
                        return account['email']
            
            # Path 2: google.services
            google_services = prefs.get('google', {}).get('services', {})
            if google_services:
                # Direct last_username
                if 'last_username' in google_services:
                    return google_services['last_username']
                
                # Nested in signin
                signin_info = google_services.get('signin', {})
                if signin_info and 'last_username' in signin_info:
                    return signin_info['last_username']

            # Path 3: sync metadata
            sync = prefs.get('sync', {})
            if sync and 'authenticated_email' in sync:
                return sync['authenticated_email']
            
            # Path 4: signin info (root level)
            signin = prefs.get('signin', {})
            if signin and 'last_username' in signin:
                return signin['last_username']
            
            # Path 5: Profile name (sometimes contains email)
            profile = prefs.get('profile', {})
            name = profile.get('name', '')
            if '@' in name:
                return name
            
        except Exception as e:
            self.logger.debug(f"Could not read profile preferences: {e}")

        # Fallback: Check Local State (contains info for ALL profiles)
        try:
            profile_dir_name = os.path.basename(profile_path)
            user_data_root = os.path.dirname(profile_path)
            local_state_file = os.path.join(user_data_root, 'Local State')
            
            if os.path.exists(local_state_file):
                with open(local_state_file, 'r', encoding='utf-8') as f:
                    local_state = json.load(f)
                
                # Path: profile.info_cache.[ProfileDir].user_name
                info_cache = local_state.get('profile', {}).get('info_cache', {})
                profile_info = info_cache.get(profile_dir_name, {})
                
                if 'user_name' in profile_info and '@' in profile_info['user_name']:
                    return profile_info['user_name']
        except Exception as e:
            self.logger.debug(f"Could not read Local State for email fallback: {e}")
        
        return None

    
    def get_browser_profiles_with_email(self, browser: str) -> List[Dict]:
        """
        Get all profiles with their email addresses.
        
        Args:
            browser: Browser key
        
        Returns:
            List of dicts: {'name': str, 'path': str, 'email': str or None}
        """
        profiles = self.get_browser_profiles(browser)
        result = []
        
        for profile_name, profile_path in profiles:
            email = self.get_profile_email(browser, profile_path)
            result.append({
                'name': profile_name,
                'path': profile_path,
                'email': email
            })
        
        return result
    
    def search_profiles_by_email(self, browser: str, email_query: str) -> List[Dict]:
        """
        Search profiles by email address.
        
        Args:
            browser: Browser key
            email_query: Email or part of email to search
        
        Returns:
            List of matching profiles
        """
        profiles = self.get_browser_profiles_with_email(browser)
        
        if not email_query:
            return profiles
        
        email_query = email_query.lower()
        matches = []
        
        for profile in profiles:
            email = profile.get('email', '') or ''
            if email_query in email.lower():
                matches.append(profile)
        
        self.logger.debug(f"Found {len(matches)} profiles matching '{email_query}'")
        return matches
    
    # ==================== Process Management ====================
    
    def is_browser_running(self, browser: str) -> bool:
        """
        Check if browser is currently running.
        
        Args:
            browser: Browser key
        
        Returns:
            True if browser is running
        """
        if browser not in self.SUPPORTED_BROWSERS:
            return False
        
        process_names = self.SUPPORTED_BROWSERS[browser]['process_names']
        
        for proc in psutil.process_iter(['name']):
            try:
                if proc.info['name'].lower() in [p.lower() for p in process_names]:
                    return True
            except (psutil.NoSuchProcess, psutil.AccessDenied):
                pass
        
        return False
    
    def get_browser_processes(self, browser: str) -> List[psutil.Process]:
        """
        Get all running processes for a browser.
        
        Args:
            browser: Browser key
        
        Returns:
            List of Process objects
        """
        if browser not in self.SUPPORTED_BROWSERS:
            return []
        
        process_names = self.SUPPORTED_BROWSERS[browser]['process_names']
        processes = []
        
        for proc in psutil.process_iter(['pid', 'name']):
            try:
                if proc.info['name'].lower() in [p.lower() for p in process_names]:
                    processes.append(proc)
            except (psutil.NoSuchProcess, psutil.AccessDenied):
                pass
        
        return processes
    
    def close_browser_gracefully(self, browser: str) -> bool:
        """
        Attempt to close browser gracefully.
        
        Args:
            browser: Browser key
        
        Returns:
            True if browser was closed successfully
        """
        self.logger.info(f"Attempting to close {browser} gracefully...")
        
        if self.dry_run:
            self.logger.info(f"[DRY RUN] Would close {browser}")
            return True
        
        processes = self.get_browser_processes(browser)
        if not processes:
            self.logger.info(f"{browser} is not running")
            return True
        
        # Try terminate first (graceful)
        for proc in processes:
            try:
                proc.terminate()
                self.logger.debug(f"Sent terminate signal to PID {proc.pid}")
            except (psutil.NoSuchProcess, psutil.AccessDenied) as e:
                self.logger.warning(f"Could not terminate PID {proc.pid}: {e}")
        
        # Wait for processes to exit
        gone, alive = psutil.wait_procs(processes, timeout=5)
        
        if alive:
            self.logger.warning(f"{len(alive)} processes still alive after graceful close")
            return False
        
        self.logger.info(f"Successfully closed {browser}")
        return True
    
    def kill_browser_processes(self, browser: str) -> bool:
        """
        Force kill browser processes.
        
        Args:
            browser: Browser key
        
        Returns:
            True if all processes were killed
        """
        self.logger.warning(f"Force killing {browser} processes...")
        
        if self.dry_run:
            self.logger.info(f"[DRY RUN] Would kill {browser} processes")
            return True
        
        processes = self.get_browser_processes(browser)
        if not processes:
            self.logger.info(f"No {browser} processes to kill")
            return True
        
        for proc in processes:
            try:
                proc.kill()
                self.logger.debug(f"Killed PID {proc.pid}")
            except (psutil.NoSuchProcess, psutil.AccessDenied) as e:
                self.logger.error(f"Could not kill PID {proc.pid}: {e}")
        
        # Verify all killed
        gone, alive = psutil.wait_procs(processes, timeout=3)
        
        if alive:
            self.logger.error(f"Failed to kill {len(alive)} processes")
            return False
        
        self.logger.info(f"Successfully killed all {browser} processes")
        return True
    
    # ==================== Backup Operations ====================
    
    def create_backup(self, file_path: str) -> Optional[str]:
        """
        Create backup of a file before modification.
        
        Args:
            file_path: Path to file to backup
        
        Returns:
            Path to backup file, or None if failed
        """
        if not os.path.exists(file_path):
            self.logger.warning(f"Cannot backup non-existent file: {file_path}")
            return None
        
        timestamp = datetime.now().strftime("%Y%m%d_%H%M%S")
        filename = os.path.basename(file_path)
        backup_path = os.path.join(self.backup_dir, f"{filename}.backup_{timestamp}")
        
        if self.dry_run:
            self.logger.info(f"[DRY RUN] Would backup {file_path} to {backup_path}")
            return backup_path
        
        try:
            shutil.copy2(file_path, backup_path)
            self.logger.info(f"Created backup: {backup_path}")
            return backup_path
        except Exception as e:
            self.logger.error(f"Backup failed: {e}")
            return None
    
    def restore_backup(self, backup_path: str, original_path: str) -> bool:
        """
        Restore a file from backup.
        
        Args:
            backup_path: Path to backup file
            original_path: Path to restore to
        
        Returns:
            True if restored successfully
        """
        if not os.path.exists(backup_path):
            self.logger.error(f"Backup file not found: {backup_path}")
            return False
        
        if self.dry_run:
            self.logger.info(f"[DRY RUN] Would restore {backup_path} to {original_path}")
            return True
        
        try:
            shutil.copy2(backup_path, original_path)
            self.logger.info(f"Restored from backup: {original_path}")
            return True
        except Exception as e:
            self.logger.error(f"Restore failed: {e}")
            return False
    
    # ==================== Selective Cleaning ====================
    
    def clean_antigravity_cookies(self, browser: str, profile_path: str) -> int:
        """
        Remove only Antigravity-related cookies from browser.
        
        Args:
            browser: Browser key
            profile_path: Path to browser profile
        
        Returns:
            Number of cookies deleted
        """
        self.logger.info(f"Cleaning Antigravity cookies from {browser} profile...")
        
        if browser == 'firefox':
            cookie_db = os.path.join(profile_path, 'cookies.sqlite')
        else:
            # Chromium-based
            cookie_db = os.path.join(profile_path, 'Network', 'Cookies')
            if not os.path.exists(cookie_db):
                cookie_db = os.path.join(profile_path, 'Cookies')  # Older Chrome versions
        
        if not os.path.exists(cookie_db):
            self.logger.warning(f"Cookie database not found: {cookie_db}")
            return 0
        
        # Create backup
        backup = self.create_backup(cookie_db)
        if not backup and not self.dry_run:
            self.logger.error("Backup failed, aborting cookie cleaning")
            return 0
        
        deleted_count = 0
        
        try:
            # Connect to SQLite database
            conn = sqlite3.connect(cookie_db)
            cursor = conn.cursor()
            
            # Find Antigravity-related cookies
            for keyword in self.ANTIGRAVITY_KEYWORDS:
                if self.dry_run:
                    cursor.execute("SELECT COUNT(*) FROM cookies WHERE host_key LIKE ? OR name LIKE ?", 
                                   (f'%{keyword}%', f'%{keyword}%'))
                    count = cursor.fetchone()[0]
                    self.logger.info(f"[DRY RUN] Would delete {count} cookies matching '{keyword}'")
                    deleted_count += count
                else:
                    cursor.execute("DELETE FROM cookies WHERE host_key LIKE ? OR name LIKE ?", 
                                   (f'%{keyword}%', f'%{keyword}%'))
                    deleted = cursor.rowcount
                    if deleted > 0:
                        self.logger.debug(f"Deleted {deleted} cookies matching '{keyword}'")
                        deleted_count += deleted
            
            if not self.dry_run:
                conn.commit()
            
            conn.close()
            
            self.logger.info(f"Cleaned {deleted_count} Antigravity cookies")
            return deleted_count
            
        except sqlite3.Error as e:
            self.logger.error(f"SQLite error: {e}")
            if backup and not self.dry_run:
                self.logger.info("Attempting to restore from backup...")
                self.restore_backup(backup, cookie_db)
            return 0
    
    def clean_antigravity_localstorage(self, browser: str, profile_path: str) -> int:
        """
        Remove Antigravity-related LocalStorage data.
        
        Args:
            browser: Browser key
            profile_path: Path to browser profile
        
        Returns:
            Number of items deleted
        """
        self.logger.info(f"Cleaning Antigravity LocalStorage from {browser} profile...")
        
        if browser == 'firefox':
            ls_path = os.path.join(profile_path, 'webappsstore.sqlite')
        else:
            ls_path = os.path.join(profile_path, 'Local Storage', 'leveldb')
        
        if not os.path.exists(ls_path):
            self.logger.warning(f"LocalStorage not found: {ls_path}")
            return 0
        
        deleted_count = 0
        
        # For Chromium browsers, LocalStorage is in LevelDB format (complex)
        # For safety, we'll only delete specific files matching keywords
        if browser != 'firefox' and os.path.isdir(ls_path):
            for file in os.listdir(ls_path):
                file_path = os.path.join(ls_path, file)
                if any(keyword in file.lower() for keyword in self.ANTIGRAVITY_KEYWORDS):
                    if self.dry_run:
                        self.logger.info(f"[DRY RUN] Would delete: {file_path}")
                        deleted_count += 1
                    else:
                        try:
                            os.remove(file_path)
                            self.logger.debug(f"Deleted: {file_path}")
                            deleted_count += 1
                        except Exception as e:
                            self.logger.error(f"Failed to delete {file_path}: {e}")
        
        self.logger.info(f"Cleaned {deleted_count} LocalStorage items")
        return deleted_count
    
    def clean_antigravity_cache_entries(self, browser: str, profile_path: str) -> int:
        """
        Remove Antigravity-related cache entries.
        
        Args:
            browser: Browser key
            profile_path: Path to browser profile
        
        Returns:
            Number of cache entries deleted
        """
        self.logger.info(f"Cleaning Antigravity cache from {browser} profile...")
        
        cache_paths = []
        
        if browser == 'firefox':
            cache_paths.append(os.path.join(profile_path, 'cache2', 'entries'))
        else:
            # Chromium-based
            cache_paths.append(os.path.join(profile_path, 'Cache', 'Cache_Data'))
            cache_paths.append(os.path.join(profile_path, 'Code Cache'))
            cache_paths.append(os.path.join(profile_path, 'GPUCache'))
        
        deleted_count = 0
        
        for cache_path in cache_paths:
            if not os.path.exists(cache_path):
                continue
            
            if os.path.isdir(cache_path):
                for root, dirs, files in os.walk(cache_path):
                    for file in files:
                        # Only delete files with Antigravity keywords in name
                        if any(keyword in file.lower() for keyword in self.ANTIGRAVITY_KEYWORDS):
                            file_path = os.path.join(root, file)
                            if self.dry_run:
                                self.logger.debug(f"[DRY RUN] Would delete cache: {file_path}")
                                deleted_count += 1
                            else:
                                try:
                                    os.remove(file_path)
                                    self.logger.debug(f"Deleted cache: {file_path}")
                                    deleted_count += 1
                                except Exception as e:
                                    self.logger.error(f"Failed to delete {file_path}: {e}")
        
        self.logger.info(f"Cleaned {deleted_count} cache entries")
        return deleted_count
    
    def clean_browser_completely(self, browser: str) -> Dict[str, int]:
        """
        Clean all Antigravity traces from a browser (all profiles).
        
        Args:
            browser: Browser key
        
        Returns:
            Dictionary with cleaning statistics
        """
        self.logger.info(f"Starting complete cleaning for {browser}...")
        
        stats = {
            'cookies': 0,
            'localstorage': 0,
            'cache': 0,
            'profiles_cleaned': 0
        }
        
        # Check if browser is running
        if self.is_browser_running(browser):
            self.logger.warning(f"{browser} is currently running")
            if not self.close_browser_gracefully(browser):
                self.logger.warning("Graceful close failed, attempting force kill...")
                if not self.kill_browser_processes(browser):
                    self.logger.error(f"Could not close {browser}. Aborting cleaning.")
                    return stats
        
        # Get all profiles
        profiles = self.get_browser_profiles(browser)
        
        if not profiles:
            self.logger.warning(f"No profiles found for {browser}")
            return stats
        
        # Clean each profile
        for profile_name, profile_path in profiles:
            self.logger.info(f"Cleaning profile: {profile_name}")
            
            stats['cookies'] += self.clean_antigravity_cookies(browser, profile_path)
            stats['localstorage'] += self.clean_antigravity_localstorage(browser, profile_path)
            stats['cache'] += self.clean_antigravity_cache_entries(browser, profile_path)
            stats['profiles_cleaned'] += 1
        
        self.logger.info(f"Cleaning complete for {browser}: {stats}")
        return stats


if __name__ == "__main__":
    # Test code
    logging.basicConfig(
        level=logging.DEBUG,
        format='[%(asctime)s] [%(levelname)s] %(message)s',
        datefmt='%Y-%m-%d %H:%M:%S'
    )
    logger = logging.getLogger(__name__)
    
    helper = BrowserHelper(logger, dry_run=True)
    browsers = helper.detect_installed_browsers()
    
    for browser in browsers:
        profiles = helper.get_browser_profiles(browser)
        print(f"\n{browser}: {len(profiles)} profiles")
        for name, path in profiles:
            print(f"  - {name}: {path}")
