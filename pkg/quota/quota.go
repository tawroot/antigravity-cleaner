package quota

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/tawroot/antigravity-cleaner/pkg/patcher"
	"github.com/zalando/go-keyring"
	_ "modernc.org/sqlite"
)

const (
	UsageURL         = "https://cloudcode-pa.googleapis.com/v1internal:retrieveUserQuotaSummary"
	DetailedQuotaURL = "https://cloudcode-pa.googleapis.com/v1internal:retrieveUserQuota"
)

var (
	UsageOAuthID     = xorDecode([]byte{107, 106, 109, 107, 106, 106, 108, 106, 108, 106, 111, 99, 107, 119, 46, 55, 50, 41, 41, 51, 52, 104, 50, 104, 107, 54, 57, 40, 63, 104, 105, 111, 44, 46, 53, 54, 53, 48, 50, 110, 61, 110, 106, 105, 63, 42, 116, 59, 42, 42, 41, 116, 61, 53, 53, 61, 54, 63, 47, 41, 63, 40, 57, 53, 52, 46, 63, 52, 46, 116, 57, 53, 55})
	UsageOAuthSecret = xorDecode([]byte{29, 21, 25, 9, 10, 2, 119, 17, 111, 98, 28, 13, 8, 110, 98, 108, 22, 62, 22, 16, 107, 55, 22, 24, 98, 41, 2, 25, 110, 32, 108, 43, 30, 27, 60})
)

func xorDecode(b []byte) string {
	res := make([]byte, len(b))
	for i, v := range b {
		res[i] = v ^ 0x5a
	}
	return string(res)
}

// QuotaBucket represents a time-window usage quota
type QuotaBucket struct {
	BucketID          string    `json:"bucketId,omitempty"`
	DisplayName       string    `json:"displayName,omitempty"`
	Description       string    `json:"description,omitempty"`
	Window            string    `json:"window"`
	RemainingFraction *float64  `json:"remainingFraction"`
	ResetTimeStr      string    `json:"resetTime"`
	ResetTime         time.Time `json:"-"`
	Disabled          bool      `json:"disabled"`
}

// QuotaGroup represents a group of quotas (e.g. Gemini, Claude/GPT)
type QuotaGroup struct {
	DisplayName string        `json:"displayName"`
	Description string        `json:"description,omitempty"`
	Buckets     []QuotaBucket `json:"buckets"`
}

// QuotaSummaryResponse maps the Google CloudCode quota summary response
type QuotaSummaryResponse struct {
	Response struct {
		Groups []QuotaGroup `json:"groups"`
	} `json:"response"`
	Groups []QuotaGroup `json:"groups"`
}

// DetailedQuotaBucket represents individual model quotas from retrieveUserQuota API
type DetailedQuotaBucket struct {
	ModelID           string    `json:"modelId"`
	TokenType         string    `json:"tokenType"`
	RemainingFraction *float64  `json:"remainingFraction"`
	ResetTimeStr      string    `json:"resetTime"`
	ResetTime         time.Time `json:"-"`
}

type DetailedQuotaResponse struct {
	Buckets []DetailedQuotaBucket `json:"buckets"`
}

