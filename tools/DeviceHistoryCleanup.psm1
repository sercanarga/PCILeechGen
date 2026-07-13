Set-StrictMode -Version Latest

$script:IrreversiblePurgePhrase = 'PURGE DEVICE HISTORY'
$script:RegistrySystemRoot = 'Registry::HKEY_LOCAL_MACHINE\SYSTEM'
$script:BackupSchemaVersion = 2
$script:Sha256Pattern = '^[0-9A-F]{64}$'

function Test-NumberedControlSetName {
    [CmdletBinding()]
    param([Parameter(Mandatory)][string]$Name)

    return $Name -cmatch '^ControlSet\d{3}$'
}

function Assert-DeviceHistoryRegistryPath {
    [CmdletBinding()]
    param([Parameter(Mandatory)][string]$Path)

    $normalizedPath = $Path -replace '^Microsoft\.PowerShell\.Core\\', ''
    $pattern = '^Registry::HKEY_LOCAL_MACHINE\\SYSTEM\\ControlSet\d{3}\\Enum\\(PCI|USB)\\[^\\]+\\[^\\]+$'
    if ($normalizedPath -cnotmatch $pattern) {
        throw "Refusing unsafe device-history registry path: $Path"
    }

    return $normalizedPath
}

function ConvertTo-NativeRegistryPath {
    [CmdletBinding()]
    param([Parameter(Mandatory)][string]$Path)

    $safePath = Assert-DeviceHistoryRegistryPath -Path $Path
    return ($safePath -replace '^Registry::HKEY_LOCAL_MACHINE\\', 'HKLM\')
}

function ConvertTo-AbsoluteFileSystemPath {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)][string]$Path,
        [string]$Description = 'Path'
    )

    if (-not [System.IO.Path]::IsPathRooted($Path)) {
        throw "$Description must be an absolute filesystem path: $Path"
    }

    try {
        $absolutePath = [System.IO.Path]::GetFullPath($Path)
        $pathRoot = [System.IO.Path]::GetPathRoot($absolutePath)
        if (-not [string]::Equals($absolutePath, $pathRoot, [System.StringComparison]::OrdinalIgnoreCase)) {
            $trimCharacters = [char[]]@(
                [System.IO.Path]::DirectorySeparatorChar,
                [System.IO.Path]::AltDirectorySeparatorChar
            )
            $absolutePath = $absolutePath.TrimEnd($trimCharacters)
        }
        return $absolutePath
    }
    catch {
        throw "$Description is not a valid filesystem path: $Path"
    }
}

function Assert-BackupFilePath {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)][string]$Root,
        [Parameter(Mandatory)][string]$Path
    )

    $absoluteRoot = ConvertTo-AbsoluteFileSystemPath -Path $Root -Description 'Backup root'
    $absolutePath = ConvertTo-AbsoluteFileSystemPath -Path $Path -Description 'Backup file'
    $parent = [System.IO.Path]::GetDirectoryName($absolutePath)
    if (-not [string]::Equals($parent, $absoluteRoot, [System.StringComparison]::OrdinalIgnoreCase)) {
        throw "Backup file must be a direct child of its recorded root: $absolutePath"
    }

    return $absolutePath
}

function Assert-NotReparsePoint {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)][string]$Path,
        [string]$Description = 'Path'
    )

    $item = Get-Item -LiteralPath $Path -Force -ErrorAction Stop
    if (($item.Attributes -band [System.IO.FileAttributes]::ReparsePoint) -ne 0) {
        throw "$Description cannot be a filesystem reparse point: $Path"
    }

    return $item
}

function Get-StreamSha256 {
    [CmdletBinding()]
    param([Parameter(Mandatory)][System.IO.Stream]$Stream)

    $algorithm = $null
    try {
        $Stream.Position = 0
        $algorithm = [System.Security.Cryptography.SHA256]::Create()
        $hash = $algorithm.ComputeHash($Stream)
        return ([System.BitConverter]::ToString($hash) -replace '-', '')
    }
    finally {
        if ($null -ne $algorithm) {
            $algorithm.Dispose()
        }
        $Stream.Position = 0
    }
}

