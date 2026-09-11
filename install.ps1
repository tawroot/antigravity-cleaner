$ErrorActionPreference = "SilentlyContinue"
$GithubBaseUrl = "https://github.com/tawroot/antigravity-cleaner/releases/download/v5.1.1"
$ExeName = "antigravity-cleaner-windows-amd64.exe"

$HomePath = if ($IsWindows -or $env:OS -like "*Windows*") { $env:USERPROFILE } else { $env:HOME }
$InstallDir = Join-Path $HomePath ".antigravity\bin"
$TargetFile = Join-Path $InstallDir "antigravity-cleaner.exe"

Write-Host "====================================================" -ForegroundColor Cyan
Write-Host "  ⚡ Antigravity Cleaner Toolkit Installer (v5.1.1)" -ForegroundColor Yellow
Write-Host "  ✨ Universal Support: Auto-detects ALL Antigravity Versions (1.x, 2.x, IDE & CLI)" -ForegroundColor Cyan
Write-Host "====================================================" -ForegroundColor Cyan

# 1. Create Directory
if (!(Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
}

# 2. Download Native Go Binary
Write-Host "Downloading Antigravity Native Engine..." -ForegroundColor Yellow
try {
    $DownloadUrl = "$GithubBaseUrl/$ExeName"
    Invoke-WebRequest -Uri $DownloadUrl -OutFile $TargetFile -UseBasicParsing
    Copy-Item -Path $TargetFile -Destination (Join-Path $InstallDir "ag-cleaner.exe") -Force -ErrorAction SilentlyContinue
    Copy-Item -Path $TargetFile -Destination (Join-Path $InstallDir "agc.exe") -Force -ErrorAction SilentlyContinue
    Write-Host "Download Complete." -ForegroundColor Green
}
catch {
    Write-Host "Falling back to legacy PowerShell engine..." -ForegroundColor Gray
    $LegacyScript = Join-Path $InstallDir "Antigravity.ps1"
    Invoke-WebRequest -Uri "https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main/legacy/Antigravity.ps1" -OutFile $LegacyScript -UseBasicParsing
    Start-Process powershell.exe -ArgumentList "-NoProfile -ExecutionPolicy Bypass -File `"$LegacyScript`""
    exit 0
}

# 3. Create Desktop Shortcut
try {
    $WshShell = New-Object -ComObject WScript.Shell
    $DesktopPath = [Environment]::GetFolderPath("Desktop")
    $ShortcutFile = "$DesktopPath\Antigravity Cleaner.lnk"
    $Shortcut = $WshShell.CreateShortcut($ShortcutFile)
    $Shortcut.TargetPath = $TargetFile
    $Shortcut.Save()
    Write-Host "Desktop shortcut created successfully!" -ForegroundColor Green
}
catch {
    Write-Host "Could not create desktop shortcut automatically." -ForegroundColor Gray
}

# 4. Add to user PATH
$UserPath = [Environment]::GetEnvironmentVariable("PATH", "User")
if ($UserPath -notlike "*$InstallDir*") {
    [Environment]::SetEnvironmentVariable("PATH", "$UserPath;$InstallDir", "User")
    Write-Host "Added $InstallDir to user PATH." -ForegroundColor Cyan
}

Write-Host ""
Write-Host "Installation Complete! Launching Antigravity Cleaner..." -ForegroundColor Green
Start-Process $TargetFile
