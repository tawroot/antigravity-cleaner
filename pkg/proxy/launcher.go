package proxy

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// FindAntigravityExecutable finds the main Antigravity GUI application
func FindAntigravityExecutable() string {
	home, _ := os.UserHomeDir()

	candidates := []string{
		filepath.Join(home, "Applications", "Antigravity", "antigravity"),
		filepath.Join(home, "Applications", "Antigravity-x64", "antigravity"),
		filepath.Join(home, "antigravity", "antigravity"),
		filepath.Join(home, "antigravity", "antigravity.exe"),
		"/usr/share/antigravity-ide/antigravity",
		"/usr/bin/antigravity",
		"/opt/Antigravity/antigravity",
		"/opt/Antigravity IDE/antigravity",
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Antigravity", "Antigravity.exe"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Antigravity IDE", "Antigravity IDE.exe"),
		"/Applications/Antigravity.app/Contents/MacOS/Electron",
		"/Applications/Antigravity.app/Contents/MacOS/Antigravity",
		"/Applications/Antigravity IDE.app/Contents/MacOS/Electron",
	}

	for _, c := range candidates {
		if fi, err := os.Stat(c); err == nil && !fi.IsDir() {
			return c
		}
	}
	return ""
}

// BuildProxyEnv constructs clean proxy environment variables for child processes
func BuildProxyEnv(proxyURL string) []string {
	env := os.Environ()

	var httpProxy string
	var socksProxy string

	if strings.HasPrefix(proxyURL, "socks") {
		socksProxy = proxyURL
		// If socks5h, ensure Go and curl understand remote DNS
		if !strings.HasPrefix(proxyURL, "socks5h://") && strings.HasPrefix(proxyURL, "socks5://") {
			socksProxy = strings.Replace(proxyURL, "socks5://", "socks5h://", 1)
		}
		// Also provide standard HTTP proxy fallback if possible
		httpProxy = socksProxy
	} else {
		httpProxy = proxyURL
		socksProxy = proxyURL
	}

	overrides := map[string]string{
		"ALL_PROXY":   socksProxy,
		"all_proxy":   socksProxy,
		"HTTPS_PROXY": httpProxy,
		"https_proxy": httpProxy,
		"HTTP_PROXY":  httpProxy,
		"http_proxy":  httpProxy,
		"GRPC_PROXY":  httpProxy,
	}

	// Filter existing proxy envs
	var cleanEnv []string
	for _, e := range env {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) > 0 {
			keyUpper := strings.ToUpper(parts[0])
			if _, isProxy := overrides[keyUpper]; isProxy {
				continue
			}
		}
		cleanEnv = append(cleanEnv, e)
	}

	for k, v := range overrides {
		cleanEnv = append(cleanEnv, fmt.Sprintf("%s=%s", k, v))
	}

	return cleanEnv
}

// CleanChromiumProxyURL normalizes proxy schemes for Chromium/Electron.
// Chromium does NOT support socks5h:// (it causes ERR_NO_SUPPORTED_PROXIES).
// Chromium's socks5:// implementation automatically performs DNS resolution on the proxy.
func CleanChromiumProxyURL(proxyURL string) string {
	return strings.Replace(proxyURL, "socks5h://", "socks5://", 1)
}

// LaunchAntigravityWithProxy launches Antigravity with injected proxy variables (No TUN needed)
func LaunchAntigravityWithProxy(appPath string, proxyURL string, extraArgs []string) error {
	if appPath == "" {
		appPath = FindAntigravityExecutable()
		if appPath == "" {
			return fmt.Errorf("Antigravity executable not found")
		}
	}

	if proxyURL == "" {
		best, err := GetBestProxy()
		if err != nil {
			return err
		}
		proxyURL = best.URL
	}

	chromiumProxy := CleanChromiumProxyURL(proxyURL)

	// Build arguments: pass --proxy-server to Chromium/Electron without quotes
	args := []string{
		fmt.Sprintf("--proxy-server=%s", chromiumProxy),
	}
	args = append(args, extraArgs...)

	cmd := exec.Command(appPath, args...)
	cmd.Dir = filepath.Dir(appPath)
	cmd.Env = BuildProxyEnv(proxyURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start Antigravity with proxy: %w", err)
	}

	return nil
}

