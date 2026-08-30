package fsops

import (
	"context"
	"crypto/rand"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func benchCopy(b *testing.B, size int) {
	dir := b.TempDir()
	src := filepath.Join(dir, "src.bin")
	buf := make([]byte, size)
	if _, err := rand.Read(buf); err != nil {
		b.Fatal(err)
	}
	if err := os.WriteFile(src, buf, 0o600); err != nil {
		b.Fatal(err)
	}
	ctx := context.Background()
	b.ResetTimer()
	b.SetBytes(int64(size))
	for i := 0; i < b.N; i++ {
		dst := filepath.Join(dir, "d")
		if _, err := CopyFile(ctx, src, dst, 0o600, time.Time{}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCopyFile1KB(b *testing.B)  { benchCopy(b, 1<<10) }
func BenchmarkCopyFile1MB(b *testing.B)  { benchCopy(b, 1<<20) }
func BenchmarkCopyFile10MB(b *testing.B) { benchCopy(b, 10<<20) }
