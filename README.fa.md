<div dir="rtl">

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

<p align="center">
  <img src="docs/images/logo.png" alt="Antigravity Cleaner Logo" width="115">
</p>

<h1 align="center">⚡ آنتی‌گراویتی کلینر، فیکسر و پچر پرو (v5.4.3)</h1>

<p align="center">
  <strong>موتور خودترمیم و ضدتحریم هوش مصنوعی، رفع ارور ۴۰۳ و سهمیه ۴۲۹، تله‌متری صفر و تزریق پروکسی بدون سربار TUN</strong>
</p>

<div align="center">

  [![Release](https://img.shields.io/badge/Release-v5.4.3-0078D6?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner/releases)
  [![Downloads](https://img.shields.io/github/downloads/tawroot/antigravity-cleaner/total?style=for-the-badge&logo=github&color=0078D6&label=Downloads)](https://github.com/tawroot/antigravity-cleaner/releases)
  [![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
  [![Security Audit](https://img.shields.io/badge/Security_Audit-Passing-brightgreen?style=for-the-badge&logo=githubactions&logoColor=white)](.github/workflows/security.yml)
  [![Telegram Community](https://img.shields.io/badge/Telegram-@dallrep-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me/dallrep)
  [![Tests](https://img.shields.io/badge/Tests-100%25_Passing-brightgreen?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Scan Speed](https://img.shields.io/badge/Scan_Speed-1.9_GB%2Fs-blueviolet?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Privacy](https://img.shields.io/badge/Privacy-Zero--Telemetry-success?style=for-the-badge)]()

</div>

> *از صمیم قلب تقدیم به تک‌تک انسان‌ها، کودکان، معلمان، دانش‌آموزان، پژوهشگران و توسعه‌دهندگانی که فراتر از مرزها برای آگاهی و یادگیری تلاش می‌کنند — به‌ویژه تمام آنان که هم‌زمان میان تیغ فیلترینگ داخلی و سد تحریم‌های ناعادلانه خارجی گرفتار شده‌اند. در 🇰🇵 کره شمالی، 🇮🇷 ایران، 🇨🇳 چین، 🇷🇺 روسیه، 🇹🇷 ترکیه، 🇨🇺 کوبا، 🇸🇾 سوریه، 🇧🇾 بلاروس، 🇸🇩 سودان، 🇻🇪 ونزوئلا، 🇦🇫 افغانستان، 🇲🇲 میانمار و تمام سرزمین‌های در حال توسعه جهان. فناوری هوش مصنوعی حاصل دسترنج مشترک کل تمدن بشری است. دسترسی آزاد به آموزش، دانش و ابزارهای هوش مصنوعی حق مسلم تمام ابناء بشر است، نه امتیازی گزینشی برای مرزهای ثروتمند.*  
> — **دل‌نوشته‌ای از @dalroot**

<div align="center">

[![GitHub Stars](https://img.shields.io/github/stars/tawroot/antigravity-cleaner?style=social)](https://github.com/tawroot/antigravity-cleaner/stargazers)
&nbsp;
[![Telegram Support](https://img.shields.io/badge/Community-t.me%2Fdallrep-blue?style=social&logo=telegram)](https://t.me/dallrep)
&nbsp;
[![GitHub Forks](https://img.shields.io/github/forks/tawroot/antigravity-cleaner?style=social)](https://github.com/tawroot/antigravity-cleaner/network/members)

</div>

> 🌟 **حمایت از دسترسی آزاد به هوش مصنوعی:**  
> اگر این ابزار به جریان کدنویسی شما کمک کرد، ثبت یک **ستاره (Star)** در گیت‌هاب باعث معرفی آن به سایر برنامه‌نویسان در مناطق تحت تحریم خواهد شد.

---

## 🚀 نصب و اجرای سریع تک‌خطی (CLI One-Liner)

بدون نیاز به نصب دستی، کامپایل یا پیش‌نیازهای نرم‌افزاری، مستقیماً از طریق ترمینال نصب و اجرا کنید:

### 🐧 لینوکس و 🍎 مک‌او‌اس (پردازنده‌های سری M اپل و اینتل)
```bash
curl -fsSL https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.sh | bash
```

### 🪟 ویندوز (PowerShell)
```powershell
irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex
```

---

## 💻 رابط‌های بصری: محیط گرافیکی دسکتاپ و داشبورد ترمینال

### 💾 پنجره گرافیکی پرتابل (Desktop GUI)
اجرای مستقل با دابل‌کلیک یا دستور `ag-cleaner gui` با فرایند سه‌مرحله‌ای شفاف (`۱. پچ خودکار ➔ ۲. اجرای ادیتور ➔ ۳. ساخت میانبر`) همراه با کلید قطع دائمی تله‌متری:

<div align="center">
  <img src="docs/images/gui.png" alt="Antigravity Cleaner Retro KeyGen Desktop GUI v5.4.0" width="460">
</div>

---

### 🖥️ داشبورد تعاملی ترمینال (TUI) و دکتر عیب‌یابی شبکه
با اجرای دستور `ag-cleaner` در ترمینال، منوی تعاملی باز شده و با انتخاب گزینه `7` آزمون کامل سلامت و ارتباط با سرورهای هوش مصنوعی انجام می‌شود:

<div align="center">
  <img src="docs/images/terminal.png" alt="Antigravity Cleaner Interactive TUI" width="720">
  <br>
  <img src="docs/images/terminal-doctor.png" alt="Antigravity Cleaner Doctor Diagnostics" width="720">
</div>

---

## 💾 دریافت مستقیم نسخه پرتابل (بدون نیاز به نصب)

> **⚡ مشخصات فنی باینری‌ها:**  
> فایل اجرایی به زبان **Go خالص (Go 1.22+)** کامپایل شده است؛ **۱۰۰٪ مستقل و پرتابل** (بدون نیاز به پایتون یا هرگونه پیش‌نیاز)، **کاملاً آفلاین**، بدون نیاز به دسترسی روت یا ترمینال، و اجراشونده در محیط لوکال (`127.0.0.1`). کافیست روی فایل دابل‌کلیک کنید:

<div align="center">

<table style="width: 100%; border-collapse: collapse; text-align: center;" dir="rtl">
  <thead>
    <tr style="background: rgba(255, 255, 255, 0.05); border-bottom: 2px solid rgba(255, 255, 255, 0.1);">
      <th style="padding: 16px; width: 33%; text-align: center;">🪟 ویندوز (Windows)</th>
      <th style="padding: 16px; width: 34%; text-align: center;">🍎 مک‌او‌اس (macOS)</th>
      <th style="padding: 16px; width: 33%; text-align: center;">🐧 لینوکس (Linux)</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="padding: 20px 14px; vertical-align: top;">
        <span style="font-size: 40px;">🪟</span><br>
        <strong>ویندوز ۱۰ و ۱۱</strong><br>
        <small style="opacity: 0.75;">نسخه‌های ۶۴ بیتی</small><br><br>
        <a href="https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-windows-amd64.exe">
          <img src="https://img.shields.io/badge/Download-.exe_(x64)-0078D6?style=for-the-badge&logo=windows&logoColor=white" alt="Download Windows">
        </a><br><br>
        <small style="opacity: 0.85;">⚡ دابل‌کلیک روی فایل .exe جهت باز شدن پنجره</small>
      </td>
      <td style="padding: 20px 14px; vertical-align: top;">
        <span style="font-size: 40px;">🍎</span><br>
        <strong>اپل مک‌او‌اس</strong><br>
        <small style="opacity: 0.75;">سری M اپل و پردازنده‌های اینتل</small><br><br>
        <a href="https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-arm64">
          <img src="https://img.shields.io/badge/Apple_Silicon-ARM64-white?style=for-the-badge&logo=apple&logoColor=black" alt="Download Mac ARM64">
        </a><br>
        <a href="https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-amd64" style="display: inline-block; margin-top: 6px;">
          <img src="https://img.shields.io/badge/Intel_Core-x86__64-gray?style=for-the-badge&logo=apple&logoColor=white" alt="Download Mac Intel">
        </a>
      </td>
      <td style="padding: 20px 14px; vertical-align: top;">
        <span style="font-size: 40px;">🐧</span><br>
        <strong>توزیع‌های لینوکس</strong><br>
        <small style="opacity: 0.75;">اوبونتو، دبیان، آرچ، فدورا و ARM</small><br><br>
        <a href="https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-amd64">
          <img src="https://img.shields.io/badge/Linux-x86__64-FCC624?style=for-the-badge&logo=linux&logoColor=black" alt="Download Linux x64">
        </a><br>
        <a href="https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-arm64" style="display: inline-block; margin-top: 6px;">
          <img src="https://img.shields.io/badge/Linux-ARM64-FCC624?style=for-the-badge&logo=linux&logoColor=black" alt="Download Linux ARM64">
        </a>
      </td>
    </tr>
  </tbody>
</table>

</div>

---

## ⚡ آنتی‌گراویتی کلینر، فیکسر و پچر دقیقاً چه کاری انجام می‌دهد؟

<p align="center">
  <img src="docs/images/error-notice.png" alt="Google Region Restriction Notice" width="680">
</p>

این جعبه‌ابزار مستقل سیستمی با **زبان Go خالص** مهندسی شده تا محدودیت‌های منطقه‌ای، ارورهای تحریمی ۴۰۳، فلگ‌های اهلیت اکانت و بن‌بست‌های سهمیه (HTTP 429) را در این بخش‌ها مرتفع سازد:
- **محیط Google Antigravity 2.x**
- **ادیتور Antigravity IDE**
- **خط فرمان رسمی Antigravity CLI (`agy`)**
- **افزونه رسمی وی‌اس‌کد (`google.google-antigravity`)**

### 🧬 کالبدشکافی لایه‌های مهندسی و پچ‌های هسته

| لایه هدف | مؤلفه اجرایی | تغییرات اعمال‌شده و استحکام‌بخشی |
| :--- | :--- | :--- |
| **۱. باینری هسته MultiGate** | `language_server` (باینری) | پچ الگوی بایت‌کد سطح ماشین (`RAX+8 hasValidAuth = 1`)؛ آزادسازی همیشگی گیت‌های پردازش کد ابری در مناطق تحریم‌شده. |
| **۲. فرانت‌اند ادیتور IDE** | `main.js` (جاوااسکریپت) | بازگشایی پرچم `isGoogleInternal` در الکترون، پاکسازی کش V8 و بی‌اثرسازی اعتبارسنجی سمت کاربر. |
| **۳. پنل گفتگو و چت** | `workbench.desktop.main.js` | رفع محدودیت‌های لوکیشن در پنل دستیار هوش مصنوعی و سایدبار. |
| **۴. سپر حفاظتی ASAR** | `app.asar` (آرشیو) | مهار کامل استثناهای لوله‌کشی ترمینال (`write EIO`) و بی‌صدا کردن لاگ‌های حجیم در پردازش ابری. |
| **۵. محافظ افزونه VS Code** | `extension.js` (اکستنشن) | تزریق گارد محافظتی عدم دانلود مجدد فایل‌های خام گوگل. |
| **۶. رپر پرچم‌های سازمانی** | رپر پروسس باینری | تزریق ۸ فلگ حریم خصوصی (`--disable_telemetry=true` و `--allow_all_sidecars=true`). |
| **۷. پروکسی مستقیم پروسس** | تزریق‌کننده SOCKS5 | هدایت مستقیم سوکت‌های پروسس به پورت‌های پروکسی محلی (`127.0.0.1:10808`، `:2080`، `:7890`) بدون فعال‌سازی سنگین TUN Mode. |

> 📖 **وایت‌پیپر فنی و تحلیل اسمبلی:** جهت مشاهده کدهای ماشین، بنچمارک‌های اسکن حافظه رم و تحلیل معکوس، مستند [**docs/ARCHITECTURE.md**](docs/ARCHITECTURE.md) را مطالعه کنید.

---

## 🛠️ مرجع جدول دستورات خط فرمان (CLI)

آنتی‌گراویتی کلینر به طور کامل بدون نیاز به رابط گرافیکی و از طریق ترمینال قابل اجرا و کنترل است:

| دستور (Command) | دسته‌بندی | شرح عملکرد و کاربرد |
| :--- | :---: | :--- |
| `ag-cleaner` | **تعاملی** | اجرای داشبورد تعاملی خط فرمان (TUI) با کلیدهای جهت‌نما. |
| `ag-cleaner gui` | **رابط کاربری** | اجرای پنجره گرافیکی پرتابل و مستقل دسکتاپ (GUI). |
| `ag-cleaner doctor` | **عیب‌یابی** | تست کامل سلامت سیستم، نشت DNS و اسکن خودکار پورت‌های فعال پروکسی. |
| `ag-cleaner doctor --verbose` | **عیب‌یابی عمیق** | تست تفصیلی با نمایش لاگ‌های زنده شبکه و مسیر سوکت‌ها. |
| `ag-cleaner patch all` | **پچ و آزادسازی** | پچ کد ماشین هسته (`hasValidAuth = 1`)، خط فرمان CLI، فرانت‌اند و سپر ASAR. |
| `ag-cleaner launch` | **اجرای هوشمند** | اجرای ادیتور با هدایت مستقیم ترافیک به پروکسی محلی (بدون نیاز به TUN Mode). |
| `ag-cleaner clean` | **ریست سهمیه** | پاکسازی جراحی ۱۶ کش سهمیه و ریت‌لیمیت بدون حذف حتی ۱ پیام از چت‌ها. |
| `ag-cleaner restore` | **بازیابی کارخانه** | بازگردانی اتمیک تمامی فایل‌ها به نسخه دست‌نخورده اولیه کارخانه (`.agybak`). |
| `ag-cleaner kill` | **مدیریت پروسس** | شناسایی و بستن تمام پروسس‌های هنگ‌کرده و زامبی `language_server`. |
| `ag-cleaner create-launcher` | **میانبر سیستم** | ساخت خودکار شورت‌کات دسکتاپ با تنظیمات آماده پروکسی و پرچم‌های سازمانی. |

---

## ❓ کلینیک خطایابی و پرسش‌های متداول (FAQ)

پاسخ سریع به متداول‌ترین سوالات کاربران در ادامه آمده است. جهت مشاهده کلینیک تخصصی ۱۰ خطا با اسکرین‌شات و تحلیل فورنزیک به [**docs/TROUBLESHOOTING.md**](docs/TROUBLESHOOTING.md) مراجعه فرمایید.

<details>
<summary><strong>🛡️ آیا استفاده از آنتی‌گراویتی کلینر امن است؟</strong></summary>
<br>

**بله، امنیت نرم‌افزار به صورت ۱۰۰٪ متن‌باز و شفاف قابل حسابرسی است:**
1. **عملکرد کاملاً آفلاین:** پچ و پاکسازی بدون نیاز به ذره‌ای اینترنت و منحصراً روی آدرس محلی `127.0.0.1` انجام می‌شود.
2. **زبان Go خالص (بدون وابستگی):** یک فایل مستقل تک‌باینری بدون فایل‌های مشکوک DLL یا اسکریپت‌های پایتونی خارجی.
3. **گذراندن ۱۰۰٪ ممیزی‌های امنیتی گیت‌هاب:** بررسی مداوم با اسکنرهای CodeQL و Govulncheck در [`.github/workflows/security.yml`](.github/workflows/security.yml).
4. **حفظ کامل تاریخچه گفتگوها:** بدون دست‌زدن به فایل‌های حیاتی چت و پروژه‌ها در `User/globalStorage`.
5. **معماری معتبر:** مستندشده در [`.github/SECURITY.md`](.github/SECURITY.md) و [`docs/ARCHITECTURE.md`](docs/ARCHITECTURE.md).

</details>

<details>
<summary><strong>⚡ ویژگی Auto-Fix چه کارهایی انجام می‌دهد و چرا دائمی است؟</strong></summary>
<br>

با کلیک روی دکمه **`[Auto-Fix]`** یا دستور `ag-cleaner patch all`:
۱. مقدار `RAX+8 hasValidAuth = 1` در باینری موتور هسته فعال می‌شود.  
۲. گیت‌های احراز هویت در خط فرمان `agy` بی‌اثر می‌شوند.  
۳. پرچم سازمانی `isGoogleInternal` در جاوااسکریپت الکترون باز می‌شود.  
۴. سپر ASAR از بروز کرش‌های `write EIO` جلوگیری می‌کند.  
۵. ۸ پرچم سازمانی حریم خصوصی و ضدتله‌متری تزریق می‌گردد.  
۶. پروکسی محلی به طور خودکار به پروسس ادیتور شناسانده می‌شود.

</details>

<details>
<summary><strong>🔄 بکاپ خودکار و بازیابی چگونه کار می‌کند؟</strong></summary>
<br>

* **بکاپ اولیه کارخانه (`.agybak`):** قبل از هر تغییری، یک نسخه دست‌نخورده از فایل‌ها ذخیره می‌شود ([`pkg/patcher/atomic.go`](pkg/patcher/atomic.go)).
* **بازگشت به حالت کارخانه با ۱ کلیک:** با کلیک روی دکمه **`[Restore]`** یا اجرای `ag-cleaner restore` فایل‌های اصلی بدون کم‌وکاست بازمی‌گردند.
* **بدون هیچ ردپای اضافه:** با حذف فایل باینری برنامه، سیستم به طور کامل و تمیز به حالت قبل برمی‌گردد.

</details>

<details>
<summary><strong>🔬 راهنمای خلاصه کدهای خطای متداول</strong></summary>
<br>

* **خطای `Agent terminated due to error` و `write EIO`:** اجرای `ag-cleaner patch all` (سپر ASAR خطاهای لوله‌کشی را خنثی کرده و پروب KeepAlive پانزده ثانیه‌ای تزریق می‌کند).
* **ارور `403 Forbidden` در آدرس `/auth-success`:** اجرای `ag-cleaner patch all` برای دور زدن فیلترهای گیت ورودی گوگل.
* **ارور `proxyconnect tcp: dial tcp 127.0.0.1:10808: connection refused`:** اجرای `ag-cleaner doctor` جهت شناسایی پورت‌های فعال کلاینت پروکسی شما.
* **خطای `Failed to fetch` و سقف سهمیه ۴۲۹:** اجرای `ag-cleaner clean` برای پاکسازی کش موقت بدون آسیب به چت‌ها.

👉 **مشاهده کلینیک تخصصی و اسکرین‌شات ۱۰ ارور در [docs/TROUBLESHOOTING.md](docs/TROUBLESHOOTING.md)**

</details>

---

## 📊 بنچمارک‌های زنده و جدول مقایسه تخصصی

<details>
<summary><strong>⚡ مشاهده بنچمارک‌های زنده عملکرد و ممیزی سیستم</strong></summary>
<br>

| شاخص سنجش | آنتی‌گراویتی کلینر پرو (نسخه ۵.۴.۰) | اسکریپت‌های شل و ابزارهای سنتی |
| :--- | :--- | :--- |
| **زبان و موتور اجرایی** | زبان مدرن Go 1.22+ (باینری کامپایل‌شده) | پایتون ۳ / بش / نود جی‌اس |
| **سرعت اسکن MultiGate** | **۱.۹۴ گیگابایت در ثانیه (۱۰۰٪ در حافظه رم)** | ۴۰ الی ۸۰ مگابایت در ثانیه (وابسته به هارد) |
| **سرعت تعویض Inode** | **۰.۰۰۱۸ ثانیه (تعویض اتمیک POSIX)** | نوشتن سنتی با خطر ارور `ETXTBSY` |
| **زمان پاکسازی ۱۶ منبع کش**| **۰.۰۰۲۱ ثانیه** | ۱.۸ الی ۳.۵ ثانیه |
| **میزان مصرف رم** | **۸.۲ مگابایت RSS** | ۸۵ الی ۱۸۰ مگابایت RSS |
| **تله‌متری و خروجی شبکه** | **۰.۰۰ کیلوبایت (تله‌متری کاملاً صفر)** | غیرقابل بررسی / دارای ریکوئست‌های خارجی |

</details>

<details>
<summary><strong>🥊 جدول مقایسه قابلیت‌ها با سایر راهکارها</strong></summary>
<br>

| قابلیت و معیار فنی | آنتی‌گراویتی کلینر پرو | وی‌پی‌ان سراسری (حالت TUN) | حذف دستی فایل‌ها |
| :--- | :---: | :---: | :---: |
| **تزریق مستقیم پروکسی بدون TUN** | ✅ اتصال داخلی سوکت SOCKS5 | ❌ اجبار به درگیر کردن کل سیستم | ❌ ندارد |
| **پچ باینری کد ماشین MultiGate** | ✅ پچ الگوی اینود رم با سرعت بالا | ❌ قادر به رفع تحریم داخلی کلاینت نیست | ❌ ندارد |
| **حفظ صددرصدی چت‌ها و پروژه‌ها** | ✅ حفظ کامل تاریخچه گفتگوها | ➖ نامربوط | ❌ خطر حذف پوشه `globalStorage` |
| **سپر ASAR ضد ارور `write EIO`** | ✅ مهار خطاهای قطع استریم خروجی | ❌ همچنان ارور می‌دهد | ❌ ندارد |
| **بازگشت به کارخانه با ۱ کلیک** | ✅ آنی و تضمینی (`.agybak`) | ❌ نیاز به نصب مجدد ادیتور | ❌ نیاز به بکاپ دستی |
| **عملکرد آفلاین و قابل حسابرسی** | ✅ ۱۰۰٪ لوکال‌هاست (`127.0.0.1`) | ⚠️ سرورهای ریموت فیلترشکن | ➖ دستی |

</details>

---

## 🤖 اسکیل هوش مصنوعی برای ایجنت‌ها (`skills/antigravity-cleaner`)

آنتی‌گراویتی کلینر مجهز به یک **اسکیل اختصاصی ایجنت هوش مصنوعی** ([`skills/antigravity-cleaner/SKILL.md`](skills/antigravity-cleaner/SKILL.md)) است که به دستیارهای کدنویسی هوشمند (Antigravity، Claude Code، Cursor و Codex) اجازه می‌دهد در صورت بروز خطاهای تحریم، قطعی شبکه یا سقف سهمیه، محیط توسعه را به صورت خودکار ترمیم کنند:

- **عیب‌یابی هوشمند بدون دخالت انسان**: اجرای اتوماتیک `ag-cleaner doctor` در زمان قطعی سوکت.
- **خودترمیمی سهمیه سقف مصرف**: پاکسازی جراحی کش هنگام بروز ارور ۴۲۹.
- **کانفیگ صفر**: شناسایی خودکار توسط ایجنت‌ها در ریشه پروژه.

---

## 🔬 آزمایشگاه تحقیقاتی امنیت هوش مصنوعی کاربردی چند-اکوسیستمی

پشت توسعه **Antigravity Cleaner Pro**، یک جریان پژوهشی فعال در حوزه امنیت ابزارهای توسعه هوش مصنوعی قرار دارد. در این آزمایشگاه تحقیقاتی، ابزارهای کدنویسی نسل جدید در چهار اکوسیستم اصلی **Google DeepMind (Antigravity و Gemini)، OpenAI، Anthropic (Claude Code) و Cursor** به صورت سامان‌مند مهندسی معکوس و کالبدشکافی می‌شوند:

- **بیش از ۱۰۰ پچ داخلی و پل ارتباطی:** مخزن تحقیقاتی محرمانه این آزمایشگاه شامل بیش از ۱۰۰ پچ بایت‌کد، مسدودکننده‌های تله‌متری، پل‌های پروتکل Stubby RPC و کانکتورهای مدل‌های شخصی (BYOM) است.
- **پروژه آغازین و پرچمدار:** آنتی‌گراویتی کلینر پروژه پرچمدار این آزمایشگاه است و ابزارهای مکمل سایر اکوسیستم‌های هوش مصنوعی نیز در دست توسعه هستند.
- **پایداری مهندسی برای توسعه‌دهندگان:** این پروژه برای فراهم ساختن تداوم کار برنامه‌نویسانی طراحی شده است که در شرایط انزوای دیجیتال و تحریم‌های ناعادلانه بین‌المللی فعالیت می‌کنند.
- **حمایت از پژوهش‌ها:** اگر این ابزار فضای کاری شما را آزاد و پایدار نگه داشته است، ثبت **ستاره (Star)** در این مخزن ⭐ مستقیماً به انتشار عمومی دستاوردهای جدید آزمایشگاه کمک خواهد کرد!

---

## 🚦 ثبت رسمی گزارش خطاها و بازخوردها

جهت تسریع در پیگیری و حل مشکلات:
۱. **فرم‌های ساختاریافته گیت‌هاب:** در صورت مشاهده هرگونه باگ، از طریق [فرم ثبت باگ گیت‌هاب](https://github.com/tawroot/antigravity-cleaner/issues/new?template=bug_report.yml) مورد را ارسال کنید.  
۲. **گفت‌وگو و راه‌اندازی پروکسی:** جهت تبادل نظر، بهینه‌سازی کانفیگ‌ها و راهنمایی‌ها به [تالار گفت‌وگوی گیت‌هاب (Discussions)](https://github.com/tawroot/antigravity-cleaner/discussions) بپیوندید.  
۳. **گزارش آسیب‌پذیری‌های امنیتی:** طبق سند [SECURITY.fa.md](.github/SECURITY.fa.md) با ما تماس بگیرید.

---

## 👥 سازندگان و حق کپی‌رایت

- **معمار و توسعه‌دهنده اصلی:** **[@dalroot](https://github.com/dalroot)**
- **دستیار و هم‌توسعه‌دهنده هوش مصنوعی:** **Antigravity AI (Google DeepMind)** — *خلق‌شده در قالب برنامه‌نویسی جفتی (Pair-Programming) با هوش مصنوعی گوگل برای آزادسازی ابزارهای گوگل برای تمام برنامه‌نویسان تحت تحریم جهان!* 🤖
- **مجوز نرم‌افزار:** لایسنس بین‌المللی آپاچی نسخه ۲.۰ (Apache-2.0)
- تقدیم به آزادی اینترنت، دسترسی برابر به دانش و همبستگی جهانی تمام برنامه‌نویسان.

---

## 🗺️ نقشه راه آینده (Roadmap)

اهداف فازهای v5.4.0 و نسخه آینده v6.0.0 در سند رسمی [**ROADMAP.fa.md**](ROADMAP.fa.md) قابل مشاهده است.

---

## 📈 تاریخچه ستاره‌های ریپازیتوری (Star History)

<a href="https://star-history.com/#tawroot/antigravity-cleaner&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=tawroot/antigravity-cleaner&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=tawroot/antigravity-cleaner&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=tawroot/antigravity-cleaner&type=Date" />
 </picture>
</a>

---

## 💬 تالار گفت‌وگو و ارتباط با جامعه کاربری

جهت ثبت سوالات، پیشنهادات و ارتباط مستقیم:  
[**ورود به تالار گفت‌وگوی گیت‌هاب (Discussions) →**](https://github.com/tawroot/antigravity-cleaner/discussions)

</div>