// GenerateNoTunDesktopLauncher creates .desktop shortcuts for both Antigravity and Antigravity Cleaner GUI
func GenerateNoTunDesktopLauncher(proxyURL string) (string, error) {
	appPath := FindAntigravityExecutable()
	if appPath == "" {
		return "", fmt.Errorf("Antigravity executable not found")
	}

	if proxyURL == "" {
		best, err := GetBestProxy()
		if err != nil {
			return "", err
		}
		proxyURL = best.URL
	}

	home, _ := os.UserHomeDir()

	if runtime.GOOS == "linux" {
		binDir := filepath.Join(home, ".local", "bin")
		_ = os.MkdirAll(binDir, 0755)

		desktopAppDir := filepath.Join(home, ".local", "share", "applications")
		_ = os.MkdirAll(desktopAppDir, 0755)

		cleanerBin, _ := os.Executable()
		if cleanerBin == "" || strings.Contains(cleanerBin, "/tmp/") || strings.Contains(cleanerBin, "go-build") {
			cleanerBin = filepath.Join(binDir, "antigravity-cleaner")
		}

		// 1. Create Smart Launcher script in ~/.local/bin/antigravity-smart-launcher.sh
		smartLauncherScript := filepath.Join(binDir, "antigravity-smart-launcher.sh")
		scriptContent := fmt.Sprintf(`#!/usr/bin/env bash
# Smart Proxy Launcher for Google Antigravity
# Auto-detects running proxy (v2ray, NekoBox, Clash, Hiddify) and launches Antigravity

APP_PATH="%s"

# 1. Clean stale singleton locks from ungraceful exits or display disconnects
LOCK_FILE="$HOME/.config/Antigravity/SingletonLock"
if [ -L "$LOCK_FILE" ]; then
    LOCK_TARGET=$(readlink "$LOCK_FILE" 2>/dev/null)
    LOCK_PID="${LOCK_TARGET##*-}"
    if [ -n "$LOCK_PID" ] && ! kill -0 "$LOCK_PID" 2>/dev/null; then
        rm -f "$LOCK_FILE" "$HOME/.config/Antigravity/SingletonSocket" "$HOME/.config/Antigravity/SingletonCookie" 2>/dev/null
    fi
fi

# 2. Dynamic proxy port detection
CANDIDATES=("10808" "2080" "7890" "1080" "10809" "2081")
ACTIVE_PORT=""

for port in "${CANDIDATES[@]}"; do
    if (echo > /dev/tcp/127.0.0.1/$port) >/dev/null 2>&1; then
        ACTIVE_PORT="$port"
        break
    fi
done

# 3. Resilient fallback: Never abort or lock out the user!
if [ -z "$ACTIVE_PORT" ]; then
    ACTIVE_PORT="10808"
    if command -v notify-send >/dev/null 2>&1; then
        notify-send "Antigravity" "⚠️ پروکسی فعالی شناسایی نشد؛ در حال استفاده از پورت پیش‌فرض 10808..." -i dialog-warning -t 3000 &
    fi
fi

PROXY_TYPE="socks5"
if [ "$ACTIVE_PORT" = "10809" ] || [ "$ACTIVE_PORT" = "2081" ]; then
    PROXY_TYPE="http"
fi

export ALL_PROXY="${PROXY_TYPE}h://127.0.0.1:${ACTIVE_PORT}"
export HTTPS_PROXY="${PROXY_TYPE}h://127.0.0.1:${ACTIVE_PORT}"
export HTTP_PROXY="${PROXY_TYPE}h://127.0.0.1:${ACTIVE_PORT}"

# Pass socks5:// without quotes to Chromium (Chromium does not support socks5h)
exec "$APP_PATH" --proxy-server="${PROXY_TYPE}://127.0.0.1:${ACTIVE_PORT}" "$@"
`, appPath)

		_ = os.WriteFile(smartLauncherScript, []byte(scriptContent), 0755)

		// 2. Setup icon
		iconPath := filepath.Join(filepath.Dir(appPath), "icon.png")
		if fi, err := os.Stat(iconPath); err != nil || fi.Size() < 100 {
			userIcon := filepath.Join(home, ".local", "share", "icons", "antigravity.png")
			iconPath = EnsureDefaultIcon(userIcon)
		}

		// 3. Create single Antigravity IDE Desktop Launcher
		appMenuPath := filepath.Join(desktopAppDir, "antigravity.desktop")
		contentAntigravity := fmt.Sprintf(`[Desktop Entry]
Version=1.0
Name=Antigravity
GenericName=AI Code Editor
Comment=Google Antigravity with Smart No-TUN Proxy
Exec=%s %%F
Icon=%s
Path=%s
Terminal=false
Type=Application
Categories=Development;IDE;
StartupNotify=true
StartupWMClass=antigravity
`, smartLauncherScript, iconPath, filepath.Dir(appPath))

		if err := os.WriteFile(appMenuPath, []byte(contentAntigravity), 0755); err != nil {
			return "", err
		}

		// 4. Also create on ~/Desktop if Desktop directory exists
		desktopDir := filepath.Join(home, "Desktop")
		createdPaths := appMenuPath
		if fi, err := os.Stat(desktopDir); err == nil && fi.IsDir() {
			desktopShortcut := filepath.Join(desktopDir, "Antigravity.desktop")
			_ = os.WriteFile(desktopShortcut, []byte(contentAntigravity), 0755)
			_ = exec.Command("gio", "set", desktopShortcut, "metadata::trusted", "true").Run()

			createdPaths = fmt.Sprintf("%s and Desktop shortcut", appMenuPath)
		}

		// Refresh GNOME/KDE application database
		_ = exec.Command("update-desktop-database", desktopAppDir).Run()

		return createdPaths, nil
	}

	if runtime.GOOS == "darwin" {
		cmdPath := filepath.Join(home, "Desktop", "Antigravity-Proxy.command")
		content := fmt.Sprintf(`#!/usr/bin/env bash
# macOS Smart Proxy Launcher for Google Antigravity
APP_PATH="%s"
if [ ! -f "$APP_PATH" ]; then
    APP_PATH="/Applications/Antigravity.app/Contents/MacOS/Antigravity"
fi
if [ ! -f "$APP_PATH" ]; then
    APP_PATH="/Applications/Antigravity.app/Contents/MacOS/Electron"
fi

CANDIDATES=("10808" "2080" "7890" "1080")
ACTIVE_PORT=""
for port in "${CANDIDATES[@]}"; do
    if nc -z -w 1 127.0.0.1 "$port" 2>/dev/null; then
        ACTIVE_PORT="$port"
        break
    fi
done

if [ -z "$ACTIVE_PORT" ]; then
    ACTIVE_PORT="10808"
fi

export ALL_PROXY="socks5h://127.0.0.1:${ACTIVE_PORT}"
export HTTPS_PROXY="socks5h://127.0.0.1:${ACTIVE_PORT}"
export HTTP_PROXY="socks5h://127.0.0.1:${ACTIVE_PORT}"

exec "$APP_PATH" --proxy-server="socks5://127.0.0.1:${ACTIVE_PORT}" "$@"
`, appPath)

		if err := os.WriteFile(cmdPath, []byte(content), 0755); err != nil {
			return "", err
		}
		return cmdPath, nil
	}

	if runtime.GOOS == "windows" {
		chromiumProxy := CleanChromiumProxyURL(proxyURL)
		batPath := filepath.Join(home, "Desktop", "Antigravity-NoTUN.bat")
		content := fmt.Sprintf(`@echo off
set "ALL_PROXY=%s"
set "HTTPS_PROXY=%s"
set "HTTP_PROXY=%s"
start "" "%s" --proxy-server=%s
`, proxyURL, proxyURL, proxyURL, appPath, chromiumProxy)

		if err := os.WriteFile(batPath, []byte(content), 0755); err != nil {
			return "", err
		}

		// Also generate native Windows shortcut (.lnk) via PowerShell if possible
		lnkScript := fmt.Sprintf(`
$WshShell = New-Object -comObject WScript.Shell
$Shortcut = $WshShell.CreateShortcut("%s\Desktop\Antigravity.lnk")
$Shortcut.TargetPath = "%s"
$Shortcut.Arguments = "--proxy-server=%s"
$Shortcut.WorkingDirectory = "%s"
$Shortcut.IconLocation = "%s,0"
$Shortcut.Description = "Google Antigravity with Smart Proxy"
$Shortcut.Save()
`, home, appPath, chromiumProxy, filepath.Dir(appPath), appPath)
		_ = exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", lnkScript).Run()

		return batPath, nil
	}

	return "", fmt.Errorf("desktop launcher generation not supported on %s", runtime.GOOS)
}
