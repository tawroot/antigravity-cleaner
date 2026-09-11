APP_NAME := antigravity-cleaner
VERSION := 5.1.0
BUILD_DIR := dist

.PHONY: all build clean test release

all: test build

build:
	@mkdir -p $(BUILD_DIR)
	go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME) ./cmd/antigravity

test:
	go test -v ./...

clean:
	rm -rf $(BUILD_DIR) bin/

release: clean
	@mkdir -p $(BUILD_DIR)
	@echo "Building Linux amd64..."
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 ./cmd/antigravity
	@echo "Building Linux arm64..."
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME)-linux-arm64 ./cmd/antigravity
	@echo "Building macOS amd64..."
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME)-darwin-amd64 ./cmd/antigravity
	@echo "Building macOS arm64 (Apple Silicon)..."
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME)-darwin-arm64 ./cmd/antigravity
	@echo "Building Windows amd64..."
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe ./cmd/antigravity
	@echo "All builds completed in $(BUILD_DIR)/"
