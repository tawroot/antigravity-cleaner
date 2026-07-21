#!/usr/bin/env sh
set -eu

base_url="https://raw.githubusercontent.com/atakhadiviom/antigravity-cleaner/fix-macos-linux-support"
install_dir="${HOME}/.antigravity"
target_file="${install_dir}/Antigravity.sh"

mkdir -p "$install_dir"
if command -v curl >/dev/null 2>&1; then
  curl -fsSL "${base_url}/Antigravity.sh" -o "$target_file"
elif command -v wget >/dev/null 2>&1; then
  wget -qO "$target_file" "${base_url}/Antigravity.sh"
else
  echo "curl or wget is required." >&2
  exit 1
fi

chmod +x "$target_file"
if [ -t 0 ]; then
  exec "$target_file"
else
  # Piped install (curl ... | sh): menu must read from the real terminal
  exec "$target_file" 0</dev/tty
fi