function Import-VerifiedRegistryKey {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)][string]$Source,
        [Parameter(Mandatory)][string]$ExpectedSha256
    )

    if ($ExpectedSha256 -cnotmatch $script:Sha256Pattern) {
        throw "Backup has an invalid SHA-256 digest: $Source"
    }

    $stream = $null
    try {
        # Deny writes and deletes between verification and reg.exe opening the file.
        $stream = [System.IO.File]::Open(
            $Source,
            [System.IO.FileMode]::Open,
            [System.IO.FileAccess]::Read,
            [System.IO.FileShare]::Read
        )
        $actualSha256 = Get-StreamSha256 -Stream $stream
        if ($actualSha256 -cne $ExpectedSha256) {
            throw "Backup SHA-256 mismatch: $Source"
        }

        Import-RegistryKey -Source $Source
    }
    finally {
        if ($null -ne $stream) {
            $stream.Dispose()
        }
    }
}

function Get-NumberedControlSetNames {
    [CmdletBinding()]
    param([string]$SystemRoot = $script:RegistrySystemRoot)

    try {
        $children = Get-ChildItem -LiteralPath $SystemRoot -ErrorAction Stop
    }
    catch {
        throw "Unable to enumerate SYSTEM control sets at ${SystemRoot}: $($_.Exception.Message)"
    }

    return @($children |
        ForEach-Object { $_.PSChildName } |
        Where-Object { Test-NumberedControlSetName -Name $_ } |
        Sort-Object -Unique)
}

function Get-DeviceHistoryEntries {
    [CmdletBinding()]
    param(
        [ValidateSet('PCI', 'USB')][string]$Hive,
        [string]$SystemRoot = $script:RegistrySystemRoot
    )

    $entries = @{}
    foreach ($controlSet in Get-NumberedControlSetNames -SystemRoot $SystemRoot) {
        $hivePath = "$SystemRoot\$controlSet\Enum\$Hive"
        try {
            $null = Get-Item -LiteralPath $hivePath -ErrorAction Stop
        }
        catch [System.Management.Automation.ItemNotFoundException] {
            continue
        }
        catch {
            throw "Unable to inspect ${hivePath}: $($_.Exception.Message)"
        }

        try {
            $hardwareKeys = Get-ChildItem -LiteralPath $hivePath -ErrorAction Stop
        }
        catch {
            throw "Unable to enumerate ${hivePath}: $($_.Exception.Message)"
        }

        foreach ($hardwareKey in $hardwareKeys) {
            try {
                $instances = Get-ChildItem -LiteralPath $hardwareKey.PSPath -ErrorAction Stop
            }
            catch {
                throw "Unable to enumerate $($hardwareKey.PSPath): $($_.Exception.Message)"
            }

            foreach ($instance in $instances) {
                $path = Assert-DeviceHistoryRegistryPath -Path $instance.PSPath
                $key = "$Hive|$($hardwareKey.PSChildName)|$($instance.PSChildName)"
                $friendlyName = $instance.PSChildName
                $driver = '-'

                try {
                    $properties = Get-ItemProperty -LiteralPath $instance.PSPath -ErrorAction Stop
                    if ($properties.FriendlyName) {
                        $friendlyName = [string]$properties.FriendlyName
                    }
                    elseif ($properties.DeviceDesc) {
                        $friendlyName = [string]$properties.DeviceDesc
                    }
                    if ($properties.Driver) {
                        $driver = [string]$properties.Driver
                    }
                }
                catch {
                    # Metadata is optional; enumeration itself remains fail-closed above.
                }

                if (-not $entries.ContainsKey($key)) {
                    $entries[$key] = [pscustomobject]@{
                        Type          = $Hive
                        HardwareId    = $hardwareKey.PSChildName
                        InstanceId    = $instance.PSChildName
                        FriendlyName  = $friendlyName
                        Driver        = $driver
                        ControlSets   = [System.Collections.Generic.List[string]]::new()
                        RegistryPaths = [System.Collections.Generic.List[string]]::new()
                    }
                }

                $entries[$key].ControlSets.Add($controlSet)
                $entries[$key].RegistryPaths.Add($path)
            }
        }
    }

    return @($entries.Values | Sort-Object Type, HardwareId, InstanceId)
}

