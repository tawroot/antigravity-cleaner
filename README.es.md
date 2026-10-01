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

<h1 align="center">⚡ Antigravity Cleaner — Toolkit definitivo para la libertad de IA (v5.2.0)</h1>

<p align="center">
  <strong>Motor empresarial de autorrecuperación, evasión de sanciones y diagnóstico para Google Antigravity IDE</strong>
</p>

<div align="center">

  [![Version](https://img.shields.io/badge/Versión-5.2.0-blue?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner/releases)
  [![Go Report](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
  [![Tests](https://img.shields.io/badge/Pruebas-100%25_Superadas-brightgreen?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Benchmark](https://img.shields.io/badge/Escaneo_Bytecode-1.9_GB%2Fs-blueviolet?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Platform](https://img.shields.io/badge/Plataforma-Linux%20%7C%20macOS%20%7C%20Windows-brightgreen.svg?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner)
  [![License](https://img.shields.io/badge/Licencia-Apache_2.0-blue?style=for-the-badge)](LICENSE)
  [![Security](https://img.shields.io/badge/Privacidad-100%25_Offline_%7C_Cero--Rastreo-success?style=for-the-badge)]()
</div>

> *Dedicado desde el corazón a todas las personas, desarrolladores y creadores atrapados entre el filtrado interno y las sanciones digitales internacionales — en 🇨🇺 Cuba, 🇻🇪 Venezuela, 🇮🇷 Irán, 🇨🇳 China, 🇷🇺 Rusia, 🇸🇾 Siria, 🇰🇵 Corea del Norte, 🇧🇾 Bielorrusia, 🇸🇩 Sudán y en cada rincón restringido del mundo. El acceso libre al conocimiento, la tecnología y la inteligencia artificial es un derecho humano fundamental, no un privilegio.*  
> — **@dalroot**

<div align="center">

[![GitHub Stars](https://img.shields.io/github/stars/tawroot/antigravity-cleaner?style=social)](https://github.com/tawroot/antigravity-cleaner/stargazers)
&nbsp;
**⭐ Si Antigravity Cleaner te ayudó a eludir sanciones o resolver errores de cuota, ¡danos una estrella en GitHub para ayudar a otros desarrolladores a descubrirlo!**
&nbsp;
[![GitHub Forks](https://img.shields.io/github/forks/tawroot/antigravity-cleaner?style=social)](https://github.com/tawroot/antigravity-cleaner/network/members)

</div>

---

## ⚡ ¿Qué es Antigravity Cleaner?

<p align="center">
  <img src="docs/images/error-notice.png" alt="Aviso de restricción regional de Google" width="680">
</p>

**Antigravity Cleaner** es una utilidad de sistemas de alto rendimiento desarrollada en **Go puro**. Ofrece una interfaz gráfica retro **1-Click KeyGen GUI** y un panel interactivo para terminal (**Lipgloss & Bubbletea**). Resuelve de forma definitiva los bloqueos regionales, restricciones geográficas (HTTP 403 Forbidden / Location not supported), agotamiento de cuotas (HTTP 429 Quota Exhausted) y caídas en la transmisión de código para:
- **Google Antigravity 2.x**
- **Antigravity IDE**
- **Antigravity CLI (`agy`)**
- **Extensión oficial para VS Code (`google.google-antigravity`)**

A diferencia de los scripts pesados que obligan a enrutar todo el sistema operativo a través de un modo **TUN (VPN)** con privilegios de root, Antigravity Cleaner integra un **Inyector Smart Proxy Sin TUN** que dirige el tráfico de Antigravity directamente a clientes proxy locales (Xray, Clash, NekoBox, Hiddify) con cero sobrecarga en el sistema.

> 🔒 **Garantía 100% Offline y Cero Rastreo:**  
> Antigravity Cleaner opera **estrictamente en local (`127.0.0.1`)**. Contiene **CERO telemetría, CERO analíticas y CERO llamadas a servidores remotos**. Tu código fuente, claves de API y conversaciones nunca se transmiten ni registran. Código 100% abierto bajo la licencia Apache 2.0.

---

## 💾 GUI Retro KeyGen de 1 Clic (Sin usar terminal)

¿Prefieres una interfaz gráfica sencilla? Descarga el ejecutable para tu sistema operativo desde [Releases](https://github.com/tawroot/antigravity-cleaner/releases) y haz **doble clic**. ¡Sin comandos, sin entorno Python y sin dependencias externas!

<div align="center">
  <img src="assets/shot.jpg" alt="Antigravity Retro KeyGen GUI (Estilo Win95)" width="440">
  <br>
  <sub><em>Patcher auténtico estilo Windows 95/98 con campo estelar interactivo, efectos de sonido Web Audio chiptune de 8 bits y reparación automática en un clic.</em></sub>
</div>

<br>

| Sistema Operativo | Binario Autónomo | Modo de Ejecución |
| :--- | :---: | :--- |
| 🪟 **Windows** (64-bit) | [![Descarga Directa](https://img.shields.io/badge/Descarga_Directa-.exe-0078D6?style=for-the-badge&logo=windows&logoColor=white)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-windows-amd64.exe) | Doble clic en `.exe` (abre la GUI automáticamente) |
| 🍎 **macOS** (Apple Silicon M1–M4) | [![Descarga Directa](https://img.shields.io/badge/Descarga_Directa-ARM64-white?style=for-the-badge&logo=apple&logoColor=black)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-arm64) | Doble clic o `./antigravity-cleaner-darwin-arm64` |
| 🍎 **macOS** (Intel) | [![Descarga Directa](https://img.shields.io/badge/Descarga_Directa-Intel_x64-gray?style=for-the-badge&logo=apple&logoColor=white)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-amd64) | Doble clic o `./antigravity-cleaner-darwin-amd64` |
| 🐧 **Linux** (x86_64) | [![Descarga Directa](https://img.shields.io/badge/Descarga_Directa-x86__64-FCC624?style=for-the-badge&logo=linux&logoColor=black)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-amd64) | Doble clic o `ag-cleaner gui` |
| 🐧 **Linux** (ARM64) | [![Descarga Directa](https://img.shields.io/badge/Descarga_Directa-ARM64-FCC624?style=for-the-badge&logo=linux&logoColor=black)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-arm64) | Doble clic o `ag-cleaner gui` |

---

## 🚀 Instalación rápida CLI (Un solo comando)

### Linux / macOS
```bash
curl -fsSL https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.sh | bash
```

### Windows (PowerShell)
```powershell
irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex
```

---

## 🛠️ Subcomandos CLI

```bash
# Diagnóstico completo de salud del sistema, sockets y variables proxy
ag-cleaner doctor

# Limpieza quirúrgica de cuota y bloqueos (sin borrar código de tus proyectos)
ag-cleaner clean

# Parcheo de bytecode de language_server para anular el bloqueo regional (403)
ag-cleaner patch

# Enrutamiento inteligente a proxy local sin modo TUN (SOCKS5 / HTTP)
ag-cleaner proxy --socks5 127.0.0.1:10808

# Iniciar Antigravity IDE con soporte KeepAlive activo
ag-cleaner run
```

---

## 🥊 Matriz comparativa

| Característica | Scripts Tradicionales | Modo TUN / VPN del Sistema | **Antigravity Cleaner (v5.2.0)** |
| :--- | :---: | :---: | :---: |
| **Arquitectura** | Python / Shell antiguo | Controlador virtual TUN | **Binario Go estático de alto rendimiento** |
| **Evasión de Sanciones (403)** | ❌ No soportado | ⚠️ Inconsistente por fugas DNS | **✅ Parcheo de bytecode in-situ (1.9 GB/s)** |
| **Restablecimiento de Cuota (429)** | ❌ Borra todo el perfil | ❌ No aplicable | **✅ Reinicio quirúrgico de tokens y bloqueos** |
| **Requiere Permisos Root** | ⚠️ Frecuente | ❌ Obligatorio | **✅ Cero privilegios root requeridos** |
| **Sobrecarga de Red** | Alta | Muy alta | **✅ Cero sobrecarga (Loopback directo)** |
| **Privacidad y Telemetría** | ❓ Desconocida | ❓ Varía según el proveedor | **✅ 100% Offline (Zero-Tracking verificado)** |

---

## 📜 Licencia
Este proyecto se distribuye bajo la licencia **Apache License 2.0**.
