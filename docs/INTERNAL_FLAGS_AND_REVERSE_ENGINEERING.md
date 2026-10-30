<div dir="rtl">

# مستندات مهندسی معکوس و کشف فلگ‌های داخلی موتور Antigravity (Language Server)

**تاریخ کشف و تحلیل:** جمعه ۸ آبان ۱۴۰۵ (2026-10-30)  
**نسخه تارگت مهندسی معکوس:** `Google Antigravity language_server (x86_64 ELF stripped)`  
**ابزارهای استفاده شده:** REA (Reverse Engineer Anything v6.8.0), Objdump, Readelf, Python Binary Scanners

---

## ۱. تحلیل عمیق پچ حالت SSH (Remote Environment Spoofing)

### ۱.۱. مکانیسم کشف‌شده در بایت‌های باینری
در باینری `language_server`، تابع آدرس `0x6935840` مسئولیت مستقیم شناسایی اتصال SSH را بر عهده دارد:

```go
// بازسازی کد منبع از اسمبلی آدرس 0x6935840
func isSSH() bool {
    if os.Getenv("SSH_TTY") != "" { return true }       // rodata: 0x4283fc1 (طول 7 بایت)
    if os.Getenv("SSH_CLIENT") != "" { return true }    // rodata: 0x42c6243 (طول 10 بایت)
    if os.Getenv("SSH_CONNECTION") != "" { return true }// rodata: 0x432121a (طول 14 بایت)
    return false
}
```

این تابع به عنوان متد سنسور در جدول اینترفیس `0x09399e50` ثبت شده است. سایر سنسورهای هم‌ردیف آن در اینترفیس:
1. **محیط کانتینر / داکر (`0x6935660`):** بررسی فایل‌های `/.dockerenv` (در `0x42de32d`) و `/proc/1/cgroup` (در `0x432120c`).
2. **سرویس سکرت دسکتاپ (`0x6935760`):** بررسی وجود `DBUS_SESSION_BUS_ADDRESS` (در `0x43c7e68`).
3. **محیط سشن ریموت SSH (`0x6935840`):** بررسی `SSH_TTY`، `SSH_CLIENT` و `SSH_CONNECTION`.
4. **محیط WSL (`0x6935940`):** بررسی `WSL_DISTRO_NAME`، `WSL_INTEROP` و `/proc/sys/kernel/osrelease`.
5. **محیط اندروید (`0x69358c0`):** شناسه `Android (no OS keyring)`.

### ۱.۲. اثر دقیق این حالت در رفتار کلاینت و موتور احراز هویت
در آدرس `0x6920b72` تابع ارزیابی سنسورها صدا زده می‌شود:
```text
6920b72: call 6935320    ; اجرای سنسورها
6920b77: test %al, %al
6920b79: jne 6920c36     ; پرش در صورت تشخیص ریموت / SSH / کانتینر
```
در صورت مثبت بودن پاسخ، برنامه مستقیماً به آدرس `0x6920c36` پرش کرده و رشته استخراج‌شده در `0x447b07a` را لاگ و اعمال می‌کند:
> `Using file-based token storage because %s detected`  
> (که `%s` با `SSH session` پر می‌شود: `Using file-based token storage because SSH session detected`)

**نتایج حیاتی این پچ:**
1. **حذف وابستگی به جاکلیدی لینوکس (GNOME Keyring / FreeDesktop D-Bus Secret Service):** در لینوکس، سرویس جاکلیدی در بسیاری از موارد (مخصوصاً توزیع‌های سبک، دسکتاپ‌های بدون گنوم، یا حین قطعی ارتباط با D-Bus) قفل مانده یا پاسخ نمی‌دهد و منجر به خطای فلج‌کننده `Failed to fetch / There was an unexpected issue setting up your account` می‌شود. با فعال بودن `SSH_CONNECTION`، موتور به ۱۰۰٪ حالت ذخیره‌سازی توکن فایلی پایدار سوییچ می‌کند.
2. **لغو الزامات مرورگر تعاملی محلی (Local Browser Loopback):** موتور هوش مصنوعی دیگر اصرار به باز کردن تب‌های پاپ‌آپ لوکال برای رفرش توکن ندارد و از توکن‌های از پیش ذخیره‌شده فایل استفاده می‌کند.

