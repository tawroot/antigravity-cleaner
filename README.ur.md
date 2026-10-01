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

<h1 align="center">⚡ Antigravity Cleaner — ڈیجیٹل آزادی برائے مصنوعی ذہانت (v5.2.0)</h1>

<p align="center">
  <strong>گوگل اینٹی گریوٹی کے لیے خودکار تشخیص، علاقائی پابندیوں کا خاتمہ اور نو-ٹن اسمارٹ پراکسی انجن</strong>
</p>

<div align="center">

  [![Version](https://img.shields.io/badge/ورژن-5.2.0-blue?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner/releases)
  [![Go Report](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
  [![Tests](https://img.shields.io/badge/ٹیسٹ-100%25_کامیاب-brightgreen?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Benchmark](https://img.shields.io/badge/اسکین_رفتار-1.9_GB%2Fs-blueviolet?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Platform](https://img.shields.io/badge/پلیٹ_فارم-Linux%20%7C%20macOS%20%7C%20Windows-brightgreen.svg?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner)
  [![License](https://img.shields.io/badge/لائسنس-Apache_2.0-blue?style=for-the-badge)](LICENSE)
  [![Security](https://img.shields.io/badge/پرائیویسی-100%25_آف_لائن_%7C_زیرو_ٹریکنگ-success?style=for-the-badge)]()
</div>

> *دلی خلوص کے ساتھ ان تمام ڈیولپرز اور محققین کے نام جو داخلی انٹرنیٹ فلٹرنگ اور بین الاقوامی ڈیجیٹل پابندیوں کی زد میں ہیں — پاکستان، ایران، چین، روس، ترکی، اور دنیا کے تمام پابندی زدہ خطوں میں۔ علم، ٹیکنالوجی اور مصنوعی ذہانت تک آزادانہ رسائی ہر انسان کا بنیادی حق ہے۔*  
> — **@dalroot**

<div align="center">

[![GitHub Stars](https://img.shields.io/github/stars/tawroot/antigravity-cleaner?style=social)](https://github.com/tawroot/antigravity-cleaner/stargazers)
&nbsp;
**⭐ اگر اس ٹول نے گوگل کی پابندیوں یا کوٹہ کی غلطیوں کو حل کرنے میں آپ کی مدد کی ہے، تو براہ کرم گٹ ہب پر ایک اسٹار ضرور دیں!**
&nbsp;
[![GitHub Forks](https://img.shields.io/github/forks/tawroot/antigravity-cleaner?style=social)](https://github.com/tawroot/antigravity-cleaner/network/members)

</div>

---

## ⚡ اینٹی گریوٹی کلینر کیا ہے؟

<p align="center">
  <img src="docs/images/error-notice.png" alt="گوگل ریجن پابندی نوٹس" width="680">
</p>

**Antigravity Cleaner** ایک جدید اور تیز رفتار سسٹم ٹول ہے جسے **خالص Go لینگویج** میں تیار کیا گیا ہے۔ یہ ون-کلک ریٹرو گرافیکل انٹرفیس (**Retro KeyGen GUI**) اور جدید ٹرمینل ڈیش بورڈ (**Lipgloss & Bubbletea**) دونوں فراہم کرتا ہے۔ یہ گوگل اینٹی گریوٹی کے تمام مسائل بشمول 403 Forbidden (User location is not supported)، 429 کوٹہ کی بندش، اور کوڈ اسٹریمنگ کے انقطاع کو مکمل طور پر حل کرتا ہے:
- **Google Antigravity 2.x**
- **Antigravity IDE**
- **Antigravity CLI (`agy`)**
- **آفیشل VS Code ایکسٹینشن (`google.google-antigravity`)**

روایتی سکرپٹس کے برعکس جو پورے آپریٹنگ سسٹم کو بھاری **TUN (VPN)** موڈ پر مجبور کرتی ہیں، یہ ٹول بغیر روٹ اور بغیر سسٹم ٹن کے، ٹریفک کو براہ راست مقامی پراکسی (Clash, Xray, Hiddify) سے جوڑتا ہے۔

> 🔒 **100٪ آف لائن اور زیرو ٹریکنگ کی ضمانت:**  
> یہ ٹول مکمل طور پر آپ کے اپنے کمپیوٹر پر لوکل طور پر چلتا ہے۔ اس میں کوئی ٹیلی میٹری، ڈیٹا جمع کرنے یا بیرونی کنکشن کا عمل موجود نہیں ہے۔ کوڈ اور چیٹس بالکل محفوظ رہتے ہیں۔

---

## 💾 1-کلک ریٹرو گرافیکل انٹرفیس (ٹرمینل کے بغیر)

کیا آپ گرافیکل انٹرفیس کو ترجیح دیتے ہیں؟ [Releases](https://github.com/tawroot/antigravity-cleaner/releases) سے اپنے سسٹم کے لیے فائل ڈاؤن لوڈ کریں اور **ڈبل کلک** کر کے چلائیں:

<div align="center">
  <img src="assets/shot.jpg" alt="اینٹی گریوٹی ریٹرو انٹرفیس" width="440">
  <br>
  <sub><em>کلاسک ونڈوز 95 اسٹائل، اینیمیٹڈ اسٹار فیلڈ اور ون-کلک آٹو فکس بٹن کے ساتھ۔</em></sub>
</div>

<br>

| آپریٹنگ سسٹم | ڈائریکٹ بائنری | طریقہ کار |
| :--- | :---: | :--- |
| 🪟 **Windows** (64-bit) | [![براہ راست ڈاؤن لوڈ](https://img.shields.io/badge/ڈاؤن_لوڈ-.exe-0078D6?style=for-the-badge&logo=windows&logoColor=white)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-windows-amd64.exe) | ڈبل کلک `.exe` (خودکار GUI کھولتا ہے) |
| 🍎 **macOS** (Apple Silicon M1–M4) | [![براہ راست ڈاؤن لوڈ](https://img.shields.io/badge/ڈاؤن_لوڈ-ARM64-white?style=for-the-badge&logo=apple&logoColor=black)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-arm64) | ڈبل کلک یا `./antigravity-cleaner-darwin-arm64` |
| 🍎 **macOS** (Intel) | [![براہ راست ڈاؤن لوڈ](https://img.shields.io/badge/ڈاؤن_لوڈ-Intel_x64-gray?style=for-the-badge&logo=apple&logoColor=white)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-amd64) | ڈبل کلک یا `./antigravity-cleaner-darwin-amd64` |
| 🐧 **Linux** (x86_64) | [![براہ راست ڈاؤن لوڈ](https://img.shields.io/badge/ڈاؤن_لوڈ-x86__64-FCC624?style=for-the-badge&logo=linux&logoColor=black)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-amd64) | ڈبل کلک یا `ag-cleaner gui` |
| 🐧 **Linux** (ARM64) | [![براہ راست ڈاؤن لوڈ](https://img.shields.io/badge/ڈاؤن_لوڈ-ARM64-FCC624?style=for-the-badge&logo=linux&logoColor=black)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-arm64) | ڈبل کلک یا `ag-cleaner gui` |

---

## 🚀 فوری ٹرمینل انسٹالیشن

### Linux / macOS
```bash
curl -fsSL https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.sh | bash
```

### Windows (PowerShell)
```powershell
irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex
```

---

## 🛠️ کمانڈ لائن سب کمانڈز

```bash
# سسٹم کی صحت اور پراکسی کنکشن کا تفصیلی معائنہ
ag-cleaner doctor

# کوٹہ اور بلاکس کی صفائی (آپ کے پروجیکٹ کوڈ کو چھیڑے بغیر)
ag-cleaner clean

# ریجن اور 403 پابندی کو ختم کرنے کے لیے بائٹ کوڈ پیچ
ag-cleaner patch

# بغیر TUN موڈ کے پراکسی کنکشن
ag-cleaner proxy --socks5 127.0.0.1:10808

# ادیتور کو محفوظ پراکسی ماحول میں چلانا
ag-cleaner run
```

---

## 📜 لائسنس
یہ پراجیکٹ بین الاقوامی **Apache License 2.0** کے تحت شائع کیا گیا ہے۔
