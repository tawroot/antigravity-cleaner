package proxy

import (
	"strings"
	"testing"
)

func TestCleanHostPort(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"socks5://127.0.0.1:10808", "127.0.0.1:10808"},
		{"http://127.0.0.1:7890/", "127.0.0.1:7890"},
		{"socks5h://127.0.0.1:2080", "127.0.0.1:2080"},
		{"127.0.0.1:10808", "127.0.0.1:10808"},
	}

	for _, c := range cases {
		res := CleanHostPort(c.input)
		if res != c.expected {
			t.Errorf("CleanHostPort(%s) = %s; want %s", c.input, res, c.expected)
		}
	}
}

func TestBuildProxyEnv(t *testing.T) {
	// 1. SOCKS proxy: ALL_PROXY should be set, but HTTPS_PROXY must not be set to socks (prevents Undici crash)
	socksEnv := BuildProxyEnv("socks5://127.0.0.1:10808")
	var foundAllProxy bool
	for _, e := range socksEnv {
		if strings.HasPrefix(e, "ALL_PROXY=") {
			foundAllProxy = true
			if !strings.Contains(e, "127.0.0.1:10808") {
				t.Errorf("ALL_PROXY does not contain target proxy: %s", e)
			}
		}
		if strings.HasPrefix(strings.ToUpper(e), "HTTPS_PROXY=") || strings.HasPrefix(strings.ToUpper(e), "HTTP_PROXY=") {
			t.Errorf("HTTPS_PROXY / HTTP_PROXY should not be set for socks proxies, got: %s", e)
		}
	}
	if !foundAllProxy {
		t.Errorf("BuildProxyEnv did not produce ALL_PROXY for socks")
	}

	// 2. HTTP proxy: ALL_PROXY, HTTPS_PROXY, and HTTP_PROXY should be set
	httpEnv := BuildProxyEnv("http://127.0.0.1:7890")
	var foundHttpsProxy bool
	for _, e := range httpEnv {
		if strings.HasPrefix(e, "HTTPS_PROXY=") {
			foundHttpsProxy = true
			if !strings.Contains(e, "127.0.0.1:7890") {
				t.Errorf("HTTPS_PROXY does not contain target proxy: %s", e)
			}
		}
	}
	if !foundHttpsProxy {
		t.Errorf("BuildProxyEnv did not produce HTTPS_PROXY for http")
	}
}

func BenchmarkCleanHostPort(b *testing.B) {
	input := "socks5h://127.0.0.1:10808"
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		res := CleanHostPort(input)
		if res == "" {
			b.Fatal("empty host port")
		}
	}
}

func BenchmarkBuildProxyEnv(b *testing.B) {
	input := "socks5h://127.0.0.1:10808"
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		env := BuildProxyEnv(input)
		if len(env) == 0 {
			b.Fatal("empty proxy env")
		}
	}
}
