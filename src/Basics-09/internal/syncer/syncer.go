// Package syncer wires scanner + differ + fsops behind a ticker-driven
// loop. Operations from each diff are dispatched to a fixed
// pool of workers guarded by sync.WaitGroup and cancelled via context.
package syncer

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"syncer/internal/differ"
	"syncer/internal/fsops"
	"syncer/internal/logger"
	"syncer/internal/scanner"
)

type Config struct {
	Src      string
	Dst      string
	Interval time.Duration
	Workers  int
}

type Syncer struct {
	cfg Config
	log *logger.Logger
}

func New(cfg Config, log *logger.Logger) *Syncer {
	if cfg.Workers <= 0 {
		cfg.Workers = 4
	}
	if cfg.Interval <= 0 {
		cfg.Interval = 5 * time.Second
	}
	return &Syncer{cfg: cfg, log: log}
}

// Run until ctx is cancelled. It performs one reconciliation
// immediately, then continues on the configured interval.
func (s *Syncer) Run(ctx context.Context) error {
	if err := os.MkdirAll(s.cfg.Dst, 0o755); err != nil {
		return fmt.Errorf("prepare dst: %w", err)
	}

	s.log.Info("syncer started",
		"src", s.cfg.Src, "dst", s.cfg.Dst,
		"interval", s.cfg.Interval, "workers", s.cfg.Workers)

	s.Reconcile(ctx)

	ticker := time.NewTicker(s.cfg.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			s.log.Info("syncer stopped", "reason", ctx.Err())
			return nil
		case <-ticker.C:
			s.Reconcile(ctx)
		}
	}
}

// Reconcile does one scan - check diff - apply cycle. Errors are logged and swallowed,
// so a single failing operation doesn't kill the loop.
func (s *Syncer) Reconcile(ctx context.Context) {
	start := time.Now()

	srcSnap, err := scanner.Scan(ctx, s.cfg.Src)
	if err != nil {
		s.log.Error("scan src", "err", err)
		return
	}
	dstSnap, err := scanner.Scan(ctx, s.cfg.Dst)
	if err != nil {
		s.log.Error("scan dst", "err", err)
		return
	}

	ops := differ.Diff(srcSnap, dstSnap)
	if len(ops) == 0 {
		s.log.Debug("reconcile: no changes", "took", time.Since(start))
		return
	}

	s.log.Info("reconcile: applying ops", "count", len(ops))
	s.apply(ctx, ops)
	s.log.Info("reconcile done", "took", time.Since(start))
}

// apply groups ops by type (mkdir - copy/symlink - chmod - delete) and runs
// each type's ops through a worker pool.
func (s *Syncer) apply(ctx context.Context, ops []differ.Op) {
	mkdirs := opsOfKind(ops, differ.OpMkdir)
	copies := append(
		opsOfKind(ops, differ.OpCopyFile),
		opsOfKind(ops, differ.OpCopySymlink)...,
	)
	chmods := opsOfKind(ops, differ.OpChmod)
	deletes := opsOfKind(ops, differ.OpDelete)

	// Mkdirs are already sorted shallow -> deep by the differ. Doing them
	// on a single goroutine avoids parent/child races and is cheap.
	for _, op := range mkdirs {
		if ctx.Err() != nil {
			return
		}
		s.runOne(ctx, op)
	}

	s.runPool(ctx, copies)
	s.runPool(ctx, chmods)

	// Deletes are sorted deep → shallow. Same reasoning: single-threaded.
	for _, op := range deletes {
		if ctx.Err() != nil {
			return
		}
		s.runOne(ctx, op)
	}
}

func (s *Syncer) runPool(ctx context.Context, ops []differ.Op) {
	if len(ops) == 0 {
		return
	}
	n := s.cfg.Workers
	if n > len(ops) {
		n = len(ops)
	}
	jobs := make(chan differ.Op)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for op := range jobs {
				if ctx.Err() != nil {
					continue // drain until channel closes
				}
				s.runOne(ctx, op)
			}
		}()
	}
	for _, op := range ops {
		select {
		case <-ctx.Done():
			// still need to close jobs and let workers exit cleanly
			close(jobs)
			wg.Wait()
			return
		case jobs <- op:
		}
	}
	close(jobs)
	wg.Wait()
}

func (s *Syncer) runOne(ctx context.Context, op differ.Op) {
	dst := filepath.Join(s.cfg.Dst, op.RelPath)
	src := filepath.Join(s.cfg.Src, op.RelPath)

	switch op.Kind {
	case differ.OpMkdir:
		if err := fsops.Mkdir(dst, op.Mode); err != nil {
			s.log.Error("mkdir", "path", op.RelPath, "err", err)
			return
		}
		s.log.Info("mkdir", "path", op.RelPath, "mode", op.Mode.Perm())
	case differ.OpCopyFile:
		n, err := fsops.CopyFile(ctx, src, dst, op.Mode, op.ModTime)
		if err != nil {
			s.log.Error("copy", "path", op.RelPath, "err", err)
			return
		}
		s.log.Info("copy", "path", op.RelPath, "size", n, "mode", op.Mode.Perm())
	case differ.OpCopySymlink:
		if err := fsops.CopySymlink(src, dst); err != nil {
			s.log.Error("symlink", "path", op.RelPath, "err", err)
			return
		}
		s.log.Info("symlink", "path", op.RelPath, "target", op.LinkTarget)
	case differ.OpChmod:
		if err := fsops.Chmod(dst, op.Mode); err != nil {
			s.log.Error("chmod", "path", op.RelPath, "err", err)
			return
		}
		s.log.Info("chmod", "path", op.RelPath, "mode", op.Mode.Perm())
	case differ.OpDelete:
		if err := fsops.Delete(dst); err != nil {
			s.log.Error("delete", "path", op.RelPath, "err", err)
			return
		}
		s.log.Info("delete", "path", op.RelPath)
	}
}

func opsOfKind(ops []differ.Op, k differ.OpKind) []differ.Op {
	var out []differ.Op
	for _, op := range ops {
		if op.Kind == k {
			out = append(out, op)
		}
	}
	return out
}
