package updater

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/tawroot/antigravity-cleaner/pkg/proxy"
)

const (
	RepoOwner = "tawroot"
	RepoName  = "antigravity-cleaner"
)

// ReleaseAsset represents a single downloadable binary release asset
type ReleaseAsset struct {
	Name        string `json:"name"`
	DownloadURL string `json:"browser_download_url"`
	Size        int64  `json:"size"`
}

// GitHubRelease represents GitHub Releases API response
type GitHubRelease struct {
	TagName     string         `json:"tag_name"`
	Name        string         `json:"name"`
	Body        string         `json:"body"`
	PublishedAt string         `json:"published_at"`
	Assets      []ReleaseAsset `json:"assets"`
}

func getHTTPClient() *http.Client {
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

	return &http.Client{
		Transport: transport,
		Timeout:   20 * time.Second,
	}
}

// GetMatchingAssetName returns the expected asset filename for the host platform
func GetMatchingAssetName() string {
	osName := runtime.GOOS
	arch := runtime.GOARCH

	switch osName {
	case "windows":
		return "antigravity-cleaner-windows-amd64.exe"
	case "darwin":
		if arch == "arm64" {
			return "antigravity-cleaner-darwin-arm64"
		}
		return "antigravity-cleaner-darwin-amd64"
	case "linux":
		if arch == "arm64" {
			return "antigravity-cleaner-linux-arm64"
		}
		return "antigravity-cleaner-linux-amd64"
	default:
		return fmt.Sprintf("antigravity-cleaner-%s-%s", osName, arch)
	}
}

// CheckLatestRelease queries GitHub API for the latest release
func CheckLatestRelease(currentVersion string) (*GitHubRelease, bool, error) {
	client := getHTTPClient()
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", RepoOwner, RepoName)

	req, err := http.NewRequestWithContext(context.Background(), "GET", apiURL, nil)
	if err != nil {
		return nil, false, err
	}
	req.Header.Set("User-Agent", "Antigravity-Cleaner-Updater/"+currentVersion)

	resp, err := client.Do(req)
	if err != nil {
		return nil, false, fmt.Errorf("failed to contact update server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("update server returned status HTTP %d", resp.StatusCode)
	}

	var release GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, false, fmt.Errorf("failed to parse release metadata: %w", err)
	}

	latestVer := strings.TrimPrefix(release.TagName, "v")
	curVer := strings.TrimPrefix(currentVersion, "v")

	hasUpdate := isNewerVersion(latestVer, curVer)
	return &release, hasUpdate, nil
}

// isNewerVersion returns true if latest is strictly newer than current
func isNewerVersion(latest, current string) bool {
	lParts := strings.Split(latest, ".")
	cParts := strings.Split(current, ".")

	for i := 0; i < len(lParts) && i < len(cParts); i++ {
		var l, c int
		fmt.Sscanf(lParts[i], "%d", &l)
		fmt.Sscanf(cParts[i], "%d", &c)
		if l > c {
			return true
		}
		if l < c {
			return false
		}
	}
	return len(lParts) > len(cParts)
}

// DownloadAndSelfUpdate replaces the current running binary with the latest release
func DownloadAndSelfUpdate(release *GitHubRelease, progressCb func(downloaded, total int64)) error {
	assetName := GetMatchingAssetName()
	var targetAsset *ReleaseAsset
	for _, a := range release.Assets {
		if a.Name == assetName {
			targetAsset = &a
			break
		}
	}

	if targetAsset == nil {
		return fmt.Errorf("no matching binary asset found for %s (%s)", assetName, release.TagName)
	}

	currentExe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to resolve current executable: %w", err)
	}
	currentExe, err = filepath.EvalSymlinks(currentExe)
	if err != nil {
		return fmt.Errorf("failed to resolve executable symlinks: %w", err)
	}

	exeDir := filepath.Dir(currentExe)
	tmpFile := filepath.Join(exeDir, fmt.Sprintf(".antigravity-cleaner-%s.tmp", release.TagName))

	// Download binary
	client := getHTTPClient()
	req, err := http.NewRequestWithContext(context.Background(), "GET", targetAsset.DownloadURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Antigravity-Cleaner-Updater")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download binary asset: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download server returned HTTP %d", resp.StatusCode)
	}

	out, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("cannot create temporary update file: %w", err)
	}

	var downloaded int64
	buf := make([]byte, 32*1024)
	for {
		n, rErr := resp.Body.Read(buf)
		if n > 0 {
			if _, wErr := out.Write(buf[:n]); wErr != nil {
				if cErr := out.Close(); cErr != nil {
					_ = os.Remove(tmpFile)
					return fmt.Errorf("write error: %v (close error: %w)", wErr, cErr)
				}
				_ = os.Remove(tmpFile)
				return wErr
			}
			downloaded += int64(n)
			if progressCb != nil {
				progressCb(downloaded, targetAsset.Size)
			}
		}
		if rErr != nil {
			if rErr == io.EOF {
				break
			}
			if cErr := out.Close(); cErr != nil {
				_ = os.Remove(tmpFile)
				return fmt.Errorf("read error: %v (close error: %w)", rErr, cErr)
			}
			_ = os.Remove(tmpFile)
			return rErr
		}
	}
	if cErr := out.Close(); cErr != nil {
		_ = os.Remove(tmpFile)
		return fmt.Errorf("failed to close update file: %w", cErr)
	}

	// Ensure executable permissions
	_ = os.Chmod(tmpFile, 0755)

	// Replace binary safely
	if runtime.GOOS == "windows" {
		oldFile := currentExe + ".old"
		_ = os.Remove(oldFile)
		if err := os.Rename(currentExe, oldFile); err != nil {
			_ = os.Remove(tmpFile)
			return fmt.Errorf("failed to replace binary on Windows (in-use lock): %w", err)
		}
		if err := os.Rename(tmpFile, currentExe); err != nil {
			_ = os.Rename(oldFile, currentExe)
			return fmt.Errorf("failed to commit new binary: %w", err)
		}
	} else {
		// POSIX atomic replacement
		if err := os.Rename(tmpFile, currentExe); err != nil {
			_ = os.Remove(tmpFile)
			return fmt.Errorf("failed to replace binary in-place: %w", err)
		}
	}

	return nil
}
