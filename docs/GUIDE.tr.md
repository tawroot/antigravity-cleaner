# 📖 Antigravity Cleaner — Kapsamlı Kullanım Kılavuzu (v5.2.0)

**Google Antigravity IDE**, `agy` CLI ve Google Antigravity eklentisi için bölgesel yaptırımları aşma, 403/429 hatalarını çözme ve TUN modu olmadan akıllı proxy yapılandırma rehberi.

---

## 📑 İçindekiler
1. [Çözülen Temel Sorunlar](#-çözülen-temel-sorunlar)
2. [Arayüz Seçenekleri](#-arayüz-seçenekleri)
   - [1-Tıkla Retro KeyGen Arayüzü (Terminal Gerektirmez)](#1-1-tıkla-retro-keygen-arayüzü-gui)
   - [Modern Etkileşimli Terminal (TUI)](#2-modern-etkileşimli-terminal-arayüzü-tui)
   - [CLI Komut Satırı](#3-cli-komut-satırı)
3. [CLI Komut Başvurusu](#-cli-komut-başvurusu)
   - [`doctor` — Sistem Sağlık ve Tanı Kontrolü](#ag-cleaner-doctor)
   - [`clean` — Cerrahi Kota ve Önbellek Sıfırlama](#ag-cleaner-clean)
   - [`patch` — Bölge Kısıtlamasını Devre Dışı Bırakma](#ag-cleaner-patch)
   - [`proxy` — TUN Olmadan Akıllı Soket Enjeksiyonu](#ag-cleaner-proxy)
   - [`run` — IDE'yi Proxy ile Başlatma](#ag-cleaner-run)
4. [Adım Adım Sorun Giderme Senaryoları](#-adım-adım-sorun-giderme-senaryoları)
   - ["User location is not supported" (403 Forbidden) Hatasını Düzeltme](#senaryo-1-bölgesel-kısıtlama-ve-403-hatasını-aşma)
   - [Sıkışan Kota ve Hız Limitlerini (429) Düzeltme](#senaryo-2-sıkışan-kota-ve-hız-limitlerini-sıfırlama-429)
   - [Kod Üretimi Sırasında Akış Donmalarını Önleme](#senaryo-3-akış-donmalarını-ve-bağlantı-kopmalarını-düzeltme)
   - [TUN Modu Olmadan Yerel Proxy'ye Bağlanma](#senaryo-4-sistem-geneli-tun-olmadan-yerel-proxyye-bağlanma)
5. [Yedekleme ve Güvenli Geri Yükleme](#-yedekleme-ve-güvenli-geri-yükleme)

---

## 🎯 Çözülen Temel Sorunlar

| Sorun / Hata | Teknik Neden | Antigravity Cleaner Çözümü |
| :--- | :--- | :--- |
| **`User location is not supported` (403)** | Google, `language_server` baytkodunda ülke uygunluğunu kısıtlar. | **`ag-cleaner patch`**: Baytkodu 1.9 GB/s hızla tarar ve coğrafi kontrolleri etkisiz hale getirir. |
| **`Quota exceeded / Rate limited` (429)** | Bozuk yerel kota sayaçları ve yetim kilit dosyaları sistemi limitli gösterir. | **`ag-cleaner clean`**: Proje kodlarına dokunmadan kota önbelleğini sıfırlar. |
| **Akışın %99'da Donması veya Kopması** | Ara güvenlik duvarları uzun süren LLM üretimlerinde boşta kalan bağlantıları keser. | **`ag-cleaner proxy`**: TCP KeepAlive sinyalleri ile bağlantıyı sürekli canlı tutar. |
| **Tarayıcıda Çalışan Proxy'nin IDE'de Çalışmaması** | Antigravity, sistem proxy ayarlarını yok sayan bağımsız bir arka plan sürecidir. | **`ag-cleaner proxy`**: Ağ soketlerini doğrudan yerel proxy bağlantı noktalarına (10808 / 7890) yönlendirir. |

---

## 🖥️ Arayüz Seçenekleri

### 1. 1-Tıkla Retro KeyGen Arayüzü (GUI)
* **Kimin İçin:** Terminal komutlarıyla uğraşmak istemeyen kullanıcılar.
* **Nasıl Çalıştırılır:** İndirilen yürütülebilir dosyaya çift tıklayın (Windows'ta `.exe`).
* **Özellikler:** Windows 95 retro arayüzü, canlı yıldız alanı efekti, retro chiptune sesleri ve tek tıkla **"Auto-Fix & Patch"** butonu.

### 2. Modern Etkileşimli Terminal Arayüzü (TUI)
* **Kimin İçin:** Terminal içi zengin gösterge panelini tercih eden geliştiriciler.
* **Nasıl Çalıştırılır:**
  ```bash
  ag-cleaner tui
  ```
* **Kısayollar:** Paneller arası geçiş için `Tab`, işlem çalıştırmak için `Enter`, çıkış için `q`.

### 3. CLI Komut Satırı
* **Kimin İçin:** Gelişmiş kullanıcılar, CI/CD ve otomasyon senaryoları.

---

## 🛠️ CLI Komut Başvurusu

### `ag-cleaner doctor`
Antigravity ortamınız için kapsamlı bir sağlık ve teşhis taraması yapar.
```bash
ag-cleaner doctor
```

### `ag-cleaner clean`
Bozuk önbellekleri ve yetim kilit dosyalarını cerrahi olarak temizler.
```bash
ag-cleaner clean
```
> [!IMPORTANT]
> `ag-cleaner clean` **ASLA** proje kodlarınızı, git depolarınızı, eklentilerinizi veya ayarlarınızı silmez.

### `ag-cleaner patch`
Bölgesel yaptırımları aşmak için `language_server` dosyasına doğrudan baytkod yaması uygular.
```bash
ag-cleaner patch
```
Orijinal dosyayı geri yüklemek için:
```bash
ag-cleaner patch --restore
```

### `ag-cleaner proxy`
Sistem genelinde TUN modunu açmadan akıllı yerel proxy yönlendirmesi sağlar.
```bash
# Otomatik tespit (10808, 7890, 1080, 2080)
ag-cleaner proxy

# Özel SOCKS5 portu belirtme
ag-cleaner proxy --socks5 127.0.0.1:10808
```

### `ag-cleaner run`
Antigravity IDE'yi doğrulanmış ağ ortam değişkenleriyle sarmalayarak başlatır.
```bash
ag-cleaner run
```

---

## 🚀 Adım Adım Sorun Giderme Senaryoları

### Senaryo 1: Bölgesel Kısıtlama ve 403 Hatasını Aşma
1. Antigravity IDE'yi tamamen kapatın.
2. Yamayı uygulayın:
   ```bash
   ag-cleaner patch
   ```
3. Durumu kontrol edin:
   ```bash
   ag-cleaner doctor
   ```
4. Proxy/VPN istemcinizi açın ve IDE'yi başlatın.

---

### Senaryo 2: Sıkışan Kota ve Hız Limitlerini Sıfırlama (429)
1. IDE'yi kapatın.
2. Cerrahi önbellek temizliğini çalıştırın:
   ```bash
   ag-cleaner clean
   ```
3. IDE'yi tekrar başlatın.

---

### Senaryo 3: Akış Donmalarını ve Bağlantı Kopmalarını Düzeltme
1. Proxy istemcinizin kararlı olduğundan emin olun.
2. IDE'yi KeepAlive desteğiyle başlatın:
   ```bash
   ag-cleaner run
   ```

---

### Senaryo 4: Sistem Geneli TUN Olmadan Yerel Proxy'ye Bağlanma
1. Proxy istemcinizi (Clash: 7890, Xray: 10808) açık tutun.
2. Şu komutu girin:
   ```bash
   ag-cleaner proxy --socks5 127.0.0.1:10808
   ```

---

## 🛡️ Yedekleme ve Güvenli Geri Yükleme
Her `patch` işleminde orijinal dosyanın tarihli tam bir yedeği alınır (`<dosya>.bak.<zaman_damgası>`).
Geri yüklemek için:
```bash
ag-cleaner patch --restore
```

---

## 📜 Lisans
Bu proje uluslararası standart **Apache License 2.0** ile lisanslanmıştır.
