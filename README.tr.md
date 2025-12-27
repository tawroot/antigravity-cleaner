# Antigravity Cleaner Shell (v3.0.0)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Platform](https://img.shields.io/badge/Platform-Windows-blue)](https://microsoft.com/windows)
[![Version](https://img.shields.io/badge/Version-3.0.0-green)]()

> **İran halkına ve dijital yaptırımlar ile internet kısıtlamaları altında sıkışıp kalan herkese ithaf edilmiştir.**
> Bilgiye, teknolojiye ve geliştirme araçlarına özgür erişimin her insanın en temel hakkı olduğuna inanıyoruz. Bu araç, dijital özgürlüğün sesidir; hem içeriden hem de dışarıdan etrafımıza örülen duvarları aşmak için tasarlandı.

---

## İçindekiler
1. [Proje Hakkında](#proje-hakkında)
2. [Kolay Kurulum (Tek Satır)](#kolay-kurulum-tek-satır)
3. [Temel Özellikler](#temel-özellikler)
4. [Güvenlik ve Gizlilik](#güvenlik-ve-gizlilik)
5. [Modül Kılavuzu](#modül-kılavuzu)
6. [Yol Haritası (Roadmap)](#yol-haritası-roadmap)
7. [Lisans](#lisans)
8. [Yazar Hakkında](#yazar-hakkında)

---

## Proje Hakkında
**Antigravity Cleaner Shell**, geliştiricilere, freelancerlara ve ileri düzey kullanıcılara yardımcı olmak amacıyla tasarlanmış, **PowerShell** tabanlı, açık kaynaklı (Open Source) ve güçlü bir araçtır. Bu aracın temel amacı, tarayıcı oturumlarını akıllıca yönetmek, Google yaptırımlarını (Google Sanctions) aşmak ve yazılım geliştirme süreçleri için sistemi optimize etmektir.

Bu sürüm (v3.0.0), Python veya ağır kütüphanelere ihtiyaç duymadan, tamamen Windows üzerinde yerel (Native) olarak çalışacak şekilde yeniden yazılmıştır.

---

## Kolay Kurulum (Tek Satır)

Kurulum ve çalıştırma için PowerShell'i açın ve sadece aşağıdaki satırı kopyalayıp Enter tuşuna basın.
(Bu komut programı otomatik olarak indirir ve masaüstünüzde bir kısayol oluşturur)

```powershell
iwr https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/antigravity-cleaner/install.ps1 -useb | iex
```

*Hata alırsanız, terminali Yönetici (Administrator) olarak çalıştırın.*

---

## Temel Özellikler

### 1. Gelişmiş Oturum Yöneticisi (Session Manager)
Antigravity'nin kalbi olan bu modül, tarayıcı profillerinizi yönetmenizi ve taşımanızı sağlar.
*   **Akıllı Algılama:** Chrome, Edge, Brave ve Opera tarayıcı profillerini ve bağlı e-posta adreslerini otomatik olarak algılar.
*   **Çift Modlu Yedekleme (Dual Mode Backup):**
    *   **Light Mod:** Sadece çerezleri, giriş bilgilerini ve temel ayarları yedekler (~20MB). Hızlı taşıma ve şifre girmeden oturum açmak için idealdir.
    *   **Full Mod:** Tüm uzantılar, geçmiş ve önbellek dahil olmak üzere tam profil yedeği alır (~500MB+).
*   **Antigravity Desktop Desteği:** Antigravity IDE (VS Code tabanlı) masaüstü sürümü için özel yedekleme.

### 2. Bölge Denetçisi (Region Inspector)
Google hesaplarının yaptırıma uğramasından endişe edenler için hayati bir araç.
*   **Durum Kontrolü:** Doğru profille doğrudan Google'ın gizli `Country Association` sayfasına yönlendirir.
*   **Sızıntı Önleme Kontrolü (Pre-Check):** Bölge değişikliği talebinde bulunmadan önce IP Sızıntısı, DNS Sızıntısı ve WebRTC testleri yaparak talebinizin Google tarafından kabul edilme şansını 10 kat artırır.

### 3. Ağ Analizörü (Systems Analysis)
*   **Bağlantı Testi:** Google servislerine (Google Developer Services, Gemini AI, Cloud Platform) bağlantı durumunu anlık olarak kontrol eder.
*   **Bağımlılık Testi:** Geliştiriciler için kritik depolara (GitHub API, VS Code Marketplace) erişimi kontrol eder.

### 4. Sistem Optimizasyonu (System Cleaner)
*   Önceki konumunuza dair izler taşıyabilecek geçici dosyaları (Temp) ve sistem önbelleklerini temizler.
*   Diskte yer açmak ve IDE hızını artırmak için geliştirme araçlarının (JetBrains, VSCode) büyük önbelleklerini temizler.

### 5. Ağ Onarım Aracı (Network Reset)
*   Tek tıkla DNS önbelleğini temizler, Winsock ve TCP/IP ayarlarını sıfırlayarak bağlantı sorunlarını çözer.

---

## Güvenlik ve Gizlilik
Güvenliğinizi ciddiye alıyoruz.
*   **%100 Çevrimdışı:** Bu betik, "veri göndermek" için internete erişmez. Tüm işlemler kendi sisteminizde gerçekleştirilir.
*   **Telemetri Yok:** Verilerinizi toplamakla ilgilenmiyoruz.
*   **Şifreleme:** Oturum yedek dosyaları, Windows DPAPI standardı ile şifrelenmiş hassas bilgiler içerir ve (Tam Profil Yedeği hariç) sadece sizin sisteminizde kullanılabilir.

---

## Modül Kılavuzu

### Profil Yedekleme
1. Ana menüde `2` (Session Manager) seçeneğini seçin.
2. `1` (Backup Browser Profile) seçeneğine tıklayın.
3. Listeden (e-postaları gösteren) istediğiniz profili seçin.
4. Hızlı yedekleme (sadece girişler) için **Light Mod**, tam taşıma için **Full Mod** seçin.

### Geri Yükleme (Restore)
1. Session Manager'da `3` seçeneğine tıklayın.
2. Yedekleme listesini göreceksiniz (Tarih ve Light/Full türü ile).
3. Bir yedek seçtiğinizde, program tarayıcıyı otomatik olarak kapatır ve dosyaları değiştirir. **Uyarı:** O profilin mevcut verileri silinecektir.

### Google Bölge Değişikliği
1. Ana menüde `5` (Region Inspector) seçeneğine tıklayın.
2. VPN'inizin düzgün çalıştığından emin olmak için önce `Pre-Check` yapmanız önerilir.
3. Tarayıcının açılıp kayıtlı ülkeyi göstermesi için bir profil seçin.

---

## Yol Haritası (Roadmap)
Sürekli yeni özellikler ekliyoruz. Gelecek planlarımız:
- [ ] **v3.1:** Firefox desteği.
- [ ] **v3.2:** Otomatik DNS değiştirme aracı (Shecan, 403, Cloudflare).
- [ ] **v3.5:** Terminal kullanmak istemeyenler için modern Grafik Arayüz (GUI).

---

## Lisans
Bu proje **Tawana Anti-Copy License (TACL)** altında yayınlanmıştır. Kopyalanması ve ticari kullanımı yasaktır. Detaylar için `docs/LICENSE.md` dosyasına bakınız.

---

## Yazar Hakkında
**Tawana Network** tarafından geliştirilmiştir.
*Başkalarının duvar ördüğü yerlere biz köprüler kuruyoruz.*
