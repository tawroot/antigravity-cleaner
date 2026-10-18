// ========================================
// Antigravity Cleaner - Core Application Script
// Multilingual Localization & Theme Engine
// ========================================

// Language Content
const translations = {
    "en": {
        "home": "Home",
        "features": "Features",
        "download": "Download",
        "pricing": "Pricing",
        "blog": "Blog",
        "faq": "FAQ",
        "get_app": "Get App",
        "hero_title": "Fix Antigravity IDE Instantly",
        "hero_subtitle": "The most powerful maintenance tool for installation errors, network interference, and region restrictions.",
        "get_started": "Get Started",
        "view_source": "View Source",
        "stars": "GitHub Stars",
        "forks": "Forks",
        "downloads": "Downloads",
        "version": "Version",
        "quick_install": "Quick Install",
        "quick_install_desc": "Install instantly via terminal without leaving your keyboard.",
        "footer_desc": "Supporting open internet and developer productivity.",
        "product": "Product",
        "support": "Support",
        "docs": "Docs",
        "feat_clean_title": "One-Click Clean",
        "feat_clean_desc": "Remove corrupted installation files, cache, and registry keys instantly.",
        "feat_region_title": "Region Bypass",
        "feat_region_desc": "Automatically fix region-locked errors for constrained networks.",
        "feat_session_title": "Session Preservation",
        "feat_session_desc": "Clean the system without losing your browser cookies or login sessions.",
        "feat_network_title": "Network Reset",
        "feat_network_desc": "Flush DNS and reset network adapters to solve connection timeouts.",
        "feat_dryrun_title": "Dry Run Mode",
        "feat_dryrun_desc": "Preview exact changes before applying them to your system.",
        "feat_portable_title": "Portable",
        "feat_portable_desc": "No installation required. Run directly from USB or any folder.",
        "faq_safe_q": "Is it safe to use?",
        "faq_safe_a": "Yes. Antigravity Cleaner specifically targets temporary files, caches, and invalid registry keys associated with known IDE errors. It does not touch your personal documents or project source codes.",
        "faq_pass_q": "Will I lose my browser passwords?",
        "faq_pass_a": "No. Our 'Session Preservation' technology ensures that critical files like Cookies and Login Data are backed up or excluded from the cleaning process, so you don't have to log in again.",
        "faq_net_q": "How do I fix 'Network Error'?",
        "faq_net_a": "Use the 'Network Reset' feature in the Tools menu. This will flush your DNS cache and reset the Winsock catalog, which often resolves connection issues caused by VPNs using the Tun/Tap adapter.",
        "faq_mac_q": "Do you support Mac M1/M2?",
        "faq_mac_a": "Yes, we have native builds for Apple Silicon (ARM64) as well as Intel Macs. Check the Download page for the correct version.",
        "price_free_title": "Community",
        "price_free_1": "Open Source Core",
        "price_free_2": "GitHub Issues Support",
        "price_free_3": "All Basic Features",
        "download_free": "Download Free",
        "price_popular": "POPULAR",
        "price_pro_title": "Pro Network",
        "price_pro_1": "Private V2Ray Servers",
        "price_pro_2": "Low Latency Routing",
        "price_pro_3": "Priority Support",
        "price_ent_title": "Enterprise",
        "price_custom": "Custom",
        "price_ent_1": "Dedicated IPs",
        "price_ent_2": "Fleet Management",
        "price_ent_3": "24/7 SLA Support",
        "contact_sales": "Contact Sales",
        "blog_1_title": "Fixing SSL Handshake Errors",
        "blog_1_desc": "A comprehensive guide to understanding and fixing SSL/TLS errors in restricted network environments.",
        "read_more": "Read More",
        "blog_2_title": "Version 4.0 Release Notes",
        "blog_2_desc": "Celebrating the release of our biggest update yet, featuring new UI and faster cleaning engines.",
        "blog_3_title": "Understanding Region Locks",
        "blog_3_desc": "How software limits access based on geography and what you can do about it legally.",
        "community_title": "Join Our Community",
        "community_subtitle": "Get help, share ideas, and stay updated with the latest news.",
        "star_title": "Star on GitHub",
        "star_desc": "Show your support by starring our repository. It helps us grow!",
        "star_btn": "Star Repository",
        "telegram_title": "Telegram Channel",
        "telegram_desc": "Join our Telegram channel for updates, tips, and community support.",
        "telegram_btn": "Join Channel",
        "donate_title": "Support Development",
        "donate_desc": "Help us maintain and improve this free tool with a donation.",
        "donate_btn": "Donate",
        "why_choose": "Why Choose Antigravity Cleaner?",
        "why_choose_desc": "Powerful features designed for developers and power users.",
        "view_all_features": "View All Features",
        "cta_title": "Ready to Fix Your IDE?",
        "cta_desc": "Download now and solve installation errors in seconds.",
        "download_now": "Download Now",
        "view_on_github": "View on GitHub",
        "freedom_badge": "🌍 UNIVERSAL AI FREEDOM & REGION BYPASS",
        "freedom_title": "Are you locked out of Antigravity IDE & AI Tools?",
        "freedom_desc": "Geographical 403 blocks and unfair digital sanctions shouldn't hold back passionate developers. Antigravity Cleaner cleans regional lock tokens, bypasses network restrictions, and restores your sovereign right to code.",
        "status_bypassed": "Bypassed",
        "freedom_quote": "Dedicated to every developer filtered from within and sanctioned from without. Code knows no borders.",
        "gui_badge": "🎨 RETRO KEYGEN DESKTOP INTERFACE",
        "gui_title": "Powerful Diagnostics, Timeless Retro Aesthetic",
        "gui_desc": "A fully offline standalone GUI built for developers who appreciate classic keygen design, chiptune audio, and one-click diagnostics.",
        "gui_f1": "1-Click Automatic Repair & Bypass",
        "gui_f2": "100% Offline & Telemetry-Free",
        "gui_f3": "Built-in 8-Bit Chiptune Synthesizer",
        "gui_f4": "Session & Credential Preservation",
        "coming_soon": "Coming Soon",
        "how_it_works_badge": "⚡ 3-STEP SURGICAL REPAIR & HEALING",
        "how_it_works_title": "How It Works — Fix Your IDE in Under 30 Seconds",
        "how_it_works_subtitle": "Systematic self-healing architecture: from deep cache de-bloat and Go core patching to lightweight proxy injection.",
        "step1_title": "01. Surgical Cache De-bloat & Reset",
        "step1_desc": "Clears bloated GPU/Dawn caches, orphaned SingletonLocks, and artificial 429 quota tokens that cause memory stutter. Google account login, session tokens, and cookies remain 100% untouched.",
        "step2_title": "02. Core Go Binary & JS Engine Patching",
        "step2_desc": "Directly patches the native Go backend daemon (agy), Electron main.js, and VS Code extension to neutralize regional locks and terminate background telemetry loops that silently hog bandwidth.",
        "step3_title": "03. Lightweight Layer-7 Routing & KeepAlive",
        "step3_desc": "Eliminates the heavy latency and packet-loss overhead of system-wide TUN adapters. Injects direct SOCKS5 with 15s TCP KeepAlive so AI reasoning streams run buttery smooth and responsive even on weak or throttled internet.",
        "perf_banner_title": "Noticeably Lighter & Smoother on Congested Networks",
        "perf_banner_desc": "By avoiding the kernel virtualization overhead of global TUN adapters, stopping aggressive telemetry pings, and keeping HTTP/2 connections alive with 15s heartbeats, Antigravity IDE consumes significantly fewer system resources and runs smoothly even on high-latency internet connections.",
        "offline_privacy": "Offline & Private",
        "troubleshooting": "Troubleshooting",
        "tb_hero_badge": "OFFICIAL TROUBLESHOOTING KNOWLEDGE BASE",
        "tb_hero_title": "Fix Google Antigravity Errors & Region Lock",
        "tb_hero_desc": "Complete step-by-step solutions for Error 403 Forbidden, 429 Quota Exhaustion, Middlebox TLS Resets, and No-TUN Smart Proxy setup.",
        "tb_cta_download": "Get Antigravity Cleaner (Free & Offline)",
        "tb_cta_star": "Star on GitHub",
        "tb_quickfix_title": "1-Click Universal Auto-Fix (Recommended)",
        "tb_quickfix_desc": "Instead of manually editing JavaScript bundles and configuring network routing, run Antigravity Cleaner to patch all 4 targets and generate an isolated proxy launcher automatically:",
        "tb_quickfix_win": "* On Windows: Open PowerShell and run <code>irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex</code>",
        "tb_err1_badge": "ERROR 403 FORBIDDEN • REGION LOCK",
        "tb_err1_title": "\"Service is not available in your country\" / Region Lockout",
        "tb_err1_cause": "<strong>Root Cause:</strong> Google DeepMind and Google Cloud enforce strict geographic restrictions on Google Antigravity in countries such as Iran, Russia, China, Cuba, and Syria. The IDE runs internal JavaScript gatekeeper functions that inspect telemetry and token responses. If blocked, Antigravity forces a session invalidation and kicks you out to the login screen.",
        "tb_err1_sol_title": "Fixing Error 403 Step-by-Step:",
        "tb_err1_sol_1": "<strong>Neutralize Client-Side Checks:</strong> Run <code>antigravity-cleaner patch</code>. This surgically unlocks <code>main.js</code>, the IDE binary, CLI (<code>agy</code>), and VS Code extension so the software never locks itself.",
        "tb_err1_sol_2": "<strong>Launch with Smart Proxy (No TUN):</strong> Run <code>antigravity-cleaner launch</code>. This passes <code>--proxy-server=socks5://127.0.0.1:10808</code> directly to the Chromium engine. Chromium will resolve Google auth endpoints through your foreign proxy IP, keeping your session intact.",
        "tb_err1_sol_3": "<strong>Avoid TUN Mode:</strong> Never use global TUN mode for Antigravity. TUN mode causes MTU packet fragmentation and triggers Middlebox DPI dropouts.",
        "tb_err2_badge": "ERROR 429 • RATE LIMIT / QUOTA EXHAUSTED",
        "tb_err2_title": "\"Too Many Requests\" & Corrupted Token Cache",
        "tb_err2_cause": "<strong>Root Cause:</strong> Autonomous agents generate dozens of tool calls, search queries, and subagent transcripts per session. Local token buckets in <code>~/.config/Antigravity/Cache</code> and <code>DawnWebGPUCache</code> get out of sync with Google's servers, causing perpetual 429 errors even when you have quota left.",
        "tb_err2_sol_title": "How to Resolve Error 429 without Losing Data:",
        "tb_err2_sol_1": "Click the <strong>[Reset 429]</strong> button in the Antigravity Cleaner GUI or run:",
        "tb_err2_sol_2": "<strong>Session Preservation Guarantee:</strong> Our surgical reset purges corrupted quota blobs and telemetry caches while strictly preserving your <code>Cookies</code>, <code>Preferences</code>, <code>app_storage.json</code>, and chat history. You do NOT have to log in again.",
        "tb_err3_badge": "NETWORK ARCHITECTURE • NO-TUN MODE",
        "tb_err3_title": "Smart Proxy Injection for v2ray, Xray, Clash, and NekoBox",
        "tb_err3_cause": "<strong>Why traditional VPNs fail:</strong> When you enable TUN mode, your entire operating system's traffic is redirected through the VPN. This increases latency, causes packet drops, and breaks local intranet and domestic banking websites.",
        "tb_err3_sol_title": "How Antigravity Cleaner's No-TUN Injection Works:",
        "tb_err3_sol_1": "The Cleaner automatically scans local loopback ports:",
        "tb_err3_sol_2": "It launches Antigravity with a dedicated pipeline: <code>--proxy-server=socks5://127.0.0.1:&lt;detected_port&gt;</code>.",
        "tb_err3_sol_3": "Only Google Antigravity traffic is routed through the proxy. Your web browser, terminal downloads, and domestic sites continue running at maximum local speed!",
        "tb_err3_note": "This generates a permanent 1-click desktop launcher (<code>Antigravity.desktop</code> on Linux, <code>Antigravity.lnk</code> on Windows, <code>Antigravity-Proxy.command</code> on macOS).",
        "tb_err4_badge": "CONNECTION DROPOUT • MIDDLEBOX DPI",
        "tb_err4_title": "Middlebox TLS Disconnects & Stream Interruptions",
        "tb_err4_cause": "<strong>The Issue:</strong> During deep thinking or multi-minute code planning, the HTTP/2 SSE connection between the IDE and Google's edge server (<code>daily-cloudcode-pa.googleapis.com</code>) remains silent. Strict ISP firewalls (Middleboxes) mistake idle TLS connections for dead tunnels and send TCP RST packets, causing \"Stream Closed / Broken Pipe\" errors.",
        "tb_err4_sol_title": "The Solution:",
        "tb_err4_sol_desc": "Antigravity Cleaner patches the socket connection layer to inject a <strong>15-second TCP KeepAlive heartbeat</strong>. This keeps the stateful connection alive across all firewalls and middleboxes without dropping your agentic sessions.",
        "tb_err5_badge": "STARTUP FAILURE • SINGLETON LOCK",
        "tb_err5_title": "Antigravity Won't Open After Crash or Sleep",
        "tb_err5_cause": "<strong>The Issue:</strong> If your system suspends or Antigravity crashes, Chromium leaves a lockfile at <code>~/.config/Antigravity/SingletonLock</code> pointing to a dead PID. When clicked, Antigravity thinks another instance is active and refuses to open.",
        "tb_err5_sol_desc": "Antigravity Cleaner's smart launcher audits and removes stale singleton locks automatically on every launch!",
        "tb_err6_badge": "STREAM CRASH • WRITE EIO / BROKEN PIPE",
        "tb_err6_title": "\"Agent terminated due to error\" & Broken Pipe Exceptions",
        "tb_err6_cause": "<strong>Root Cause:</strong> Sudden pipe breakdown caused by verbose raw logging (<code>write EIO</code>) or middlebox connection timeouts during streaming agent thought loops.",
        "tb_err6_sol_title": "Instant Fix with Antigravity Cleaner:",
        "tb_err6_sol_desc": "Click <strong>[ ① ⚡ Auto-Fix ]</strong> in the Retro GUI or run <code>antigravity-cleaner patch all</code>. Our Zero-Log ASAR Shield silences verbose console outputs, eliminates broken pipe exceptions, and enforces 15s TCP KeepAlive heartbeat tuning to keep streaming inference alive.",
        "tb_err7_badge": "CONNECTION DROP • DPI & TELEMETRY LEAK",
        "tb_err7_title": "\"Unknown: There was a network issue connecting to the server\"",
        "tb_err7_cause": "<strong>Root Cause:</strong> Middlebox/DPI firewalls terminating idle TLS/gRPC streams during prolonged reasoning phases, or background telemetry probes leaking real unmasked IP midway through a task.",
        "tb_err7_sol_desc": "Launch Antigravity via <strong>[ ② 🚀 Launch IDE ]</strong> or <code>antigravity-cleaner launch</code>. Traffic is isolated over a direct SOCKS5h loopback socket with the 8-flag Zero-Telemetry wrapper active.",
        "tb_err8_badge": "AUTH BLOCK • 403 THAT’S AN ERROR",
        "tb_err8_title": "\"403. That’s an error. /auth-success\"",
        "tb_err8_cause": "<strong>Root Cause:</strong> Google regional authentication gateway flagging the browser login redirect from sanctioned regions due to historical device fingerprints or unshielded telemetry beacons.",
        "tb_err8_sol_desc": "Run <strong>[ ① ⚡ Auto-Fix ]</strong>. Cleaner patches the core binary machine code (<code>hasValidAuth = true</code>), injects enterprise auth mode flags, and purges tainted authentication tokens.",
        "tb_err9_badge": "PROXY MISMATCH • DIAL TCP 127.0.0.1:10808",
        "tb_err9_title": "\"proxyconnect tcp: dial tcp 127.0.0.1:10808: connect: connection refused\"",
        "tb_err9_cause": "<strong>Root Cause:</strong> Local proxy client (Xray, Clash, NekoBox) is either closed or listening on an unexpected port.",
        "tb_err9_sol_desc": "Run <code>antigravity-cleaner doctor</code> (or click <strong>[ 🩺 Doctor ]</strong> in GUI). The diagnostics engine dynamically scans active ports and aligns your IDE proxy configuration automatically.",
        "tb_err10_badge": "STATE CORRUPTION • FAILED TO FETCH",
        "tb_err10_title": "\"There was an unexpected issue setting up your account. Failed to fetch\"",
        "tb_err10_cause": "<strong>Root Cause:</strong> Corrupted DIPS database records and orphaned WebSocket caches in application state.",
        "tb_err10_sol_desc": "Click <strong>[ 🧹 Reset 429 ]</strong> in the GUI or execute <code>antigravity-cleaner clean</code>. Purges 16 cached quota artifacts in 0.002s while 100% preserving your projects and chat history.",
        "footer_note": "Antigravity Cleaner is a community-driven open-source project released under the GNU General Public License v3.0.<br>Dedicated to developers facing geographic restrictions and digital isolation worldwide. 🕊️",
        "tb_err7_sol_title": "Instant Fix:",
        "tb_err8_sol_title": "Instant Fix:",
        "tb_err9_sol_title": "Instant Fix:",
        "tb_err10_sol_title": "Instant Fix:"
    },
    "fa": {
        "home": "خانه",
        "features": "ویژگی‌ها",
        "download": "دانلود",
        "pricing": "تعرفه‌ها",
        "blog": "وبلاگ",
        "faq": "سوالات",
        "get_app": "دانلود برنامه",
        "hero_title": "رفع مشکلات آنتی‌گرویتی، آنی و سریع",
        "hero_subtitle": "قدرتمندترین ابزار مدیریت برای رفع خطاهای نصب، تداخل شبکه و محدودیت‌های منطقه‌ای.",
        "get_started": "شروع کنید",
        "view_source": "کد منبع",
        "stars": "ستاره‌ها",
        "forks": "فورک‌ها",
        "downloads": "دانلودها",
        "version": "نسخه",
        "quick_install": "نصب سریع",
        "quick_install_desc": "نصب فوری از طریق ترمینال بدون نیاز به دانلود دستی فایل.",
        "footer_desc": "حمایت از اینترنت آزاد و بهره‌وری توسعه‌دهندگان.",
        "product": "محصول",
        "support": "پشتیبانی",
        "docs": "مستندات",
        "feat_clean_title": "پاکسازی تک‌کلیک",
        "feat_clean_desc": "حذف فوری فایل‌های نصبی خراب، کش‌ها و کلیدهای رجیستری معیوب سیستم.",
        "feat_region_title": "دور زدن تحریم",
        "feat_region_desc": "رفع خودکار خطاهای Region Lock و محدودیت‌های جغرافیایی برای شبکه‌های محدود.",
        "feat_session_title": "حفظ نشست‌ها",
        "feat_session_desc": "پاکسازی عمیق سیستم بدون از دست دادن کوکی‌ها، تاریخچه و لاگین‌های مرورگر.",
        "feat_network_title": "ریست شبکه",
        "feat_network_desc": "فلاش DNS و ریست کامل آداپتورها برای حل مشکل تایم‌آوت و تداخل VPN.",
        "feat_dryrun_title": "حالت تست",
        "feat_dryrun_desc": "مشاهده و بررسی دقیق تغییرات احتمالی قبل از اعمال نهایی روی سیستم.",
        "feat_portable_title": "نسخه پرتابل",
        "feat_portable_desc": "بدون نیاز به نصب. اجرا به صورت مستقیم از روی فلش USB یا هر پوشه‌ای.",
        "faq_safe_q": "آیا امن است؟",
        "faq_safe_a": "بله. آنتی‌گرویتی فقط فایل‌های موقت و کلیدهای خراب مربوط به خطاهای IDE را هدف می‌گیرد و کاری به فایل‌های شخصی شما ندارد.",
        "faq_pass_q": "آیا پسوردهای مرورگر پاک می‌شوند؟",
        "faq_pass_a": "خیر. تکنولوژی «حفظ نشست» ما فایل‌های حساس مثل کوکی‌ها و دیتای لاگین را بک‌آپ گرفته یا مستثنی می‌کند.",
        "faq_net_q": "چطور خطای Network Error را حل کنم؟",
        "faq_net_a": "از گزینه Network Reset در منوی Tools استفاده کنید. این کار DNS را فلاش کرده و کاتالوگ Winsock را ریست می‌کند.",
        "faq_mac_q": "آیا مک M1/M2 پشتیبانی می‌شود؟",
        "faq_mac_a": "بله، ما بیلد نیتیو (Native) برای پردازنده‌های اپل سیلیکون (ARM64) و همچنین اینتل داریم.",
        "price_free_title": "جامعه",
        "price_free_1": "هسته متن‌باز",
        "price_free_2": "پشتیبانی در گیت‌هاب",
        "price_free_3": "تمام امکانات پایه",
        "download_free": "دانلود رایگان",
        "price_popular": "محبوب",
        "price_pro_title": "شبکه حرفه‌ای",
        "price_pro_1": "سرورهای اختصاصی V2Ray",
        "price_pro_2": "روتینگ با تاخیر کم",
        "price_pro_3": "پشتیبانی اولویت‌دار",
        "price_ent_title": "سازمانی",
        "price_custom": "توافقی",
        "price_ent_1": "آی‌پی اختصاصی",
        "price_ent_2": "مدیریت گروهی کاربران",
        "price_ent_3": "پشتیبانی ۲۴/۷ SLA",
        "contact_sales": "تماس با فروش",
        "blog_1_title": "حل خطاهای SSL Handshake",
        "blog_1_desc": "راهنمای جامع برای رفع خطاهای SSL/TLS در محیط‌های شبکه محدود.",
        "read_more": "ادامه مطلب",
        "blog_2_title": "یادداشت‌های انتشار نسخه ۴.۰",
        "blog_2_desc": "انتشار بزرگترین آپدیت ما با رابط کاربری جدید و موتور پاکسازی سریع‌تر.",
        "blog_3_title": "درک قفل‌های منطقه‌ای",
        "blog_3_desc": "چگونه نرم‌افزارها دسترسی را محدود می‌کنند و راه حل‌ها چیست.",
        "community_title": "به جامعه ما بپیوندید",
        "community_subtitle": "کمک بگیرید، ایده‌ها را به اشتراک بگذارید و از آخرین اخبار مطلع شوید.",
        "star_title": "ستاره در گیت‌هاب",
        "star_desc": "با ستاره دادن به مخزن ما، حمایت خود را نشان دهید. این به رشد ما کمک می‌کند!",
        "star_btn": "ستاره دادن",
        "telegram_title": "کانال تلگرام",
        "telegram_desc": "برای دریافت به‌روزرسانی‌ها، نکات و پشتیبانی جامعه به کانال تلگرام ما بپیوندید.",
        "telegram_btn": "عضویت در کانال",
        "donate_title": "حمایت از توسعه",
        "donate_desc": "با کمک مالی به ما در نگهداری و بهبود این ابزار رایگان کمک کنید.",
        "donate_btn": "حمایت مالی",
        "why_choose": "چرا آنتی‌گرویتی کلینر؟",
        "why_choose_desc": "امکانات قدرتمند طراحی شده برای توسعه‌دهندگان و کاربران حرفه‌ای.",
        "view_all_features": "مشاهده تمام ویژگی‌ها",
        "cta_title": "آماده رفع مشکلات IDE خود هستید؟",
        "cta_desc": "همین حالا دانلود کنید و خطاهای نصب را در چند ثانیه حل کنید.",
        "download_now": "دانلود فوری",
        "view_on_github": "مشاهده در گیت‌هاب",
        "freedom_badge": "🌍 آزادی جهانی هوش مصنوعی و رفع تحریم‌های منطقه‌ای",
        "freedom_title": "آیا دسترسی شما به آنتی‌گرویتی و ابزارهای هوش مصنوعی تحریم شده است؟",
        "freedom_desc": "خطاهای ۴۰۳ جغرافیایی و تحریم‌های ناعادلانه نباید مانع پیشرفت برنامه‌نویسان شود. آنتی‌گرویتی کلینر توکن‌های قفل منطقه‌ای را پاکسازی کرده، تداخل شبکه را رفع می‌کند و دسترسی کامل را به شما بازمی‌گرداند.",
        "status_bypassed": "رفع تحریم",
        "freedom_quote": "تقدیم به تمام برنامه‌نویسانی که از داخل فیلتر و از بیرون تحریم شده‌اند. دانش و کد مرز نمی‌شناسد.",
        "gui_badge": "🎨 رابط گرافیکی نوستالژیک رترو (KeyGen Style)",
        "gui_title": "عیب‌یابی قدرتمند در کنار حس ناب و نوستالژیک",
        "gui_desc": "یک نرم‌افزار گرافیکی سبک و کاملاً آفلاین با الهام از کی‌جن‌های کلاسیک، همراه با موزیک چیپ‌تیون ۸-بیتی و امکان پاکسازی تک‌کلیک.",
        "gui_f1": "پاکسازی خودکار و رفع قفل با یک کلیک",
        "gui_f2": "۱۰۰٪ آفلاین و بدون هیچ ارسال داده یا ردگیری",
        "gui_f3": "پخش‌کننده داخلی موزیک خاطره‌انگیز ۸-بیتی",
        "gui_f4": "حفظ کامل سشن‌ها، کوکی‌ها و پروژه‌های شما",
        "coming_soon": "به‌زودی",
        "how_it_works_badge": "⚡ فرآیند ۳ مرحله‌ای بازیابی و رفع خطا",
        "how_it_works_title": "رفع اختلالات Antigravity در کمتر از ۳۰ ثانیه",
        "how_it_works_subtitle": "معماری خودترمیم‌شونده: از پاکسازی کش و رفع قفل‌ها تا خنثی‌سازی تحریم و اتصال مستقیم به پروکسی.",
        "step1_title": "۱. پاکسازی کش‌های معیوب و قفل‌های نشست (حفظ حساب کاربری)",
        "step1_desc": "حذف کش‌های سنگین GPU، رفع فایل‌های قفل‌شده سیستم (SingletonLock) و صفر کردن شمارنده خطای ۴۲۹؛ اطلاعات ورود به حساب گوگل و فایل‌های پروژه بدون تغییر باقی می‌مانند.",
        "step2_title": "۲. خنثی‌سازی محدودیت‌های منطقه‌ای در هسته برنامه",
        "step2_desc": "غیرفعال‌سازی بررسی‌های مسدودکننده جغرافیایی در ماژول‌های کلاینت و هسته باینری ادیتور، همراه با جلوگیری از هدررفت ترافیک اینترنت در پس‌زمینه.",
        "step3_title": "۳. اتصال پایدار به پروکسی بدون سربار شبکه (ضد قطعی استریم)",
        "step3_desc": "اتصال مستقیم فرآیند ادیتور به پروکسی محلی بدون نیاز به فعال‌سازی حالت سنگین TUN در کل سیستم؛ به همراه ارسال سیگنال زنده (KeepAlive) برای تداوم نگارش کد توسط مدل‌های هوش مصنوعی.",
        "perf_banner_title": "چرا محیط توسعه به‌طور ملموس سبک‌تر و روان‌تر اجرا می‌شود؟",
        "perf_banner_desc": "به دلیل عدم درگیر شدن کارت شبکه با آداپتور مجازی TUN، قطع حلقه‌های تکراری تله‌متری در پس‌زمینه و حفظ اتصال پایدار با سرور از طریق هارت‌بیت، ادیتور منابع بسیار کمتری از سخت‌افزار و پهنای باند را مصرف می‌کند.",
        "offline_privacy": "۱۰۰٪ آفلاین و خصوصی",
        "troubleshooting": "عیب‌یابی",
        "tb_hero_badge": "مرجع رسمی مستندات و راهنمای حل خطاهای آنتی‌گرویتی",
        "tb_hero_title": "حل خطاهای Google Antigravity، قفل منطقه‌ای و تداخل شبکه",
        "tb_hero_desc": "راهنمای جامع گام‌به‌گام برای رفع خطای ۴۰۳ (تحریم)، اتمام ساختگی سهمیه ۴۲۹، قطعی استریم توسط فایروال و اتصال پروکسی بدون حالت TUN.",
        "tb_cta_download": "دریافت آنتی‌گرویتی کلینر (رایگان و آفلاین)",
        "tb_cta_star": "ستاره در گیت‌هاب",
        "tb_quickfix_title": "رفع خودکار و جامع تمام خطاها با ۱ کلیک (پیشنهادی)",
        "tb_quickfix_desc": "به جای ویرایش دستی فایل‌های پیچیده جاوااسکریپت و تنظیمات سخت شبکه، با اجرای آنتی‌گرویتی کلینر هر ۴ جزء اصلی را پچ کرده و میانبر لانچر پروکسی اختصاصی بسازید:",
        "tb_quickfix_win": "* در ویندوز: پاورشل (PowerShell) را باز کرده و این دستور را اجرا نمایید: <code>irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex</code>",
        "tb_err1_badge": "خطای ۴۰۳ • قفل منطقه‌ای و تحریم جغرافیایی",
        "tb_err1_title": "«سرویس در کشور شما در دسترس نیست» (قفل تحریم منطقه‌ای)",
        "tb_err1_cause": "<strong>علت اصلی:</strong> شرکت گوگل محدودیت‌های سرسختانه منطقه‌ای بر روی ادیتور Antigravity اعمال کرده است (به‌ویژه در ایران، روسیه و چین). ادیتور دارای توابع داخلی در جاوااسکریپت و هسته برنامه است که وضعیت جغرافیایی شما را ارزیابی می‌کنند. در صورت شناسایی، سشن باطل شده و کاربر به صفحه لاگین بازگردانده می‌شود.",
        "tb_err1_sol_title": "مراحل حل گام‌به‌گام خطای ۴۰۳:",
        "tb_err1_sol_1": "<strong>خنثی‌سازی بررسی‌های کلاینت:</strong> دستور <code>antigravity-cleaner patch</code> را اجرا کنید. این کار به صورت خودکار <code>main.js</code>، هسته باینری، ابزار CLI و اکستنشن را پچ کرده تا بررسی موقعیت مکانی بی‌اثر شود.",
        "tb_err1_sol_2": "<strong>اجرا با پروکسی هوشمند (بدون حالت TUN):</strong> دستور <code>antigravity-cleaner launch</code> را اجرا نمایید. این کار پارامتر اختصاصی پروکسی لوکال را به موتور نرم‌افزار تزریق می‌کند تا تبادلات هویتی بدون نشت آی‌پی و پایدار انجام شود.",
        "tb_err1_sol_3": "<strong>اجتناب از مود TUN:</strong> هرگز برای کار با این برنامه از حالت سراسری TUN فیلترشکن استفاده نکنید، زیرا موجب خرد شدن پکت‌ها و قطع ارتباط توسط فایروال می‌شود.",
        "tb_err2_badge": "خطای ۴۲۹ • اتمام سهمیه یا لیمیت کاذب درخواست‌ها",
        "tb_err2_title": "پیام «Too Many Requests» و کش معیوب توکن‌ها",
        "tb_err2_cause": "<strong>علت اصلی:</strong> دستیارهای هوشمند در هر جلسه تعداد زیادی فراخوانی ابزار و ترنسکریپت ثبت می‌کنند. کش‌های محلی دیتابیس توکن در پوشه کانفیگ با سرورهای گوگل ناهماهنگ شده و حتی با وجود داشتن سهمیه خالی، خطای کاذب ۴۲۹ می‌دهند.",
        "tb_err2_sol_title": "روش حل خطای ۴۲۹ بدون از دست رفتن اطلاعات شخصی:",
        "tb_err2_sol_1": "روی دکمه <strong>[Reset 429]</strong> در برنامه گرافیکی کلیک کنید یا این دستور را اجرا نمایید:",
        "tb_err2_sol_2": "<strong>تضمین حفظ نشست:</strong> این پاکسازی صرفاً کش‌های معیوب و داده‌های خراب سهمیه را تخلیه کرده و کوکی‌ها، لاگین اکانت گوگل و تاریخچه پروژه‌ها را به طور ۱۰۰٪ حفظ می‌کند (نیازی به لاگین مجدد نیست).",
        "tb_err3_badge": "معماری شبکه • اتصال بدون حالت TUN",
        "tb_err3_title": "تزریق هوشمند پروکسی برای v2ray، Xray، Clash و NekoBox",
        "tb_err3_cause": "<strong>چرا وی‌پی‌ان‌های معمولی با TUN با شکست مواجه می‌شوند:</strong> در حالت TUN کل ترافیک سیستم‌عامل درگیر تونل شده، پینگ به شدت بالا می‌رود و سایت‌های داخلی و بانکی قطع می‌شوند.",
        "tb_err3_sol_title": "مکانیزم تزریق پروکسی بدون TUN در آنتی‌گرویتی کلینر:",
        "tb_err3_sol_1": "نرم‌افزار به صورت خودکار پورت‌های فعال محلی پروکسی شما را شناسایی می‌کند:",
        "tb_err3_sol_2": "ادیتور مستقیماً با پایپ‌لاین ایزوله اجرا می‌شود: <code>--proxy-server=socks5://127.0.0.1:&lt;detected_port&gt;</code>.",
        "tb_err3_sol_3": "فقط و فقط ترافیک آنتی‌گرویتی از پروکسی عبور می‌کند؛ مرورگر، دانلودها و سایر ابزارهای شما با بالاترین سرعت مستقیم به کار خود ادامه می‌دهند!",
        "tb_err3_note": "با اجرای این دستور، یک آیکون شورتکات دائمی روی دسکتاپ ساخته می‌شود تا با یک دابل‌کلیک ادیتور با پروکسی متصل شود (<code>Antigravity.desktop</code> در لینوکس، <code>Antigravity.lnk</code> در ویندوز).",
        "tb_err4_badge": "قطعی استریم • فیلترینگ هوشمند DPI",
        "tb_err4_title": "قطعی‌های ناگهانی TLS و بسته شدن استریم کدنویسی هوش مصنوعی",
        "tb_err4_cause": "<strong>مشکل:</strong> هنگام پردازش‌های سنگین یا تفکر هوش مصنوعی، ارتباط HTTP/2 برای لحظاتی بدون ارسال داده ساکت می‌ماند. فایروال‌های میانی (Middlebox) این حالت را با اتصال قطع‌شده اشتباه گرفته و پکت ریست (TCP RST) می‌فرستند که موجب خطای Broken Pipe می‌شود.",
        "tb_err4_sol_title": "راه حل قطعی:",
        "tb_err4_sol_desc": "آنتی‌گرویتی کلینر یک سیگنال هارت‌بیت ۱۵ ثانیه‌ای (<strong>TCP KeepAlive</strong>) به هسته سوکت تزریق می‌کند تا تونل ارتباطی همیشه در نگاه فایروال‌ها زنده بماند و استریم هرگز قطع نشود.",
        "tb_err5_badge": "خطای باز نشدن برنامه • قفل فایل Singleton",
        "tb_err5_title": "باز نشدن برنامه پس از اسلیپ (Sleep) یا بسته شدن ناگهانی",
        "tb_err5_cause": "<strong>مشکل:</strong> اگر سیستم به خواب برود یا برنامه کرش کند، یک فایل قفل در مسیر پیکربندی به جا می‌ماند. ادیتور تصور می‌کند نمونه دیگری در حال اجراست و از بالا آمدن مجدد خودداری می‌کند.",
        "tb_err5_sol_desc": "لانچر هوشمند آنتی‌گرویتی کلینر در هر بار اجرا، فایل‌های قفل متروکه را فوراً شناسایی و پاک می‌کند تا نرم‌افزار بدون مشکل اجرا شود!",
        "tb_err6_badge": "کرش استریم • خطای Write EIO / Broken Pipe",
        "tb_err6_title": "پیام خطای «Agent terminated due to error» و خطاهای پایپ",
        "tb_err6_cause": "<strong>علت اصلی:</strong> لاگ‌گیری سنگین کنسول یا تایم‌اوت ارتباط با سرور در حین استریم مداوم، باعث شکستن پایپ داده و متوقف شدن دستیار کدنویسی می‌شود.",
        "tb_err6_sol_title": "رفع فوری با آنتی‌گرویتی کلینر:",
        "tb_err6_sol_desc": "روی دکمه <strong>[ ① ⚡ Auto-Fix ]</strong> در محیط کاربری کلیک کنید یا دستور <code>antigravity-cleaner patch all</code> را اجرا نمایید. شیلد بدون‌لاگ (Zero-Log ASAR) جلوی سربار کنسول را گرفته و با هارت‌بیت ۱۵ ثانیه‌ای، جریان استریم را کاملاً پایدار نگه می‌دارد.",
        "tb_err7_badge": "قطعی اتصال • نشت تله‌متری و مسدودسازی فایروال",
        "tb_err7_title": "خطای «Unknown: There was a network issue connecting to the server»",
        "tb_err7_cause": "<strong>علت اصلی:</strong> قطع کردن نشست‌های gRPC و استریم توسط فایروال یا نشت آی‌پی واقعی از طریق پکت‌های تله‌متری پس‌زمینه در حین کار با ادیتور.",
        "tb_err7_sol_desc": "برنامه را از طریق دکمه <strong>[ ② 🚀 Launch IDE ]</strong> یا دستور <code>antigravity-cleaner launch</code> اجرا کنید. ترافیک کاملاً روی سوکت امن SOCKS5 با مسدودسازی ۸ مرحله‌ای تله‌متری ایزوله می‌شود.",
        "tb_err8_badge": "مسدودی ورود • خطای 403 /auth-success",
        "tb_err8_title": "خطای «403. That’s an error. /auth-success» در صفحه مرورگر",
        "tb_err8_cause": "<strong>علت اصلی:</strong> گیت‌وی احراز هویت گوگل ریدایرکت مرورگر را به دلیل ردپای موقعیت جغرافیایی یا توکن‌های آلوده شناسایی کرده و اجازه لاگین نمی‌دهد.",
        "tb_err8_sol_desc": "گزینه <strong>[ ① ⚡ Auto-Fix ]</strong> را اجرا کنید. کلینر هسته باینری برنامه را پچ کرده، پرچم‌های حالت Enterprise Auth را تزریق نموده و توکن‌های آلوده را پاکسازی می‌کند.",
        "tb_err9_badge": "خطای عدم تطابق پورت پروکسی • پورت 10808",
        "tb_err9_title": "خطای «proxyconnect tcp: dial tcp 127.0.0.1:10808: connect: connection refused»",
        "tb_err9_cause": "<strong>علت اصلی:</strong> کلاینت پروکسی شما (مانند V2Ray، Clash، NekoBox یا Hiddify) بسته است یا روی پورت دیگری (مثلاً ۲۰۸۰ یا ۷۸۹۰) تنظیم شده است.",
        "tb_err9_sol_desc": "دستور <code>antigravity-cleaner doctor</code> را اجرا کنید (یا روی دکمه <strong>[ 🩺 Doctor ]</strong> کلیک کنید). موتور عیب‌یاب پورت‌های فعال را بررسی کرده و تنظیمات را هماهنگ می‌کند.",
        "tb_err10_badge": "خطای دیتابیس وضعیت • Failed to fetch",
        "tb_err10_title": "پیام خطای «There was an unexpected issue setting up your account. Failed to fetch»",
        "tb_err10_cause": "<strong>علت اصلی:</strong> رکوردهای خراب در دیتابیس داخلی DIPS و کش‌های به جا مانده از وب‌سوکت در وضعیت برنامه.",
        "tb_err10_sol_desc": "روی دکمه <strong>[ 🧹 Reset 429 ]</strong> کلیک کنید یا دستور <code>antigravity-cleaner clean</code> را اجرا کنید. ۱۶ دیتابیس موقت خراب در کسری از ثانیه پاک شده در حالی که پروژه‌ها و تاریخچه چت کاملاً دست‌نخورده باقی می‌ماند.",
        "footer_note": "آنتی‌گرویتی کلینر یک پروژه متن‌باز جامعه‌محور تحت مجوز GNU GPL v3.0 است.<br>تقدیم به تمام برنامه‌نویسانی که با محدودیت‌های منطقه‌ای و انزوای دیجیتال در سراسر جهان مواجه‌اند. 🕊️",
        "tb_err7_sol_title": "رفع فوری:",
        "tb_err8_sol_title": "رفع فوری:",
        "tb_err9_sol_title": "رفع فوری:",
        "tb_err10_sol_title": "رفع فوری:"
    },
    "ru": {
        "home": "Главная",
        "features": "Функции",
        "download": "Скачать",
        "pricing": "Цены",
        "blog": "Блог",
        "faq": "FAQ",
        "hero_title": "Исправьте ошибки Antigravity мгновенно",
        "hero_subtitle": "Самый мощный инструмент для обслуживания ошибок установки и сетевых помех.",
        "coming_soon": "Скоро",
        "offline_privacy": "Офлайн и приватно",
        "troubleshooting": "Устранение неполадок",
        "tb_hero_badge": "ОФИЦИАЛЬНАЯ БАЗА ЗНАНИЙ ПО УСТРАНЕНИЮ ОШИБОК",
        "tb_hero_title": "Исправление ошибок Google Antigravity и региональной блокировки",
        "tb_hero_desc": "Пошаговые решения для Ошибки 403 Forbidden, 429 Превышение квоты, сбросов Middlebox TLS и настройки умного прокси без TUN.",
        "tb_cta_download": "Скачать Antigravity Cleaner (Бесплатно и офлайн)",
        "tb_cta_star": "Звезда на GitHub",
        "tb_quickfix_title": "Универсальное автоисправление в 1 клик (Рекомендуется)",
        "tb_quickfix_desc": "Вместо ручного редактирования бандлов JavaScript и настройки маршрутизации, запустите Antigravity Cleaner для пропатчивания всех 4 целей и генерации изолированного лаунчера:",
        "tb_quickfix_win": "* В Windows: Откройте PowerShell и выполните <code>irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex</code>",
        "tb_err1_badge": "ОШИБКА 403 FORBIDDEN • РЕГИОНАЛЬНЫЙ БЛОК",
        "tb_err1_title": "\"Сервис недоступен в вашей стране\" / Блокировка региона",
        "tb_err1_cause": "<strong>Причина:</strong> Google применяет строгие географические ограничения на Google Antigravity в ряде стран (Иран, Россия, Китай и др.). IDE содержит внутренние функции проверки, которые анализируют телеметрию и ответы токенов, сбрасывая сессию при блокировке.",
        "tb_err1_sol_title": "Пошаговое исправление ошибки 403:",
        "tb_err1_sol_1": "<strong>Нейтрализация проверок:</strong> Выполните <code>antigravity-cleaner patch</code>. Это снимет ограничения в <code>main.js</code>, бинарнике IDE, CLI и расширении.",
        "tb_err1_sol_2": "<strong>Запуск через умный прокси (без TUN):</strong> Запустите <code>antigravity-cleaner launch</code>, чтобы направить запросы к серверам Google через локальный SOCKS5 прокси без включения TUN.",
        "tb_err1_sol_3": "<strong>Избегайте режима TUN:</strong> Глобальный TUN приводит к фрагментации MTU и обрывам соединений на DPI-фильтрах.",
        "tb_err2_badge": "ОШИБКА 429 • ПРЕВЫШЕНИЕ ЛИМИТА / КВОТА",
        "tb_err2_title": "\"Too Many Requests\" и поврежденный кэш токенов",
        "tb_err2_cause": "<strong>Причина:</strong> Агенты генерируют множество запросов. Локальные кэши токенов рассинхронизируются с серверами Google, вызывая постоянную ошибку 429 даже при наличии квоты.",
        "tb_err2_sol_title": "Как исправить ошибку 429 без потери данных:",
        "tb_err2_sol_1": "Нажмите кнопку <strong>[Reset 429]</strong> в интерфейсе Antigravity Cleaner или выполните:",
        "tb_err2_sol_2": "<strong>Гарантия сохранения сессии:</strong> Сброс очищает поврежденные кэши квоты, строго сохраняя ваши <code>Cookies</code>, историю чатов и аккаунт.",
        "tb_err3_badge": "СЕТЕВАЯ АРХИТЕКТУРА • БЕЗ РЕЖИМА TUN",
        "tb_err3_title": "Умное подключение прокси для v2ray, Xray, Clash и NekoBox",
        "tb_err3_cause": "<strong>Почему обычные VPN дают сбой:</strong> Режим TUN перенаправляет весь трафик системы, повышая пинг и ломая доступ к локальным сервисам.",
        "tb_err3_sol_title": "Как работает лаунчер без TUN в Antigravity Cleaner:",
        "tb_err3_sol_1": "Утилита автоматически сканирует локальные порты прокси:",
        "tb_err3_sol_2": "Запускает Antigravity с выделенным каналом: <code>--proxy-server=socks5://127.0.0.1:&lt;detected_port&gt;</code>.",
        "tb_err3_sol_3": "Только трафик Antigravity проходит через прокси. Браузер и остальные приложения работают на максимальной локальной скорости!",
        "tb_err3_note": "Создает постоянный ярлык на рабочем столе для запуска в 1 клик (<code>Antigravity.desktop</code> в Linux, <code>Antigravity.lnk</code> в Windows, <code>Antigravity-Proxy.command</code> в macOS).",
        "tb_err4_badge": "ОБРЫВ СОЕДИНЕНИЯ • MIDDLEBOX DPI",
        "tb_err4_title": "Сбросы TLS Middlebox и прерывание потока",
        "tb_err4_cause": "<strong>Проблема:</strong> Во время глубоких вычислений модели поток HTTP/2 затихает. Сетевые фильтры провайдера принимают бездействующее соединение за мертвое и отправляют RST-пакеты.",
        "tb_err4_sol_title": "Решение:",
        "tb_err4_sol_desc": "Antigravity Cleaner внедряет 15-секундный сигнал <strong>TCP KeepAlive</strong>, поддерживая соединение активным через любые фильтры.",
        "tb_err5_badge": "СБОЙ ЗАПУСКА • SINGLETON LOCK",
        "tb_err5_title": "Antigravity не открывается после сбоя или сна",
        "tb_err5_cause": "<strong>Проблема:</strong> При сбое Chromium оставляет файл <code>SingletonLock</code> с несуществующим PID, из-за чего приложение считает себя уже запущенным.",
        "tb_err5_sol_desc": "Лаунчер Antigravity Cleaner автоматически удаляет устаревшие файлы блокировки при каждом запуске!",
        "tb_err6_badge": "СБОЙ ПОТОКА • WRITE EIO / BROKEN PIPE",
        "tb_err6_title": "\"Agent terminated due to error\" и ошибки Broken Pipe",
        "tb_err6_cause": "<strong>Причина:</strong> Разрыв канала из-за избыточного логирования (<code>write EIO</code>) или тайم-аута соединения во время генерации ответа моделью.",
        "tb_err6_sol_title": "Мгновенное исправление с Antigravity Cleaner:",
        "tb_err6_sol_desc": "Нажмите <strong>[ ① ⚡ Auto-Fix ]</strong> в GUI или выполните <code>antigravity-cleaner patch all</code>. Защита Zero-Log глушит лишние логи и удерживает соединение активным.",
        "tb_err7_badge": "ОБРЫВ СВЯЗИ • DPI И УТЕЧКА ТЕЛЕМЕТРИИ",
        "tb_err7_title": "\"Unknown: There was a network issue connecting to the server\"",
        "tb_err7_cause": "<strong>Причина:</strong> Разрыв TLS/gRPC потоков сетевыми фильтрами или утечка реального IP через телеметрию посреди выполнения задачи.",
        "tb_err7_sol_desc": "Запустите Antigravity через <strong>[ ② 🚀 Launch IDE ]</strong> или <code>antigravity-cleaner launch</code>. Трафик изолируется через SOCKS5h с блокировкой телеметрии.",
        "tb_err8_badge": "БЛОК АВТОРИЗАЦИИ • 403 THAT’S AN ERROR",
        "tb_err8_title": "\"403. That’s an error. /auth-success\"",
        "tb_err8_cause": "<strong>Причина:</strong> Шлюз авторизации Google блокирует перенаправление из-за региональных отпечатков устройства или следов телеметрии.",
        "tb_err8_sol_desc": "Запустите <strong>[ ① ⚡ Auto-Fix ]</strong>. Cleaner пропатчит машинный код (<code>hasValidAuth = true</code>) и очистит следы старой авторизации.",
        "tb_err9_badge": "НЕСООТВЕТСТВИЕ ПРОКСИ • DIAL TCP 10808",
        "tb_err9_title": "\"proxyconnect tcp: dial tcp 127.0.0.1:10808: connect: connection refused\"",
        "tb_err9_cause": "<strong>Причина:</strong> Локальный прокси-клиент (Xray, Clash, NekoBox) закрыт или слушает другой порт.",
        "tb_err9_sol_desc": "Запустите <code>antigravity-cleaner doctor</code> (или <strong>[ 🩺 Doctor ]</strong> в GUI). Утилита просканирует активные порты и выставит правильный адрес.",
        "tb_err10_badge": "ПОВРЕЖДЕНИЕ СОСТОЯНИЯ • FAILED TO FETCH",
        "tb_err10_title": "\"There was an unexpected issue setting up your account. Failed to fetch\"",
        "tb_err10_cause": "<strong>Причина:</strong> Поврежденные записи базы данных DIPS и зависшие кэши WebSocket в состоянии приложения.",
        "tb_err10_sol_desc": "Нажмите <strong>[ 🧹 Reset 429 ]</strong> в GUI или выполните <code>antigravity-cleaner clean</code>. Очищает 16 временных артефактов за доли секунды без потери проектов.",
        "footer_note": "Antigravity Cleaner — это проект с открытым исходным кодом под лицензией GNU GPL v3.0.<br>Посвящается разработчикам, сталкивающимся с региональными ограничениями по всему миру. 🕊️",
        "docs": "Документация",
        "get_app": "Скачать",
        "tb_err7_sol_title": "Мгновенное решение:",
        "tb_err8_sol_title": "Мгновенное решение:",
        "tb_err9_sol_title": "Мгновенное решение:",
        "tb_err10_sol_title": "Мгновенное решение:"
    },
    "zh": {
        "home": "首页",
        "features": "功能",
        "download": "下载",
        "pricing": "价格",
        "blog": "博客",
        "faq": "常见问题",
        "hero_title": "立即修复 Antigravity IDE",
        "hero_subtitle": "最强大的安装错误和网络干扰维护工具。",
        "coming_soon": "即将推出",
        "offline_privacy": "100% 离线与隐私安全",
        "troubleshooting": "故障排除",
        "tb_hero_badge": "官方故障排除知识库",
        "tb_hero_title": "修复 Google Antigravity 错误与区域锁定",
        "tb_hero_desc": "针对 403 错误、429 配额用尽、中间件 TLS 重置及免 TUN 智能代理配置的完整逐步解决方案。",
        "tb_cta_download": "获取 Antigravity Cleaner（免费离线）",
        "tb_cta_star": "在 GitHub 上点赞",
        "tb_quickfix_title": "一键全能自动修复（推荐）",
        "tb_quickfix_desc": "无需手动修改复杂的 JavaScript 代码包和配置网络路由，运行 Antigravity Cleaner 即可自动对 4 个核心目标打补丁并生成独立的代理启动器：",
        "tb_quickfix_win": "* 在 Windows 上：打开 PowerShell 并运行 <code>irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex</code>",
        "tb_err1_badge": "错误 403 FORBIDDEN • 区域锁定",
        "tb_err1_title": "“该服务在您的国家/地区不可用” / 区域锁定",
        "tb_err1_cause": "<strong>根本原因：</strong>Google 在伊朗、俄罗斯、中国等地区对 Antigravity 实施了严格的地理访问限制。IDE 内部包含用于检查遥测和令牌响应的验证函数，一旦触发封锁便会强制使会话失效并注销登录。",
        "tb_err1_sol_title": "逐步修复 403 错误：",
        "tb_err1_sol_1": "<strong>解除客户端验证：</strong>运行 <code>antigravity-cleaner patch</code>，精确定位并解锁 <code>main.js</code>、IDE 二进制、CLI 及扩展中的限制。",
        "tb_err1_sol_2": "<strong>智能代理启动（免 TUN）：</strong>运行 <code>antigravity-cleaner launch</code>，直接向引擎传递本地 SOCKS5 代理，确保身份认证平稳通过且不发生泄漏。",
        "tb_err1_sol_3": "<strong>避免全局 TUN 模式：</strong>全局 TUN 易导致 MTU 数据包分片并触发防火墙深度检测导致的流中断。",
        "tb_err2_badge": "错误 429 • 速率超限 / 配额用尽",
        "tb_err2_title": "“Too Many Requests” 与损坏的令牌缓存",
        "tb_err2_cause": "<strong>根本原因：</strong>智能体在运行中产生大量调用记录，本地令牌缓存与云端服务器失去同步，导致在配额充足的情况下依然频繁弹出 429 错误。",
        "tb_err2_sol_title": "如何在不丢失数据的前提下解决 429 错误：",
        "tb_err2_sol_1": "在图形界面中点击 <strong>[Reset 429]</strong> 按钮或执行命令：",
        "tb_err2_sol_2": "<strong>会话安全保障：</strong>定向重置仅清除已损坏的配额状态及遥测缓存，全面保留登录凭证、<code>Cookie</code>、偏好设置和项目历史，无需重新登录。",
        "tb_err3_badge": "网络架构 • 免 TUN 模式",
        "tb_err3_title": "适用于 v2ray、Xray、Clash 与 NekoBox 的智能代理注入",
        "tb_err3_cause": "<strong>为什么传统 VPN 容易失效：</strong>开启全局 TUN 模式会导致操作系统全部流量被强行劫持，大幅增加延迟并破坏局域网及本地网站连接。",
        "tb_err3_sol_title": "Antigravity Cleaner 免 TUN 注入原理：",
        "tb_err3_sol_1": "工具自动扫描本地监听端口：",
        "tb_err3_sol_2": "使用专属管道启动 Antigravity：<code>--proxy-server=socks5://127.0.0.1:&lt;detected_port&gt;</code>。",
        "tb_err3_sol_3": "仅 Antigravity 流量经由代理路由，浏览器与日常网络应用继续保持最高直连速度！",
        "tb_err3_note": "这会生成一个永久的桌面一键启动快捷方式（Linux 下为 <code>Antigravity.desktop</code>，Windows 下为 <code>Antigravity.lnk</code>，macOS 下为 <code>Antigravity-Proxy.command</code>）。",
        "tb_err4_badge": "连接中断 • 深度包检测 DPI",
        "tb_err4_title": "中间件 TLS 断开与推流中断",
        "tb_err4_cause": "<strong>问题表现：</strong>在模型深度思考时，长连接保持静默。部分网络防火墙会误判静默的 TLS 连接并发送重置包，导致流断开错误。",
        "tb_err4_sol_title": "解决方案：",
        "tb_err4_sol_desc": "Antigravity Cleaner 在套接字层注入 <strong>15 秒心跳保活信号（TCP KeepAlive）</strong>，确保在各种复杂网络环境下长连接永不中断。",
        "tb_err5_badge": "启动失败 • 实例锁定",
        "tb_err5_title": "系统休眠或崩溃后 Antigravity 无法打开",
        "tb_err5_cause": "<strong>问题表现：</strong>系统休眠或程序异常崩溃后会残留 <code>SingletonLock</code> 文件，导致再次点击时程序误判已有实例在运行而拒绝启动。",
        "tb_err5_sol_desc": "Antigravity Cleaner 智能启动器会在每次启动时自动检测并清除残留的死锁文件！",
        "tb_err6_badge": "推流崩溃 • WRITE EIO / 管道中断",
        "tb_err6_title": "“Agent terminated due to error” 与管道中断异常",
        "tb_err6_cause": "<strong>根本原因：</strong>输出日志过量（<code>write EIO</code>）或在智能体思考推流时网络超时，导致数据管道突然断开。",
        "tb_err6_sol_title": "使用 Antigravity Cleaner 瞬间修复：",
        "tb_err6_sol_desc": "在图形界面点击 <strong>[ ① ⚡ Auto-Fix ]</strong> 或运行 <code>antigravity-cleaner patch all</code>。零日志屏蔽层将静默冗余日志，并通过 15 秒心跳彻底解决管道中断。",
        "tb_err7_badge": "连接掉线 • DPI 与遥测泄漏",
        "tb_err7_title": "“Unknown: There was a network issue connecting to the server” 错误",
        "tb_err7_cause": "<strong>根本原因：</strong>防火墙掐断了长时推理的 gRPC 连接，或后台遥测探测在执行任务中途泄露了真实 IP。",
        "tb_err7_sol_desc": "通过 <strong>[ ② 🚀 Launch IDE ]</strong> 或 <code>antigravity-cleaner launch</code> 启动。所有流量均经由本地安全套接字隔离并激活零遥测封装。",
        "tb_err8_badge": "认证拦截 • 403 THAT’S AN ERROR",
        "tb_err8_title": "“403. That’s an error. /auth-success” 授权错误",
        "tb_err8_cause": "<strong>根本原因：</strong>Google 区域认证网关通过历史设备指纹或未屏蔽的标记检测到受限区域，从而拦截登录重定向。",
        "tb_err8_sol_desc": "运行 <strong>[ ① ⚡ Auto-Fix ]</strong>。清理工具将修补底层机器码（强制 <code>hasValidAuth = true</code>），注入企业认证模式并清理损坏的令牌。",
        "tb_err9_badge": "代理端口不匹配 • DIAL TCP 10808",
        "tb_err9_title": "“proxyconnect tcp: dial tcp 127.0.0.1:10808: connect: connection refused” 错误",
        "tb_err9_cause": "<strong>根本原因：</strong>本地代理客户端（Xray、Clash、NekoBox 等）未开启或配置在其他端口上。",
        "tb_err9_sol_desc": "运行 <code>antigravity-cleaner doctor</code>（或在界面点击 <strong>[ 🩺 Doctor ]</strong>），诊断引擎将自动扫描活动端口并调整配置。",
        "tb_err10_badge": "状态损坏 • FAILED TO FETCH",
        "tb_err10_title": "“There was an unexpected issue setting up your account. Failed to fetch” 错误",
        "tb_err10_cause": "<strong>根本原因：</strong>应用状态中受损的 DIPS 数据库记录和孤立的 WebSocket 缓存。",
        "tb_err10_sol_desc": "在图形界面点击 <strong>[ 🧹 Reset 429 ]</strong> 或执行 <code>antigravity-cleaner clean</code>。在 0.002 秒内清除 16 个缓存项，同时 100% 保留项目与历史记录。",
        "footer_note": "Antigravity Cleaner 是一个基于 GNU GPL v3.0 许可发布的开源社区项目。<br>献给全球所有面临地理封锁与数字孤立的开发者。🕊️",
        "docs": "文档",
        "get_app": "获取应用",
        "tb_err7_sol_title": "即时修复方案：",
        "tb_err8_sol_title": "即时修复方案：",
        "tb_err9_sol_title": "即时修复方案：",
        "tb_err10_sol_title": "即时修复方案："
    },
    "tr": {
        "offline_privacy": "Çevrimdışı ve Gizli",
        "troubleshooting": "Sorun Giderme",
        "tb_hero_badge": "RESMİ SORUN GİDERME BİLGİ BANKASI",
        "tb_hero_title": "Google Antigravity Hatalarını ve Bölge Kilidini Düzeltin",
        "tb_hero_desc": "403 Forbidden, 429 Kota Aşımı, Middlebox TLS Sıfırlamaları ve TUN Olmayan Akıllı Proxy kurulumu için adım adım çözümler.",
        "tb_cta_download": "Antigravity Cleaner İndir (Ücretsiz ve Çevrimdışı)",
        "tb_cta_star": "GitHub'da Yıldız Ver",
        "tb_quickfix_title": "1 Tıkla Evrensel Otomatik Düzeltme (Önerilen)",
        "tb_quickfix_desc": "JavaScript paketlerini elle düzenlemek ve ağ yönlendirmesini yapılandırmak yerine, Antigravity Cleaner'ı çalıştırarak 4 hedefin tümünü yamalayın ve otomatik olarak yalıtılmış bir proxy başlatıcısı oluşturun:",
        "tb_quickfix_win": "* Windows'ta: PowerShell'i açın ve çalıştırın: <code>irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex</code>",
        "tb_err1_badge": "HATA 403 FORBIDDEN • BÖLGE KİLİDİ",
        "tb_err1_title": "\"Hizmet ülkenizde kullanılamıyor\" / Bölge Kilidi",
        "tb_err1_cause": "<strong>Temel Neden:</strong> Google DeepMind, belirli ülkelerde Antigravity üzerinde sıkı coğrafi kısıtlamalar uygular. IDE, telemetri ve belirteç yanıtlarını inceleyen dahili doğrulama işlevleri çalıştırır ve engellendiğinde oturumu kapatır.",
        "tb_err1_sol_title": "Adım Adım Hata 403 Çözümü:",
        "tb_err1_sol_1": "<strong>İstemci Kontrollerini Kaldırın:</strong> <code>antigravity-cleaner patch</code> komutunu çalıştırarak <code>main.js</code>, IDE ikili dosyası, CLI ve eklentideki kontrolleri devre dışı bırakın.",
        "tb_err1_sol_2": "<strong>Akıllı Proxy ile Başlatın (TUN Yok):</strong> <code>antigravity-cleaner launch</code> komutunu çalıştırarak doğrudan Chromium motoruna SOCKS5 proxy parametresini iletin.",
        "tb_err1_sol_3": "<strong>TUN Modundan Kaçının:</strong> Genel TUN modu paket parçalanmasına ve Middlebox DPI bağlantı kopmalarına neden olur.",
        "tb_err2_badge": "HATA 429 • İSTEK SINIRI / KOTA TÜKENDİ",
        "tb_err2_title": "\"Too Many Requests\" ve Bozuk Belirteç Önbelleği",
        "tb_err2_cause": "<strong>Temel Neden:</strong> Otonom ajanlar oturum başına onlarca çağrı üretir. Yerel önbellekler sunucuyla senkronizasyonunu kaybederek kotanız olmasına rağmen 429 hatası verir.",
        "tb_err2_sol_title": "Veri Kaybetmeden 429 Hatasını Çözme:",
        "tb_err2_sol_1": "Antigravity Cleaner arayüzündeki <strong>[Reset 429]</strong> butonuna tıklayın veya çalıştırın:",
        "tb_err2_sol_2": "<strong>Oturum Koruma Garantisi:</strong> Sıfırlama işlemi çerezlerinizi ve geçmişinizi koruyarak yalnızca bozuk önbellekleri temizler; tekrar giriş yapmanız gerekmez.",
        "tb_err3_badge": "AĞ MİMARİSİ • TUN MODSUZ",
        "tb_err3_title": "v2ray, Xray, Clash ve NekoBox için Akıllı Proxy Enjeksiyonu",
        "tb_err3_cause": "<strong>Geleneksel VPN'ler neden başarısız olur:</strong> TUN modu tüm sistem trafiğini yönlendirir; bu da gecikmeyi artırır ve yerel sitelere erişimi keser.",
        "tb_err3_sol_title": "Antigravity Cleaner TUN Olmayan Enjeksiyon Nasıl Çalışır:",
        "tb_err3_sol_1": "Yazılım yerel loopback bağlantı noktalarını otomatik olarak tarar:",
        "tb_err3_sol_2": "Antigravity'yi özel parametreyle başlatır: <code>--proxy-server=socks5://127.0.0.1:&lt;detected_port&gt;</code>.",
        "tb_err3_sol_3": "Yalnızca Antigravity trafiği proxy'den geçer; tarayıcınız ve diğer uygulamalarınız tam hızda yerel olarak çalışmaya devam eder!",
        "tb_err3_note": "Masaüstünde tek tıkla çalışan kalıcı bir başlatıcı oluşturur (Linux'ta <code>Antigravity.desktop</code>, Windows'ta <code>Antigravity.lnk</code>, macOS'ta <code>Antigravity-Proxy.command</code>).",
        "tb_err4_badge": "BAĞLANTI KOPMASI • MIDDLEBOX DPI",
        "tb_err4_title": "Middlebox TLS Bağlantı Kopmaları ve Akış Kesintileri",
        "tb_err4_cause": "<strong>Sorun:</strong> Modelin düşünme aşamasında bağlantı sessiz kalır. Katı ISP güvenlik duvarları boşta kalan bağlantıyı keserek akış kopmalarına yol açar.",
        "tb_err4_sol_title": "Çözüm:",
        "tb_err4_sol_desc": "Antigravity Cleaner, <strong>15 saniyelik TCP KeepAlive</strong> sinyali enjekte ederek oturumlarınızın kesilmesini önler.",
        "tb_err5_badge": "BAŞLATMA HATASI • SINGLETON LOCK",
        "tb_err5_title": "Kilitlenme veya Uyku Sonrası Antigravity Açılmıyor",
        "tb_err5_cause": "<strong>Sorun:</strong> Çökme sonrası Chromium kilit dosyası bırakır ve program başka bir örneğin açık olduğunu düşünerek başlamayı reddeder.",
        "tb_err5_sol_desc": "Antigravity Cleaner akıllı başlatıcısı, her başlatmada eski kilitleri otomatik olarak denetler ve temizler!",
        "tb_err6_badge": "AKIŞ ÇÖKMESİ • WRITE EIO / BROKEN PIPE",
        "tb_err6_title": "\"Agent terminated due to error\" ve Broken Pipe Hataları",
        "tb_err6_cause": "<strong>Temel Neden:</strong> Ayrıntılı ham günlük kaydı veya aracı döngüleri sırasında bağlantı zaman aşımlarından kaynaklanan ani boru hattı arızası.",
        "tb_err6_sol_title": "Antigravity Cleaner ile Anında Düzeltme:",
        "tb_err6_sol_desc": "GUI'de <strong>[ ① ⚡ Auto-Fix ]</strong> butonuna tıklayın veya <code>antigravity-cleaner patch all</code> çalıştırın. Sıfır-Günlük kalkanı akışı canlı tutar.",
        "tb_err7_badge": "BAĞLANTI KOPMASI • DPI VE TELEMETRİ SIZINTISI",
        "tb_err7_title": "\"Unknown: There was a network issue connecting to the server\"",
        "tb_err7_cause": "<strong>Temel Neden:</strong> DPI güvenlik duvarlarının gRPC akışlarını sonlandırması veya arka plan telemetrisinin gerçek IP'yi sızdırması.",
        "tb_err7_sol_desc": "Antigravity'yi <strong>[ ② 🚀 Launch IDE ]</strong> veya <code>antigravity-cleaner launch</code> ile başlatın; telemetri sızıntısı engellenir.",
        "tb_err8_badge": "GİRİŞ ENGELİ • 403 THAT’S AN ERROR",
        "tb_err8_title": "\"403. That’s an error. /auth-success\"",
        "tb_err8_cause": "<strong>Temel Neden:</strong> Google kimlik doğrulama ağ geçidinin kısıtlı bölgelerden gelen tarayıcı yönlendirmesini engellemesi.",
        "tb_err8_sol_desc": "<strong>[ ① ⚡ Auto-Fix ]</strong> seçeneğini çalıştırın; makine kodunu yamalar ve bozuk kimlik doğrulama belirteçlerini temizler.",
        "tb_err9_badge": "PROXY UYUŞMAZLIĞI • DIAL TCP 10808",
        "tb_err9_title": "\"proxyconnect tcp: dial tcp 127.0.0.1:10808: connect: connection refused\"",
        "tb_err9_cause": "<strong>Temel Neden:</strong> Yerel proxy istemciniz (Xray, Clash, NekoBox) kapalı veya farklı bir bağlantı noktasını dinliyor.",
        "tb_err9_sol_desc": "<code>antigravity-cleaner doctor</code> çalıştırın; tanı motoru aktif portları tarayarak ayarları otomatik olarak eşleştirecektir.",
        "tb_err10_badge": "DURUM BOZULMASI • FAILED TO FETCH",
        "tb_err10_title": "\"There was an unexpected issue setting up your account. Failed to fetch\"",
        "tb_err10_cause": "<strong>Temel Neden:</strong> Uygulama durumundaki bozuk DIPS veritabanı kayıtları ve yetim WebSocket önbellekleri.",
        "tb_err10_sol_desc": "GUI'de <strong>[ 🧹 Reset 429 ]</strong> butonuna tıklayın veya <code>antigravity-cleaner clean</code> çalıştırın; projeleriniz korunarak önbellek temizlenir.",
        "footer_note": "Antigravity Cleaner, GNU GPL v3.0 altında yayınlanan topluluk odaklı bir açık kaynak projesidir.<br>Dünya genelinde bölgesel kısıtlamalarla karşılaşan geliştiricilere adanmıştır. 🕊️",
        "home": "Ana Sayfa",
        "features": "Özellikler",
        "download": "İndir",
        "docs": "Belgeler",
        "faq": "SSS",
        "blog": "Blog",
        "pricing": "Fiyatlandırma",
        "get_app": "Uygulamayı İndir",
        "tb_err7_sol_title": "Anında Düzeltme:",
        "tb_err8_sol_title": "Anında Düzeltme:",
        "tb_err9_sol_title": "Anında Düzeltme:",
        "tb_err10_sol_title": "Anında Düzeltme:"
    },
    "es": {
        "offline_privacy": "100% Desconectado y Privado",
        "troubleshooting": "Solución de problemas",
        "tb_hero_badge": "BASE DE CONOCIMIENTOS OFICIAL PARA SOLUCIÓN DE PROBLEMAS",
        "tb_hero_title": "Solucionar Errores de Google Antigravity y Bloqueo de Región",
        "tb_hero_desc": "Soluciones completas paso a paso para el Error 403 Forbidden, 429 Agotamiento de Cuota, Reinicios TLS de Middlebox y configuración de Proxy Inteligente sin TUN.",
        "tb_cta_download": "Obtener Antigravity Cleaner (Gratis y Offline)",
        "tb_cta_star": "Estrella en GitHub",
        "tb_quickfix_title": "Reparación Automática Universal en 1 Clic (Recomendado)",
        "tb_quickfix_desc": "En lugar de editar manualmente paquetes de JavaScript y configurar el enrutamiento de red, ejecute Antigravity Cleaner para parchear los 4 objetivos y generar un iniciador de proxy aislado automáticamente:",
        "tb_quickfix_win": "* En Windows: Abra PowerShell y ejecute <code>irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex</code>",
        "tb_err1_badge": "ERROR 403 FORBIDDEN • BLOQUEO REGIONAL",
        "tb_err1_title": "\"El servicio no está disponible en su país\" / Bloqueo Regional",
        "tb_err1_cause": "<strong>Causa Principal:</strong> Google aplica restricciones geográficas estrictas en Antigravity en ciertos países. El IDE ejecuta funciones internas que inspeccionan la telemetría y los tokens, cerrando la sesión si se detecta una región bloqueada.",
        "tb_err1_sol_title": "Solución Paso a Paso para el Error 403:",
        "tb_err1_sol_1": "<strong>Neutralizar Verificaciones:</strong> Ejecute <code>antigravity-cleaner patch</code> para desbloquear con precisión <code>main.js</code>, el binario del IDE, la CLI y la extensión.",
        "tb_err1_sol_2": "<strong>Iniciar con Proxy Inteligente (Sin TUN):</strong> Ejecute <code>antigravity-cleaner launch</code> para inyectar la configuración SOCKS5 directamente en el motor de la aplicación.",
        "tb_err1_sol_3": "<strong>Evite el Modo TUN:</strong> El modo TUN global provoca fragmentación de paquetes MTU y cortes por cortafuegos DPI.",
        "tb_err2_badge": "ERROR 429 • LÍMITE DE TASA / CUOTA AGOTADA",
        "tb_err2_title": "\"Too Many Requests\" y Caché de Tokens Dañada",
        "tb_err2_cause": "<strong>Causa Principal:</strong> Los agentes generan múltiples llamadas y transcripciones. Las cachés locales de tokens pierden sincronía con los servidores, causando errores 429 continuos aun con cuota disponible.",
        "tb_err2_sol_title": "Cómo Resolver el Error 429 sin Perder Datos:",
        "tb_err2_sol_1": "Haga clic en el botón <strong>[Reset 429]</strong> en la interfaz o ejecute:",
        "tb_err2_sol_2": "<strong>Garantía de Preservación de Sesión:</strong> La limpieza purga cachés dañadas manteniendo al 100% sus <code>Cookies</code>, cuentas e historial de proyectos.",
        "tb_err3_badge": "ARQUITECTURA DE RED • MODO SIN TUN",
        "tb_err3_title": "Inyección Inteligente de Proxy para v2ray, Xray, Clash y NekoBox",
        "tb_err3_cause": "<strong>Por qué fallan las VPN tradicionales:</strong> El modo TUN redirige todo el tráfico del sistema, aumentando la latencia y rompiendo el acceso a sitios locales.",
        "tb_err3_sol_title": "Cómo Funciona la Inyección sin TUN de Antigravity Cleaner:",
        "tb_err3_sol_1": "El limpiador escanea automáticamente los puertos locales activos:",
        "tb_err3_sol_2": "Inicia Antigravity con un canal dedicado: <code>--proxy-server=socks5://127.0.0.1:&lt;detected_port&gt;</code>.",
        "tb_err3_sol_3": "Solo el tráfico de Antigravity pasa por el proxy. El navegador y el resto de apps continúan a la máxima velocidad local.",
        "tb_err3_note": "Genera un acceso directo permanente de 1 clic en el escritorio (<code>Antigravity.desktop</code> en Linux, <code>Antigravity.lnk</code> en Windows, <code>Antigravity-Proxy.command</code> en macOS).",
        "tb_err4_badge": "PÉRDIDA DE CONEXIÓN • DPI DE MIDDLEBOX",
        "tb_err4_title": "Desconexiones TLS de Middlebox e Interrupciones de Flujo",
        "tb_err4_cause": "<strong>El Problema:</strong> Durante el razonamiento del modelo, la conexión HTTP/2 permanece en silencio. Los cortafuegos del ISP cierran la conexión por inactividad.",
        "tb_err4_sol_title": "La Solución:",
        "tb_err4_sol_desc": "Antigravity Cleaner inyecta un latido <strong>TCP KeepAlive de 15 segundos</strong> para mantener la conexión activa frente a cortafuegos.",
        "tb_err5_badge": "ERROR DE INICIO • BLOQUEO DE INSTANCIA",
        "tb_err5_title": "Antigravity no Abre Tras un Cierre Inesperado o Suspensión",
        "tb_err5_cause": "<strong>El Problema:</strong> Tras un bloqueo, queda un archivo <code>SingletonLock</code> residual que impide volver a abrir la aplicación.",
        "tb_err5_sol_desc": "¡El iniciador inteligente de Antigravity Cleaner audita y elimina bloqueos obsoletos automáticamente en cada inicio!",
        "tb_err6_badge": "CORTE DE FLUJO • WRITE EIO / CONDUCTO ROTO",
        "tb_err6_title": "\"Agent terminated due to error\" y Excepciones de Tubería Rota",
        "tb_err6_cause": "<strong>Causa Principal:</strong> Ruptura repentina de la canalización por registro excesivo o tiempos de espera de la red durante el streaming.",
        "tb_err6_sol_title": "Solución Instantánea con Antigravity Cleaner:",
        "tb_err6_sol_desc": "Haga clic en <strong>[ ① ⚡ Auto-Fix ]</strong> en la GUI o ejecute <code>antigravity-cleaner patch all</code> para silenciar registros y mantener el streaming activo.",
        "tb_err7_badge": "CAÍDA DE CONEXIÓN • DPI Y FUGA DE TELEMETRÍA",
        "tb_err7_title": "\"Unknown: There was a network issue connecting to the server\"",
        "tb_err7_cause": "<strong>Causa Principal:</strong> Cortafuegos que terminan flujos gRPC o telemetría que filtra la IP real a mitad de una tarea.",
        "tb_err7_sol_desc": "Inicie Antigravity mediante <strong>[ ② 🚀 Launch IDE ]</strong> o <code>antigravity-cleaner launch</code> con el aislamiento Zero-Telemetry activo.",
        "tb_err8_badge": "BLOQUEO DE AUTENTICACIÓN • 403 THAT’S AN ERROR",
        "tb_err8_title": "\"403. That’s an error. /auth-success\"",
        "tb_err8_cause": "<strong>Causa Principal:</strong> La pasarela de Google bloquea la redirección del navegador por huellas del dispositivo o balizas de telemetría.",
        "tb_err8_sol_desc": "Ejecute <strong>[ ① ⚡ Auto-Fix ]</strong>; parchea el código máquina binario y purga los tokens de autenticación corruptos.",
        "tb_err9_badge": "DESAJUSTE DE PROXY • DIAL TCP 10808",
        "tb_err9_title": "\"proxyconnect tcp: dial tcp 127.0.0.1:10808: connect: connection refused\"",
        "tb_err9_cause": "<strong>Causa Principal:</strong> El cliente de proxy local (Xray, Clash, NekoBox) está apagado o en un puerto diferente.",
        "tb_err9_sol_desc": "Ejecute <code>antigravity-cleaner doctor</code> para escanear los puertos activos y sincronizar la configuración del proxy.",
        "tb_err10_badge": "CORRUPCIÓN DE ESTADO • FAILED TO FETCH",
        "tb_err10_title": "\"There was an unexpected issue setting up your account. Failed to fetch\"",
        "tb_err10_cause": "<strong>Causa Principal:</strong> Registros corruptos en la base de datos DIPS y cachés huérfanas de WebSocket.",
        "tb_err10_sol_desc": "Haga clic en <strong>[ 🧹 Reset 429 ]</strong> o ejecute <code>antigravity-cleaner clean</code>; elimina 16 artefactos dañados preservando su historial.",
        "footer_note": "Antigravity Cleaner es un proyecto de código abierto bajo licencia GNU GPL v3.0.<br>Dedicado a los desarrolladores que enfrentan restricciones geográficas y aislamiento digital en todo el mundo. 🕊️",
        "home": "Inicio",
        "features": "Características",
        "download": "Descargar",
        "docs": "Documentación",
        "faq": "Preguntas",
        "blog": "Blog",
        "pricing": "Precios",
        "get_app": "Obtener App",
        "tb_err7_sol_title": "Solución Instantánea:",
        "tb_err8_sol_title": "Solución Instantánea:",
        "tb_err9_sol_title": "Solución Instantánea:",
        "tb_err10_sol_title": "Solución Instantánea:"
    },
    "ar": {
        "offline_privacy": "أوفلاين وخاص بالكامل",
        "troubleshooting": "حل المشكلات",
        "tb_hero_badge": "قاعدة المعرفة الرسمية لحل المشكلات",
        "tb_hero_title": "إصلاح أخطاء Google Antigravity وقفل المنطقة",
        "tb_hero_desc": "حلول شاملة خطوة بخطوة للخطأ 403 Forbidden، ونفاد الحصة 429، وإعادة تعيين Middlebox TLS، وإعداد البروكسي الذكي بدون وضع TUN.",
        "tb_cta_download": "احصل على Antigravity Cleaner (مجاني وأوفلاين)",
        "tb_cta_star": "نجمة على GitHub",
        "tb_quickfix_title": "إصلاح تلقائي شامل بنقرة واحدة (موصى به)",
        "tb_quickfix_desc": "بدلاً من تعديل ملفات JavaScript يدوياً وإعدادات التوجيه المعقدة، قم بتشغيل Antigravity Cleaner لتصحيح جميع الأهداف الأربعة وإنشاء مشغل بروكسي معزول تلقائياً:",
        "tb_quickfix_win": "* على نظام Windows: افتح PowerShell وشغّل <code>irm https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex</code>",
        "tb_err1_badge": "خطأ 403 FORBIDDEN • قفل المنطقة الجغرافية",
        "tb_err1_title": "\"الخدمة غير متوفرة في بلدك\" / حظر جغرافي",
        "tb_err1_cause": "<strong>السبب الجذري:</strong> تفرض Google قيوداً جغرافية صارمة على Antigravity في دول مثل إيران وروسيا والصين. يحتوي المحرر على دوال داخلية تفحص الموقع الجغرافي وتلغي الجلسة تلقائياً عند اكتشاف منطقة محظورة.",
        "tb_err1_sol_title": "خطوات حل الخطأ 403 خطوة بخطوة:",
        "tb_err1_sol_1": "<strong>إلغاء فحوصات العميل:</strong> نفّذ الأمر <code>antigravity-cleaner patch</code> لتجاوز القيود في <code>main.js</code> وملفات المحرر والإضافة تلقائياً.",
        "tb_err1_sol_2": "<strong>التشغيل عبر بروكسي ذكي (بدون TUN):</strong> شغّل <code>antigravity-cleaner launch</code> لتمرير إعدادات SOCKS5 المحلية إلى محرك البرنامج مباشرة.",
        "tb_err1_sol_3": "<strong>تجنب وضع TUN العام:</strong> يؤدي وضع TUN الشامل إلى تجزئة الحزم وانقطاع الاتصال بواسطة جدران الحماية.",
        "tb_err2_badge": "خطأ 429 • تجاوز حد الطلبات أو الحصة",
        "tb_err2_title": "\"Too Many Requests\" وذاكرة التخزين المؤقت التالفة",
        "tb_err2_cause": "<strong>السبب الجذري:</strong> تولد نماذج الذكاء الاصطناعي استدعاءات مكثفة. تفقد الذاكرة المؤقتة المحلية مزامنتها مع خوادم Google مما يؤدي إلى ظهور الخطأ 429 بالرغم من توفر الرصيد.",
        "tb_err2_sol_title": "طريقة حل الخطأ 429 دون فقدان البيانات:",
        "tb_err2_sol_1": "انقر على زر <strong>[Reset 429]</strong> في الواجهة الرسومية أو نفّذ:",
        "tb_err2_sol_2": "<strong>ضمان الحفاظ على الجلسة:</strong> يقوم التنظيف بإزالة الملفات التالفة مع الحفاظ الكامل على بيانات تسجيل الدخول وملفات تعريف الارتباط وتاريخ المحادثات.",
        "tb_err3_badge": "بنية الشبكة • وضع بدون TUN",
        "tb_err3_title": "حقن البروكسي الذكي لبرامج v2ray و Xray و Clash و NekoBox",
        "tb_err3_cause": "<strong>لماذا تفشل شبكات VPN التقليدية:</strong> يؤدي تفعيل وضع TUN إلى توجيه كل حركة بيانات الجهاز، مما يزيد وقت الاستجابة ويعطل الخدمات المحلية.",
        "tb_err3_sol_title": "كيف تعمل آلية الربط بدون TUN في البرنامج:",
        "tb_err3_sol_1": "يقوم البرنامج بفحص المنافذ المحلية النشطة تلقائياً:",
        "tb_err3_sol_2": "يشغل Antigravity بقناة مخصصة: <code>--proxy-server=socks5://127.0.0.1:&lt;detected_port&gt;</code>.",
        "tb_err3_sol_3": "تمر حركة بيانات المحرر فقط عبر البروكسي، بينما يستمر المتصفح والتطبيقات الأخرى بالعمل بالسرعة القصوى المباشرة!",
        "tb_err3_note": "ينشئ اختصاراً دائماً على سطح المكتب للتشغيل بنقرة واحدة (<code>Antigravity.desktop</code> في Linux، و<code>Antigravity.lnk</code> في Windows).",
        "tb_err4_badge": "انقطاع الاتصال • فحص الحزم العميق DPI",
        "tb_err4_title": "انقطاعات TLS وتعطل تدفق استجابة الذكاء الاصطناعي",
        "tb_err4_cause": "<strong>المشكلة:</strong> أثناء معالجة النموذج للأوامر المعقدة، يبقى الاتصال صامتاً لفترة، مما يدفع جدران الحماية لقطع الاتصال ظناً منها أنه غير نشط.",
        "tb_err4_sol_title": "الحل النهائي:",
        "tb_err4_sol_desc": "يحقن البرنامج إشارة نبضات <strong>TCP KeepAlive كل 15 ثانية</strong> للحفاظ على استمرارية الاتصال ومنع انقطاع الاستجابة.",
        "tb_err5_badge": "فشل الإقلاع • ملف القفل Singleton",
        "tb_err5_title": "عدم استجابة البرنامج بعد إيقاف التشغيل المفاجئ أو وضع السكون",
        "tb_err5_cause": "<strong>المشكلة:</strong> عند إغلاق النظام المفاجئ، يبقى ملف القفل معلقاً مما يمنع البرنامج من الفتح مجدداً لاعتقاده بوجود نسخة تعمل.",
        "tb_err5_sol_desc": "يقوم مشغل Antigravity Cleaner الذكي بالتحقق من ملفات القفل العالقة وحذفها تلقائياً عند كل تشغيل!",
        "tb_err6_badge": "انهيار التدفق • خطأ WRITE EIO / انقطاع القناة",
        "tb_err6_title": "رسالة الخطأ \"Agent terminated due to error\" واستثناءات القناة",
        "tb_err6_cause": "<strong>السبب الجذري:</strong> انهيار مفاجئ في قناة تدفق البيانات ناتج عن تسجيل السجلات المفرط أو انتهاء مهلة الاتصال أثناء التفكير.",
        "tb_err6_sol_title": "الحل الفوري عبر البرنامج:",
        "tb_err6_sol_desc": "انقر على زر <strong>[ ① ⚡ Auto-Fix ]</strong> في الواجهة أو شغّل <code>antigravity-cleaner patch all</code> لكتم السجلات الزائدة والحفاظ على ثبات التدفق.",
        "tb_err7_badge": "انقطاع الاتصال • تسريب بيانات التتبع والحظر",
        "tb_err7_title": "خطأ \"Unknown: There was a network issue connecting to the server\"",
        "tb_err7_cause": "<strong>السبب الجذري:</strong> إنهاء جدران الحماية لتدفقات gRPC أو تسريب الـ IP الحقيقي عبر تتبع الخلفية أثناء العمل.",
        "tb_err7_sol_desc": "شغّل البرنامج عبر زر <strong>[ ② 🚀 Launch IDE ]</strong> أو الأمر <code>antigravity-cleaner launch</code> لعزل البيانات بالكامل مع إيقاف التتبع.",
        "tb_err8_badge": "حظر تسجيل الدخول • 403 THAT’S AN ERROR",
        "tb_err8_title": "خطأ \"403. That’s an error. /auth-success\" أثناء التفويض",
        "tb_err8_cause": "<strong>السبب الجذري:</strong> حظر بوابة تفويض Google لإعادة توجيه المتصفح بسبب بصمة الجهاز أو تتبع الموقع الجغرافي.",
        "tb_err8_sol_desc": "شغّل خيار <strong>[ ① ⚡ Auto-Fix ]</strong>؛ يقوم البرنامج بتعديل التعليمات البرمجية الأساسية ومسح رموز الدخول التالفة.",
        "tb_err9_badge": "عدم تطابق منفذ البروكسي • 10808",
        "tb_err9_title": "خطأ \"proxyconnect tcp: dial tcp 127.0.0.1:10808: connect: connection refused\"",
        "tb_err9_cause": "<strong>السبب الجذري:</strong> برنامج البروكسي المحلي مغلق أو يستمع إلى منفذ آخر مختلف.",
        "tb_err9_sol_desc": "نفّذ <code>antigravity-cleaner doctor</code> أو انقر <strong>[ 🩺 Doctor ]</strong>؛ يقوم الفاحص باكتشاف المنافذ ومطابقتها تلقائياً.",
        "tb_err10_badge": "تلف حالة الجلسة • Failed to fetch",
        "tb_err10_title": "خطأ \"There was an unexpected issue setting up your account. Failed to fetch\"",
        "tb_err10_cause": "<strong>السبب الجذري:</strong> سجلات قاعدة بيانات DIPS التالفة وبقايا ذاكرة WebSocket المؤقتة.",
        "tb_err10_sol_desc": "انقر على <strong>[ 🧹 Reset 429 ]</strong> أو نفّذ <code>antigravity-cleaner clean</code>؛ يمسح 16 عنصراً تالفاً مع الحفاظ التام على المشاريع والمحادثات.",
        "footer_note": "Antigravity Cleaner هو مشروع مفتوح المصدر تم إصداره بموجب ترخيص GNU GPL v3.0.<br>مُهدى إلى المطورين الذين يواجهون القيود الجغرافية والعزلة الرقمية حول العالم. 🕊️",
        "home": "الرئيسية",
        "features": "المميزات",
        "download": "تنزيل",
        "docs": "المستندات",
        "faq": "الأسئلة الشائعة",
        "blog": "المدونة",
        "pricing": "الأسعار",
        "get_app": "تحميل البرنامج",
        "tb_err7_sol_title": "الحل الفوري:",
        "tb_err8_sol_title": "الحل الفوري:",
        "tb_err9_sol_title": "الحل الفوري:",
        "tb_err10_sol_title": "الحل الفوري:"
    }
};

