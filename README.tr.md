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

<h1 align="center">⚡ Antigravity Cleaner — Yapay Zeka Özgürlük Araç Seti (v5.2.0)</h1>

<p align="center">
  <strong>Google Antigravity IDE İçin Kurumsal Kendi Kendini Onarma, Yaptırım Aşma ve Teşhis Motoru</strong>
</p>

<div align="center">

  [![Version](https://img.shields.io/badge/Sürüm-5.2.0-blue?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner/releases)
  [![Go Report](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
  [![Tests](https://img.shields.io/badge/Testler-100%25_Geçti-brightgreen?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Benchmark](https://img.shields.io/badge/Baytkod_Tarama-1.9_GB%2Fs-blueviolet?style=for-the-badge)](docs/test-reports/benchmarks.log)
  [![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-brightgreen.svg?style=for-the-badge)](https://github.com/tawroot/antigravity-cleaner)
  [![License](https://img.shields.io/badge/Lisans-Apache_2.0-blue?style=for-the-badge)](LICENSE)
  [![Security](https://img.shields.io/badge/Gizlilik-100%25_Çevrimdışı_%7C_Sıfır_Takip-success?style=for-the-badge)]()
</div>

> *Yurt içi internet filtrelemesi ile uluslararası dijital yaptırımların arasında sıkışıp kalan tüm geliştiricilere, araştırmacılara ve içerik üreticilerine içtenlikle ithaf edilmiştir — 🇹🇷 Türkiye, 🇮🇷 İran, 🇨🇳 Çin, 🇷🇺 Rusya, 🇧🇾 Belarus, 🇨🇺 Küba, 🇸🇾 Suriye ve dünyanın tüm kısıtlı bölgelerinde. Bilgiye, teknolojiye ve yapay zekaya özgür erişim temel bir insan hakkıdır.*  
> — **@dalroot**

<div align="center">

[![GitHub Stars](https://img.shields.io/github/stars/tawroot/antigravity-cleaner?style=social)](https://github.com/tawroot/antigravity-cleaner/stargazers)
&nbsp;
**⭐ Antigravity Cleaner, Google yaptırımlarını veya kota hatalarını aşmanıza yardımcı olduysa, lütfen bir Star verin — bu, kısıtlı bölgelerdeki diğer geliştiricilerin projeyi keşfetmesine yardımcı olur!**
&nbsp;
[![GitHub Forks](https://img.shields.io/github/forks/tawroot/antigravity-cleaner?style=social)](https://github.com/tawroot/antigravity-cleaner/network/members)

</div>

---

## ⚡ Antigravity Cleaner Nedir?

<p align="center">
  <img src="docs/images/error-notice.png" alt="Google Bölge Kısıtlama Bildirimi" width="680">
</p>

**Antigravity Cleaner**, **saf Go** ile geliştirilmiş kurumsal düzeyde yüksek performanslı bir sistem aracıdır. Hem nostaljik **1-Tıkla Retro KeyGen Arayüzü (GUI)** hem de modern interaktif terminal paneli (**Lipgloss & Bubbletea**) sunar. Aşağıdaki araçlarda bölgesel engellemeleri, 403 Forbidden ("User location is not supported") hatalarını, 429 kota kilitlenmelerini ve kod akış kesintilerini kökten çözer:
- **Google Antigravity 2.x**
- **Antigravity IDE**
- **Antigravity CLI (`agy`)**
- **Resmi VS Code Eklentisi (`google.google-antigravity`)**

Kullanıcıları tüm işletim sistemini ağır ve root yetkisi gerektiren **TUN Modu (VPN)** üzerinden yönlendirmeye zorlayan geleneksel araçların aksine Antigravity Cleaner, Antigravity trafiğini doğrudan yerel proxy istemcilerine (Xray, Clash, NekoBox, Hiddify) sıfır sistem yüküyle ileten **TUN Gerektirmeyen Akıllı Proxy Enjektörü** içerir.

> 🔒 **%100 Çevrimdışı ve Sıfır Takip Garantisi:**  
> Antigravity Cleaner **yalnızca yerel makinenizde (`127.0.0.1`)** çalışır. İçerisinde **SIFIR telemetri, SIFIR analiz ve SIFIR uzak ağ çağrısı** bulunur. Kaynak kodlarınız, API anahtarlarınız ve sohbetleriniz asla kaydedilmez veya iletilmez. Apache License 2.0 altında tamamen açık kaynaklıdır.

---

## 💾 1-Tıkla Retro KeyGen GUI (Terminal Gerektirmez)

Basit bir grafik arayüz mü tercih ediyorsunuz? İşletim sisteminize uygun dosyayı [Releases](https://github.com/tawroot/antigravity-cleaner/releases) sayfasından indirin ve **çift tıklayın**. Terminal komutu, Python veya harici bağımlılık gerekmez!

<div align="center">
  <img src="assets/shot.jpg" alt="Antigravity Retro Arayüzü (Win95 Stili)" width="440">
  <br>
  <sub><em>Windows 95/98 klasik stili, canlı yıldız alanı efekti, 8-bit retro sesler ve tek tıkla otomatik onarım.</em></sub>
</div>

<br>

| İşletim Sistemi | Bağımsız Dosya | Çalıştırma Modu |
| :--- | :---: | :--- |
| 🪟 **Windows** (64-bit) | [![Doğrudan İndir](https://img.shields.io/badge/Doğrudan_İndir-.exe-0078D6?style=for-the-badge&logo=windows&logoColor=white)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-windows-amd64.exe) | `.exe` dosyasına çift tıklayın (GUI otomatik açılır) |
| 🍎 **macOS** (Apple Silicon M1–M4) | [![Doğrudan İndir](https://img.shields.io/badge/Doğrudan_İndir-ARM64-white?style=for-the-badge&logo=apple&logoColor=black)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-arm64) | Çift tıklayın veya `./antigravity-cleaner-darwin-arm64` |
| 🍎 **macOS** (Intel) | [![Doğrudan İndir](https://img.shields.io/badge/Doğrudan_İndir-Intel_x64-gray?style=for-the-badge&logo=apple&logoColor=white)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-darwin-amd64) | Çift tıklayın veya `./antigravity-cleaner-darwin-amd64` |
| 🐧 **Linux** (x86_64) | [![Doğrudan İndir](https://img.shields.io/badge/Doğrudan_İndir-x86__64-FCC624?style=for-the-badge&logo=linux&logoColor=black)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-amd64) | Çift tıklayın veya `ag-cleaner gui` |
| 🐧 **Linux** (ARM64) | [![Doğrudan İndir](https://img.shields.io/badge/Doğrudan_İndir-ARM64-FCC624?style=for-the-badge&logo=linux&logoColor=black)](https://github.com/tawroot/antigravity-cleaner/releases/latest/download/antigravity-cleaner-linux-arm64) | Çift tıklayın veya `ag-cleaner gui` |

---

## 🚀 Hızlı Terminal Kurulumu (Tek Satır)

### Linux / macOS
```bash
curl -fsSL https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.sh | bash
```

### Windows (PowerShell)
```powershell
irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex
```

---

## 🛠️ Temel Terminal Komutları

```bash
# Sistem sağlığı, soketler ve proxy değişkenlerini kapsamlı kontrol etme
ag-cleaner doctor

# Kota ve kilit dosyalarını cerrahi temizleme (kodlarınıza dokunmaz)
ag-cleaner clean

# Bölgesel kısıtlamayı (403) aşmak için doğrudan baytkod yaması uygulama
ag-cleaner patch

# Sistem geneli TUN açmadan doğrudan yerel proxy'ye bağlanma
ag-cleaner proxy --socks5 127.0.0.1:10808

# Antigravity IDE'yi kararlı bağlantı ile başlatma
ag-cleaner run
```

---

## 📜 Lisans
Bu proje uluslararası standart **Apache License 2.0** ile lisanslanmıştır.
