// Package fsops contains the primitive filesystem operations used by the
// syncer: atomic file copy, symlink copy, directory create, chmod, delete.
package fsops

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// bufPool holds reusable 32KB buffers for io.CopyBuffer calls, which noticeably
// reduces allocations for many small files.
var bufPool = sync.Pool{
	New: func() any {
		b := make([]byte, 32*1024)
		return &b
	},
}

// CopyFile copies src -> dst atomically. It writes to a temp file in the target
// directory, syncs it, then renames over dst. Permissions and mtime are
// preserved. The context is checked before starting and while copying.
func CopyFile(ctx context.Context, src, dst string, mode os.FileMode, mtime time.Time) (int64, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	in, err := os.Open(src)
	if err != nil {
		return 0, fmt.Errorf("open src: %w", err)
	}
	defer in.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return 0, fmt.Errorf("mkdir parent: %w", err)
	}

	tmp, err := os.CreateTemp(filepath.Dir(dst), ".sync-*.tmp")
	if err != nil {
		return 0, fmt.Errorf("create tmp: %w", err)
	}
	tmpPath := tmp.Name()
	// If anything below fails, clean up the temp file.
	cleanup := true
	defer func() {
		if cleanup {
			_ = os.Remove(tmpPath)
		}
	}()

	bufPtr := bufPool.Get().(*[]byte)
	defer bufPool.Put(bufPtr)

	n, err := copyWithCancel(ctx, tmp, in, *bufPtr)
	if err != nil {
		tmp.Close()
		return n, fmt.Errorf("copy: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return n, fmt.Errorf("sync tmp: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return n, fmt.Errorf("close tmp: %w", err)
	}
	if err := os.Chmod(tmpPath, mode.Perm()); err != nil {
		return n, fmt.Errorf("chmod tmp: %w", err)
	}
	if err := os.Rename(tmpPath, dst); err != nil {
		return n, fmt.Errorf("rename: %w", err)
	}
	cleanup = false
	if !mtime.IsZero() {
		_ = os.Chtimes(dst, time.Now(), mtime)
	}
	return n, nil
}

// copyWithCancel is like io.CopyBuffer but yields to ctx cancellation.
func copyWithCancel(ctx context.Context, dst io.Writer, src io.Reader, buf []byte) (int64, error) {
	var total int64
	for {
		if err := ctx.Err(); err != nil {
			return total, err
		}
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[:nr])
			total += int64(nw)
			if ew != nil {
				return total, ew
			}
			if nw < nr {
				return total, io.ErrShortWrite
			}
		}
		if er != nil {
			if errors.Is(er, io.EOF) {
				return total, nil
			}
			return total, er
		}
	}
}

// CopySymlink reads the link target from src and creates a symlink at dst
// pointing at the same target. It does not follow the link.
func CopySymlink(src, dst string) error {
	target, err := os.Readlink(src)
	if err != nil {
		return fmt.Errorf("readlink: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return fmt.Errorf("mkdir parent: %w", err)
	}
	// If a symlink already exists at dst we replace it. Regular files/dirs
	// at dst are the caller's problem - the differ should have removed them.
	if _, err := os.Lstat(dst); err == nil {
		if err := os.Remove(dst); err != nil {
			return fmt.Errorf("remove existing: %w", err)
		}
	}
	if err := os.Symlink(target, dst); err != nil {
		return fmt.Errorf("symlink: %w", err)
	}
	return nil
}

// Mkdir creates the directory with the given mode, plus any missing parents.
func Mkdir(path string, mode os.FileMode) error {
	if err := os.MkdirAll(path, mode.Perm()); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}
	// MkdirAll may not chmod an already-existing dir; force it.
	return os.Chmod(path, mode.Perm())
}

// Delete removes a file or directory (recursively for directories).
func Delete(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}

// Chmod changes permissions on path. Symlinks are left alone since most
// filesystems do not honor chmod on the link itself.
func Chmod(path string, mode os.FileMode) error {
	info, err := os.Lstat(path)
	if err != nil {
		return fmt.Errorf("lstat: %w", err)
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return nil
	}
	if err := os.Chmod(path, mode.Perm()); err != nil {
		return fmt.Errorf("chmod: %w", err)
	}
	return nil
}
