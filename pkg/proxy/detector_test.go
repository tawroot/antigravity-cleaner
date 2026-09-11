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
	env := BuildProxyEnv("socks5h://127.0.0.1:10808")
	var foundAllProxy bool
	for _, e := range env {
		if strings.HasPrefix(e, "ALL_PROXY=") {
			foundAllProxy = true
			if !strings.Contains(e, "127.0.0.1:10808") {
				t.Errorf("ALL_PROXY does not contain target proxy: %s", e)
			}
		}
	}
	if !foundAllProxy {
		t.Errorf("BuildProxyEnv did not produce ALL_PROXY")
	}
}
