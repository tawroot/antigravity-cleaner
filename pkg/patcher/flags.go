package patcher

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// defaultFlagsLogger provides structured logging with [FLAGS] tag
var defaultFlagsLogger LogFunc = func(format string, args ...interface{}) {
	log.Printf("[FLAGS] "+format, args...)
}

// EnterpriseFlagsOptions specifies options for tuning language_server runtime behavior
type EnterpriseFlagsOptions struct {
	// DisableTelemetry when true prevents tracking, git logs, and IP leakage
	DisableTelemetry bool
	// IsGoogleInternal when true unlocks enterprise tier, Mendel experiments, and URL denylist bypass
	IsGoogleInternal bool
	// JetboxReleaseChannel sets the release channel (e.g. "insiders", "autopush", "stable")
	JetboxReleaseChannel string
	// StubbyAuthMode sets the enterprise auth mode (e.g. "business_licensed_enterprise")
	StubbyAuthMode string
	// UseLocalChrome enables local Google Chrome integration
	UseLocalChrome bool
	// BrowserEvalEnv enables the browser evaluation environment
	BrowserEvalEnv bool
	// LocalChromeHeadless controls whether local Chrome is headless
	LocalChromeHeadless bool
	// AllowAllSidecars permits running any sidecar or MCP process without signature restrictions
	AllowAllSidecars bool
	// DisableUnleash when true halts remote Unleash feature-gating & region-locks
	DisableUnleash bool
	// UseSnapshotMendelfp when true prevents live Mendel experimentation polling
	UseSnapshotMendelfp bool
	// BypassSidecarUserConfig when true keeps all AI sidecars and companions permanently enabled
	BypassSidecarUserConfig bool
	// DisableLogToDisk prevents write EIO and massive local disk logs
	DisableLogToDisk bool
	// SuppressCppLoggerCheck suppresses C++ stderr flood
	SuppressCppLoggerCheck bool
}

// DefaultEnterpriseFlagsOptions returns the recommended production defaults based on REA findings
func DefaultEnterpriseFlagsOptions() EnterpriseFlagsOptions {
	return EnterpriseFlagsOptions{
		DisableTelemetry:        true,
		IsGoogleInternal:        true,
		JetboxReleaseChannel:   "insiders",
		StubbyAuthMode:          "business_licensed_enterprise",
		UseLocalChrome:          true,
		BrowserEvalEnv:          true,
		LocalChromeHeadless:     false,
		AllowAllSidecars:        true,
		DisableUnleash:          true,
		UseSnapshotMendelfp:     true,
		BypassSidecarUserConfig: true,
		DisableLogToDisk:        true,
		SuppressCppLoggerCheck:  false, // Issue #42: Removed in Antigravity 2.28.2+, causes fatal exit status 2 if passed
	}
}


// BuildServerFlags generates the slice of CLI arguments to pass to language_server
func BuildServerFlags(opts EnterpriseFlagsOptions, logger LogFunc) []string {
	if logger == nil {
		logger = defaultFlagsLogger
	}

	var flags []string

	if opts.DisableTelemetry {
		flags = append(flags, "--disable_telemetry=true")
		logger("Applied privacy flag: --disable_telemetry=true (Halts telemetry, Git logging and error traces)")
	}

	if opts.IsGoogleInternal {
		flags = append(flags, "--is_google_internal=true")
		logger("Applied enterprise flag: --is_google_internal=true (Unlocks internal payload metadata and experiments)")
	}

	if opts.JetboxReleaseChannel != "" {
		flags = append(flags, "--jetbox_release_channel="+opts.JetboxReleaseChannel)
		logger("Applied release channel: --jetbox_release_channel=" + opts.JetboxReleaseChannel)
	}

	if opts.StubbyAuthMode != "" {
		flags = append(flags, "--stubby_auth_mode="+opts.StubbyAuthMode)
		logger("Applied stubby auth mode: --stubby_auth_mode=" + opts.StubbyAuthMode)
	}

	if opts.UseLocalChrome {
		flags = append(flags, "--use_local_chrome=true")
		logger("Applied local Chrome flag: --use_local_chrome=true")
	}

	if opts.BrowserEvalEnv {
		flags = append(flags, "--browser_eval_env=true")
		logger("Applied browser eval flag: --browser_eval_env=true")
	}

	if !opts.LocalChromeHeadless {
		flags = append(flags, "--local_chrome_headless=false")
		logger("Applied visual browser flag: --local_chrome_headless=false")
	}

	if opts.AllowAllSidecars {
		flags = append(flags, "--allow_all_sidecars=true")
		logger("Applied sidecar policy flag: --allow_all_sidecars=true (Unlocks all local companion processes and MCPs)")
	}

	if opts.DisableUnleash {
		flags = append(flags, "--disable_unleash=true")
		logger("Applied region gate bypass flag: --disable_unleash=true (Halts remote Unleash geographic gating)")
	}

	if opts.UseSnapshotMendelfp {
		flags = append(flags, "--use_snapshot_mendelfp=true")
		logger("Applied experiment shield flag: --use_snapshot_mendelfp=true (Forces embedded snapshot, halts server polling)")
	}

	if opts.BypassSidecarUserConfig {
		flags = append(flags, "--bypass_sidecar_user_config=true")
		logger("Applied sidecar config bypass: --bypass_sidecar_user_config=true (Forces all sidecars enabled)")
	}

	if opts.DisableLogToDisk {
		flags = append(flags, "--disable_log_to_disk=true")
		logger("Applied disk shield flag: --disable_log_to_disk=true (Prevents massive local log dumps and write EIO)")
	}

	if opts.SuppressCppLoggerCheck {
		flags = append(flags, "--suppress_cpp_logger_module_check_and_dont_log_cpp_to_stderr=true")
		logger("Applied stderr shield: --suppress_cpp_logger_module_check_and_dont_log_cpp_to_stderr=true")
	}

	return flags
}

