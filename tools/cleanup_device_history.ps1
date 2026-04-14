#Requires -RunAsAdministrator

$ErrorActionPreference = 'SilentlyContinue'

function Write-Color($text, $color) {
    Write-Host $text -ForegroundColor $color
}

function Write-Title($title) {
    Write-Host ""
    Write-Host " ============================================" -ForegroundColor Cyan
    Write-Host "   $title" -ForegroundColor Cyan
    Write-Host " ============================================" -ForegroundColor Cyan
}

function Get-EnumDevices($hive) {
    $rawDevices = @()
    $controlSets = @('ControlSet001', 'ControlSet002', 'ControlSet003', 'CurrentControlSet')
    foreach ($cs in $controlSets) {
        $path = "HKLM:\SYSTEM\$cs\Enum\$hive"
        if (Test-Path $path) {
            $items = Get-ChildItem $path -ErrorAction SilentlyContinue
            foreach ($item in $items) {
                $subItems = Get-ChildItem $item.PSPath -ErrorAction SilentlyContinue
                foreach ($sub in $subItems) {
                    $friendlyName = $null
                    $driver = $null
                    try { $friendlyName = (Get-ItemProperty $sub.PSPath -Name 'FriendlyName' -ErrorAction Stop).FriendlyName } catch {}
                    if (-not $friendlyName) {
                        try { $friendlyName = (Get-ItemProperty $sub.PSPath -Name 'DeviceDesc' -ErrorAction Stop).DeviceDesc } catch {}
                    }
                    if (-not $friendlyName) { $friendlyName = $sub.PSChildName }
                    try { $driver = (Get-ItemProperty $sub.PSPath -Name 'Driver' -ErrorAction Stop).Driver } catch {}
                    if (-not $driver) { $driver = '-' }
                    $rawDevices += [PSCustomObject]@{
                        Type         = $hive
                        ControlSet   = $cs
                        FullPath     = $sub.PSPath
                        HardwareId   = $item.PSChildName
                        InstanceId   = $sub.PSChildName
                        FriendlyName = $friendlyName
                        Driver       = $driver
                    }
                }
            }
        }
    }
    return $rawDevices
}

function Deduplicate($rawDevices) {
    $grouped = $rawDevices | Group-Object { "$($_.HardwareId)|$($_.InstanceId)" }
    $devices = @()
    foreach ($g in $grouped) {
        $first = $g.Group[0]
        $controlSets = ($g.Group | ForEach-Object { $_.ControlSet }) -join ', '
        $fullPaths = $g.Group | ForEach-Object { $_.FullPath }
        $drivers = ($g.Group | ForEach-Object { if ($_.Driver -ne '-') { $_.Driver } } | Sort-Object -Unique) -join ', '
        if (-not $drivers) { $drivers = '-' }
        $typeLabel = if ($first.Type -eq 'PCI') { 'PCI' } else { 'USB' }
        $devices += [PSCustomObject]@{
            Type         = $typeLabel
            FullPath     = $first.FullPath
            FullPaths    = $fullPaths
            HardwareId   = $first.HardwareId
            InstanceId   = $first.InstanceId
            FriendlyName = $first.FriendlyName
            Driver       = $drivers
            ControlSets  = $controlSets
        }
    }
    return $devices
}

function Show-Menu {
    param([array]$devices)
    $cursorPos = 0

    while ($true) {
        Clear-Host
        Write-Host " ============================================" -ForegroundColor Cyan
        Write-Host "   pcileech-gen device history cleanup" -ForegroundColor Cyan
        Write-Host "   Up/Down: navigate  Enter: delete  Q: quit" -ForegroundColor Cyan
        Write-Host " ============================================" -ForegroundColor Cyan
        Write-Host ""

        $start = [Math]::Max(0, $cursorPos - 10)
        $end   = [Math]::Min($devices.Count - 1, $start + 20)

        for ($i = $start; $i -le $end; $i++) {
            $dev = $devices[$i]
            $vendor = ($dev.HardwareId -split '&')[0]
            $devPart = ($dev.HardwareId -split '&')[1]
            $driverShort = if ($dev.Driver -ne '-') { ($dev.Driver -split '\\')[0] } else { '-' }
            $typeTag = "[$($dev.Type)]"

            if ($i -eq $cursorPos) {
                Write-Host "> $typeTag $($dev.FriendlyName)" -ForegroundColor Yellow -NoNewline
                Write-Host " ($vendor $devPart)" -ForegroundColor DarkYellow -NoNewline
                Write-Host "  driver: $driverShort" -ForegroundColor Magenta
            } else {
                Write-Host "  $typeTag $($dev.FriendlyName)" -NoNewline
                Write-Host " ($vendor $devPart)" -ForegroundColor DarkGray -NoNewline
                Write-Host "  driver: $driverShort" -ForegroundColor DarkGray
            }
        }

        Write-Host ""
        Write-Host "  total: $($devices.Count) devices" -ForegroundColor Gray

        $key = $host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown')
        switch ($key.VirtualKeyCode) {
            38 { if ($cursorPos -gt 0) { $cursorPos-- } }
            40 { if ($cursorPos -lt ($devices.Count - 1)) { $cursorPos++ } }
            13 { return $devices[$cursorPos] }
            81 { return $null }
        }
    }
}

function Backup-Registry {
    $timestamp = Get-Date -Format "yyyyMMdd_HHmmss"
    $backupDir = "$env:USERPROFILE\Desktop\reg_backup_$timestamp"
    New-Item -ItemType Directory -Path $backupDir -Force | Out-Null
    reg export "HKLM\SYSTEM\CurrentControlSet\Enum" "$backupDir\enum_backup.reg" /y 2>$null
    Write-Color "[+] registry backup: $backupDir" "Green"
    return $backupDir
}

