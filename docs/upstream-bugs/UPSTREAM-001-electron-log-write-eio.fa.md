# 🐛 گزارش خطای بالادستی UPSTREAM-001: خطای کنترل‌نشده `write EIO` در ترنسپورت کنسول `electron-log`

- **شناسه باگ:** `UPSTREAM-001`
- **مؤلفه تحت تأثیر:** `resources/app.asar/node_modules/electron-log/src/node/transports/console.js`
- **محصول هدف:** برنامه دسکتاپ Google Antigravity IDE (پروسس اصلی الکترون)
- **محیط بازتولید:** لینوکس (تمام توزیع‌ها)، مک و ویندوز در اجرای بدون ترمینال (Detached Launches)
- **وضعیت در پروژه:** ✅ مهارشده از طریق ایزولاسیون Stdio در `pkg/proxy/launcher.go` + برنامه‌ریزی غیرفعال‌سازی ترنسپورت کنسول در نسخه v5.4.0

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

## 🛡️ استراتژی خنثی‌سازی در Antigravity Cleaner Pro

### لایه ۱: ایزولاسیون در سطح سیستم‌عامل (`pkg/proxy/launcher.go`)
در لانچر Go و اسکریپت‌های دسکتاپ:
- `cmd.Stdout = nil`
- `cmd.Stderr = nil`
- `cmd.Stdin = nil`
- هدایت خروجی در اسکریپت‌های شل با `exec "$APP_PATH" ... >/dev/null 2>&1`

با اتصال استریم‌های خروجی به `/dev/null`، بسته شدن ترمینال والد هیچ سیگنال شکستگی لوله‌ای تولید نمی‌کند.

### لایه ۲: خاموش‌سازی ترنسپورت کنسول در فرانت‌اند (نسخه v5.4.0)
تنظیم مقدار `log.transports.console.level = false` در داخل `app.asar` که فراخوانی متد `console.warn` در پروسس اصلی الکترون را به طور کامل متوقف می‌کند، بدون اینکه هیچ تداخلی با کارکرد ادیتور داشته باشد.
