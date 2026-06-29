#!/usr/bin/env bash
set -u

APP_TITLE="ANTIGRAVITY CLEANER"
VERSION="4.1.0-native"
SCRIPT_DIR="$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)"
DATA_PATH="$SCRIPT_DIR/Data"
SESSION_DATA_PATH="$DATA_PATH/Sessions"
OS_NAME="$(uname -s 2>/dev/null || echo unknown)"
HOME_PATH="${HOME:-}"

case "$OS_NAME" in
  Darwin) PLATFORM="mac" ;;
  Linux) PLATFORM="linux" ;;
  *) PLATFORM="unsupported" ;;
esac

mkdir -p "$SESSION_DATA_PATH"

show_header() {
  clear 2>/dev/null || true
  printf '\n  ================================================================\n'
  printf '       %s v%s\n' "$APP_TITLE" "$VERSION"
  printf '       Native macOS/Linux Shell Edition\n'
  printf '  ================================================================\n\n'
}

ok() { printf '  [OK] %s\n' "$*"; }
err() { printf '  [ERROR] %s\n' "$*" >&2; }
info() { printf '  [INFO] %s\n' "$*"; }
warn() { printf '  [WARN] %s\n' "$*"; }
wait_key() { printf '\n  Press Enter to continue... '; read -r _ || true; }

open_url() {
  url="$1"
  if [ "$PLATFORM" = "mac" ]; then
    open "$url" >/dev/null 2>&1 || return 1
  elif command -v xdg-open >/dev/null 2>&1; then
    xdg-open "$url" >/dev/null 2>&1 || return 1
  else
    return 1
  fi
}

open_browser_profile() {
  browser="$1"; profile="$2"; url="$3"
  if [ "$PLATFORM" = "mac" ]; then
    case "$browser" in
      *Chrome*) open -a "Google Chrome" --args --profile-directory="$profile" "$url" ;;
      *Edge*) open -a "Microsoft Edge" --args --profile-directory="$profile" "$url" ;;
      *Brave*) open -a "Brave Browser" --args --profile-directory="$profile" "$url" ;;
      *Opera*) open -a "Opera" "$url" ;;
      *) open_url "$url" ;;
    esac
  else
    case "$browser" in
      *Chrome*) command -v google-chrome >/dev/null 2>&1 && google-chrome --profile-directory="$profile" "$url" >/dev/null 2>&1 & ;;
      *Edge*) command -v microsoft-edge >/dev/null 2>&1 && microsoft-edge --profile-directory="$profile" "$url" >/dev/null 2>&1 & ;;
      *Brave*) command -v brave-browser >/dev/null 2>&1 && brave-browser --profile-directory="$profile" "$url" >/dev/null 2>&1 & ;;
      *Opera*) command -v opera >/dev/null 2>&1 && opera "$url" >/dev/null 2>&1 & ;;
      *) open_url "$url" ;;
    esac
  fi
}

human_size_mb() {
  path="$1"
  if [ ! -e "$path" ]; then echo "0.00"; return; fi
  if command -v du >/dev/null 2>&1; then
    du -sm "$path" 2>/dev/null | awk '{printf "%.2f", $1}'
  else
    echo "0.00"
  fi
}

safe_name() {
  printf '%s' "$1" | tr -c 'A-Za-z0-9@._-' '_'
}

json_email_from_preferences() {
  pref="$1"
  [ -f "$pref" ] || { echo "No Login"; return; }
  if command -v python3 >/dev/null 2>&1; then
    python3 - "$pref" <<'PY' 2>/dev/null || echo "No Login"
import json, sys
try:
    data=json.load(open(sys.argv[1], encoding='utf-8'))
    email=""
    ai=data.get('account_info') or []
    if ai and isinstance(ai, list):
        email=ai[0].get('email') or ""
    if not email:
        email=((data.get('google') or {}).get('services') or {}).get('username') or ""
    print(email or "No Login")
except Exception:
    print("No Login")
PY
  else
    grep -Eo '"email"[[:space:]]*:[[:space:]]*"[^"]+"|"username"[[:space:]]*:[[:space:]]*"[^"]+"' "$pref" 2>/dev/null | head -1 | sed 's/.*: *"//; s/"$//' || echo "No Login"
  fi
}

