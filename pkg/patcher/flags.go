package patcher

import (
	"log"
	"strings"
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
}

// DefaultEnterpriseFlagsOptions returns the recommended production defaults based on REA findings
func DefaultEnterpriseFlagsOptions() EnterpriseFlagsOptions {
	return EnterpriseFlagsOptions{
		DisableTelemetry:      true,
		IsGoogleInternal:      true,
		JetboxReleaseChannel: "insiders",
		StubbyAuthMode:        "business_licensed_enterprise",
		UseLocalChrome:        true,
		BrowserEvalEnv:        true,
		LocalChromeHeadless:   false,
		AllowAllSidecars:      true,
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

	return flags
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
