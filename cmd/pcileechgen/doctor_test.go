package main

import (
	"errors"
	"strings"
	"testing"
)

func TestDoctorIOMMUFailureReportsKernelArgFix(t *testing.T) {
	probes := baselineDoctorProbes()
	probes.checkIOMMU = func() error { return errors.New("IOMMU is disabled") }

	checks, err := runDoctor(baselineDoctorOptions(), probes)
	if err != nil {
		t.Fatalf("runDoctor returned error: %v", err)
	}

	check := requireDoctorCheck(t, checks, "IOMMU")
	requireDoctorStatus(t, check, "fail")
	requireDoctorText(t, check.message, "IOMMU is disabled")
	requireDoctorText(t, check.fix, "intel_iommu=on")
	requireDoctorText(t, check.fix, "amd_iommu=on")
}

func TestDoctorSharedIOMMUGroupReportsPeerDevice(t *testing.T) {
	probes := baselineDoctorProbes()
	probes.readIOMMUGroup = func(bdf string) ([]string, error) {
		if bdf != "0000:03:00.0" {
			t.Fatalf("readIOMMUGroup bdf = %q, want 0000:03:00.0", bdf)
		}
		return []string{"0000:03:00.0", "0000:00:1f.0"}, nil
	}

	checks, err := runDoctor(baselineDoctorOptions(), probes)
	if err != nil {
		t.Fatalf("runDoctor returned error: %v", err)
	}

	check := requireDoctorCheck(t, checks, "IOMMU group")
	requireDoctorStatus(t, check, "warn")
	requireDoctorText(t, check.message, "0000:00:1f.0")
	requireDoctorText(t, check.fix, "unbind")
	requireDoctorText(t, check.fix, "vfio-pci")
}

func TestDoctorNonD0PowerStateReportsRecoveryFix(t *testing.T) {
	probes := baselineDoctorProbes()
	probes.readPowerState = func(bdf string) (string, error) {
		if bdf != "0000:03:00.0" {
			t.Fatalf("readPowerState bdf = %q, want 0000:03:00.0", bdf)
		}
		return "D3hot", nil
	}

	checks, err := runDoctor(baselineDoctorOptions(), probes)
	if err != nil {
		t.Fatalf("runDoctor returned error: %v", err)
	}

	check := requireDoctorCheck(t, checks, "Power state")
	requireDoctorStatus(t, check, "fail")
	requireDoctorText(t, check.message, "D3hot")
	requireDoctorText(t, check.fix, "/sys/bus/pci/devices/0000:03:00.0/power/control")
	requireDoctorText(t, check.fix, "on")
}

func TestDoctorBoardCompatibilityFailsWhenDonorBARExceedsBRAM(t *testing.T) {
	probes := baselineDoctorProbes()
	probes.readBARs = func(bdf string) ([]doctorBAR, error) {
		if bdf != "0000:03:00.0" {
			t.Fatalf("readBARs bdf = %q, want 0000:03:00.0", bdf)
		}
		return []doctorBAR{{index: 0, size: 16 * 1024}}, nil
	}
	probes.boardBRAMSize = func(board string) (int, error) {
		if board != "PCIeSquirrel" {
			t.Fatalf("boardBRAMSize board = %q, want PCIeSquirrel", board)
		}
		return 4 * 1024, nil
	}

	checks, err := runDoctor(baselineDoctorOptions(), probes)
	if err != nil {
		t.Fatalf("runDoctor returned error: %v", err)
	}

	check := requireDoctorCheck(t, checks, "Board compatibility")
	requireDoctorStatus(t, check, "fail")
	requireDoctorText(t, check.message, "BAR0")
	requireDoctorText(t, check.message, "16384")
	requireDoctorText(t, check.message, "4096")
	requireDoctorText(t, check.message, "PCIeSquirrel")
	requireDoctorText(t, check.fix, "larger BRAM")
}

func TestDoctorMissingVivadoIsWarningNotFatal(t *testing.T) {
	probes := baselineDoctorProbes()
	probes.findVivado = func(path string) (string, error) {
		return "", errors.New("vivado not found")
	}

	checks, err := runDoctor(baselineDoctorOptions(), probes)
	if err != nil {
		t.Fatalf("runDoctor returned error for missing Vivado: %v", err)
	}

	check := requireDoctorCheck(t, checks, "Vivado")
	requireDoctorStatus(t, check, "warn")
	requireDoctorText(t, check.message, "vivado not found")
	requireDoctorText(t, check.fix, "--skip-vivado")
	requireDoctorText(t, check.fix, "--vivado-path")
}

func baselineDoctorOptions() doctorOptions {
	return doctorOptions{
		bdf:        "0000:03:00.0",
		board:      "PCIeSquirrel",
		vivadoPath: "",
		skipVivado: false,
	}
}

func baselineDoctorProbes() doctorProbes {
	return doctorProbes{
		checkIOMMU:    func() error { return nil },
		checkVFIOMods: func() error { return nil },
		readDeviceInfo: func(string) (doctorDeviceInfo, error) {
			return doctorDeviceInfo{vendorID: 0x8086, deviceID: 0x100e, className: "Ethernet controller"}, nil
		},
		readIOMMUGroup: func(string) ([]string, error) { return []string{"0000:03:00.0"}, nil },
		readPowerState: func(string) (string, error) { return "D0", nil },
		readDriver:     func(string) (string, error) { return "vfio-pci", nil },
		readBARs:       func(string) ([]doctorBAR, error) { return []doctorBAR{{index: 0, size: 4096}}, nil },
		boardBRAMSize:  func(string) (int, error) { return 4096, nil },
		findVivado:     func(string) (string, error) { return "/opt/Xilinx/Vivado/bin/vivado", nil },
	}
}

func requireDoctorCheck(t *testing.T, checks []doctorCheck, name string) doctorCheck {
	t.Helper()
	for _, check := range checks {
		if check.name == name {
			return check
		}
	}
	t.Fatalf("doctor checks = %+v, want check named %q", checks, name)
	return doctorCheck{}
}

func requireDoctorStatus(t *testing.T, check doctorCheck, want string) {
	t.Helper()
	if check.status != want {
		t.Fatalf("%s status = %q, want %q (check: %+v)", check.name, check.status, want, check)
	}
}

func requireDoctorText(t *testing.T, got, wantSubstring string) {
	t.Helper()
	if !strings.Contains(got, wantSubstring) {
		t.Fatalf("text = %q, want substring %q", got, wantSubstring)
	}
}
