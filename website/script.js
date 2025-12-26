// Language Content
const translations = {
    en: {
        home: "Home", features: "Features", download: "Download", pricing: "Pricing", blog: "Blog", faq: "FAQ", get_app: "Get App",
        hero_title: "Fix Antigravity IDE Instantly",
        hero_subtitle: "The most powerful maintenance tool for installation errors, network interference, and region restrictions.",
        get_started: "Get Started", view_source: "View Source", stars: "GitHub Stars", forks: "Forks", downloads: "Downloads", version: "Version",
        quick_install: "Quick Install", quick_install_desc: "Install instantly via terminal without leaving your keyboard.",
        footer_desc: "Supporting open internet and developer productivity.", product: "Product", support: "Support", docs: "Documentation",
        feat_clean_title: "One-Click Clean", feat_clean_desc: "Remove corrupted installation files, cache, and registry keys instantly.",
        feat_region_title: "Region Bypass", feat_region_desc: "Automatically fix region-locked errors for constrained networks.",
        feat_session_title: "Session Preservation", feat_session_desc: "Clean the system without losing your browser cookies or login sessions.",
        feat_network_title: "Network Reset", feat_network_desc: "Flush DNS and reset network adapters to solve connection timeouts.",
        feat_dryrun_title: "Dry Run Mode", feat_dryrun_desc: "Preview exact changes before applying them to your system.",
        feat_portable_title: "Portable", feat_portable_desc: "No installation required. Run directly from USB or any folder.",
        faq_safe_q: "Is it safe to use?", faq_safe_a: "Yes. Antigravity Cleaner specifically targets temporary files, caches, and invalid registry keys associated with known IDE errors. It does not touch your personal documents or project source codes.",
        faq_pass_q: "Will I lose my browser passwords?", faq_pass_a: "No. Our 'Session Preservation' technology ensures that critical files like Cookies and Login Data are backed up or excluded from the cleaning process, so you don't have to log in again.",
        faq_net_q: "How do I fix 'Network Error'?", faq_net_a: "Use the 'Network Reset' feature in the Tools menu. This will flush your DNS cache and reset the Winsock catalog, which often resolves connection issues caused by VPNs using the Tun/Tap adapter.",
        faq_mac_q: "Do you support Mac M1/M2?", faq_mac_a: "Yes, we have native builds for Apple Silicon (ARM64) as well as Intel Macs. Check the Download page for the correct version.",
        price_free_title: "Community", price_free_1: "Open Source Core", price_free_2: "GitHub Issues Support", price_free_3: "All Basic Features", download_free: "Download Free",
        price_popular: "POPULAR", price_pro_title: "Pro Network", price_pro_1: "Private V2Ray Servers", price_pro_2: "Low Latency Routing", price_pro_3: "Priority Support",
        price_ent_title: "Enterprise", price_custom: "Custom", price_ent_1: "Dedicated IPs", price_ent_2: "Fleet Management", price_ent_3: "24/7 SLA Support", contact_sales: "Contact Sales",
        blog_1_title: "Fixing SSL Handshake Errors", blog_1_desc: "A comprehensive guide to understanding and fixing SSL/TLS errors in restricted network environments.", read_more: "Read More",
        blog_2_title: "Version 4.0 Release Notes", blog_2_desc: "Celebrating the release of our biggest update yet, featuring new UI and faster cleaning engines.",
        blog_3_title: "Understanding Region Locks", blog_3_desc: "How software limits access based on geography and what you can do about it legally.",

        // Community & Support
        community_title: "Join Our Community", community_subtitle: "Get help, share ideas, and stay updated with the latest news.",
        star_title: "Star on GitHub", star_desc: "Show your support by starring our repository. It helps us grow!", star_btn: "Star Repository",
        telegram_title: "Telegram Channel", telegram_desc: "Join our Telegram channel for updates, tips, and community support.", telegram_btn: "Join Channel",
        donate_title: "Support Development", donate_desc: "Help us maintain and improve this free tool with a donation.", donate_btn: "Donate",

        // Features Preview
        why_choose: "Why Choose Antigravity Cleaner?", why_choose_desc: "Powerful features designed for developers and power users.", view_all_features: "View All Features",

        // CTA
        cta_title: "Ready to Fix Your IDE?", cta_desc: "Download now and solve installation errors in seconds.",
        download_now: "Download Now", view_on_github: "View on GitHub",

        // Campaign
        campaign_badge: "🚀 FUNDING CAMPAIGN", campaign_title: "Fund Antigravity Cleaner v5.0",
        campaign_desc: "Help us build the next generation with AI-powered cleaning, advanced automation, and enterprise features.",
        campaign_raised: "raised", campaign_goal: "Goal:", campaign_backers: "Backers", campaign_days: "Days Left", campaign_version: "Target Version",
        campaign_features_title: "What's Coming in v5.0?",
        feature_ai: "AI-Powered Cleaning", feature_ai_desc: "Smart detection and removal",
        feature_auto: "Auto-Scheduler", feature_auto_desc: "Automated maintenance tasks",
        feature_enterprise: "Enterprise Tools", feature_enterprise_desc: "Fleet management & reporting",
        feature_ui: "Modern GUI", feature_ui_desc: "Beautiful desktop interface",
        campaign_support: "Support v5.0", campaign_discuss: "Join Discussion",
        campaign_note: "💡 All donations go directly to development. We accept BTC, ETH, USDT, and more.",
        backer_perks_title: "🎁 Backer Benefits",
        perk_support: "Priority Support", perk_feature: "Feature Requests", perk_early: "Early Access to v5.0", perk_credit: "Credits in Release Notes"
    },
    fa: {
        home: "خانه", features: "ویژگی‌ها", download: "دانلود", pricing: "تعرفه‌ها", blog: "وبلاگ", faq: "سوالات", get_app: "دانلود برنامه",
        hero_title: "رفع مشکلات آنتی‌گرویتی، آنی و سریع",
        hero_subtitle: "قدرتمندترین ابزار مدیریت برای رفع خطاهای نصب، تداخل شبکه و محدودیت‌های منطقه‌ای.",
        get_started: "شروع کنید", view_source: "کد منبع", stars: "ستاره‌ها", forks: "فورک‌ها", downloads: "دانلودها", version: "نسخه",
        quick_install: "نصب سریع", quick_install_desc: "نصب فوری از طریق ترمینال بدون نیاز به دانلود دستی فایل.",
        footer_desc: "حمایت از اینترنت آزاد و بهره‌وری توسعه‌دهندگان.", product: "محصول", support: "پشتیبانی", docs: "مستندات",
        feat_clean_title: "پاکسازی تک‌کلیک", feat_clean_desc: "حذف فوری فایل‌های نصبی خراب، کش‌ها و کلیدهای رجیستری معیوب سیستم.",
        feat_region_title: "دور زدن تحریم", feat_region_desc: "رفع خودکار خطاهای Region Lock و محدودیت‌های جغرافیایی برای شبکه‌های محدود.",
        feat_session_title: "حفظ نشست‌ها", feat_session_desc: "پاکسازی عمیق سیستم بدون از دست دادن کوکی‌ها، تاریخچه و لاگین‌های مرورگر.",
        feat_network_title: "ریست شبکه", feat_network_desc: "فلاش DNS و ریست کامل آداپتورها برای حل مشکل تایم‌آوت و تداخل VPN.",
        feat_dryrun_title: "حالت تست", feat_dryrun_desc: "مشاهده و بررسی دقیق تغییرات احتمالی قبل از اعمال نهایی روی سیستم.",
        feat_portable_title: "نسخه پرتابل", feat_portable_desc: "بدون نیاز به نصب. اجرا به صورت مستقیم از روی فلش USB یا هر پوشه‌ای.",
        faq_safe_q: "آیا امن است؟", faq_safe_a: "بله. آنتی‌گرویتی فقط فایل‌های موقت و کلیدهای خراب مربوط به خطاهای IDE را هدف می‌گیرد و کاری به فایل‌های شخصی شما ندارد.",
        faq_pass_q: "آیا پسوردهای مرورگر پاک می‌شوند؟", faq_pass_a: "خیر. تکنولوژی «حفظ نشست» ما فایل‌های حساس مثل کوکی‌ها و دیتای لاگین را بک‌آپ گرفته یا مستثنی می‌کند.",
        faq_net_q: "چطور خطای Network Error را حل کنم؟", faq_net_a: "از گزینه Network Reset در منوی Tools استفاده کنید. این کار DNS را فلاش کرده و کاتالوگ Winsock را ریست می‌کند.",
        faq_mac_q: "آیا مک M1/M2 پشتیبانی می‌شود؟", faq_mac_a: "بله، ما بیلد نیتیو (Native) برای پردازنده‌های اپل سیلیکون (ARM64) و همچنین اینتل داریم.",
        price_free_title: "جامعه", price_free_1: "هسته متن‌باز", price_free_2: "پشتیبانی در گیت‌هاب", price_free_3: "تمام امکانات پایه", download_free: "دانلود رایگان",
        price_popular: "محبوب", price_pro_title: "شبکه حرفه‌ای", price_pro_1: "سرورهای اختصاصی V2Ray", price_pro_2: "روتینگ با تاخیر کم", price_pro_3: "پشتیبانی اولویت‌دار",
        price_ent_title: "سازمانی", price_custom: "توافقی", price_ent_1: "آی‌پی اختصاصی", price_ent_2: "مدیریت گروهی کاربران", price_ent_3: "پشتیبانی ۲۴/۷ SLA", contact_sales: "تماس با فروش",
        blog_1_title: "حل خطاهای SSL Handshake", blog_1_desc: "راهنمای جامع برای رفع خطاهای SSL/TLS در محیط‌های شبکه محدود.", read_more: "ادامه مطلب",
        blog_2_title: "یادداشت‌های انتشار نسخه ۴.۰", blog_2_desc: "انتشار بزرگترین آپدیت ما با رابط کاربری جدید و موتور پاکسازی سریع‌تر.",
        blog_3_title: "درک قفل‌های منطقه‌ای", blog_3_desc: "چگونه نرم‌افزارها دسترسی را محدود می‌کنند و راه حل‌ها چیست.",

        // Community & Support
        community_title: "به جامعه ما بپیوندید", community_subtitle: "کمک بگیرید، ایده‌ها را به اشتراک بگذارید و از آخرین اخبار مطلع شوید.",
        star_title: "ستاره در گیت‌هاب", star_desc: "با ستاره دادن به مخزن ما، حمایت خود را نشان دهید. این به رشد ما کمک می‌کند!", star_btn: "ستاره دادن",
        telegram_title: "کانال تلگرام", telegram_desc: "برای دریافت به‌روزرسانی‌ها، نکات و پشتیبانی جامعه به کانال تلگرام ما بپیوندید.", telegram_btn: "عضویت در کانال",
        donate_title: "حمایت از توسعه", donate_desc: "با کمک مالی به ما در نگهداری و بهبود این ابزار رایگان کمک کنید.", donate_btn: "حمایت مالی",

        // Features Preview
        why_choose: "چرا آنتی‌گرویتی کلینر؟", why_choose_desc: "امکانات قدرتمند طراحی شده برای توسعه‌دهندگان و کاربران حرفه‌ای.", view_all_features: "مشاهده تمام ویژگی‌ها",

        // CTA
        cta_title: "آماده رفع مشکلات IDE خود هستید؟", cta_desc: "همین حالا دانلود کنید و خطاهای نصب را در چند ثانیه حل کنید.",
        download_now: "دانلود فوری", view_on_github: "مشاهده در گیت‌هاب",

        // Campaign
        campaign_badge: "🚀 کمپین تامین مالی", campaign_title: "حمایت از نسخه ۵.۰ آنتی‌گرویتی",
        campaign_desc: "به ما کمک کنید نسل بعدی را با پاکسازی هوشمند AI، اتوماسیون پیشرفته و امکانات سازمانی بسازیم.",
        campaign_raised: "جمع‌آوری شده", campaign_goal: "هدف:", campaign_backers: "حامیان", campaign_days: "روز باقی‌مانده", campaign_version: "نسخه هدف",
        campaign_features_title: "چه چیزی در نسخه ۵.۰ می‌آید؟",
        feature_ai: "پاکسازی هوشمند AI", feature_ai_desc: "تشخیص و حذف هوشمند",
        feature_auto: "زمان‌بند خودکار", feature_auto_desc: "وظایف نگهداری خودکار",
        feature_enterprise: "ابزارهای سازمانی", feature_enterprise_desc: "مدیریت ناوگان و گزارش‌دهی",
        feature_ui: "رابط گرافیکی مدرن", feature_ui_desc: "رابط کاربری زیبا",
        campaign_support: "حمایت از نسخه ۵.۰", campaign_discuss: "پیوستن به بحث",
        campaign_note: "💡 تمام کمک‌ها مستقیماً صرف توسعه می‌شود. ما BTC، ETH، USDT و سایر ارزها را قبول می‌کنیم.",
        backer_perks_title: "🎁 مزایای حامیان",
        perk_support: "پشتیبانی اولویت‌دار", perk_feature: "درخواست ویژگی", perk_early: "دسترسی زودهنگام به نسخه ۵.۰", perk_credit: "نام در یادداشت‌های انتشار"
    },
    ru: {
        home: "Главная", features: "Функции", download: "Скачать", pricing: "Цены", blog: "Блог", faq: "FAQ",
        hero_title: "Исправьте ошибки Antigravity мгновенно",
        hero_subtitle: "Самый мощный инструмент для обслуживания ошибок установки и сетевых помех."
    },
    zh: {
        home: "首页", features: "功能", download: "下载", pricing: "价格", blog: "博客", faq: "常见问题",
        hero_title: "立即修复 Antigravity IDE",
        hero_subtitle: "最强大的安装错误和网络干扰维护工具。"
    }
};

