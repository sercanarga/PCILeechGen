// Package color provides ANSI terminal colors.
package color

import (
	"fmt"
	"os"
)

// ANSI color codes
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
	bold   = "\033[1m"
	dimmed = "\033[2m"
)

// enabled tracks whether color output is active.
var enabled = isTerminal()

// isTerminal returns true if stdout appears to be a terminal.
func isTerminal() bool {
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

// Disable turns off color output (useful for piped/redirected output).
func Disable() { enabled = false }

// Enable turns on color output.
func Enable() { enabled = true }

func wrap(c, s string) string {
	if !enabled {
		return s
	}
	return c + s + reset
}

// OK formats a success marker.
func OK(msg string) string { return wrap(green, "[OK] "+msg) }

// Fail formats a failure marker.
func Fail(msg string) string { return wrap(red, "[FAIL] "+msg) }

// Warn formats a warning marker.
func Warn(msg string) string { return wrap(yellow, "[WARN] "+msg) }

// Info formats an info marker.
func Info(msg string) string { return wrap(cyan, "[INFO] "+msg) }

// Bold formats text as bold.
func Bold(s string) string { return wrap(bold, s) }

// Dim formats text as dimmed.
func Dim(s string) string { return wrap(dimmed, s) }

// Header formats a section header.
func Header(s string) string { return wrap(bold+cyan, "--- "+s+" ---") }

// Okf is a formatted OK printf.
func Okf(format string, a ...any) string { return OK(fmt.Sprintf(format, a...)) }

// Failf is a formatted Fail printf.
func Failf(format string, a ...any) string { return Fail(fmt.Sprintf(format, a...)) }

// Warnf is a formatted Warn printf.
func Warnf(format string, a ...any) string { return Warn(fmt.Sprintf(format, a...)) }
