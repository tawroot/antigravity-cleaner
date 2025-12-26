# ============================================================================
# ANTIGRAVITY CLEANER - PORTABLE BUILD SCRIPT (v4.0.0 Unified)
# ============================================================================
# Copyright (c) 2024-2025 Tawana Mohammadi / Tawana Network
# All Rights Reserved
# ============================================================================

$ErrorActionPreference = "Stop"

Write-Host ""
Write-Host "============================================================" -ForegroundColor Cyan
Write-Host "   ANTIGRAVITY CLEANER - PORTABLE BUILD v4.0.0" -ForegroundColor White
Write-Host "   (c) Tawana Network - All Rights Reserved" -ForegroundColor Gray
Write-Host "============================================================" -ForegroundColor Cyan
Write-Host ""

# Check Python
if (-not (Get-Command python -ErrorAction SilentlyContinue)) {
    Write-Host "[ERROR] Python not found!" -ForegroundColor Red
    exit 1
}

# Install critical build dependencies
Write-Host "[*] Installing build dependencies..." -ForegroundColor Cyan
python -m pip install --upgrade pip
python -m pip install setuptools wheel
python -m pip install pyinstaller customtkinter psutil pycryptodome rich pillow

# Clean previous builds
Write-Host "[*] Cleaning previous builds..." -ForegroundColor Cyan
if (Test-Path "dist") { Remove-Item -Path "dist" -Recurse -Force }
if (Test-Path "build") { Remove-Item -Path "build" -Recurse -Force }
if (Test-Path "*.spec") { Remove-Item -Path "*.spec" -Force }

# Create output structure
$portableDir = "AntigravityCleaner-Portable"
if (Test-Path $portableDir) { Remove-Item $portableDir -Recurse -Force }
New-Item -ItemType Directory -Path $portableDir -Force | Out-Null

# Build command - The EXACT command that works locally
Write-Host ""
Write-Host "[*] Building Portable EXE..." -ForegroundColor Cyan

$buildArgs = @(
    "-m", "PyInstaller",
    "--onefile",
    "--windowed",
    "--name", "AntigravityCleaner",
    "--clean",
    "--noupx",
    "--version-file=version_info.txt",
    "--manifest=app.manifest",
    "--icon=icon.ico",
    "--collect-all", "customtkinter",
    "--add-data", "src;src",
    "--add-data", "LICENSE;.",
    "--hidden-import=customtkinter",
    "--hidden-import=tkinter",
    "--hidden-import=tkinter.ttk",
    "--hidden-import=tkinter.messagebox",
    "--hidden-import=PIL",
    "--hidden-import=PIL._tkinter_finder",
    "src/gui_apple.py"
)

# Run Build
& python $buildArgs

if ($LASTEXITCODE -ne 0) {
    Write-Host "[ERROR] Build Failed!" -ForegroundColor Red
    exit 1
}

# Move Artifacts
Write-Host "[*] Packaging artifacts..." -ForegroundColor Cyan
Move-Item -Path "dist/AntigravityCleaner.exe" -Destination "$portableDir/" -Force
Copy-Item "README.md", "LICENSE" "$portableDir/" -Force

# Create Portable Marker
$portableContent = @"
Antigravity Cleaner - Portable Edition
======================================

Version: 4.0.0
Platform: Windows
Build Date: $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss UTC')

Copyright (c) 2024-2025 Tawana Mohammadi / Tawana Network
All Rights Reserved

This is the portable version. No installation required.
Just run AntigravityCleaner.exe

Support: https://github.com/tawroot/antigravity-cleaner/issues
"@
$portableContent | Out-File -FilePath "$portableDir/PORTABLE.txt" -Encoding UTF8

New-Item -ItemType Directory -Path "$portableDir/data" -Force | Out-Null

# Create ZIP
$zipName = "AntigravityCleaner-Windows-Portable.zip"
if (Test-Path $zipName) { Remove-Item $zipName -Force }
Compress-Archive -Path "$portableDir/*" -DestinationPath $zipName -Force

Write-Host ""
Write-Host "============================================================" -ForegroundColor Green
Write-Host "   BUILD SUCCESSFUL!" -ForegroundColor White
Write-Host "   Output: $zipName" -ForegroundColor Yellow
Write-Host "============================================================" -ForegroundColor Green
