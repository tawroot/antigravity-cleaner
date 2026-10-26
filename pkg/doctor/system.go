package doctor

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tawroot/antigravity-cleaner/pkg/patcher"
)

// DetectOSInfo returns a comprehensive, human-friendly OS & architecture description
func DetectOSInfo() string {
	arch := runtime.GOARCH
	switch runtime.GOOS {
	case "linux":
		if data, err := os.ReadFile("/etc/os-release"); err == nil {
			for _, line := range strings.Split(string(data), "\n") {
				if strings.HasPrefix(line, "PRETTY_NAME=") {
					name := strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), `"`)
					return fmt.Sprintf("%s [%s]", name, arch)
				}
			}
		}
		return fmt.Sprintf("Linux [%s]", arch)
	case "darwin":
		out, err := exec.Command("sw_vers", "-productVersion").Output()
		if err == nil {
			ver := strings.TrimSpace(string(out))
			archDesc := "Intel"
			if arch == "arm64" {
				archDesc = "Apple Silicon"
			}
			return fmt.Sprintf("macOS %s [%s %s]", ver, archDesc, arch)
		}
		return fmt.Sprintf("macOS [%s]", arch)
	case "windows":
		// Query Windows build details via registry or powershell
		if out, err := exec.Command("powershell", "-NoProfile", "-Command", "[System.Environment]::OSVersion.VersionString").Output(); err == nil {
			trimmed := strings.TrimSpace(string(out))
			if trimmed != "" {
				return fmt.Sprintf("%s [%s]", trimmed, arch)
			}
		}
		return fmt.Sprintf("Windows [%s]", arch)
	default:
		return fmt.Sprintf("%s [%s]", runtime.GOOS, arch)
	}
}

// DetectAntigravityVersion extracts the installed version of Antigravity IDE, 2.x, or CLI
func DetectAntigravityVersion() (string, string) {
	// 1. Try reading package.json inside app.asar
	if asarPath := patcher.FindAntigravityAsar(); asarPath != "" {
		if prod, ver, err := readAsarPackageJSON(asarPath); err == nil && ver != "" {
			return prod, ver
		}
	}

	// 2. Try reading package.json next to main.js (standalone/extracted IDE layout)
	if mainJs := patcher.FindIdeMainJs(); mainJs != "" {
		candidates := []string{
			filepath.Join(filepath.Dir(mainJs), "..", "package.json"),
			filepath.Join(filepath.Dir(mainJs), "package.json"),
			filepath.Join(filepath.Dir(mainJs), "..", "..", "package.json"),
		}
		for _, c := range candidates {
			if data, err := os.ReadFile(c); err == nil {
				var pkg struct {
					Name        string `json:"name"`
					ProductName string `json:"productName"`
					Version     string `json:"version"`
				}
				if json.Unmarshal(data, &pkg) == nil && pkg.Version != "" {
					prod := pkg.ProductName
					if prod == "" {
						prod = pkg.Name
					}
					if prod == "" {
						prod = "Antigravity IDE"
					}
					return prod, pkg.Version
				}
			}
		}
	}

	// 3. Try querying agy CLI
	if agyPath := patcher.FindAgy(); agyPath != "" {
		if out, err := exec.Command(agyPath, "--version").Output(); err == nil {
			v := strings.TrimSpace(string(out))
			if v != "" {
				return "Antigravity CLI", v
			}
		}
	}

	return "Antigravity", "Installed (Version Unspecified)"
}

func readAsarPackageJSON(asarPath string) (string, string, error) {
	f, err := os.Open(asarPath)
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	buf := make([]byte, 16)
	if _, err := f.ReadAt(buf, 0); err != nil {
		return "", "", err
	}
	headerSize := binary.LittleEndian.Uint32(buf[12:16])

	headerBuf := make([]byte, headerSize)
	if _, err := f.ReadAt(headerBuf, 16); err != nil {
		return "", "", err
	}

	var header struct {
		Files map[string]struct {
			Offset string `json:"offset"`
			Size   int    `json:"size"`
		} `json:"files"`
	}
	if err := json.Unmarshal(headerBuf, &header); err != nil {
		return "", "", err
	}

	pkgInfo, ok := header.Files["package.json"]
	if !ok {
		return "", "", fmt.Errorf("package.json not in asar root")
	}

	var offset int64
	fmt.Sscanf(pkgInfo.Offset, "%d", &offset)
	contentBuf := make([]byte, pkgInfo.Size+16)
	contentOffset := int64(16) + int64(headerSize) + offset
	if _, err := f.ReadAt(contentBuf, contentOffset); err != nil {
		return "", "", err
	}

	idx := bytes.IndexByte(contentBuf, '{')
	if idx == -1 {
		return "", "", fmt.Errorf("no json object found")
	}

	var pkg struct {
		Name        string `json:"name"`
		ProductName string `json:"productName"`
		Version     string `json:"version"`
	}
	if err := json.Unmarshal(contentBuf[idx:idx+pkgInfo.Size], &pkg); err != nil {
		return "", "", err
	}

	prod := pkg.ProductName
	if prod == "" {
		prod = pkg.Name
	}
	if prod == "" {
		prod = "Antigravity"
	}
	return prod, pkg.Version, nil
}
