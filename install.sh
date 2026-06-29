#!/usr/bin/env sh
set -eu

if ! command -v pwsh >/dev/null 2>&1; then
  echo "PowerShell 7+ (pwsh) is required. Install it first: https://learn.microsoft.com/powershell/scripting/install/installing-powershell" >&2
  exit 1
fi

base_url="https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main"
install_dir="${HOME}/.antigravity"
target_file="${install_dir}/Antigravity.ps1"

mkdir -p "$install_dir"
if command -v curl >/dev/null 2>&1; then
  curl -fsSL "${base_url}/Antigravity.ps1" -o "$target_file"
elif command -v wget >/dev/null 2>&1; then
  wget -qO "$target_file" "${base_url}/Antigravity.ps1"
else
  echo "curl or wget is required." >&2
  exit 1
fi

exec pwsh -NoProfile -File "$target_file"
