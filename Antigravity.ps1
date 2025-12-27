<#
.SYNOPSIS
    Antigravity Cleaner - Professional System Optimization Tool
.DESCRIPTION
    A powerful shell-based tool for system cleaning, session management, and network optimization.
.AUTHOR
    Tawana Network
.VERSION
    4.0.0 (Shell Edition)
#>

$host.UI.RawUI.WindowTitle = "Antigravity Cleaner Shell v4.0.0"
$ErrorActionPreference = "SilentlyContinue"

# --- Configuration ---
$AppTitle = "ANTIGRAVITY CLEANER"
$Version = "4.0.0"
$DataPath = Join-Path $PSScriptRoot "Data"
$SessionDataPath = Join-Path $DataPath "Sessions"
$LogsPath = Join-Path $DataPath "Logs"

# --- UI Helpers ---

function Show-Header {
    Clear-Host
    Write-Host ""
    Write-Host "  ================================================================" -ForegroundColor Cyan
    Write-Host "       $AppTitle v$Version" -ForegroundColor White
    Write-Host "       (c) Tawana Network - Professional Shell Edition" -ForegroundColor DarkGray
    Write-Host "  ================================================================" -ForegroundColor Cyan
    Write-Host ""
}

function Show-Success { param($Message) Write-Host "  [OK] $Message" -ForegroundColor Green }
function Show-Error { param($Message) Write-Host "  [ERROR] $Message" -ForegroundColor Red }
function Show-Info { param($Message) Write-Host "  [INFO] $Message" -ForegroundColor Cyan }
function Show-Warning { param($Message) Write-Host "  [WARN] $Message" -ForegroundColor Yellow }

function Wait-Key {
    Write-Host ""
    Write-Host "  Press any key to continue..." -ForegroundColor Gray -NoNewline
    $null = $host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")
}

# --- modules: Stubbed for now ---
function Clear-Junk {
    param (
        [string]$Path,
        [string]$Description,
        [switch]$DryRun
    )
    if (Test-Path $Path) {
        $files = Get-ChildItem -Path $Path -Recurse -Force -ErrorAction SilentlyContinue
        $count = ($files | Measure-Object).Count
        $size = ($files | Measure-Object -Property Length -Sum).Sum / 1MB
        
        if ($count -gt 0) {
            $msg = "$Description ($Path) - Found $count items ({0:N2} MB)" -f $size
            if ($DryRun) {
                Show-Info "[DRY RUN] $msg"
            } else {
                Try {
                    Remove-Item -Path "$Path\*" -Recurse -Force -ErrorAction Stop
                    Show-Success "Cleaned: $msg"
                } Catch {
                    Show-Warning "Partial Clean: $Description (Locked files skipped)"
                }
            }
        } else {
            Show-Info "Empty: $Description"
        }
    } else {
        Show-Info "Skip: $Description (Not Found)"
    }
}

function Invoke-Cleaner {
    param([switch]$DryRun)
    
    Show-Header
    if ($DryRun) { Write-Host "  *** DRY RUN MODE - NO FILES WILL BE DELETED ***" -ForegroundColor Yellow; Write-Host "" }
    Show-Info "Starting System Clean..."
    
    # System Paths
    Clear-Junk -Path $env:TEMP -Description "User Temp" -DryRun:$DryRun
    Clear-Junk -Path $env:windir\Temp -Description "Windows Temp" -DryRun:$DryRun
    Clear-Junk -Path $env:windir\Prefetch -Description "Prefetch Cache" -DryRun:$DryRun
    
    # Application Paths (Antigravity/IDE Traces)
    $AppData = $env:LOCALAPPDATA
    $Roaming = $env:APPDATA
    
    $TargetPaths = @(
        "$AppData\Antigravity",
        "$Roaming\Antigravity",
        "$AppData\JetBrains",
        "$AppData\VSCode\Cache",
        "$AppData\Google\Chrome\User Data\Default\Cache"
    )
    
    foreach ($path in $TargetPaths) {
        Clear-Junk -Path $path -Description "App Trace" -DryRun:$DryRun
    }
    
    Show-Success "Cleaning Completed."
    Wait-Key
}

