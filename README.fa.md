<p align="center">
  <a href="README.fa.md">🇮🇷 فارسی</a> •
  <a href="README.md">🇺🇸 English</a> •
  <a href="README.ar.md">🇸🇦 العربية</a> •
  <a href="README.ru.md">🇷🇺 Русский</a> •
  <a href="README.es.md">🇪🇸 Español</a> •
  <a href="README.tr.md">🇹🇷 Türkçe</a> •
  <a href="README.zh.md">🇨🇳 简体中文</a> •
  <a href="README.ur.md">🇵🇰 اردو</a>
</p>

# ⚡ آنتی‌گراویتی کلینر (Antigravity Cleaner) — جعبه‌ابزار آزادی هوش مصنوعی (نسخه ۵.۲.۰)

<div align="center">
  <img src="docs/images/banner.png" alt="Antigravity Cleaner Banner" width="100%">
  <br>
  
  [![Version](https://img.shields.io/badge/نسخه-5.2.0-blue?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner/releases)
  [![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
  [![پشتیبانی نسخه‌ها](https://img.shields.io/badge/پشتیبانی-تمام_نسخه‌های_انتی‌گراویتی-purple?style=for-the-badge)]()
  [![تست‌ها](https://img.shields.io/badge/تست‌ها-۱۰۰٪_موفق-brightgreen?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![بنچمارک](https://img.shields.io/badge/سرعت_اسکن-۱.۹_گیگابایت_بر_ثانیه-blueviolet?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![امنیت](https://img.shields.io/badge/امنیت-۱۰۰٪_آفلاین_|_زیرو--ترک-success?style=for-the-badge)]()
</div>

> *از صمیم قلب تقدیم به تمام مردم، توسعه‌دهندگان و پژوهشگرانی که هم‌زمان میان تیغ فیلترینگ داخلی و سد تحریم‌های ناعادلانه خارجی گرفتار شده‌اند — در 🇮🇷 ایران، 🇨🇳 چین، 🇷🇺 روسیه، 🇹🇷 ترکیه، 🇨🇺 کوبا، 🇸🇾 سوریه، 🇰🇵 کره شمالی، 🇧🇾 بلاروس، 🇸🇩 سودان، 🇻🇪 ونزوئلا، 🇦🇫 افغانستان، 🇲🇲 میانمار و تمام سرزمین‌های محدود شده جهان. دسترسی آزاد به دانش، هوش مصنوعی و فناوری، حق ذاتی تمام بشریت است، نه امتیازی گزینشی.*  
> — **دل‌نوشته‌ای از @dalroot**

---

## ⚡ آنتی‌گراویتی کلینر چیست؟

ابزار **Antigravity Cleaner** یک جعبه‌ابزار سیستمی مدرن و سریع با زبان Go است که به دو صورت گرافیکی (رابط نوستالژیک Retro KeyGen) و خط فرمان ترمینال ارائه می‌شود. این ابزار بدون دستکاری تنظیمات کاربری یا کد پروژه‌ها، نسخه‌های محیط توسعه گوگل آنتی‌گراویتی (نسخه‌های ۱ و ۲، ادیتور، خط فرمان agy و افزونه VS Code) را بررسی و موانع دسترسی را برطرف می‌کند:

- **رفع قطعی خطاهای محدودیت منطقه‌ای و تحریم:** خنثی‌سازی چک‌های اعتبارسنجی لوکال (Region Not Supported)
- **رفع خطای سقف سهمیه و ریت‌لیمیت:** پاکسازی هوشمند کش سهمیه بدون ابطال نشست‌ها (HTTP 429 Quota Exhausted)
- **پایداری استریم پاسخ‌ها:** فعال‌سازی هارت‌بیت فعال (TCP KeepAlive) جهت جلوگیری از قطع ارتباط با فایروال‌های میانی
- **اتصال مستقیم به پروکسی بدون حالت TUN:** هدایت ترافیک ادیتور به پورت‌های لوکال کلاینت‌های فیلترشکن بدون سنگین‌شدن شبکه سیستم

> 🔒 **ضمانت امنیت، حفظ حریم خصوصی و عملکرد کاملاً آفلاین:**  
> آنتی‌گراویتی کلینر به صورت ۱۰۰٪ آفلاین روی رایانه شما اجرا می‌شود. این ابزار فاقد هرگونه تله‌متری، جمع‌آوری داده، ارسال لاگ یا ارتباط با سرورهای خارجی است. کدهای پروژه، درخواست‌ها و تاریخچه گفتگوهای شما دست‌نخورده باقی می‌مانند. سورس‌کد این ابزار آزاد و تحت مجوز بین‌المللی GPL-3.0 منتشر شده است.

---

## 💾 اجرای گرافیکی با یک دابل‌کلیک (بدون نیاز به ترمینال)

اگر ترجیح می‌دهید بدون درگیر شدن با خط فرمان و دستورات ترمینال کار کنید، کافیست فایل نسخه سیستم‌عامل خود را از بخش [Releases گیت‌هاب](https://github.com/tawroot/antigravity-cleaner/releases) دانلود کرده و روی آن **دابل‌کلیک** کنید تا پنجره گرافیکی و خاطره‌انگیز پچر باز شود:

<div align="center">
  <img src="assets/shot.jpg" alt="پنجره گرافیکی رترو آنتی‌گراویتی پچر" width="440">
  <br>
  <sub><em>پچر گرافیکی کلاسیک طرح ویندوز ۹۵ با انیمیشن ستاره‌ای، افکت‌های صوتی نوستالژیک ۸-بیتی و پچ ۱-کلیکه تمام ۴ هدف.</em></sub>
</div>

<br>

<div dir="rtl">

<table style="width: 100%; border-collapse: collapse; text-align: right;" dir="rtl">
  <thead>
    <tr style="background: rgba(255, 255, 255, 0.05); border-bottom: 2px solid rgba(255, 255, 255, 0.1);">
      <th style="padding: 12px 16px; text-align: right; width: 30%;">🖥️ سیستم‌عامل</th>
      <th style="padding: 12px 16px; text-align: center; width: 38%;">⬇️ دریافت فایل باینری</th>
      <th style="padding: 12px 16px; text-align: right; width: 32%;">⚡ نحوه اجرا</th>
    </tr>
  </thead>
  <tbody>
    <tr style="border-bottom: 1px solid rgba(255, 255, 255, 0.08);">
      <td style="padding: 14px 16px;">🪟 <strong>ویندوز</strong><br><small style="opacity: 0.75;">(نسخه‌های ۱۰ و ۱۱ - ۶۴ بیتی)</small></td>
      <td style="padding: 14px 16px; text-align: center;">
        <a href="https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-windows-amd64.exe"><img src="https://img.shields.io/badge/دانلود_مستقیم-.exe-0078D6?style=for-the-badge&logo=windows&logoColor=white" alt="دانلود ویندوز"></a>
        <br>
        <small><a href="https://github.com/tawroot/antigravity-cleaner/releases/latest" style="opacity: 0.85;">📦 مشاهده در صفحه ریلیزها</a></small>
      </td>
      <td style="padding: 14px 16px;">دابل‌کلیک روی فایل اجرایی (باز شدن خودکار محیط گرافیکی)</td>
    </tr>
    <tr style="border-bottom: 1px solid rgba(255, 255, 255, 0.08);">
      <td style="padding: 14px 16px;">🍎 <strong>مک‌او‌اس Apple Silicon</strong><br><small style="opacity: 0.75;">(تراشه‌های M1 تا M4)</small></td>
      <td style="padding: 14px 16px; text-align: center;">
        <a href="https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-arm64"><img src="https://img.shields.io/badge/دانلود_مستقیم-ARM64-white?style=for-the-badge&logo=apple&logoColor=black" alt="دانلود مک آرم"></a>
        <br>
        <small><a href="https://github.com/tawroot/antigravity-cleaner/releases/latest" style="opacity: 0.85;">📦 مشاهده در صفحه ریلیزها</a></small>
      </td>
      <td style="padding: 14px 16px;">دابل‌کلیک یا اجرای مستقیم فایل باینری</td>
    </tr>
    <tr style="border-bottom: 1px solid rgba(255, 255, 255, 0.08);">
      <td style="padding: 14px 16px;">🍎 <strong>مک‌او‌اس Intel</strong><br><small style="opacity: 0.75;">(پردازنده‌های اینتل)</small></td>
      <td style="padding: 14px 16px; text-align: center;">
        <a href="https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-amd64"><img src="https://img.shields.io/badge/دانلود_مستقیم-Intel_x64-gray?style=for-the-badge&logo=apple&logoColor=white" alt="دانلود مک اینتل"></a>
        <br>
        <small><a href="https://github.com/tawroot/antigravity-cleaner/releases/latest" style="opacity: 0.85;">📦 مشاهده در صفحه ریلیزها</a></small>
      </td>
      <td style="padding: 14px 16px;">دابل‌کلیک یا اجرای مستقیم فایل باینری</td>
    </tr>
    <tr style="border-bottom: 1px solid rgba(255, 255, 255, 0.08);">
      <td style="padding: 14px 16px;">🐧 <strong>لینوکس x86_64</strong><br><small style="opacity: 0.75;">(اوبونتو، دبیان، آرچ، فدورا)</small></td>
      <td style="padding: 14px 16px; text-align: center;">
        <a href="https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-amd64"><img src="https://img.shields.io/badge/دانلود_مستقیم-x86__64-FCC624?style=for-the-badge&logo=linux&logoColor=black" alt="دانلود لینوکس"></a>
        <br>
        <small><a href="https://github.com/tawroot/antigravity-cleaner/releases/latest" style="opacity: 0.85;">📦 مشاهده در صفحه ریلیزها</a></small>
      </td>
      <td style="padding: 14px 16px;">اجرای مستقیم یا دستور <code>ag-cleaner gui</code></td>
    </tr>
    <tr>
      <td style="padding: 14px 16px;">🐧 <strong>لینوکس ARM64</strong><br><small style="opacity: 0.75;">(رزبری‌پای و سرورهای ARM)</small></td>
      <td style="padding: 14px 16px; text-align: center;">
        <a href="https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-arm64"><img src="https://img.shields.io/badge/دانلود_مستقیم-ARM64-FCC624?style=for-the-badge&logo=linux&logoColor=black" alt="دانلود لینوکس آرم"></a>
        <br>
        <small><a href="https://github.com/tawroot/antigravity-cleaner/releases/latest" style="opacity: 0.85;">📦 مشاهده در صفحه ریلیزها</a></small>
      </td>
      <td style="padding: 14px 16px;">اجرای مستقیم یا دستور <code>ag-cleaner gui</code></td>
    </tr>
  </tbody>
</table>

</div>

<br>

### 📑 فهرست دسترسی سریع
- [💾 دانلود و اجرای گرافیکی (بدون نیاز به ترمینال)](#-اجرای-گرافیکی-با-یک-دابلکلیک-بدون-نیاز-به-ترمینال)
- [🚀 نصب سریع ترمینالی (یک خطی)](#-نصب-سریع-یک-خطی)
- [💻 داشبورد تعاملی ترمینال (TUI)](#-داشبورد-بصری-و-مدرن-سال-۲۰۲۶-tui)
- [🛠️ دستورات خط فرمان و ابزار دکتر](#️-دستورات-خط-فرمان)
- [🏛️ معماری و نمودار گردش ترافیک](#️-معماری-و-نمودار-گردش-ترافیک)
- [🛡️ تضمین امنیت، حریم خصوصی و عدم دستکاری چت‌ها](#️-تضمین-امنیت-حریم-خصوصی-و-سلامت-دادهها)
- [📊 بنچمارک‌ها و گزارش آزمون زنده](#-بنچمارکها-و-گزارش-آزمونهای-واقعی-سیستم)
- [🥊 جدول مقایسه با سایر روش‌ها](#-جدول-مقایسه-جامع)
- [🔧 جدول رفع اشکال و خطاهای رایج](#-جدول-عیبیابی-سریع-خطاهای-رایج)
- [👥 تقدیر و سازندگان](#-سازندگان-و-توسعهدهندگان)

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
 
<div dir="rtl">

| شاخص عملکردی | مقدار اندازه‌گیری‌شده | سند لاگ رسمی |
| :--- | :--- | :--- |
| **سرعت اسکن بایت‌کد در حافظه RAM** | **۱,۹۰۱ مگابایت بر ثانیه (~۱.۹ گیگابایت/ثانیه)** | [`benchmarks.log`](docs/test-reports/benchmarks.log) |
| **زمان اعمال پچ روی باینری ۱۷۱ مگابایتی** | **۰.۲۳ ثانیه** | [`benchmarks.log`](docs/test-reports/benchmarks.log) |
| **بررسی کامل اتصال شبکه و سرورها (Doctor)** | **۲.۵ ثانیه (هم‌زمان با Goroutines)** | [`live_system_test.log`](docs/test-reports/live_system_test.log) |
| **پاکسازی دقیق کش سهمیه (ارور ۴۲۹)** | **حذف ۱۶ کش معیوب (بدون حذف حتی یک چت)** | [`live_system_test.log`](docs/test-reports/live_system_test.log) |
| **تست‌های واحد سیستمی (Unit Tests)** | **۱۰۰٪ موفق و پاس‌شده** | [`benchmarks.log`](docs/test-reports/benchmarks.log) |

</div>

---

## 🥊 جدول مقایسه جامع با سایر ابزارها

<div dir="rtl">

| شاخص و قابلیت فنی | پچرهای تک‌منظوره پایتونی (`open-patcher` و غیره) | اسکریپت‌های متفرقه (پاورشل و شل) | **⚡ Antigravity Cleaner (نسخه ۵.۲)** |
| :--- | :---: | :---: | :---: |
| **زبان و معماری توسعه** | پایتون (نیازمند نصب پایتون، pip و کتابخانه‌ها) | پاورشل یا بچ‌فایل (محدود به یک سیستم‌عامل) | **Go خالص (تک فایل باینری سبک، بدون نیاز به نصب هیچ پیش‌نیاز)** |
| **طراحی رابط کاربری** | خط فرمان متنی ساده | پنجره کنسول سیاه متفرقه | **محیط گرافیکی رترو (Win95 GUI) + منوی مدرن ترمینال** |
| **اتصال به فیلترشکن بدون حالت TUN** | ❌ (اجبار به درگیر کردن کارت شبکه کل سیستم) | ❌ ندارد | ✅ **شناسایی خودکار پورت‌های پروکسی لوکال و تزریق اختصاصی به ادیتور** |
| **پایداری استریم و رفع قطعی‌های DPI** | ❌ (قطع مکرر ارتباط حین استریم چت) | ❌ ندارد | ✅ **فعال‌سازی KeepAlive روی سوکت‌ها جهت حفظ ارتباط پایدار** |
| **رفع خطای ۴۲۹ بدون پاک شدن تاریخچه** | ❌ (حذف کامل دایرکتوری تنظیمات و چت‌ها) | ❌ (پاکسازی فله‌ای فایل‌ها) | ✅ **پاکسازی هوشمند فقط کش سهمیه + حفظ ۱۰۰٪ چت‌ها و لاگین** |
| **پچ عمیق هسته باینری (`language_server`)** | پشتیبانی محدود | ❌ ندارد | ✅ **پچ مستقیم ماشین‌کد به سبک مهندسی معکوس IDA** |
| **پچ فایل‌های کلاینت و فرانت‌اند (`main.js`)** | ✅ فقط پچ رشته‌ای Regex | ❌ ندارد | ✅ **پچ کامل فرانت‌اند + تخلیه کش‌های معیوب گرافیکی** |
| **جلوگیری از بازدانلود خودکار ادیتور** | ✅ نسخه اولیه | ❌ ندارد | ✅ **قفل کانال به‌روزرسانی و جلوگیری از لغو پچ‌ها** |
| **ابزار بررسی سلامت سیستم (`doctor`)** | ❌ ندارد | ⚠️ محدود | ✅ **تست هم‌زمان تاخیر پینگ، سلامت DNS و ارتباط با مدل Gemini** |
| **جایگزینی اتمیک فایل‌ها روی دیسک** | فایل‌های موقت پایتون | ❌ ندارد | ✅ **تعویض اتمیک Inode بدون کرش یا مسدود شدن برنامه حین کار** |

</div>

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

این ابزار برای حل دقیق و فوری پرتکرارترین ارورهای گزارش‌شده توسعه‌دهندگان طراحی شده است:

<div dir="rtl">

| متن دقیق ارور (Error Message) | ریشه فنی مشکل | راه‌حل فوری و خودکار |
| :--- | :--- | :--- |
| **`Your current account is not eligible for Antigravity, because it is not currently available in your location`** | اعتبارسنجی کلاینت‌ساید ریجن در `main.js`، باینری `agy` و اکستنشن VS Code | اجرای دستور `ag-cleaner patch all` (یا کلیک روی **[Auto-Fix]** در رابط کاربری) |
| **`Antigravity Gemini: opaque "User location is not supported" (HTTP 400)`** | قفل منطقه‌ای سرور گوگل در درخواست‌های مستقیم شبکه داخلی | اجرای برنامه با `ag-cleaner launch` (تزریق مستقیم پروکسی بدون فعال‌سازی TUN) |
| **`Unknown: There was a network issue connecting to the server, please try again.`** | قطع ارتباط استریم‌های بلاتکلیف (Idle) توسط فایروال DPI مخابرات حین پردازش مدل | اجرای برنامه با `ag-cleaner launch` (تنظیم سوکت فعال با TCP KeepAlive ۱۵ ثانیه‌ای) |
| **`HTTP 429 Quota Exceeded` / `Too Many Requests`** | آسیب دیدن کش توکن‌ها و شمارنده‌های موقت پس از سشن‌های سنگین کدنویسی | اجرای دستور `ag-cleaner clean` (پاکسازی جراحی کش سهمیه، **حفظ ۱۰۰٪ چت‌ها و لاگین**) |
| **خطای مسدود شدن بررسی‌های `isGoogleInternal` و چک‌های محلی** | بررسی‌های تعبیه‌شده در باینری ماشین‌کد موتور اصلی (`language_server`) | اجرای دستور `ag-cleaner patch all` (پچ چندمرحله‌ای باینری به سبک مهندسی معکوس IDA) |
| **عدم باز شدن برنامه / خطای `SingletonLock`** | بسته شدن ناگهانی سیستم یا سرور گرافیکی و باقی ماندن فایل قفل روی دیسک | اجرای دستور `ag-cleaner clean` یا `ag-cleaner launch` (حذف خودکار قفل‌های معلق) |
| **`Failed to fetch` / کرش کردن یا پاسخ ندادن `language_server`** | فرآیندهای زامبی در پس‌زمینه سیستم | اجرای دستور `ag-cleaner kill` برای توقف پردازش‌های مرده و راه‌اندازی تمیز |
| **افت فریم، کندی محیط و تأخیر در تایپ پاسخ‌ها** | تورم کش دیتای GPU/Dawn و وب‌سوکت‌های قطع‌شده در پس‌زمینه | اجرای دستور `ag-cleaner clean` برای تخلیه ایمن حافظه موقت |

</div>

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
- **دستیار و برنامه‌نویسی زوجی (Pair-Programming):** **هوش مصنوعی Antigravity (Google DeepMind)** — *با طنزی معنادار: ساخت ابزاری با خود هوش مصنوعی گوگل برای رهایی برنامه‌نویسان تحت تحریم از قفل‌های منطقه‌ای گوگل!* 🤖
- **مجوز نرم‌افزار:** GNU General Public License v3.0 (GPL-3.0)
- تقدیم به تمام توسعه‌دهندگانی که برای دسترسی آزاد به دانش و ابزارهای فناوری ایستادگی می‌کنند.

---

## 📈 روند رشد ستاره‌های گیت‌هاب (Star History)

اگر این ابزار به رفع مشکلات و صرفه‌جویی در زمان شما کمک کرد، ثبت یک ستاره (⭐ Star) در گیت‌هاب باعث معرفی آن به سایر برنامه‌نویسان خواهد شد:

<a href="https://star-history.com/#tawroot/antigravity-cleaner&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=tawroot/antigravity-cleaner&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=tawroot/antigravity-cleaner&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=tawroot/antigravity-cleaner&type=Date" />
 </picture>
</a>

---

## 💬 گفت‌وگو و تبادل نظر (Discussions)

برای طرح سوالات، پیشنهاد قابلیت‌های جدید، یا اشتراک‌گذاری تنظیمات پروکسی:  
[**ورود به تالار گفت‌وگوی گیت‌هاب (Discussions) →**](https://github.com/tawroot/antigravity-cleaner/discussions)