function Export-RegistryKey {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)][string]$NativePath,
        [Parameter(Mandatory)][string]$Destination
    )

    & reg.exe export $NativePath $Destination /y | Out-Null
    if ($LASTEXITCODE -ne 0) {
        throw "reg.exe export failed for $NativePath with exit code $LASTEXITCODE"
    }
}

function Import-RegistryKey {
    [CmdletBinding()]
    param([Parameter(Mandatory)][string]$Source)

    & reg.exe import $Source | Out-Null
    if ($LASTEXITCODE -ne 0) {
        throw "reg.exe import failed for $Source with exit code $LASTEXITCODE"
    }
}

function New-DeviceHistoryBackup {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)]$Device,
        [string]$BackupRoot = (Join-Path $env:USERPROFILE ("Desktop\PCILeechGen-DeviceHistory-{0}-{1}" -f (Get-Date -Format 'yyyyMMdd-HHmmss'), [Guid]::NewGuid().ToString('N')))
    )

    $paths = @($Device.RegistryPaths | ForEach-Object { [string]$_ })
    if ($paths.Count -eq 0) {
        throw 'Refusing to create a backup without exact device registry paths.'
    }

    $absoluteBackupRoot = ConvertTo-AbsoluteFileSystemPath -Path $BackupRoot -Description 'Backup root'
    if (Test-Path -LiteralPath $absoluteBackupRoot) {
        throw "Refusing to reuse an existing backup root: $absoluteBackupRoot"
    }

    $null = New-Item -ItemType Directory -Path $absoluteBackupRoot -ErrorAction Stop
    $null = Assert-NotReparsePoint -Path $absoluteBackupRoot -Description 'Backup root'
    $entries = [System.Collections.Generic.List[object]]::new()
    $index = 0
    foreach ($path in $paths) {
        $safePath = Assert-DeviceHistoryRegistryPath -Path $path
        $file = Assert-BackupFilePath -Root $absoluteBackupRoot -Path (Join-Path $absoluteBackupRoot ("registry-{0:D3}.reg" -f $index))
        Export-RegistryKey -NativePath (ConvertTo-NativeRegistryPath -Path $safePath) -Destination $file
        $fileItem = Assert-NotReparsePoint -Path $file -Description 'Registry backup file'
        $sha256 = (Get-FileHash -LiteralPath $file -Algorithm SHA256 -ErrorAction Stop).Hash.ToUpperInvariant()
        if ($sha256 -cnotmatch $script:Sha256Pattern) {
            throw "Unable to calculate a valid SHA-256 digest for registry backup: $file"
        }
        $entries.Add([pscustomobject]@{
                RegistryPath = $safePath
                File         = $file
                Length       = [long]$fileItem.Length
                Sha256       = $sha256
            })
        $index++
    }

    $manifestPath = Assert-BackupFilePath -Root $absoluteBackupRoot -Path (Join-Path $absoluteBackupRoot 'manifest.json')
    @{
        SchemaVersion = $script:BackupSchemaVersion
        CreatedAt     = (Get-Date).ToUniversalTime().ToString('o')
        Entries       = $entries.ToArray()
    } | ConvertTo-Json -Depth 4 | Set-Content -LiteralPath $manifestPath -Encoding UTF8 -ErrorAction Stop

    return [pscustomobject]@{
        Root         = $absoluteBackupRoot
        ManifestPath = $manifestPath
        Entries      = $entries.ToArray()
    }
}

