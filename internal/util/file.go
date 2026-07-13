package util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CopyFile copies a single file from src to dst.
func CopyFile(src, dst string) error {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)
	if src == dst {
		return nil
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	srcInfo, err := in.Stat()
	if err != nil {
		return err
	}
	if !srcInfo.Mode().IsRegular() {
		return fmt.Errorf("copy source %q: not a regular file", src)
	}

	parent := filepath.Dir(dst)
	parentInfo, err := os.Lstat(parent)
	if err != nil {
		return fmt.Errorf("inspect destination directory %q: %w", parent, err)
	}
	if parentInfo.Mode()&os.ModeSymlink != 0 || !parentInfo.IsDir() {
		return fmt.Errorf("destination directory %q is not a real directory", parent)
	}
	if existing, err := os.Lstat(dst); err == nil {
		if !existing.Mode().IsRegular() {
			return fmt.Errorf("refusing to replace non-regular destination %q", dst)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("inspect destination %q: %w", dst, err)
	}

	// Never open dst for truncation: a pre-existing symlink could otherwise
	// redirect a copy into an arbitrary file. Write a sibling temporary file
	// and atomically replace only the directory entry at dst instead.
	out, err := os.CreateTemp(parent, "."+filepath.Base(dst)+".tmp-*")
	if err != nil {
		return fmt.Errorf("create temporary destination: %w", err)
	}
	tmpName := out.Name()
	defer func() {
		_ = out.Close()
		_ = os.Remove(tmpName)
	}()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	if err := out.Chmod(srcInfo.Mode().Perm() & 0o666); err != nil {
		return fmt.Errorf("set destination mode: %w", err)
	}
	if err := out.Sync(); err != nil {
		return fmt.Errorf("sync destination: %w", err)
	}
	if err := out.Close(); err != nil {
		return fmt.Errorf("close destination: %w", err)
	}
	if err := os.Rename(tmpName, dst); err != nil {
		return fmt.Errorf("replace destination: %w", err)
	}
	return nil
}

// CopyDir recursively copies a directory from src to dst.
func CopyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if mkdirErr := os.MkdirAll(dst, srcInfo.Mode()); mkdirErr != nil {
		return mkdirErr
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := CopyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := CopyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}