const langSelect = document.getElementById('langSelect');
const menuToggle = document.getElementById('menuToggle');
const navLinks = document.getElementById('navLinks');
const header = document.querySelector('header');

function setLanguage(lang) {
    document.documentElement.lang = lang;
    const rtlLanguages = ['fa', 'ar', 'ur'];
    document.documentElement.dir = rtlLanguages.includes(lang) ? 'rtl' : 'ltr';

    document.querySelectorAll('[data-lang-key]').forEach(elem => {
        const key = elem.getAttribute('data-lang-key');
        if (translations[lang] && translations[lang][key]) {
            elem.innerText = translations[lang][key];
        } else if (translations['en'][key]) {
            elem.innerText = translations['en'][key];
        }
    });

    localStorage.setItem('ag_lang', lang);
    if (langSelect) langSelect.value = lang;
}

if (langSelect) {
    langSelect.addEventListener('change', (e) => setLanguage(e.target.value));
    const savedLang = localStorage.getItem('ag_lang') || (navigator.language.startsWith('fa') ? 'fa' : 'en');
    setLanguage(savedLang);
}

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

if (header) {
    window.addEventListener('scroll', () => {
        if (window.scrollY > 20) header.classList.add('scrolled');
        else header.classList.remove('scrolled');
    });
}