const langMap = {
    en: { flag: '🇺🇸', name: 'English' },
    fa: { flag: '🇮🇷', name: 'فارسی' },
    ru: { flag: '🇷🇺', name: 'Русский' },
    zh: { flag: '🇨🇳', name: '简体中文' },
    tr: { flag: '🇹🇷', name: 'Türkçe' },
    es: { flag: '🇪🇸', name: 'Español' },
    ar: { flag: '🇸🇦', name: 'العربية' }
};

// ========================================
// THEME MANAGEMENT SYSTEM
// ========================================

function initThemeToggle() {
    const savedTheme = localStorage.getItem('ag_theme') || (window.matchMedia && window.matchMedia('(prefers-color-scheme: light)').matches ? 'light' : 'dark');
    document.documentElement.setAttribute('data-theme', savedTheme);
    if (document.body) {
        document.body.setAttribute('data-theme', savedTheme);
    }

    const toggleBtn = document.getElementById('themeToggleBtn');
    if (toggleBtn) {
        toggleBtn.onclick = function (e) {
            e.preventDefault();
            const currentTheme = document.documentElement.getAttribute('data-theme') || 'dark';
            const nextTheme = currentTheme === 'light' ? 'dark' : 'light';
            document.documentElement.setAttribute('data-theme', nextTheme);
            if (document.body) {
                document.body.setAttribute('data-theme', nextTheme);
            }
            localStorage.setItem('ag_theme', nextTheme);
        };
    }
}