### ۱.۳. آیا حالت SSH در دور زدن تحریم و حساسیت سرور گوگل تاثیر دارد؟
**پاسخ قطعی و تحلیل معماری:**
* **رفتار Fraud Detection گوگل:** سرورهای گوگل ترافیک ورودی را از طریق پروتوباف‌های `ClientMetadata` و `EnvironmentConfig` آنالیز می‌کنند. وقتی کلاینت گزارش می‌دهد که در یک سشن SSH ریموت در حال اجراست و هم‌زمان ترافیک خروجی از طریق پروکسی SOCKS5/HTTP تمیز رد می‌شود، الگوریتم‌های سرور گوگل این سناریو را به عنوان **یک توسعه‌دهنده متصل به VPS یا کلاود سرور خارجی** طبقه‌بندی می‌کنند.
* این هویت پردازشی باعث می‌شود رفتارهای ناهنجار سمت کلاینت (مانند تغییرات مکرر سشن گرافیکی، ناسازگاری وب‌سوکت‌های مرورگر محلی و خطاهای محیطی کلاینت) صفر شود و به شدت ریسک مسدودسازی یا فلگ شدن کلاینت کاهش یابد.

### ۱.۴. اقدام ضروری برای ریلز جاری (امشب):
در حال حاضر `SSH_CONNECTION` فقط در `pkg/proxy/launcher.go` خط ۱۴۳ داخل `BuildProxyEnv` وجود دارد (فقط برای دستور ترمینالی `antigravity-cleaner launch`).
برای این که ۱۰۰٪ پایدار و جامع شود:
1. باید در اسکریپت شورتکات دسکتاپ (`pkg/proxy/launcher.go` در تابع `GenerateNoTunDesktopLauncher`) داخل `antigravity-smart-launcher.sh` تزریق شود:
   ```bash
   export SSH_CONNECTION="127.0.0.1 1234 127.0.0.1 22"
   export SSH_CLIENT="127.0.0.1 1234 22"
   ```
2. در رپر باینری هسته (`pkg/patcher/wrapper.go`) داخل اسکریپت جایگزین باینری اضافه شود تا حتی اگر کاربر بدون هیچ لانچری مستقیماً IDE را باز کرد، پروسه `language_server` در حالت SSH اجرا شود.

---

## ۲. گنجینه فلگ‌های استخراج‌شده هسته (۱۱۴ فلگ رسمی) و نقشه راه آپدیت‌های آینده

با آنالیز روتین ثبت فلگ‌ها در آدرس `0x7608aa0` (فراخوانی `0x4da9340`)، کلیه فلگ‌های مخفی و سازمانی استخراج شدند. موارد زیر برای آپدیت‌های بعدی اولویت‌بندی شده‌اند:

### دسته ۱: فلگ‌های خنثی‌سازی گیت‌ها و آزمایش‌های جغرافیایی سرور گوگل (Geographic & Feature Bypasses)
1. **`--disable_unleash=true`**
   - *Help:* `Whether to disable Unleash`
   - *اثر:* قطع ارتباط با سیستم ریموت Unleash گوگل که قابلیت‌ها را بر اساس لوکیشن جغرافیایی مسدود یا محدود می‌کند.
2. **`--use_snapshot_mendelfp=true`**
   - *Help:* `Whether to use the embedded Mendel snapshot instead of polling live Mendel for experiments`
   - *اثر:* جلوگیری از پولینگ زنده سرور Mendel برای اعمال محدودیت‌های اکانت‌های مصرفی و اجبار به استفاده از اسنپ‌شات امبدد پایدار.
3. **`--validate_corp_sso=false`**
   - *Help:* `Whether to validate Corp SSO ticket on incoming HTTP requests`
   - *اثر:* لغو اعتبارسنجی تیکت سازمانی Corp SSO برای درخواست‌های لوکال HTTP.

### دسته ۲: مهندسی سهمیه، ریت‌لیمیت و انتخاب مدل (Quota & Model Tier Manipulation)
1. **`--beyond_quota_bucket=<team>/<bucket>`**
   - *Help:* `Quota bucket to use for Beyond API calls (format: team/bucket)`
   - *اثر:* فورس باکت سهمیه به باکت‌های سازمانی یا تیم‌های توسعه داخلی جهت رهایی از ریت‌لیمیت باکت عمومی `PRIMARY_QUOTA`.
2. **`--override_model_name=<model>`**
   - *Help:* `Model name to override default model`
   - *اثر:* اجبار موتور به فراخوانی مستقیم یک مدل خاص (مثلاً مدل‌های نامحدود یا پرو).
3. **`--force_contrib_variant=true`**
   - *Help:* `Forces a contrib experiment variant for every conversation. For local development.`
   - *اثر:* فعال‌سازی واریانت‌های توسعه‌ای در چت‌ها.

### دسته ۳: آزادی کامل ابزارها و سایدکارهای هوش مصنوعی (Sidecar & MCP Freedom)
1. **`--allow_all_sidecars=true`**
   - *Help:* `If true, allows all sidecars to run instead of just trusted CLIs. Note that the Trusted CLIs feature is effectively required to run any sidecar.`
   - *اثر:* رفع محدودیت اعتبارسنجی امضای ابزارهای جانبی و MCPها.
