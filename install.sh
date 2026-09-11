#!/usr/bin/env bash
set -e

APP_NAME="antigravity-cleaner"
INSTALL_DIR="${HOME}/.local/bin"
REPO="tawroot/antigravity-cleaner"
VERSION="v5.1.1"

echo "===================================================="
echo "  ⚡ Antigravity Cleaner — Universal AI Freedom Toolkit (${VERSION})"
echo "  ✨ Universal Support: Auto-detects ALL Antigravity Versions (1.x, 2.x, IDE & CLI)"
echo "===================================================="

# 1. Detect OS and Architecture
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case "${ARCH}" in
    x86_64|amd64)
        ARCH="amd64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo "❌ Unsupported architecture: ${ARCH}"
        exit 1
        ;;
esac

TARGET_NAME="${APP_NAME}-${OS}-${ARCH}"
DEST="${INSTALL_DIR}/${APP_NAME}"

mkdir -p "${INSTALL_DIR}"

# 2. Check if local build exists first
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" 2>/dev/null && pwd || echo "")"
if [ -f "${SCRIPT_DIR}/dist/${APP_NAME}" ]; then
    echo "📦 Installing from local build (dist/${APP_NAME})..."
    cp "${SCRIPT_DIR}/dist/${APP_NAME}" "${DEST}"
elif [ -f "${SCRIPT_DIR}/dist/${TARGET_NAME}" ]; then
    echo "📦 Installing from local build (${TARGET_NAME})..."
    cp "${SCRIPT_DIR}/dist/${TARGET_NAME}" "${DEST}"
elif [ -f "${SCRIPT_DIR}/bin/antigravity" ]; then
    echo "📦 Installing from local bin/antigravity..."
    cp "${SCRIPT_DIR}/bin/antigravity" "${DEST}"
else
    echo "🌐 Downloading ${TARGET_NAME} from GitHub Releases..."
    DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/${TARGET_NAME}"
    curl -fsSL "${DOWNLOAD_URL}" -o "${DEST}" || {
        echo "❌ Download failed. Please build locally with: make release"
        exit 1
    }
fi

chmod +x "${DEST}"

# Create convenient aliases
ln -sf "${DEST}" "${INSTALL_DIR}/ag-cleaner"
ln -sf "${DEST}" "${INSTALL_DIR}/agc"

echo "===================================================="
echo "✅ Installed successfully to: ${DEST}"
echo ""
echo "To run Antigravity Cleaner, type in your terminal:"
echo "    antigravity-cleaner"
echo "    (or short alias: ag-cleaner / agc)"
echo ""
echo "To run health & diagnostics:"
echo "    ag-cleaner doctor"
echo "===================================================="

# Check if ~/.local/bin is in PATH
if [[ ":$PATH:" != *":${INSTALL_DIR}:"* ]]; then
    echo "⚠️  Note: ${INSTALL_DIR} is not in your current PATH."
    echo "   Add it by running: export PATH=\"\$HOME/.local/bin:\$PATH\""
fi