browser_roots() {
  if [ "$PLATFORM" = "mac" ]; then
    printf 'Google Chrome|%s\n' "$HOME_PATH/Library/Application Support/Google/Chrome"
    printf 'Microsoft Edge|%s\n' "$HOME_PATH/Library/Application Support/Microsoft Edge"
    printf 'Brave Browser|%s\n' "$HOME_PATH/Library/Application Support/BraveSoftware/Brave-Browser"
    printf 'Opera|%s\n' "$HOME_PATH/Library/Application Support/com.operasoftware.Opera"
  else
    printf 'Google Chrome|%s\n' "$HOME_PATH/.config/google-chrome"
    printf 'Brave Browser|%s\n' "$HOME_PATH/.config/BraveSoftware/Brave-Browser"
    printf 'Microsoft Edge|%s\n' "$HOME_PATH/.config/microsoft-edge"
    printf 'Opera|%s\n' "$HOME_PATH/.config/opera"
  fi
}

collect_profiles() {
  profiles=()
  while IFS='|' read -r browser root; do
    [ -d "$root" ] || continue
    for dir in "$root"/Default "$root"/Profile\ *; do
      [ -d "$dir" ] || continue
      name="$(basename "$dir")"
      email="$(json_email_from_preferences "$dir/Preferences")"
      profiles+=("$browser|$name|$dir|$email")
    done
  done < <(browser_roots)
}

show_profiles() {
  collect_profiles
  i=1
  for row in "${profiles[@]}"; do
    IFS='|' read -r browser name path email <<< "$row"
    printf '  [%d] %s - %s (%s)\n' "$i" "$browser" "$email" "$name"
    i=$((i+1))
  done
}

clean_path() {
  path="$1"; desc="$2"; dry_run="$3"
  if [ ! -d "$path" ]; then info "Skip: $desc (Not Found)"; return; fi
  size="$(human_size_mb "$path")"
  count="$(find "$path" -mindepth 1 2>/dev/null | wc -l | tr -d ' ')"
  if [ "$count" = "0" ]; then info "Empty: $desc"; return; fi
  if [ "$dry_run" = "1" ]; then
    info "[DRY RUN] $desc ($path) - Found $count items (${size} MB)"
  else
    find "$path" -mindepth 1 -maxdepth 1 -exec rm -rf -- {} + 2>/dev/null && ok "Cleaned: $desc ($path) - $count items (${size} MB)" || warn "Partial Clean: $desc (locked/protected files skipped)"
  fi
}

invoke_cleaner() {
  dry_run="${1:-0}"
  show_header
  [ "$dry_run" = "1" ] && warn "DRY RUN MODE - NO FILES WILL BE DELETED"
  info "Starting System Cleaner"
  clean_path "/tmp" "Temp Files" "$dry_run"
  if [ "$PLATFORM" = "mac" ]; then
    clean_path "$HOME_PATH/Library/Caches/JetBrains" "JetBrains Cache" "$dry_run"
    clean_path "$HOME_PATH/Library/Application Support/Code/User/workspaceStorage" "VS Code Workspace Storage" "$dry_run"
    clean_path "$HOME_PATH/Library/Preferences/Antigravity" "Antigravity Preferences" "$dry_run"
  else
    clean_path "$HOME_PATH/.cache/JetBrains" "JetBrains Cache" "$dry_run"
    clean_path "$HOME_PATH/.config/Antigravity" "Antigravity Config" "$dry_run"
  fi
  ok "Cleaning Completed."
  wait_key
}

cleaner_menu() {
  show_header
  printf '  System Cleaner\n\n  [1] Quick Clean\n  [2] Dry Run\n  [0] Back\n  > '
  read -r choice || true
  case "$choice" in
    1) invoke_cleaner 0 ;;
    2) invoke_cleaner 1 ;;
  esac
}

copy_dir_contents() {
  src="$1"; dest="$2"
  mkdir -p "$dest"
  cp -a "$src"/. "$dest"/
}

backup_browser_profile() {
  show_header
  info "Backup Browser Profile"
  show_profiles
  if [ "${#profiles[@]}" -eq 0 ]; then warn "No supported browser profiles found."; wait_key; return; fi
  printf '  [0] Back\n  > Select Profile to Backup: '
  read -r sel || true
  [ "$sel" = "0" ] && return
  case "$sel" in ''|*[!0-9]*) warn "Invalid selection"; wait_key; return ;; esac
  [ "$sel" -ge 1 ] && [ "$sel" -le "${#profiles[@]}" ] || { warn "Invalid selection"; wait_key; return; }
  IFS='|' read -r browser name path email <<< "${profiles[$((sel-1))]}"
  printf '\n  [1] Light (Login & Sessions Only)\n  [2] Full (Entire Profile)\n  > Select Mode: '
  read -r mode || true
  timestamp="$(date +%Y-%m-%d_%H-%M-%S)"
  tag="Full"; [ "$mode" = "1" ] && tag="Light"
  dest="$SESSION_DATA_PATH/$browser/$(safe_name "$email")/_${tag}_${timestamp}"
  mkdir -p "$dest"
  info "Backing up $browser ($email)..."
  if [ "$tag" = "Light" ]; then
    for file in Cookies "Login Data" "Web Data" Preferences "Secure Preferences" "Extension Cookies" "Local State"; do
      [ -e "$path/$file" ] && cp -a "$path/$file" "$dest/" 2>/dev/null || true
    done
    for folder in "Local Extension Settings" "Sync Data" "Local Storage" Databases; do
      [ -d "$path/$folder" ] && copy_dir_contents "$path/$folder" "$dest/$folder" 2>/dev/null || true
    done
  else
    copy_dir_contents "$path" "$dest" || { err "Backup failed"; wait_key; return; }
  fi
  cat > "$dest/meta.json" <<EOF
{"Browser":"$browser","ProfileName":"$name","Email":"$email","Date":"$timestamp","Mode":"$tag"}
EOF
  ok "Backup saved to: $dest"
  wait_key
}

