# ============================================================================
# ANTIGRAVITY CLEANER - PORTABLE BUILD SCRIPT
# ============================================================================
# Copyright (c) 2024-2025 Tawana Mohammadi / Tawana Network
# All Rights Reserved
# ============================================================================

Write-Host ""
Write-Host "============================================================" -ForegroundColor Cyan
Write-Host "   ANTIGRAVITY CLEANER - PORTABLE BUILD" -ForegroundColor White
Write-Host "   (c) Tawana Network - All Rights Reserved" -ForegroundColor Gray
Write-Host "============================================================" -ForegroundColor Cyan
Write-Host ""

# Check Python
if (-not (Get-Command python -ErrorAction SilentlyContinue)) {
    Write-Host "[ERROR] Python not found!" -ForegroundColor Red
    exit 1
}

# Check PyInstaller
$pyinstaller = python -c "import PyInstaller" 2>&1
if ($LASTEXITCODE -ne 0) {
    Write-Host "[!] Installing PyInstaller..." -ForegroundColor Yellow
    python -m pip install pyinstaller
}

# Check dependencies
Write-Host "[*] Checking dependencies..." -ForegroundColor Cyan
$deps = @("customtkinter", "psutil", "pycryptodome", "rich")
foreach ($dep in $deps) {
    $check = python -c "import $dep" 2>&1
    if ($LASTEXITCODE -ne 0) {
        Write-Host "  Installing $dep..." -ForegroundColor Yellow
        python -m pip install $dep
    }
}


# Clean previous builds
Write-Host "[*] Cleaning previous builds..." -ForegroundColor Cyan
Remove-Item -Path "dist" -Recurse -Force -ErrorAction SilentlyContinue
Remove-Item -Path "build" -Recurse -Force -ErrorAction SilentlyContinue
Remove-Item -Path "*.spec" -Force -ErrorAction SilentlyContinue

# Create output directory
$outputDir = "dist\AntigravityCleaner-Portable"
New-Item -ItemType Directory -Path $outputDir -Force | Out-Null

# Build command - Optimized for minimal AV detection
Write-Host ""
Write-Host "[*] Building Portable EXE..." -ForegroundColor Cyan
Write-Host "    This may take a few minutes..." -ForegroundColor Gray

$buildArgs = @(
    "--onefile",
    "--windowed",
    "--name", "AntigravityCleaner",
    "--clean",
    "--noupx",  # Disable UPX to reduce false positives
    "--version-file", "version_info.txt",
    "--manifest", "app.manifest"
)

# Add icon if exists
if (Test-Path "icon.ico") {
    $buildArgs += "--icon=icon.ico"
    Write-Host "    Using custom icon: icon.ico" -ForegroundColor Green
}

# Add hidden imports
$hiddenImports = @(
    "customtkinter",
    "tkinter",
    "tkinter.ttk",
    "tkinter.messagebox",
    "PIL",
    "PIL._tkinter_finder"
)
foreach ($hi in $hiddenImports) {
    $buildArgs += "--hidden-import=$hi"
}

# Add data files
$buildArgs += "--add-data", "src;src"
$buildArgs += "--add-data", "LICENSE;."

# Entry point
$buildArgs += "src/gui_apple.py"

# Run PyInstaller
Write-Host ""
Write-Host "[*] Executing PyInstaller..." -ForegroundColor DarkGray

# Build the command as an array of arguments for better reliability
$allArgs = @("-m", "PyInstaller") + $buildArgs

# Run using python -m PyInstaller to avoid PATH issues
& python @allArgs

if ($LASTEXITCODE -ne 0) {
    Write-Host ""
    Write-Host "[ERROR] PyInstaller failed!" -ForegroundColor Red
    exit 1
}


# Move to portable folder
Write-Host ""
Write-Host "[*] Creating Portable package..." -ForegroundColor Cyan

Move-Item -Path "dist\AntigravityCleaner.exe" -Destination "$outputDir\" -Force

# Copy additional files
Copy-Item -Path "LICENSE" -Destination "$outputDir\" -Force
Copy-Item -Path "README.md" -Destination "$outputDir\" -Force -ErrorAction SilentlyContinue

# Create portable marker file
@"
Antigravity Cleaner - Portable Edition
======================================

Version: 3.1.0
Copyright: (c) 2024-2025 Tawana Mohammadi / Tawana Network
License: Proprietary - All Rights Reserved

This is the portable version. No installation required.
Just run AntigravityCleaner.exe

Support: https://github.com/tawroot/antigravity-cleaner/issues
"@ | Out-File -FilePath "$outputDir\PORTABLE.txt" -Encoding UTF8

# Create settings folder
New-Item -ItemType Directory -Path "$outputDir\data" -Force | Out-Null

Write-Host ""
Write-Host "============================================================" -ForegroundColor Green
Write-Host "   BUILD SUCCESSFUL!" -ForegroundColor White
Write-Host "============================================================" -ForegroundColor Green
Write-Host ""
Write-Host "Output: $outputDir\AntigravityCleaner.exe" -ForegroundColor Cyan
Write-Host ""
Write-Host "Package contents:" -ForegroundColor Yellow
Get-ChildItem -Path $outputDir | ForEach-Object { Write-Host "  - $($_.Name)" }
Write-Host ""
Write-Host "[TIP] To reduce AV warnings, sign with a certificate:" -ForegroundColor Yellow
Write-Host "      signtool sign /f cert.pfx /p password AntigravityCleaner.exe" -ForegroundColor Gray
Write-Host ""
