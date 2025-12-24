import os
import sys
import platform
import shutil
import subprocess
import time
import glob
import logging
from logging.handlers import RotatingFileHandler
from datetime import datetime

# Import new modules
try:
    from browser_helper import BrowserHelper
    from network_optimizer import NetworkOptimizer
    from session_manager import SessionManager
except ImportError as e:
    # Modules not yet available, will be created
    BrowserHelper = None
    NetworkOptimizer = None
    SessionManager = None

# Try imports for runtime (UI and Process handling)
try:
    import psutil
    from rich.console import Console
    from rich.progress import Progress, SpinnerColumn, TextColumn
    from rich.panel import Panel
    from rich.prompt import Prompt, Confirm
    from rich.table import Table
    from rich.style import Style
    from rich import print as rprint
except ImportError:
    print("Missing dependencies. Please run: pip install -r requirements.txt")
    sys.exit(1)

# Platform check
CURRENT_OS = platform.system()
IS_WINDOWS = CURRENT_OS == "Windows"
IS_MAC = CURRENT_OS == "Darwin"
IS_LINUX = CURRENT_OS == "Linux"

if IS_WINDOWS:
    import winreg

# Setup Console
console = Console()

# --- Configuration & Constants ---

APP_NAME = "Antigravity"
LOG_FILE = os.path.join(os.path.expanduser("~"), "Desktop", "Antigravity-Cleaner.log")

