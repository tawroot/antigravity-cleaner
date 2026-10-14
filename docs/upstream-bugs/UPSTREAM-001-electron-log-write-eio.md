# 🐛 UPSTREAM-001: Uncaught Exception `write EIO` in `electron-log` Console Transport

- **Bug Identifier:** `UPSTREAM-001`
- **Affected Component:** `resources/app.asar/node_modules/electron-log/src/node/transports/console.js`
- **Upstream Products:** Google Antigravity IDE (Electron Main Process)
- **Runtime Environment:** Linux (all distributions), macOS, Windows (terminal-detached launches)
- **Toolkit Status:** 🟡 **In Progress (Workaround implemented at launcher OS layer; root cause patch inside ASAR actively planned for v5.4.0 — NOT yet marked as fully resolved)**

---

## 🔍 Upstream Error Stack Trace

```text
Uncaught Exception:
Error: write EIO
    at afterWriteDispatched (node:internal/stream_base_commons:159:15)
    at writeGeneric (node:internal/stream_base_commons:150:3)
    at Socket._writeGeneric (node:net:966:11)
    at Socket._write (node:net:978:8)
    at writeOrBuffer (node:internal/streams/writable:570:12)
    at _write (node:internal/streams/writable:499:10)
    at Writable.write (node:internal/streams/writable:508:10)
    at console.value (node:internal/console/constructor:313:16)
    at console.warn (node:internal/console/constructor:437:26)
    at transport.writeFn (/home/hokar/Applications/Antigravity/resources/app.asar/node_modules/electron-log/src/node/transports/console.js:60:7)
```

---

## 🔬 REA Reverse Engineering Evidence

Using the `analyze_javascript_application` module in the REA reverse engineering framework:

* **REA Evidence Identifier:** `ev_0b0572bf2f393d9d64e50173b4095cb2f4366a26b919886eafb0fa476ec1d9da`
* **Root Artifact SHA-256 (`app.asar`):** `a75535fec1f039be8e78c6a66f3e41f15969101fcf44a0690b7dce4d998b875b`
* **AST Graph Scope:**
  - Parsed JavaScript Files: **682 files**
  - Visited AST Nodes: **2,194,232 nodes**
  - Recovered Modules & Findings: **561 modules**, **3,431 findings**

### Exact Code Extraction (`node_modules/electron-log/src/node/transports/console.js`):
From lines 45–63 in `console.js`:

```javascript
// Line 46: Default level is set to 'silly', broadcasting every message to console
level: 'silly',
transforms: [
  addTemplateColors,
  format,
  formatStyles,
  concatFirstStringElements,
  maxDepth,
  toJSON,
],
useStyles: process.env.FORCE_STYLES,

// Lines 58–61: writeFn lacks try/catch boundary
writeFn({ message }) {
  const consoleLogFn = consoleMethods[message.level] || consoleMethods.info;
  consoleLogFn(...message.data); // <-- Line 60: Uncaught exception on detached stream!
},
```

---

## 🔬 Root Cause Analysis

### 1. The Broken Pipe Defect
In Node.js / Electron, writes to standard output (`process.stdout`) or standard error (`process.stderr`) are asynchronous socket streams. When an application is launched from a terminal or child process, its file descriptors (fd 1 and fd 2) are connected to that terminal.

If the launching terminal is subsequently closed or the process group is orphaned without detaching standard I/O streams:
1. The kernel marks the pipe as closed / detached.
2. The next write attempt to `process.stdout` returns the OS error code `EIO` (Input/output error on broken pipe).
3. The `console.warn` invocation inside `transport.writeFn` in `electron-log` lacks a `try / catch` boundary or a `process.stdout.on('error', ...)` handler.
4. Electron catches this as an unhandled exception in the main process and spawns the modal JavaScript Error dialog.

### 2. Upstream Trigger Events Inside Antigravity
Why does `electron-log` attempt to write to console after launch?
1. **Telemetry Drop Events:** When Zero-Telemetry flags (`--disable_telemetry=true`) are injected, Google's telemetry emitter fails to broadcast beacons to `cloudcode-pa.googleapis.com` and logs a dropped-event warning locally.
2. **Background AutoUpdater Check:** The application attempts periodic background checks against Google update endpoints; under network proxying or air-gapped environments, the timeout triggers a local `log.warn()`.
3. **Window Focus / IPC Heartbeats:** Idle socket disconnects during long periods of developer inactivity.

---

## 🛡️ Mitigation Architecture & Phased Roadmap

### Phase 1: OS-Level Stdio Isolation (`pkg/proxy/launcher.go`) — *Applied*
- `cmd.Stdout = nil`, `cmd.Stderr = nil`, `cmd.Stdin = nil`
- Shell launcher redirects `exec "$APP_PATH" ... >/dev/null 2>&1`
- Prevents broken-pipe popups from parent terminal exits, but leaves the underlying upstream defect unpatched in `app.asar`.

### Phase 2: Root-Cause ASAR Bytecode Patch (Planned v5.4.0) — *In Progress*
- Surgical AST/bytecode modification to set `level: false` or wrap Line 60 with `try { consoleLogFn(...); } catch (_) {}` directly in `app.asar`.
- Eliminates console log dispatch at the root before Node.js attempts stream writes.