# --- Module Wrappers for Menu ---
function Menu-Cleaner {
    Show-Header
    Show-Info "System Cleaning Module"
    Write-Host "  [1] Quick Clean (Standard)"
    Write-Host "  [2] Dry Run (Analyze Only)"
    Write-Host "  [0] Back"
    
    $c = Read-Host "  > Select"
    switch ($c) {
        "1" { Invoke-Cleaner }
        "2" { Invoke-Cleaner -DryRun }
        "0" { return }
    }
}

function Get-Browsers {
    $AppData = $env:LOCALAPPDATA
    @{
        "Google Chrome" = "$AppData\Google\Chrome\User Data"
        "Microsoft Edge" = "$AppData\Microsoft\Edge\User Data"
        "Brave Browser" = "$AppData\BraveSoftware\Brave-Browser\User Data"
        "Opera"         = "$env:APPDATA\Opera Software\Opera Stable"
    }
}

function Get-BrowserAccounts {
    param($UserDataPath)
    $accounts = @()
    
    if (!(Test-Path $UserDataPath)) { return $accounts }

    $profileDirs = Get-ChildItem -Path $UserDataPath -Directory | Where-Object { $_.Name -eq "Default" -or $_.Name -like "Profile *" }
    
    foreach ($dir in $profileDirs) {
        $prefPath = Join-Path $dir.FullName "Preferences"
        if (Test-Path $prefPath) {
            try {
                # Read JSON with error suppression for locked files
                $content = Get-Content -Path $prefPath -Raw -ErrorAction SilentlyContinue
                if ($content) {
                    $json = $content | ConvertFrom-Json
                    
                    $email = $null
                    # Try detection method 1 (Modern Chrome/Edge)
                    if ($json.account_info -and $json.account_info.Count -gt 0) {
                        $email = $json.account_info[0].email
                    }
                    # Method 2 (Legacy)
                    if ([string]::IsNullOrWhiteSpace($email) -and $json.google -and $json.google.services) {
                        $email = $json.google.services.username
                    }

                    if (-not [string]::IsNullOrWhiteSpace($email)) {
                        $accounts += $email
                    }
                }
            } catch {}
        }
    }
    return $accounts | Select-Object -Unique
}

function Invoke-BackupSession {
    Show-Header
    Show-Info "Backup Browser Sessions"
    
    $browsers = Get-Browsers
    $available = @()
    $i = 1
    
    foreach ($key in $browsers.Keys) {
        if (Test-Path $browsers[$key]) {
            $emails = Get-BrowserAccounts -UserDataPath $browsers[$key]
            $emailStr = if ($emails) { " [Accounts: $($emails -join ', ')]" } else { "" }
            
            Write-Host "  [$i] $key" -NoNewline
            if ($emailStr) { Write-Host $emailStr -ForegroundColor DarkGray } else { Write-Host "" }
            
            $available += $key
            $i++
        }
    }
    
    if ($available.Count -eq 0) {
        Show-Warning "No supported browsers found."
        Wait-Key
        return
    }
    
    Write-Host "  [A] All Browsers"
    Write-Host "  [0] Back"
    
    $sel = Read-Host "  > Select Browser to Backup"
    if ($sel -eq "0") { return }
    
    $toBackup = @()
    if ($sel -eq "A" -or $sel -eq "a") {
        $toBackup = $available
    } elseif ($sel -match "^\d+$" -and [int]$sel -le $available.Count -and [int]$sel -gt 0) {
        $toBackup = @($available[[int]$sel - 1])
    } else {
        Show-Error "Invalid Selection"
        Wait-Key
        return
    }
    
    $timestamp = Get-Date -Format "yyyy-MM-dd_HH-mm-ss"
    
    foreach ($bName in $toBackup) {
        $src = $browsers[$bName]
        $dest = Join-Path $SessionDataPath "$bName\$timestamp"
        
        Show-Info "Backing up $bName..."
        try {
            if (!(Test-Path $dest)) { New-Item -ItemType Directory -Path $dest -Force > $null }
            Copy-Item -Path "$src\*" -Destination $dest -Recurse -Force -ErrorAction Stop
            Show-Success "Backup Complete: $dest"
        } catch {
            Show-Error "Failed to backup $bName`n  $_"
        }
    }
    Wait-Key
}

