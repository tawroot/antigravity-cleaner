package doctor

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/tawroot/antigravity-cleaner/pkg/patcher"
	"github.com/tawroot/antigravity-cleaner/pkg/proxy"
)

// DiagnosticItem represents a friendly diagnostic check
type DiagnosticItem struct {
	Name         string
	Status       string   // "OK", "WARN", "FAIL"
	HumanSummary string   // Short friendly text for regular users
	DebugDetails string   // Raw technical string for --verbose
	TargetPaths  []string // Explicit paths (App path, Core path, Backup) to render directly in markdown!
}

// RunDiagnostics executes concurrent, fast, non-intimidating diagnostics
func RunDiagnostics(verbose bool) []DiagnosticItem {
	var items []DiagnosticItem
	var mu sync.Mutex
	var wg sync.WaitGroup

	// 0. System Environment
	osInfo := DetectOSInfo()
	items = append(items, DiagnosticItem{
		Name:         "OS Environment",
		Status:       "OK",
		HumanSummary: osInfo,
		DebugDetails: fmt.Sprintf("OS: %s, Arch: %s, CPU: %d cores", runtime.GOOS, runtime.GOARCH, runtime.NumCPU()),
	})

	// 1. Antigravity Installation & Version
	prodName, appVersion := DetectAntigravityVersion()
	appPath := proxy.FindAntigravityExecutable()
	if appPath != "" {
		items = append(items, DiagnosticItem{
			Name:         "Antigravity Build",
			Status:       "OK",
			HumanSummary: fmt.Sprintf("%s (v%s)", prodName, appVersion),
			DebugDetails: appPath,
			TargetPaths:  []string{fmt.Sprintf("App Path: `%s`", appPath)},
		})
	} else {
		items = append(items, DiagnosticItem{
			Name:         "Antigravity Build",
			Status:       "FAIL",
			HumanSummary: "NOT INSTALLED on this machine!",
			DebugDetails: "Please download and install Google Antigravity first from https://antigravity.google",
			TargetPaths:  []string{"Note: Antigravity Cleaner is a patcher. Google Antigravity IDE itself must be installed first."},
		})
	}

	// 1.5. macOS App Management Writeability Check
	if runtime.GOOS == "darwin" && appPath != "" {
		writeStatus, writeErr := patcher.CheckAppBundleWriteability(appPath)
		if writeErr != nil {
			items = append(items, DiagnosticItem{
				Name:         "App Management",
				Status:       "WARN",
				HumanSummary: "Write Blocked by macOS (Grant App Management)",
				DebugDetails: writeErr.Error(),
			})
		} else {
			items = append(items, DiagnosticItem{
				Name:         "App Management",
				Status:       "OK",
				HumanSummary: "Permissions OK (Bundle Writeable)",
				DebugDetails: writeStatus,
			})
		}
	}

	// 2. Language Server Core
	lsPath := patcher.FindLanguageServerModular(nil)
	if lsPath == "" {
		lsPath = patcher.FindLanguageServer()
	}
	if lsPath != "" {
		status := patcher.CheckLanguageServerStatus(lsPath)
		itemStatus := "WARN"
		human := "Official (Unpatched)"
		if strings.HasPrefix(status, "Patched") {
			itemStatus = "OK"
			human = "Unlocked (Patched)"
		}
		paths := []string{fmt.Sprintf("Core Binary: `%s`", lsPath)}
		bakPath := lsPath + ".agybak"
		if fi, err := os.Stat(bakPath); err == nil && !fi.IsDir() {
			paths = append(paths, fmt.Sprintf("Backup: `Present (%s, %d bytes)`", filepath.Base(bakPath), fi.Size()))
		} else {
			paths = append(paths, "Backup: `None (Will be created on first patch)`")
		}
		items = append(items, DiagnosticItem{
			Name:         "Core Engine",
			Status:       itemStatus,
			HumanSummary: human,
			DebugDetails: fmt.Sprintf("%s (%s)", filepath.Base(lsPath), status),
			TargetPaths:  paths,
		})
	} else {
		mainJsCheck := patcher.FindIdeMainJs()
		asarPathCheck := patcher.FindAntigravityAsar()
		if mainJsCheck != "" || asarPathCheck != "" {
			items = append(items, DiagnosticItem{
				Name:         "Core Engine",
				Status:       "OK",
				HumanSummary: "Integrated / Bundled (Antigravity IDE)",
				DebugDetails: "Embedded AI engine managed via IDE internal bundle",
				TargetPaths:  []string{"Core logic managed via Antigravity IDE bundle"},
			})
		} else {
			items = append(items, DiagnosticItem{
				Name:         "Core Engine",
				Status:       "WARN",
				HumanSummary: "Not Found (language_server binary missing)",
				DebugDetails: "Check resources/bin or resources/app/bin",
			})
		}
	}

	// 3. IDE Frontend
	mainJs := patcher.FindIdeMainJs()
	asarPath := patcher.FindAntigravityAsar()
	targetAsset := asarPath
	if targetAsset == "" {
		targetAsset = mainJs
	}
	if targetAsset != "" {
		itemStatus := "WARN"
		human := "Standard"
		if mainJs != "" {
			if content, err := os.ReadFile(mainJs); err == nil && strings.Contains(string(content), "resetIsTierGCPTos(),true") {
				itemStatus = "OK"
				human = "Unlocked (Internal Mode)"
			}
		}
		if asarPath != "" {
			itemStatus = "OK"
			human = "ASAR Patch & Privacy Active"
		}
		items = append(items, DiagnosticItem{
			Name:         "IDE Interface",
			Status:       itemStatus,
			HumanSummary: human,
			DebugDetails: targetAsset,
			TargetPaths:  []string{fmt.Sprintf("Interface Asset: `%s`", targetAsset)},
		})
	}

	// 3.5. Active Running Processes
	runningPID := patcher.FindRunningLanguageServerPID()
	procSummary := "Idle / Closed"
	procStatus := "OK"
	if runningPID > 0 {
		procSummary = fmt.Sprintf("Running (PID: %d)", runningPID)
	}
	items = append(items, DiagnosticItem{
		Name:         "Active Process",
		Status:       procStatus,
		HumanSummary: procSummary,
		DebugDetails: fmt.Sprintf("language_server PID: %d", runningPID),
	})

	// 3.6. Safe Recent Error Logs Sniffer
	logStatus, logSummary, logDetails := SafeScanRecentErrorLogs()
	var logPaths []string
	if logDetails != "" {
		logPaths = append(logPaths, fmt.Sprintf("Last Trace: `%s`", logDetails))
	}
	items = append(items, DiagnosticItem{
		Name:         "Recent Logs",
		Status:       logStatus,
		HumanSummary: logSummary,
		DebugDetails: logDetails,
		TargetPaths:  logPaths,
	})

	// 4. Local Proxy Detection
	proxies := proxy.DetectActiveProxies()
	if len(proxies) > 0 {
		best, _ := proxy.GetBestProxy()
		summary := fmt.Sprintf("Active on %s (%s)", best.Address, best.Provider)
		items = append(items, DiagnosticItem{
			Name:         "Proxy Connection",
			Status:       "OK",
			HumanSummary: summary,
			DebugDetails: best.URL,
		})
	} else {
		items = append(items, DiagnosticItem{
			Name:         "Proxy Connection",
			Status:       "FAIL",
			HumanSummary: "No Proxy Found (Start v2ray/Xray/Clash)",
			DebugDetails: "Scanned loopback ports 10808, 2080, 7890",
		})
	}

	// 5. DNS Sanity Check (Fast 1s timeout)
	wg.Add(1)
	go func() {
		defer wg.Done()
		r := &net.Resolver{
			PreferGo: true,
		}
		ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
		defer cancel()

		ips, err := r.LookupIP(ctx, "ip", "generativelanguage.googleapis.com")
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			items = append(items, DiagnosticItem{
				Name:         "Google DNS",
				Status:       "FAIL",
				HumanSummary: "Lookup Blocked / Failed",
				DebugDetails: err.Error(),
			})
		} else {
			items = append(items, DiagnosticItem{
				Name:         "Google DNS",
				Status:       "OK",
				HumanSummary: fmt.Sprintf("Healthy (%d IPs resolved)", len(ips)),
				DebugDetails: fmt.Sprintf("%v", ips),
			})
		}
	}()

	// 6. Google Endpoints Reachability (Concurrently executed!)
	bestProxy, _ := proxy.GetBestProxy()
	var transport *http.Transport
	if bestProxy != nil {
		proxyURL, _ := url.Parse(bestProxy.URL)
		transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		}
	} else {
		transport = &http.Transport{}
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   3 * time.Second, // Fast 3s max timeout
	}

	endpoints := []struct {
		Name string
		URL  string
	}{
		{Name: "Gemini AI API", URL: "https://generativelanguage.googleapis.com"},
		{Name: "Cloud Code Service", URL: "https://daily-cloudcode-pa.googleapis.com"},
		{Name: "Google Accounts", URL: "https://accounts.google.com"},
	}

	for _, ep := range endpoints {
		wg.Add(1)
		go func(targetName, targetURL string) {
			defer wg.Done()
			start := time.Now()
			req, err := http.NewRequestWithContext(context.Background(), "GET", targetURL, nil)
			if err != nil {
				return
			}
			// Add browser-like User-Agent to prevent middlebox blocking
			req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) Antigravity-Doctor/5.0")

			resp, err := client.Do(req)
			elapsed := time.Since(start)

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				items = append(items, DiagnosticItem{
					Name:         targetName,
					Status:       "WARN",
					HumanSummary: "Slow / Timed Out",
					DebugDetails: err.Error(),
				})
			} else {
				resp.Body.Close()
				speedLabel := "Fast"
				if elapsed.Milliseconds() > 2000 {
					speedLabel = "Moderate"
				}
				items = append(items, DiagnosticItem{
					Name:         targetName,
					Status:       "OK",
					HumanSummary: fmt.Sprintf("Connected (%dms, %s)", elapsed.Milliseconds(), speedLabel),
					DebugDetails: fmt.Sprintf("HTTP %d, Latency %v", resp.StatusCode, elapsed),
				})
			}
		}(ep.Name, ep.URL)
	}

	wg.Wait()
	return items
}

