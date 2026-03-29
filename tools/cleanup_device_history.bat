@echo off
:: device history cleanup for dma testing
:: run as administrator before connecting device
:: creates registry backup first

net session >nul 2>&1
if %errorlevel% neq 0 (
    echo [!] run as administrator
    pause
    exit /b 1
)

echo ============================================
echo   device history cleanup
echo   creating registry backup first...
echo ============================================

:: backup registry before changes
set BACKUP_DIR=%USERPROFILE%\Desktop\reg_backup_%date:~-4,4%%date:~-7,2%%date:~-10,2%
mkdir "%BACKUP_DIR%" >nul 2>&1
reg export "HKLM\SYSTEM\CurrentControlSet\Enum" "%BACKUP_DIR%\enum_backup.reg" /y >nul 2>&1
echo [+] registry backup saved to %BACKUP_DIR%

echo.
echo [*] phase 1: stopping services...
net stop DeviceInstall >nul 2>&1

echo [*] phase 2: cleaning setupapi logs...
del /f /q "%SystemRoot%\inf\setupapi.dev.log" >nul 2>&1
del /f /q "%SystemRoot%\inf\setupapi.dev.log.old" >nul 2>&1
del /f /q "%SystemRoot%\inf\setupapi.app.log" >nul 2>&1
del /f /q "%SystemRoot%\inf\setupapi.app.log.old" >nul 2>&1
del /f /q "%SystemRoot%\inf\setupapi.offline.log" >nul 2>&1
del /f /q "%SystemRoot%\inf\setupapi.setup.log" >nul 2>&1

echo [*] phase 3: cleaning device install logs...
del /f /q "%ProgramData%\Microsoft\Windows\DeviceMetadataStore\*.*" /s >nul 2>&1

echo [*] phase 4: cleaning pnp event logs...
wevtutil cl Microsoft-Windows-Kernel-PnP/Configuration >nul 2>&1
wevtutil cl Microsoft-Windows-Kernel-PnP/Device Management >nul 2>&1
wevtutil cl Microsoft-Windows-DeviceSetupManager/Admin >nul 2>&1
wevtutil cl Microsoft-Windows-DeviceSetupManager/Operational >nul 2>&1
wevtutil cl Microsoft-Windows-UserPnp/DeviceInstall >nul 2>&1
wevtutil cl Microsoft-Windows-UserPnp/ActionCenter >nul 2>&1
wevtutil cl Microsoft-Windows-USB-USBHUB3-Analytic >nul 2>&1
wevtutil cl Microsoft-Windows-USB-UCX-Analytic >nul 2>&1

echo [*] phase 5: cleaning pci device registry (ghost devices)...
:: remove disconnected pci device entries across all controlsets
for %%C in (ControlSet001 ControlSet002 ControlSet003) do (
    for /f "tokens=*" %%i in ('reg query "HKLM\SYSTEM\%%C\Enum\PCI" 2^>nul ^| findstr /i "VEN_"') do (
        reg delete "%%i" /f >nul 2>&1
    )
)

echo [*] phase 6: cleaning usb device history (ftdi ft601)...
for %%C in (ControlSet001 ControlSet002 ControlSet003) do (
    for /f "tokens=*" %%i in ('reg query "HKLM\SYSTEM\%%C\Enum\USB" 2^>nul ^| findstr /i "VID_0403"') do (
        reg delete "%%i" /f >nul 2>&1
    )
)

echo [*] phase 7: cleaning thunderbolt device cache...
for %%C in (ControlSet001 ControlSet002 ControlSet003) do (
    for /f "tokens=*" %%i in ('reg query "HKLM\SYSTEM\%%C\Enum\THUNDERBOLT" 2^>nul') do (
        reg delete "%%i" /f >nul 2>&1
    )
)

echo [*] phase 8: cleaning device migration cache...
reg delete "HKLM\SYSTEM\CurrentControlSet\Control\DeviceMigration" /f >nul 2>&1

echo [*] phase 9: cleaning driver store temp...
del /f /q "%SystemRoot%\System32\DriverStore\Temp\*.*" >nul 2>&1

echo [*] phase 10: cleaning windows error reporting...
del /f /q /s "%ProgramData%\Microsoft\Windows\WER\ReportQueue\*" >nul 2>&1

echo [*] phase 11: cleaning prefetch...
del /f /q "%SystemRoot%\Prefetch\DEVICESENSUS*" >nul 2>&1
del /f /q "%SystemRoot%\Prefetch\DEVICECENSUS*" >nul 2>&1

echo [*] phase 12: deleting volume shadow copies...
vssadmin delete shadows /all /quiet >nul 2>&1

echo [*] phase 13: cleaning user device association...
reg delete "HKCU\Software\Microsoft\Plug and Play\Device Association Framework\Store" /f >nul 2>&1

echo [*] phase 14: restarting services...
net start DeviceInstall >nul 2>&1

echo.
echo ============================================
echo   [+] cleanup complete
echo   [!] reboot before connecting device
echo   [!] backup: %BACKUP_DIR%
echo ============================================
echo.
pause
