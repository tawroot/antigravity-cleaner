package proxy

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// ProxyEndpoint holds discovered proxy information
type ProxyEndpoint struct {
	Type     string // "socks5", "http", "mixed"
	Address  string // "127.0.0.1:10808"
	URL      string // "socks5h://127.0.0.1:10808"
	Provider string // "v2rayN / Xray", "NekoBox / Hiddify", "Clash / Mihomo"
}

// Known default proxy ports and software
var defaultCandidates = []struct {
	Port     int
	Type     string
	Provider string
}{
	{Port: 10808, Type: "socks5", Provider: "v2ray / Xray"},
	{Port: 10809, Type: "http", Provider: "v2ray / Xray HTTP"},
	{Port: 2080, Type: "socks5", Provider: "NekoBox / Hiddify"},
	{Port: 2081, Type: "http", Provider: "Hiddify HTTP"},
	{Port: 7890, Type: "mixed", Provider: "Clash / Mihomo"},
	{Port: 1080, Type: "socks5", Provider: "Standard SOCKS5"},
	{Port: 8888, Type: "http", Provider: "Custom HTTP"},
	{Port: 8080, Type: "http", Provider: "HTTP Proxy"},
}

// CleanHostPort strips protocol prefixes and trailing slashes for concise display
func CleanHostPort(rawURL string) string {
	s := strings.TrimSpace(rawURL)
	s = strings.TrimPrefix(s, "socks5h://")
	s = strings.TrimPrefix(s, "socks5://")
	s = strings.TrimPrefix(s, "socks://")
	s = strings.TrimPrefix(s, "http://")
	s = strings.TrimPrefix(s, "https://")
	s = strings.TrimSuffix(s, "/")
	return s
}

// IsPortOpen tests if a TCP port is responding on 127.0.0.1
func IsPortOpen(port int, timeout time.Duration) bool {
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// DetectActiveProxies scans loopback ports and environment variables
func DetectActiveProxies() []ProxyEndpoint {
	var found []ProxyEndpoint

	// 1. Scan loopback ports in parallel with short timeout
	type probeResult struct {
		endpoint ProxyEndpoint
		open     bool
	}

	ch := make(chan probeResult, len(defaultCandidates))
	for _, c := range defaultCandidates {
		go func(port int, pType, provider string) {
			open := IsPortOpen(port, 250*time.Millisecond)
			urlPrefix := "http://"
			if pType == "socks5" || pType == "mixed" {
				urlPrefix = "socks5h://" // 'socks5h' forces remote DNS resolution!
			}
			ch <- probeResult{
				open: open,
				endpoint: ProxyEndpoint{
					Type:     pType,
					Address:  fmt.Sprintf("127.0.0.1:%d", port),
					URL:      fmt.Sprintf("%s127.0.0.1:%d", urlPrefix, port),
					Provider: provider,
				},
			}
		}(c.Port, c.Type, c.Provider)
	}

	for i := 0; i < len(defaultCandidates); i++ {
		res := <-ch
		if res.open {
			found = append(found, res.endpoint)
		}
	}

	// 2. Fallback to check environment variables if no local port detected
	if len(found) == 0 {
		for _, envVar := range []string{"ALL_PROXY", "all_proxy", "HTTPS_PROXY", "https_proxy"} {
			val := os.Getenv(envVar)
			if val != "" {
				val = strings.TrimSpace(val)
				typ := "http"
				if strings.HasPrefix(val, "socks") {
					typ = "socks5"
				}
				cleanAddr := CleanHostPort(val)
				found = append(found, ProxyEndpoint{
					Type:     typ,
					Address:  cleanAddr,
					URL:      val,
					Provider: "System Proxy",
				})
				break
			}
		}
	}

	return found
}

// GetBestProxy returns the optimal proxy (preferring SOCKS5 with remote DNS)
func GetBestProxy() (*ProxyEndpoint, error) {
	active := DetectActiveProxies()
	if len(active) == 0 {
		return nil, fmt.Errorf("no active proxy detected on 127.0.0.1 (please start your proxy client)")
	}

	// Prefer socks5/socks5h first for remote DNS, then mixed, then http
	for _, p := range active {
		if p.Type == "socks5" || p.Type == "mixed" {
			return &p, nil
		}
	}

	return &active[0], nil
}