// DetectSupportedFlags inspects a language_server binary to determine which CLI flags it natively supports.
// It executes '<binPath> --help' with a short timeout and parses the usage declaration.
// This prevents fatal Go runtime exits (status 2: 'flags provided but not defined') across different core builds (Issue #48).
func DetectSupportedFlags(binPath string) (map[string]bool, error) {
	if binPath == "" {
		return nil, fmt.Errorf("empty binary path")
	}
	if fi, err := os.Stat(binPath); err != nil || fi.IsDir() {
		return nil, fmt.Errorf("invalid binary path %q: %w", binPath, err)
	}

	supported := make(map[string]bool)

	// 1. Proactive execution probe with --help
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, binPath, "--help")
	output, _ := cmd.CombinedOutput()
	if len(output) > 0 {
		scanner := bufio.NewScanner(bytes.NewReader(output))
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if strings.HasPrefix(line, "-") {
				parts := strings.SplitN(line, ":", 2)
				flagPart := strings.TrimLeft(parts[0], "-")
				flagName := strings.TrimSpace(strings.SplitN(flagPart, "=", 2)[0])
				if flagName != "" {
					supported[flagName] = true
				}
			}
		}
	}

	if len(supported) > 0 {
		return supported, nil
	}

	// 2. Fallback: Scan binary content for known flag string markers if --help yielded nothing
	data, err := os.ReadFile(binPath)
	if err != nil {
		return nil, err
	}

	candidateMarkers := []string{
		"disable_telemetry",
		"is_google_internal",
		"jetbox_release_channel",
		"stubby_auth_mode",
		"use_stubby_auth",
		"use_local_chrome",
		"browser_eval_env",
		"local_chrome_headless",
	}

	for _, m := range candidateMarkers {
		if bytes.Contains(data, []byte(m)) {
			supported[m] = true
		}
	}

	return supported, nil
}

// FilterSupportedFlags validates candidate flags against the binary's actual registered flags.
// If supported flags cannot be probed (e.g. non-binary or mock test file), candidates are retained.
func FilterSupportedFlags(binPath string, candidateFlags []string, logger LogFunc) []string {
	if binPath == "" || len(candidateFlags) == 0 {
		return candidateFlags
	}

	supported, err := DetectSupportedFlags(binPath)
	if err != nil || len(supported) == 0 {
		return candidateFlags
	}

	binBase := filepath.Base(binPath)
	var filtered []string
	for _, flag := range candidateFlags {
		flagName := strings.TrimLeft(flag, "-")
		flagName = strings.TrimSpace(strings.SplitN(flagName, "=", 2)[0])
		if supported[flagName] {
			filtered = append(filtered, flag)
		} else {
			if logger != nil {
				logger("Probed %s: filtered unsupported runtime flag %s (preventing status 2 crash, closes #48)", binBase, flag)
			}
		}
	}

	return filtered
}

// BuildServerFlagsForBinary generates safe, verified CLI arguments specifically tailored for the target binary.
func BuildServerFlagsForBinary(binPath string, opts EnterpriseFlagsOptions, logger LogFunc) []string {
	candidateFlags := BuildServerFlags(opts, logger)
	return FilterSupportedFlags(binPath, candidateFlags, logger)
}


// MergeLaunchArguments merges recommended flags into an existing argument list without duplicates
func MergeLaunchArguments(existingArgs []string, opts EnterpriseFlagsOptions, logger LogFunc) []string {
	recommended := BuildServerFlags(opts, logger)
	result := make([]string, len(existingArgs))
	copy(result, existingArgs)

	existingSet := make(map[string]bool)
	for _, arg := range existingArgs {
		key := strings.SplitN(arg, "=", 2)[0]
		existingSet[key] = true
	}

	for _, rec := range recommended {
		key := strings.SplitN(rec, "=", 2)[0]
		if !existingSet[key] {
			result = append(result, rec)
		}
	}

	return result
}
