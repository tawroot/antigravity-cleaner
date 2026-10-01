<p align="center">
  <a href="README.md">🇺🇸 English</a> •
  <a href="README.fa.md">🇮🇷 فارسی</a> •
  <a href="README.ar.md">🇸🇦 العربية</a> •
  <a href="README.ru.md">🇷🇺 Русский</a> •
  <a href="README.es.md">🇪🇸 Español</a> •
  <a href="README.tr.md">🇹🇷 Türkçe</a> •
  <a href="README.zh.md">🇨🇳 简体中文</a> •
  <a href="README.ur.md">🇵🇰 اردو</a>
</p>

<p align="center">
  <img src="docs/images/logo.png" alt="Antigravity Cleaner Logo" width="115">
</p>

<h1 align="center">⚡ Antigravity Cleaner — 终极 AI 自由工具包 (v5.2.0)</h1>

<p align="center">
  <strong>专为 Google Antigravity IDE 打造的企业级自愈、区域封锁绕过与诊断引擎</strong>
</p>

<div align="center">

  [![Version](https://img.shields.io/badge/版本-5.2.0-blue?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner/releases)
  [![Go Report](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
  [![Tests](https://img.shields.io/badge/测试-100%25_通过-brightgreen?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Benchmark](https://img.shields.io/badge/字节码扫描-1.9_GB%2Fs-blueviolet?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Platform](https://img.shields.io/badge/平台-Linux%20%7C%20macOS%20%7C%20Windows-brightgreen.svg?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner)
  [![License](https://img.shields.io/badge/许可证-Apache_2.0-blue?style=for-the-badge)](LICENSE)
  [![Security](https://img.shields.io/badge/隐私-100%25_离线_%7C_零追踪-success?style=for-the-badge)]()
</div>

> *谨以此项目献给所有身处网络审查与国际数字制裁双重限制下的开发者、研究者与创作者 — 无论是在 🇨🇳 中国、🇮🇷 伊朗、🇷🇺 俄罗斯、🇧🇾 白俄罗斯、🇹🇷 土耳其、🇨🇺 古巴、🇸🇾 叙利亚、🇰🇵 朝鲜、🇸🇩 苏丹还是世界任何受到限制的角落。自由获取知识、技术与人工智能是每个人的基本权利，而非特权。*  
> — **@dalroot**

---

## ⚡ 什么是 Antigravity Cleaner？

<p align="center">
  <img src="docs/images/error-notice.png" alt="Google 区域限制错误提示" width="680">
</p>

**Antigravity Cleaner** 是一款采用**纯 Go 语言**开发的企业级高性系统工具。它同时提供复古的 **1-Click Retro KeyGen 图形界面** 与符合 2026 年标准的现代终端控制台 (**Lipgloss & Bubbletea**)。它彻底解决了以下组件中的区域封锁、账户资格限制 (403 Location Not Supported)、配额耗尽 (429 Quota Exhausted) 以及代码流中断问题：
- **Google Antigravity 2.x**
- **Antigravity IDE**
- **Antigravity CLI (`agy`)**
- **VS Code 官方扩展 (`google.google-antigravity`)**

与传统的重型脚本（强制要求系统全局启用高权限 **TUN 模式 (VPN)**）不同，Antigravity Cleaner 引入了自动化的 **No-TUN 智能代理注入器**，将 Antigravity 的流量直接精准注入本地代理客户端（如 Clash、Xray、v2rayN、NekoBox、Hiddify），系统开销为零。

> 🔒 **100% 离线与零追踪保证：**
> Antigravity Cleaner 仅在本地环回地址 **`127.0.0.1`** 运行。它包含 **零遥测、零数据统计与零远程网络请求**。您的项目源代码、API 密钥与会话记录绝不会被上传或追踪。完全开源并遵循 Apache License 2.0 许可证。

---

## 💾 一键复古图形界面 1-Click Retro GUI（无需打开终端）

喜欢图形界面操作？直接从 [Releases](https://github.com/tawroot/antigravity-cleaner/releases) 下载适用于您系统的独立可执行文件，**双击**即可运行！无需终端命令、无需安装 Python 环境，没有任何外部依赖。

<div align="center">
  <img src="assets/shot.jpg" alt="Antigravity Retro KeyGen GUI (Win95 风格)" width="440">
  <br>
  <sub><em>纯正的 Windows 95/98 经典破解器风格界面，配有 Demoscene 星空动画、Web Audio 8-bit 音效与一键自动化修复。</em></sub>
</div>

<br>

| 操作系统 | 独立免安装文件 | 运行模式 |
| :--- | :---: | :--- |
| 🪟 **Windows** (64-bit) | [![直接下载](https://img.shields.io/badge/直接下载-.exe-0078D6?style=for-the-badge&logo=windows&logoColor=white)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-windows-amd64.exe) | 双击 `.exe`（自动启动 GUI） |
| 🍎 **macOS** (Apple Silicon M1–M4) | [![直接下载](https://img.shields.io/badge/直接下载-ARM64-white?style=for-the-badge&logo=apple&logoColor=black)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-arm64) | 双击或命令行 `./antigravity-cleaner-darwin-arm64` |
| 🍎 **macOS** (Intel) | [![直接下载](https://img.shields.io/badge/直接下载-Intel_x64-gray?style=for-the-badge&logo=apple&logoColor=white)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-amd64) | 双击或命令行 `./antigravity-cleaner-darwin-amd64` |
| 🐧 **Linux** (x86_64) | [![直接下载](https://img.shields.io/badge/直接下载-x86__64-FCC624?style=for-the-badge&logo=linux&logoColor=black)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-amd64) | 双击或执行 `ag-cleaner gui` |
| 🐧 **Linux** (ARM64) | [![直接下载](https://img.shields.io/badge/直接下载-ARM64-FCC624?style=for-the-badge&logo=linux&logoColor=black)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-arm64) | 双击或执行 `ag-cleaner gui` |

---

## 🚀 终端一键快速安装 (One-Liner)

### Linux 与 macOS
```bash
curl -fsSL https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.sh | bash
```

### Windows (PowerShell 普通用户模式)
```powershell
irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex
```

---

## 💻 交互式终端控制台 (TUI Dashboard)

无参数直接运行即可进入交互式控制面板：

```text
 ┌──────────────────────────────────────────────────────────┐
 │  ANTIGRAVITY CLEANER v5.2.0                              │
 │  AI Freedom & Self-Healing Engine for Restricted Regions │
 └──────────────────────────────────────────────────────────┘
 [1] ⚡ 快速一键自动修复 (Full Patch + Injected Proxy)
 [2] 🧬 修复核心引擎 (language_server 与 agy 机器码)
 [3] 🪟 修复 IDE 界面 (main.js 与 VS Code 扩展保护)
 [4] 🚀 智能代理模式启动 Antigravity (无需 TUN 驱动!)
 [5] 📌 创建带代理参数的桌面专属快捷方式
 [6] 🧹 精准重置 429 配额错误 (保留所有聊天与工作区设置)
 [7] 🩺 运行 Antigravity Doctor 系统与网络体检
 [8] 📦 还原备份文件 (.agybak)
 [9] 🕹️ 启动复古 GUI (经典 Win95 风格界面)
 [0] 🚪 退出程序
```

---

## 🛠️ CLI 命令行自动化子命令

支持在自动化脚本或开发环境构建流水线中调用：

| 子命令 | 功能说明 |
| :--- | :--- |
| `ag-cleaner doctor` | 运行全面的 7 项系统、DNS、代理与 Google Gemini 服务器连通性体检 |
| `ag-cleaner clean` | 精准清理损坏的 429 配额标记与死锁进程文件 (SingletonLock) |
| `ag-cleaner patch` | 自动扫描并为 `language_server` 和 `main.js` 打上解除限制补丁 |
| `ag-cleaner proxy` | 注入本地 SOCKS5 代理环境变量并启动 Antigravity |
| `ag-cleaner gui` | 在本地浏览器窗口中启动独立运行的复古图形界面 |

---

## 🩺 内置医生诊断系统 (Antigravity Doctor)

只需运行 `ag-cleaner doctor` 即可秒级排查连接问题：

```text
🩺 Antigravity Cleaner — 系统健康与网络连接体检报告
───────────────────────────────────────────────────────────
 Antigravity IDE         READY   已安装并正确检测
 Core Engine             READY   已解除限制 (Patched)
 Proxy Connection        READY   已监听 127.0.0.1:7890 (Clash / Mihomo)
 Google DNS              READY   正常 (成功解析 16 个官方 IP)
 Gemini AI API           READY   已连接 (590ms, 极佳)
 Cloud Code Service      READY   已连接 (597ms, 极佳)
 Google Accounts         READY   已连接 (1039ms, 极佳)

 系统已优化完毕，可享受无中断开发体验！
```

---

## 🛡️ 精准清理机制 (Surgical Quota Reset)

不同于常规脚本直接简单粗暴删除用户的整个配置文件夹，Antigravity Cleaner 采用**外科手术级别的精准清理**：
- ❌ **仅清理：** GPU 编译缓存、残留的 `SingletonLock` 单例锁、临时 Cookies 与配额追踪数据。
- ✅ **100% 完整保留：** 所有的历史聊天记录、Google 登录凭据、工作区代码状态、编辑器偏好与安装的插件。
- 💾 在修改任何二进制文件之前，工具会自动创建带有时间戳的 `.agybak` 备份，随时可一键无损回滚。

---

## 👥 开发者与开源许可证

- **核心架构与维护者：** **[@dalroot](https://github.com/dalroot)**
- **结对编程协助：** **Antigravity AI (Google DeepMind)**
- **开源许可证：** **Apache License 2.0 (Apache-2.0)**
- 致力于打破信息技术壁垒，为全球开发者创造自由平等的 AI 编程环境。