// FindStateVscdb finds the newest state.vscdb across standard install locations
func FindStateVscdb() string {
	home, _ := os.UserHomeDir()
	var candidates []string

	if runtime.GOOS == "windows" {
		appData := os.Getenv("APPDATA")
		userProfile := os.Getenv("USERPROFILE")
		candidates = append(candidates,
			filepath.Join(appData, "Antigravity IDE", "User", "globalStorage", "state.vscdb"),
			filepath.Join(appData, "Antigravity", "User", "globalStorage", "state.vscdb"),
			filepath.Join(appData, "Code", "User", "globalStorage", "state.vscdb"),
			filepath.Join(appData, "Code - Insiders", "User", "globalStorage", "state.vscdb"),
			filepath.Join(userProfile, "scoop", "persist", "antigravity-ide", "data", "user-data", "User", "globalStorage", "state.vscdb"),
		)
	} else if runtime.GOOS == "darwin" {
		candidates = append(candidates,
			filepath.Join(home, "Library", "Application Support", "Antigravity IDE", "User", "globalStorage", "state.vscdb"),
			filepath.Join(home, "Library", "Application Support", "Antigravity", "User", "globalStorage", "state.vscdb"),
			filepath.Join(home, "Library", "Application Support", "Code", "User", "globalStorage", "state.vscdb"),
			filepath.Join(home, "Library", "Application Support", "Code - Insiders", "User", "globalStorage", "state.vscdb"),
		)
	} else {
		// Linux
		candidates = append(candidates,
			filepath.Join(home, ".config", "Antigravity IDE", "User", "globalStorage", "state.vscdb"),
			filepath.Join(home, ".config", "Antigravity", "User", "globalStorage", "state.vscdb"),
			filepath.Join(home, ".config", "Code", "User", "globalStorage", "state.vscdb"),
			filepath.Join(home, ".config", "Code - Insiders", "User", "globalStorage", "state.vscdb"),
			filepath.Join(home, ".config", "VSCodium", "User", "globalStorage", "state.vscdb"),
		)
	}

	var newestPath string
	var newestTime int64

	for _, c := range candidates {
		if fi, err := os.Stat(c); err == nil && !fi.IsDir() {
			if hasAntigravityToken(c) {
				if fi.ModTime().Unix() > newestTime {
					newestTime = fi.ModTime().Unix()
					newestPath = c
				}
			}
		}
	}
	return newestPath
}

func hasAntigravityToken(dbPath string) bool {
	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s?mode=ro", dbPath))
	if err != nil {
		return false
	}
	defer db.Close()
	var val string
	err = db.QueryRow("SELECT value FROM ItemTable WHERE key = 'antigravityUnifiedStateSync.oauthToken'").Scan(&val)
	return err == nil && val != ""
}

// FindHubLoginUsername checks if an Antigravity Standalone Hub user is logged in
func FindHubLoginUsername() string {
	home, _ := os.UserHomeDir()
	var candidates []string

	if runtime.GOOS == "windows" {
		candidates = append(candidates, filepath.Join(os.Getenv("APPDATA"), "Antigravity", "app_storage.json"))
	} else if runtime.GOOS == "darwin" {
		candidates = append(candidates, filepath.Join(home, "Library", "Application Support", "Antigravity", "app_storage.json"))
	} else {
		candidates = append(candidates, filepath.Join(home, ".config", "Antigravity", "app_storage.json"))
	}

	for _, c := range candidates {
		if data, err := os.ReadFile(c); err == nil {
			var d map[string]interface{}
			if err := json.Unmarshal(data, &d); err == nil {
				if userVal, ok := d["jetski.onboarding.lastLoginUsername"].(string); ok && userVal != "" {
					return userVal
				}
			}
		}
	}
	return ""
}

var tokenRe = regexp.MustCompile(`1//[a-zA-Z0-9_-]+`)

