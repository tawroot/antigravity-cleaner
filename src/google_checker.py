"""
Google Services Connectivity Checker
=====================================

Tests connectivity to Google's sensitive domains and services.
Useful for diagnosing Antigravity/Gemini connection issues.

License: MIT
"""

import socket
import ssl
import time
import platform
import subprocess
from typing import Dict, List, Tuple, Optional
from datetime import datetime
import logging
import threading

try:
    import urllib.request
    import urllib.error
except ImportError:
    urllib = None


class GoogleServicesChecker:
    """
    Tests connectivity to Google's critical services.
    
    Features:
    - Ping/Traceroute to Google domains
    - HTTPS connectivity test
    - SSL certificate validation
    - DNS resolution check
    - Response time measurement
    """
    
    # Google's sensitive/strict domains
    GOOGLE_SERVICES = {
        'accounts': {
            'name': 'Google Accounts',
            'domain': 'accounts.google.com',
            'icon': '🔐',
            'critical': True
        },
        'gemini': {
            'name': 'Gemini AI',
            'domain': 'gemini.google.com',
            'icon': '🤖',
            'critical': True
        },
        'aistudio': {
            'name': 'AI Studio',
            'domain': 'aistudio.google.com',
            'icon': '🧪',
            'critical': True
        },
        'cloud': {
            'name': 'Google Cloud',
            'domain': 'cloud.google.com',
            'icon': '☁️',
            'critical': True
        },
        'apis': {
            'name': 'Google APIs',
            'domain': 'www.googleapis.com',
            'icon': '🔌',
            'critical': True
        },
        'oauth': {
            'name': 'OAuth Server',
            'domain': 'oauth2.googleapis.com',
            'icon': '🔑',
            'critical': True
        },
        'play': {
            'name': 'Google Play',
            'domain': 'play.google.com',
            'icon': '🎮',
            'critical': False
        },
        'drive': {
            'name': 'Google Drive',
            'domain': 'drive.google.com',
            'icon': '📁',
            'critical': False
        },
        'youtube': {
            'name': 'YouTube',
            'domain': 'www.youtube.com',
            'icon': '📺',
            'critical': False
        },
        'google': {
            'name': 'Google Search',
            'domain': 'www.google.com',
            'icon': '🔍',
            'critical': False
        },
    }
    
    def __init__(self, logger: Optional[logging.Logger] = None, timeout: int = 10):
        """
        Initialize GoogleServicesChecker.
        
        Args:
            logger: Optional logger instance
            timeout: Connection timeout in seconds
        """
        self.logger = logger or logging.getLogger(__name__)
        self.timeout = timeout
        self.results = {}
        self.is_checking = False
        
        self.logger.info("GoogleServicesChecker initialized")
    
    def check_all_services(self, callback=None) -> Dict:
        """
        Check connectivity to all Google services.
        
        Args:
            callback: Optional callback function(service_key, result) for progress updates
        
        Returns:
            Dictionary with all results
        """
        self.logger.info("Starting connectivity check for all Google services...")
        self.is_checking = True
        self.results = {}
        
        for key, service in self.GOOGLE_SERVICES.items():
            if not self.is_checking:
                break
                
            result = self.check_service(service['domain'])
            result['service'] = service
            self.results[key] = result
            
            if callback:
                callback(key, result)
        
        self.is_checking = False
        
        # Calculate summary
        total = len(self.results)
        success = sum(1 for r in self.results.values() if r.get('https_ok'))
        critical_ok = sum(1 for k, r in self.results.items() 
                        if self.GOOGLE_SERVICES[k]['critical'] and r.get('https_ok'))
        critical_total = sum(1 for s in self.GOOGLE_SERVICES.values() if s['critical'])
        
        self.results['_summary'] = {
            'total': total,
            'success': success,
            'failed': total - success,
            'critical_ok': critical_ok,
            'critical_total': critical_total,
            'score': int((success / total) * 100) if total > 0 else 0,
            'timestamp': datetime.now().isoformat()
        }
        
        self.logger.info(f"Check complete: {success}/{total} services OK")
        return self.results
    
    def check_service(self, domain: str) -> Dict:
        """
        Check connectivity to a single service.
        
        Args:
            domain: Domain to check
        
        Returns:
            Dictionary with check results
        """
        result = {
            'domain': domain,
            'dns_ok': False,
            'dns_ip': None,
            'dns_time': None,
            'ping_ok': False,
            'ping_time': None,
            'https_ok': False,
            'https_time': None,
            'https_status': None,
            'ssl_ok': False,
            'ssl_issuer': None,
            'ssl_expiry': None,
            'error': None
        }
        
        # 1. DNS Resolution
        try:
            start = time.time()
            ip = socket.gethostbyname(domain)
            result['dns_time'] = round((time.time() - start) * 1000, 1)
            result['dns_ip'] = ip
            result['dns_ok'] = True
        except socket.gaierror as e:
            result['error'] = f"DNS failed: {e}"
            return result
        
        # 2. Ping test
        try:
            result['ping_ok'], result['ping_time'] = self._ping(domain)
        except Exception as e:
            self.logger.debug(f"Ping failed: {e}")
        
        # 3. HTTPS connectivity
        try:
            start = time.time()
            url = f"https://{domain}/"
            
            # Create SSL context
            context = ssl.create_default_context()
            
            req = urllib.request.Request(url, headers={'User-Agent': 'Mozilla/5.0'})
            response = urllib.request.urlopen(req, timeout=self.timeout, context=context)
            
            result['https_time'] = round((time.time() - start) * 1000, 1)
            result['https_status'] = response.getcode()
            result['https_ok'] = True
            
            # 4. SSL certificate info
            try:
                with socket.create_connection((domain, 443), timeout=self.timeout) as sock:
                    with context.wrap_socket(sock, server_hostname=domain) as ssock:
                        cert = ssock.getpeercert()
                        result['ssl_ok'] = True
                        result['ssl_issuer'] = dict(x[0] for x in cert.get('issuer', []))
                        result['ssl_expiry'] = cert.get('notAfter')
            except Exception as e:
                self.logger.debug(f"SSL info failed: {e}")
                
        except urllib.error.HTTPError as e:
            result['https_time'] = round((time.time() - start) * 1000, 1)
            result['https_status'] = e.code
            # Some codes are still "OK" (redirects, etc)
            result['https_ok'] = e.code < 500
            result['error'] = f"HTTP {e.code}"
        except urllib.error.URLError as e:
            result['error'] = f"URL error: {e.reason}"
        except Exception as e:
            result['error'] = str(e)
        
        return result
    
    def _ping(self, host: str) -> Tuple[bool, Optional[float]]:
        """
        Ping a host and return (success, time_ms).
        """
        param = '-n' if platform.system().lower() == 'windows' else '-c'
        timeout_param = '-w' if platform.system().lower() == 'windows' else '-W'
        
        try:
            output = subprocess.run(
                ['ping', param, '1', timeout_param, str(self.timeout * 1000 if platform.system().lower() == 'windows' else self.timeout), host],
                capture_output=True,
                text=True,
                timeout=self.timeout + 2
            )
            
            if output.returncode == 0:
                # Parse time from output
                import re
                if platform.system().lower() == 'windows':
                    match = re.search(r'time[=<](\d+)ms', output.stdout)
                else:
                    match = re.search(r'time=(\d+\.?\d*)', output.stdout)
                
                if match:
                    return True, float(match.group(1))
                return True, None
            
            return False, None
        except Exception:
            return False, None
    
    def traceroute(self, domain: str, callback=None) -> List[Dict]:
        """
        Run traceroute to a domain.
        
        Args:
            domain: Target domain
            callback: Optional callback for each hop
        
        Returns:
            List of hops with timing info
        """
        hops = []
        
        if platform.system().lower() == 'windows':
            cmd = ['tracert', '-d', '-w', '1000', '-h', '20', domain]
        else:
            cmd = ['traceroute', '-n', '-w', '1', '-m', '20', domain]
        
        try:
            process = subprocess.Popen(
                cmd,
                stdout=subprocess.PIPE,
                stderr=subprocess.PIPE,
                text=True
            )
            
            import re
            hop_num = 0
            
            for line in process.stdout:
                line = line.strip()
                if not line:
                    continue
                
                # Parse hop
                if platform.system().lower() == 'windows':
                    # Windows: "  1    <1 ms    <1 ms    <1 ms  192.168.1.1"
                    match = re.match(r'\s*(\d+)\s+(.+)', line)
                else:
                    # Linux/Mac: " 1  192.168.1.1  0.5 ms  0.4 ms  0.3 ms"
                    match = re.match(r'\s*(\d+)\s+(.+)', line)
                
                if match:
                    hop_num = int(match.group(1))
                    hop_data = match.group(2)
                    
                    hop = {
                        'hop': hop_num,
                        'raw': hop_data,
                        'ip': None,
                        'time': None
                    }
                    
                    # Extract IP
                    ip_match = re.search(r'(\d+\.\d+\.\d+\.\d+)', hop_data)
                    if ip_match:
                        hop['ip'] = ip_match.group(1)
                    
                    # Extract time
                    time_match = re.search(r'(\d+\.?\d*)\s*ms', hop_data)
                    if time_match:
                        hop['time'] = float(time_match.group(1))
                    
                    hops.append(hop)
                    
                    if callback:
                        callback(hop)
            
            process.wait(timeout=60)
            
        except Exception as e:
            self.logger.error(f"Traceroute failed: {e}")
        
        return hops
    
    def get_status_emoji(self, result: Dict) -> str:
        """Get status emoji for a result."""
        if result.get('https_ok'):
            return "✅"
        elif result.get('dns_ok'):
            return "⚠️"
        else:
            return "❌"
    
    def generate_report(self) -> str:
        """Generate a text report of results."""
        if not self.results:
            return "No results. Run check_all_services() first."
        
        lines = [
            "=" * 60,
            "🌐 GOOGLE SERVICES CONNECTIVITY REPORT",
            "=" * 60,
            f"📅 {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}",
            ""
        ]
        
        summary = self.results.get('_summary', {})
        if summary:
            lines.extend([
                f"📊 Score: {summary['score']}%",
                f"✅ Success: {summary['success']}/{summary['total']}",
                f"🔴 Critical: {summary['critical_ok']}/{summary['critical_total']}",
                ""
            ])
        
        lines.append("─" * 60)
        
        for key, service in self.GOOGLE_SERVICES.items():
            result = self.results.get(key, {})
            status = self.get_status_emoji(result)
            
            line = f"{status} {service['icon']} {service['name']}"
            if result.get('https_time'):
                line += f" ({result['https_time']}ms)"
            if result.get('error'):
                line += f" - {result['error']}"
            
            lines.append(line)
        
        lines.extend(["", "=" * 60])
        
        return "\n".join(lines)
    
    def stop(self):
        """Stop ongoing check."""
        self.is_checking = False


# ==================== Test ====================

if __name__ == "__main__":
    logging.basicConfig(level=logging.DEBUG)
    
    checker = GoogleServicesChecker()
    
    def on_progress(key, result):
        print(f"{checker.get_status_emoji(result)} {key}: {result.get('https_time', 'N/A')}ms")
    
    results = checker.check_all_services(callback=on_progress)
    print("\n" + checker.generate_report())
