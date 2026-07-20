//go:build linux

package vfio

import (
	"fmt"
	"os"
	"syscall"
)

const nsGetParent = 0xb702 // NS_GET_PARENT from linux/nsfs.h

func checkInitialPIDNamespace() error {
	f, err := os.Open("/proc/self/ns/pid")
	if err != nil {
		return fmt.Errorf("cannot verify PID namespace: %w", err)
	}
	defer f.Close()

	parent, _, errno := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), nsGetParent, 0)
	if errno == 0 {
		_ = syscall.Close(int(parent))
		return fmt.Errorf("current process is in a nested PID namespace; refusing VFIO bind against incomplete host mount data")
	}
	if errno != syscall.EPERM {
		return fmt.Errorf("cannot verify PID namespace parent: %w", errno)
	}
	// EPERM means either the initial PID namespace has no parent or its parent
	// is outside this process's user-namespace authority. In the latter case the
	// process also lacks authority to write the parent namespace's host sysfs.
	return nil
}