backup_antigravity_app() {
  show_header
  ag_path="$HOME_PATH/.config/Antigravity"
  [ "$PLATFORM" = "mac" ] && ag_path="$HOME_PATH/Library/Application Support/Antigravity"
  [ -d "$ag_path" ] || { warn "Antigravity Desktop App data not found."; wait_key; return; }
  timestamp="$(date +%Y-%m-%d_%H-%M-%S)"
  dest="$SESSION_DATA_PATH/Antigravity_Desktop/_$timestamp"
  copy_dir_contents "$ag_path" "$dest" || { err "Backup failed"; wait_key; return; }
  cat > "$dest/meta.json" <<EOF
{"Browser":"Antigravity Desktop","ProfileName":"AppData","Email":"Desktop App","Date":"$timestamp","Mode":"Full"}
EOF
  ok "App backup saved to: $dest"
  wait_key
}

meta_value() {
  key="$1"; file="$2"
  if command -v python3 >/dev/null 2>&1; then
    python3 - "$key" "$file" <<'PY' 2>/dev/null
import json, sys
print(json.load(open(sys.argv[2], encoding='utf-8')).get(sys.argv[1], ''))
PY
  else
    sed -n "s/.*\"$key\":\"\([^\"]*\)\".*/\1/p" "$file" | head -1
  fi
}

restore_profile() {
  show_header
  info "Restore Profile/App"
  backups=()
  while IFS= read -r meta_file; do
    backups+=("$meta_file")
  done < <(find "$SESSION_DATA_PATH" -name meta.json -type f 2>/dev/null | sort)
  if [ "${#backups[@]}" -eq 0 ]; then warn "No backups found."; wait_key; return; fi
  i=1
  for meta in "${backups[@]}"; do
    browser="$(meta_value Browser "$meta")"; email="$(meta_value Email "$meta")"; date="$(meta_value Date "$meta")"; mode="$(meta_value Mode "$meta")"
    printf '  [%d] [%s] %s - %s (%s)\n' "$i" "$mode" "$browser" "$email" "$date"
    i=$((i+1))
  done
  printf '  [0] Back\n  > Select Backup to Restore: '
  read -r sel || true
  [ "$sel" = "0" ] && return
  case "$sel" in ''|*[!0-9]*) warn "Invalid selection"; wait_key; return ;; esac
  [ "$sel" -ge 1 ] && [ "$sel" -le "${#backups[@]}" ] || { warn "Invalid selection"; wait_key; return; }
  meta="${backups[$((sel-1))]}"; src="$(dirname "$meta")"
  browser="$(meta_value Browser "$meta")"; profile="$(meta_value ProfileName "$meta")"
  if [ "$browser" = "Antigravity Desktop" ]; then
    dest="$HOME_PATH/.config/Antigravity"; [ "$PLATFORM" = "mac" ] && dest="$HOME_PATH/Library/Application Support/Antigravity"
  else
    dest=""
    while IFS='|' read -r b root; do [ "$b" = "$browser" ] && dest="$root/$profile"; done < <(browser_roots)
  fi
  [ -n "$dest" ] || { err "Could not resolve destination for $browser"; wait_key; return; }
  warn "Restoring will overwrite data in: $dest"
  printf '  > Are you sure? (Y/N): '
  read -r confirm || true
  [ "$confirm" = "Y" ] || [ "$confirm" = "y" ] || return
  case "$browser" in
    *Chrome*) pkill -x "Google Chrome" 2>/dev/null || pkill chrome 2>/dev/null || true ;;
    *Edge*) pkill -x "Microsoft Edge" 2>/dev/null || pkill microsoft-edge 2>/dev/null || true ;;
    *Brave*) pkill -x "Brave Browser" 2>/dev/null || pkill brave-browser 2>/dev/null || true ;;
    *Opera*) pkill -x Opera 2>/dev/null || pkill opera 2>/dev/null || true ;;
  esac
  copy_dir_contents "$src" "$dest" && ok "Restore Successful!" || err "Restore failed"
  wait_key
}

