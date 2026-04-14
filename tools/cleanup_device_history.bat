@echo off
:: pcileechgen device history cleanup - launcher
:: Double-click to run. Requires Administrator privileges.
:: This script launches cleanup_device_history.ps1 with elevation.

net session >nul 2>&1
if %errorlevel% neq 0 (
    echo [!] This script requires Administrator privileges.
    echo.
    echo     Right-click this file -^> "Run as administrator"
    echo.
    pause
    exit /b 1
)

:: find the script next to this batch file
set SCRIPT_DIR=%~dp0
set SCRIPT=%SCRIPT_DIR%cleanup_device_history.ps1

if not exist "%SCRIPT%" (
    echo [!] cleanup_device_history.ps1 not found in %SCRIPT_DIR%
    pause
    exit /b 1
)

powershell -ExecutionPolicy Bypass -File "%SCRIPT%"
pause
