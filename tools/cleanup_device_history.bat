@echo off
setlocal DisableDelayedExpansion

rem Keep the path out of PowerShell source so apostrophes and spaces are safe.
set "PCILEECHGEN_CLEANUP_SCRIPT=%~dp0cleanup_device_history.ps1"

if not exist "%PCILEECHGEN_CLEANUP_SCRIPT%" (
    echo [!] cleanup_device_history.ps1 was not found next to this launcher.
    exit /b 2
)

powershell.exe -NoProfile -ExecutionPolicy Bypass -Command "$ErrorActionPreference='Stop'; try { $scriptPath=[Environment]::GetEnvironmentVariable('PCILEECHGEN_CLEANUP_SCRIPT','Process'); if (-not (Test-Path -LiteralPath $scriptPath -PathType Leaf)) { exit 2 }; $process=Start-Process -FilePath 'powershell.exe' -Verb RunAs -Wait -PassThru -ArgumentList @('-NoProfile','-ExecutionPolicy','Bypass','-File',$scriptPath); if ($null -eq $process) { exit 1 }; exit [int]$process.ExitCode } catch { Write-Error $_; exit 1 }"
set "EXIT_CODE=%ERRORLEVEL%"
endlocal & exit /b %EXIT_CODE%
