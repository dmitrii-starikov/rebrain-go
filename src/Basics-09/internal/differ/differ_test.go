package differ

import (
	"os"
	"testing"
	"time"

	"syncer/internal/scanner"
)

func file(rel string, size int64, mtime time.Time, mode os.FileMode) scanner.Entry {
	return scanner.Entry{RelPath: rel, Size: size, ModTime: mtime, Mode: mode}
}
func dir(rel string, mode os.FileMode) scanner.Entry {
	return scanner.Entry{RelPath: rel, IsDir: true, Mode: mode | os.ModeDir}
}
func link(rel, target string) scanner.Entry {
	return scanner.Entry{RelPath: rel, IsSymlink: true, LinkTarget: target, Mode: os.ModeSymlink}
}

func TestDiffAllNew(t *testing.T) {
	t0 := time.Unix(1000, 0)
	src := scanner.Snapshot{
		"a":     dir("a", 0o755),
		"a/b":   dir("a/b", 0o755),
		"a/b/f": file("a/b/f", 3, t0, 0o644),
		"lnk":   link("lnk", "a/b/f"),
	}
	ops := Diff(src, scanner.Snapshot{})
	if len(ops) != 4 {
		t.Fatalf("expected 4 ops, got %d: %+v", len(ops), ops)
	}
	// mkdirs first, in shallow → deep order
	if ops[0].Kind != OpMkdir || ops[0].RelPath != "a" {
		t.Errorf("op[0] = %+v", ops[0])
	}
	if ops[1].Kind != OpMkdir || ops[1].RelPath != "a/b" {
		t.Errorf("op[1] = %+v", ops[1])
	}
}

func TestDiffDeleteExtras(t *testing.T) {
	dst := scanner.Snapshot{
		"a":     dir("a", 0o755),
		"a/x":   file("a/x", 1, time.Unix(1, 0), 0o644),
		"a/y/z": file("a/y/z", 1, time.Unix(1, 0), 0o644),
		"a/y":   dir("a/y", 0o755),
	}
	ops := Diff(scanner.Snapshot{}, dst)
	// All deletes; deepest first.
	if ops[0].Kind != OpDelete || ops[0].RelPath != "a/y/z" {
		t.Errorf("first delete = %+v", ops[0])
	}
	last := ops[len(ops)-1]
	if last.RelPath != "a" {
		t.Errorf("last delete = %+v (want top-level)", last)
	}
}

func TestDiffFileChanged(t *testing.T) {
	t0 := time.Unix(1000, 0)
	t1 := time.Unix(2000, 0)
	src := scanner.Snapshot{"f": file("f", 5, t1, 0o644)}
	dst := scanner.Snapshot{"f": file("f", 5, t0, 0o644)}
	ops := Diff(src, dst)
	if len(ops) != 1 || ops[0].Kind != OpCopyFile {
		t.Fatalf("expected 1 copy, got %+v", ops)
	}
}

func TestDiffFileUnchanged(t *testing.T) {
	t0 := time.Unix(1, 0)
	src := scanner.Snapshot{"f": file("f", 5, t0, 0o644)}
	dst := scanner.Snapshot{"f": file("f", 5, t0, 0o644)}
	if ops := Diff(src, dst); len(ops) != 0 {
		t.Fatalf("expected 0 ops, got %+v", ops)
	}
}

func TestDiffModeOnly(t *testing.T) {
	t0 := time.Unix(1, 0)
	src := scanner.Snapshot{"f": file("f", 5, t0, 0o755)}
	dst := scanner.Snapshot{"f": file("f", 5, t0, 0o644)}
	ops := Diff(src, dst)
	if len(ops) != 1 || ops[0].Kind != OpChmod {
		t.Fatalf("expected 1 chmod, got %+v", ops)
	}
}

func TestDiffTypeMismatch(t *testing.T) {
	src := scanner.Snapshot{"x": dir("x", 0o755)}
	dst := scanner.Snapshot{"x": file("x", 1, time.Unix(1, 0), 0o644)}
	ops := Diff(src, dst)
	// Expect delete + mkdir.
	if len(ops) != 2 {
		t.Fatalf("expected 2 ops, got %+v", ops)
	}
	kinds := map[OpKind]bool{}
	for _, o := range ops {
		kinds[o.Kind] = true
	}
	if !kinds[OpDelete] || !kinds[OpMkdir] {
		t.Errorf("expected delete+mkdir, got %+v", ops)
	}
}

func TestDiffSymlinkTargetChanged(t *testing.T) {
	src := scanner.Snapshot{"l": link("l", "new")}
	dst := scanner.Snapshot{"l": link("l", "old")}
	ops := Diff(src, dst)
	if len(ops) != 1 || ops[0].Kind != OpCopySymlink || ops[0].LinkTarget != "new" {
		t.Fatalf("unexpected ops: %+v", ops)
	}
}
