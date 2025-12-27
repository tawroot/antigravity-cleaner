$ErrorActionPreference = "SilentlyContinue"
$GithubBaseUrl = "https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main"
$InstallDir = Join-Path $env:USERPROFILE ".antigravity"
$TargetFile = Join-Path $InstallDir "Antigravity.ps1"

# 1. Create Directory
if (!(Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
    Write-Host "Created installation directory: $InstallDir" -ForegroundColor Cyan
}

# 2. Download Main Script
Write-Host "Downloading Antigravity Shell..." -ForegroundColor Yellow
try {
    Invoke-WebRequest -Uri "$GithubBaseUrl/antigravity-cleaner/Antigravity.ps1" -OutFile $TargetFile -UseBasicParsing
    Write-Host "Download Complete." -ForegroundColor Green
}
catch {
    Write-Host "Failed to download script. Check internet connection." -ForegroundColor Red
    exit 1
}

# 3. Create Desktop Shortcut
$WshShell = New-Object -ComObject WScript.Shell
$DesktopPath = [Environment]::GetFolderPath("Desktop")
$ShortcutFile = "$DesktopPath\Antigravity Shell.lnk"
$Shortcut = $WshShell.CreateShortcut($ShortcutFile)
$Shortcut.TargetPath = "powershell.exe"
$Shortcut.Arguments = "-NoProfile -ExecutionPolicy Bypass -File `"$TargetFile`""
$Shortcut.IconLocation = "shell32.dll,238" # Shield Icon
$Shortcut.Save()
Write-Host "Shortcut created on Desktop." -ForegroundColor Green

# 4. Launch
Write-Host "Launching Antigravity..." -ForegroundColor Cyan
Start-Process "powershell.exe" -ArgumentList "-NoProfile -ExecutionPolicy Bypass -File `"$TargetFile`""