window.toggleTerminal = function (type) {
    const win = document.getElementById('cmdWin');
    const unix = document.getElementById('cmdUnix');
    const tabWin = document.getElementById('tabWin');
    const tabUnix = document.getElementById('tabUnix');
    if (!win || !unix || !tabWin || !tabUnix) return;

    if (type === 'win') {
        win.style.display = 'block';
        unix.style.display = 'none';
        tabWin.classList.add('active');
        tabUnix.classList.remove('active');
    } else {
        win.style.display = 'none';
        unix.style.display = 'block';
        tabWin.classList.remove('active');
        tabUnix.classList.add('active');
    }
};

window.copyTerminal = function () {
    const activeTab = document.querySelector('.t-tab.active');
    if (!activeTab) return;

    const text = activeTab.id === 'tabWin'
        ? document.getElementById('winTxt')?.innerText
        : document.getElementById('unixTxt')?.innerText;
    if (!text) return;

    navigator.clipboard.writeText(text).then(() => {
        const btn = document.querySelector('.t-copy-btn');
        if (!btn) return;
        const original = btn.innerHTML;
        btn.innerHTML = '<i class="fas fa-check"></i>';
        btn.style.background = 'rgba(16, 185, 129, 0.2)';
        setTimeout(() => {
            btn.innerHTML = original;
            btn.style.background = '';
        }, 2000);
    }).catch(err => console.error('Failed to copy:', err));
};

