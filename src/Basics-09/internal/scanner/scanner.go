// Package scanner walks a directory tree and produces a Snapshot
package scanner

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

// Entry describes a single filesystem object in a snapshot.
// For symlinks, Mode has os.ModeSymlink set and LinkTarget is populated
// with the link's contents (not the resolved path).
type Entry struct {
	RelPath    string
	Mode       os.FileMode
	Size       int64
	ModTime    time.Time
	IsDir      bool
	IsSymlink  bool
	LinkTarget string
}

// Snapshot is a directory tree keyed by path relative to the scan root.
// The root itself is not included.
type Snapshot map[string]Entry

// Scan walks root and returns a Snapshot. It uses Lstat so symlinks are
// captured as symlinks rather than what they point to.
func Scan(ctx context.Context, root string) (Snapshot, error) {
	snap := make(Snapshot)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		if path == root {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return fmt.Errorf("rel: %w", err)
		}

		info, err := os.Lstat(path)
		if err != nil {
			return fmt.Errorf("lstat %s: %w", path, err)
		}
		entry := Entry{
			RelPath: rel,
			Mode:    info.Mode(),
			Size:    info.Size(),
			ModTime: info.ModTime(),
			IsDir:   info.IsDir(),
		}
		if info.Mode()&os.ModeSymlink != 0 {
			entry.IsSymlink = true
			target, err := os.Readlink(path)
			if err != nil {
				return fmt.Errorf("readlink %s: %w", path, err)
			}
			entry.LinkTarget = target
		}
		snap[rel] = entry
		return nil
	})
	if err != nil {
		return nil, err
	}
	return snap, nil
}
