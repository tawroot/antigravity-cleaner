# 🚀 Release Guide - Antigravity Cleaner v4.0

این راهنما نحوه ساخت و انتشار نسخه‌های جدید را توضیح می‌دهد.

---

## 📋 پیش‌نیازها

قبل از شروع، مطمئن شوید:
- ✅ تمام تغییرات commit و push شده‌اند
- ✅ `CHANGELOG.md` به‌روزرسانی شده است
- ✅ شماره نسخه در `version_info.txt` صحیح است
- ✅ تست‌های محلی انجام شده‌اند

---

## 🎯 روش 1: Release اتوماتیک با Tag

### مرحله 1: به‌روزرسانی شماره نسخه

فایل `version_info.txt` را ویرایش کنید:

```python
# تغییر شماره نسخه
filevers=(4, 0, 1, 0),  # نسخه جدید
prodvers=(4, 0, 1, 0),
# ...
StringStruct(u'FileVersion', u'4.0.1.0'),
StringStruct(u'ProductVersion', u'4.0.1.0'),
```

### مرحله 2: به‌روزرسانی CHANGELOG

فایل `CHANGELOG.md` را ویرایش کنید:

```markdown
## [4.0.1] - 2025-12-26

### ✨ New Features
- ویژگی جدید 1
- ویژگی جدید 2

### 🐛 Bug Fixes
- رفع باگ 1
- رفع باگ 2
```

### مرحله 3: Commit و Push

```bash
git add .
git commit -m "chore: bump version to v4.0.1"
git push origin main
```

### مرحله 4: ساخت Tag

```bash
# ساخت tag
git tag -a v4.0.1 -m "Release v4.0.1"

# Push tag به GitHub
git push origin v4.0.1
```

### مرحله 5: منتظر بمانید! 🎉

GitHub Actions به صورت خودکار:
1. ✅ بیلد برای Windows, macOS, Linux می‌گیرد
2. ✅ پکیج‌های Portable ZIP/TAR.GZ می‌سازد
3. ✅ Release در GitHub ایجاد می‌کند
4. ✅ فایل‌ها را آپلود می‌کند

**زمان تقریبی:** 15-20 دقیقه