session_manager() {
  show_header
  printf '  Session Manager\n\n  [1] Backup Browser Profile\n  [2] Backup Antigravity App\n  [3] Restore Profile/App\n  [0] Back\n  > '
  read -r choice || true
  case "$choice" in
    1) backup_browser_profile ;;
    2) backup_antigravity_app ;;
    3) restore_profile ;;
  esac
}

system_analysis() {
  show_header
  info "System Analysis"
  for item in \
    "Google Search|https://www.google.com" \
    "Google Identity|https://accounts.google.com" \
    "Gemini AI API|https://generativelanguage.googleapis.com" \
    "Google AI Studio|https://ai.google.dev" \
    "Google Cloud Platform|https://cloud.google.com" \
    "Colab|https://colab.research.google.com" \
    "VSCode Marketplace|https://marketplace.visualstudio.com" \
    "OpenVSX Registry|https://open-vsx.org" \
    "Github API|https://api.github.com"; do
    name="${item%%|*}"; url="${item#*|}"
    printf '    testing %s... ' "$name"
    if command -v curl >/dev/null 2>&1 && curl -fsSIL --max-time 5 "$url" >/dev/null 2>&1; then echo OK; else echo FAIL; fi
  done
  wait_key
}

network_reset() {
  show_header
  warn "This will flush local DNS caches. You may be asked for sudo."
  printf '  > Continue? (Y/N): '
  read -r confirm || true
  [ "$confirm" = "Y" ] || [ "$confirm" = "y" ] || return
  if [ "$PLATFORM" = "mac" ]; then
    sudo dscacheutil -flushcache 2>/dev/null || true
    sudo killall -HUP mDNSResponder 2>/dev/null || true
  else
    if command -v resolvectl >/dev/null 2>&1; then sudo resolvectl flush-caches
    elif command -v systemd-resolve >/dev/null 2>&1; then sudo systemd-resolve --flush-caches
    else warn "No supported DNS cache command found."; fi
  fi
  ok "Network reset complete. Restart if problems persist."
  wait_key
}

network_tools() {
  show_header
  printf '  Network Optimizer\n\n  [1] Full System Analysis\n  [2] Reset Network / Flush DNS\n  [0] Back\n  > '
  read -r choice || true
  case "$choice" in
    1) system_analysis ;;
    2) network_reset ;;
  esac
}

region_inspector() {
  show_header
  info "Region Inspector"
  printf '  > Launch browser leak pre-check? (Y/N): '
  read -r check || true
  [ "$check" = "Y" ] || [ "$check" = "y" ] && open_url "https://browserleaks.com/ip" || true
  show_profiles
  if [ "${#profiles[@]}" -eq 0 ]; then warn "No supported browser profiles found."; wait_key; return; fi
  printf '  [0] Back\n  > Select Profile to Inspect: '
  read -r sel || true
  [ "$sel" = "0" ] && return
  case "$sel" in ''|*[!0-9]*) warn "Invalid selection"; wait_key; return ;; esac
  [ "$sel" -ge 1 ] && [ "$sel" -le "${#profiles[@]}" ] || { warn "Invalid selection"; wait_key; return; }
  IFS='|' read -r browser name path email <<< "${profiles[$((sel-1))]}"
  url="https://policies.google.com/country-association-form"
  info "Opening Region Settings for: $email"
  open_browser_profile "$browser" "$name" "$url" && ok "Browser opened." || { warn "Could not open automatically. Open manually: $url"; }
  wait_key
}

main() {
  if [ "$PLATFORM" = "unsupported" ]; then err "Unsupported OS: $OS_NAME"; exit 1; fi
  while true; do
    show_header
    printf '  MAIN MENU\n\n'
    printf '  [1] System Cleaner        (Clean Junk & Antigravity Traces)\n'
    printf '  [2] Session Manager       (Backup/Restore Browser Profiles)\n'
    printf '  [3] Network Optimizer     (Fix Connection & DNS)\n'
    printf '  [4] System Analysis       (Check Google & Antigravity Services)\n'
    printf '  [5] Region Inspector      (Check/Change Account Region)\n'
    printf '  [0] Exit\n\n  > Enter Choice: '
    read -r selection || exit 0
    case "$selection" in
      1) cleaner_menu ;;
      2) session_manager ;;
      3) network_tools ;;
      4) system_analysis ;;
      5) region_inspector ;;
      0) echo "Goodbye!"; exit 0 ;;
      *) err "Invalid selection"; wait_key ;;
    esac
  done
}

main "$@"