class Cleaner:
    def __init__(self):
        self.dry_run = False
        self.found_items = []

    def log(self, message, style="dim"):
        """Log to file and console."""
        timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        log_entry = f"[{timestamp}] {message}"
        
        # Write to file
        with open(LOG_FILE, "a", encoding="utf-8") as f:
            f.write(log_entry + "\n")
            
        # Write to console (fancy)
        console.print(f"[{style}]{message}[/{style}]")

    def get_user_confirmation(self, question):
        return Confirm.ask(question)

    # --- Scanning Logic ---

    def scan_processes(self):
        """Check if Antigravity is running."""
        self.log("Scanning for running processes...", style="cyan")
        running = []
        for proc in psutil.process_iter(['pid', 'name']):
            try:
                if APP_NAME.lower() in proc.info['name'].lower():
                    running.append(proc)
            except (psutil.NoSuchProcess, psutil.AccessDenied, psutil.ZombieProcess):
                pass
        return running

    def kill_processes(self, processes):
        if not processes:
            return
        
        self.log(f"Found {len(processes)} running Antigravity processes.", style="yellow")
        if self.dry_run:
            self.log("[Dry Run] Would terminate processes.", style="yellow")
            return

        for proc in processes:
            try:
                proc.kill()
                self.log(f"Killed process {proc.info['name']} (PID: {proc.info['pid']})", style="green")
            except Exception as e:
                self.log(f"Failed to kill {proc.info['name']}: {e}", style="red")

    def find_uninstallers_windows(self):
        """Find uninstall strings in Windows Registry."""
        self.log("Scanning Windows Registry for uninstallers...", style="cyan")
        uninstallers = []
        roots = [
            (winreg.HKEY_CURRENT_USER, r"SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall"),
            (winreg.HKEY_LOCAL_MACHINE, r"SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall"),
            (winreg.HKEY_LOCAL_MACHINE, r"SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall"),
        ]

        for hive, path in roots:
            try:
                with winreg.OpenKey(hive, path) as key:
                    for i in range(0, winreg.QueryInfoKey(key)[0]):
                        try:
                            subkey_name = winreg.EnumKey(key, i)
                            with winreg.OpenKey(key, subkey_name) as subkey:
                                try:
                                    display_name = winreg.QueryValueEx(subkey, "DisplayName")[0]
                                    if APP_NAME.lower() in display_name.lower():
                                        uninstall_string = winreg.QueryValueEx(subkey, "UninstallString")[0]
                                        uninstallers.append({
                                            "name": display_name,
                                            "key": path + "\\" + subkey_name,
                                            "cmd": uninstall_string
                                        })
                                except FileNotFoundError:
                                    pass
                        except OSError:
                            continue
            except OSError:
                continue
        return uninstallers

    def get_cleanup_paths(self, deep=False):
        """Return list of paths to check based on OS."""
        paths = []
        home = os.path.expanduser("~")

        if IS_WINDOWS:
            local_appdata = os.environ.get("LOCALAPPDATA", os.path.join(home, "AppData", "Local"))
            appdata = os.environ.get("APPDATA", os.path.join(home, "AppData", "Roaming"))
            temp = os.environ.get("TEMP", os.path.join(local_appdata, "Temp"))
            
            paths.extend([
                os.path.join(local_appdata, "Programs", "Antigravity"),
                os.path.join(local_appdata, "Antigravity"),
                os.path.join(appdata, "Antigravity"),
                os.path.join(appdata, "Google", "Antigravity"), # Based on legacy script
                os.path.join(local_appdata, "Google", "Antigravity"),
            ])
            
            if deep:
                paths.extend([
                    os.path.join(temp, "antigravity-stable-user-x64"),
                    os.path.join(temp, "is-*.tmp"), # Inno Setup temp files
                    # Chrome Extension Trace (Wildcard structure handled by expand_globs)
                    os.path.join(local_appdata, "Google", "Chrome", "User Data", "*", "Extensions", "*", "*", "*antigravity*"),
                    # Python Lib Trace
                    os.path.join(local_appdata, "Python", "pythoncore-*", "Lib", "antigravity.py")
                ])

        elif IS_MAC:
            paths.extend([
                os.path.join(home, "Library", "Application Support", "Antigravity"),
                os.path.join(home, "Library", "Caches", "Antigravity"),
                os.path.join(home, "Library", "Preferences", "com.antigravity.plist"), # Hypothetical
                os.path.join(home, "Library", "Saved Application State", "com.antigravity.savedState"),
                "/Applications/Antigravity.app"
            ])

        elif IS_LINUX:
            paths.extend([
                os.path.join(home, ".config", "Antigravity"),
                os.path.join(home, ".local", "share", "Antigravity"),
                os.path.join(home, ".cache", "Antigravity"),
            ])

        return paths

    def expand_globs(self, paths):
        """Expand wildcard paths."""
        expanded = []
        for p in paths:
            # Simple check if it's a glob pattern or exact path
            if "*" in p:
                expanded.extend(glob.glob(p))
            else:
                expanded.append(p)
        return list(set(expanded)) # Unique

    def clean_paths(self, paths):
        found_any = False
        for p in paths:
            if os.path.exists(p):
                found_any = True
                if self.dry_run:
                    self.log(f"[Dry Run] Would remove: {p}", style="yellow")
                else:
                    try:
                        if os.path.isdir(p):
                            shutil.rmtree(p)
                        else:
                            os.remove(p)
                        self.log(f"Removed: {p}", style="green")
                    except Exception as e:
                        self.log(f"Error removing {p}: {e}", style="red")
        
        if not found_any:
            self.log("No leftover files found in standard paths.", style="dim")


    def run_windows_uninstallers(self, uninstallers):
        for item in uninstallers:
            cmd = item['cmd']
            self.log(f"Running uninstaller for: {item['name']}", style="bold white")
            if self.dry_run:
                self.log(f"[Dry Run] CMD: {cmd}", style="yellow")
                continue

            # Attempt to parse quiet flags
            # This is heuristic based on the legacy script
            final_cmd = cmd
            args = []
            
            if "msiexec" in cmd.lower():
                args = ["/qn", "/norestart"]
                # We need to restructure for subprocess
                # cmd usually: msiexec /x {GUID}
                parts = cmd.split()
                exe = parts[0]
                arguments = parts[1:] + args
                try:
                    subprocess.run([exe] + arguments, check=True)
                    self.log("MSI Uninstall complete.", style="green")
                except subprocess.CalledProcessError as e:
                    self.log(f"Uninstall failed: {e}", style="red")
            
            elif "unins" in cmd.lower() and ".exe" in cmd.lower():
                 # Ex: "C:\path\unins000.exe"
                 # We simply run it with silent flags
                 # Need to extract the exe path carefully if quoted
                 import shlex
                 parts = shlex.split(cmd)
                 exe = parts[0]
                 # Common InnoSetup/NSIS silent flags
                 silent_args = ["/VERYSILENT", "/SUPPRESSMSGBOXES", "/NORESTART"]
                 try:
                    subprocess.run([exe] + silent_args, check=True)
                    self.log("Uninstaller executed successfully.", style="green")
                 except subprocess.CalledProcessError as e:
                    self.log(f"Uninstaller failed: {e}", style="red")
            else:
                self.log(f"Unknown uninstaller type. Running manually: {cmd}", style="yellow")
                subprocess.run(cmd, shell=True)


    def network_reset(self):
        self.log("Resetting Network Settings...", style="bold magenta")
        
        commands = []
        if IS_WINDOWS:
            commands = [
                "ipconfig /flushdns",
                "netsh winsock reset",
                "netsh int ip reset"
            ]
        elif IS_MAC:
            commands = [
                "dscacheutil -flushcache",
                "killall -HUP mDNSResponder"
            ]
        elif IS_LINUX:
            # Distro dependent, try systemd-resolve or just simple flush
            commands = [
                "resolvectl flush-caches" 
            ]

        if self.dry_run:
            for cmd in commands:
                self.log(f"[Dry Run] Would run: {cmd}", style="yellow")
        else:
            for cmd in commands:
                self.log(f"Executing: {cmd}", style="cyan")
                try:
                    subprocess.run(cmd, shell=True, check=False) # Check false to ignore errors on missing cmds
                except Exception as e:
                    self.log(f"Error running {cmd}: {e}", style="red")
            self.log("Network reset complete. Restart recommended.", style="green")


    # --- Main Actions ---

    def run_clean(self, deep=False):
        # 1. Check processes
        procs = self.scan_processes()
        if procs:
            if self.dry_run or self.get_user_confirmation(f"Found {len(procs)} running instances. Kill them?"):
                self.kill_processes(procs)

        # 2. Uninstall (Windows only usually has registry uninstallers)
        if IS_WINDOWS:
            uninstallers = self.find_uninstallers_windows()
            if uninstallers:
                self.log(f"Found {len(uninstallers)} matching uninstallers.", style="bold white")
                if self.dry_run:
                    for u in uninstallers: self.log(f" - {u['name']}", style="dim")
                else:
                    if self.get_user_confirmation("Run uninstallers first?"):
                        self.run_windows_uninstallers(uninstallers)
            else:
                self.log("No uninstallers found in registry.", style="dim")

        # 3. Clean files
        self.log("Scanning for leftovers...", style="bold white")
        target_paths = self.get_cleanup_paths(deep=deep)
        target_paths = self.expand_globs(target_paths)
        
        # Filter existing
        existing = [p for p in target_paths if os.path.exists(p) or glob.glob(p)]
        
        if existing:
            self.log(f"Found {len(existing)} locations to clean.", style="yellow")
            self.clean_paths(existing)
        else:
            self.log("No leftovers found.", style="green")

        if deep:
             self.log("Deep scan complete.", style="bold green")

    def run_network_reset(self):
        self.network_reset()


