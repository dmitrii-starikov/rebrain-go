package fsops

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCopyFileBasic(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "src.txt")
	dst := filepath.Join(dir, "sub", "dst.txt")
	content := []byte("hello syncer")
	if err := os.WriteFile(src, content, 0o600); err != nil {
		t.Fatal(err)
	}
	mtime := time.Now().Add(-time.Hour).Truncate(time.Second)

	n, err := CopyFile(context.Background(), src, dst, 0o640, mtime)
	if err != nil {
		t.Fatalf("CopyFile: %v", err)
	}
	if int(n) != len(content) {
		t.Errorf("copied %d bytes, want %d", n, len(content))
	}
	got, err := os.ReadFile(dst)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(content) {
		t.Errorf("content mismatch: %q", got)
	}
	info, err := os.Lstat(dst)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0o640 {
		t.Errorf("mode = %v, want 0640", info.Mode().Perm())
	}
	if !info.ModTime().Equal(mtime) {
		t.Errorf("mtime = %v, want %v", info.ModTime(), mtime)
	}
}

func TestCopyFileOverwrite(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "src.txt")
	dst := filepath.Join(dir, "dst.txt")
	if err := os.WriteFile(src, []byte("new"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(dst, []byte("old-old-old"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := CopyFile(context.Background(), src, dst, 0o600, time.Time{}); err != nil {
		t.Fatalf("CopyFile: %v", err)
	}
	got, _ := os.ReadFile(dst)
	if string(got) != "new" {
		t.Errorf("dst = %q, want new", got)
	}
}

func TestCopyFileCancelled(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "s")
	if err := os.WriteFile(src, []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := CopyFile(ctx, src, filepath.Join(dir, "d"), 0o600, time.Time{}); err == nil {
		t.Errorf("expected error on cancelled ctx")
	}
}

func TestCopySymlink(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "link")
	if err := os.Symlink("/tmp/nowhere", src); err != nil {
		t.Skipf("symlinks not supported: %v", err)
	}
	dst := filepath.Join(dir, "sub", "link-copy")
	if err := CopySymlink(src, dst); err != nil {
		t.Fatalf("CopySymlink: %v", err)
	}
	target, err := os.Readlink(dst)
	if err != nil {
		t.Fatal(err)
	}
	if target != "/tmp/nowhere" {
		t.Errorf("target = %q", target)
	}
}

func TestCopySymlinkReplacesExisting(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "link")
	if err := os.Symlink("/a", src); err != nil {
		t.Skipf("symlinks not supported: %v", err)
	}
	dst := filepath.Join(dir, "dst")
	if err := os.Symlink("/b", dst); err != nil {
		t.Fatal(err)
	}
	if err := CopySymlink(src, dst); err != nil {
		t.Fatalf("CopySymlink: %v", err)
	}
	target, _ := os.Readlink(dst)
	if target != "/a" {
		t.Errorf("target = %q", target)
	}
}

func TestMkdirAndChmod(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "a", "b")
	if err := Mkdir(p, 0o750); err != nil {
		t.Fatalf("Mkdir: %v", err)
	}
	info, _ := os.Stat(p)
	if info.Mode().Perm() != 0o750 {
		t.Errorf("mkdir mode = %v", info.Mode().Perm())
	}
	if err := Chmod(p, 0o700); err != nil {
		t.Fatalf("Chmod: %v", err)
	}
	info, _ = os.Stat(p)
	if info.Mode().Perm() != 0o700 {
		t.Errorf("chmod mode = %v", info.Mode().Perm())
	}
}

func TestDelete(t *testing.T) {
	dir := t.TempDir()
	nested := filepath.Join(dir, "a", "b", "c.txt")
	_ = os.MkdirAll(filepath.Dir(nested), 0o755)
	_ = os.WriteFile(nested, []byte("x"), 0o600)
	if err := Delete(filepath.Join(dir, "a")); err != nil {
		t.Fatalf("Delete: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "a")); !os.IsNotExist(err) {
		t.Errorf("expected not-exists, got %v", err)
	}
}
