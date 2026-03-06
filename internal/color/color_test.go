package color

import (
	"strings"
	"testing"
)

func TestOK(t *testing.T) {
	Enable()
	s := OK("test message")
	if !strings.Contains(s, "test message") {
		t.Error("OK should contain the message")
	}
	if !strings.Contains(s, "[OK]") {
		t.Error("OK should contain [OK] prefix")
	}
}

func TestFail(t *testing.T) {
	Enable()
	s := Fail("error occurred")
	if !strings.Contains(s, "error occurred") {
		t.Error("Fail should contain the message")
	}
	if !strings.Contains(s, "[FAIL]") {
		t.Error("Fail should contain [FAIL] prefix")
	}
}

func TestWarn(t *testing.T) {
	Enable()
	s := Warn("warning msg")
	if !strings.Contains(s, "warning msg") {
		t.Error("Warn should contain the message")
	}
	if !strings.Contains(s, "[WARN]") {
		t.Error("Warn should contain [WARN] prefix")
	}
}

func TestInfo(t *testing.T) {
	Enable()
	s := Info("info msg")
	if !strings.Contains(s, "info msg") {
		t.Error("Info should contain the message")
	}
	if !strings.Contains(s, "[INFO]") {
		t.Error("Info should contain [INFO] prefix")
	}
}

func TestBold(t *testing.T) {
	s := Bold("bold text")
	if !strings.Contains(s, "bold text") {
		t.Error("Bold should contain the text")
	}
}

func TestDim(t *testing.T) {
	s := Dim("dimmed")
	if !strings.Contains(s, "dimmed") {
		t.Error("Dim should contain the text")
	}
}

func TestHeader(t *testing.T) {
	s := Header("Section")
	if !strings.Contains(s, "Section") {
		t.Error("Header should contain the text")
	}
	if !strings.Contains(s, "---") {
		t.Error("Header should contain dashes")
	}
}

func TestOkf(t *testing.T) {
	s := Okf("count: %d", 42)
	if !strings.Contains(s, "count: 42") {
		t.Error("Okf should format arguments")
	}
}

func TestFailf(t *testing.T) {
	s := Failf("err: %s", "boom")
	if !strings.Contains(s, "err: boom") {
		t.Error("Failf should format arguments")
	}
}

func TestWarnf(t *testing.T) {
	s := Warnf("warn: %d%%", 50)
	if !strings.Contains(s, "warn: 50%") {
		t.Error("Warnf should format arguments")
	}
}

func TestDisable_StripsColor(t *testing.T) {
	Enable()
	colored := OK("test")
	Disable()
	plain := OK("test")
	Enable() // restore

	if !strings.Contains(plain, "[OK] test") {
		t.Error("disabled color should still contain text & prefix")
	}
	if plain == colored && strings.Contains(colored, "\033") {
		t.Error("disabled should differ from enabled when terminal colors are on")
	}
}

func TestDisableEnable_RoundTrip(t *testing.T) {
	Disable()
	s := Bold("hello")
	if s != "hello" {
		t.Errorf("Disable: Bold = %q, want %q", s, "hello")
	}

	Enable()
	s = Bold("hello")
	if !strings.Contains(s, "hello") {
		t.Error("Enable: Bold should still contain text")
	}
}
