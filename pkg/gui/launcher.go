package gui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// LaunchWindow attempts to open the specified URL in a dedicated app-mode window
// (without browser tabs/address bar) across Linux, macOS, and Windows.
// If an app-mode browser is not found, it falls back to the system default browser.
func LaunchWindow(url string) error {
	var cmd *exec.Cmd

	// Temporary isolated directory ensures Chromium respects --window-size and does not inherit maximized state
	tempProfile := os.TempDir() + "/antigravity-gui-profile"

	switch runtime.GOOS {
	case "windows":
		// Microsoft Edge is guaranteed on Windows 10/11
		browsers := []string{
			os.Getenv("ProgramFiles(x86)") + `\Microsoft\Edge\Application\msedge.exe`,
			os.Getenv("ProgramFiles") + `\Microsoft\Edge\Application\msedge.exe`,
			os.Getenv("ProgramFiles") + `\Google\Chrome\Application\chrome.exe`,
			os.Getenv("LocalAppData") + `\Google\Chrome\Application\chrome.exe`,
			os.Getenv("ProgramFiles") + `\BraveSoftware\Brave-Browser\Application\brave.exe`,
		}
		for _, b := range browsers {
			if _, err := os.Stat(b); err == nil {
				cmd = exec.Command(b,
					fmt.Sprintf("--app=%s", url),
					"--window-size=440,510",
					fmt.Sprintf("--user-data-dir=%s", tempProfile),
					"--no-first-run",
					"--no-default-browser-check",
				)
				return cmd.Start()
			}
		}
		// Fallback for Windows
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
		return cmd.Start()

	case "darwin": // macOS
		macBrowsers := []string{
			"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
			"/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge",
			"/Applications/Brave Browser.app/Contents/MacOS/Brave Browser",
		}
		for _, b := range macBrowsers {
			if _, err := os.Stat(b); err == nil {
				cmd = exec.Command(b,
					fmt.Sprintf("--app=%s", url),
					"--window-size=440,510",
					fmt.Sprintf("--user-data-dir=%s", tempProfile),
					"--no-first-run",
					"--no-default-browser-check",
				)
				return cmd.Start()
			}
		}
		// Fallback for macOS
		cmd = exec.Command("open", url)
		return cmd.Start()

	default: // Linux / BSD
		linuxBrowsers := []string{
			"google-chrome",
			"google-chrome-stable",
			"chromium",
			"chromium-browser",
			"brave-browser",
			"microsoft-edge",
			"microsoft-edge-stable",
		}
		for _, b := range linuxBrowsers {
			if p, err := exec.LookPath(b); err == nil {
				cmd = exec.Command(p,
					fmt.Sprintf("--app=%s", url),
					"--window-size=440,510",
					fmt.Sprintf("--user-data-dir=%s", tempProfile),
					"--class=antigravity-patcher",
					"--no-first-run",
					"--no-default-browser-check",
				)
				return cmd.Start()
			}
		}
		// Fallback for Linux
		cmd = exec.Command("xdg-open", url)
		return cmd.Start()
	}
}
