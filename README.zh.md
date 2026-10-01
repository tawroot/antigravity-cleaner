# Antigravity Cleaner v2.1.1 🌍
### 跨平台清理和修复工具 | 修复安装和网络问题

[![Python](https://img.shields.io/badge/Made%20with-Python-blue?style=for-the-badge&logo=python)](https://python.org)
[![OS](https://img.shields.io/badge/Platform-Win%20|%20Mac%20|%20Linux-lightgrey?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue?style=for-the-badge)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge)](http://makeapullrequest.com)

<div align="center">
  <img src="https://capsule-render.vercel.app/api?type=waving&color=00BFFF&height=300&section=header&text=Antigravity%20Cleaner&fontSize=90&animation=fadeIn&fontAlignY=38&desc=Fix%20Install%20Errors%20%26%20Network%20Issues&descAlignY=51&descAlign=62" alt="Antigravity Header" />
</div>

---

## 🌐 可用语言 | Available Languages

| 🇬🇧 [English](README.md) | 🇮🇷 [فارسی](README.fa.md) | 🇨🇳 **中文** | 🇷🇺 [Русский](README.ru.md) |
|:---:|:---:|:---:|:---:|
| 🇹🇷 [Türkçe](README.tr.md) | 🇪🇸 [Español](README.es.md) | 🇦🇪 [العربية](README.ar.md) | 🇵🇰 [اردو](README.ur.md) |

---

## 🌍 我们解决的全球性问题

**Antigravity Cleaner** 最初是为伊朗用户开发的，但它解决的问题**在全球范围内普遍存在**。

### 🎯 我们解决的问题：

- **🚫 区域限制**：由于制裁，Google服务在某些国家不可用（伊朗、古巴、叙利亚、朝鲜、克里米亚）
- **🌐 网络干扰**：DNS操纵、代理冲突、防火墙限制
- **🔒 区域锁定错误**：登录时显示"您所在地区不可用"
- **📡 连接问题**：网络错误、超时、SSL证书失败
- **🔄 安装失败**：下载损坏、进程卡住、残留文件
- **🔐 会话过期**：重复登录提示、Cookie损坏

### 👥 谁会使用这个工具：

- **🚨 重度限制网络**：全面互联网过滤的国家（中国、土库曼斯坦、朝鲜、伊朗）
- **⚖️ 受制裁地区**：Google服务被正式封锁的地区（伊朗、古巴、叙利亚、克里米亚）
- **🌐 部分过滤网络**：选择性服务限制的地区（俄罗斯、土耳其、埃及、巴基斯坦）
- **🏢 网络受限环境**：企业网络、教育机构、公共WiFi
- **✈️ 旅行者和外籍人士**：旅行时遇到区域锁定的用户
- **🔧 开发者和IT专业人员**：处理安装问题和网络诊断的系统管理员

---

### 🌟 解决普遍问题的综合方案：

**Antigravity Cleaner** 提供了一套强大的工具，旨在解决受限制或不稳定网络环境中用户面临的最令人沮丧的障碍。

- **🔐 Google 登录修复与优化**：
  - 自动修复由于区域封锁导致的常见 Google 登录错误（403 Forbidden, 400 Bad Request）。
  - 清理阻止 Google 身份验证尝试的浏览器特定痕迹。
  - 专门优化网络设置，以绕过与登录相关的连接中断。

- **🌐 网络堆栈与 DNS 优化**：
  - **刷新与重置 (Flush & Reset)**：完整重置网络堆栈 (Winsock) 并刷新 DNS 缓存以消除干扰。
  - **DNS 健康检查**：识别您的 DNS 提供商是否在干扰服务连接。
  - **代理冲突解决**：检测并帮助解决干扰浏览器连接的隐藏代理设置。

- **💾 本地加密会话管理**：
  - **安全提取**：从浏览器配置文件中安全提取会话 Cookie，以便转移或备份。
  - **零云隐私 (Zero-Cloud)**：所有数据都**本地**存储在您的机器上。**绝不向任何外部服务器发送任何数据。**
  - **AES-256 加密**：会话使用行业标准加密存储，并由您系统唯一的互斥主密钥保护。
  - **配置文件便携性**：轻松跨配置文件或设备转移已验证的会话，无需重新输入密码。

---

## 🖥️ 终端截图

```console
+-------------------------------------------------------------+
|               ANTIGRAVITY CLEANER v2.1.1                    |
|        Running on Windows 10 | Python 3.12.0                |
+-------------------------------------------------------------+
| 选项:                                                        |
|  [1] 快速清理  (标准路径)                                      |
|  [2] 深度清理  (注册表 + Temp + Cache + Extensions)           |
|  [3] 网络重置  (DNS Flush + Winsock Reset)                   |
|  [4] 完全修复  (深度清理 + 网络重置)                            |
|  [5] 切换测试模式  (Dry Run)                                  |
|  [6] 浏览器登录助手  ⭐ 新功能                                  |
|  [7] 会话管理器  (备份/恢复) ⭐ 新功能                          |
|  [0] 退出                                                    |
+-------------------------------------------------------------+
```

---

## 🚀 一行命令安装运行

### 🪟 Windows (PowerShell)
```powershell
iwr -useb https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex
```

### 🍎 macOS / 🐧 Linux (Terminal)
```bash
curl -sL https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.sh | bash
```

---

## ✨ 主要功能

### 🧹 深度清理
- 从AppData、Temp、注册表和**Chrome扩展**中删除痕迹
- 完全清除残留文件
- 删除前自动备份

### 🌐 网络重置
- 自动执行`netsh`和`flushdns`
- 修复"网络错误"或"区域锁定"问题
- 检测和修复DNS问题

### 🔧 进程管理
- 自动终止卡住的进程
- 检测运行中的服务
- 安全清理不丢失数据

### 🌐 浏览器登录助手 (v2.1+)

登录Antigravity有问题？这个功能可以帮助你：

**安全的浏览器清理：**
- 只清除与Antigravity相关的数据
- 其他浏览器数据保持不变
- 删除前自动备份
- 支持Chrome、Edge、Brave和Firefox

**会话管理：**
- 成功登录后备份会话
- 恢复会话避免重复登录
- 加密存储（AES-256）
- 30天有效期

**网络优化：**
- 测试Google服务连接
- 诊断DNS和代理问题
- 修复SSL证书问题
- 详细的诊断报告

---

## 📋 功能完整指南

### 1️⃣ 快速清理（标准路径）
**功能：**
- 扫描标准Antigravity安装文件夹
- 删除AppData和Roaming中的残留文件
- 清除基本配置文件

**何时使用：**
- 正常卸载Antigravity后
- 常规清理
- 想要快速安全地清理时

**安全性：** ✅ 非常安全 - 只删除Antigravity文件

---

### 2️⃣ 深度清理（激进扫描）
**功能：**
- 快速清理的所有内容，加上：
- 扫描Windows注册表中的Antigravity条目
- 删除临时安装文件
- 清理与Antigravity相关的浏览器扩展
- 搜索隐藏的缓存文件

**何时使用：**
- 快速清理没有解决问题时
- 全新重装之前
- 有顽固的安装错误时

**安全性：** ✅ 安全 - 删除前创建备份

---

### 3️⃣ 网络重置
**功能：**
- 清除DNS缓存
- 重置Windows网络堆栈
- 重置TCP/IP设置

**何时使用：**
- 登录时出现"网络错误"
- 连接超时问题
- 更改VPN/代理后

**安全性：** ⚠️ 需要重启 - 但完全安全

---

### 4️⃣ 完全修复
**功能：**
- 运行深度清理（选项2）
- 然后运行网络重置（选项3）
- 完整的系统清理和网络修复

**何时使用：**
- 需要最大程度清理
- 多次安装尝试失败
- 同时存在文件和网络问题

**安全性：** ✅ 安全 - 两个安全操作的组合

---

### 5️⃣ 测试模式（Dry Run）
**功能：**
- 启用"测试模式"
- 显示将被删除的内容而不实际删除
- 允许预览所有操作

**何时使用：**
- 第一次使用工具时
- 想要查看将清理什么
- 实际清理前的测试

**安全性：** ✅ 完全安全 - 没有实际更改

---

### 6️⃣ 浏览器登录助手 ⭐ 新功能

**选项1：清理浏览器中的Antigravity痕迹（安全）**
- 扫描Chrome、Edge、Brave和Firefox
- 只找到与Antigravity相关的Cookie/缓存
- 删除卡住的登录数据
- 删除前自动备份

**选项2：优化网络以便登录**
- 清除DNS缓存
- 测试Google服务连接
- 修复代理冲突

**选项3：网络诊断报告**
- 测试Google服务器连接
- 检查DNS解析
- 检测代理/VPN问题
- 验证SSL证书

**选项4：完全登录修复**
- 结合选项1、2和3
- 完整的浏览器和网络修复

**支持的浏览器：**
- ✅ Google Chrome
- ✅ Microsoft Edge
- ✅ Brave Browser
- ✅ Mozilla Firefox

---

### 7️⃣ 会话管理器 ⭐ 新功能

**选项1：备份当前会话**
- 保存当前登录会话
- 使用AES-256加密Cookie
- 安全存储在`~/.antigravity-cleaner/sessions/`

**选项2：恢复保存的会话**
- 恢复之前保存的会话
- 避免重新输入凭据
- 自动登录

**选项3：列出所有保存的会话**
- 显示所有备份的会话
- 显示会话年龄和状态

**选项4：删除旧会话**
- 删除过期会话（超过30天）
- 释放磁盘空间

**安全性：**
- 🔒 AES-256-GCM加密
- 🔒 仅本地存储（无云端）
- 🔒 30天自动过期
- 🔒 限制性文件权限

---

## 🔗 有用链接

- **检查Google区域**：如果您遇到"您所在地区不可用"问题：
  [Google国家关联表单](https://policies.google.com/country-association-form)

- **项目网站**：[tawroot.github.io/antigravity-cleaner](https://tawroot.github.io/antigravity-cleaner)

- **报告问题**：[GitHub Issues](https://github.com/tawroot/antigravity-cleaner/issues)

---

## 🛡️ 安全功能

### 自动备份
- ✅ 删除前备份
- ✅ 带时间戳的备份文件
- ✅ 需要时轻松回滚
- ✅ 存储在`~/.antigravity-cleaner/backups/`

### 测试模式
- ✅ 预览所有操作
- ✅ 没有实际更改
- ✅ 安全测试
- ✅ 使用`python src/main.py --dry-run`启用

### 选择性删除
- ✅ 仅Antigravity相关数据
- ✅ 其他浏览器数据安全
- ✅ 书签保留
- ✅ 密码保留

### 加密
- 🔒 会话使用AES-256-GCM
- 🔒 行业标准安全
- 🔒 仅本地存储
- 🔒 无云传输

---

## 🛠️ 手动安装（开发者）

```bash
git clone https://github.com/tawroot/antigravity-cleaner.git
cd antigravity-cleaner
pip install -r src/requirements.txt
python src/main.py
```

---

## 🔍 SEO关键词

<details>
<summary>点击查看关键词</summary>

**安装和错误:**
antigravity IDE 修复, 卸载 antigravity, 清洁重装, 安装错误, 下载失败, 损坏安装, 卡住进程, 残留文件, 注册表清理

**网络和连接:**
网络错误 antigravity, 谷歌登录修复, 区域锁定修复, 受制裁国家谷歌, DNS刷新工具, 网络重置, 代理冲突, SSL证书错误

**会话和浏览器:**
浏览器会话管理器, Cookie备份, 会话恢复, 加密会话存储, Chrome配置文件管理器

**区域和访问:**
受限网络工具, 区域解锁, 谷歌账户区域更改, VPN替代品, 审查绕过, 互联网自由工具

**技术:**
Windows清理器, macOS清理器, Linux清理器, Python脚本, 跨平台工具, 自动化工具

</details>

---

<div align="center">
  <b>开源项目</b><br>
  <i>用开源工具赋能用户</i><br><br>
  <a href="https://github.com/tawroot/antigravity-cleaner">⭐ 在GitHub上加星</a>
</div>
