package gui

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/tawroot/antigravity-cleaner/pkg/cleaner"
	"github.com/tawroot/antigravity-cleaner/pkg/doctor"
	"github.com/tawroot/antigravity-cleaner/pkg/patcher"
	"github.com/tawroot/antigravity-cleaner/pkg/proxy"
)

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
	App      string `json:"app"`
	Path     string `json:"path"`
	Detected bool   `json:"detected"`
	Proxy    string `json:"proxy"`
	OS       string `json:"os"`
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
	targetPath := patcher.FindLanguageServer()
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

	resp := DetectResponse{
		App:      "Google Antigravity 2.x (IDE & Language Server)",
		Path:     targetPath,
		Detected: detected,
		Proxy:    proxyDesc,
		OS:       runtime.GOOS,
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

	res, err := patcher.PatchLanguageServer("")
	if err != nil {
		msg = err.Error()
	} else {
		success = true
		msg = res
		// Full Auto-Fix: Patch all 4 targets + Proxy Launcher
		_, _ = patcher.PatchAgy("")
		_, _ = patcher.PatchIdeMainJs("")
		_, _ = patcher.PatchVsCodeExtension("")
		_, _ = proxy.GenerateNoTunDesktopLauncher("")
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