// ========================================
// LANGUAGE LOCALIZATION SYSTEM
// ========================================

function updateCustomDropdownUI(lang) {
    const flagElem = document.getElementById('currLangFlag');
    const nameElem = document.getElementById('currLangName');
    if (langMap[lang]) {
        if (flagElem) flagElem.innerText = langMap[lang].flag;
        if (nameElem) nameElem.innerText = langMap[lang].name;
    }
    document.querySelectorAll('.lang-option').forEach(opt => {
        if (opt.getAttribute('data-lang') === lang) {
            opt.classList.add('active');
        } else {
            opt.classList.remove('active');
        }
    });
}

function setLanguage(lang) {
    document.documentElement.lang = lang;
    const rtlLanguages = ['fa', 'ar', 'ur'];
    document.documentElement.dir = rtlLanguages.includes(lang) ? 'rtl' : 'ltr';

    document.querySelectorAll('[data-lang-key]').forEach(elem => {
        const key = elem.getAttribute('data-lang-key');
        const text = (translations[lang] && translations[lang][key]) || (translations['en'] && translations['en'][key]);
        if (text !== undefined && text !== null) {
            if (elem.getAttribute('data-lang-html') === 'true' || text.includes('<') || text.includes('&')) {
                elem.innerHTML = text;
            } else {
                elem.innerText = text;
            }
        }
    });

    localStorage.setItem('ag_lang', lang);
    const langSelect = document.getElementById('langSelect');
    if (langSelect) langSelect.value = lang;
    updateCustomDropdownUI(lang);
}

