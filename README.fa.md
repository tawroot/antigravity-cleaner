# ⚡ آنتی‌گراویتی کلینر (Antigravity Cleaner) — جعبه‌ابزار آزادی هوش مصنوعی (نسخه ۵.۲.۰)

<div align="center">
  <img src="docs/images/banner.png" alt="Antigravity Cleaner Banner" width="100%">
  <br>
  
  [![Version](https://img.shields.io/badge/نسخه-5.2.0-blue?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner/releases)
  [![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
  [![پشتیبانی نسخه‌ها](https://img.shields.io/badge/پشتیبانی-تمام_نسخه‌های_انتی‌گراویتی-purple?style=for-the-badge)]()
  [![تست‌ها](https://img.shields.io/badge/تست‌ها-۱۰۰٪_موفق-brightgreen?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![بنچمارک](https://img.shields.io/badge/سرعت_اسکن-۱.۹_گیگابایت_بر_ثانیه-blueviolet?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![امنیت](https://img.shields.io/badge/تله‌متری-صفر_(۱۰۰٪_آفلاین)-success?style=for-the-badge)]()
</div>

> *تقدیم به جامعه توسعه‌دهندگان ایران، کوبا، سوریه و تمام نقاط تحت تحریم و سانسور دیجیتال. دسترسی آزاد به ابزارهای هوش مصنوعی و کدنویسی یک حق مسلم انسانی است.*

---

## ⚡ آنتی‌گراویتی کلینر چیست؟

**Antigravity Cleaner** یک ابزار سیستمی مدرن و فوق‌العاده سریع با زبان **Go** است که هم دارای **رابط کاربری گرافیکی کلاسیک (Retro KeyGen GUI)** و هم رابط ترمینال سال ۲۰۲۶ (**Lipgloss**) می‌باشد. این ابزار **به صورت کاملاً خودکار و هوشمند تمام نسخه‌های گوگل آنتی‌گراویتی (Antigravity 1.x / 2.x / محیط IDE / خط فرمان `agy` / اکستنشن VS Code)** را بدون نیاز به تنظیم دستی شناسایی و قفل‌گشایی می‌کند:
- **حل قطعی ارورهای ریجن و تحریم (Region Not Supported)**
- **حل محدودیت کوتا و خطای ۴۲۹ (HTTP 429 Quota Exhausted)**
- **جلوگیری از قطعی استریم چت در حین کار با تنظیم سوکت KeepAlive**
- **تزریق مستقیم پروکسی لوکال بدون نیاز به VPN یا روت کل سیستم (No-TUN)**

> 🔒 **تضمین امنیت، عدم رهگیری و ۱۰۰٪ آفلاین (Zero-Track Guarantee):**
> آنتی‌گراویتی کلینر **کاملاً آفلاین روی لوکال‌هاست (`127.0.0.1`)** اجرا می‌شود. این ابزار **فاقد هرگونه تله‌متری، جمع‌آوری داده، لاگ از راه دور یا اتصال به سرور خارجی** است. کدهای پروژه، پرامپت‌ها و چت‌های شما هرگز خوانده یا ارسال نمی‌شوند. سورس‌کد پروژه کاملاً باز، شفاف و تحت لایسنس بین‌المللی GPL-3.0 قابل بررسی است.

---

## 💾 اجرای گرافیکی با یک دابل‌کلیک (بدون نیاز به ترمینال)

اگر ترجیح می‌دهید بدون درگیر شدن با خط فرمان و دستورات ترمینال کار کنید، کافیست فایل نسخه سیستم‌عامل خود را از بخش [Releases گیت‌هاب](https://github.com/tawroot/antigravity-cleaner/releases) دانلود کرده و روی آن **دابل‌کلیک** کنید تا پنجره گرافیکی و خاطره‌انگیز پچر باز شود:

<div align="center">
  <img src="assets/shot.jpg" alt="پنجره گرافیکی رترو آنتی‌گراویتی پچر" width="440">
  <br>
  <sub><em>پچر گرافیکی کلاسیک طرح ویندوز ۹۵ با انیمیشن ستاره‌ای، افکت‌های صوتی نوستالژیک ۸-بیتی و پچ ۱-کلیکه تمام ۴ هدف.</em></sub>
</div>

<br>

| سیستم‌عامل | لینک دانلود فایل مستقل (بدون نیاز به نصب) | نحوه اجرا |
| :--- | :--- | :--- |
| 🪟 **ویندوز** (۶۴ بیتی) | [**`antigravity-cleaner-windows-amd64.exe`**](https://github.com/tawroot/antigravity-cleaner/releases/latest) | دابل‌کلیک روی فایل `.exe` (اجرای خودکار GUI) |
| 🍎 **مک** (اپل سیلیکون M1 تا M4) | [**`antigravity-cleaner-darwin-arm64`**](https://github.com/tawroot/antigravity-cleaner/releases/latest) | دابل‌کلیک یا اجرای مستقیم فایل |
| 🍎 **مک** (اینتل) | [**`antigravity-cleaner-darwin-amd64`**](https://github.com/tawroot/antigravity-cleaner/releases/latest) | دابل‌کلیک یا اجرای مستقیم فایل |
| 🐧 **لینوکس** (x86_64) | [**`antigravity-cleaner-linux-amd64`**](https://github.com/tawroot/antigravity-cleaner/releases/latest) | دابل‌کلیک یا دستور `ag-cleaner gui` |
| 🐧 **لینوکس** (ARM64) | [**`antigravity-cleaner-linux-arm64`**](https://github.com/tawroot/antigravity-cleaner/releases/latest) | دابل‌کلیک یا دستور `ag-cleaner gui` |

---

## 🏛️ معماری و نمودار گردش ترافیک

```mermaid
graph TD
    subgraph فضای اجرایی آنتی‌گراویتی
        IDE[محیط Antigravity IDE / خط فرمان agy]
        LS[هسته باینری language_server]
    end

    subgraph موتور پاکساز و پچر Go
        PAT[اسکنر حافظه IDA با سرعت ۱.۹ گیگابایت/ثانیه]
        PROXY[شناسایی خودکار پورت‌های پروکسی :10808 / :2080 / :7890]
        KEEPALIVE[تنظیم‌کننده سوکت KeepAlive ۱۵ ثانیه‌ای]
        CLEAN[پاکسازی جراحی توکن‌های فاسد ۴۲۹]
    end

    subgraph کلاینت پروکسی کاربر
        XRAY[Xray / v2rayN / Clash / NekoBox / Hiddify]
    end

    subgraph سرورهای هوش مصنوعی گوگل
        GOOGLE[مدل‌های Gemini AI و سرویس‌های Cloud Code]
    end

    PAT -->|پچ اتمیک Inode بدون بستن برنامه| LS
    LS -->|فعال‌سازی لایسنس و بای‌پس| IDE
    PROXY -->|تزریق متغیرهای محیطی پروکسی| IDE
    IDE -->|هدایت ترافیک به پورت لوکال| XRAY
    KEEPALIVE -.->|جلوگیری از قطع پکت‌های استریم با فایروال| XRAY
    XRAY -->|کانکشن پایدار و بدون قفل| GOOGLE
```

---

## 🛡️ ضمانت امنیت، حریم خصوصی و حفظ اطلاعات

- 🔒 **اجرای ۱۰۰٪ محلی و آفلاین:** این ابزار کاملاً لوکال است، هیچ سرور ریموت یا تحلیلی (Telemetry) ندارد و هیچ دیتایی از سیستمتان خارج نمی‌شود.
- 💬 **حفظ ۱۰۰٪ تاریخچه چت‌ها و پروژه‌ها:** پاکسازی جراحی تنها کش‌های سوخته و کوکی‌های نتورک را پاک می‌کند و به دایرکتوری تاریخچه چت‌ها (`User/globalStorage`) هرگز دست نمی‌زند.
- ⚛️ **پچ اتمیک بدون توقف سیستم:** تعویض باینری به سبک اتمیک لینوکس/ویندوز صورت می‌گیرد، به طوری که حتی اگر آنتی‌گراویتی باز باشد، برنامه کرش نکرده و سشن فعال قطع نمی‌شود.

---

## 📊 گزارش تست‌های زنده و بنچمارک سیستم

| شاخص عملکردی | مقدار اندازه‌گیری‌شده | لینک لاگ رسمی |
| :--- | :--- | :--- |
| **سرعت اسکن بایت‌کد در رم** | **۱,۹۰۱.۸۴ مگابایت بر ثانیه (~۱.۹ گیگابایت/ثانیه)** | [`benchmarks.log`](docs/test-reports/benchmarks.log) |
| **زمان پچ کامل باینری ۱۷۱ مگابایتی** | **۰.۲۳ ثانیه** | [`benchmarks.log`](docs/test-reports/benchmarks.log) |
| **بررسی کامل سلامت شبکه (Doctor)** | **۲.۵ ثانیه (هم‌زمان با Goroutines)** | [`live_system_test.log`](docs/test-reports/live_system_test.log) |
| **پاکسازی جراحی کوتا ۴۲۹** | **۱۶ فایل کش فاسد پاک شد (۰ چت پاک شد)** | [`live_system_test.log`](docs/test-reports/live_system_test.log) |
| **تست‌های واحد پروژه** | **۱۰۰٪ موفق و پاس‌شده** | [`benchmarks.log`](docs/test-reports/benchmarks.log) |

---

## 🥊 جدول مقایسه

| قابلیت | سایر پچرها | اسکریپت‌های متفرقه | **⚡ Antigravity Cleaner (نسخه ۵.۱)** |
| :--- | :---: | :---: | :---: |
| **زبان و معماری** | پایتون (نیازمند `pip` و مفسر) | پاورشل (محدود به ویندوز) | **Go خالص (تک فایل باینری بدون هیچ پیش‌نیاز)** |
| **طراحی بصری ترمینال (TUI)** | لاگ‌های خام و نامنظم | ترمینال قدیمی متنی | **داشبورد کارت‌محور با لبه‌های گرد Lipgloss ۲۰۲۶** |
| **تزریق مستقیم پروکسی (بدون TUN)** | ❌ (اجبار به TUN کل سیستم) | ❌ ندارد | ✅ **کشف خودکار پورت‌های پروکسی و تزریق سوکت** |
| **پایداری استریم (ضد اختلال DPI)** | ❌ (قطع مکرر استریم با DPI) | ❌ ندارد | ✅ **تنظیم سوکت KeepAlive برای رفع قطعی استریم** |
| **ریست جراحی کوتا و ارور ۴۲۹** | ❌ (پاک شدن کل چت‌ها و پروژه‌ها) | ❌ (حذف فله‌ای فایل‌ها) | ✅ **پاکسازی توکن‌های خراب + حفظ ۱۰۰٪ چت‌ها و تنظیمات** |
| **پچ باینری هسته (`language_server`)** | x86_64 / ARM64 | ❌ ندارد | ✅ **پچ ماشین‌کد به سبک IDA با الگوریتم MultiGate** |
| **پچ محیط IDE (`main.js`)** | ✅ Regex پچ | ❌ ندارد | ✅ **پچ خودکار + پاکسازی کش VS Code** |
| **حفاظت از بازدانلود VS Code** | ✅ پایه | ❌ ندارد | ✅ **قفل کانال و گارد کامل بازدانلود** |
| **ابزار عیب‌یاب داکتر (`doctor`)** | ❌ ندارد | ⚠️ محدود | ✅ **تست هم‌زمان ۲ ثانیه‌ای پینگ گوگل، سلامت پروکسی و DNS** |
| **جایگزینی اتمیک (POSIX Atomic)** | پایتون | ❌ ندارد | ✅ **تعویض Inode بدون ایجاد خطا در حین اجرای برنامه** |

---

## 🚀 نصب و اجرای سریع (تک‌خطی)

### 🐧 لینوکس و 🍎 مک‌بوک (Apple Silicon M1/M2/M3/M4 و اینتل)
ترمینال را باز کرده و دستور زیر را وارد کنید:
```bash
curl -fsSL https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.sh | bash
```

### 🪟 ویندوز (PowerShell)
پاورشل را باز کرده و دستور زیر را اجرا کنید:
```powershell
irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex
```

---

## 💻 محیط گرافیکی ۲۰۲۶ در ترمینال

با اجرای دستور `ag-cleaner doctor`، یک کارت وضعیت شیک و آرامش‌بخش بدون شلوغی‌های ترساننده دریافت می‌کنید:

```text
╭────────────────────────────────────────────────────────────────────────────╮
│  🩺 Antigravity Cleaner — Health & Connectivity Audit                      │
│ ────────────────────────────────────────────────────────────────────────   │
│                                                                            │
│  Antigravity IDE         READY   Installed & Detected                      │
│  Core Engine             READY   Unlocked (Patched)                        │
│  Proxy Connection        READY   Active on 127.0.0.1:10808 (v2ray / Xray)  │
│  Google DNS              READY   Healthy (16 IPs resolved)                 │
│  Cloud Code Service      READY   Connected (1100ms, Fast)                  │
│  Google Accounts         READY   Connected (1800ms, Fast)                  │
│  Gemini AI API           READY   Connected (950ms, Fast)                   │
│                                                                            │
╰────────────────────────────────────────────────────────────────────────────╯
                                                                              
   ONLINE  System is optimized. You can launch Antigravity with Smart Proxy!
```

---

## 🛠️ دستورات خط فرمان (CLI)

```bash
# باز کردن منوی تعاملی و تمام‌صفحه
ag-cleaner

# آزمایش سلامت سیستم و پینگ سرورهای گوگل
ag-cleaner doctor

# نمایش گزارش کامل با جزئیات فنی (برای دیباگ)
ag-cleaner doctor --verbose

# اعمال تمام پچ‌ها با یک دستور (هسته، CLI و IDE)
ag-cleaner patch all

# اجرای Antigravity با پروکسی هوشمند (بدون نیاز به TUN)
ag-cleaner launch

# ریست هدفمند ارور ۴۲۹ بدون پاک شدن تاریخچه چت‌ها
ag-cleaner clean

# ساخت لانچر دسکتاپ با پروکسی از پیش تنظیم‌شده
ag-cleaner create-launcher

# مشاهده شناسنامه پروژه و سازندگان
ag-cleaner about
```

---

## 🤖 اسکیل هوش مصنوعی برای ایجنت‌ها (`skills/antigravity-cleaner`)

این پروژه به یک **Agent Skill** رسمی (`skills/antigravity-cleaner/SKILL.md`) مجهز شده است که به دستیارهای هوش مصنوعی (Google Antigravity, Claude Code, Cursor, Codex) این امکان را می‌دهد که به محض مواجهه با ارورهای لایسنس، تحریم، پر شدن کوتا یا قطع شدن استریم، **محیط اجرایی خود را به صورت خودکار تعمیر کنند**:

- **عیب‌یابی خودکار محیط:** ایجنت در صورت قطعی، دستور `ag-cleaner doctor` را فراخوانی می‌کند.
- **رفع خودکار خطای ۴۲۹:** در صورت پر شدن توکن‌ها، ایجنت کش‌های سوخته را بدون دستکاری چت‌های کاربر تمیز می‌کند.
- **تنظیمات صفر:** اسکیل مستقیماً در ریشه پروژه قرار دارد و برای تمام ایجنت‌ها قابل کشف است.

---

## 🔧 جدول راهنمای رفع سریع ارورها (Troubleshooting)

| نوع ارور / مشکل | علت فنی | راه‌حل فوری |
| :--- | :--- | :--- |
| **خطای ریجن یا عدم صلاحیت حساب (Gate 1)** | بررسی بایت‌کد محلی توسط هسته رسمی گوگل | دستور `ag-cleaner patch all` را اجرا کنید |
| **خطای محدودیت سهمیه (HTTP 429)** | کوکی‌های منقضی و انباشت توکن‌های خراب | دستور `ag-cleaner clean` را اجرا کنید (چت‌ها حفظ می‌شوند) |
| **قطع شدن مکرر استریم پاسخ‌ها** | فایروال شبکه و قطع کانکشن‌های بدون KeepAlive | اجرای برنامه با `ag-cleaner launch` (سوکت ۱۵ ثانیه‌ای) |
| **کندی و لگ در یک پروژه خاص (مانند فریلنس)** | تورم کش دیتای نتورک و وب‌سوکت‌های مرده | اجرای دستور `ag-cleaner clean` |

---

## 🏗️ کامپایل از سورس

نیازمند **Go نسخه 1.22 یا بالاتر**:

```bash
git clone https://github.com/tawroot/antigravity-cleaner.git
cd antigravity-cleaner

# بیلد برای سیستم جاری
make build

# بیلد کراس‌پلتفرم برای لینوکس، مک (Intel و Apple Silicon) و ویندوز
make release
```

---

## 👥 سازندگان و حق کپی‌رایت

- **معمار و توسعه‌دهنده اصلی:** **[@dalroot](https://github.com/dalroot)**
- **دستیار و برنامه‌نویسی زوجی (Pair-Programming):** **Antigravity AI (Google DeepMind)**
- **مجوز نرم‌افزار:** GNU General Public License v3.0 (GPL-3.0)
- تقدیم به تمام توسعه‌دهندگانی که برای آزادی دانش و ابزارهای فناوری تلاش می‌کنند.
