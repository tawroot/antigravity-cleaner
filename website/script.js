/**
 * Antigravity Cleaner — VitePress / Marzban Interactive Script
 */

(function () {
  'use strict';

  // 1. Theme Management (Light mode default)
  const root = document.documentElement;
  const themeToggle = document.getElementById('themeToggle');
  const THEME_KEY = 'vp_theme_mode';

  function initTheme() {
    const urlParams = new URLSearchParams(window.location.search);
    const urlTheme = urlParams.get('theme');
    const saved = urlTheme || localStorage.getItem(THEME_KEY);
    if (saved === 'dark') {
      root.setAttribute('data-theme', 'dark');
    } else {
      root.setAttribute('data-theme', 'light');
    }
  }

  function toggleTheme() {
    const isDark = root.getAttribute('data-theme') === 'dark';
    const nextTheme = isDark ? 'light' : 'dark';
    root.setAttribute('data-theme', nextTheme);
    localStorage.setItem(THEME_KEY, nextTheme);
  }

  if (themeToggle) {
    themeToggle.addEventListener('click', toggleTheme);
  }
  initTheme();

  // 2. View Switcher (Home vs Docs)
  const viewHome = document.getElementById('view-home');
  const viewDocs = document.getElementById('view-docs');
  const navHomeBtn = document.getElementById('navHomeBtn');
  const navDocsBtn = document.getElementById('navDocsBtn');
  const brandLink = document.getElementById('brandLink');

  function showView(viewName, targetId) {
    if (viewName === 'docs') {
      viewHome.classList.remove('active');
      viewDocs.classList.add('active');
      if (navHomeBtn) navHomeBtn.classList.remove('active');
      if (navDocsBtn) navDocsBtn.classList.add('active');
      window.scrollTo(0, 0);

      if (targetId) {
        setTimeout(() => {
          const isFa = root.getAttribute('lang') === 'fa';
          const faTarget = targetId.startsWith('fa-') ? targetId : ('fa-' + targetId);
          const el = (isFa && document.getElementById(faTarget)) 
            ? document.getElementById(faTarget) 
            : (document.getElementById(targetId) || document.getElementById(targetId.replace('fa-', '')));
          if (el) {
            el.scrollIntoView({ behavior: 'auto' });
          }
        }, 30);
      } else {
        window.scrollTo({ top: 0, behavior: 'smooth' });
      }
    } else {
      viewDocs.classList.remove('active');
      viewHome.classList.add('active');
      if (navDocsBtn) navDocsBtn.classList.remove('active');
      if (navHomeBtn) navHomeBtn.classList.add('active');
      window.scrollTo({ top: 0, behavior: 'smooth' });
    }
  }

  if (navHomeBtn) navHomeBtn.addEventListener('click', (e) => { e.preventDefault(); showView('home'); });
  if (brandLink) brandLink.addEventListener('click', (e) => { e.preventDefault(); showView('home'); });
  if (navDocsBtn) navDocsBtn.addEventListener('click', (e) => { e.preventDefault(); showView('docs'); });

  // Global Anchor navigation for smooth view switching (Header, Hero, Footer, Docs)
  document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', (e) => {
      const hash = anchor.getAttribute('href');
      if (hash === '#home') {
        e.preventDefault();
        showView('home');
      } else if (hash === '#download') {
        e.preventDefault();
        showView('home');
        setTimeout(() => {
          const el = document.getElementById('download');
          if (el) el.scrollIntoView({ behavior: 'smooth' });
        }, 50);
      } else if (hash === '#quick-install') {
        e.preventDefault();
        showView('home');
        setTimeout(() => {
          const el = document.getElementById('quick-install');
          if (el) el.scrollIntoView({ behavior: 'smooth' });
        }, 50);
      } else if (hash === '#docs') {
        e.preventDefault();
        showView('docs');
        } else if (hash.startsWith('#doc-') || hash.startsWith('#fa-') || hash === '#gui-guide' || hash === '#cli-guide' || hash === '#benchmarks') {
        e.preventDefault();
        const target = (hash === '#benchmarks') ? 'doc-benchmarks' : hash.replace('#', '');
        showView('docs', target);
      }
    });
  });

  // Handle URL hash changes
  function checkHash() {
    const hash = window.location.hash;
    if (hash === '#download') {
      showView('home');
      setTimeout(() => {
        const el = document.getElementById('download');
        if (el) el.scrollIntoView({ behavior: 'smooth' });
      }, 50);
    } else if (hash === '#quick-install') {
      showView('home');
      setTimeout(() => {
        const el = document.getElementById('quick-install');
        if (el) el.scrollIntoView({ behavior: 'smooth' });
      }, 50);
    } else if (hash === '#docs' || hash.startsWith('#doc-') || hash.startsWith('#fa-') || hash === '#gui-guide' || hash === '#cli-guide' || hash === '#benchmarks') {
      const target = (hash === '#docs' || hash === '#home') ? null : hash.replace('#', '');
      showView('docs', target);
    } else if (hash === '#home') {
      showView('home');
    }
  }
  window.addEventListener('hashchange', checkHash);

  // 3. Quick Terminal Tabs & Copy
  const tabCurl = document.getElementById('tab-curl');
  const tabPs = document.getElementById('tab-ps');
  const codeCurl = document.getElementById('code-curl');
  const codePs = document.getElementById('code-ps');
  const copyCodeBtn = document.getElementById('copyCodeBtn');

  if (tabCurl && tabPs) {
    tabCurl.addEventListener('click', () => {
      tabCurl.classList.add('active');
      tabPs.classList.remove('active');
      codeCurl.style.display = 'block';
      codePs.style.display = 'none';
    });

    tabPs.addEventListener('click', () => {
      tabPs.classList.add('active');
      tabCurl.classList.remove('active');
      codePs.style.display = 'block';
      codeCurl.style.display = 'none';
    });
  }

  // Pre-select Windows tab if on Windows
  if (navigator.userAgent && navigator.userAgent.toLowerCase().includes('windows') && tabPs) {
    tabPs.click();
  }

  if (copyCodeBtn) {
    copyCodeBtn.addEventListener('click', async () => {
      const activeCodeEl = (codePs && codePs.style.display !== 'none') ? codePs : codeCurl;
      const cleanCmd = activeCodeEl ? (activeCodeEl.getAttribute('data-raw') || activeCodeEl.textContent) : '';
      try {
        await navigator.clipboard.writeText(cleanCmd.trim());
        copyCodeBtn.classList.add('copied');
        const label = copyCodeBtn.querySelector('span');
        const orig = label ? label.textContent : 'Copy';
        if (label) label.textContent = 'Copied!';
        setTimeout(() => {
          copyCodeBtn.classList.remove('copied');
          if (label) label.textContent = orig;
        }, 2000);
      } catch (err) {
        console.error('Copy failed:', err);
      }
    });
  }

  // 4. Copy Code Buttons inside Documentation
  document.querySelectorAll('.copy-code-btn').forEach(btn => {
    btn.addEventListener('click', async () => {
      const pre = btn.closest('.vp-code-block').querySelector('pre code');
      if (pre) {
        try {
          await navigator.clipboard.writeText(pre.textContent.trim());
          const orig = btn.textContent;
          btn.textContent = 'Copied!';
          setTimeout(() => { btn.textContent = orig; }, 2000);
        } catch (err) {
          console.error('Copy failed:', err);
        }
      }
    });
  });

  // 5. Docs Sidebar Links & TOC ScrollSpy
  window.addEventListener('scroll', () => {
    if (!viewDocs.classList.contains('active')) return;

    const isFa = root.getAttribute('lang') === 'fa';
    const activeHeadings = document.querySelectorAll(isFa ? '.doc-fa h1, .doc-fa h2' : '.doc-en h1, .doc-en h2');
    const activeTocLinks = document.querySelectorAll(isFa ? '.toc-fa .toc-link' : '.toc-en .toc-link');
    const activeSidebarLinks = document.querySelectorAll(isFa ? '.nav-fa .link' : '.nav-en .link');

    let activeId = '';
    activeHeadings.forEach(h => {
      const top = h.getBoundingClientRect().top;
      if (top <= 140) {
        activeId = h.id;
      }
    });

    if (activeId) {
      activeTocLinks.forEach(tl => {
        if (tl.getAttribute('href') === '#' + activeId) {
          tl.classList.add('active');
        } else {
          tl.classList.remove('active');
        }
      });

      activeSidebarLinks.forEach(sl => {
        if (sl.getAttribute('href') === '#' + activeId) {
          sl.classList.add('active');
        } else {
          sl.classList.remove('active');
        }
      });
    }
  });

  // 7. Multi-language (i18n) Engine
  const langToggleBtn = document.getElementById('langToggleBtn');
  const langDropdown = document.getElementById('langDropdown');
  const currentLangLabel = document.getElementById('currentLangLabel');
  const langOptions = document.querySelectorAll('.lang-option');
  const LANG_KEY = 'vp_lang_mode';

  let locales = {};

  async function loadLocale(lang) {
    if (locales[lang]) return locales[lang];
    try {
      const res = await fetch(`locales/${lang}.json`);
      if (res.ok) {
        locales[lang] = await res.json();
        return locales[lang];
      }
    } catch (e) {
      console.warn('Failed to load locale:', lang, e);
    }
    return null;
  }

  function getNestedTranslation(obj, path) {
    return path.split('.').reduce((acc, part) => acc && acc[part], obj);
  }

  async function setLanguage(lang) {
    const data = await loadLocale(lang);
    if (!data) return;

    // Set html lang and direction (Persian = RTL)
    root.setAttribute('lang', lang);
    if (lang === 'fa' || lang === 'ar') {
      root.setAttribute('dir', 'rtl');
    } else {
      root.setAttribute('dir', 'ltr');
    }

    if (currentLangLabel) {
      currentLangLabel.textContent = lang === 'fa' ? 'فارسی' : 'English';
    }

    langOptions.forEach(opt => {
      if (opt.getAttribute('data-lang') === lang) {
        opt.classList.add('active');
      } else {
        opt.classList.remove('active');
      }
    });

    // Translate all elements with data-i18n
    document.querySelectorAll('[data-i18n]').forEach(el => {
      const key = el.getAttribute('data-i18n');
      const val = getNestedTranslation(data, key);
      if (val) {
        el.textContent = val;
      }
    });

    localStorage.setItem(LANG_KEY, lang);
  }

  if (langToggleBtn && langDropdown) {
    langToggleBtn.addEventListener('click', (e) => {
      e.stopPropagation();
      langDropdown.classList.toggle('show');
    });

    document.addEventListener('click', () => {
      langDropdown.classList.remove('show');
    });

    langOptions.forEach(opt => {
      opt.addEventListener('click', () => {
        const selectedLang = opt.getAttribute('data-lang');
        setLanguage(selectedLang);
        langDropdown.classList.remove('show');
      });
    });
  }

  // Init language from URL param or saved preference
  const urlParams = new URLSearchParams(window.location.search);
  const urlLang = urlParams.get('lang');
  const savedLang = urlLang || localStorage.getItem(LANG_KEY) || 'en';
  if (savedLang !== 'en') {
    setLanguage(savedLang).then(() => {
      if (window.location.hash) {
        checkHash();
      }
    });
  } else {
    if (window.location.hash) {
      checkHash();
    }
  }

})();