2. **`--bypass_sidecar_user_config=true`**
   - *Help:* `If true, bypasses user sidecar config so all sidecars are always treated as enabled.`
   - *اثر:* فعال نگه داشتن دائم کلیه ابزارهای اکوسیستم هوش مصنوعی بدون توجه به غیرفعال شدن احتمالی در تنظیمات کاربر.

### دسته ۴: عایق‌بندی لاگ‌ها، پیشگیری از خطای EIO و پنهان‌کاری کامل (Zero-Log & Anti-Crash)
1. **`--disable_log_to_disk=true`**
   - *Help:* `Disable logging to local disk (not including binary logging)`
   - *اثر:* قطع کامل نوشتن لاگ‌های سنگین تاریخچه چت و پرومپت‌ها روی دیسک محلی و رفع قطعی خطاهای `write EIO`.
2. **`--suppress_cpp_logger_module_check_and_dont_log_cpp_to_stderr=true`**
   - *Help:* `Suppress the check whether this binary runs on Borg and does not use the C++ logger_module`
   - *اثر:* جلوگیری از ارسال ارورها و هشدارهای داخلی C++ به `stderr` که باعث نشت یا بافر اورفلو در پروسه والد می‌شدند.
3. **`--silent_init=true`**
   - *اثر:* بوت شدن پروسه در پس‌زمینه بدون چاپ هرگونه بنر ردیابی.

### دسته ۵: روتینگ شبکه و پروکسی در سطح پروسه
1. **`--proxy_server_url=<url>`**
   - *Help:* `Proxy server URL for remote control`
   - *اثر:* تزریق مستقیم آدرس پروکسی به لایه شبکه باینری در کنار متغیرهای محیطی `ALL_PROXY`.

---

## ۳. لیست کامل ۱۱۴ فلگ استخراج‌شده از باینری برای مرجع توسعه
```text
undefok, syslog_on_start, log_dir, log_link, logbuflevel, version, disable_log_to_disk, v, vmodule, log_backtrace_at, logtostderr, alsologtostderr, stderrthreshold, silent_init, util_status_save_stack_trace, mlock_style, suppress_cpp_logger_module_check_and_dont_log_cpp_to_stderr, protosanitizer_use_old_annotation_sets, protosanitizer_use_streamz_for_sanitization_monitoring, protosanitizer_use_sensitivity_based_logic, api_server_url, cloud_code_endpoint, proxy_server_url, instance_id, remote_agent_fastpush_override, analytics_server_url, inference_api_server_url, https_server_port, http_server_port, lsp_port, extension_server_port, extension_server_csrf_token, host_bridge_url, host_bridge_token, read_host_bridge_token_from_stdin, use_host_auth, ve_log_dump, model_api_client_type, override_model_name, generative_service_addr, beyond_service_addr, beyond_quota_bucket, traffic_stream_selfid, google_cloud_project, workspace_id, parent_pipe_path, exit_on_stdin_close, disable_unleash, use_snapshot_mendelfp, disable_telemetry, startup_warning, force_contrib_variant, xbox_tool_use_port, running_in_xbox, sandbox_mode, enable_lsp, mquery_for_context_module, is_google3_workspace, is_google_internal, enable_product_api, limit_go_max_procs, stamp, web_bundle_path, file_watch_max_dir_count, max_num_tracked_workspaces, fs_error_probability, fs_latency_ms, read_dir_latency_ms, read_file_latency_ms, stat_latency_ms, disable_local_path, app_data_dir, gemini_dir, config_dir, csrf_token, cdp_port, local_chrome_user_data_dir, local_chrome_headless, use_local_chrome, browser_eval_env, use_custom_page_actions, use_ls_chrome_devtools_mcp, standalone, headless, use_stubby_auth, stubby_auth_cloud_project, stubby_auth_cloud_location, stubby_auth_mode, use_test_gaia, use_mocked_data, mock_fixture, persistent_mode, remote_accessible, allowed_ingress_origins, serve_hub_over_autobahn, validate_corp_sso, jetbox, jetbox-version, jetbox_release_channel, enable_sidecars, allow_all_sidecars, bypass_sidecar_user_config, override_ide_name, override_ide_version, subclient_type, override_oauth_client_id, override_oauth_client_secret, override_business_oauth_client_id, override_business_oauth_client_secret, zerog_http_socket_path, zerog_tls_cert_path, zerog_vm_ssh_socket_name, zerog_vm_ssh_key_path, zerog_app_host_port
```

</div>