function Restore-DeviceHistoryBackup {
    [CmdletBinding()]
    param([Parameter(Mandatory)]$Backup)

    $backupRoot = ConvertTo-AbsoluteFileSystemPath -Path ([string]$Backup.Root) -Description 'Backup root'
    if (-not (Test-Path -LiteralPath $backupRoot -PathType Container)) {
        throw "Backup root is missing: $backupRoot"
    }
    $null = Assert-NotReparsePoint -Path $backupRoot -Description 'Backup root'

    $errors = [System.Collections.Generic.List[System.Exception]]::new()
    foreach ($entry in $Backup.Entries) {
        try {
            $null = Assert-DeviceHistoryRegistryPath -Path $entry.RegistryPath
            $file = Assert-BackupFilePath -Root $backupRoot -Path ([string]$entry.File)
            if (-not (Test-Path -LiteralPath $file -PathType Leaf)) {
                throw "Backup file is missing: $file"
            }
            $fileItem = Assert-NotReparsePoint -Path $file -Description 'Registry backup file'
            if ([long]$fileItem.Length -ne [long]$entry.Length) {
                throw "Backup length mismatch: $file"
            }
            Import-VerifiedRegistryKey -Source $file -ExpectedSha256 ([string]$entry.Sha256)
        }
        catch {
            $errors.Add($_.Exception)
        }
    }

    if ($errors.Count -gt 0) {
        throw [System.AggregateException]::new('One or more registry backups could not be restored.', [System.Exception[]]$errors.ToArray())
    }
}

function Get-DeviceInstallServiceState {
    [CmdletBinding()]
    param()

    try {
        $service = Get-Service -Name DeviceInstall -ErrorAction Stop
    }
    catch {
        throw "Unable to query DeviceInstall service state: $($_.Exception.Message)"
    }

    return [pscustomobject]@{ Status = [string]$service.Status }
}

function Set-DeviceInstallServiceState {
    [CmdletBinding()]
    param([Parameter(Mandatory)]$State)

    $service = Get-Service -Name DeviceInstall -ErrorAction Stop
    if ($State.Status -eq 'Running' -and $service.Status -ne 'Running') {
        Start-Service -Name DeviceInstall -ErrorAction Stop
    }
    elseif ($State.Status -ne 'Running' -and $service.Status -eq 'Running') {
        Stop-Service -Name DeviceInstall -ErrorAction Stop
    }
}

function Remove-DeviceRegistryKeys {
    [CmdletBinding()]
    param([Parameter(Mandatory)]$Device)

    foreach ($path in @($Device.RegistryPaths)) {
        $safePath = Assert-DeviceHistoryRegistryPath -Path $path
        Remove-Item -LiteralPath $safePath -Recurse -Force -ErrorAction Stop
    }
}

function Confirm-IrreversiblePurge {
    [CmdletBinding()]
    param()

    $response = Read-Host "Type '$script:IrreversiblePurgePhrase' to permanently purge Windows device-install logs and metadata"
    if ($response -cne $script:IrreversiblePurgePhrase) {
        throw 'Irreversible purge was not confirmed with the exact phrase.'
    }
}

function Clear-DeviceSystemHistory {
    [CmdletBinding()]
    param()

    $files = @(
        (Join-Path $env:SystemRoot 'inf\setupapi.dev.log'),
        (Join-Path $env:SystemRoot 'inf\setupapi.dev.log.old'),
        (Join-Path $env:SystemRoot 'inf\setupapi.app.log'),
        (Join-Path $env:SystemRoot 'inf\setupapi.app.log.old')
    )
    foreach ($file in $files) {
        if (Test-Path -LiteralPath $file -PathType Leaf) {
            Remove-Item -LiteralPath $file -Force -ErrorAction Stop
        }
    }

    $metadataRoot = Join-Path $env:ProgramData 'Microsoft\Windows\DeviceMetadataStore'
    if (Test-Path -LiteralPath $metadataRoot -PathType Container) {
        Get-ChildItem -LiteralPath $metadataRoot -File -ErrorAction Stop |
            Where-Object { $_.Extension -in '.dmetainf', '.devicemetadata-ms' } |
            Remove-Item -Force -ErrorAction Stop
    }

    foreach ($channel in @(
            'Microsoft-Windows-Kernel-PnP/Configuration',
            'Microsoft-Windows-Kernel-PnP/Device Management',
            'Microsoft-Windows-DeviceSetupManager/Admin',
            'Microsoft-Windows-DeviceSetupManager/Operational',
            'Microsoft-Windows-UserPnp/DeviceInstall',
            'Microsoft-Windows-UserPnp/ActionCenter'
        )) {
        & wevtutil.exe cl $channel | Out-Null
        if ($LASTEXITCODE -ne 0) {
            throw "wevtutil.exe failed to clear $channel with exit code $LASTEXITCODE"
        }
    }
}

