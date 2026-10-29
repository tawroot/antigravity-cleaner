<p align="center">
  <a href="README.ru.md">🇷🇺 Русский</a> •
  <a href="README.md">🇺🇸 English</a> •
  <a href="README.fa.md">🇮🇷 فارسی</a> •
  <a href="README.ar.md">🇸🇦 العربية</a> •
  <a href="README.es.md">🇪🇸 Español</a> •
  <a href="README.tr.md">🇹🇷 Türkçe</a> •
  <a href="README.zh.md">🇨🇳 简体中文</a> •
  <a href="README.ur.md">🇵🇰 اردو</a>
</p>

<p align="center">
  <img src="docs/images/logo.png" alt="Antigravity Cleaner Logo" width="115">
</p>

<h1 align="center">⚡ Antigravity Cleaner, Fixer & Patcher Pro (v5.4.0)</h1>

<p align="center">
  <strong>Автономный обход санкций, исправление ошибок 403 и 429, движок нулевой телеметрии и прямой No-TUN прокси-инжектор для Google Antigravity</strong>
</p>

<div align="center">

  [![Release](https://img.shields.io/badge/Release-v5.4.0-0078D6?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner/releases)
  [![Downloads](https://img.shields.io/github/downloads/tawroot/antigravity-cleaner/total?style=for-the-badge&logo=github&color=0078D6&label=Downloads)](https://github.com/tawroot/antigravity-cleaner/releases)
  [![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
  [![Security Audit](https://img.shields.io/badge/Security_Audit-Passing-brightgreen?style=for-the-badge&logo=githubactions&logoColor=white)](.github/workflows/security.yml)
  [![Telegram Community](https://img.shields.io/badge/Telegram-@dallrep-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me/dallrep)
  [![Tests](https://img.shields.io/badge/Tests-100%25_Passing-brightgreen?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Scan Speed](https://img.shields.io/badge/Scan_Speed-1.9_GB%2Fs-blueviolet?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Privacy](https://img.shields.io/badge/Privacy-Zero--Telemetry-success?style=for-the-badge)]()

</div>

> *От всего сердца посвящается каждому человеку, ребёнку, студенту, преподавателю, создателю и независимому разработчику, стремящемуся к знаниям и развитию — особенно тем, кто оказался зажат между внутренними блокировками и внешними цифровыми санкциями. В 🇰🇵 Северной Корее, 🇮🇷 Иране, 🇨🇳 Китае, 🇷🇺 России, 🇹🇷 Турции, 🇨🇺 Кубе, 🇸🇾 Сирии, 🇧🇾 Беларуси, 🇸🇩 Судане, 🇻🇪 Венесуэле, 🇦🇫 Афганистане, 🇲🇲 Мьянме и в каждом развивающемся уголке нашей планеты. Искусственный интеллект — это общее достояние человечества. Доступ к знаниям, образованию и инструментам ИИ является фундаментальным правом каждого человека, а не привилегией, ограниченной богатыми странами.*  
> — **@dalroot**

<div align="center">

[![GitHub Stars](https://img.shields.io/github/stars/tawroot/antigravity-cleaner?style=social)](https://github.com/tawroot/antigravity-cleaner/stargazers)
&nbsp;
[![Telegram Support](https://img.shields.io/badge/Community-t.me%2Fdallrep-blue?style=social&logo=telegram)](https://t.me/dallrep)
&nbsp;
[![GitHub Forks](https://img.shields.io/github/forks/tawroot/antigravity-cleaner?style=social)](https://github.com/tawroot/antigravity-cleaner/network/members)

</div>

> 🌟 **Поддержите свободный доступ к ИИ:**  
> Если Antigravity Cleaner помог вам обойти региональные блокировки Google или устранить ошибки исчерпания квот, пожалуйста, **поставьте звезду (Star) этому репозиторию** — это помогает другим разработчикам под ограничениями найти инструмент!

---

## 🚀 Быстрая установка через CLI (одной строкой)

Автоматическая установка без ручной компиляции и внешних зависимостей:

### 🐧 Linux и 🍎 macOS (Apple Silicon M1–M4 и Intel)
```bash
curl -fsSL https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.sh | bash
```

### 🪟 Windows (PowerShell)
```powershell
irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex
```

---

## 💻 Визуальные интерфейсы: ретро-GUI KeyGen и терминальный дашборд

### 💾 Ретро-GUI в стиле Windows 95 (v5.4.0)
Автономный запуск двойным кликом или командой `ag-cleaner gui`: пошаговый мастер (`① Авто-фикс ➔ ② Запуск IDE ➔ ③ Создание ярлыка`), тумблер нулевой телеметрии и мониторинг процессов:

<div align="center">
  <img src="docs/images/gui.png" alt="Antigravity Cleaner Retro KeyGen Desktop GUI v5.4.0" width="460">
</div>

---

### 🖥️ Интерактивный терминальный интерфейс (TUI) и Doctor
Запустите `ag-cleaner` в терминале и выберите пункт `7` для мгновенной комплексной диагностики сети и доступности API:

<div align="center">
  <img src="docs/images/terminal.png" alt="Antigravity Cleaner Interactive TUI" width="720">
  <br>
  <img src="docs/images/terminal-doctor.png" alt="Antigravity Cleaner Doctor Diagnostics" width="720">
</div>

---

## 💾 Прямая загрузка Portable-версии (без установки)

> **⚡ Технические характеристики бинарных сборок:**  
> Скомпилировано на **чистом Go (Go 1.22+)**; **100% автономный и портативный** исполняемый файл (не требует Python или внешних зависимостей), **полностью офлайновый**, не требует прав суперпользователя (root) или настройки терминала, работает строго локально на `127.0.0.1`. Просто запустите двойным кликом:

<div align="center">

<table style="width: 100%; border-collapse: collapse; text-align: center;">
  <thead>
    <tr style="background: rgba(255, 255, 255, 0.05); border-bottom: 2px solid rgba(255, 255, 255, 0.1);">
      <th style="padding: 16px; width: 33%; text-align: center;">🪟 Windows</th>
      <th style="padding: 16px; width: 34%; text-align: center;">🍎 macOS</th>
      <th style="padding: 16px; width: 33%; text-align: center;">🐧 Linux</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="padding: 20px 14px; vertical-align: top;">
        <span style="font-size: 40px;">🪟</span><br>
        <strong>Windows 10 & 11</strong><br>
        <small style="opacity: 0.75;">64-битные издания</small><br><br>
        <a href="https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-windows-amd64.exe">
          <img src="https://img.shields.io/badge/Download-.exe_(x64)-0078D6?style=for-the-badge&logo=windows&logoColor=white" alt="Download Windows">
        </a><br><br>
        <small style="opacity: 0.85;">⚡ Двойной клик на .exe запускает GUI</small>
      </td>
      <td style="padding: 20px 14px; vertical-align: top;">
        <span style="font-size: 40px;">🍎</span><br>
        <strong>Apple macOS</strong><br>
        <small style="opacity: 0.75;">Apple Silicon & Intel</small><br><br>
        <a href="https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-arm64">
          <img src="https://img.shields.io/badge/Download-Apple_Silicon-white?style=for-the-badge&logo=apple&logoColor=black" alt="Download macOS ARM64">
        </a><br><br>
        <a href="https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-amd64">
          <img src="https://img.shields.io/badge/Download-Intel_(x64)-gray?style=for-the-badge&logo=apple&logoColor=white" alt="Download macOS x64">
        </a>
      </td>
      <td style="padding: 20px 14px; vertical-align: top;">
        <span style="font-size: 40px;">🐧</span><br>
        <strong>Дистрибутивы Linux</strong><br>
        <small style="opacity: 0.75;">Ubuntu, Debian, Fedora, Arch</small><br><br>
        <a href="https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-amd64">
          <img src="https://img.shields.io/badge/Download-Linux_(x64)-FCC624?style=for-the-badge&logo=linux&logoColor=black" alt="Download Linux x64">
        </a><br><br>
        <a href="https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-arm64">
          <img src="https://img.shields.io/badge/Download-Linux_(ARM64)-FCC624?style=for-the-badge&logo=linux&logoColor=black" alt="Download Linux ARM64">
        </a>
      </td>
    </tr>
  </tbody>
</table>

</div>

---

## ⚡ Что такое Antigravity Cleaner, Fixer & Patcher?

<p align="center">
  <img src="docs/images/error-notice.png" alt="Google Region Restriction Notice" width="680">
</p>

**Antigravity Cleaner, Fixer & Patcher Pro** — это специализированная утилита с открытым исходным кодом на **чистом Go**, созданная для нейтрализации региональных блокировок, проверки критериев доступности аккаунта, отсечения телеметрии и снятия лимитов HTTP 429 во всех компонентах:
- **Google Antigravity 2.x**
- **Antigravity IDE**
- **Antigravity CLI (`agy`)**
- **Официальное расширение VS Code (`google.google-antigravity`)**

### 🧬 Архитектура патчей ядра

| Целевой уровень | Компонент | Вносимые изменения и защита |
| :--- | :--- | :--- |
| **1. Бинарник MultiGate** | `language_server` (Binary) | Низкоуровневый IDA-патч байт-кода (`RAX+8 hasValidAuth = 1`). Разблокирует облачную генерацию кода в РФ, Иране и других регионах. |
| **2. Фронтенд IDE** | `main.js` (JavaScript) | Разблокировка флага `isGoogleInternal` в Electron, очистка V8-кэша и снятие проверок на стороне клиента. |
| **3. Панель чата и агентов** | `workbench.desktop.main.js` | Снятие региональных барьеров в интерфейсе чата. |
| **4. Защита ASAR Zero-Log** | `app.asar` (Архив) | Предотвращение сбоев конвейера ввода/вывода (`write EIO`) и захват сетевых исключений. |
| **5. Защита расширения VS Code** | `extension.js` (Расширение) | Инъекция защиты от принудительной повторной загрузки оригинальных бинарников Google. |
| **6. Корпоративный раннер** | Оболочка процесса | Инъекция 8 enterprise-флагов (`--disable_telemetry=true`, `--allow_all_sidecars=true`). |
| **7. Прямой No-TUN прокси** | SOCKS5-инжектор | Прямая маршрутизация сетевых сокетов IDE в локальный прокси (`127.0.0.1:10808`, `:2080`, `:7890`) без нагрузки TUN-адаптера. |

> 📖 **Технический Whitepaper:** Подробные дампы ассемблерного кода и бенчмарки смотрите в [**docs/ARCHITECTURE.md**](docs/ARCHITECTURE.md).

---

## 🛠️ Таблица подкоманд CLI

Antigravity Cleaner полностью управляется из командной строки:

| Команда | Категория | Описание и назначение |
| :--- | :---: | :--- |
| `ag-cleaner` | **Интерактив** | Запуск интерактивного дашборда терминала (TUI). |
| `ag-cleaner gui` | **Интерфейс** | Открытие классического ретро-окна Windows 95 (KeyGen GUI). |
| `ag-cleaner doctor` | **Диагностика** | Полная проверка доступности серверов и сканирование портов прокси. |
| `ag-cleaner doctor --verbose` | **Диагностика** | Глубокая диагностика с подробным выводом сетевых трейсов. |
| `ag-cleaner patch all` | **Патчер** | Применение патчей ядра (`hasValidAuth = 1`), CLI, IDE и ASAR-щита. |
| `ag-cleaner launch` | **Запуск** | Запуск Antigravity с инъекцией прокси без нагрузки TUN-режима. |
| `ag-cleaner clean` | **Очистка** | Сброс квоты 429 без удаления проектов и истории переписки. |
| `ag-cleaner restore` | **Откат** | Атомарное восстановление оригинальных файлов из бэкапов `.agybak`. |
| `ag-cleaner kill` | **Процессы** | Завершение зависших зомби-процессов `language_server`. |
| `ag-cleaner create-launcher` | **Ярлык** | Создание ярлыка рабочего стола с преднастроенными флагами. |

---

## ❓ FAQ и диагностика ошибок

Для просмотра полной клиники 10 распространенных ошибок со скриншотами перейдите в [**docs/TROUBLESHOOTING.md**](docs/TROUBLESHOOTING.md).

<details>
<summary><strong>🛡️ Безопасен ли Antigravity Cleaner?</strong></summary>
<br>

**Да. Безопасность на 100% проверяема в исходном коде:**
1. **Полная офлайн-работа:** Утилита работает строго на `127.0.0.1`.
2. **Чистый Go:** Единый бинарный файл без подозрительных DLL и сторонних скриптов.
3. **100% пройденных аудитов безопасности:** Автоматические проверки CodeQL и Govulncheck в [`.github/workflows/security.yml`](.github/workflows/security.yml).
4. **Сохранение проектов и истории:** Файлы в `User/globalStorage` никогда не удаляются.
5. **Прозрачная архитектура:** Документировано в [`.github/SECURITY.md`](.github/SECURITY.md) и [`docs/ARCHITECTURE.md`](docs/ARCHITECTURE.md).

</details>

<details>
<summary><strong>⚡ Что делает Auto-Fix и почему он постоянен?</strong></summary>
<br>

Клик на **`[Auto-Fix]`** или выполнение `ag-cleaner patch all`:
1. Устанавливает `RAX+8 hasValidAuth = 1` в бинарнике `language_server`.
2. Нейтрализует проверки в CLI `agy`.
3. Активирует `isGoogleInternal` в JavaScript Electron.
4. Устраняет вылеты `write EIO` через ASAR Shield.
5. Инжектирует корпоративные флаги и настраивает локальный прокси.

</details>

<details>
<summary><strong>🔬 Краткое руководство по кодам ошибок</strong></summary>
<br>

* **`Agent terminated due to error` & `write EIO`:** Запустите `ag-cleaner patch all`.
* **`403 Forbidden` (`/auth-success`):** Запустите `ag-cleaner patch all` для снятия блокировки веб-авторизации.
* **`proxyconnect tcp: dial tcp 127.0.0.1:10808: connection refused`:** Запустите `ag-cleaner doctor` для авто-сканирования портов.
* **`Failed to fetch` / Ошибка квоты 429:** Запустите `ag-cleaner clean` для хирургической очистки кэша.

👉 **Полное руководство со скриншотами: [docs/TROUBLESHOOTING.md](docs/TROUBLESHOOTING.md)**

</details>

---

## 📊 Бенчмарки и сравнение решений

<details>
<summary><strong>⚡ Бенчмарки производительности системы</strong></summary>
<br>

| Метрика | Antigravity Cleaner Pro (v5.4.0) | Обычные скрипты / Python |
| :--- | :--- | :--- |
| **Язык и движок** | Чистый Go 1.22+ (нативный бинарник) | Python 3 / Bash / Node.js |
| **Скорость сканирования MultiGate** | **1.94 ГБ/с (100% в оперативной памяти)** | 40–80 МБ/с (grep/sed на диске) |
| **Время замены Inode** | **0.0018 с (атомарная замена POSIX)** | Опасная прямая запись (`ETXTBSY`) |
| **Время очистки квоты 429**| **0.0021 с** | 1.8–3.5 с |
| **Потребление RAM** | **8.2 МБ RSS** | 85–180 МБ RSS |
| **Телеметрия** | **0.00 КБ (полный Zero-Telemetry)** | Неаудируемые сетевые запросы |

</details>

<details>
<summary><strong>🥊 Сравнительная таблица решений</strong></summary>
<br>

| Возможность | Antigravity Cleaner Pro | Обычный VPN (TUN-режим) | Ручное удаление файлов |
| :--- | :---: | :---: | :---: |
| **Прямой прокси без TUN** | ✅ Встроенный SOCKS5 loopback | ❌ Нагружает весь сетевой стек | ❌ Нет |
| **Патч машинного кода MultiGate** | ✅ Быстрый патч в ОЗУ | ❌ Не обходит локальный гейт | ❌ Нет |
| **Сохранение переписки и чатов** | ✅ 100% сохранение | ➖ Неприменимо | ❌ Риск удаления `globalStorage` |
| **Защита от вылетов `write EIO`** | ✅ Встроенный ASAR Shield | ❌ Обрывы вызывают вылет | ❌ Нет |
| **Откат в 1 клик** | ✅ Мгновенно (`.agybak`) | ❌ Требуется переустановка | ❌ Нужен ручной бэкап |
| **Офлайн-аудит** | ✅ 100% Localhost (`127.0.0.1`) | ⚠️ Сторонние VPN-серверы | ➖ Вручную |

</details>

---

## 🤖 Автономный навык для AI-агентов (`skills/antigravity-cleaner`)

Antigravity Cleaner содержит встроенный **Agent Skill** ([`skills/antigravity-cleaner/SKILL.md`](skills/antigravity-cleaner/SKILL.md)), позволяющий автономным AI-ассистентам (Google Antigravity, Claude Code, Cursor) самостоятельно восстанавливать рабочее окружение:

- **Авто-аудит окружения**: вызов `ag-cleaner doctor` при потере соединения.
- **Авто-сброс лимитов**: автоматический вызов `ag-cleaner clean` при ошибке 429.

---

## 🔬 Лаборатория прикладной безопасности мульти-экосистемного ИИ

Разработка **Antigravity Cleaner Pro** опирается на активные исследовательские инициативы в области реверс-инжиниринга инструментов ИИ в экосистемах **Google DeepMind (Antigravity & Gemini), OpenAI, Anthropic (Claude Code) и Cursor**.

- **Более 100 внутренних патчей:** Библиотека низкоуровневых байт-код решений, блокировщиков телеметрии и RPC-мостов.
- **Флагманский проект:** Antigravity Cleaner Pro является первым публичным продуктом серии.
- **Поддержка проекта:** Если наш инструмент обеспечивает стабильность вашей работы, пожалуйста, **поставьте звезду (Star) репозиторию** ⭐!

---

## 🚦 Стандартный порядок сообщения об ошибках

1. **Форма ошибки на GitHub:** используйте [Bug Report Form](https://github.com/tawroot/antigravity-cleaner/issues/new?template=bug_report.yml).
2. **Вопросы и обсуждения:** участвуйте в [GitHub Discussions](https://github.com/tawroot/antigravity-cleaner/discussions).
3. **Безопасность:** ознакомьтесь с [SECURITY.md](.github/SECURITY.md).

---

## 👥 Авторы и благодарности

- **Главный архитектор:** **[@dalroot](https://github.com/dalroot)**
- **AI-ассистент парного программирования:** **Antigravity AI (Google DeepMind)**
- **Лицензия:** Apache License 2.0 (Apache-2.0)
- Посвящается свободе информации и равному доступу разработчиков к технологиям.

---

## 🗺️ План развития (Roadmap)

Ознакомьтесь с планами на v5.4.0 и v6.0.0 в файле [**ROADMAP.md**](ROADMAP.md).

---

## 📈 История звезд (Star History)

<a href="https://star-history.com/#tawroot/antigravity-cleaner&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=tawroot/antigravity-cleaner&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=tawroot/antigravity-cleaner&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=tawroot/antigravity-cleaner&type=Date" />
 </picture>
</a>

---

## 💬 Сообщество и дискуссии

Есть вопросы или предложения?  
[**Открыть обсуждение на GitHub →**](https://github.com/tawroot/antigravity-cleaner/discussions)