// SafeScanRecentErrorLogs scans Antigravity log files for recent crashes or errors safely
func SafeScanRecentErrorLogs() (status string, summary string, details string) {
	defer func() {
		if r := recover(); r != nil {
			status = "OK"
			summary = "Clean (No fatal crashes)"
			details = ""
		}
	}()

	home, _ := os.UserHomeDir()
	var logDirs []string
	switch runtime.GOOS {
	case "windows":
		appData := os.Getenv("APPDATA")
		localAppData := os.Getenv("LOCALAPPDATA")
		if appData != "" {
			logDirs = append(logDirs, filepath.Join(appData, "Antigravity", "logs"))
		}
		if localAppData != "" {
			logDirs = append(logDirs, filepath.Join(localAppData, "Antigravity", "logs"))
		}
	case "darwin":
		if home != "" {
			logDirs = append(logDirs, filepath.Join(home, "Library", "Application Support", "Antigravity", "logs"))
		}
	default: // Linux
		if home != "" {
			logDirs = append(logDirs, filepath.Join(home, ".config", "Antigravity", "logs"))
		}
	}

	for _, dir := range logDirs {
		for _, name := range []string{"main.log", "language_server.log"} {
			logFile := filepath.Join(dir, name)
			fi, err := os.Stat(logFile)
			if err != nil || fi.Size() == 0 {
				continue
			}

			f, err := os.Open(logFile)
			if err != nil {
				continue
			}
			offset := int64(0)
			if fi.Size() > 8192 {
				offset = fi.Size() - 8192
			}
			_, _ = f.Seek(offset, 0)
			buf := make([]byte, 8192)
			n, _ := f.Read(buf)
			_ = f.Close()

			content := string(buf[:n])
			lines := strings.Split(content, "\n")
			for i := len(lines) - 1; i >= 0; i-- {
				l := strings.TrimSpace(lines[i])
				if l == "" {
					continue
				}
				lower := strings.ToLower(l)
				if strings.Contains(lower, "status 403") || strings.Contains(lower, "write eio") || strings.Contains(lower, "fatal") || strings.Contains(lower, "panic") || strings.Contains(lower, "connection refused") {
					if len(l) > 80 {
						l = l[:77] + "..."
					}
					return "WARN", fmt.Sprintf("Recent trace in %s", name), l
				}
			}
		}
	}

	return "OK", "Clean (No fatal crashes in logs)", ""
}
