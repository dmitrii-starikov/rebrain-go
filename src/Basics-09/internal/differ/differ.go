// Package differ compares a source snapshot with a destination snapshot and
// produces an ordered list of operations that would make dst equal to src.
package differ

import (
	"os"
	"sort"
	"strings"
	"time"

	"syncer/internal/scanner"
)

type OpKind int

const (
	OpMkdir OpKind = iota
	OpCopyFile
	OpCopySymlink
	OpChmod
	OpDelete
)

func (k OpKind) String() string {
	switch k {
	case OpMkdir:
		return "mkdir"
	case OpCopyFile:
		return "copy"
	case OpCopySymlink:
		return "symlink"
	case OpChmod:
		return "chmod"
	case OpDelete:
		return "delete"
	}
	return "?"
}

// Op is a single operation the syncer must perform against the destination.
// RelPath is relative to the destination root. Size/Mode/ModTime carry the
// intent from the source (filled in only where meaningful).
type Op struct {
	Kind    OpKind
	RelPath string
	Mode    os.FileMode
	Size    int64
	ModTime time.Time
	// LinkTarget is only set for OpCopySymlink.
	LinkTarget string
}

// Diff returns a list of operations ordered so that they can be executed
// sequentially or concurrently within each phase:
//  1. mkdir (parents before children)
//  2. copy files / copy symlinks
//  3. chmod (existing files whose only difference is permissions)
//  4. delete (children before parents)
func Diff(src, dst scanner.Snapshot) []Op {
	var mkdirs, copies, chmods, deletes []Op

	for rel, s := range src {
		d, existsInDst := dst[rel]

		if !existsInDst {
			mkdirs, copies = appendCreate(mkdirs, copies, s)
			continue
		}

		// Type mismatch → delete the wrong-type entry, then recreate.
		if entryKind(s) != entryKind(d) {
			deletes = append(deletes, Op{Kind: OpDelete, RelPath: rel})
			mkdirs, copies = appendCreate(mkdirs, copies, s)
			continue
		}

		switch {
		case s.IsDir:
			if s.Mode.Perm() != d.Mode.Perm() {
				chmods = append(chmods, Op{Kind: OpChmod, RelPath: rel, Mode: s.Mode})
			}
		case s.IsSymlink:
			if s.LinkTarget != d.LinkTarget {
				copies = append(copies, Op{
					Kind: OpCopySymlink, RelPath: rel, LinkTarget: s.LinkTarget,
				})
			}
		default: // regular file
			if s.Size != d.Size || !s.ModTime.Equal(d.ModTime) {
				copies = append(copies, Op{
					Kind: OpCopyFile, RelPath: rel,
					Mode: s.Mode, Size: s.Size, ModTime: s.ModTime,
				})
			} else if s.Mode.Perm() != d.Mode.Perm() {
				chmods = append(chmods, Op{Kind: OpChmod, RelPath: rel, Mode: s.Mode})
			}
		}
	}

	for rel := range dst {
		if _, ok := src[rel]; !ok {
			deletes = append(deletes, Op{Kind: OpDelete, RelPath: rel})
		}
	}

	// Order matters: mkdirs shallowest first, deletes deepest first.
	sort.Slice(mkdirs, func(i, j int) bool {
		return depth(mkdirs[i].RelPath) < depth(mkdirs[j].RelPath)
	})
	sort.Slice(deletes, func(i, j int) bool {
		return depth(deletes[i].RelPath) > depth(deletes[j].RelPath)
	})

	out := make([]Op, 0, len(mkdirs)+len(copies)+len(chmods)+len(deletes))
	out = append(out, mkdirs...)
	out = append(out, copies...)
	out = append(out, chmods...)
	out = append(out, deletes...)
	return out
}

func appendCreate(mkdirs, copies []Op, s scanner.Entry) ([]Op, []Op) {
	switch {
	case s.IsDir:
		mkdirs = append(mkdirs, Op{Kind: OpMkdir, RelPath: s.RelPath, Mode: s.Mode})
	case s.IsSymlink:
		copies = append(copies, Op{
			Kind: OpCopySymlink, RelPath: s.RelPath, LinkTarget: s.LinkTarget,
		})
	default:
		copies = append(copies, Op{
			Kind: OpCopyFile, RelPath: s.RelPath,
			Mode: s.Mode, Size: s.Size, ModTime: s.ModTime,
		})
	}
	return mkdirs, copies
}

func entryKind(e scanner.Entry) int {
	switch {
	case e.IsSymlink:
		return 2
	case e.IsDir:
		return 1
	default:
		return 0
	}
}

func depth(p string) int {
	if p == "" || p == "." {
		return 0
	}
	return strings.Count(p, string(os.PathSeparator)) + 1
}
