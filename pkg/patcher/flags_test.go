package patcher

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildServerFlags(t *testing.T) {
	opts := DefaultEnterpriseFlagsOptions()
	var logs []string
	logger := func(format string, args ...interface{}) {
		logs = append(logs, format)
	}

	flags := BuildServerFlags(opts, logger)
	if len(flags) != 12 {
		t.Fatalf("expected 12 flags, got %d: %v", len(flags), flags)
	}

	var hasTelemetry, hasInternal, hasChannel, hasStubby, hasChrome, hasBrowserEval, hasVisualHeadless, hasSidecars bool
	var hasUnleash, hasMendel, hasBypassSidecars, hasLogDisk, hasCppLogger bool
	for _, f := range flags {
		switch f {
		case "--disable_telemetry=true":
			hasTelemetry = true
		case "--is_google_internal=true":
			hasInternal = true
		case "--jetbox_release_channel=insiders":
			hasChannel = true
		case "--stubby_auth_mode=business_licensed_enterprise":
			hasStubby = true
		case "--use_local_chrome=true":
			hasChrome = true
		case "--browser_eval_env=true":
			hasBrowserEval = true
		case "--local_chrome_headless=false":
			hasVisualHeadless = true
		case "--allow_all_sidecars=true":
			hasSidecars = true
		case "--disable_unleash=true":
			hasUnleash = true
		case "--use_snapshot_mendelfp=true":
			hasMendel = true
		case "--bypass_sidecar_user_config=true":
			hasBypassSidecars = true
		case "--disable_log_to_disk=true":
			hasLogDisk = true
		case "--suppress_cpp_logger_module_check_and_dont_log_cpp_to_stderr=true":
			hasCppLogger = true
		}
	}

	if !hasTelemetry {
		t.Errorf("missing --disable_telemetry=true flag")
	}
	if !hasInternal {
		t.Errorf("missing --is_google_internal=true flag")
	}
	if !hasChannel {
		t.Errorf("missing --jetbox_release_channel=insiders flag")
	}
	if !hasStubby {
		t.Errorf("missing --stubby_auth_mode=business_licensed_enterprise flag")
	}
	if !hasChrome {
		t.Errorf("missing --use_local_chrome=true flag")
	}
	if !hasBrowserEval {
		t.Errorf("missing --browser_eval_env=true flag")
	}
	if !hasVisualHeadless {
		t.Errorf("missing --local_chrome_headless=false flag")
	}
	if !hasSidecars {
		t.Errorf("missing --allow_all_sidecars=true flag")
	}
	if !hasUnleash {
		t.Errorf("missing --disable_unleash=true flag")
	}
	if !hasMendel {
		t.Errorf("missing --use_snapshot_mendelfp=true flag")
	}
	if !hasBypassSidecars {
		t.Errorf("missing --bypass_sidecar_user_config=true flag")
	}
	if !hasLogDisk {
		t.Errorf("missing --disable_log_to_disk=true flag")
	}
	if hasCppLogger {
		t.Errorf("expected --suppress_cpp_logger_module_check_and_dont_log_cpp_to_stderr flag to be disabled by default (Issue #42)")
	}
	if len(logs) != 12 {
		t.Errorf("expected 12 log entries, got %d", len(logs))
	}
}

func TestBuildServerFlagsWithCppLoggerExplicit(t *testing.T) {
	opts := DefaultEnterpriseFlagsOptions()
	opts.SuppressCppLoggerCheck = true

	var logs []string
	logger := func(format string, a ...interface{}) {
		logs = append(logs, fmt.Sprintf(format, a...))
	}

	flags := BuildServerFlags(opts, logger)
	hasCppLogger := false
	for _, f := range flags {
		if f == "--suppress_cpp_logger_module_check_and_dont_log_cpp_to_stderr=true" {
			hasCppLogger = true
		}
	}
	if !hasCppLogger {
		t.Errorf("expected cpp logger flag when SuppressCppLoggerCheck is true")
	}
	if len(logs) != 13 {
		t.Errorf("expected 13 logs, got %d", len(logs))
	}
}


func TestMergeLaunchArguments(t *testing.T) {
	existing := []string{"--standalone", "--override_ide_name=antigravity", "--disable_telemetry=false"}
	opts := DefaultEnterpriseFlagsOptions()

	merged := MergeLaunchArguments(existing, opts, nil)

	// Since --disable_telemetry was already present in existing, it should keep existing to prevent duplicate flags
	var countTelemetry, countInternal bool
	for _, a := range merged {
		if a == "--disable_telemetry=false" {
			countTelemetry = true
		}
		if a == "--is_google_internal=true" {
			countInternal = true
		}
	}

	if !countTelemetry {
		t.Errorf("expected existing telemetry flag to be preserved")
	}
	if !countInternal {
		t.Errorf("expected new internal flag to be merged")
	}
}

