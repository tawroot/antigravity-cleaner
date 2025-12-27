$ErrorActionPreference = "SilentlyContinue"
$GithubBaseUrl = "https://raw.githubusercontent.com/tawroot/antigravity-cleaner/main"

# Detect OS
$IsWin = if ($null -ne $IsWindows) { $IsWindows } else { $env:OS -like "*Windows*" }
$HomePath = if ($IsWin) { $env:USERPROFILE } else { $env:HOME }
$PSExe = if ($IsWin) { "powershell.exe" } else { "pwsh" }

$InstallDir = Join-Path $HomePath ".antigravity"
$TargetFile = Join-Path $InstallDir "Antigravity.ps1"

# 1. Create Directory
if (!(Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
    Write-Host "Created installation directory: $InstallDir" -ForegroundColor Cyan
}

# 2. Download Main Script
Write-Host "Downloading Antigravity Shell..." -ForegroundColor Yellow
try {
    Invoke-WebRequest -Uri "$GithubBaseUrl/Antigravity.ps1" -OutFile $TargetFile -UseBasicParsing
    Write-Host "Download Complete." -ForegroundColor Green
}
catch {
    Write-Host "Failed to download script. Check internet connection." -ForegroundColor Red
    exit 1
}

# 3. Create Desktop Shortcut (Windows Only)
if ($IsWin) {
    try {
        $WshShell = New-Object -ComObject WScript.Shell
        $DesktopPath = [Environment]::GetFolderPath("Desktop")
        $ShortcutFile = "$DesktopPath\Antigravity Shell.lnk"
        $Shortcut = $WshShell.CreateShortcut($ShortcutFile)
        $Shortcut.TargetPath = "powershell.exe"
        $Shortcut.Arguments = "-NoProfile -ExecutionPolicy Bypass -File `"$TargetFile`""
        $Shortcut.IconLocation = "shell32.dll,238" 
        $Shortcut.Save()
        Write-Host "Shortcut created on Desktop." -ForegroundColor Green
    }
    catch {
        Write-Host "Could not create shortcut automatically." -ForegroundColor Gray
    }
}
else {
    Write-Host ""
    Write-Host "Installation Note (macOS/Linux):" -ForegroundColor Cyan
    Write-Host "To run Antigravity in the future, use:" -ForegroundColor White
    Write-Host "  pwsh $TargetFile" -ForegroundColor Yellow
    Write-Host ""
}

# 4. Launch
Write-Host "Launching Antigravity..." -ForegroundColor Cyan
Start-Process $PSExe -ArgumentList "-NoProfile -ExecutionPolicy Bypass -File `"$TargetFile`""