function Invoke-DeviceHistoryCleanup {
    [CmdletBinding(SupportsShouldProcess = $true, ConfirmImpact = 'High')]
    param(
        [Parameter(Mandatory)]$Device,
        [switch]$PurgeSystemHistory
    )

    $description = "$($Device.Type) $($Device.HardwareId) / $($Device.InstanceId)"
    if (-not $PSCmdlet.ShouldProcess($description, 'Remove exact device-history registry keys')) {
        return [pscustomobject]@{ Performed = $false; Backup = $null }
    }

    if ($PurgeSystemHistory) {
        Confirm-IrreversiblePurge
    }

    $errors = [System.Collections.Generic.List[System.Exception]]::new()
    $backup = $null
    $serviceState = $null
    try {
        # Backup is intentionally completed before any service or registry mutation.
        $backup = New-DeviceHistoryBackup -Device $Device
        $serviceState = Get-DeviceInstallServiceState
        if ($serviceState.Status -eq 'Running') {
            Stop-Service -Name DeviceInstall -ErrorAction Stop
        }
        Remove-DeviceRegistryKeys -Device $Device
        if ($PurgeSystemHistory) {
            Clear-DeviceSystemHistory
        }
    }
    catch {
        $errors.Add($_.Exception)
    }
    finally {
        $rollbackAttempted = $false
        if ($errors.Count -gt 0 -and $null -ne $backup) {
            $rollbackAttempted = $true
            try {
                Restore-DeviceHistoryBackup -Backup $backup
            }
            catch {
                $errors.Add($_.Exception)
            }
        }

        if ($null -ne $serviceState) {
            try {
                Set-DeviceInstallServiceState -State $serviceState
            }
            catch {
                $errors.Add($_.Exception)
            }
        }

        # If service restoration is the first failure, registry mutation must
        # still be rolled back. A prior cleanup failure already triggered it.
        if ($errors.Count -gt 0 -and $null -ne $backup -and -not $rollbackAttempted) {
            try {
                Restore-DeviceHistoryBackup -Backup $backup
            }
            catch {
                $errors.Add($_.Exception)
            }
        }
    }

    if ($errors.Count -gt 0) {
        throw [System.AggregateException]::new('Device-history cleanup failed; recovery was attempted wherever prior state was available.', [System.Exception[]]$errors.ToArray())
    }

    return [pscustomobject]@{ Performed = $true; Backup = $backup }
}

function Test-DeviceHistoryAdministrator {
    [CmdletBinding()]
    param()

    if ([Environment]::OSVersion.Platform -ne [PlatformID]::Win32NT) {
        return $false
    }
    $identity = [Security.Principal.WindowsIdentity]::GetCurrent()
    $principal = [Security.Principal.WindowsPrincipal]::new($identity)
    return $principal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
}

function Select-DeviceHistoryEntry {
    [CmdletBinding()]
    param([Parameter(Mandatory)][object[]]$Devices)

    for ($index = 0; $index -lt $Devices.Count; $index++) {
        $device = $Devices[$index]
        Write-Host ("[{0}] [{1}] {2} ({3})" -f $index, $device.Type, $device.FriendlyName, $device.HardwareId)
    }
    $selection = Read-Host 'Choose a device number, or Q to quit'
    if ($selection -ceq 'Q') {
        return $null
    }
    $number = 0
    if (-not [int]::TryParse($selection, [ref]$number) -or $number -lt 0 -or $number -ge $Devices.Count) {
        throw 'Invalid device selection.'
    }
    return $Devices[$number]
}

Export-ModuleMember -Function @(
    'Get-DeviceHistoryEntries',
    'Invoke-DeviceHistoryCleanup',
    'Select-DeviceHistoryEntry',
    'Test-DeviceHistoryAdministrator'
)
