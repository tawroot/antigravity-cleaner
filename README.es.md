# Antigravity Cleaner v2.1.1 🌍
### Herramienta de limpieza y reparación multiplataforma | Solucionar errores de instalación y red

[![Python](https://img.shields.io/badge/Made%20with-Python-blue?style=for-the-badge&logo=python)](https://python.org)
[![OS](https://img.shields.io/badge/Platform-Win%20|%20Mac%20|%20Linux-lightgrey?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue?style=for-the-badge)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge)](http://makeapullrequest.com)

<div align="center">
  <img src="https://capsule-render.vercel.app/api?type=waving&color=00BFFF&height=300&section=header&text=Antigravity%20Cleaner&fontSize=90&animation=fadeIn&fontAlignY=38&desc=Fix%20Install%20Errors%20%26%20Network%20Issues&descAlignY=51&descAlign=62" alt="Antigravity Header" />
</div>

---

## 🌐 Idiomas disponibles | Available Languages

| 🇬🇧 [English](README.md) | 🇮🇷 [فارسی](README.fa.md) | 🇨🇳 [中文](README.zh.md) | 🇷🇺 [Русский](README.ru.md) |
|:---:|:---:|:---:|:---:|
| 🇹🇷 [Türkçe](README.tr.md) | 🇪🇸 **Español** | 🇦🇪 [العربية](README.ar.md) | 🇵🇰 [اردو](README.ur.md) |

---

## 🌍 El problema global que resolvemos

**Antigravity Cleaner** fue desarrollado originalmente para usuarios en Irán, pero los problemas que resuelve son **universales en todo el mundo**.

### 🎯 Qué problemas resolvemos:

- **🚫 Restricciones regionales**: Los servicios de Google no están disponibles en ciertos países debido a sanciones (Irán, Cuba, Siria, Corea del Norte, Crimea)
- **🌐 Interferencia de red**: Manipulación de DNS, conflictos de proxy, restricciones de firewall
- **🔒 Errores de bloqueo regional**: Mensaje "No disponible en tu región" al iniciar sesión
- **📡 Problemas de conectividad**: Errores de red, tiempos de espera, fallos de certificado SSL
- **🔄 Errores de instalación**: Descargas corruptas, procesos atascados, archivos residuales
- **🔐 Expiración de sesión**: Solicitudes repetidas de inicio de sesión, corrupción de cookies

### 👥 Quién usa esta herramienta:

- **🚨 Redes altamente restringidas**: Países con filtrado de internet integral (China, Turkmenistán, Corea del Norte, Irán)
- **⚖️ Regiones sancionadas**: Áreas donde los servicios de Google están oficialmente bloqueados (Irán, Cuba, Siria, Crimea)
- **🌐 Redes parcialmente filtradas**: Regiones con restricciones selectivas de servicios (Rusia, Turquía, Egipto, Pakistán)
- **🏢 Entornos con red limitada**: Redes corporativas, instituciones educativas, WiFi público
- **✈️ Viajeros y expatriados**: Usuarios que experimentan bloqueo regional al viajar
- **🔧 Desarrolladores y profesionales de TI**: Administradores de sistemas que solucionan problemas de instalación y red

---

## 🖥️ Captura de pantalla del terminal

```console
+-------------------------------------------------------------+
|               ANTIGRAVITY CLEANER v2.1.1                    |
|        Running on Windows 10 | Python 3.12.0                |
+-------------------------------------------------------------+
| Opciones:                                                   |
|  [1] Limpieza rápida  (Rutas estándar)                      |
|  [2] Limpieza profunda  (Registro + Temp + Cache + Ext)     |
|  [3] Reinicio de red  (DNS Flush + Winsock Reset)           |
|  [4] Reparación completa  (Limpieza profunda + Reinicio red)|
|  [5] Modo de prueba  (Dry Run)                              |
|  [6] Ayudante de inicio de sesión  ⭐ NUEVO                 |
|  [7] Gestor de sesiones  (Respaldo/Restaurar) ⭐ NUEVO      |
|  [0] Salir                                                  |
+-------------------------------------------------------------+
```

---

## 🚀 Instalación y ejecución con un solo comando

### 🪟 Windows (PowerShell)
```powershell
iwr -useb https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex
```

### 🍎 macOS / 🐧 Linux (Terminal)
```bash
curl -sL https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.sh | bash
```

---

## ✨ Características principales

### 🧹 Limpieza profunda
- Elimina rastros de AppData, Temp, Registro y **extensiones de Chrome**
- Limpieza completa de archivos residuales
- Respaldo automático antes de eliminar

### 🌐 Reinicio de red
- Ejecución automática de `netsh` y `flushdns`
- Corrección de errores "Error de red" o "Bloqueo regional"
- Detección y corrección de problemas de DNS

### 🔧 Gestión de procesos
- Terminación automática de procesos atascados
- Detección de servicios en ejecución
- Limpieza segura sin pérdida de datos

### 🌐 Ayudante de inicio de sesión del navegador (v2.1+)

¿Tienes problemas para iniciar sesión en Antigravity? Esta función ayuda:

**Limpieza segura del navegador:**
- Solo se eliminan los datos relacionados con Antigravity
- Los demás datos del navegador permanecen intactos
- Respaldo automático antes de eliminar
- Soporte para Chrome, Edge, Brave y Firefox

**Gestión de sesiones:**
- Respaldo de sesión después de un inicio de sesión exitoso
- Restauración de sesión para evitar inicios de sesión repetidos
- Almacenamiento cifrado (AES-256)
- Validez de 30 días

**Optimización de red:**
- Prueba de conexión a servicios de Google
- Diagnóstico de problemas de DNS y proxy
- Corrección de problemas de certificado SSL
- Informes de diagnóstico detallados

---

## 📋 Guía completa de funciones

### 1️⃣ Limpieza rápida (Rutas estándar)
**Qué hace:**
- Escanea las carpetas de instalación estándar de Antigravity
- Elimina archivos residuales de AppData y Roaming
- Limpia archivos de configuración básicos

**Cuándo usar:**
- Después de desinstalar Antigravity normalmente
- Para limpieza rutinaria
- Cuando quieres limpiar rápida y seguramente

**Seguridad:** ✅ Muy seguro - solo archivos de Antigravity

---

### 2️⃣ Limpieza profunda (Escaneo agresivo)
**Qué hace:**
- Todo lo de Limpieza rápida, más:
- Escaneo del Registro de Windows para entradas de Antigravity
- Eliminación de archivos temporales de instalación
- Limpieza de extensiones del navegador relacionadas con Antigravity
- Búsqueda de archivos de caché ocultos

**Cuándo usar:**
- Cuando la Limpieza rápida no resolvió el problema
- Antes de una reinstalación limpia
- Cuando tienes errores de instalación persistentes

**Seguridad:** ✅ Seguro - crea respaldo antes de eliminar

---

### 3️⃣ Reinicio de red
**Qué hace:**
- Limpia la caché de DNS
- Reinicia la pila de red de Windows
- Reinicia la configuración TCP/IP

**Cuándo usar:**
- "Error de red" al iniciar sesión
- Problemas de tiempo de espera de conexión
- Después de cambiar VPN/proxy

**Seguridad:** ⚠️ Requiere reinicio - pero completamente seguro

---

### 4️⃣ Reparación completa
**Qué hace:**
- Ejecuta Limpieza profunda (opción 2)
- Luego ejecuta Reinicio de red (opción 3)
- Limpieza completa del sistema y corrección de red

**Cuándo usar:**
- Necesitas limpieza máxima
- Múltiples intentos de instalación fallidos
- Problemas tanto de archivos como de red

**Seguridad:** ✅ Seguro - combinación de dos operaciones seguras

---

### 5️⃣ Modo de prueba (Dry Run)
**Qué hace:**
- Activa el "modo de prueba"
- Muestra qué se eliminará sin eliminar realmente
- Permite previsualizar todas las operaciones

**Cuándo usar:**
- Primera vez que usas la herramienta
- Quieres ver qué se limpiará
- Prueba antes de la limpieza real

**Seguridad:** ✅ Completamente seguro - sin cambios reales

---

### 6️⃣ Ayudante de inicio de sesión del navegador ⭐ NUEVO

**Opción 1: Limpiar rastros de Antigravity en el navegador (Seguro)**
- Escanea Chrome, Edge, Brave y Firefox
- Encuentra solo cookies/caché relacionados con Antigravity
- Elimina datos de inicio de sesión atascados
- Respaldo automático antes de eliminar

**Opción 2: Optimizar red para inicio de sesión**
- Limpia la caché de DNS
- Prueba la conexión a servicios de Google
- Corrige conflictos de proxy

**Opción 3: Informe de diagnóstico de red**
- Prueba la conexión a servidores de Google
- Verifica la resolución de DNS
- Detecta problemas de proxy/VPN
- Verifica certificados SSL

**Opción 4: Reparación completa de inicio de sesión**
- Combina las opciones 1, 2 y 3
- Corrección completa de navegador y red

**Navegadores soportados:**
- ✅ Google Chrome
- ✅ Microsoft Edge
- ✅ Brave Browser
- ✅ Mozilla Firefox

---

### 7️⃣ Gestor de sesiones ⭐ NUEVO

**Opción 1: Respaldar sesión actual**
- Guarda tu sesión de inicio de sesión actual
- Cifra las cookies con AES-256
- Almacena de forma segura en `~/.antigravity-cleaner/sessions/`

**Opción 2: Restaurar sesión guardada**
- Restaura la sesión guardada previamente
- Evita volver a ingresar credenciales
- Inicio de sesión automático

**Opción 3: Listar todas las sesiones guardadas**
- Muestra todas las sesiones respaldadas
- Muestra la antigüedad y el estado de la sesión

**Opción 4: Eliminar sesiones antiguas**
- Elimina sesiones expiradas (más de 30 días)
- Libera espacio en disco

**Seguridad:**
- 🔒 Cifrado AES-256-GCM
- 🔒 Solo almacenamiento local (sin nube)
- 🔒 Expiración automática de 30 días
- 🔒 Permisos de archivo restrictivos

---

## 🔗 Enlaces útiles

- **Verificar región de Google**: Si tienes el problema "No disponible en tu región":
  [Formulario de asociación de país de Google](https://policies.google.com/country-association-form)

- **Sitio web del proyecto**: [tawroot.github.io/antigravity-cleaner](https://tawroot.github.io/antigravity-cleaner)

- **Reportar problemas**: [GitHub Issues](https://github.com/tawroot/antigravity-cleaner/issues)

---

## 🛡️ Características de seguridad

### Respaldo automático
- ✅ Respaldo antes de cualquier eliminación
- ✅ Archivos de respaldo con marca de tiempo
- ✅ Reversión fácil si es necesario
- ✅ Almacenamiento en `~/.antigravity-cleaner/backups/`

### Modo de prueba
- ✅ Previsualización de todas las operaciones
- ✅ Sin cambios reales
- ✅ Prueba segura
- ✅ Activar con `python src/main.py --dry-run`

### Eliminación selectiva
- ✅ Solo datos relacionados con Antigravity
- ✅ Otros datos del navegador seguros
- ✅ Marcadores preservados
- ✅ Contraseñas preservadas

### Cifrado
- 🔒 AES-256-GCM para sesiones
- 🔒 Seguridad de estándar industrial
- 🔒 Solo almacenamiento local
- 🔒 Sin transmisión a la nube

---

## 🛠️ Instalación manual (Desarrolladores)

```bash
git clone https://github.com/tawroot/antigravity-cleaner.git
cd antigravity-cleaner
pip install -r src/requirements.txt
python src/main.py
```

---

## 🔍 Palabras clave SEO

<details>
<summary>Haz clic para ver las palabras clave</summary>

**Instalación y errores:**
corrección antigravity IDE, desinstalar antigravity, reinstalación limpia, error de instalación, descarga fallida, instalación corrupta, proceso atascado, archivos residuales, limpiador de registro

**Red y conectividad:**
error de red antigravity, corrección inicio sesión google, corrección bloqueo regional, países sancionados google, herramienta limpieza DNS, reinicio de red, conflicto proxy, error certificado SSL

**Sesión y navegador:**
administrador sesiones navegador, respaldo cookies, restauración sesión, almacenamiento sesión cifrado, administrador perfiles chrome

**Regional y acceso:**
herramienta redes restringidas, desbloqueo regional, cambio región cuenta google, alternativa VPN, elusión censura, herramienta libertad internet

**Técnico:**
limpiador windows, limpiador macos, limpiador linux, script python, herramienta multiplataforma, herramienta automatización

</details>

---

<div align="center">
  <b>Proyecto de código abierto</b><br>
  <i>Empoderando a los usuarios con herramientas de código abierto</i><br><br>
  <a href="https://github.com/tawroot/antigravity-cleaner">⭐ Estrella en GitHub</a>
</div>