function initCustomLangDropdown() {
    const customLangSelector = document.getElementById('customLangSelector');
    const langDropdownBtn = document.getElementById('langDropdownBtn');

    if (langDropdownBtn && customLangSelector) {
        langDropdownBtn.onclick = function (e) {
            e.stopPropagation();
            customLangSelector.classList.toggle('open');
        };

        customLangSelector.querySelectorAll('.lang-option').forEach(opt => {
            opt.onclick = function (e) {
                e.stopPropagation();
                const lang = opt.getAttribute('data-lang');
                setLanguage(lang);
                customLangSelector.classList.remove('open');
            };
        });
    }
}

document.addEventListener('click', (e) => {
    const customLangSelector = document.getElementById('customLangSelector');
    if (customLangSelector && !customLangSelector.contains(e.target)) {
        customLangSelector.classList.remove('open');
    }
});

// Native hidden select fallback
const langSelect = document.getElementById('langSelect');
if (langSelect) {
    langSelect.addEventListener('change', (e) => setLanguage(e.target.value));
}

// Mobile Menu
const menuToggle = document.getElementById('menuToggle');
const navLinks = document.getElementById('navLinks');

if (menuToggle && navLinks) {
    menuToggle.addEventListener('click', () => {
        navLinks.classList.toggle('active');
        menuToggle.classList.toggle('active');
    });

    document.addEventListener('click', (e) => {
        if (!menuToggle.contains(e.target) && !navLinks.contains(e.target)) {
            navLinks.classList.remove('active');
            menuToggle.classList.remove('active');
        }
    });

    navLinks.querySelectorAll('a').forEach(link => {
        link.addEventListener('click', () => {
            navLinks.classList.remove('active');
            menuToggle.classList.remove('active');
        });
    });
}

