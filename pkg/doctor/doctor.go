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
	"strings"
	"sync"
	"time"

	"github.com/tawroot/antigravity-cleaner/pkg/patcher"
	"github.com/tawroot/antigravity-cleaner/pkg/proxy"
)

// DiagnosticItem represents a friendly diagnostic check
type DiagnosticItem struct {
	Name         string
	Status       string // "OK", "WARN", "FAIL"
	HumanSummary string // Short friendly text for regular users
	DebugDetails string // Raw technical string for --verbose
}

// RunDiagnostics executes concurrent, fast, non-intimidating diagnostics
func RunDiagnostics(verbose bool) []DiagnosticItem {
	var items []DiagnosticItem
	var mu sync.Mutex
	var wg sync.WaitGroup

	// 1. Antigravity Installation
	appPath := proxy.FindAntigravityExecutable()
	if appPath != "" {
		items = append(items, DiagnosticItem{
			Name:         "Antigravity IDE",
			Status:       "OK",
			HumanSummary: "Installed & Detected",
			DebugDetails: appPath,
		})
	} else {
		items = append(items, DiagnosticItem{
			Name:         "Antigravity IDE",
			Status:       "WARN",
			HumanSummary: "Not Found in Standard Path",
			DebugDetails: "Check /opt, /usr/share or ~/.local",
		})
	}

	// 2. Language Server Core
	lsPath := patcher.FindLanguageServer()
	if lsPath != "" {
		status := patcher.CheckLanguageServerStatus(lsPath)
		itemStatus := "WARN"
		human := "Official (Unpatched)"
		if strings.HasPrefix(status, "Patched") {
			itemStatus = "OK"
			human = "Unlocked (Patched)"
		}
		items = append(items, DiagnosticItem{
			Name:         "Core Engine",
			Status:       itemStatus,
			HumanSummary: human,
			DebugDetails: fmt.Sprintf("%s (%s)", filepath.Base(lsPath), status),
		})
	} else {
		items = append(items, DiagnosticItem{
			Name:         "Core Engine",
			Status:       "WARN",
			HumanSummary: "Not Found",
			DebugDetails: "language_server binary missing",
		})
	}

	// 3. IDE Frontend
	mainJs := patcher.FindIdeMainJs()
	if mainJs != "" {
		content, err := os.ReadFile(mainJs)
		itemStatus := "WARN"
		human := "Standard"
		if err == nil && strings.Contains(string(content), "resetIsTierGCPTos(),true") {
			itemStatus = "OK"
			human = "Unlocked (Internal Mode)"
		}
		items = append(items, DiagnosticItem{
			Name:         "IDE Interface",
			Status:       itemStatus,
			HumanSummary: human,
			DebugDetails: mainJs,
		})
	}

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
