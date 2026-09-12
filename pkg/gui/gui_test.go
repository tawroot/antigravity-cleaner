package gui

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestGuiServerAndEndpoints(t *testing.T) {
	server, err := NewServer()
	if err != nil {
		t.Fatalf("Failed to create GUI server: %v", err)
	}
	defer server.Close()

	go func() {
		_ = server.Start()
	}()

	client := &http.Client{}

	// Test 1: Root embedded HTML page
	resp, err := client.Get(server.URL())
	if err != nil {
		t.Fatalf("Failed to GET root URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read body: %v", err)
	}

	bodyStr := string(bodyBytes)
	if !strings.Contains(bodyStr, "DALROOT") {
		t.Errorf("Expected embedded HTML to contain DALROOT logo text")
	}
	if !strings.Contains(bodyStr, "Antigravity PATCH") {
		t.Errorf("Expected embedded HTML to contain Antigravity PATCH title")
	}

	// Test 2: /api/detect endpoint
	detectResp, err := client.Get(server.URL() + "api/detect")
	if err != nil {
		t.Fatalf("Failed to GET /api/detect: %v", err)
	}
	defer detectResp.Body.Close()

	if detectResp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK from /api/detect, got %d", detectResp.StatusCode)
	}

	var dResp DetectResponse
	if err := json.NewDecoder(detectResp.Body).Decode(&dResp); err != nil {
		t.Fatalf("Failed to decode /api/detect JSON: %v", err)
	}

	if dResp.App == "" {
		t.Errorf("Expected non-empty App in detect response")
	}
	if dResp.OS == "" {
		t.Errorf("Expected non-empty OS in detect response")
	}
}