# --- Agent Logging Setup ---

def setup_agent_logging():
    """Setup detailed logging for agent operations"""
    log_dir = os.path.join(os.path.dirname(__file__), '..', '.agent', 'logs')
    os.makedirs(log_dir, exist_ok=True)
    
    log_file = os.path.join(log_dir, 'browser-helper-operations.log')
    
    logger = logging.getLogger('antigravity_agent')
    logger.setLevel(logging.DEBUG)
    
    # Rotating file handler (10MB max, keep 3 backups)
    handler = RotatingFileHandler(log_file, maxBytes=10*1024*1024, backupCount=3)
    handler.setFormatter(logging.Formatter(
        '[%(asctime)s] [%(levelname)s] %(message)s',
        datefmt='%Y-%m-%d %H:%M:%S'
    ))
    
    logger.addHandler(handler)
    return logger


# --- Browser Login Helper Submenu ---

def browser_login_helper_menu(browser_helper, network_optimizer, logger):
    """Browser Login Helper submenu"""
    while True:
        console.print("\n" + "="*70)
        console.print("[bold cyan]BROWSER LOGIN HELPER | کمک‌کننده ورود[/bold cyan]")
        console.print("="*70)
        console.print("\n1. [green]Clean Antigravity Browser Traces (Safe)[/green]")
        console.print("   [dim]پاک‌سازی ردهای Antigravity در مرورگر (ایمن)[/dim]")
        console.print("\n2. [yellow]Optimize Network for Login[/yellow]")
        console.print("   [dim]بهینه‌سازی شبکه برای ورود[/dim]")
        console.print("\n3. [magenta]Network Diagnostic Report[/magenta]")
        console.print("   [dim]گزارش تشخیصی شبکه[/dim]")
        console.print("\n4. [cyan]Run Full Login Repair (1+2)[/cyan]")
        console.print("   [dim]اجرای تعمیر کامل ورود[/dim]")
        console.print("\n0. [dim]Back to Main Menu[/dim]")
        
        choice = Prompt.ask("\nEnter choice", choices=["0", "1", "2", "3", "4"], default="0")
        
        if choice == "0":
            break
        elif choice == "1":
            # Clean browser traces
            browsers = browser_helper.detect_installed_browsers()
            if not browsers:
                console.print("[red]No supported browsers found.[/red]")
                continue
            
            console.print(f"\n[cyan]Found browsers: {', '.join(browsers)}[/cyan]")
            browser = Prompt.ask("Select browser to clean", choices=browsers + ["all"], default="all")
            
            if browser == "all":
                for b in browsers:
                    stats = browser_helper.clean_browser_completely(b)
                    console.print(f"\n[green]✓ {b}: {stats['cookies']} cookies, {stats['cache']} cache items cleaned[/green]")
            else:
                stats = browser_helper.clean_browser_completely(browser)
                console.print(f"\n[green]✓ Cleaned {stats['cookies']} cookies, {stats['cache']} cache items[/green]")
        
        elif choice == "2":
            # Optimize network
            console.print("\n[cyan]Optimizing network settings...[/cyan]")
            network_optimizer.clear_dns_cache()
            if IS_WINDOWS:
                if Confirm.ask("Reset network stack? (Requires restart)"):
                    network_optimizer.reset_network_stack()
            console.print("[green]✓ Network optimization complete[/green]")
        
        elif choice == "3":
            # Diagnostic report
            console.print("\n[cyan]Generating diagnostic report...[/cyan]")
            report = network_optimizer.generate_diagnostic_report()
            console.print("\n" + report)
            
            # Save to file
            report_file = os.path.join(os.path.expanduser("~"), "Desktop", "Antigravity-Network-Diagnostic.txt")
            with open(report_file, 'w', encoding='utf-8') as f:
                f.write(report)
            console.print(f"\n[green]✓ Report saved to: {report_file}[/green]")
        
        elif choice == "4":
            # Full repair
            console.print("\n[cyan]Running full login repair...[/cyan]")
            
            # Clean browsers
            browsers = browser_helper.detect_installed_browsers()
            for b in browsers:
                stats = browser_helper.clean_browser_completely(b)
                console.print(f"[green]✓ {b}: {stats['cookies']} cookies cleaned[/green]")
            
            # Optimize network
            network_optimizer.clear_dns_cache()
            console.print("[green]✓ DNS cache cleared[/green]")
            
            # Diagnostic
            console.print("\n[cyan]Running diagnostics...[/cyan]")
            connectivity = network_optimizer.test_google_connectivity()
            console.print(f"[green]✓ Google connectivity: {connectivity['overall_status']}[/green]")
            
            console.print("\n[bold green]✓ Full login repair complete![/bold green]")
        
        if choice != "0":
            if not Confirm.ask("\nContinue in Browser Helper?"):
                break


