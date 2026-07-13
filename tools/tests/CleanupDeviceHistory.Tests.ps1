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
        }

        It 'performs no writes for WhatIf' {
            $device = [pscustomobject]@{
                Type          = 'PCI'
                HardwareId    = 'VEN_1234&DEV_5678'
                InstanceId    = '1'
                RegistryPaths = @('Registry::HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\PCI\VEN_1234&DEV_5678\1')
            }

            Mock New-DeviceHistoryBackup {}
            Mock Get-DeviceInstallServiceState { [pscustomobject]@{ Status = 'Running' } }
            Mock Stop-Service {}
            Mock Remove-Item {}
            Mock Restore-DeviceHistoryBackup {}
            Mock Set-DeviceInstallServiceState {}

            $result = Invoke-DeviceHistoryCleanup -Device $device -WhatIf

            $result.Performed | Should -BeFalse
            Assert-MockCalled New-DeviceHistoryBackup -Times 0 -Exactly
            Assert-MockCalled Stop-Service -Times 0 -Exactly
            Assert-MockCalled Remove-Item -Times 0 -Exactly
        }

        It 'exports every registry key before returning a backup manifest' {
            $device = [pscustomobject]@{
                RegistryPaths = @(
                    'Registry::HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\PCI\VEN_1\A',
                    'Registry::HKEY_LOCAL_MACHINE\SYSTEM\ControlSet002\Enum\PCI\VEN_1\A'
                )
            }

            Mock New-Item { [pscustomobject]@{ FullName = 'TestDrive:\backup' } }
            Mock Export-RegistryKey {}
            Mock Set-Content {}

            $backup = New-DeviceHistoryBackup -Device $device -BackupRoot 'TestDrive:\backup'

            $backup.Entries.Count | Should -Be 2
            Assert-MockCalled Export-RegistryKey -Times 2 -Exactly
            Assert-MockCalled Set-Content -Times 1 -Exactly
        }

        It 'requires the exact phrase for irreversible purge' {
            Mock Read-Host { 'delete device history' }
            { Confirm-IrreversiblePurge } | Should -Throw

            Mock Read-Host { 'PURGE DEVICE HISTORY' }
            { Confirm-IrreversiblePurge } | Should -Not -Throw
        }
    }
}
