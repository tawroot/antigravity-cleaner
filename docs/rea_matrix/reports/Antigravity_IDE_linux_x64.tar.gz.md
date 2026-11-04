# Reverse Engineering Report: Antigravity_IDE_linux_x64.tar.gz

## 1. Archive Metadata & Forensics
- **Filename:** `Antigravity_IDE_linux_x64.tar.gz`
- **File Size:** 187 MB (195,496,575 bytes compressed, 610,662,400 bytes uncompressed)
- **SHA-256 Checksum:** `b060a7c7679103b2b619752f8ee7a3088dd241a7d207d4311974b56ea6f94356`
- **MIME / Format:** `application/gzip`, POSIX tar archive (GNU / Unix)
- **Package Type:** Linux x86_64 Desktop IDE Distribution

---

## 2. Directory Layout & Architecture
The tarball extracts into a top-level root directory named `Antigravity IDE/`:

```
Antigravity IDE/
├── antigravity-ide              # Main Electron wrapper binary (ELF 64-bit x86-64)
├── bin/
│   └── antigravity-ide          # Shell launcher script
├── chrome-sandbox               # Chromium setuid SUID sandbox binary
├── chrome_crashpad_handler      # Google Crashpad crash handler binary
├── icudtl.dat                   # ICU Unicode data table
├── libEGL.so, libGLESv2.so      # OpenGL / ANGLE shared libraries
├── libffmpeg.so                 # FFmpeg media codec shared library
├── libvk_swiftshader.so         # Software Vulkan SwiftShader renderer
├── libvulkan.so.1               # Vulkan loader library
├── locales/                     # 55+ language localization pak files
├── resources/
│   ├── app/                     # Unpacked Electron app directory
│   │   ├── out/                 # Compiled workbench, main process, and jetskiAgent
│   │   ├── extensions/          # Prebundled extensions (agy, code-executor, remote, etc.)
│   │   ├── node_modules/        # Bundled Node.js runtime dependencies
│   │   ├── package.json         # Electron entry definition ("main": "./out/main.js")
│   │   └── product.json         # Product configuration, build commit, telemetry
│   ├── completions/             # Shell completions (bash, zsh)
│   └── resources.pak            # Chromium UI resources
└── v8_context_snapshot.bin      # V8 V8 snapshot blob
```

### Binary Architecture Verification:
- **Binary:** `Antigravity IDE/antigravity-ide`
- **ELF Class:** `ELF64`, 2's complement, little endian
- **Machine:** Advanced Micro Devices X86-64
- **Interpreter:** `/lib64/ld-linux-x86-64.so.2`
- **Type:** `DYN` (Position-Independent Executable - PIE), stripped
- **Build ID:** sha1 `940048b723fe705976921f71928b1ed32a45b6b0`

---

## 3. Product & Engine Identity
From `resources/app/product.json` and `resources/app/package.json`:
- **Product Name:** `Antigravity IDE`
- **Application Name:** `antigravity-ide`
- **Core Engine Version:** `1.107.0` (Code-OSS upstream baseline)
- **IDE Build Version (`ideVersion`):** `2.6.0`
- **Build Commit:** `9e3a5da5d6731effd57d47d88bbe5df14ac6ccdd`
- **Build Date:** `2026-10-26T18:06:23.613Z`
- **Internal Language Server CL:** `1007378935`
- **AI Config Aria Key:** `antigravity-ide`
- **Extension Gallery:** Open-VSX gallery (`https://open-vsx.org/vscode/gallery`)

---

## 4. Subsystems, Bundled Extensions & Core Components

### 4.1 Internal Codename: "Jetski"
The AI agent workbench inside Antigravity is codenamed **Jetski**.
- Located at: `resources/app/out/jetskiAgent/main.js` (11.8 MB compiled React/TS bundle)
- Companion Tailwind stylesheet: `resources/app/out/jetskiAgent/main.css` / `jetskiMain.tailwind.css`

### 4.2 Built-in Extensions:
- `resources/app/extensions/agy`: The core Antigravity extension bridge
  - `dist/extension.js`: Extension wrapper & command registrar
  - `vscode/extension.js`: Core supervisor for downloading, verifying, lifecycle managing, and spawning the `agy` backend binary
  - `bridge.js`: UI / IFrame message bridge communicating with webview agent panels
- `resources/app/extensions/antigravity-code-executor`: Agent sandbox shell execution runner
- `resources/app/extensions/antigravity-remote-openssh`: Drop-in replacement for MS Remote-SSH
- `resources/app/extensions/antigravity-remote-wsl`: Drop-in replacement for MS Remote-WSL
- `resources/app/extensions/antigravity-dev-containers`: Dev containers support

---

## 5. Backend Language Server Launch Chain & Lifecycle

