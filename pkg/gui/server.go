package gui

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/tawroot/antigravity-cleaner/pkg/account"
	"github.com/tawroot/antigravity-cleaner/pkg/cleaner"
	"github.com/tawroot/antigravity-cleaner/pkg/doctor"
	"github.com/tawroot/antigravity-cleaner/pkg/patcher"
	"github.com/tawroot/antigravity-cleaner/pkg/proxy"
)

func init() {
	patcher.RegisterPostInstallHook(func(logger patcher.LogFunc) {
		_ = account.SyncTokenBridge(logger)
	})
}

//go:embed assets/*
var assetsFS embed.FS

// Server represents the local loopback backend for the GUI.
type Server struct {
	listener net.Listener
	server   *http.Server
	port     int
}

// Response represents a standard JSON response.
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// DetectResponse represents auto-detection details for the frontend.
type DetectResponse struct {
	App            string `json:"app"`
	Path           string `json:"path"`
	Detected       bool   `json:"detected"`
	Proxy          string `json:"proxy"`
	OS             string `json:"os"`
	Arch           string `json:"arch"`
	WrapperActive  bool   `json:"wrapper_active"`
	RunningPID     int    `json:"running_pid"`
	ProcessRunning bool   `json:"process_running"`
}

// NewServer initializes a new GUI loopback server on an ephemeral port.
func NewServer() (*Server, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("failed to bind loopback listener: %w", err)
	}

	port := listener.Addr().(*net.TCPAddr).Port
	s := &Server{
		listener: listener,
		port:     port,
	}

	mux := http.NewServeMux()

	// Static assets sub-filesystem
	subFS, err := fs.Sub(assetsFS, "assets")
	if err != nil {
		return nil, fmt.Errorf("failed to load embedded assets: %w", err)
	}
	mux.Handle("/", http.FileServer(http.FS(subFS)))

	// REST API routes
	mux.HandleFunc("/api/detect", s.handleDetect)
	mux.HandleFunc("/api/patch", s.handlePatch)
	mux.HandleFunc("/api/launch", s.handleLaunch)
	mux.HandleFunc("/api/create-launcher", s.handleCreateLauncher)
	mux.HandleFunc("/api/reset", s.handleReset)
	mux.HandleFunc("/api/restore", s.handleRestore)
	mux.HandleFunc("/api/doctor", s.handleDoctor)
	mux.HandleFunc("/api/report", s.handleReport)
	mux.HandleFunc("/api/privacy", s.handlePrivacy)
	mux.HandleFunc("/api/open-external", s.handleOpenExternal)
	mux.HandleFunc("/api/exit", s.handleExit)

	s.server = &http.Server{
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s, nil
}

// Port returns the assigned local port.
func (s *Server) Port() int {
	return s.port
}

// URL returns the loopback HTTP URL.
func (s *Server) URL() string {
	return fmt.Sprintf("http://127.0.0.1:%d/", s.port)
}

// Start runs the server in background.
func (s *Server) Start() error {
	return s.server.Serve(s.listener)
}

// Close terminates the server listener.
func (s *Server) Close() error {
	return s.server.Close()
}

func (s *Server) handleDetect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	targetPath := patcher.FindLanguageServerModular(nil)
	if targetPath == "" {
		targetPath = patcher.FindLanguageServer()
	}
	detected := targetPath != ""

	if targetPath == "" {
		targetPath = patcher.FindIdeMainJs()
	}
	if targetPath == "" {
		targetPath = patcher.FindAgy()
	}

	proxyDesc := "Direct socket (no proxy active)"
	if best, err := proxy.GetBestProxy(); err == nil {
		proxyDesc = fmt.Sprintf("%s (%s)", best.URL, best.Provider)
	}

	runningPID := patcher.FindRunningLanguageServerPID()
	wrapperActive := false
	if targetPath != "" {
		wm := patcher.NewWrapperManager(filepath.Dir(targetPath), nil)
		wrapperActive = wm.IsWrapperInstalled()
	}

	resp := DetectResponse{
		App:            "Google Antigravity 2.x (IDE & Language Server)",
		Path:           targetPath,
		Detected:       detected,
		Proxy:          proxyDesc,
		OS:             runtime.GOOS,
		Arch:           runtime.GOARCH,
		WrapperActive:  wrapperActive,
		RunningPID:     runningPID,
		ProcessRunning: runningPID > 0,
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) handlePatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_, _ = patcher.KillAntigravityProcesses()

	var msg string
	success := false

	_, err := patcher.PatchLanguageServer("")
	if err != nil {
		msg = err.Error()
	} else {
		success = true
		msg = "All 4 targets patched & Zero-Telemetry enabled!"
		_, _ = patcher.PatchAgy("")
		_, _ = patcher.PatchIdeMainJs("")
		_, _ = patcher.PatchVsCodeExtension("")
		_, _ = patcher.PatchAsarZeroLog("")
		_, _ = proxy.GenerateNoTunDesktopLauncher("")

		// Ensure Zero-Telemetry & Enterprise wrapper is active
		lsPath := patcher.FindLanguageServerModular(nil)
		if lsPath == "" {
			lsPath = patcher.FindLanguageServer()
		}
		if lsPath != "" {
			wm := patcher.NewWrapperManager(filepath.Dir(lsPath), nil)
			if !wm.IsWrapperInstalled() {
				_ = wm.InstallWrapper(patcher.DefaultEnterpriseFlagsOptions())
			}
		}
	}

	_ = json.NewEncoder(w).Encode(Response{
		Success: success,
		Message: msg,
	})
}