// ExtractRefreshToken reads the Google OAuth refresh token from OS Keyring or state.vscdb
func ExtractRefreshToken(dbPath string) (string, error) {
	// Mode 1: Check Antigravity Standalone Hub credentials from OS Keyring (Linux/macOS/Windows) when dbPath is not explicitly specified
	if dbPath == "" {
		if krVal, err := keyring.Get("gemini", "antigravity"); err == nil && krVal != "" {
			var payload struct {
				Token struct {
					RefreshToken string `json:"refresh_token"`
				} `json:"token"`
			}
			if err := json.Unmarshal([]byte(krVal), &payload); err == nil && payload.Token.RefreshToken != "" {
				return payload.Token.RefreshToken, nil
			}
			if match := tokenRe.FindString(krVal); match != "" {
				return match, nil
			}
		}
	}

	// Mode 2: Check IDE (VS Code state.vscdb)
	if dbPath == "" {
		dbPath = FindStateVscdb()
		if dbPath == "" {
			if hubUser := FindHubLoginUsername(); hubUser != "" {
				return "", fmt.Errorf("active session is in Antigravity Standalone Hub (logged in as %s).\n    The Standalone Hub encrypts tokens in browser cookies.\n    To query CloudCode Quota API via CLI, sign in to Antigravity IDE (VS Code extension) or agy CLI", hubUser)
			}
			return "", fmt.Errorf("state.vscdb not found — please launch Antigravity IDE and sign in first")
		}
	}

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s?mode=ro", dbPath))
	if err != nil {
		return "", fmt.Errorf("failed to open %s: %w", dbPath, err)
	}
	defer db.Close()

	var rawValue string
	err = db.QueryRow("SELECT value FROM ItemTable WHERE key = 'antigravityUnifiedStateSync.oauthToken'").Scan(&rawValue)
	if err != nil {
		return "", fmt.Errorf("antigravityUnifiedStateSync.oauthToken not found in database: %w", err)
	}

	// Token might be base64-encoded or raw JSON
	var searchPool []byte
	if decoded, err := base64.StdEncoding.DecodeString(rawValue); err == nil {
		searchPool = decoded
	} else {
		searchPool = []byte(rawValue)
	}

	// Try finding 1//... directly
	match := tokenRe.Find(searchPool)
	if len(match) > 0 {
		return string(match), nil
	}

	// Check for URL-safe base64 sub-blocks
	wordRe := regexp.MustCompile(`[a-zA-Z0-9_-]{32,}`)
	words := wordRe.FindAll(searchPool, -1)
	for _, w := range words {
		padded := string(w)
		for len(padded)%4 != 0 {
			padded += "="
		}
		if dec, err := base64.URLEncoding.DecodeString(padded); err == nil {
			subMatch := tokenRe.Find(dec)
			if len(subMatch) > 0 {
				return string(subMatch), nil
			}
		}
	}

	return "", fmt.Errorf("no usable Google OAuth refresh token found in database")
}

// RefreshAccessToken exchanges the refresh token for a live Google access token
func RefreshAccessToken(refreshToken string) (string, error) {
	if refreshToken == "" {
		return "", fmt.Errorf("refresh token is empty")
	}

	form := url.Values{}
	form.Set("client_id", UsageOAuthID)
	form.Set("client_secret", UsageOAuthSecret)
	form.Set("grant_type", "refresh_token")
	form.Set("refresh_token", refreshToken)

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.PostForm("https://oauth2.googleapis.com/token", form)
	if err != nil {
		return "", fmt.Errorf("failed to contact oauth2.googleapis.com: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("OAuth refresh failed (HTTP %d): %s", resp.StatusCode, string(body))
	}

	var res struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &res); err != nil {
		return "", fmt.Errorf("failed to parse token response: %w", err)
	}
	if res.AccessToken == "" {
		return "", fmt.Errorf("empty access token in OAuth response")
	}

	return res.AccessToken, nil
}

// FetchQuotaSummary retrieves the quota summary.
// It first attempts to query the active local language_server Connect RPC (which contains both Gemini and 3P weekly+5h quotas).
// If language_server is not running or fails, it falls back to Google CloudCode UsageURL.
func FetchQuotaSummary(accessToken string) ([]QuotaGroup, error) {
	if groups, err := DiscoverLanguageServerRPC(); err == nil && len(groups) > 0 {
		return groups, nil
	}
	return FetchQuotaSummaryFromCloud(accessToken)
}

