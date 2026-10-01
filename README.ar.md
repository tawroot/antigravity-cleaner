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

<h1 align="center">⚡ Antigravity Cleaner — أداة الحرية الرقمية للذكاء الاصطناعي (v5.2.0)</h1>

<p align="center">
  <strong>محرك التشخيص الذاتي، وتجاوز العقوبات الإقليمية، وتصفير الحصص لبيئة Google Antigravity IDE</strong>
</p>

<div align="center">

  [![Version](https://img.shields.io/badge/الإصدار-5.2.0-blue?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner/releases)
  [![Go Report](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
  [![Tests](https://img.shields.io/badge/الاختبارات-100%25_ناجحة-brightgreen?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Benchmark](https://img.shields.io/badge/سرعة_الفحص-1.9_جيجابايت%2Fثانية-blueviolet?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Platform](https://img.shields.io/badge/الأنظمة-Linux%20%7C%20macOS%20%7C%20Windows-brightgreen.svg?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner)
  [![License](https://img.shields.io/badge/الترخيص-Apache_2.0-blue?style=for-the-badge)](LICENSE)
  [![Security](https://img.shields.io/badge/الخصوصية-100%25_أوفلاين_%7C_بدون_تتبع-success?style=for-the-badge)]()
</div>

> *إهداء من صميم القلب إلى جميع المطورين والباحثين والمبدعين المحاصرين بين قيود الحجب الداخلية والعقوبات الرقمية الدولية — في 🇸🇾 سوريا، 🇸🇩 السودان، 🇮🇷 إيران، 🇨🇺 كوبا، 🇨🇳 الصين، 🇷🇺 روسيا، 🇾🇪 اليمن، وجميع المناطق الخاضعة للقيود في العالم. إن الوصول الحر إلى المعرفة والذكاء الاصطناعي والتقنية هو حق أساسي من حقوق الإنسان وليس امتيازاً.*  
> — **@dalroot**

<div align="center">

[![GitHub Stars](https://img.shields.io/github/stars/tawroot/antigravity-cleaner?style=social)](https://github.com/tawroot/antigravity-cleaner/stargazers)
&nbsp;
**⭐ إذا ساعدتك هذه الأداة في تخطي عقوبات جوجل أو حل أخطاء الحصص، يرجى دعمنا بنجمة (Star) لمساعدة المطورين الآخرين في العثور عليها!**
&nbsp;
[![GitHub Forks](https://img.shields.io/github/forks/tawroot/antigravity-cleaner?style=social)](https://github.com/tawroot/antigravity-cleaner/network/members)

</div>

---

## ⚡ ما هي أداة Antigravity Cleaner؟

<p align="center">
  <img src="docs/images/error-notice.png" alt="تنبيه قيود المنطقة الجغرافية من جوجل" width="680">
</p>

أداة **Antigravity Cleaner** هي حزمة برمجية عالية الأداء مكتوبة بلغة **Go النقية**. توفر واجهة رسومية كلاسيكية بنقرة واحدة **1-Click Retro KeyGen GUI** إلى جانب لوحة تحكم تفاعلية في الطرفية (**Lipgloss & Bubbletea**). تحل الأداة بشكل جذري مشاكل الحظر الجغرافي وقيود الحسابات (HTTP 403 Forbidden / Location not supported) واستنفاد الحصص (HTTP 429 Quota Exhausted) وانقطاع تدفق الأكواد في:
- **Google Antigravity 2.x**
- **محرر Antigravity IDE**
- **واجهة الأوامر Antigravity CLI (`agy`)**
- **إضافة VS Code الرسمية (`google.google-antigravity`)**

على عكس السكربتات القديمة التي تفرض تشغيل وضع **TUN (VPN)** الثقيل لكامل النظام، تتضمن الأداة **حاقن بروكسي ذكي بدون TUN** يوجه ترافيك المحرر مباشرة إلى تطبيقات البروكسي المحلية (Xray, Clash, NekoBox, Hiddify) دون أي إبطاء للنظام.

> 🔒 **ضمان العمل دون اتصال 100% وانعدام التتبع:**  
> تعمل الأداة محلياً بالكامل على جهازك (`127.0.0.1`). لا تحتوي على أي تتبع عن بُعد (Telemetry) أو إحصائيات أو اتصالات خارجية. تبقى مشاريعك ومحادثاتك آمنة تماماً. المشروع مفتوح المصدر بالكامل وفق رخصة Apache License 2.0.

---

## 💾 تشغيل الواجهة الرسومية بنقرة واحدة (دون الحاجة للطرفية)

هل تفضل استخدام واجهة رسومية بسيطة؟ حمّل الملف المخصص لنظامك من صفحة [Releases](https://github.com/tawroot/antigravity-cleaner/releases) وافتحه بنقرة مزدوجة (Double-click):

<div align="center">
  <img src="assets/shot.jpg" alt="واجهة Antigravity Retro KeyGen GUI (ستايل ويندوز 95)" width="440">
  <br>
  <sub><em>واجهة كلاسيكية بنمط Windows 95 مع خلفية متحركة ومؤثرات صوتية ريترو 8-بت وزر إصلاح تلقائي بنقرة واحدة.</em></sub>
</div>

<br>

| نظام التشغيل | الملف التنفيذي المستقل | طريقة التشغيل |
| :--- | :---: | :--- |
| 🪟 **Windows** (64-bit) | [![تحميل مباشر](https://img.shields.io/badge/تحميل_مباشر-.exe-0078D6?style=for-the-badge&logo=windows&logoColor=white)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-windows-amd64.exe) | نقرة مزدوجة على `.exe` (تفتح الواجهة تلقائياً) |
| 🍎 **macOS** (Apple Silicon M1–M4) | [![تحميل مباشر](https://img.shields.io/badge/تحميل_مباشر-ARM64-white?style=for-the-badge&logo=apple&logoColor=black)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-arm64) | نقرة مزدوجة أو تشغيل `./antigravity-cleaner-darwin-arm64` |
| 🍎 **macOS** (Intel) | [![تحميل مباشر](https://img.shields.io/badge/تحميل_مباشر-Intel_x64-gray?style=for-the-badge&logo=apple&logoColor=white)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-amd64) | نقرة مزدوجة أو تشغيل `./antigravity-cleaner-darwin-amd64` |
| 🐧 **Linux** (x86_64) | [![تحميل مباشر](https://img.shields.io/badge/تحميل_مباشر-x86__64-FCC624?style=for-the-badge&logo=linux&logoColor=black)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-amd64) | نقرة مزدوجة أو تشغيل `ag-cleaner gui` |
| 🐧 **Linux** (ARM64) | [![تحميل مباشر](https://img.shields.io/badge/تحميل_مباشر-ARM64-FCC624?style=for-the-badge&logo=linux&logoColor=black)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-arm64) | نقرة مزدوجة أو تشغيل `ag-cleaner gui` |

---

## 🚀 التثبيت السريع للطرفية (أمر واحد)

### أنظمة Linux و macOS
```bash
curl -fsSL https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.sh | bash
```

### نظام Windows (PowerShell)
```powershell
irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex
```

---

## 🛠️ أوامر الطرفية الرئيسية

```bash
# فحص وتشخيص شامل لصحة البيئة والمنافذ والاتصال
ag-cleaner doctor

# تصفير جراحي للحصص ومسح الملفات العالقة (دون حذف مشاريعك)
ag-cleaner clean

# تطبيق باتش البايت كود لتعطيل فحص الحظر الإقليمي (403)
ag-cleaner patch

# توجيه ترافيك المحرر مباشرة إلى البروكسي المحلي دون نمط TUN
ag-cleaner proxy --socks5 127.0.0.1:10808

# تشغيل المحرر في بيئة محمية مع استقرار التدفق
ag-cleaner run
```

---

## 📜 الترخيص
هذا المشروع منشور تحت ترخيص **Apache License 2.0** الدولي المفتوح.
