//go:build unix

/*
Copyright © 2026 envaar
SPDX-License-Identifier: Apache-2.0
*/

package fs

import (
	"errors"
	"path/filepath"
	"syscall"
	"testing"
)

func TestValidateRegularFileRejectsNonRegular(t *testing.T) {
	dir := t.TempDir()
	// A FIFO is neither a directory nor a regular file, so
	// ValidateRegularFile must reject it with ErrNotRegularFile.
	fifo := filepath.Join(dir, "pipe")
	if err := syscall.Mkfifo(fifo, 0o644); err != nil {
		t.Skipf("mkfifo unsupported: %v", err)
	}

	if _, err := ValidateRegularFile(fifo); !errors.Is(err, ErrNotRegularFile) {
		t.Errorf("a FIFO should return ErrNotRegularFile, got %v", err)
	}
}

func TestReadFileRejectsNonRegular(t *testing.T) {
	dir := t.TempDir()
	fifo := filepath.Join(dir, "pipe")
	if err := syscall.Mkfifo(fifo, 0o644); err != nil {
		t.Skipf("mkfifo unsupported: %v", err)
	}

	if _, err := ReadFile(fifo); !errors.Is(err, ErrNotRegularFile) {
		t.Errorf("ReadFile should reject a FIFO with ErrNotRegularFile, got %v", err)
	}
}