// DiscoverLanguageServerRPC probes the local running language_server Connect RPC
func DiscoverLanguageServerRPC() ([]QuotaGroup, error) {
	pid := patcher.FindRunningLanguageServerPID()
	if pid == 0 {
		return nil, fmt.Errorf("language_server process not running")
	}

	cmdlineBytes, err := os.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
	if err != nil {
		return nil, fmt.Errorf("failed to read /proc/%d/cmdline: %w", pid, err)
	}

	args := strings.Split(string(cmdlineBytes), "\x00")
	var csrfToken string
	for i, arg := range args {
		if (arg == "--csrf_token" || arg == "-csrf_token") && i+1 < len(args) {
			csrfToken = args[i+1]
			break
		}
		if strings.HasPrefix(arg, "--csrf_token=") {
			csrfToken = strings.TrimPrefix(arg, "--csrf_token=")
			break
		}
	}

	if csrfToken == "" {
		return nil, fmt.Errorf("csrf_token not found in language_server arguments")
	}

	ports := findPIDListeningPorts(pid)
	if len(ports) == 0 {
		return nil, fmt.Errorf("no listening ports found for language_server PID %d", pid)
	}

	client := &http.Client{Timeout: 3 * time.Second}
	for _, port := range ports {
		url := fmt.Sprintf("http://127.0.0.1:%d/exa.language_server_pb.LanguageServerService/RetrieveUserQuotaSummary", port)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte("{}")))
		if err != nil {
			continue
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-codeium-csrf-token", csrfToken)

		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			var summary QuotaSummaryResponse
			if err := json.Unmarshal(body, &summary); err == nil {
				groups := summary.Response.Groups
				if len(groups) == 0 {
					groups = summary.Groups
				}
				if len(groups) > 0 {
					for i := range groups {
						for j := range groups[i].Buckets {
							b := &groups[i].Buckets[j]
							if b.ResetTimeStr != "" {
								if t, err := time.Parse(time.RFC3339, b.ResetTimeStr); err == nil {
									b.ResetTime = t
								}
							}
						}
					}
					return groups, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("none of the language_server ports responded to RetrieveUserQuotaSummary")
}

func findPIDListeningPorts(pid int) []int {
	data, err := os.ReadFile(fmt.Sprintf("/proc/%d/net/tcp", pid))
	if err != nil {
		return nil
	}

	var ports []int
	lines := strings.Split(string(data), "\n")
	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) >= 4 {
			if fields[3] == "0A" { // 0A = TCP_LISTEN
				localAddr := fields[1]
				parts := strings.Split(localAddr, ":")
				if len(parts) == 2 {
					if portNum, err := strconv.ParseInt(parts[1], 16, 64); err == nil {
						ports = append(ports, int(portNum))
					}
				}
			}
		}
	}
	return ports
}

// FetchQuotaSummaryFromCloud retrieves the quota summary from CloudCode API
func FetchQuotaSummaryFromCloud(accessToken string) ([]QuotaGroup, error) {
	client := &http.Client{Timeout: 15 * time.Second}

	req, err := http.NewRequest("POST", UsageURL, bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "antigravity")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("quota service request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("quota service returned HTTP %d: %s", resp.StatusCode, string(body))
	}

	var summary QuotaSummaryResponse
	if err := json.Unmarshal(body, &summary); err != nil {
		return nil, fmt.Errorf("failed to parse quota summary JSON: %w", err)
	}

	groups := summary.Response.Groups
	if len(groups) == 0 {
		groups = summary.Groups
	}

	// Parse timestamps
	for i := range groups {
		for j := range groups[i].Buckets {
			b := &groups[i].Buckets[j]
			if b.ResetTimeStr != "" {
				t, err := time.Parse(time.RFC3339, b.ResetTimeStr)
				if err == nil {
					b.ResetTime = t
				}
			}
		}
	}

	return groups, nil
}

// FetchDetailedUserQuota retrieves individual model-by-model quotas from retrieveUserQuota API
func FetchDetailedUserQuota(accessToken string) ([]DetailedQuotaBucket, error) {
	client := &http.Client{Timeout: 15 * time.Second}

	req, err := http.NewRequest("POST", DetailedQuotaURL, bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "antigravity")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("detailed quota request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("detailed quota returned HTTP %d: %s", resp.StatusCode, string(body))
	}

	var res DetailedQuotaResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("failed to parse detailed quota JSON: %w", err)
	}

	for i := range res.Buckets {
		b := &res.Buckets[i]
		if b.ResetTimeStr != "" {
			if t, err := time.Parse(time.RFC3339, b.ResetTimeStr); err == nil {
				b.ResetTime = t
			}
		}
	}

	return res.Buckets, nil
}

