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

func TestStatRegularFileRejectsNonRegular(t *testing.T) {
	dir := t.TempDir()
	// A FIFO is neither a directory nor a regular file, so StatRegularFile must
	// reject it with ErrNotRegular (which readDiffFile relies on for its wording).
	fifo := filepath.Join(dir, "pipe")
	if err := syscall.Mkfifo(fifo, 0o644); err != nil {
		t.Skipf("mkfifo unsupported: %v", err)
	}

	if _, err := StatRegularFile(fifo); !errors.Is(err, ErrNotRegular) {
		t.Errorf("a FIFO should return ErrNotRegular, got %v", err)
	}
}
