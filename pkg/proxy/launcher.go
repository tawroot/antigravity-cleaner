package proxy

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// AppTarget describes an installed Antigravity application target (Standard 2.0 or IDE)
type AppTarget struct {
	Name    string // "Antigravity" or "Antigravity IDE"
	ExePath string
	IsIDE   bool
}

// FindAllAntigravityExecutables finds all distinct Antigravity GUI application targets installed
func FindAllAntigravityExecutables() []AppTarget {
	home, _ := os.UserHomeDir()
	var targets []AppTarget
	seen := make(map[string]bool)

	add := func(name, path string, isIDE bool) {
		if path == "" {
			return
		}
		if real, err := filepath.EvalSymlinks(path); err == nil && real != "" {
			path = real
		}
		clean := filepath.Clean(path)
		if !seen[clean] {
			if fi, err := os.Stat(clean); err == nil && !fi.IsDir() {
				seen[clean] = true
				targets = append(targets, AppTarget{
					Name:    name,
					ExePath: clean,
					IsIDE:   isIDE,
				})
			}
		}
	}

	// 1. Windows direct program checks
	if runtime.GOOS == "windows" {
		localAppData := os.Getenv("LOCALAPPDATA")
		progFiles := os.Getenv("PROGRAMFILES")
		progFilesX86 := os.Getenv("ProgramFiles(x86)")

		// Antigravity 2.0 (Official app)
		add("Antigravity", filepath.Join(localAppData, "Programs", "antigravity", "Antigravity.exe"), false)
		add("Antigravity", filepath.Join(localAppData, "Programs", "Antigravity", "Antigravity.exe"), false)
		add("Antigravity", filepath.Join(progFiles, "Antigravity", "Antigravity.exe"), false)
		add("Antigravity", filepath.Join(progFilesX86, "Antigravity", "Antigravity.exe"), false)

		// Antigravity IDE (VS Code based)
		add("Antigravity IDE", filepath.Join(localAppData, "Programs", "Antigravity IDE", "Antigravity IDE.exe"), true)
		add("Antigravity IDE", filepath.Join(progFiles, "Antigravity IDE", "Antigravity IDE.exe"), true)
	}

	// 2. Linux checks
	if runtime.GOOS == "linux" {
		// Antigravity 2.0
		add("Antigravity", filepath.Join(home, "Applications", "Antigravity", "antigravity"), false)
		add("Antigravity", filepath.Join(home, "Applications", "antigravity", "antigravity"), false)
		add("Antigravity", filepath.Join(home, "Applications", "Antigravity-x64", "antigravity"), false)
		add("Antigravity", "/usr/share/antigravity/bin/antigravity", false)
		add("Antigravity", "/usr/share/antigravity/antigravity", false)
		add("Antigravity", "/opt/antigravity/antigravity", false)
		add("Antigravity", "/opt/Antigravity/antigravity", false)
		add("Antigravity", "/usr/local/bin/antigravity", false)
		add("Antigravity", "/usr/bin/antigravity", false)

		// Antigravity IDE
		add("Antigravity IDE", "/usr/share/antigravity-ide/antigravity", true)
		add("Antigravity IDE", "/opt/antigravity-ide/antigravity", true)
		add("Antigravity IDE", "/opt/Antigravity IDE/antigravity", true)
	}

	// 3. macOS checks
	if runtime.GOOS == "darwin" {
		add("Antigravity", "/Applications/Antigravity.app/Contents/MacOS/Antigravity", false)
		add("Antigravity", "/Applications/Antigravity.app/Contents/MacOS/Electron", false)
		add("Antigravity IDE", "/Applications/Antigravity IDE.app/Contents/MacOS/Electron", true)
	}

	// 4. Dynamic PATH discovery
	for _, name := range []string{"antigravity", "antigravity.exe"} {
		if path, err := exec.LookPath(name); err == nil && path != "" {
			add("Antigravity", path, false)
		}
	}
	if path, err := exec.LookPath("antigravity-ide"); err == nil && path != "" {
		add("Antigravity IDE", path, true)
	}

	return targets
}