function Remove-Device {
    param($dev)
    $count = 0
    $failed = 0
    foreach ($path in $dev.FullPaths) {
        $regPath = $path -replace 'HKEY_LOCAL_MACHINE', 'HKLM'
        $parentPath = $regPath.Substring(0, $regPath.LastIndexOf('\'))
        $childName = $regPath.Substring($regPath.LastIndexOf('\') + 1)
        try {
            if (Test-Path "HKLM:$($parentPath.Substring(4))") {
                Remove-Item "HKLM:$($parentPath.Substring(4))\$childName" -Recurse -Force -ErrorAction Stop
                $count++
            } else {
                $failed++
            }
        } catch {
            $failed++
        }
    }
    Write-Color "  deleted: $count control set(s)" "Green"
    if ($failed -gt 0) { Write-Color "  failed: $failed" "Red" }
}

Write-Title "pcileech-gen device history cleanup"
Write-Host ""
Write-Host "[*] scanning PCI device history..." -ForegroundColor White
$pciRaw = Get-EnumDevices 'PCI'
$pciCount = ($pciRaw | Group-Object { "$($_.HardwareId)|$($_.InstanceId)" }).Count
Write-Color "[*] found $pciCount unique PCI device(s)" "Green"

Write-Host "[*] scanning USB device history..." -ForegroundColor White
$usbRaw = Get-EnumDevices 'USB'
$usbCount = ($usbRaw | Group-Object { "$($_.HardwareId)|$($_.InstanceId)" }).Count
Write-Color "[*] found $usbCount unique USB device(s)" "Green"

$allRaw = $pciRaw + $usbRaw
if ($allRaw.Count -eq 0) {
    Write-Color "[*] no device history found" "Green"
    Read-Host "Press Enter to exit"
    exit 0
}

$devices = Deduplicate $allRaw
Write-Color "[*] total: $($devices.Count) unique device(s)" "Green"
Write-Host ""
Write-Host "Press any key to open the selection menu..." -ForegroundColor Gray
$null = $host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown')

while ($true) {
    $target = Show-Menu -devices $devices
    if ($null -eq $target) {
        Write-Color "[*] exiting" "Yellow"
        break
    }

    Write-Host ""
    Write-Color "delete this device from all control sets?" "Yellow"
    Write-Host "    type: $($target.Type)" -ForegroundColor White
    Write-Host "    $($target.FriendlyName) ($($target.HardwareId))" -ForegroundColor White
    Write-Host "    driver: $($target.Driver)" -ForegroundColor White
    Write-Host "    control sets: $($target.ControlSets)" -ForegroundColor White
    Write-Host ""
    $confirm = Read-Host "confirm? (yes/no)"

    if ($confirm -eq 'yes') {
        Write-Host ""
        $backupDir = Backup-Registry
        Write-Host "[*] stopping device install service..." -ForegroundColor White
        Stop-Service DeviceInstall -Force -ErrorAction SilentlyContinue
        Write-Host "[*] cleaning setupapi logs..." -ForegroundColor White
        Remove-Item "$env:SystemRoot\inf\setupapi.dev.log" -Force -ErrorAction SilentlyContinue
        Remove-Item "$env:SystemRoot\inf\setupapi.dev.log.old" -Force -ErrorAction SilentlyContinue
        Remove-Item "$env:SystemRoot\inf\setupapi.app.log" -Force -ErrorAction SilentlyContinue
        Remove-Item "$env:SystemRoot\inf\setupapi.app.log.old" -Force -ErrorAction SilentlyContinue
        Write-Host "[*] cleaning device metadata store..." -ForegroundColor White
        Remove-Item "$env:ProgramData\Microsoft\Windows\DeviceMetadataStore\*.dmetainf" -Force -ErrorAction SilentlyContinue
        Remove-Item "$env:ProgramData\Microsoft\Windows\DeviceMetadataStore\*.devicemetadata-ms" -Force -ErrorAction SilentlyContinue
        Write-Host "[*] cleaning PnP event logs..." -ForegroundColor White
        wevtutil cl Microsoft-Windows-Kernel-PnP/Configuration 2>$null
        wevtutil cl Microsoft-Windows-Kernel-PnP/Device\ Management 2>$null
        wevtutil cl Microsoft-Windows-DeviceSetupManager/Admin 2>$null
        wevtutil cl Microsoft-Windows-DeviceSetupManager/Operational 2>$null
        wevtutil cl Microsoft-Windows-UserPnp/DeviceInstall 2>$null
        wevtutil cl Microsoft-Windows-UserPnp/ActionCenter 2>$null
        Write-Host "[*] removing device..." -ForegroundColor White
        Remove-Device -dev $target
        Write-Host "[*] cleaning device migration cache..." -ForegroundColor White
        Remove-Item "HKLM:\SYSTEM\CurrentControlSet\Control\DeviceMigration" -Recurse -Force -ErrorAction SilentlyContinue
        Write-Host "[*] restarting device install service..." -ForegroundColor White
        Start-Service DeviceInstall -ErrorAction SilentlyContinue

        Write-Title "cleanup complete"
        Write-Color "  backup: $backupDir" "Green"
        Write-Host ""
        Write-Color "  [!] REBOOT before connecting device" "Red"
        Write-Host ""
        Read-Host "Press Enter to exit"
        break
    }
}