# --- Session Manager Submenu ---

def session_manager_menu(session_manager, browser_helper, logger):
    """Session Manager submenu with email search support"""
    while True:
        console.print("\n" + "="*70)
        console.print("[bold green]SESSION MANAGER | مدیریت نشست‌ها[/bold green]")
        console.print("="*70)
        console.print("\n1. [cyan]Backup Current Session[/cyan]")
        console.print("   [dim]پشتیبان‌گیری از Session فعلی[/dim]")
        console.print("\n2. [yellow]Restore Saved Session[/yellow]")
        console.print("   [dim]بازیابی Session ذخیره‌شده[/dim]")
        console.print("\n3. [magenta]List All Saved Sessions[/magenta]")
        console.print("   [dim]لیست تمام Session های ذخیره‌شده[/dim]")
        console.print("\n4. [red]Delete Old Sessions[/red]")
        console.print("   [dim]حذف Session های قدیمی[/dim]")
        console.print("\n5. [blue]Search Profiles by Email[/blue]")
        console.print("   [dim]جستجوی پروفایل بر اساس ایمیل[/dim]")
        console.print("\n0. [dim]Back to Main Menu[/dim]")
        
        choice = Prompt.ask("\nEnter choice", choices=["0", "1", "2", "3", "4", "5"], default="0")
        
        if choice == "0":
            break
        elif choice == "1":
            # Backup session with email search
            browsers = browser_helper.detect_installed_browsers()
            if not browsers:
                console.print("[red]No supported browsers found.[/red]")
                continue
            
            console.print(f"\n[cyan]Found browsers: {', '.join(browsers)}[/cyan]")
            browser = Prompt.ask("Select browser", choices=browsers)
            
            # Check if browser is running
            if browser_helper.is_browser_running(browser):
                console.print(f"\n[yellow]⚠ {browser.capitalize()} is currently running.[/yellow]")
                console.print("[dim]Browsers lock their session data while running. To backup, the browser must be closed.[/dim]")
                if Confirm.ask(f"Close {browser.capitalize()} now? (Recommended)", default=True):
                    browser_helper.close_browser_gracefully(browser)
                    # Force kill if still running
                    if browser_helper.is_browser_running(browser):
                        browser_helper.kill_browser_processes(browser)
            
            # Ask if user wants to search by email
            search_by_email = Confirm.ask("Search profile by email?", default=True)
            
            if search_by_email:
                email_query = Prompt.ask("Enter email (or part of it)")
                profiles = browser_helper.search_profiles_by_email(browser, email_query)
            else:
                profiles = browser_helper.get_browser_profiles_with_email(browser)
            
            if not profiles:
                console.print("[red]No profiles found.[/red]")
                continue
            
            # Display profiles with emails
            console.print("\n[cyan]Available profiles:[/cyan]")
            table = Table()
            table.add_column("#", style="dim", width=4)
            table.add_column("Profile", style="cyan")
            table.add_column("Email", style="yellow")
            
            for i, p in enumerate(profiles, 1):
                email_display = p.get('email') or '[dim]No email[/dim]'
                table.add_row(str(i), p['name'], email_display)
            
            console.print(table)
            
            # Select profile
            profile_idx = int(Prompt.ask("Select profile number", 
                                         choices=[str(i) for i in range(1, len(profiles)+1)]))
            selected_profile = profiles[profile_idx - 1]
            
            session_name = Prompt.ask("Session name (optional, press Enter for auto)", default="")
            if not session_name:
                # Auto-generate with email if available
                if selected_profile.get('email'):
                    email_prefix = selected_profile['email'].split('@')[0]
                    session_name = f"{browser}_{email_prefix}_{datetime.now().strftime('%Y%m%d')}"
                else:
                    session_name = None
            
            backup_result = session_manager.backup_session(browser, selected_profile['path'], session_name)
            if backup_result:
                console.print(f"[green]✓ Session backed up successfully![/green]")
                console.print(f"[dim]Path: {backup_result}[/dim]")
            else:
                console.print("[red]✗ Session backup failed.[/red]")

        
        elif choice == "2":
            # Restore session
            sessions = session_manager.list_saved_sessions()
            if not sessions:
                console.print("[red]No saved sessions found.[/red]")
                continue
            
            console.print("\n[cyan]Saved sessions:[/cyan]")
            for i, s in enumerate(sessions, 1):
                status = "[red](expired)[/red]" if s.get('expired') else "[green](valid)[/green]"
                console.print(f"{i}. {s['name']} - {s['browser']} - {s['cookie_count']} cookies {status}")
            
            session_idx = int(Prompt.ask("Select session number", choices=[str(i) for i in range(1, len(sessions)+1)]))
            selected_session = sessions[session_idx - 1]
            
            # Get browser and profile with email search
            browsers = browser_helper.detect_installed_browsers()
            browser = Prompt.ask("Select browser", choices=browsers, default=selected_session['browser'])
            
            # Check if browser is running
            if browser_helper.is_browser_running(browser):
                console.print(f"\n[yellow]⚠ {browser.capitalize()} is currently running.[/yellow]")
                console.print("[dim]Browsers lock their session data while running. To restore, the browser must be closed.[/dim]")
                if Confirm.ask(f"Close {browser.capitalize()} now? (Recommended)", default=True):
                    browser_helper.close_browser_gracefully(browser)
                    if browser_helper.is_browser_running(browser):
                        browser_helper.kill_browser_processes(browser)

            # Ask if user wants to search by email
            search_by_email = Confirm.ask("Search profile by email?", default=True)
            
            if search_by_email:
                email_query = Prompt.ask("Enter email (or part of it)")
                profiles = browser_helper.search_profiles_by_email(browser, email_query)
            else:
                profiles = browser_helper.get_browser_profiles_with_email(browser)
            
            if not profiles:
                console.print("[red]No profiles found.[/red]")
                continue
            
            # Display profiles with emails
            console.print("\n[cyan]Available profiles:[/cyan]")
            table = Table()
            table.add_column("#", style="dim", width=4)
            table.add_column("Profile", style="cyan")
            table.add_column("Email", style="yellow")
            
            for i, p in enumerate(profiles, 1):
                email_display = p.get('email') or '[dim]No email[/dim]'
                table.add_row(str(i), p['name'], email_display)
            
            console.print(table)
            
            profile_idx = int(Prompt.ask("Select profile number", 
                                         choices=[str(i) for i in range(1, len(profiles)+1)]))
            selected_profile = profiles[profile_idx - 1]
            
            if session_manager.restore_session(selected_session['name'], browser, selected_profile['path']):
                console.print("[green]✓ Session restored successfully![/green]")
            else:
                console.print("[red]✗ Session restore failed.[/red]")
        
        elif choice == "3":
            # List sessions
            sessions = session_manager.list_saved_sessions()
            if not sessions:
                console.print("[yellow]No saved sessions found.[/yellow]")
                continue
            
            table = Table(title="Saved Sessions")
            table.add_column("Name", style="cyan")
            table.add_column("Browser", style="magenta")
            table.add_column("Backup Time", style="yellow")
            table.add_column("Cookies", justify="right", style="green")
            table.add_column("Status", style="white")
            
            for s in sessions:
                status = "[red]Expired[/red]" if s.get('expired') else "[green]Valid[/green]"
                table.add_row(
                    s['name'],
                    s['browser'],
                    s['backup_time'],
                    str(s['cookie_count']),
                    status
                )
            
            console.print(table)
        
        elif choice == "4":
            # Delete old sessions
            if Confirm.ask("Delete all expired sessions?"):
                count = session_manager.delete_expired_sessions()
                console.print(f"[green]✓ Deleted {count} expired sessions[/green]")
        
        elif choice == "5":
            # Search profiles by email
            browsers = browser_helper.detect_installed_browsers()
            if not browsers:
                console.print("[red]No supported browsers found.[/red]")
                continue
            
            console.print(f"\n[cyan]Found browsers: {', '.join(browsers)}[/cyan]")
            browser = Prompt.ask("Select browser", choices=browsers + ["all"], default="all")
            
            email_query = Prompt.ask("Enter email to search (or part of it)")
            
            if browser == "all":
                all_profiles = []
                for b in browsers:
                    profiles = browser_helper.search_profiles_by_email(b, email_query)
                    for p in profiles:
                        p['browser'] = b
                    all_profiles.extend(profiles)
                profiles = all_profiles
            else:
                profiles = browser_helper.search_profiles_by_email(browser, email_query)
                for p in profiles:
                    p['browser'] = browser
            
            if not profiles:
                console.print(f"[yellow]No profiles found matching '{email_query}'[/yellow]")
                continue
            
            console.print(f"\n[green]Found {len(profiles)} matching profiles:[/green]")
            table = Table()
            table.add_column("#", style="dim", width=4)
            table.add_column("Browser", style="magenta")
            table.add_column("Profile", style="cyan")
            table.add_column("Email", style="yellow")
            
            for i, p in enumerate(profiles, 1):
                email_display = p.get('email') or '[dim]No email[/dim]'
                table.add_row(str(i), p.get('browser', ''), p['name'], email_display)
            
            console.print(table)
            
            # Ask if user wants to backup
            if Confirm.ask("Backup one of these profiles?"):
                profile_idx = int(Prompt.ask("Select profile number", 
                                             choices=[str(i) for i in range(1, len(profiles)+1)]))
                selected_profile = profiles[profile_idx - 1]
                
                session_name = Prompt.ask("Session name (optional)", default="")
                if not session_name:
                    if selected_profile.get('email'):
                        email_prefix = selected_profile['email'].split('@')[0]
                        session_name = f"{selected_profile['browser']}_{email_prefix}_{datetime.now().strftime('%Y%m%d')}"
                    else:
                        session_name = None
                
                # Check if browser is running
                target_browser = selected_profile['browser']
                if browser_helper.is_browser_running(target_browser):
                    console.print(f"\n[yellow]⚠ {target_browser.capitalize()} is currently running.[/yellow]")
                    if Confirm.ask(f"Close {target_browser.capitalize()} now to allow backup?", default=True):
                        browser_helper.close_browser_gracefully(target_browser)
                        if browser_helper.is_browser_running(target_browser):
                            browser_helper.kill_browser_processes(target_browser)

                backup_result = session_manager.backup_session(selected_profile['browser'], selected_profile['path'], session_name)
                if backup_result:
                    console.print("[green]✓ Session backed up successfully![/green]")
                    console.print(f"[dim]Path: {backup_result}[/dim]")
                else:
                    console.print("[red]✗ Session backup failed.[/red]")

        
        if choice != "0":
            if not Confirm.ask("\nContinue in Session Manager?"):
                break