func BenchmarkBuildServerFlags(b *testing.B) {
	opts := DefaultEnterpriseFlagsOptions()
	noopLogger := func(format string, args ...interface{}) {}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		flags := BuildServerFlags(opts, noopLogger)
		if len(flags) != 13 {
			b.Fatalf("unexpected flag count: %d", len(flags))
		}
	}
}

func BenchmarkMergeLaunchArguments(b *testing.B) {
	existing := []string{"--standalone", "--override_ide_name=antigravity", "--disable_telemetry=false"}
	opts := DefaultEnterpriseFlagsOptions()
	noopLogger := func(format string, args ...interface{}) {}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		merged := MergeLaunchArguments(existing, opts, noopLogger)
		if len(merged) == 0 {
			b.Fatal("empty merged flags")
		}
	}
}

func TestDetectAndFilterSupportedFlags(t *testing.T) {
	tempDir := t.TempDir()
	mockBin := filepath.Join(tempDir, "mock_language_server")

	// Create a mock executable that returns standard Go flag usage on --help
	mockUsage := `#!/bin/sh
if [ "$1" = "--help" ]; then
  echo "Usage of mock_language_server:"
  echo "  -disable_telemetry=false: If true, disable telemetry logging"
  echo "  -is_google_internal=false: Whether running in internal environment"
  echo "  -jetbox_release_channel=\"stable\": Release channel"
  echo "  -use_local_chrome=false: Whether to use local chrome"
  exit 0
fi
exit 0
`
	if err := os.WriteFile(mockBin, []byte(mockUsage), 0755); err != nil {
		t.Fatalf("failed to write mock binary: %v", err)
	}

	supported, err := DetectSupportedFlags(mockBin)
	if err != nil {
		t.Fatalf("DetectSupportedFlags failed: %v", err)
	}

	if !supported["disable_telemetry"] {
		t.Errorf("expected disable_telemetry to be detected")
	}
	if !supported["is_google_internal"] {
		t.Errorf("expected is_google_internal to be detected")
	}
	if supported["stubby_auth_mode"] {
		t.Errorf("did not expect stubby_auth_mode in mock usage")
	}
	if supported["use_snapshot_mendelfp"] {
		t.Errorf("did not expect use_snapshot_mendelfp in mock usage")
	}

	// Test FilterSupportedFlags
	candidates := []string{
		"--disable_telemetry=true",
		"--is_google_internal=true",
		"--stubby_auth_mode=business_licensed_enterprise",
		"--use_snapshot_mendelfp=true",
	}

	var logs []string
	filtered := FilterSupportedFlags(mockBin, candidates, func(format string, args ...interface{}) {
		logs = append(logs, fmt.Sprintf(format, args...))
	})

	if len(filtered) != 2 {
		t.Fatalf("expected 2 filtered flags, got %d: %v", len(filtered), filtered)
	}
	if filtered[0] != "--disable_telemetry=true" || filtered[1] != "--is_google_internal=true" {
		t.Errorf("unexpected filtered flags: %v", filtered)
	}

	if len(logs) != 2 {
		t.Errorf("expected 2 log messages for filtered flags, got %d: %v", len(logs), logs)
	}
}

func TestBuildServerFlagsForBinaryRealHostTargets(t *testing.T) {
	opts := DefaultEnterpriseFlagsOptions()
	targets := []string{
		"/opt/antigravity/resources/bin/language_server",
		"/home/hokar/Applications/Antigravity/resources/bin/language_server",
	}

	for _, target := range targets {
		if _, err := os.Stat(target); err != nil {
			continue
		}

		var filteredLogs []string
		safeFlags := BuildServerFlagsForBinary(target, opts, func(format string, args ...interface{}) {
			filteredLogs = append(filteredLogs, fmt.Sprintf(format, args...))
		})

		if len(safeFlags) == 0 {
			t.Errorf("expected non-empty safe flags for %s", target)
		}

		// Ensure that unsupported flags are NEVER present for the older binary (/opt/antigravity)
		if strings.Contains(target, "/opt/antigravity") {
			for _, f := range safeFlags {
				if strings.Contains(f, "stubby_auth_mode") {
					t.Errorf("target %s should not have stubby_auth_mode flag: %s", target, f)
				}
				if strings.Contains(f, "use_snapshot_mendelfp") {
					t.Errorf("target %s should not have use_snapshot_mendelfp flag: %s", target, f)
				}
			}
		}
	}
}