window.fetchStats = async function () {
    try {
        const res = await fetch('https://api.github.com/repos/tawroot/antigravity-cleaner');
        const data = await res.json();

        const starCount = document.getElementById('star-count');
        const forkCount = document.getElementById('fork-count');

        if (starCount && data.stargazers_count) {
            starCount.innerText = data.stargazers_count.toLocaleString();
        }
        if (forkCount && data.forks_count) {
            forkCount.innerText = data.forks_count.toLocaleString();
        }
    } catch (e) {
        console.error('Failed to fetch GitHub stats:', e);
    }
};

// ========================================
// CAMPAIGN PROGRESS UPDATER
// ========================================

// Manual update function - call this when you receive donations
window.updateCampaignProgress = function (raised, backers, daysLeft) {
    const goal = 500; // $500 goal
    const percentage = Math.min(Math.round((raised / goal) * 100), 100);

    // Update raised amount
    const raisedEl = document.getElementById('campaign-raised');
    if (raisedEl) raisedEl.innerText = `$${raised.toLocaleString()}`;

    // Update progress bar
    const progressEl = document.getElementById('campaign-progress');
    const percentageEl = document.getElementById('campaign-percentage');
    if (progressEl) progressEl.style.width = `${percentage}%`;
    if (percentageEl) percentageEl.innerText = `${percentage}%`;

    // Update backers
    const backersEl = document.getElementById('campaign-backers');
    if (backersEl) backersEl.innerText = backers;

    // Update days left
    const daysEl = document.getElementById('campaign-days');
    if (daysEl) daysEl.innerText = daysLeft;

    // Save to localStorage
    localStorage.setItem('campaign_data', JSON.stringify({ raised, backers, daysLeft, updated: Date.now() }));
};

