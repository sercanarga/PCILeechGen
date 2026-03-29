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
reg export "HKLM\SYSTEM\CurrentControlSet\Control" "%BACKUP_DIR%\control_backup.reg" /y >nul 2>&1
echo [+] registry backup saved to %BACKUP_DIR%

echo.
echo [*] phase 1: stopping services...
net stop DeviceInstall >nul 2>&1
net stop PlugPlay >nul 2>&1
net stop WMPNetworkSvc >nul 2>&1

echo [*] phase 2: cleaning setupapi logs...
del /f /q "%SystemRoot%\inf\setupapi.dev.log" >nul 2>&1
del /f /q "%SystemRoot%\inf\setupapi.dev.log.old" >nul 2>&1
del /f /q "%SystemRoot%\inf\setupapi.app.log" >nul 2>&1
del /f /q "%SystemRoot%\inf\setupapi.app.log.old" >nul 2>&1
del /f /q "%SystemRoot%\inf\setupapi.offline.log" >nul 2>&1
del /f /q "%SystemRoot%\inf\setupapi.setup.log" >nul 2>&1
del /f /q "%SystemRoot%\inf\setupapi.upgrade.log" >nul 2>&1

echo [*] phase 3: cleaning device install logs...
rd /s /q "%SystemRoot%\INF\WmiApRpl" >nul 2>&1
del /f /q "%SystemRoot%\Logs\MoSetup\*.log" >nul 2>&1
del /f /q "%ProgramData%\Microsoft\Windows\DeviceMetadataStore\*.*" /s >nul 2>&1
del /f /q "%SystemRoot%\System32\LogFiles\setupcln\*.log" >nul 2>&1

echo [*] phase 4: cleaning pnp and system event logs...
wevtutil cl Microsoft-Windows-Kernel-PnP/Configuration >nul 2>&1
wevtutil cl Microsoft-Windows-Kernel-PnP/Device Management >nul 2>&1
wevtutil cl Microsoft-Windows-DeviceSetupManager/Admin >nul 2>&1
wevtutil cl Microsoft-Windows-DeviceSetupManager/Operational >nul 2>&1
wevtutil cl Microsoft-Windows-PnPDev/InfDev >nul 2>&1
wevtutil cl Microsoft-Windows-UserPnp/DeviceInstall >nul 2>&1
wevtutil cl Microsoft-Windows-UserPnp/ActionCenter >nul 2>&1
wevtutil cl Microsoft-Windows-Kernel-Boot/Operational >nul 2>&1
wevtutil cl Microsoft-Windows-USB-USBHUB3-Analytic >nul 2>&1
wevtutil cl Microsoft-Windows-USB-UCX-Analytic >nul 2>&1
wevtutil cl System >nul 2>&1
wevtutil cl Application >nul 2>&1

echo [*] phase 5: cleaning pci device registry entries...
:: remove disconnected pci device entries across all controlsets
for %%C in (ControlSet001 ControlSet002 ControlSet003) do (
    for /f "tokens=*" %%i in ('reg query "HKLM\SYSTEM\%%C\Enum\PCI" 2^>nul ^| findstr /i "VEN_"') do (
        reg delete "%%i" /f >nul 2>&1
    )
)

echo [*] phase 6: cleaning usb device history (ftdi ft601)...
:: clean ft601 usb traces across all controlsets
for %%C in (ControlSet001 ControlSet002 ControlSet003) do (
    for /f "tokens=*" %%i in ('reg query "HKLM\SYSTEM\%%C\Enum\USB" 2^>nul ^| findstr /i "VID_0403"') do (
        reg delete "%%i" /f >nul 2>&1
    )
    :: clean USBSTOR entries
    reg delete "HKLM\SYSTEM\%%C\Enum\USBSTOR" /f >nul 2>&1
)

echo [*] phase 7: cleaning thunderbolt device cache...
for %%C in (ControlSet001 ControlSet002 ControlSet003) do (
    for /f "tokens=*" %%i in ('reg query "HKLM\SYSTEM\%%C\Enum\THUNDERBOLT" 2^>nul') do (
        reg delete "%%i" /f >nul 2>&1
    )
    reg delete "HKLM\SYSTEM\%%C\Services\ucx01000\TrustDB" /f >nul 2>&1
    reg delete "HKLM\SYSTEM\%%C\Services\TbtP2pBridgeClass" /f >nul 2>&1
)

echo [*] phase 8: cleaning device migration and setup cache...
reg delete "HKLM\SYSTEM\CurrentControlSet\Control\DeviceMigration" /f >nul 2>&1
reg delete "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Setup\PnpResources" /f >nul 2>&1
reg delete "HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion\EMDMgmt" /f >nul 2>&1

echo [*] phase 9: cleaning driver store staging...
del /f /q "%SystemRoot%\System32\DriverStore\Temp\*.*" >nul 2>&1

echo [*] phase 10: cleaning device class coinstallers cache...
:: forces windows to re-evaluate device classes on next plug
reg delete "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\DeviceAccess" /f >nul 2>&1

echo [*] phase 11: cleaning windows error reporting device traces...
del /f /q /s "%ProgramData%\Microsoft\Windows\WER\ReportQueue\*" >nul 2>&1
del /f /q /s "%LOCALAPPDATA%\Microsoft\Windows\WER\ReportQueue\*" >nul 2>&1

echo [*] phase 12: cleaning prefetch cache...
del /f /q "%SystemRoot%\Prefetch\DEVICESENSUS*" >nul 2>&1
del /f /q "%SystemRoot%\Prefetch\DEVICECENSUS*" >nul 2>&1
del /f /q "%SystemRoot%\Prefetch\WMIPRVSE*" >nul 2>&1

echo [*] phase 13: cleaning wmi repository...
net stop winmgmt /y >nul 2>&1
rd /s /q "%SystemRoot%\System32\wbem\Repository" >nul 2>&1
net start winmgmt >nul 2>&1
:: winmgmt will rebuild repository on next boot

echo [*] phase 14: deleting volume shadow copies (restore points)...
vssadmin delete shadows /all /quiet >nul 2>&1

echo [*] phase 15: cleaning usb connection history (user hive)...
reg delete "HKCU\Software\Microsoft\Plug and Play\Device Association Framework\Store" /f >nul 2>&1
reg delete "HKCU\Software\Microsoft\Windows\CurrentVersion\DeviceAccess" /f >nul 2>&1

echo [*] phase 16: cleaning setupdi device interface cache...
:: remove cached device interfaces that hold old hardware ids
for /f "tokens=*" %%i in ('reg query "HKLM\SYSTEM\CurrentControlSet\Control\DeviceClasses" 2^>nul') do (
    reg delete "%%i" /f >nul 2>&1
)

echo [*] phase 17: restarting services...
net start PlugPlay >nul 2>&1
net start DeviceInstall >nul 2>&1

echo.
echo ============================================
echo   [+] cleanup complete
echo   [!] REBOOT REQUIRED before connecting device
echo   [!] registry backup saved to:
echo       %BACKUP_DIR%
echo ============================================
echo.
pause
