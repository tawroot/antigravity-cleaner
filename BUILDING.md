# Building for Different Platforms

## 🪟 Windows

### Requirements:
- Python 3.8+
- PyInstaller

### Build Command:
```powershell
python -m PyInstaller --onefile --windowed --name "AntigravityCleaner" --icon=icon.ico src/gui_cleaner.py
```

### Output:
- `dist/AntigravityCleaner.exe` (Portable EXE file)

---

## 🍎 macOS

### Requirements:
- Python 3.8+
- PyInstaller

### Build Command:
```bash
python3 -m pip install pyinstaller
python3 -m PyInstaller --onefile --windowed --name "AntigravityCleaner" --icon=icon.icns src/gui_cleaner.py
```

### Output:
- `dist/AntigravityCleaner.app` (macOS Application Bundle)

### Notes:
- Use `.icns` format for macOS icons
- May need to sign the app for distribution
- Users may need to allow in Security & Privacy settings

---

## 🐧 Linux

### Requirements:
- Python 3.8+
- PyInstaller
- tkinter (install via package manager)

### Install Dependencies:
```bash
# Ubuntu/Debian
sudo apt-get install python3-tk python3-pip

# Fedora
sudo dnf install python3-tkinter python3-pip

# Arch
sudo pacman -S tk python-pip
```

### Build Command:
```bash
python3 -m pip install pyinstaller
python3 -m PyInstaller --onefile --windowed --name "AntigravityCleaner" src/gui_cleaner.py
```

### Output:
- `dist/AntigravityCleaner` (Executable binary)

### Notes:
- No icon needed (uses system default)
- May need to make executable: `chmod +x dist/AntigravityCleaner`

---

## 📦 Cross-Platform Building

### Using GitHub Actions (Recommended)

You can build for all platforms automatically using GitHub Actions:

1. Create `.github/workflows/build.yml`
2. Push to GitHub
3. Download builds from Releases

### Manual Building

To build for a specific platform, you **must** build on that platform:
- Windows EXE → Build on Windows
- macOS APP → Build on macOS
- Linux Binary → Build on Linux

**You cannot cross-compile** (e.g., build macOS app on Windows).

---

## 🎨 Icon Formats

- **Windows**: `.ico` (256x256 or multi-size)
- **macOS**: `.icns` (512x512 or multi-size)
- **Linux**: `.png` (optional, 256x256)

### Converting Icons:

```bash
# PNG to ICO (Windows)
# Use online converter or ImageMagick:
magick convert icon.png -define icon:auto-resize=256,128,64,48,32,16 icon.ico

# PNG to ICNS (macOS)
# Use iconutil or online converter:
mkdir icon.iconset
sips -z 512 512 icon.png --out icon.iconset/icon_512x512.png
iconutil -c icns icon.iconset
```

---

## 📝 Distribution

### Windows:
- Distribute the `.exe` file
- Optional: Create installer with Inno Setup or NSIS

### macOS:
- Distribute the `.app` bundle
- Optional: Create DMG file
- Optional: Sign and notarize for Gatekeeper

### Linux:
- Distribute the binary
- Optional: Create `.deb`, `.rpm`, or AppImage
- Optional: Publish to Snap Store or Flathub

---

## 🚀 Automated Building with GitHub Actions

See `.github/workflows/build-releases.yml` for automated multi-platform builds.

---

## ⚠️ Common Issues

### Windows:
- Antivirus may flag the EXE (false positive)
- Solution: Sign the EXE or add exception

### macOS:
- "App is damaged" error
- Solution: Sign the app or use `xattr -cr AntigravityCleaner.app`

### Linux:
- Missing tkinter
- Solution: Install `python3-tk` package

---

## 📊 File Sizes

Approximate sizes:
- Windows EXE: 15-25 MB
- macOS APP: 20-30 MB
- Linux Binary: 15-25 MB

These include the Python runtime and all dependencies.