// Load saved campaign data on page load
function loadCampaignData() {
    const saved = localStorage.getItem('campaign_data');
    if (saved) {
        try {
            const data = JSON.parse(saved);
            updateCampaignProgress(data.raised || 0, data.backers || 0, data.daysLeft || 30);
        } catch (e) {
            console.error('Failed to load campaign data:', e);
        }
    }
}

document.addEventListener('DOMContentLoaded', () => {
    if (document.getElementById('star-count')) fetchStats();
    loadCampaignData(); // Load campaign progress
    initHybridNavigation();
});

// ========================================
// HYBRID NAVIGATION SYSTEM (PJAX)
// ========================================

function initHybridNavigation() {
    // Get main content container
    const mainContent = document.querySelector('body');
    if (!mainContent) return;

    // Track current page for analytics
    let currentPage = window.location.pathname;

    // Handle all internal navigation links
    document.addEventListener('click', (e) => {
        const link = e.target.closest('a');

        // Check if it's an internal navigation link
        if (!link) return;
        if (link.target === '_blank') return;
        if (link.href.startsWith('http') && !link.href.includes(window.location.host)) return;
        if (link.href.startsWith('#')) return;
        if (link.href.startsWith('mailto:')) return;
        if (link.href.startsWith('tel:')) return;

        // Check if it's a local HTML file
        const url = new URL(link.href);
        if (!url.pathname.endsWith('.html') && url.pathname !== '/') return;

        // Prevent default navigation
        e.preventDefault();

        // Don't reload if clicking current page
        if (url.pathname === currentPage) return;

        // Load new page with smooth transition
        loadPage(link.href);
    });

    // Handle browser back/forward buttons
    window.addEventListener('popstate', (e) => {
        if (e.state && e.state.path) {
            loadPage(e.state.path, false);
        }
    });

    // Save initial state
    history.replaceState({ path: window.location.pathname }, '', window.location.pathname);
}