// FindAntigravityExecutable finds the main Antigravity GUI application
func FindAntigravityExecutable() string {
	targets := FindAllAntigravityExecutables()
	if len(targets) == 0 {
		return ""
	}
	// Prefer standard Antigravity if available, otherwise IDE
	for _, t := range targets {
		if !t.IsIDE {
			return t.ExePath
		}
	}
	return targets[0].ExePath
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
	// Detach stdio pipes to prevent broken-pipe write EIO crashes when parent terminal closes
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Stdin = nil

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

	// Resolve symlinks on appPath before extracting icons/dir
	if realPath, err := filepath.EvalSymlinks(appPath); err == nil && realPath != "" {
		appPath = realPath
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

		appDir := filepath.Dir(appPath)
		workDir := appDir
		if strings.EqualFold(filepath.Base(appDir), "bin") {
			workDir = filepath.Dir(appDir)
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
# Redirect stdout/stderr to /dev/null to prevent electron-log write EIO broken-pipe popups
exec "$APP_PATH" --proxy-server="${PROXY_TYPE}://127.0.0.1:${ACTIVE_PORT}" "$@" >/dev/null 2>&1
`, appPath)

		_ = os.WriteFile(smartLauncherScript, []byte(scriptContent), 0755)

		// 2. Setup icon
		iconCandidates := []string{
			filepath.Join(appDir, "icon.png"),
			filepath.Join(filepath.Dir(appDir), "icon.png"),
			filepath.Join(appDir, "antigravity.png"),
			filepath.Join(filepath.Dir(appDir), "antigravity.png"),
			filepath.Join(filepath.Dir(appDir), "resources", "app", "resources", "linux", "code.png"),
			"/usr/share/pixmaps/antigravity.png",
		}

		var iconPath string
		for _, cand := range iconCandidates {
			if fi, err := os.Stat(cand); err == nil && fi.Size() >= 100 {
				iconPath = cand
				break
			}
		}

		if iconPath == "" {
			userIcon := filepath.Join(home, ".local", "share", "icons", "antigravity.png")
			iconPath = EnsureDefaultIcon(userIcon)
			hicolorIcon := filepath.Join(home, ".local", "share", "icons", "hicolor", "256x256", "apps", "antigravity.png")
			EnsureDefaultIcon(hicolorIcon)
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
`, smartLauncherScript, iconPath, workDir)

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

		// Generate native Windows shortcuts (.lnk) for all installed targets
		targets := FindAllAntigravityExecutables()
		if len(targets) == 0 {
			targets = []AppTarget{{Name: "Antigravity", ExePath: appPath}}
		}

		for _, t := range targets {
			lnkPath := filepath.Join(home, "Desktop", fmt.Sprintf("%s.lnk", t.Name))
			lnkScript := fmt.Sprintf(`
$WshShell = New-Object -comObject WScript.Shell
$Shortcut = $WshShell.CreateShortcut("%s")
$Shortcut.TargetPath = "%s"
$Shortcut.Arguments = "--proxy-server=%s"
$Shortcut.WorkingDirectory = "%s"
$Shortcut.IconLocation = "%s,0"
$Shortcut.Description = "Google %s with Smart Proxy"
$Shortcut.Save()
`, lnkPath, t.ExePath, chromiumProxy, filepath.Dir(t.ExePath), t.ExePath, t.Name)
			_ = exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", lnkScript).Run()
		}

		if len(targets) > 1 {
			return fmt.Sprintf("Desktop shortcuts created for Antigravity & Antigravity IDE (%s)", batPath), nil
		}
		return batPath, nil
	}

	return "", fmt.Errorf("desktop launcher generation not supported on %s", runtime.GOOS)
}
