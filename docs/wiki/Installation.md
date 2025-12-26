# Installation Guide

## Windows

### Method 1: Portable EXE (Recommended)
1. Download [AntigravityCleaner.exe](https://github.com/tawroot/antigravity-cleaner/releases/latest)
2. Run the file - no installation needed!

### Method 2: PowerShell
```powershell
iwr -useb https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.ps1 | iex
```

---

## macOS

### Terminal Installation
```bash
curl -sSL https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.sh | bash
```

### Manual Installation
1. Download from [Releases](https://github.com/tawroot/antigravity-cleaner/releases)
2. Extract the archive
3. Run `python3 src/gui_apple.py`

**Note:** macOS may show a security warning. Right-click the app and select "Open" to bypass.

---

## Linux

### Ubuntu/Debian
```bash
curl -sSL https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/install.sh | bash
```

### Requirements
- Python 3.10+
- Tkinter (`sudo apt install python3-tk`)

### Manual Installation
```bash
git clone https://github.com/tawroot/antigravity-cleaner.git
cd antigravity-cleaner
pip install -r requirements.txt  # If exists
python3 src/gui_apple.py
```

---

## System Requirements

| | Minimum | Recommended |
|-|---------|-------------|
| **RAM** | 4 GB | 8 GB |
| **Storage** | 100 MB | 500 MB |
| **Python** | 3.10 | 3.11+ |

---

## Verification

After installation, you should see the Dashboard with your system health score.

![Dashboard Screenshot](https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/website/og-image.png)

---

## Need Help?

- 📢 [Join Telegram](https://t.me/panbehnet)
- 🐛 [Report Issues](https://github.com/tawroot/antigravity-cleaner/issues)
