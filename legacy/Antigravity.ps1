<#
.SYNOPSIS
    Antigravity Cleaner - Professional System Optimization Tool
.DESCRIPTION
    A powerful shell-based tool for system cleaning, session management, and network optimization.
.AUTHOR
    Tawana Network
.VERSION
    4.1.0 (Shell Edition)
#>

$host.UI.RawUI.WindowTitle = "Antigravity Cleaner Shell v4.1.0"
$ErrorActionPreference = "SilentlyContinue"

# --- Configuration & Platform Detection ---
$AppTitle = "ANTIGRAVITY CLEANER"
$Version = "4.1.0"

# Detect OS safely (avoiding conflict with built-in readonly variables in PS 6+)
$Global:OS_Win = if ($null -ne $IsWindows) { $IsWindows } else { $true }
$Global:OS_Mac = if ($null -ne $IsMacOS) { $IsMacOS } else { $false }
$Global:OS_Lin = if ($null -ne $IsLinux) { $IsLinux } else { $false }

# Fallback for PowerShell 5.1
if ($null -eq $IsWindows) {
    $Global:OS_Win = $env:OS -like "*Windows*"
    $Global:OS_Mac = [Runtime.InteropServices.RuntimeInformation]::IsOSPlatform([Runtime.InteropServices.OSPlatform]::OSX) -eq $true
    $Global:OS_Lin = [Runtime.InteropServices.RuntimeInformation]::IsOSPlatform([Runtime.InteropServices.OSPlatform]::Linux) -eq $true
}

$HomePath = if ($OS_Win) { $env:USERPROFILE } else { $env:HOME }

$DataPath = Join-Path $PSScriptRoot "Data"
$SessionDataPath = Join-Path $DataPath "Sessions"

# --- OS Specific Paths ---
$Paths = @{
    Temp     = if ($OS_Win) { $env:TEMP } else { "/tmp" }
    WinTemp  = if ($OS_Win) { "$env:windir\Temp" } else { $null }
    Prefetch = if ($OS_Win) { "$env:windir\Prefetch" } else { $null }
    
    # AppData equivalents
    Local    = if ($OS_Win) { $env:LOCALAPPDATA } elseif ($OS_Mac) { "$HomePath/Library/Application Support" } else { "$HomePath/.config" }
    Roaming  = if ($OS_Win) { $env:APPDATA } elseif ($OS_Mac) { "$HomePath/Library/Application Support" } else { "$HomePath/.config" }
}

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
            }
            else {
                Try {
                    Remove-Item -Path "$Path\*" -Recurse -Force -ErrorAction Stop
                    Show-Success "Cleaned: $msg"
                }
                Catch {
                    Show-Warning "Partial Clean: $Description (Locked files skipped)"
                }
            }
        }
        else {
            Show-Info "Empty: $Description"
        }
    }
    else {
        Show-Info "Skip: $Description (Not Found)"
    }
}

function Invoke-Cleaner {
    param([switch]$DryRun)
    
    Show-Header
    if ($DryRun) { Write-Host "  *** DRY RUN MODE - NO FILES WILL BE DELETED ***" -ForegroundColor Yellow; Write-Host "" }
    Show-Info "Starting System Clean..."
    
    # System Paths
    Clear-Junk -Path $Paths.Temp -Description "Temp Files" -DryRun:$DryRun
    if ($Paths.WinTemp) { Clear-Junk -Path $Paths.WinTemp -Description "OS Temp" -DryRun:$DryRun }
    if ($Paths.Prefetch) { Clear-Junk -Path $Paths.Prefetch -Description "Prefetch" -DryRun:$DryRun }
    
    # Application Paths (Antigravity/IDE Traces)
    $TargetPaths = @()
    
    if ($OS_Win) {
        $TargetPaths += "$($Paths.Local)\Antigravity"
        $TargetPaths += "$($Paths.Roaming)\Antigravity"
        $TargetPaths += "$($Paths.Local)\JetBrains"
        $TargetPaths += "$($Paths.Local)\VSCode\Cache"
    }
    elseif ($OS_Mac) {
        $TargetPaths += "$HomePath/Library/Preferences/Antigravity"
        $TargetPaths += "$HomePath/Library/Caches/JetBrains"
        $TargetPaths += "$HomePath/Library/Application Support/Code/User/workspaceStorage"
    }
    else {
        # Linux
        $TargetPaths += "$HomePath/.config/Antigravity"
        $TargetPaths += "$HomePath/.cache/JetBrains"
    }
    
    foreach ($path in $TargetPaths) {
        Clear-Junk -Path $path -Description "App Trace" -DryRun:$DryRun
    }
    
    Show-Success "Cleaning Completed."
    Wait-Key
}

