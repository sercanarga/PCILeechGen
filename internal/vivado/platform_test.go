package vivado

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestFindDiscoversPlatformLauncherInPathWithLiteralMetacharacters(t *testing.T) {
	installDir := filepath.Join(t.TempDir(), "Xilinx %VIVADO_DISCOVERY_ROOT% & (native)^! Suite", "2025.1")
	binaryPath, _ := writeFakeVivadoLauncher(t, installDir)

	found, err := Find(installDir)
	if err != nil {
		t.Fatalf("Find(custom path) returned error: %v", err)
	}
	if found.Path != installDir {
		t.Fatalf("Find(custom path).Path = %q, want %q", found.Path, installDir)
	}
	if found.BinaryPath() != binaryPath {
		t.Fatalf("BinaryPath() = %q, want %q", found.BinaryPath(), binaryPath)
	}

	t.Setenv("PATH", filepath.Dir(binaryPath)+string(os.PathListSeparator)+os.Getenv("PATH"))
	found, err = Find("")
	if err != nil {
		t.Fatalf("Find(PATH) returned error: %v", err)
	}
	if found.Path != installDir {
		t.Fatalf("Find(PATH).Path = %q, want %q", found.Path, installDir)
	}
	if found.BinaryPath() != binaryPath {
		t.Fatalf("Find(PATH).BinaryPath() = %q, want %q", found.BinaryPath(), binaryPath)
	}
}

func TestRunTCLPreservesLiteralMetacharactersArgumentsAndEnvironment(t *testing.T) {
	t.Setenv("VIVADO_TEST_ROOT", "EXPANDED_INSTALL")
	t.Setenv("VIVADO_TEST_WORK", "EXPANDED_WORK")
	t.Setenv("VIVADO_TEST_SCRIPT", "EXPANDED_SCRIPT")
	installDir := filepath.Join(t.TempDir(), "Xilinx %VIVADO_TEST_ROOT% & (native)^! Suite", "2025.1")
	_, invocationPath := writeFakeVivadoLauncher(t, installDir)
	workDir := filepath.Join(t.TempDir(), "firmware %VIVADO_TEST_WORK% & (offline)^! output")
	if err := os.MkdirAll(workDir, 0o755); err != nil {
		t.Fatalf("create working directory: %v", err)
	}
	tclPath := filepath.Join(workDir, "generate %VIVADO_TEST_SCRIPT% & project^!.tcl")
	if err := os.WriteFile(tclPath, nil, 0o644); err != nil {
		t.Fatalf("write Tcl fixture: %v", err)
	}

	v := &Vivado{Path: installDir, Version: "2025.1"}
	if err := v.RunTCL(tclPath, workDir, 5*time.Second); err != nil {
		t.Fatalf("RunTCL returned error: %v", err)
	}

	data, err := os.ReadFile(invocationPath)
	if err != nil {
		t.Fatalf("read launcher invocation: %v", err)
	}
	got := strings.Split(strings.TrimSpace(string(data)), "\n")
	for i := range got {
		got[i] = strings.TrimSuffix(got[i], "\r")
		if len(got[i]) >= 2 && got[i][0] == '"' && got[i][len(got[i])-1] == '"' {
			got[i] = got[i][1 : len(got[i])-1]
		}
	}
	want := []string{"-mode", "batch", "-notrace", "-source", tclPath, workDir, installDir}
	if len(got) != len(want) {
		t.Fatalf("launcher invocation = %#v, want %#v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("launcher invocation[%d] = %q, want %q; full invocation %#v", i, got[i], want[i], got)
		}
	}
}

func writeFakeVivadoLauncher(t *testing.T, installDir string) (string, string) {
	t.Helper()
	binDir := filepath.Join(installDir, "bin")
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		t.Fatalf("create fake Vivado bin directory: %v", err)
	}
	invocationPath := filepath.Join(binDir, "invocation.txt")
	if runtime.GOOS == "windows" {
		binaryPath := filepath.Join(binDir, "vivado.bat")
		script := "@echo off\r\necho \"%~1\">\"%~dp0invocation.txt\"\r\necho \"%~2\">>\"%~dp0invocation.txt\"\r\necho \"%~3\">>\"%~dp0invocation.txt\"\r\necho \"%~4\">>\"%~dp0invocation.txt\"\r\necho \"%~5\">>\"%~dp0invocation.txt\"\r\necho \"%CD%\">>\"%~dp0invocation.txt\"\r\necho \"%XILINX_VIVADO%\">>\"%~dp0invocation.txt\"\r\nexit /b 0\r\n"
		if err := os.WriteFile(binaryPath, []byte(script), 0o644); err != nil {
			t.Fatalf("write fake Vivado batch launcher: %v", err)
		}
		return binaryPath, invocationPath
	}
	binaryPath := filepath.Join(binDir, "vivado")
	script := "#!/bin/sh\nprintf '%s\\n' \"$1\" \"$2\" \"$3\" \"$4\" \"$5\" \"$PWD\" \"$XILINX_VIVADO\" > \"$(dirname \"$0\")/invocation.txt\"\n"
	if err := os.WriteFile(binaryPath, []byte(script), 0o755); err != nil {
		t.Fatalf("write fake Vivado launcher: %v", err)
	}
	return binaryPath, invocationPath
}