// Header Scroll Effect
const header = document.querySelector('header');
if (header) {
    const scrollHandler = () => {
        if (window.scrollY > 20) header.classList.add('scrolled');
        else header.classList.remove('scrolled');
    };
    window.addEventListener('scroll', scrollHandler);
    scrollHandler();
}

// Copy Code Helper
window.copyCode = function (btn, text) {
    navigator.clipboard.writeText(text).then(() => {
        const originalHtml = btn.innerHTML;
        btn.innerHTML = '<i class="fas fa-check"></i> Copied!';
        btn.style.borderColor = '#00ff9f';
        btn.style.color = '#00ff9f';
        setTimeout(() => {
            btn.innerHTML = originalHtml;
            btn.style.borderColor = '';
            btn.style.color = '';
        }, 2000);
    }).catch(err => console.error('Failed to copy:', err));
};

// GitHub Stats Fetcher
window.fetchStats = async function () {
    try {
        const [repoRes, relRes] = await Promise.all([
            fetch('https://api.github.com/repos/tawroot/antigravity-cleaner'),
            fetch('https://api.github.com/repos/tawroot/antigravity-cleaner/releases')
        ]);
        const data = await repoRes.json();
        const releases = await relRes.json();

        const starCount = document.getElementById('star-count');
        const forkCount = document.getElementById('fork-count');
        const dlCount = document.getElementById('download-count');

        if (starCount && data.stargazers_count) {
            starCount.innerText = data.stargazers_count.toLocaleString();
        }
        document.querySelectorAll('.github-stars-val').forEach(el => {
            if (data.stargazers_count) el.innerText = data.stargazers_count.toLocaleString();
        });
        if (forkCount && data.forks_count) {
            forkCount.innerText = data.forks_count.toLocaleString();
        }

        if (dlCount) {
            if (Array.isArray(releases)) {
                let totalDl = 0;
                releases.forEach(r => (r.assets || []).forEach(a => totalDl += (a.download_count || 0)));
                if (totalDl > 0) {
                    dlCount.innerText = totalDl > 4900 ? totalDl.toLocaleString() + '+' : '+4.9K';
                } else {
                    dlCount.innerText = '+4.9K';
                }
            } else {
                dlCount.innerText = '+4.9K';
            }
        }
    } catch (e) {
        console.error('Failed to fetch GitHub stats:', e);
        const dlCount = document.getElementById('download-count');
        if (dlCount) dlCount.innerText = '+4.9K';
    }
};