# --- Module Wrappers for Menu ---
function Invoke-CleanerMenu {
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
    if ($OS_Win) {
        return @{
            "Google Chrome"  = "$($Paths.Local)\Google\Chrome\User Data"
            "Microsoft Edge" = "$($Paths.Local)\Microsoft\Edge\User Data"
            "Brave Browser"  = "$($Paths.Local)\BraveSoftware\Brave-Browser\User Data"
            "Opera"          = "$($Paths.Roaming)\Opera Software\Opera Stable"
        }
    }
    elseif ($OS_Mac) {
        return @{
            "Google Chrome"  = "$HomePath/Library/Application Support/Google/Chrome"
            "Microsoft Edge" = "$HomePath/Library/Application Support/Microsoft Edge"
            "Brave Browser"  = "$HomePath/Library/Application Support/BraveSoftware/Brave-Browser"
            "Opera"          = "$HomePath/Library/Application Support/com.operasoftware.Opera"
        }
    }
    else {
        # Linux
        return @{
            "Google Chrome"  = "$HomePath/.config/google-chrome"
            "Brave Browser"  = "$HomePath/.config/BraveSoftware/Brave-Browser"
            "Microsoft Edge" = "$HomePath/.config/microsoft-edge"
        }
    }
}

# --- Helper for Profile Detection ---
function Get-BrowserProfiles {
    param($BrowserName, $UserDataPath)
    $profiles = @()
    
    if (!(Test-Path $UserDataPath)) { return $profiles }

    # Standard "Default" and "Profile X" folders
    $dirs = Get-ChildItem -Path $UserDataPath -Directory | Where-Object { $_.Name -eq "Default" -or $_.Name -like "Profile *" }
    
    foreach ($dir in $dirs) {
        $email = "No Login"
        $prefPath = Join-Path $dir.FullName "Preferences"
        
        if (Test-Path $prefPath) {
            try {
                $content = Get-Content -Path $prefPath -Raw -ErrorAction SilentlyContinue
                if ($content) {
                    $json = $content | ConvertFrom-Json
                    if ($json.account_info -and $json.account_info.Count -gt 0) {
                        $email = $json.account_info[0].email
                    }
                    elseif ($json.google -and $json.google.services -and $json.google.services.username) {
                        $email = $json.google.services.username
                    }
                }
            }
            catch {}
        }
        
        $profiles += [PSCustomObject]@{
            Browser = $BrowserName
            Name    = $dir.Name
            Path    = $dir.FullName
            Email   = $email
        }
    }
    return $profiles
}