// RenderQuota renders a colorful terminal representation matching the official Antigravity IDE UI
func RenderQuota(groups []QuotaGroup, detailed []DetailedQuotaBucket) string {
	if len(groups) == 0 && len(detailed) == 0 {
		return "  [!] No quota data returned by Google CloudCode API.\n"
	}

	var sb strings.Builder
	sb.WriteString("\n  ═══════════════════════════════════════════════════════════════\n")
	sb.WriteString("    ANTIGRAVITY GOOGLE CLOUD LIVE QUOTA MONITOR\n")
	sb.WriteString("  ═══════════════════════════════════════════════════════════════\n\n")

	now := time.Now().UTC()

	// 1. Check if Gemini group is already in groups (e.g. from paid/custom tiers)
	var geminiGroup *QuotaGroup
	for i := range groups {
		if strings.Contains(strings.ToLower(groups[i].DisplayName), "gemini") {
			geminiGroup = &groups[i]
			break
		}
	}

	// 2. Extract representative Gemini bucket and image generation bucket from detailed
	var primaryGemini *DetailedQuotaBucket
	var imageGemini *DetailedQuotaBucket

	for i := range detailed {
		b := &detailed[i]
		if strings.Contains(strings.ToLower(b.ModelID), "image") {
			if imageGemini == nil {
				imageGemini = b
			}
		} else if strings.HasPrefix(b.ModelID, "gemini-") {
			if primaryGemini == nil {
				primaryGemini = b
			}
		}
	}

	// Card 1: Gemini Models (matching Antigravity IDE UI)
	sb.WriteString("  • \x1b[1;32mGemini Models\x1b[0m\n")
	sb.WriteString("    \x1b[90mModels within this group: Gemini 3.8 Flash, Pro, Thinking\x1b[0m\n")

	if geminiGroup != nil && len(geminiGroup.Buckets) > 0 {
		for _, b := range geminiGroup.Buckets {
			bar, pctStr := renderProgressBar(b.RemainingFraction)
			resetInfo := formatResetTime(b.ResetTime, now)
			windowLabel := fmt.Sprintf("%-6s", b.Window)
			sb.WriteString(fmt.Sprintf("    %-8s %s %s %s\n", windowLabel, bar, pctStr, resetInfo))
		}
	} else if primaryGemini != nil {
		bar, pctStr := renderProgressBar(primaryGemini.RemainingFraction)
		resetInfo := formatResetTime(primaryGemini.ResetTime, now)
		sb.WriteString(fmt.Sprintf("    %-8s %s %s %s\n", "5h", bar, pctStr, resetInfo))
	} else {
		sb.WriteString("    No active quota bucket reported.\n")
	}
	sb.WriteString("\n")

	// Card 2: Claude and GPT models (matching Antigravity IDE UI)
	for _, g := range groups {
		if strings.Contains(strings.ToLower(g.DisplayName), "gemini") {
			continue // Already rendered in Card 1
		}
		sb.WriteString(fmt.Sprintf("  • \x1b[1;36m%s\x1b[0m\n", g.DisplayName))
		if g.Description != "" {
			sb.WriteString(fmt.Sprintf("    \x1b[90m%s\x1b[0m\n", g.Description))
		}
		for _, b := range g.Buckets {
			bar, pctStr := renderProgressBar(b.RemainingFraction)
			resetInfo := formatResetTime(b.ResetTime, now)
			windowLabel := fmt.Sprintf("%-6s", b.Window)
			sb.WriteString(fmt.Sprintf("    %-8s %s %s %s\n", windowLabel, bar, pctStr, resetInfo))
		}
		sb.WriteString("\n")
	}

	// Card 3: Gemini Image Generation (Dedicated field for image generation / editing)
	if imageGemini != nil {
		sb.WriteString("  • \x1b[1;35mGemini Image Generation\x1b[0m\n")
		sb.WriteString("    \x1b[90mSpecial quota bucket for image generation & editing\x1b[0m\n")
		bar, pctStr := renderProgressBar(imageGemini.RemainingFraction)
		resetInfo := formatResetTime(imageGemini.ResetTime, now)
		sb.WriteString(fmt.Sprintf("    %-24s %s %s %s\n", imageGemini.ModelID, bar, pctStr, resetInfo))
		sb.WriteString("\n")
	}

	sb.WriteString("  ───────────────────────────────────────────────────────────────\n")
	return sb.String()
}

