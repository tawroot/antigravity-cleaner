@echo off
REM ========================================
REM Antigravity Cleaner - EXE Builder
REM ========================================

title Building Antigravity Cleaner EXE

echo.
echo ========================================
echo   Antigravity Cleaner - EXE Builder
echo ========================================
echo.

REM Run the PowerShell build script
powershell -ExecutionPolicy Bypass -File "%~dp0build_exe.ps1"

exit /b 0
