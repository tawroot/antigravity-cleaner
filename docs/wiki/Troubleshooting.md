# Troubleshooting

## Common Errors

### 403 Forbidden Error
**Problem:** Access denied when trying to use Antigravity IDE or Google services.

**Solution:**
1. Run Antigravity Cleaner
2. Go to **Network** tab
3. Click **Flush DNS**
4. Click **Reset Network**
5. Restart your computer
6. Try again

---

### Region Lock / Access Denied
**Problem:** Service not available in your region.

**Solution:**
1. Use **Google Test** page to diagnose
2. Note which services are blocked
3. Consider using a VPN in a supported region
4. Try a different Google account

---

### Login Loop / Stuck on "Setting up account"
**Problem:** Can't complete login, stuck on loading screen.

**Solution:**
1. Run **Quick Clean** from Cleaner tab
2. Flush DNS from Network tab
3. Clear all browser cache and cookies
4. Use a personal @gmail.com account (not Workspace)
5. Try in Incognito/Private mode

---

### Gemini Code Assist Conflict
**Problem:** Installed Gemini extension in VS Code, now can't access Antigravity.

**Solution:**
1. Go to [Google Account Security](https://myaccount.google.com/permissions)
2. Revoke access for "Gemini Code Assist"
3. Clear browser cookies
4. Try Antigravity again

---

### Application Won't Start
**Problem:** Nothing happens when running the executable.

**Solutions:**

**Windows:**
- Right-click → Run as Administrator
- Check Windows Defender didn't block it
- Install Visual C++ Redistributable

**macOS:**
- Right-click → Open (bypass Gatekeeper)
- Allow in System Preferences → Security

**Linux:**
- Make executable: `chmod +x ./antigravity`
- Install Tkinter: `sudo apt install python3-tk`

---

### Session Restore Failed
**Problem:** Can't restore browser session backup.

**Solution:**
1. Close the target browser completely
2. Check backup file exists in sessions folder
3. Try restoring to a different profile
4. Verify backup file isn't corrupted

---

## Still Need Help?

- 📢 [Ask in Telegram](https://t.me/panbehnet)
- 🐛 [Open GitHub Issue](https://github.com/tawroot/antigravity-cleaner/issues)
- 📧 Include your OS version and error message
