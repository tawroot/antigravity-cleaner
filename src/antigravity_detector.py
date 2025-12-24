"""
Antigravity Detector Module
===========================

Detects installation status and traces of Antigravity IDE.
Provides health score and detailed status report.

License: MIT
"""

import os
import platform
import glob
import json
from datetime import datetime
from typing import Dict, List, Tuple, Optional
import logging

try:
    import psutil
except ImportError:
    psutil = None


class AntigravityDetector:
    """
    Detects Antigravity IDE installation and traces.
    
    Features:
    - Check if Antigravity is installed
    - Detect running processes
    - Find leftover files and folders
    - Calculate system health score
    - Generate detailed status report
    """
    
    # Possible installation paths
    INSTALL_PATHS = {
        'windows': [
            os.path.join(os.environ.get('LOCALAPPDATA', ''), 'Programs', 'Antigravity'),
            os.path.join(os.environ.get('LOCALAPPDATA', ''), 'Antigravity'),
            os.path.join(os.environ.get('APPDATA', ''), 'Antigravity'),
            os.path.join(os.environ.get('PROGRAMFILES', ''), 'Antigravity'),
            os.path.join(os.environ.get('PROGRAMFILES(X86)', ''), 'Antigravity'),
        ],
        'darwin': [
            '/Applications/Antigravity.app',
            os.path.join(os.path.expanduser('~'), 'Applications', 'Antigravity.app'),
            os.path.join(os.path.expanduser('~'), 'Library', 'Application Support', 'Antigravity'),
        ],
        'linux': [
            '/usr/share/antigravity',
            '/opt/antigravity',
            os.path.join(os.path.expanduser('~'), '.local', 'share', 'Antigravity'),
            os.path.join(os.path.expanduser('~'), '.config', 'Antigravity'),
        ]
    }
    
    # Cache and temp paths
    CACHE_PATHS = {
        'windows': [
            os.path.join(os.environ.get('LOCALAPPDATA', ''), 'Antigravity', 'Cache'),
            os.path.join(os.environ.get('TEMP', ''), 'antigravity*'),
            os.path.join(os.environ.get('LOCALAPPDATA', ''), 'Google', 'CrashReports', '*antigravity*'),
        ],
        'darwin': [
            os.path.join(os.path.expanduser('~'), 'Library', 'Caches', 'Antigravity'),
            os.path.join(os.path.expanduser('~'), 'Library', 'Caches', 'com.google.antigravity'),
            '/tmp/antigravity*',
        ],
        'linux': [
            os.path.join(os.path.expanduser('~'), '.cache', 'Antigravity'),
            '/tmp/antigravity*',
        ]
    }
    
    # Process names to look for
    PROCESS_NAMES = [
        'antigravity',
        'Antigravity',
        'antigravity.exe',
        'Antigravity.exe',
    ]
    
    # Registry keys (Windows only)
    REGISTRY_KEYS = [
        r'SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\Antigravity',
        r'SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall\Antigravity',
        r'SOFTWARE\Antigravity',
    ]
    
    def __init__(self, logger: Optional[logging.Logger] = None):
        """
        Initialize AntigravityDetector.
        
        Args:
            logger: Optional logger instance
        """
        self.logger = logger or logging.getLogger(__name__)
        self.current_os = platform.system().lower()
        
        # Detection results
        self.is_installed = False
        self.is_running = False
        self.install_path = None
        self.version = None
        self.leftover_paths = []
        self.leftover_size = 0
        self.registry_entries = []
        self.health_score = 100
        
        self.logger.info(f"AntigravityDetector initialized (OS: {self.current_os})")
    
    def detect_all(self) -> Dict:
        """
        Run full detection and return results.
        
        Returns:
            Dictionary with all detection results
        """
        self.logger.info("Starting full Antigravity detection...")
        
        # Check installation
        self.is_installed, self.install_path = self._check_installed()
        
        # Check running processes
        self.is_running = self._check_running()
        
        # Find leftovers
        self.leftover_paths, self.leftover_size = self._find_leftovers()
        
        # Check registry (Windows only)
        if self.current_os == 'windows':
            self.registry_entries = self._check_registry()
        
        # Calculate health score
        self.health_score = self._calculate_health_score()
        
        results = {
            'is_installed': self.is_installed,
            'install_path': self.install_path,
            'version': self.version,
            'is_running': self.is_running,
            'leftover_paths': self.leftover_paths,
            'leftover_size': self.leftover_size,
            'leftover_size_human': self._format_size(self.leftover_size),
            'registry_entries': self.registry_entries,
            'health_score': self.health_score,
            'status': self._get_status_text(),
            'timestamp': datetime.now().isoformat()
        }
        
        self.logger.info(f"Detection complete: installed={self.is_installed}, running={self.is_running}, health={self.health_score}")
        return results
    
    def _check_installed(self) -> Tuple[bool, Optional[str]]:
        """Check if Antigravity is installed."""
        self.logger.debug("Checking installation...")
        
        paths = self.INSTALL_PATHS.get(self.current_os, [])
        
        for path in paths:
            if path and os.path.exists(path):
                self.logger.debug(f"Found installation at: {path}")
                
                # Try to find version
                version_file = os.path.join(path, 'version.json')
                if os.path.exists(version_file):
                    try:
                        with open(version_file, 'r') as f:
                            data = json.load(f)
                            self.version = data.get('version')
                    except:
                        pass
                
                return True, path
        
        return False, None
    
    def _check_running(self) -> bool:
        """Check if Antigravity process is running."""
        self.logger.debug("Checking running processes...")
        
        if not psutil:
            self.logger.warning("psutil not available, skipping process check")
            return False
        
        for proc in psutil.process_iter(['name']):
            try:
                name = proc.info['name'].lower()
                for target in self.PROCESS_NAMES:
                    if target.lower() in name:
                        self.logger.debug(f"Found running process: {name}")
                        return True
            except (psutil.NoSuchProcess, psutil.AccessDenied):
                pass
        
        return False
    
    def _find_leftovers(self) -> Tuple[List[str], int]:
        """Find leftover files and folders."""
        self.logger.debug("Finding leftover files...")
        
        leftovers = []
        total_size = 0
        
        # Check install paths
        for path in self.INSTALL_PATHS.get(self.current_os, []):
            if path and os.path.exists(path):
                leftovers.append(path)
                total_size += self._get_dir_size(path)
        
        # Check cache paths (expand globs)
        for pattern in self.CACHE_PATHS.get(self.current_os, []):
            if pattern:
                for path in glob.glob(pattern):
                    if os.path.exists(path):
                        leftovers.append(path)
                        total_size += self._get_dir_size(path) if os.path.isdir(path) else os.path.getsize(path)
        
        self.logger.debug(f"Found {len(leftovers)} leftover paths ({self._format_size(total_size)})")
        return leftovers, total_size
    
    def _check_registry(self) -> List[Dict]:
        """Check Windows registry for Antigravity entries."""
        entries = []
        
        if self.current_os != 'windows':
            return entries
        
        try:
            import winreg
            
            for key_path in self.REGISTRY_KEYS:
                for hive in [winreg.HKEY_CURRENT_USER, winreg.HKEY_LOCAL_MACHINE]:
                    try:
                        with winreg.OpenKey(hive, key_path) as key:
                            entries.append({
                                'hive': 'HKCU' if hive == winreg.HKEY_CURRENT_USER else 'HKLM',
                                'path': key_path
                            })
                            self.logger.debug(f"Found registry entry: {key_path}")
                    except FileNotFoundError:
                        pass
                    except Exception as e:
                        self.logger.debug(f"Registry error: {e}")
        except ImportError:
            pass
        
        return entries
    
    def _calculate_health_score(self) -> int:
        """
        Calculate system health score (0-100).
        
        100 = Clean (no Antigravity traces)
        0 = Fully infected (running with many leftovers)
        """
        score = 100
        
        # Deduct for installation
        if self.is_installed:
            score -= 30
        
        # Deduct for running process
        if self.is_running:
            score -= 20
        
        # Deduct for leftovers
        leftover_count = len(self.leftover_paths)
        if leftover_count > 0:
            score -= min(30, leftover_count * 5)
        
        # Deduct for registry entries
        registry_count = len(self.registry_entries)
        if registry_count > 0:
            score -= min(20, registry_count * 10)
        
        return max(0, score)
    
    def _get_status_text(self) -> str:
        """Get human-readable status text."""
        if self.health_score >= 90:
            return "✅ Clean - No Antigravity detected"
        elif self.health_score >= 70:
            return "🟡 Minor traces found"
        elif self.health_score >= 50:
            return "🟠 Moderate traces found"
        elif self.health_score >= 30:
            return "🔴 Significant traces detected"
        else:
            return "⛔ Heavy infection - Full cleanup needed"
    
    def _get_dir_size(self, path: str) -> int:
        """Get total size of directory."""
        total = 0
        try:
            if os.path.isfile(path):
                return os.path.getsize(path)
            
            for dirpath, dirnames, filenames in os.walk(path):
                for f in filenames:
                    fp = os.path.join(dirpath, f)
                    try:
                        total += os.path.getsize(fp)
                    except:
                        pass
        except:
            pass
        return total
    
    def _format_size(self, size: int) -> str:
        """Format size in human-readable format."""
        for unit in ['B', 'KB', 'MB', 'GB']:
            if size < 1024:
                return f"{size:.1f} {unit}"
            size /= 1024
        return f"{size:.1f} TB"
    
    def generate_report(self) -> str:
        """
        Generate a detailed text report.
        
        Returns:
            Formatted text report
        """
        results = self.detect_all()
        
        lines = [
            "=" * 60,
            "🔍 ANTIGRAVITY DETECTION REPORT",
            "=" * 60,
            f"📅 Generated: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}",
            f"💻 System: {platform.system()} {platform.release()}",
            "",
            "─" * 60,
            "📊 SUMMARY",
            "─" * 60,
            f"Health Score: {results['health_score']}/100",
            f"Status: {results['status']}",
            "",
            f"Installed: {'Yes' if results['is_installed'] else 'No'}",
            f"Running: {'Yes' if results['is_running'] else 'No'}",
            f"Leftover Files: {len(results['leftover_paths'])}",
            f"Total Size: {results['leftover_size_human']}",
        ]
        
        if results['is_installed']:
            lines.extend([
                "",
                "─" * 60,
                "📂 INSTALLATION",
                "─" * 60,
                f"Path: {results['install_path']}",
                f"Version: {results['version'] or 'Unknown'}",
            ])
        
        if results['leftover_paths']:
            lines.extend([
                "",
                "─" * 60,
                "🗂️ LEFTOVER FILES",
                "─" * 60,
            ])
            for path in results['leftover_paths'][:10]:  # Limit to 10
                lines.append(f"  • {path}")
            if len(results['leftover_paths']) > 10:
                lines.append(f"  ... and {len(results['leftover_paths']) - 10} more")
        
        if results['registry_entries']:
            lines.extend([
                "",
                "─" * 60,
                "🔑 REGISTRY ENTRIES (Windows)",
                "─" * 60,
            ])
            for entry in results['registry_entries']:
                lines.append(f"  • [{entry['hive']}] {entry['path']}")
        
        lines.extend([
            "",
            "=" * 60,
            "💡 RECOMMENDATION",
            "=" * 60,
        ])
        
        if results['health_score'] >= 90:
            lines.append("Your system is clean! No action needed.")
        elif results['health_score'] >= 70:
            lines.append("Run Quick Clean to remove minor traces.")
        elif results['health_score'] >= 50:
            lines.append("Run Deep Clean to fully remove Antigravity.")
        else:
            lines.append("Run Full Repair to completely clean your system.")
        
        lines.append("")
        
        return "\n".join(lines)
    
    def get_health_emoji(self) -> str:
        """Get emoji representing health score."""
        if self.health_score >= 90:
            return "✅"
        elif self.health_score >= 70:
            return "🟡"
        elif self.health_score >= 50:
            return "🟠"
        elif self.health_score >= 30:
            return "🔴"
        else:
            return "⛔"


# ==================== Test ====================

if __name__ == "__main__":
    logging.basicConfig(
        level=logging.DEBUG,
        format='[%(asctime)s] %(levelname)s: %(message)s'
    )
    
    detector = AntigravityDetector()
    print(detector.generate_report())
