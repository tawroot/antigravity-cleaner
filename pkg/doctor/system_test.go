package doctor

import (
	"testing"
)

func TestDetectOSInfo(t *testing.T) {
	info := DetectOSInfo()
	if info == "" {
		t.Fatalf("expected non-empty OS info")
	}
	t.Logf("Detected OS Info: %s", info)
}

func TestDetectAntigravityVersion(t *testing.T) {
	prod, ver := DetectAntigravityVersion()
	if prod == "" {
		t.Fatalf("expected non-empty product name")
	}
	t.Logf("Detected Antigravity: %s (%s)", prod, ver)
}
