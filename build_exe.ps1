# ========================================
# Antigravity Cleaner - Build EXE Script
# ========================================
# This script builds a standalone EXE file
# ========================================

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Antigravity Cleaner - EXE Builder" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# Check if Python is installed
try {
    $pythonVersion = python --version 2>&1
    Write-Host "OK Python found: $pythonVersion" -ForegroundColor Green
}
catch {
    Write-Host "ERROR Python is not installed!" -ForegroundColor Red
    Write-Host "Please install Python 3.8+ from https://www.python.org/" -ForegroundColor Yellow
    pause
    exit 1
}

# Check if PyInstaller is installed
Write-Host ""
Write-Host "Checking PyInstaller..." -ForegroundColor Yellow

try {
    $pyinstallerVersion = pyinstaller --version 2>&1
    Write-Host "OK PyInstaller found: $pyinstallerVersion" -ForegroundColor Green
}
catch {
    Write-Host "WARNING PyInstaller not found. Installing..." -ForegroundColor Yellow
    pip install pyinstaller
    
    if ($LASTEXITCODE -ne 0) {
        Write-Host "ERROR Failed to install PyInstaller!" -ForegroundColor Red
        pause
        exit 1
    }
    
    Write-Host "OK PyInstaller installed successfully!" -ForegroundColor Green
}

# Build the EXE
Write-Host ""
Write-Host "Building EXE file..." -ForegroundColor Yellow
Write-Host "This may take a few minutes..." -ForegroundColor Gray
Write-Host ""

# Create build directory
if (!(Test-Path "dist")) {
    New-Item -ItemType Directory -Path "dist" | Out-Null
}

# Build command - check if icon exists
if (Test-Path "icon.ico") {
    Write-Host "Building with custom icon..." -ForegroundColor Cyan
    pyinstaller --onefile --windowed --name "AntigravityCleaner" --icon=icon.ico --add-data "src;src" --hidden-import=customtkinter src/gui_apple.py
}
else {
    Write-Host "Building without icon..." -ForegroundColor Cyan
    pyinstaller --onefile --windowed --name "AntigravityCleaner" --add-data "src;src" --hidden-import=customtkinter src/gui_apple.py
}




# Check if build was successful
if ($LASTEXITCODE -eq 0) {
    Write-Host ""
    Write-Host "========================================" -ForegroundColor Green
    Write-Host "  BUILD SUCCESSFUL!" -ForegroundColor Green
    Write-Host "========================================" -ForegroundColor Green
    Write-Host ""
    Write-Host "EXE file created at:" -ForegroundColor Cyan
    Write-Host "   dist\AntigravityCleaner.exe" -ForegroundColor White
    Write-Host ""
    Write-Host "You can now distribute this file to users!" -ForegroundColor Yellow
    Write-Host "No Python installation required!" -ForegroundColor Yellow
    Write-Host ""
    
    # Ask if user wants to run the EXE
    $run = Read-Host "Do you want to run the EXE now? (Y/N)"
    if ($run -eq "Y" -or $run -eq "y") {
        Start-Process "dist\AntigravityCleaner.exe"
    }
}
else {
    Write-Host ""
    Write-Host "BUILD FAILED!" -ForegroundColor Red
    Write-Host "Check the error messages above." -ForegroundColor Yellow
}

Write-Host ""
pause