function Invoke-RestoreSession {
    Show-Header
    Show-Info "Restore Browser Sessions"
    
    $browsers = Get-Browsers
    # List backups in SessionDataPath
    $backups = Get-ChildItem -Path $SessionDataPath -Directory
    
    if ($backups.Count -eq 0) {
        Show-Warning "No backups found in $SessionDataPath"
        Wait-Key
        return
    }
    
    $i = 1
    foreach ($b in $backups) {
        Write-Host "  [$i] $($b.Name)"
        $i++
    }
    Write-Host "  [0] Back"
    
    $sel = Read-Host "  > Select Browser Type"
    if ($sel -eq "0") { return }
    
    if ($sel -match "^\d+$" -and [int]$sel -le $backups.Count -and [int]$sel -gt 0) {
        $browserName = $backups[[int]$sel - 1].Name
        $browserPath = $browsers[$browserName]
        
        # List Timestamps
        $timestamps = Get-ChildItem -Path (Join-Path $SessionDataPath $browserName) -Directory
        $j = 1
        Show-Info "Available Backups for $($browserName):"
        foreach ($ts in $timestamps) {
            # Scan backup for emails
            $emails = Get-BrowserAccounts -UserDataPath $ts.FullName
            $emailStr = if ($emails) { " [Accounts: $($emails -join ', ')]" } else { " [No Login Detected]" }

            Write-Host "    [$j] $($ts.Name)" -NoNewline
            Write-Host $emailStr -ForegroundColor DarkGray
            $j++
        }
        
        $tsSel = Read-Host "    > Select Backup to Restore"
        if ($tsSel -match "^\d+$" -and [int]$tsSel -le $timestamps.Count -and [int]$tsSel -gt 0) {
            $restoreSrc = $timestamps[[int]$tsSel - 1].FullName
            
            Show-Warning "Restoring will OVERWRITE current $browserName profile."
            $confirm = Read-Host "    > Are you sure? (Y/N)"
            if ($confirm -eq "Y") {
                Show-Info "Restoring from $restoreSrc..."
                try {
                    # Kill browser process first
                    $procName = switch -Wildcard ($browserName) { "*Chrome*" {"chrome"} "*Edge*" {"msedge"} "*Brave*" {"brave"} "*Opera*" {"opera"} }
                    Stop-Process -Name $procName -Force -ErrorAction SilentlyContinue
                    
                    Copy-Item -Path "$restoreSrc\*" -Destination $browserPath -Recurse -Force -ErrorAction Stop
                    Show-Success "Restore Successful!"
                } catch {
                    Show-Error "Restore Failed: $_"
                }
            }
        }
    }
    Wait-Key
}

function Invoke-SessionManager {
    Show-Header
    Show-Info "Session Manager Module"
    Write-Host "  [1] Backup Browser Sessions"
    Write-Host "  [2] Restore Browser Sessions"
    Write-Host "  [0] Back"
    
    $choice = Read-Host "  > Select"
    switch ($choice) {
        "1" { Invoke-BackupSession }
        "2" { Invoke-RestoreSession }
        "0" { return }
    }
}

