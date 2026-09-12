package gui

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Run starts the embedded HTTP server and opens the retro desktop window.
func Run() error {
	server, err := NewServer()
	if err != nil {
		return fmt.Errorf("failed to initialize GUI server: %w", err)
	}

	go func() {
		_ = server.Start()
	}()
	defer server.Close()

	// Launch window in desktop app mode
	if err := LaunchWindow(server.URL()); err != nil {
		return fmt.Errorf("failed to open GUI window: %w", err)
	}

	// Wait for termination signal or exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	return nil
}