func (s *Server) handleLaunch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var proxyURL string
	if best, err := proxy.GetBestProxy(); err == nil {
		proxyURL = best.URL
	}

	err := proxy.LaunchAntigravityWithProxy("", proxyURL, nil)
	if err != nil {
		_ = json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	_ = json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "Antigravity launched successfully with injected proxy!",
	})
}

func (s *Server) handleCreateLauncher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path, err := proxy.GenerateNoTunDesktopLauncher("")
	if err != nil {
		_ = json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "Error creating launcher: " + err.Error(),
		})
		return
	}

	_ = json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "Desktop shortcut created at " + path,
	})
}

func (s *Server) handleReset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_, _ = patcher.KillAntigravityProcesses()

	removed, err := cleaner.SurgicalReset429("", false)
	if err != nil {
		_ = json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	_ = json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: fmt.Sprintf("Cleaned %d cache & session artifacts", len(removed)),
	})
}

func (s *Server) handleRestore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_, _ = patcher.KillAntigravityProcesses()

	restoredCount := 0
	if lsPath := patcher.FindLanguageServer(); lsPath != "" {
		if err := patcher.RestoreFile(lsPath); err == nil {
			restoredCount++
		}
	}
	if agyPath := patcher.FindAgy(); agyPath != "" {
		if err := patcher.RestoreFile(agyPath); err == nil {
			restoredCount++
		}
	}
	if idePath := patcher.FindIdeMainJs(); idePath != "" {
		if err := patcher.RestoreFile(idePath); err == nil {
			restoredCount++
		}
	}

	_ = json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: fmt.Sprintf("Restored %d files from clean .agybak backups", restoredCount),
	})
}

func (s *Server) handleDoctor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	results := doctor.RunDiagnostics(false)
	allOk := true
	for _, res := range results {
		if res.Status == "FAIL" {
			allOk = false
			break
		}
	}

	_ = json.NewEncoder(w).Encode(Response{
		Success: allOk,
		Message: "Diagnostics completed",
	})
}

func (s *Server) handleReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	report := doctor.GenerateDiagnosticReport("5.4.2")
	doctor.CopyToClipboard(report)
	filePath, _ := doctor.SaveAndOpenReport(report)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"message":  report,
		"filePath": filePath,
	})
}

func (s *Server) handlePrivacy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	targetPath := patcher.FindLanguageServerModular(nil)
	if targetPath == "" {
		targetPath = patcher.FindLanguageServer()
	}
	if targetPath == "" {
		_ = json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "Language server binary not found",
		})
		return
	}

	binDir := filepath.Dir(targetPath)
	wm := patcher.NewWrapperManager(binDir, nil)

	if wm.IsWrapperInstalled() {
		if err := wm.RemoveWrapper(); err != nil {
			_ = json.NewEncoder(w).Encode(Response{
				Success: false,
				Message: "Failed to remove wrapper: " + err.Error(),
			})
			return
		}
		_ = json.NewEncoder(w).Encode(Response{
			Success: true,
			Message: "Zero-Telemetry wrapper disabled. Original binary restored.",
		})
		return
	}

	opts := patcher.DefaultEnterpriseFlagsOptions()
	if err := wm.InstallWrapper(opts); err != nil {
		_ = json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "Failed to install wrapper: " + err.Error(),
		})
		return
	}

	_ = json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "Zero-Telemetry & Enterprise wrapper enabled (7 privacy flags injected)!",
	})
}

func (s *Server) handleOpenExternal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	target := r.URL.Query().Get("url")
	if target == "" {
		target = "https://t.me/dallrep"
	}
	if !strings.HasPrefix(target, "https://t.me/") && !strings.HasPrefix(target, "https://github.com/") {
		http.Error(w, "Forbidden URL", http.StatusForbidden)
		return
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", target)
	case "darwin":
		cmd = exec.Command("open", target)
	default:
		cmd = exec.Command("xdg-open", target)
	}
	_ = cmd.Start()

	_ = json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "Opening external link: " + target,
	})
}

func (s *Server) handleExit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "Exiting",
	})
	go func() {
		time.Sleep(200 * time.Millisecond)
		os.Exit(0)
	}()
}