function Test-SystemAnalysis {
    Show-Header
    Show-Info "Antigravity & Google Services Analysis"
    Write-Host ""
    
    # 1. Google Services
    Write-Host "  [Google Services]" -ForegroundColor Cyan
    $googleEndpoints = @{
        "Google Search"      = "https://www.google.com"
        "Google Identity"    = "https://accounts.google.com"
        "Gemini AI API"      = "https://generativelanguage.googleapis.com"
        "Google AI Studio"   = "https://ai.google.dev"
        "Google Cloud Platform" = "https://cloud.google.com"
        "Colab"              = "https://colab.research.google.com"
    }

    foreach ($name in $googleEndpoints.Keys) {
        Write-Host "    testing $name... " -NoNewline
        try {
            $req = Invoke-WebRequest -Uri $googleEndpoints[$name] -UseBasicParsing -TimeoutSec 5 -ErrorAction Stop
            if ($req.StatusCode -eq 200) { Write-Host "OK" -ForegroundColor Green }
            else { Write-Host "WARN ($($req.StatusCode))" -ForegroundColor Yellow }
        } catch {
            Write-Host "FAIL" -ForegroundColor Red
        }
    }
    Write-Host ""

    # 2. Antigravity Dependencies
    Write-Host "  [Antigravity Dependencies]" -ForegroundColor Cyan
    $agEndpoints = @{
        "VSCode Marketplace" = "https://marketplace.visualstudio.com"
        "OpenVSX Registry"   = "https://open-vsx.org"
        "Github API"         = "https://api.github.com"
    }

    foreach ($name in $agEndpoints.Keys) {
        Write-Host "    testing $name... " -NoNewline
        try {
            $req = Invoke-WebRequest -Uri $agEndpoints[$name] -UseBasicParsing -TimeoutSec 5 -ErrorAction Stop
            if ($req.StatusCode -eq 200) { Write-Host "OK" -ForegroundColor Green }
            else { Write-Host "WARN ($($req.StatusCode))" -ForegroundColor Yellow }
        } catch {
            Write-Host "FAIL" -ForegroundColor Red
        }
    }
    Wait-Key
}

function Invoke-NetworkReset {
    Show-Header
    Show-Info "Network Reset Utility"
    Show-Warning "This will reset your network adapters and DNS settings."
    Show-Warning "You may lose internet connection briefly."
    
    $confirm = Read-Host "  > Continue? (Y/N)"
    if ($confirm -ne "Y") { return }
    
    Show-Info "Flushing DNS..."
    cmd /c "ipconfig /flushdns" | Out-Null
    
    Show-Info "Resetting Winsock Catalog..."
    Start-Process "netsh" -ArgumentList "winsock reset catalog" -Wait -NoNewWindow
    
    Show-Info "Resetting TCP/IP..."
    Start-Process "netsh" -ArgumentList "int ip reset" -Wait -NoNewWindow
    
    Show-Success "Network Reset Complete!"
    Show-Info "PLEASE RESTART YOUR COMPUTER for changes to take effect."
    Wait-Key
}

function Invoke-NetworkTools {
    Show-Header
    Show-Info "Network Tools Module"
    Write-Host "  [1] Full System Analysis (Google + Antigravity)"
    Write-Host "  [2] Reset Network (Flush DNS, Winsock)"
    Write-Host "  [0] Back"
    
    $choice = Read-Host "  > Select"
    switch ($choice) {
        "1" { Test-SystemAnalysis }
        "2" { Invoke-NetworkReset }
        "0" { return }
    }
}

# --- Main Loop ---

function Main {
    # Ensure Data Dirs
    if (!(Test-Path $DataPath)) { New-Item -ItemType Directory -Path $DataPath > $null }
    if (!(Test-Path $SessionDataPath)) { New-Item -ItemType Directory -Path $SessionDataPath > $null }

    while ($true) {
        Show-Header
        
        Write-Host "  MAIN MENU" -ForegroundColor Yellow
        Write-Host ""
        Write-Host "  [1] System Cleaner        " -ForegroundColor White -NoNewline; Write-Host "(Clean Junk & Antigravity Traces)" -ForegroundColor DarkGray
        Write-Host "  [2] Session Manager       " -ForegroundColor White -NoNewline; Write-Host "(Backup/Restore Browser Profiles)" -ForegroundColor DarkGray
        Write-Host "  [3] Network Optimizer     " -ForegroundColor White -NoNewline; Write-Host "(Fix Connection & DNS)" -ForegroundColor DarkGray
        Write-Host "  [0] Exit"
        Write-Host ""
        
        $selection = Read-Host "  > Enter Choice"
        
        switch ($selection) {
            "1" { Menu-Cleaner }
            "2" { Invoke-SessionManager }
            "3" { Invoke-NetworkTools }
            "0" { Write-Host "Goodbye!"; exit }
            default { Show-Error "Invalid selection" }
        }
    }
}

# Start
try {
    Main
} catch {
    Show-Error "Critical Engine Failure: $_"
    Wait-Key
}
