#Requires -Version 5.1
#Requires -Modules @{ ModuleName = 'Pester'; RequiredVersion = '5.7.1' }

$modulePath = Join-Path $PSScriptRoot '..' 'DeviceHistoryCleanup.psm1'
Import-Module $modulePath -Force

Describe 'DeviceHistoryCleanup safety contract' {
    InModuleScope DeviceHistoryCleanup {
        It 'accepts only numbered ControlSet keys' {
            Test-NumberedControlSetName -Name 'ControlSet001' | Should -BeTrue
            Test-NumberedControlSetName -Name 'ControlSet999' | Should -BeTrue
            Test-NumberedControlSetName -Name 'CurrentControlSet' | Should -BeFalse
            Test-NumberedControlSetName -Name 'ControlSet01' | Should -BeFalse
            Test-NumberedControlSetName -Name 'Select' | Should -BeFalse
        }

        It 'rejects aliases and non-PCI/USB registry paths' {
            { Assert-DeviceHistoryRegistryPath -Path 'Registry::HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Enum\PCI\VEN_1234\1' } |
                Should -Throw
            { Assert-DeviceHistoryRegistryPath -Path 'Registry::HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\ACPI\VEN_1234\1' } |
                Should -Throw
            { Assert-DeviceHistoryRegistryPath -Path 'Registry::HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\USB\VID_1234\1' } |
                Should -Not -Throw
            { Assert-DeviceHistoryRegistryPath -Path 'Microsoft.PowerShell.Core\Registry::HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\PCI\VEN_1234\1' } |
                Should -Not -Throw
        }

        It 'performs no writes or purge confirmation for WhatIf' {
            $device = [pscustomobject]@{
                Type          = 'PCI'
                HardwareId    = 'VEN_1234&DEV_5678'
                InstanceId    = '1'
                RegistryPaths = @('Registry::HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\PCI\VEN_1234&DEV_5678\1')
            }

            Mock New-DeviceHistoryBackup {}
            Mock Get-DeviceInstallServiceState { [pscustomobject]@{ Status = 'Running' } }
            Mock Remove-DeviceRegistryKeys {}
            Mock Restore-DeviceHistoryBackup {}
            Mock Set-DeviceInstallServiceState {}
            Mock Confirm-IrreversiblePurge {}
            Mock Clear-DeviceSystemHistory {}

            $result = Invoke-DeviceHistoryCleanup -Device $device -PurgeSystemHistory -WhatIf

            $result.Performed | Should -BeFalse
            Assert-MockCalled New-DeviceHistoryBackup -Times 0 -Exactly
            Assert-MockCalled Remove-DeviceRegistryKeys -Times 0 -Exactly
            Assert-MockCalled Confirm-IrreversiblePurge -Times 0 -Exactly
            Assert-MockCalled Clear-DeviceSystemHistory -Times 0 -Exactly
        }

        It 'exports and hashes every registry key before returning a backup manifest' {
            $device = [pscustomobject]@{
                RegistryPaths = @(
                    'Registry::HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\PCI\VEN_1\A',
                    'Registry::HKEY_LOCAL_MACHINE\SYSTEM\ControlSet002\Enum\PCI\VEN_1\A'
                )
            }

            $backupRoot = Join-Path $TestDrive 'backup'
            Mock Export-RegistryKey {
                param($NativePath, $Destination)
                [System.IO.File]::WriteAllText(
                    $Destination,
                    "Windows Registry Editor Version 5.00`r`n`r`n[$NativePath]`r`n"
                )
            }

            $backup = New-DeviceHistoryBackup -Device $device -BackupRoot $backupRoot

            $backup.Entries.Count | Should -Be 2
            $backup.Entries[0].Sha256 | Should -Match '^[0-9A-F]{64}$'
            $backup.Entries[0].File | Should -Be (Join-Path $backup.Root 'registry-000.reg')
            Test-Path -LiteralPath $backup.ManifestPath -PathType Leaf | Should -BeTrue
            Assert-MockCalled Export-RegistryKey -Times 2 -Exactly
        }

        It 'rejects a backup file outside its absolute recorded root' {
            $root = Join-Path $TestDrive 'contained-root'
            $outside = Join-Path $TestDrive 'outside.reg'

            { Assert-BackupFilePath -Root $root -Path $outside } | Should -Throw
            { Assert-BackupFilePath -Root 'relative-root' -Path $outside } | Should -Throw
        }

        It 'does not import a registry backup after its contents change' {
            $file = Join-Path $TestDrive 'tampered.reg'
            [System.IO.File]::WriteAllText($file, 'original')
            $expectedSha256 = (Get-FileHash -LiteralPath $file -Algorithm SHA256).Hash
            [System.IO.File]::WriteAllText($file, 'tampered')
            Mock Import-RegistryKey {}

            { Import-VerifiedRegistryKey -Source $file -ExpectedSha256 $expectedSha256 } | Should -Throw
            Assert-MockCalled Import-RegistryKey -Times 0 -Exactly
        }

        It 'imports an intact backup only from its recorded root' {
            $root = Join-Path $TestDrive 'restore-root'
            $null = New-Item -ItemType Directory -Path $root
            $file = Join-Path $root 'registry-000.reg'
            [System.IO.File]::WriteAllText($file, 'verified backup')
            $fileItem = Get-Item -LiteralPath $file
            $backup = [pscustomobject]@{
                Root    = $root
                Entries = @(
                    [pscustomobject]@{
                        RegistryPath = 'Registry::HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\PCI\VEN_1234\1'
                        File         = $file
                        Length       = [long]$fileItem.Length
                        Sha256       = (Get-FileHash -LiteralPath $file -Algorithm SHA256).Hash
                    }
                )
            }
            Mock Import-RegistryKey {}

            { Restore-DeviceHistoryBackup -Backup $backup } | Should -Not -Throw
            Assert-MockCalled Import-RegistryKey -Times 1 -Exactly
        }

        It 'rolls registry state back when service restoration is the first failure' {
            $device = [pscustomobject]@{
                Type          = 'PCI'
                HardwareId    = 'VEN_1234&DEV_5678'
                InstanceId    = '1'
                RegistryPaths = @('Registry::HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\PCI\VEN_1234&DEV_5678\1')
            }
            Mock New-DeviceHistoryBackup { [pscustomobject]@{ Root = 'unused'; Entries = @() } }
            Mock Get-DeviceInstallServiceState { [pscustomobject]@{ Status = 'Stopped' } }
            Mock Remove-DeviceRegistryKeys {}
            Mock Set-DeviceInstallServiceState { throw 'service restore failed' }
            Mock Restore-DeviceHistoryBackup {}

            { Invoke-DeviceHistoryCleanup -Device $device -Confirm:$false } | Should -Throw
            Assert-MockCalled Set-DeviceInstallServiceState -Times 1 -Exactly
            Assert-MockCalled Restore-DeviceHistoryBackup -Times 1 -Exactly
        }

        It 'requires the exact phrase for irreversible purge' {
            Mock Read-Host { 'delete device history' }
            { Confirm-IrreversiblePurge } | Should -Throw

            Mock Read-Host { 'PURGE DEVICE HISTORY' }
            { Confirm-IrreversiblePurge } | Should -Not -Throw
        }
    }
}
