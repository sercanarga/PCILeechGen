//go:build !linux

package vfio

func checkInitialPIDNamespace() error { return nil }
