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

	// Build arguments: pass --proxy-server to Chromium/Electron
	args := []string{
		fmt.Sprintf("--proxy-server=%s", proxyURL),
	}
	args = append(args, extraArgs...)

	cmd := exec.Command(appPath, args...)
	cmd.Env = BuildProxyEnv(proxyURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start Antigravity with proxy: %w", err)
	}

	return nil
}

// GenerateNoTunDesktopLauncher creates a .desktop shortcut or Windows script
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
		desktopAppDir := filepath.Join(home, ".local", "share", "applications")
		_ = os.MkdirAll(desktopAppDir, 0755)
		appMenuPath := filepath.Join(desktopAppDir, "antigravity.desktop")

		iconPath := filepath.Join(filepath.Dir(appPath), "icon.png")
		if fi, err := os.Stat(iconPath); err != nil || fi.Size() < 100 {
			userIcon := filepath.Join(home, ".local", "share", "icons", "antigravity.png")
			iconPath = EnsureDefaultIcon(userIcon)
		}

		content := fmt.Sprintf(`[Desktop Entry]
Version=1.0
Name=Antigravity
GenericName=AI Code Editor
Comment=Google Antigravity with Smart No-TUN Proxy
Exec=env ALL_PROXY=%s HTTPS_PROXY=%s HTTP_PROXY=%s %s --proxy-server="%s"
Icon=%s
Terminal=false
Type=Application
Categories=Development;IDE;
StartupNotify=true
StartupWMClass=antigravity
`, proxyURL, proxyURL, proxyURL, appPath, proxyURL, iconPath)

		if err := os.WriteFile(appMenuPath, []byte(content), 0755); err != nil {
			return "", err
		}

		// Also create on ~/Desktop if Desktop directory exists
		desktopDir := filepath.Join(home, "Desktop")
		createdPaths := appMenuPath
		if fi, err := os.Stat(desktopDir); err == nil && fi.IsDir() {
			desktopShortcut := filepath.Join(desktopDir, "Antigravity.desktop")
			_ = os.WriteFile(desktopShortcut, []byte(content), 0755)
			_ = exec.Command("gio", "set", desktopShortcut, "metadata::trusted", "true").Run()
			createdPaths = fmt.Sprintf("%s and Desktop shortcut", appMenuPath)
		}

		// Refresh GNOME/KDE application database
		_ = exec.Command("update-desktop-database", desktopAppDir).Run()

		return createdPaths, nil
	}

	if runtime.GOOS == "windows" {
		batPath := filepath.Join(home, "Desktop", "Antigravity-NoTUN.bat")
		content := fmt.Sprintf(`@echo off
set "ALL_PROXY=%s"
set "HTTPS_PROXY=%s"
set "HTTP_PROXY=%s"
start "" "%s" --proxy-server="%s"
`, proxyURL, proxyURL, proxyURL, appPath, proxyURL)

		if err := os.WriteFile(batPath, []byte(content), 0755); err != nil {
			return "", err
		}
		return batPath, nil
	}

	return "", fmt.Errorf("desktop launcher generation not supported on %s", runtime.GOOS)
}