# --- Region Inspector Module ---
function Invoke-RegionInspector {
    Show-Header
    Show-Info "Google Region Inspector"
    Write-Host "  Use this tool to check or change your Google account region."
    Write-Host "  [IMPORTANT] ACTIONS TO INCREASE SUCCESS RATE:" -ForegroundColor Yellow
    Write-Host "    1. VPN must be ON (US/EU IP)."
    Write-Host "    2. Set System Timezone to match VPN location (e.g. EST/PST)."
    Write-Host "    3. Disable WebRTC in browser if possible."
    Write-Host ""
    
    $check = Read-Host "  > Launch Pre-Check? (Opens IP/Leak Test) (Y/N)"
    if ($check -eq "Y") {
        Start-Process "https://browserleaks.com/ip"
        Write-Host "  Checking leaks... ensure 'WebRTC Leak' is NOT showing your real IP." -ForegroundColor Cyan
        Wait-Key
    }

    $browsers = Get-Browsers
    $allProfiles = @()
    $i = 1
    
    foreach ($bKey in $browsers.Keys) {
        $profs = Get-BrowserProfiles -BrowserName $bKey -UserDataPath $browsers[$bKey]
        foreach ($p in $profs) {
            Write-Host "  [$i] $($p.Browser) - $($p.Email)" -NoNewline
            Write-Host " ($($p.Name))" -ForegroundColor DarkGray
            $allProfiles += $p
            $i++
        }
    }
    
    if ($allProfiles.Count -eq 0) {
        Show-Warning "No supported browser profiles found."
        Wait-Key
        return
    }
    
    Write-Host "  [0] Back"
    $sel = Read-Host "  > Select Profile to Inspect"
    if ($sel -eq "0") { return }
    
    if ($sel -match "^\d+$" -and [int]$sel -le $allProfiles.Count -and [int]$sel -gt 0) {
        $target = $allProfiles[[int]$sel - 1]
        
        Show-Info "Opening Region Settings for: $($target.Email)"
        
        # Launch URL with specific profile
        $url = "https://policies.google.com/country-association-form"
        $LaunchCmd = ""
        $LaunchArgs = @()
        
        if ($OS_Win) {
            switch -Wildcard ($target.Browser) {
                "*Chrome*" { $LaunchCmd = "chrome"; $LaunchArgs = "--profile-directory=`"$($target.Name)`" `"$url`"" }
                "*Edge*" { $LaunchCmd = "msedge"; $LaunchArgs = "--profile-directory=`"$($target.Name)`" `"$url`"" }
                "*Brave*" { $LaunchCmd = "brave"; $LaunchArgs = "--profile-directory=`"$($target.Name)`" `"$url`"" }
                "*Opera*" { $LaunchCmd = "launcher"; $LaunchArgs = "`"$url`"" }
            }
        }
        elseif ($OS_Mac) {
            switch -Wildcard ($target.Browser) {
                "*Chrome*" { $LaunchCmd = "open"; $LaunchArgs = @("-a", "Google Chrome", "--args", "--profile-directory=`"$($target.Name)`"", "$url") }
                "*Edge*" { $LaunchCmd = "open"; $LaunchArgs = @("-a", "Microsoft Edge", "--args", "--profile-directory=`"$($target.Name)`"", "$url") }
                "*Brave*" { $LaunchCmd = "open"; $LaunchArgs = @("-a", "Brave Browser", "--args", "--profile-directory=`"$($target.Name)`"", "$url") }
            }
        }
        else {
            # Linux
            switch -Wildcard ($target.Browser) {
                "*Chrome*" { $LaunchCmd = "google-chrome"; $LaunchArgs = "--profile-directory=`"$($target.Name)`" `"$url`"" }
                "*Edge*" { $LaunchCmd = "microsoft-edge"; $LaunchArgs = "--profile-directory=`"$($target.Name)`" `"$url`"" }
                "*Brave*" { $LaunchCmd = "brave-browser"; $LaunchArgs = "--profile-directory=`"$($target.Name)`" `"$url`"" }
            }
        }
        
        try {
            Start-Process $LaunchCmd -ArgumentList $LaunchArgs
            Show-Success "Browser opened. Check the page for 'Country Association'."
        }
        catch {
            Show-Error "Could not launch browser automatically. Please open manually:"
            Write-Host "  $url" -ForegroundColor White
        }
    }
    Wait-Key
}

# --- Enhanced Session Manager ---
# --- Enhanced Session Manager ---
function Invoke-BackupSession {
    Show-Header
    Show-Info "Backup Browser Sessions"
    
    $browsers = Get-Browsers
    $allProfiles = @()
    $i = 1
    
    foreach ($bKey in $browsers.Keys) {
        $profs = Get-BrowserProfiles -BrowserName $bKey -UserDataPath $browsers[$bKey]
        foreach ($p in $profs) {
            Write-Host "  [$i] $($p.Browser) - $($p.Email)" -NoNewline
            Write-Host " ($($p.Name))" -ForegroundColor DarkGray
            $allProfiles += $p
            $i++
        }
    }
    
    if ($allProfiles.Count -eq 0) {
        Show-Warning "No profiles found."
        Wait-Key
        return
    }
    
    Write-Host "  [0] Back"
    $sel = Read-Host "  > Select Profile to Backup"
    if ($sel -eq "0") { return }
    
    if ($sel -match "^\d+$" -and [int]$sel -le $allProfiles.Count -and [int]$sel -gt 0) {
        $p = $allProfiles[[int]$sel - 1]
        
        Write-Host ""
        Show-Info "Select Backup Mode:"
        Write-Host "  [1] Light (Login & Sessions Only) - ~20MB - Quick"
        Write-Host "  [2] Full (Entire Profile) - ~500MB+ - Complete"
        $mode = Read-Host "  > Select Mode"
        
        $timestamp = Get-Date -Format "yyyy-MM-dd_HH-mm-ss"
        $safeEmail = $p.Email -replace "[^a-zA-Z0-9@-]", "_"
        $tag = if ($mode -eq "1") { "Light" } else { "Full" }
        $dest = Join-Path $SessionDataPath "$($p.Browser)\$safeEmail\_$tag`_$timestamp"
        
        Show-Info "Backing up $($p.Browser) ($($p.Email))..."
        try {
            if (!(Test-Path $dest)) { New-Item -ItemType Directory -Path $dest -Force > $null }
            
            if ($mode -eq "1") {
                # Light Mode: Specific files only
                $essentialFiles = @(
                    "Cookies", "Login Data", "Web Data", "Preferences", "Secure Preferences", "Extension Cookies", "Local State"
                )
                $essentialFolders = @("Local Extension Settings", "Sync Data", "Local Storage", "Databases")
                
                foreach ($file in $essentialFiles) {
                    $fPath = Join-Path $p.Path $file
                    if (Test-Path $fPath) { Copy-Item -Path $fPath -Destination $dest -Force -ErrorAction SilentlyContinue }
                }
                foreach ($folder in $essentialFolders) {
                    $dPath = Join-Path $p.Path $folder
                    if (Test-Path $dPath) { 
                        $targetDir = Join-Path $dest $folder
                        New-Item -ItemType Directory -Path $targetDir -Force > $null
                        Copy-Item -Path "$dPath\*" -Destination $targetDir -Recurse -Force -ErrorAction SilentlyContinue 
                    }
                }
                
            }
            else {
                # Full Mode
                Copy-Item -Path "$($p.Path)\*" -Destination $dest -Recurse -Force -ErrorAction Stop
            }
            
            # Save metadata
            $meta = @{
                Browser     = $p.Browser
                ProfileName = $p.Name
                Email       = $p.Email
                Date        = $timestamp
                Mode        = $tag
            } | ConvertTo-Json
            $meta | Out-File (Join-Path $dest "meta.json")
            
            Show-Success "Backup ($tag) Saved to:"
            Write-Host "  $dest" -ForegroundColor White
        }
        catch {
            Show-Error "Backup Failed: $_"
        }
    }
    Wait-Key
}

function Invoke-BackupAntigravityApp {
    Show-Header
    Show-Info "Backup Antigravity Desktop App"
    
    $agPath = if ($OS_Win) { "$env:APPDATA\Antigravity" } elseif ($OS_Mac) { "$HomePath/Library/Application Support/Antigravity" } else { "$HomePath/.config/Antigravity" }
    
    if (!(Test-Path $agPath)) {
        Show-Warning "Antigravity Desktop App data not found."
        Wait-Key
        return
    }
    
    $timestamp = Get-Date -Format "yyyy-MM-dd_HH-mm-ss"
    $dest = Join-Path $SessionDataPath "Antigravity_Desktop\_$timestamp"
    
    Show-Info "Backing up Antigravity App Data..."
    try {
        if (!(Test-Path $dest)) { New-Item -ItemType Directory -Path $dest -Force > $null }
        Copy-Item -Path "$agPath\*" -Destination $dest -Recurse -Force -ErrorAction Stop
        
        $meta = @{
            Browser     = "Antigravity Desktop"
            ProfileName = "Roaming"
            Email       = "Desktop App"
            Date        = $timestamp
            Mode        = "Full"
        } | ConvertTo-Json
        $meta | Out-File (Join-Path $dest "meta.json")
        
        Show-Success "App Backup Complete!"
    }
    catch {
        Show-Error "Backup Failed: $_"
    }
    Wait-Key
}

function Invoke-RestoreSession {
    Show-Header
    Show-Info "Restore Browser Sessions"
    
    # List all backups recursively
    $backups = Get-ChildItem -Path $SessionDataPath -Recurse -File -Filter "meta.json"
    
    if ($backups.Count -eq 0) {
        Show-Warning "No backups found."
        Wait-Key
        return
    }
    
    $i = 1
    $restoreList = @()
    
    foreach ($metaFile in $backups) {
        $json = Get-Content $metaFile.FullName | ConvertFrom-Json
        $folder = $metaFile.Directory.FullName
        $modeStr = if ($json.Mode) { "[$($json.Mode)]" } else { "[Full]" }
        
        Write-Host "  [$i] $modeStr $($json.Browser)" -NoNewline
        Write-Host " - $($json.Email)" -ForegroundColor Cyan -NoNewline
        Write-Host " ($($json.Date))" -ForegroundColor DarkGray
        
        $restoreList += [PSCustomObject]@{
            Header = $json
            Path   = $folder
        }
        $i++
    }
    Write-Host "  [0] Back"
    
    $sel = Read-Host "  > Select Backup to Restore"
    if ($sel -eq "0") { return }
    
    if ($sel -match "^\d+$" -and [int]$sel -le $restoreList.Count -and [int]$sel -gt 0) {
        $target = $restoreList[[int]$sel - 1]
        
        $destPath = ""
        
        if ($target.Header.Browser -eq "Antigravity Desktop") {
            $destPath = if ($OS_Win) { "$env:APPDATA\Antigravity" } elseif ($OS_Mac) { "$HomePath/Library/Application Support/Antigravity" } else { "$HomePath/.config/Antigravity" }
            Show-Info "Restoring Antigravity Desktop App..."
        }
        else {
            # Determine browser path
            $browsers = Get-Browsers
            $browserRoot = $browsers[$target.Header.Browser]
            $destPath = Join-Path $browserRoot $target.Header.ProfileName
        }
        
        Show-Warning "Restoring will OVERWRITE data in: $destPath"
        $confirm = Read-Host "  > Are you sure? (Y/N)"
        if ($confirm -eq "Y") {
            try {
                if ($target.Header.Browser -ne "Antigravity Desktop") {
                    # Stop browser
                    $procName = switch -Wildcard ($target.Header.Browser) { "*Chrome*" { "chrome" } "*Edge*" { "msedge" } "*Brave*" { "brave" } "*Opera*" { "opera" } }
                    if ($OS_Mac) {
                        $procName = switch -Wildcard ($target.Header.Browser) { "*Chrome*" { "Google Chrome" } "*Edge*" { "Microsoft Edge" } "*Brave*" { "Brave Browser" } "*Opera*" { "Opera" } }
                        Start-Process "pkill" -ArgumentList "-x", "$procName" -NoNewWindow -Wait
                    }
                    elseif ($OS_Lin) {
                        Start-Process "pkill" -ArgumentList "$procName" -NoNewWindow -Wait
                    }
                    else {
                        Stop-Process -Name $procName -Force -ErrorAction SilentlyContinue
                    }
                }
                
                # Restore
                Copy-Item -Path "$($target.Path)\*" -Destination $destPath -Recurse -Force -ErrorAction Stop
                Show-Success "Restore Successful!"
            }
            catch {
                Show-Error "Restore Failed: $_"
            }
        }
    }
    Wait-Key
}

function Invoke-SessionManager {
    Show-Header
    Show-Info "Session Manager Module"
    Write-Host "  [1] Backup Browser Profile"
    Write-Host "  [2] Backup Antigravity App"
    Write-Host "  [3] Restore Profile/App"
    Write-Host "  [0] Back"
    
    $choice = Read-Host "  > Select"
    switch ($choice) {
        "1" { Invoke-BackupSession }
        "2" { Invoke-BackupAntigravityApp }
        "3" { Invoke-RestoreSession }
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
        "Google Search"         = "https://www.google.com"
        "Google Identity"       = "https://accounts.google.com"
        "Gemini AI API"         = "https://generativelanguage.googleapis.com"
        "Google AI Studio"      = "https://ai.google.dev"
        "Google Cloud Platform" = "https://cloud.google.com"
        "Colab"                 = "https://colab.research.google.com"
    }

    foreach ($name in $googleEndpoints.Keys) {
        Write-Host "    testing $name... " -NoNewline
        try {
            $req = Invoke-WebRequest -Uri $googleEndpoints[$name] -UseBasicParsing -TimeoutSec 5 -ErrorAction Stop
            if ($req.StatusCode -eq 200) { Write-Host "OK" -ForegroundColor Green }
            else { Write-Host "WARN ($($req.StatusCode))" -ForegroundColor Yellow }
        }
        catch {
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
        }
        catch {
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
    
    if ($OS_Win) {
        Show-Info "Flushing DNS..."
        cmd /c "ipconfig /flushdns" | Out-Null
        
        Show-Info "Resetting Winsock Catalog..."
        Start-Process "netsh" -ArgumentList "winsock reset catalog" -Wait -NoNewWindow
        
        Show-Info "Resetting TCP/IP..."
        Start-Process "netsh" -ArgumentList "int ip reset" -Wait -NoNewWindow
    }
    elseif ($OS_Mac) {
        Show-Info "Flushing DNS (macOS)..."
        Start-Process "sudo" -ArgumentList "killall", "-HUP", "mDNSResponder" -Wait -NoNewWindow
    }
    elseif ($OS_Lin) {
        # Linux
        Show-Info "Flushing DNS (Linux)..."
        if (Get-Command "resolvectl" -ErrorAction SilentlyContinue) {
            Start-Process "sudo" -ArgumentList "resolvectl", "flush-caches" -Wait -NoNewWindow
        }
        else {
            Start-Process "sudo" -ArgumentList "systemd-resolve", "--flush-caches" -Wait -NoNewWindow
        }
    }
    
    Show-Success "Network Reset Complete!"
    Show-Info "If problems persist, please RESTART YOUR COMPUTER."
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
        Write-Host "  [4] System Analysis       " -ForegroundColor White -NoNewline; Write-Host "(Check Google & Antigravity Services)" -ForegroundColor DarkGray
        Write-Host "  [5] Region Inspector      " -ForegroundColor White -NoNewline; Write-Host "(Check/Change Account Region)" -ForegroundColor DarkGray
        Write-Host "  [0] Exit"
        Write-Host ""
        
        $selection = Read-Host "  > Enter Choice"
        
        switch ($selection) {
            "1" { Invoke-CleanerMenu }
            "2" { Invoke-SessionManager }
            "3" { Invoke-NetworkTools }
            "4" { Test-SystemAnalysis }
            "5" { Invoke-RegionInspector }
            "0" { Write-Host "Goodbye!"; exit }
            default { Show-Error "Invalid selection" }
        }
    }
}

# Start
try {
    Main
}
catch {
    Show-Error "Critical Engine Failure: $_"
    Wait-Key
}