پیشرفت را در [Actions](https://github.com/tawroot/antigravity-cleaner/actions) دنبال کنید.

---

## 🔧 روش 2: Build دستی (بدون Release)

برای تست بیلد بدون ایجاد release:

### از طریق GitHub UI:

1. به [Actions](https://github.com/tawroot/antigravity-cleaner/actions) بروید
2. روی **"Manual Build & Test"** کلیک کنید
3. **"Run workflow"** را بزنید
4. پلتفرم را انتخاب کنید:
   - `all` - همه پلتفرم‌ها
   - `windows` - فقط Windows
   - `macos` - فقط macOS
   - `linux` - فقط Linux
5. اگر می‌خواهید release هم بسازد، `Create GitHub Release` را فعال کنید
6. **"Run workflow"** را بزنید

### از طریق GitHub CLI:

```bash
# بیلد همه پلتفرم‌ها (بدون release)
gh workflow run manual-build.yml \
  -f platform=all \
  -f create_release=false

# بیلد فقط Windows
gh workflow run manual-build.yml \
  -f platform=windows \
  -f create_release=false

# بیلد همه + ساخت release
gh workflow run manual-build.yml \
  -f platform=all \
  -f create_release=true \
  -f release_tag=v4.0.1-beta
```

---

## 📦 ساختار پکیج‌های Portable

هر پکیج شامل موارد زیر است:

```
AntigravityCleaner-Portable/
├── AntigravityCleaner.exe (یا binary)
├── README.md
├── LICENSE
├── PORTABLE.txt
└── data/
```

### ویژگی‌های پکیج:

- ✅ **Portable کامل** - نیاز به نصب ندارد
- ✅ **بدون وابستگی** - همه چیز در یک فایل
- ✅ **Data محلی** - تنظیمات در پوشه `data/`
- ✅ **ZIP فشرده** - حجم کم برای دانلود سریع

---

## 🛡️ کاهش هشدارهای امنیتی

### Windows SmartScreen

برای جلوگیری از هشدار SmartScreen:

1. **امضای کد (Code Signing)** - بهترین راه‌حل:
   ```powershell
   # نیاز به گواهی Code Signing
   signtool sign /f cert.pfx /p password /t http://timestamp.digicert.com AntigravityCleaner.exe
   ```

2. **افزایش تعداد دانلود** - SmartScreen با افزایش دانلود، اعتماد می‌کند

3. **راهنمای کاربر** - در README توضیح دهید:
   ```markdown
   ⚠️ **Windows SmartScreen Warning**
   
   اولین بار که برنامه را اجرا می‌کنید، ممکن است هشدار ببینید.
   این طبیعی است چون برنامه امضا نشده است.
   
   برای اجرا:
   1. روی "More info" کلیک کنید
   2. "Run anyway" را بزنید
   ```

### macOS Gatekeeper

```bash
# حذف quarantine attribute
xattr -cr AntigravityCleaner.app

# یا در README:
chmod +x AntigravityCleaner
xattr -d com.apple.quarantine AntigravityCleaner
```

---

## 📊 بررسی وضعیت Build

### در GitHub Actions:

1. به [Actions](https://github.com/tawroot/antigravity-cleaner/actions) بروید
2. آخرین workflow را باز کنید
3. وضعیت هر job را بررسی کنید:
   - ✅ سبز = موفق
   - ❌ قرمز = خطا
   - 🟡 زرد = در حال اجرا

### لاگ‌های مفید:

- **Build logs** - جزئیات PyInstaller
- **Package logs** - ساخت ZIP/TAR
- **Release logs** - آپلود به GitHub

---

## 🐛 عیب‌یابی

### خطای "Module not found"

```yaml
# اضافه کردن به hidden-imports در workflow
--hidden-import=module_name
```

### خطای "Permission denied" در macOS/Linux

```bash
# اضافه کردن chmod در workflow
chmod +x dist/AntigravityCleaner
```

### خطای "Release already exists"

Workflow به صورت خودکار release قبلی را حذف می‌کند. اگر مشکل داشت:

```bash
# حذف دستی tag و release
git tag -d v4.0.1
git push origin :refs/tags/v4.0.1
gh release delete v4.0.1 --yes
```

---

## 📝 Checklist قبل از Release

- [ ] تست محلی در Windows انجام شده
- [ ] تست محلی در macOS/Linux (اختیاری)
- [ ] CHANGELOG.md به‌روزرسانی شده
- [ ] version_info.txt به‌روزرسانی شده
- [ ] README.md بررسی شده
- [ ] تمام تغییرات commit شده‌اند
- [ ] Tag ساخته و push شده

---

## 🎯 نکات مهم

### 1. شماره‌گذاری نسخه (Semantic Versioning)

```
v4.0.1
│ │ │
│ │ └─ Patch (رفع باگ)
│ └─── Minor (ویژگی جدید، سازگار با قبلی)
└───── Major (تغییرات بزرگ، ممکن است ناسازگار باشد)
```

### 2. Tag Pattern

- ✅ `v4.0.0` - Release اصلی
- ✅ `v4.0.1-beta` - نسخه Beta
- ✅ `v4.0.1-rc1` - Release Candidate
- ❌ `4.0.0` - بدون `v`

### 3. Workflow Triggers

- `push: tags: v4.*` - فقط tag‌های v4.x.x
- `workflow_dispatch` - اجرای دستی

---

## 🚀 دستورات سریع

### Release کامل:

```bash
# 1. به‌روزرسانی نسخه
# ویرایش version_info.txt و CHANGELOG.md

# 2. Commit
git add .
git commit -m "chore: release v4.0.1"
git push

# 3. Tag و Release
git tag -a v4.0.1 -m "Release v4.0.1"
git push origin v4.0.1

# 4. منتظر GitHub Actions بمانید (15-20 دقیقه)
```

### بررسی وضعیت:

```bash
# مشاهده workflow‌های در حال اجرا
gh run list --workflow=release-v4.yml

# مشاهده جزئیات آخرین run
gh run view

# دانلود artifacts
gh run download
```

### حذف Release ناموفق:

```bash
# حذف release و tag
gh release delete v4.0.1 --yes
git tag -d v4.0.1
git push origin :refs/tags/v4.0.1

# ساخت مجدد
git tag -a v4.0.1 -m "Release v4.0.1"
git push origin v4.0.1
```

---

## 📞 پشتیبانی

اگر مشکلی داشتید:

1. 📋 لاگ‌های GitHub Actions را بررسی کنید
2. 🐛 Issue در GitHub باز کنید
3. 💬 در Telegram پیام دهید: [@RAHBARUSD](https://t.me/RAHBARUSD)

---

**Powered by TAWANA NETWORK**  
© 2024-2025 Tawana Mohammadi. All Rights Reserved.