async function loadPage(url, pushState = true) {
    try {
        // Show loading state
        document.body.style.opacity = '0.7';
        document.body.style.pointerEvents = 'none';

        // Fetch new page
        const response = await fetch(url);
        if (!response.ok) throw new Error('Page not found');

        const html = await response.text();

        // Parse the HTML
        const parser = new DOMParser();
        const newDoc = parser.parseFromString(html, 'text/html');

        // Extract new content
        const newBody = newDoc.querySelector('body');
        const newTitle = newDoc.querySelector('title');

        if (!newBody) throw new Error('Invalid page structure');

        // Smooth fade out
        await new Promise(resolve => {
            document.body.style.transition = 'opacity 0.2s ease';
            document.body.style.opacity = '0';
            setTimeout(resolve, 200);
        });

        // Scroll to top
        window.scrollTo({ top: 0, behavior: 'instant' });

        // Replace content
        document.body.innerHTML = newBody.innerHTML;
        document.body.className = newBody.className;
        document.body.setAttribute('data-theme', newBody.getAttribute('data-theme') || 'dark');

        // Update title
        if (newTitle) {
            document.title = newTitle.textContent;
        }

        // Update URL
        if (pushState) {
            history.pushState({ path: url }, '', url);
        }

        // Re-initialize all scripts
        reinitializeScripts();

        // Smooth fade in
        document.body.style.opacity = '0';
        requestAnimationFrame(() => {
            document.body.style.transition = 'opacity 0.3s ease';
            document.body.style.opacity = '1';
            document.body.style.pointerEvents = '';
        });

        // Update current page
        currentPage = new URL(url, window.location.origin).pathname;

    } catch (error) {
        console.error('Navigation error:', error);
        // Fallback to normal navigation
        window.location.href = url;
    }
}

function reinitializeScripts() {
    // Re-attach language selector
    const newLangSelect = document.getElementById('langSelect');
    if (newLangSelect) {
        newLangSelect.addEventListener('change', (e) => setLanguage(e.target.value));
        const savedLang = localStorage.getItem('ag_lang') || 'en';
        newLangSelect.value = savedLang;
        setLanguage(savedLang);
    }

    // Re-attach mobile menu
    const newMenuToggle = document.getElementById('menuToggle');
    const newNavLinks = document.getElementById('navLinks');
    if (newMenuToggle && newNavLinks) {
        newMenuToggle.addEventListener('click', () => {
            newNavLinks.classList.toggle('active');
            newMenuToggle.classList.toggle('active');
        });

        document.addEventListener('click', (e) => {
            if (!newMenuToggle.contains(e.target) && !newNavLinks.contains(e.target)) {
                newNavLinks.classList.remove('active');
                newMenuToggle.classList.remove('active');
            }
        });

        newNavLinks.querySelectorAll('a').forEach(link => {
            link.addEventListener('click', () => {
                newNavLinks.classList.remove('active');
                newMenuToggle.classList.remove('active');
            });
        });
    }

    // Re-attach scroll handler for header
    const newHeader = document.querySelector('header');
    if (newHeader) {
        const scrollHandler = () => {
            if (window.scrollY > 20) newHeader.classList.add('scrolled');
            else newHeader.classList.remove('scrolled');
        };
        window.addEventListener('scroll', scrollHandler);
        scrollHandler(); // Initial check
    }

    // Fetch GitHub stats if on homepage
    if (document.getElementById('star-count')) {
        fetchStats();
    }

    // Update active nav link
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
