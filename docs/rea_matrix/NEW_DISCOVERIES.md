# Antigravity Reverse Engineering - New Discoveries & Undocumented Findings

This document records unexpected findings, undocumented CLI flags, secret environment variables, hidden APIs, telemetry endpoints, and architectural quirks discovered during the systematic reverse engineering of all 21 Google Antigravity packages (IDE, CLI, Hub, Extensions, SDK).

---

## Discovery 1: Undocumented Control Environment Variables
- `ANTIGRAVITY_SERVER_URL`: Overrides language server startup and directs the IDE to connect to an existing HTTP/gRPC server endpoint (e.g. `http://127.0.0.1:4242`). When defined, process spawning of `agy` is bypassed.
- `AGY_RELEASE_BASE_URL`: Overrides the default backend auto-updater URL (`https://antigravity-cli-auto-updater-974169037036.us-central1.run.app`). Allows 100% offline air-gapped deployments or local mirror redirection.
- `ANTIGRAVITY_SERVER_READY_TIMEOUT_MS`: Overrides the 30,000 ms startup probe timeout for slow machines or heavy proxies.
- `ANTIGRAVITY_CDE`, `AGY_ALLOWED_DOMAINS_ENV`, `AGY_CDE_SEAMLESS_AUTH_ENV`: Environment injection variables passed when running inside Google Cloud Workstations / CDE environments.

## Discovery 2: Competitor Extension Blacklist
In `out/main.js`, Antigravity IDE enforces a hardcoded ban on 19 rival extensions:
- GitHub Copilot (`github.copilot`, `github.copilot-chat`)
- Codeium (`codeium.codeium`, `codeium.codeium-enterprise`, `codeium.codeium-enterprise-updater`)
- Sourcegraph Cody (`sourcegraph.cody-ai`)
- Tabnine (`tabnine.tabnine-vscode`)
- Amazon Q / CodeWhisperer (`amazonwebservices.amazon-q-vscode`, `amazonwebservices.codewhisperer-for-command-line-companion`)
- Supermaven (`supermaven.supermaven`)
- Continue (`continue.continue`)
- Microsoft Remote SSH / WSL / Containers (`ms-vscode-remote.remote-ssh`, `ms-vscode-remote.remote-wsl`, `ms-vscode-remote.remote-containers`, `ms-vsliveshare.vsliveshare`, `ms-python.vscode-pylance`)

## Discovery 3: Internal Codename "Jetski" & Eligibility Architecture
- The entire agent engine in Antigravity is codenamed **Jetski** internally.
- Client-side gatekeeper uses a 13-state enum:
  `UNKNOWN (0)`, `INELIGIBLE_ACCOUNT (1)`, `UNKNOWN_LOCATION (2)`, `UNSUPPORTED_LOCATION (3)`, `RESTRICTED_NETWORK (4)`, `RESTRICTED_AGE (5)`, `NON_USER_ACCOUNT (6)`, `DASHER_USER (7)`, `BYOID_USER (8)`, `RESTRICTED_DASHER_USER (9)`, `VALIDATION_REQUIRED (10)`, `GOOGLE_TOS_NOT_SUPPORTED_BY_CLIENT (11)`, `UNSUPPORTED_CLIENT (12)`.
- State 3 triggers HTTP 403 / location unavailable errors.