document.addEventListener('DOMContentLoaded', () => {
    initThemeToggle();
    initCustomLangDropdown();
    const savedLang = localStorage.getItem('ag_lang') || 'en';
    setLanguage(savedLang);
    if (document.getElementById('star-count')) fetchStats();
    initHybridNavigation();
});

// ========================================
// HYBRID NAVIGATION SYSTEM (PJAX)
// ========================================

function initHybridNavigation() {
    const mainContent = document.querySelector('body');
    if (!mainContent) return;

    let currentPage = window.location.pathname;

    document.addEventListener('click', (e) => {
        const link = e.target.closest('a');

        if (!link) return;
        if (link.target === '_blank') return;
        if (link.href.startsWith('http') && !link.href.includes(window.location.host)) return;
        if (link.href.startsWith('#')) return;
        if (link.href.startsWith('mailto:')) return;
        if (link.href.startsWith('tel:')) return;

        const url = new URL(link.href);
        if (!url.pathname.endsWith('.html') && url.pathname !== '/') return;

        e.preventDefault();
        if (url.pathname === currentPage) return;

        loadPage(link.href);
    });

    window.addEventListener('popstate', (e) => {
        if (e.state && e.state.path) {
            loadPage(e.state.path, false);
        } else {
            loadPage(window.location.pathname, false);
        }
    });
}