// RenderQuotaRaw renders every raw quota bucket from Google CloudCode API (for debugging / verbose mode)
func RenderQuotaRaw(groups []QuotaGroup, detailed []DetailedQuotaBucket) string {
	if len(groups) == 0 && len(detailed) == 0 {
		return "  [!] No quota data returned by Google CloudCode API.\n"
	}

	var sb strings.Builder
	sb.WriteString("\n  ═══════════════════════════════════════════════════════════════\n")
	sb.WriteString("    ANTIGRAVITY GOOGLE CLOUD RAW DETAILED QUOTA DUMP\n")
	sb.WriteString("  ═══════════════════════════════════════════════════════════════\n\n")

	now := time.Now().UTC()

	if len(groups) > 0 {
		for _, g := range groups {
			sb.WriteString(fmt.Sprintf("  • \x1b[1;36m%s\x1b[0m\n", g.DisplayName))
			if g.Description != "" {
				sb.WriteString(fmt.Sprintf("    \x1b[90m%s\x1b[0m\n", g.Description))
			}
			for _, b := range g.Buckets {
				bar, pctStr := renderProgressBar(b.RemainingFraction)
				resetInfo := formatResetTime(b.ResetTime, now)
				windowLabel := fmt.Sprintf("%-6s", b.Window)
				sb.WriteString(fmt.Sprintf("    %-8s %s %s %s\n", windowLabel, bar, pctStr, resetInfo))
			}
			sb.WriteString("\n")
		}
	}

	if len(detailed) > 0 {
		var geminiBuckets []DetailedQuotaBucket
		var otherBuckets []DetailedQuotaBucket

		for _, b := range detailed {
			if strings.HasPrefix(b.ModelID, "gemini-") {
				geminiBuckets = append(geminiBuckets, b)
			} else {
				otherBuckets = append(otherBuckets, b)
			}
		}

		if len(geminiBuckets) > 0 {
			sb.WriteString("  • \x1b[1;32mGoogle Gemini Models (Raw API Buckets)\x1b[0m\n")
			for _, b := range geminiBuckets {
				sb.WriteString(renderBucketLine(b, now))
			}
			sb.WriteString("\n")
		}

		if len(otherBuckets) > 0 {
			sb.WriteString("  • \x1b[1;35mCompanion & Special Models (Raw API Buckets)\x1b[0m\n")
			for _, b := range otherBuckets {
				sb.WriteString(renderBucketLine(b, now))
			}
			sb.WriteString("\n")
		}
	}

	sb.WriteString("  ───────────────────────────────────────────────────────────────\n")
	return sb.String()
}

func renderProgressBar(rem *float64) (string, string) {
	if rem == nil {
		return "[░░░░░░░░░░]", "N/A "
	}
	pct := int(*rem * 100)
	if pct > 100 {
		pct = 100
	} else if pct < 0 {
		pct = 0
	}
	pctStr := fmt.Sprintf("%3d%%", pct)
	filled := pct / 10
	if filled > 10 {
		filled = 10
	}
	empty := 10 - filled
	color := "\x1b[32m" // Green
	if pct < 20 {
		color = "\x1b[31m" // Red
	} else if pct < 50 {
		color = "\x1b[33m" // Yellow
	}
	bar := fmt.Sprintf("[%s%s\x1b[90m%s\x1b[0m]", color, strings.Repeat("█", filled), strings.Repeat("░", empty))
	return bar, pctStr
}

func formatResetTime(reset time.Time, now time.Time) string {
	if reset.IsZero() {
		return ""
	}
	diff := reset.Sub(now)
	if diff > 0 {
		hours := int(diff.Hours())
		mins := int(diff.Minutes()) % 60
		if hours >= 24 {
			days := hours / 24
			remHours := hours % 24
			return fmt.Sprintf("(resets in %dd %dh)", days, remHours)
		}
		return fmt.Sprintf("(resets in %dh %dm)", hours, mins)
	}
	return "(reset pending)"
}

func renderBucketLine(b DetailedQuotaBucket, now time.Time) string {
	bar, pctStr := renderProgressBar(b.RemainingFraction)
	resetInfo := formatResetTime(b.ResetTime, now)
	return fmt.Sprintf("    %-28s %s %s %s\n", b.ModelID, bar, pctStr, resetInfo)
}