### 5.1 Binary Path Resolution (`Ub()`):
The supervisor in `resources/app/extensions/agy/vscode/extension.js` resolves the backend binary dynamically:
```javascript
function Ub() {
  var a = (0, Ia.homedir)();
  return (0, Ja.join)(a, ".gemini", "bin", `agy${process.platform === "win32" ? ".exe" : ""}`);
}
```
Target path: `~/.gemini/bin/agy` (or `%USERPROFILE%\.gemini\bin\agy.exe`).

### 5.2 Auto-Updater & Download Channels:
If the binary is missing or outdated (< `Na.MIN_AGY_VERSION = "1.1.11"`):
- **Default Production Base URL:** `https://antigravity-cli-auto-updater-974169037036.us-central1.run.app`
- **Dogfood Base URL:** `https://storage.googleapis.com/antigravity-public/antigravity-cli`
- **Environment Override:** `AGY_RELEASE_BASE_URL` (allows 100% offline or mirror redirection)
- **Manifest Format:** Fetches release manifest matching `${platform}-${arch}`
- **Sidecar Cache Metadata:** `~/.gemini/bin/agy.version.json` stores size, mtime, inode, version, and sha512.
- **Concurrency Lock:** `~/.gemini/bin/.agy-install.lock` prevents multiple IDE windows from corrupting downloads.

### 5.3 Spawn Arguments (`spawnAndWaitForPort`):
The IDE spawns the binary with `child_process.spawn`:
```javascript
let d = [
  "--hub",
  `--hub-port=${targetPort}`,
  "--app_data_dir=antigravity",
  `--csrf_token=${this.csrfToken}`,
  ...baseExtraArgs // includes `--add-dir=${workspacePath}` and user-configured `serverArgs`
];
```

### 5.4 Network Probing & Liveness:
- **Port Allocation:** Ephemeral (random free port) by default (`serverPort: 0`). User can configure a fixed port in workspace settings.
- **Startup Ready Probe:** Alternates HTTP GET between `http://127.0.0.1:${port}/healthz` and `http://127.0.0.1:${port}/` every 250ms with 2s socket timeout.
- **Timeout Override:** `ANTIGRAVITY_SERVER_READY_TIMEOUT_MS` (defaults to 30,000 ms).
- **Post-Startup Heartbeat:** Polled every 15,000 ms (`antigravity.heartbeatIntervalMs`). 3 consecutive failures triggers automated process kill and restart.
- **Direct Remote Attach Bypass:** If `ANTIGRAVITY_SERVER_URL` is set, local process spawning is completely bypassed, and the IDE connects to the remote server directly.

---

## 6. Gatekeepers, Geo-Checks & Sanctions

### 6.1 Account & Location Eligibility Enums (`jetskiAgent/main.js`):
The agent runtime defines explicit eligibility states:
```typescript
enum EligibilityState {
  UNKNOWN = 0,
  INELIGIBLE_ACCOUNT = 1,
  UNKNOWN_LOCATION = 2,
  UNSUPPORTED_LOCATION = 3,
  RESTRICTED_NETWORK = 4,
  RESTRICTED_AGE = 5,
  NON_USER_ACCOUNT = 6,
  DASHER_USER = 7,
  BYOID_USER = 8,
  RESTRICTED_DASHER_USER = 9,
  VALIDATION_REQUIRED = 10,
  GOOGLE_TOS_NOT_SUPPORTED_BY_CLIENT = 11,
  UNSUPPORTED_CLIENT = 12
}
```
State `3` (`UNSUPPORTED_LOCATION`) triggers the 403 / 400 client-side error:
`"Your current account is not eligible for Antigravity, because it is not currently available in your location"`.

### 6.2 Competitor Extension Blacklist (`out/main.js`):
Antigravity IDE hardcodes an intentional blocklist preventing rival AI assistants and specific MS extensions from loading:
```javascript
blacklistedExtensionIds: [
  "codeium.codeium",
  "codeium.codeium-enterprise-updater",
  "codeium.codeium-enterprise",
  "sourcegraph.cody-ai",
  "github.copilot",
  "github.copilot-chat",
  "tabnine.tabnine-vscode",
  "amazonwebservices.amazon-q-vscode",
  "supermaven.supermaven",
  "continue.continue",
  "ms-vscode-remote.remote-ssh",
  "ms-vscode-remote.remote-ssh-editor",
  "amazonwebservices.codewhisperer-for-command-line-companion",
  "ms-vscode-remote.remote-wsl",
  "ms-vsliveshare.vsliveshare",
  "ms-python.vscode-pylance",
  "ms-vscode-remote.remote-containers",
  "visualstudioexptteam.vscodeintellicode",
  "jeanp413.open-remote-ssh"
]
```

---

## 7. OS-Specific Behaviors (Linux x64)
- Uses standard Linux standard path: `~/.gemini/bin/agy`
- Requires setuid sandbox binary `chrome-sandbox` with mode `4755` if user namespaces are disabled in kernel.
- Relies on system `ld-linux-x86-64.so.2` and dynamically links standard system libraries (`libpthread`, `librt`, `libdl`).