async function loadPage(url, pushState = true) {
    try {
        document.body.classList.add('page-transitioning');

        const response = await fetch(url);
        if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);

        const html = await response.text();
        const parser = new DOMParser();
        const doc = parser.parseFromString(html, 'text/html');

        const newBody = doc.querySelector('body');
        const newTitle = doc.querySelector('title');

        if (!newBody) throw new Error('No body found in response');

        document.body.innerHTML = newBody.innerHTML;
        document.body.className = newBody.className;

        // Maintain user theme across page loads
        const currentTheme = localStorage.getItem('ag_theme') || 'dark';
        document.documentElement.setAttribute('data-theme', currentTheme);
        document.body.setAttribute('data-theme', currentTheme);

        if (newTitle) {
            document.title = newTitle.textContent;
        }

        if (pushState) {
            history.pushState({ path: url }, '', url);
        }

        reinitializeScripts();

        document.body.style.opacity = '0';
        requestAnimationFrame(() => {
            document.body.style.transition = 'opacity 0.3s ease';
            document.body.style.opacity = '1';
            document.body.style.pointerEvents = '';
        });

        currentPage = new URL(url, window.location.origin).pathname;

    } catch (error) {
        console.error('Navigation error:', error);
        window.location.href = url;
    }
}

function reinitializeScripts() {
    initThemeToggle();
    initCustomLangDropdown();

    const newLangSelect = document.getElementById('langSelect');
    if (newLangSelect) {
        newLangSelect.addEventListener('change', (e) => setLanguage(e.target.value));
    }
    const savedLang = localStorage.getItem('ag_lang') || 'en';
    if (newLangSelect) newLangSelect.value = savedLang;
    setLanguage(savedLang);

    const newMenuToggle = document.getElementById('menuToggle');
    const newNavLinks = document.getElementById('navLinks');
    if (newMenuToggle && newNavLinks) {
        newMenuToggle.onclick = function () {
            newNavLinks.classList.toggle('active');
            newMenuToggle.classList.toggle('active');
        };

        newNavLinks.querySelectorAll('a').forEach(link => {
            link.onclick = function () {
                newNavLinks.classList.remove('active');
                newMenuToggle.classList.remove('active');
            };
        });
    }

    const newHeader = document.querySelector('header');
    if (newHeader) {
        const scrollHandler = () => {
            if (window.scrollY > 20) newHeader.classList.add('scrolled');
            else newHeader.classList.remove('scrolled');
        };
        window.addEventListener('scroll', scrollHandler);
        scrollHandler();
    }

    if (document.getElementById('star-count')) {
        fetchStats();
    }

    updateActiveNavLink();
}

function updateActiveNavLink() {
    const currentPath = window.location.pathname;
    const currentFile = currentPath.split('/').pop() || 'index.html';

    document.querySelectorAll('.nav-links a').forEach(link => {
        link.classList.remove('active');
        const linkPath = new URL(link.href).pathname;
        const linkFile = linkPath.split('/').pop();

        if (linkFile === currentFile || (currentFile === '' && linkFile === 'index.html')) {
            link.classList.add('active');
        }
    });
}