# --- CLI Menu ---

def main():
    # Setup logging
    agent_logger = setup_agent_logging()
    agent_logger.info("=== Antigravity Cleaner Started ===")
    
    cleaner = Cleaner()
    
    # Initialize new helpers (if modules available)
    browser_helper = None
    network_optimizer = None
    session_manager = None
    
    if BrowserHelper and NetworkOptimizer and SessionManager:
        try:
            browser_helper = BrowserHelper(agent_logger, dry_run=cleaner.dry_run)
            network_optimizer = NetworkOptimizer(agent_logger, dry_run=cleaner.dry_run)
            session_storage = os.path.join(os.path.expanduser('~'), '.antigravity-cleaner', 'sessions')
            session_manager = SessionManager(session_storage, agent_logger, dry_run=cleaner.dry_run)
            agent_logger.info("Browser helper modules initialized successfully")
        except Exception as e:
            agent_logger.error(f"Failed to initialize browser helper modules: {e}")
            console.print(f"[yellow]Warning: Browser helper features unavailable: {e}[/yellow]")
    
    # Check args
    if len(sys.argv) > 1:
        arg = sys.argv[1].lower()
        if arg == "--dry-run":
            cleaner.dry_run = True
            console.print(Panel.fit("DRY RUN MODE ENABLED", style="bold yellow"))
            # Update helpers dry_run mode
            if browser_helper:
                browser_helper.dry_run = True
            if network_optimizer:
                network_optimizer.dry_run = True
            if session_manager:
                session_manager.dry_run = True
        elif arg == "--auto":
            cleaner.run_clean(deep=True)
            cleaner.run_network_reset()
            sys.exit(0)

    # Header
    grid = Table.grid(expand=True)
    grid.add_column(justify="center", ratio=1)
    grid.add_row(f"[bold cyan]ANTIGRAVITY CLEANER[/bold cyan] v{platform.python_version()}")
    grid.add_row(f"[dim]Running on {CURRENT_OS}[/dim]")
    grid.add_row(f"[dim]Log: {LOG_FILE}[/dim]")
    console.print(Panel(grid, style="blue", border_style="blue"))

    while True:
        console.print("\n[bold white]Select an Option:[/bold white]")
        console.print("1. [green]Quick Clean[/green] (Standard paths)")
        console.print("2. [yellow]Deep Clean[/yellow] (Aggressive scan + Temp)")
        console.print("3. [magenta]Network Reset[/magenta] (Fix connection issues)")
        console.print("4. [cyan]Full Repair[/cyan] (Deep Clean + Network Reset)")
        console.print("5. [dim]Toggle Dry Run[/dim] " + (f"(Currently: [bold red]ON[/bold red])" if cleaner.dry_run else "(Currently: OFF)"))
        
        # New options (if modules available)
        if browser_helper and network_optimizer:
            console.print("6. [blue]Browser Login Helper[/blue] (Clean browser traces)")
        if session_manager:
            console.print("7. [green]Session Manager[/green] (Backup/Restore sessions)")
        
        console.print("0. Exit")

        choices = ["0", "1", "2", "3", "4", "5"]
        if browser_helper and network_optimizer:
            choices.append("6")
        if session_manager:
            choices.append("7")
        
        choice = Prompt.ask("Enter choice", choices=choices, default="0")

        if choice == "0":
            agent_logger.info("=== Antigravity Cleaner Exited ===")
            sys.exit(0)
        elif choice == "1":
            cleaner.run_clean(deep=False)
        elif choice == "2":
            cleaner.run_clean(deep=True)
        elif choice == "3":
            cleaner.run_network_reset()
        elif choice == "4":
            cleaner.run_clean(deep=True)
            cleaner.run_network_reset()
        elif choice == "5":
            cleaner.dry_run = not cleaner.dry_run
            # Update helpers dry_run mode
            if browser_helper:
                browser_helper.dry_run = cleaner.dry_run
            if network_optimizer:
                network_optimizer.dry_run = cleaner.dry_run
            if session_manager:
                session_manager.dry_run = cleaner.dry_run
            status = "[bold red]ON[/bold red]" if cleaner.dry_run else "OFF"
            console.print(f"Dry Run is now {status}")
        elif choice == "6" and browser_helper and network_optimizer:
            browser_login_helper_menu(browser_helper, network_optimizer, agent_logger)
        elif choice == "7" and session_manager:
            session_manager_menu(session_manager, browser_helper, agent_logger)

        if choice != "5":
            if not Confirm.ask("Run another task?"):
                break

if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        console.print("\n[red]Cancelled by user.[/red]")
        sys.exit(0)
