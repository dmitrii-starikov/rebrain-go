package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"syncer/internal/logger"
	"syncer/internal/syncer"
)

func main() {
	var (
		src      = flag.String("src", "", "source directory (required)")
		dst      = flag.String("dst", "", "replica directory (required)")
		interval = flag.Duration("interval", 5*time.Second, "frequency of checking")
		workers  = flag.Int("workers", 4, "worker count")
		logPath  = flag.String("log", "log.txt", "log file path")
		level    = flag.String("level", "info", "log level: debug|info|error")
	)
	flag.Parse()

	if *src == "" || *dst == "" {
		fmt.Fprintln(os.Stderr, "-src and -dst are required")
		flag.Usage()
		os.Exit(2)
	}

	lvl, err := logger.ParseLevel(*level)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	log, err := logger.New(*logPath, lvl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer log.Close()

	// Create a context that cancels on SIGINT/SIGTERM for graceful shutdown.
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Initialize the syncer with CLI config and logger.
	s := syncer.New(syncer.Config{
		Src:      *src,
		Dst:      *dst,
		Interval: *interval,
		Workers:  *workers,
	}, log)

	// Run the syncer; it blocks until the context is canceled/error occurs.
	if err := s.Run(ctx); err != nil {
		log.Error("syncer exited with error", "err", err)
		os.Exit(1)
	}
}
