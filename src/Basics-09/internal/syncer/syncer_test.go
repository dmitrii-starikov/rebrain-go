package syncer

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"syncer/internal/logger"
)

func newTestSyncer(t *testing.T, src, dst string) *Syncer {
	t.Helper()
	log := logger.NewWithWriter(&bytes.Buffer{}, logger.LevelError)
	return New(Config{
		Src:      src,
		Dst:      dst,
		Interval: 10 * time.Millisecond,
		Workers:  2,
	}, log)
}

func TestReconcileInitialCopy(t *testing.T) {
	src := t.TempDir()
	dst := t.TempDir()
	writeFile(t, filepath.Join(src, "a.txt"), "hello")
	writeFile(t, filepath.Join(src, "sub", "b.txt"), "world")

	s := newTestSyncer(t, src, dst)
	s.Reconcile(context.Background())

	assertFile(t, filepath.Join(dst, "a.txt"), "hello")
	assertFile(t, filepath.Join(dst, "sub", "b.txt"), "world")
}

func TestReconcileUpdate(t *testing.T) {
	src := t.TempDir()
	dst := t.TempDir()
	writeFile(t, filepath.Join(src, "f"), "v1")
	writeFile(t, filepath.Join(dst, "f"), "v0")
	// Backdate dst so size+mtime comparison spots the change.
	old := time.Now().Add(-time.Hour)
	if err := os.Chtimes(filepath.Join(dst, "f"), old, old); err != nil {
		t.Fatal(err)
	}

	s := newTestSyncer(t, src, dst)
	s.Reconcile(context.Background())

	assertFile(t, filepath.Join(dst, "f"), "v1")
}

func TestReconcileDelete(t *testing.T) {
	src := t.TempDir()
	dst := t.TempDir()
	writeFile(t, filepath.Join(dst, "gone.txt"), "x")
	writeFile(t, filepath.Join(dst, "sub", "still.txt"), "x")

	s := newTestSyncer(t, src, dst)
	s.Reconcile(context.Background())

	if _, err := os.Stat(filepath.Join(dst, "gone.txt")); !os.IsNotExist(err) {
		t.Errorf("gone.txt should be removed, got %v", err)
	}
	if _, err := os.Stat(filepath.Join(dst, "sub")); !os.IsNotExist(err) {
		t.Errorf("sub should be removed, got %v", err)
	}
}

func TestReconcileSymlink(t *testing.T) {
	src := t.TempDir()
	dst := t.TempDir()
	writeFile(t, filepath.Join(src, "real"), "hi")
	if err := os.Symlink("real", filepath.Join(src, "link")); err != nil {
		t.Skipf("symlinks not supported: %v", err)
	}

	s := newTestSyncer(t, src, dst)
	s.Reconcile(context.Background())

	target, err := os.Readlink(filepath.Join(dst, "link"))
	if err != nil {
		t.Fatal(err)
	}
	if target != "real" {
		t.Errorf("target = %q, want real", target)
	}
}

func TestReconcileIdempotent(t *testing.T) {
	src := t.TempDir()
	dst := t.TempDir()
	writeFile(t, filepath.Join(src, "a"), "one")
	writeFile(t, filepath.Join(src, "b/c"), "two")

	s := newTestSyncer(t, src, dst)
	s.Reconcile(context.Background())

	// Snapshot dst mtimes, run again, ensure nothing changed.
	before := statAll(t, dst)
	time.Sleep(20 * time.Millisecond)
	s.Reconcile(context.Background())
	after := statAll(t, dst)

	for k, v := range before {
		if after[k] != v {
			t.Errorf("mtime changed for %s: %v → %v", k, v, after[k])
		}
	}
}

func TestRunCancels(t *testing.T) {
	src := t.TempDir()
	dst := t.TempDir()
	writeFile(t, filepath.Join(src, "a"), "x")

	s := newTestSyncer(t, src, dst)
	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan error, 1)
	go func() { done <- s.Run(ctx) }()

	time.Sleep(30 * time.Millisecond)
	cancel()

	select {
	case err := <-done:
		if err != nil {
			t.Errorf("Run returned err: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("Run did not exit after cancel")
	}
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func assertFile(t *testing.T, path, want string) {
	t.Helper()
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if string(got) != want {
		t.Errorf("%s = %q, want %q", path, got, want)
	}
}

func statAll(t *testing.T, root string) map[string]time.Time {
	t.Helper()
	out := map[string]time.Time{}
	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		out[path] = info.ModTime()
		return nil
	})
	return out
}
