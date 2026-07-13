[CmdletBinding(SupportsShouldProcess = $true, ConfirmImpact = 'High')]
param(
    [switch]$PurgeSystemHistory
)

Set-StrictMode -Version Latest
$modulePath = Join-Path $PSScriptRoot 'DeviceHistoryCleanup.psm1'
Import-Module $modulePath -Force -ErrorAction Stop

if (-not (Test-DeviceHistoryAdministrator)) {
    throw 'This utility must be run from an elevated Windows PowerShell session.'
}

$devices = @(
    Get-DeviceHistoryEntries -Hive PCI
    Get-DeviceHistoryEntries -Hive USB
)

if ($devices.Count -eq 0) {
    Write-Host 'No PCI or USB device-history entries were found.' -ForegroundColor Green
    exit 0
}

$selected = Select-DeviceHistoryEntry -Devices $devices
if ($null -eq $selected) {
    Write-Host 'No device selected.' -ForegroundColor Yellow
    exit 0
}

Write-Host ''
Write-Host "Selected: [$($selected.Type)] $($selected.FriendlyName) ($($selected.HardwareId))" -ForegroundColor Yellow
Write-Host "Registry copies: $($selected.RegistryPaths.Count) across $($selected.ControlSets -join ', ')" -ForegroundColor Gray
if ($PurgeSystemHistory) {
    Write-Host 'The optional system-history purge is irreversible and will require its exact confirmation phrase.' -ForegroundColor Red
}

$result = Invoke-DeviceHistoryCleanup -Device $selected -PurgeSystemHistory:$PurgeSystemHistory -WhatIf:$WhatIfPreference
if ($result.Performed) {
    Write-Host "Cleanup completed. Per-key registry backups: $($result.Backup.Root)" -ForegroundColor Green
    Write-Host 'Reboot Windows before reconnecting the device.' -ForegroundColor Yellow
}
elseif ($WhatIfPreference) {
    Write-Host 'WhatIf completed: no registry, service, file, or event-log writes were made.' -ForegroundColor Cyan
}
