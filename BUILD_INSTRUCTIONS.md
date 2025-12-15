# Building EXE File

## Prerequisites

1. Python 3.8 or higher
2. PyInstaller (will be installed automatically)

## How to Build

### Method 1: Double-click (Easiest)
```
Just double-click on BUILD_EXE.bat
```

### Method 2: PowerShell
```powershell
.\build_exe.ps1
```

### Method 3: Manual
```powershell
# Install PyInstaller
pip install pyinstaller

# Build the EXE
pyinstaller AntigravityCleaner.spec
```

## Output

The EXE file will be created in the `dist` folder:
```
dist/AntigravityCleaner.exe
```

## Distribution

You can distribute this single EXE file to users. They don't need Python installed!

## File Size

The EXE will be approximately 15-25 MB (includes Python runtime and all dependencies).

## Notes

- The EXE is standalone and portable
- No installation required
- Works on Windows 7, 8, 10, 11
- Antivirus may flag it (false positive) - this is normal for PyInstaller EXEs

## Customization

To add a custom icon:
1. Place an `icon.ico` file in the project root
2. Rebuild the EXE

## Troubleshooting

If build fails:
1. Make sure Python is in PATH
2. Try: `pip install --upgrade pyinstaller`
3. Check antivirus isn't blocking PyInstaller
