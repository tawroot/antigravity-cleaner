@echo off
REM ========================================
REM Antigravity Cleaner - GUI Launcher
REM ========================================

title Antigravity Cleaner GUI

echo.
echo ========================================
echo   Antigravity Cleaner v2.1 - GUI
echo ========================================
echo.

REM Check if Python is installed
where python >nul 2>nul
if %errorlevel% neq 0 (
    echo Error: Python is not installed!
    echo Please install Python 3.8 or higher.
    pause
    exit /b 1
)

REM Check Python version
python --version | findstr /R "3\.[89] 3\.1[0-9]" >nul
if %errorlevel% neq 0 (
    echo Warning: Python 3.8+ is recommended
)

echo Starting GUI...
echo.

REM Run the GUI
cd /d "%~dp0"
python src\gui_cleaner.py

if %errorlevel% neq 0 (
    echo.
    echo Error: Failed to start GUI
    pause
)

exit /b 0
