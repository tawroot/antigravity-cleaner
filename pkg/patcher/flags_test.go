package patcher

import (
	"testing"
)

func TestBuildServerFlags(t *testing.T) {
	opts := DefaultEnterpriseFlagsOptions()
	var logs []string
	logger := func(format string, args ...interface{}) {
		logs = append(logs, format)
	}

	flags := BuildServerFlags(opts, logger)
	if len(flags) != 13 {
		t.Fatalf("expected 13 flags, got %d: %v", len(flags), flags)
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
	if !hasCppLogger {
		t.Errorf("missing --suppress_cpp_logger_module_check_and_dont_log_cpp_to_stderr=true flag")
	}
	if len(logs) != 13 {
		t.Errorf("expected 13 log entries, got %d", len(logs))
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
