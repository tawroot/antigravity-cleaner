# 🐛 گزارش خطای بالادستی UPSTREAM-001: خطای کنترل‌نشده `write EIO` در ترنسپورت کنسول `electron-log`

- **شناسه باگ:** `UPSTREAM-001`
- **مؤلفه تحت تأثیر:** `resources/app.asar/node_modules/electron-log/src/node/transports/console.js`
- **محصول هدف:** برنامه دسکتاپ Google Antigravity IDE (پروسس اصلی الکترون)
- **محیط بازتولید:** لینوکس (تمام توزیع‌ها)، مک و ویندوز در اجرای بدون ترمینال (Detached Launches)
- **وضعیت در پروژه:** 🟡 **در دست اقدام (در سطح لانچر مهار موقت شده؛ پچ ریشه‌ای در سطح باینری/ASAR در فاز v5.4.0 در حال توسعه است و هنوز باگ کاملاً برطرف‌شده محسوب نمی‌شود)**

---

## 🔍 متن دقیق خطای پشته (Stack Trace)

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

## 🔬 شواهد و تحلیل مهندسی معکوس با موتور REA

با استفاده از ماژول `analyze_javascript_application` در فریم‌ورک مهندسی معکوس REA، کل پکیج `app.asar` مورد کالبدشکافی ایستا قرار گرفت:

* **شناسه مدرک مهندسی معکوس (REA Evidence ID):**  
  `ev_0b0572bf2f393d9d64e50173b4095cb2f4366a26b919886eafb0fa476ec1d9da`
* **هش SHA-256 آرشیو ریشه (`app.asar`):**  
  `a75535fec1f039be8e78c6a66f3e41f15969101fcf44a0690b7dce4d998b875b`
* **آمار درخت AST استخراج‌شده:**  
  - تعداد فایل‌های جاوااسکریپت پارس‌شده: **۶۸۲ فایل**
  - نودهای درخت نحوی (Visited AST Nodes): **۲,۱۹۴,۲۳۲ نود**
  - ارتباطات معنایی و ماژول‌ها: **۵۶۱ ماژول** و **۳,۴۳۱ یافته تحلیلی**

### کد دقیق استخراج‌شده از درون `app.asar` (سورس فایل `console.js`):
در خطوط ۴۵ تا ۶۳ فایل `node_modules/electron-log/src/node/transports/console.js`، باگ دقیقاً در خط ۶۰ قرار دارد:

```javascript
// خط ۴۶: سطح لاگ به صورت پیش‌فرض روی 'silly' قرار دارد (چاپ تمام هشدارهای ریز و درشت)
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

// خطوط ۵۸ تا ۶۱: عدم وجود بلوک try/catch روی خروجی متد کنسول
writeFn({ message }) {
  const consoleLogFn = consoleMethods[message.level] || consoleMethods.info;
  consoleLogFn(...message.data); // <-- خط ۶۰: شلیک خطا روی پایپ بسته ترمینال!
},
```

---

## 🔬 تحلیل ریشه‌ای (Root Cause Analysis)

### ۱. نقص لوله شکسته (Broken Pipe Defect)
در ساختار Node.js / Electron، نوشتن روی خروجی استاندارد (`process.stdout`) یک استریم غیرهم‌زمان سوکتی است. وقتی برنامه از طریق یک ترمینال، شل یا پروسس فرزند باز می‌شود، دیسکریپتورهای ورودی و خروجی برنامه به آن ترمینال متصل می‌مانند.

اگر ترمینال والد بعداً بسته شود بدون اینکه این لوله‌ها ایزوله شده باشند:
1. هسته سیستم‌عامل وضعیت لوله را «بسته/قطع‌شده» علامت‌گذاری می‌کند.
2. اولین تلاش بعدی برای نوشتن روی `process.stdout` با خطای سیستمی `EIO` (خطای Input/Output روی پایپ قطع‌شده) متوقف می‌شود.
3. در متد `transport.writeFn` کتابخانه `electron-log`، هیچ بلوک `try / catch` یا لیسنر خطای `process.stdout.on('error', ...)` قرار داده نشده است.
4. رانتایم الکترون این رویداد را به عنوان یک استثنای کنترل‌نشده (Uncaught Exception) در نظر گرفته و پنجره اخطار جاوااسکریپت را باز می‌کند.

### ۲. رویدادهای محرک شلیک لاگ در آنتی‌گراویتی (Trigger Events)
چرا کتابخانه لاگ گوگل در پس‌زمینه سعی می‌کند متنی در کنسول بنویسد؟
1. **ثبت رخداد افت تله‌متری:** با فعال بودن پرچم‌های Zero-Telemetry (`--disable_telemetry=true`)، ماژول ردگیری گوگل نمی‌تواند دیتا را به سرورهای `cloudcode-pa.googleapis.com` بفرستد، بنابراین یک اخطار داخلی در کنسول ثبت می‌کند.
2. **بررسی آپدیت خودکار (AutoUpdater):** برنامه در پس‌زمینه تلاش می‌کند برای دریافت نگارش جدید چک کند؛ در صورت تاخیر یا ایزولاسیون شبکه، پیام `log.warn()` صادر می‌کند.
3. **قطع موقت سوکت‌های IPC:** در زمان عدم فعالیت کاربر در ادیتور.

---

## 🛡️ استراتژی خنثی‌سازی و فازبندی در Antigravity Cleaner Pro

### فاز ۱ (اقدام انجام‌شده - مهار در لایه سیستم‌عامل):
در لانچر Go (`pkg/proxy/launcher.go`) و اسکریپت‌های دسکتاپ:
- `cmd.Stdout = nil`
- `cmd.Stderr = nil`
- `cmd.Stdin = nil`
- هدایت خروجی در اسکریپت‌های شل با `exec "$APP_PATH" ... >/dev/null 2>&1`
این کار از باز شدن دیالوگ جلوگیری می‌کند، اما باگ در ریشه کدهای الکترون کماکان زنده است.

### فاز ۲ (در دست توسعه در v5.4.0 - خنثی‌سازی ریشه‌ای در سطح ASAR):
* پچ جراحی بایت‌کد در `app.asar` روی فایل `console.js`:  
  تغییر مقدار `level: 'silly'` به `level: false`، یا پیچیدن خط ۶۰ در بلوک `try { consoleLogFn(...); } catch (_) {}`.  
  این پچ ریشه‌ای تضمین می‌کند که حتی اگر برنامه در ترمینال باز و بسته شود، ادیتور ابداً اقدام به چاپ هشدار در کنسول نخواهد کرد.
