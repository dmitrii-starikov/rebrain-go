package scanner

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestScanBasic(t *testing.T) {
	root := t.TempDir()
	writeTree(t, root, map[string]string{
		"a.txt":       "hello",
		"sub/b.txt":   "world",
		"sub/c/d.txt": "deep",
	})

	snap, err := Scan(context.Background(), root)
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	want := []string{"a.txt", "sub", "sub/b.txt", "sub/c", "sub/c/d.txt"}
	for _, rel := range want {
		if _, ok := snap[filepath.FromSlash(rel)]; !ok {
			t.Errorf("missing %q in snapshot", rel)
		}
	}
	if _, ok := snap[""]; ok {
		t.Errorf("root should not appear in snapshot")
	}

	if e := snap[filepath.FromSlash("sub/b.txt")]; e.Size != 5 {
		t.Errorf("size of sub/b.txt = %d, want 5", e.Size)
	}
	if e := snap["sub"]; !e.IsDir {
		t.Errorf("sub should be a directory")
	}
}

func TestScanSymlink(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "target"), []byte("t"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink("target", filepath.Join(root, "link")); err != nil {
		t.Skipf("symlinks not supported: %v", err)
	}

	snap, err := Scan(context.Background(), root)
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	link, ok := snap["link"]
	if !ok {
		t.Fatalf("link not in snapshot")
	}
	if !link.IsSymlink {
		t.Errorf("IsSymlink=false")
	}
	if link.LinkTarget != "target" {
		t.Errorf("LinkTarget = %q, want target", link.LinkTarget)
	}
}

func TestScanCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := Scan(ctx, t.TempDir()); err == nil {
		t.Errorf("expected cancelled error")
	}
}

func writeTree(t *testing.T, root string, files map[string]string) {
	t.Helper()
	for rel, content := range files {
		full := filepath.Join(root, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}
