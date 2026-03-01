package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/dmitrii-starikov/semaphore"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	linksMap       = make(map[string]bool, 10)
	fileExtensions = []string{
		"jpg", "jpeg", "png", "gif", "bmp", "webp", "svg",
		"pdf", "doc", "docx", "xls", "xlsx", "ppt", "pptx",
		"txt", "csv", "zip", "rar", "7z", "tar", "gz",
		"mp3", "mp4", "avi", "mov", "wmv", "flv",
		"exe", "msi", "dmg", "apk", "iso", "bin",
	}
	fileRegex  *regexp.Regexp
	httpClient *http.Client
)

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:    10,
	}

	httpClient = &http.Client{
		Transport: tr,
	}

	fileRegex = regexp.MustCompile(`^https?://[^\s]+\.(` + strings.Join(fileExtensions, "|") + `)(\?[^\s]*)?$`)
}

func main() {
	threads := flag.Int("threads", 2, "number of threads")
	timeout := flag.Int("timeout", 120, "general timeout in seconds")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("No links passed")
		return
	}

	addedCount := 0
	for _, arg := range args {
		if fileRegex.MatchString(strings.ToLower(arg)) && !linksMap[arg] {
			linksMap[arg] = true
			addedCount++
		}
	}

	if 0 == addedCount {
		fmt.Println("No valid file links passed")
		return
	}

	fmt.Printf("Starting downloading %d links in %d threads\n", addedCount, *threads)

	s := semaphore.NewSemaphore(*threads)
	errGr, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(*timeout))
	defer cancel()

	for url := range linksMap {
		url := url // closure capture

		errGr.Go(func() error {
			_, err := downloadFile(url, ctx, s)
			if err != nil {
				return fmt.Errorf("error while download %s: %w", url, err)
			}
			return nil
		})
	}

	if err := errGr.Wait(); err != nil {
		fmt.Printf("\nError: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nAll downloads completed successfully!")
}

func downloadFile(url string, ctx context.Context, s *semaphore.Semaphore) (string, error) {
	s.Acquire()
	defer s.Release()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	filename := filepath.Base(url)
	if filename == "" || filename == "/" || filename == "." {
		filename = fmt.Sprintf("file_%d", time.Now().UnixNano())
	}

	filename = strings.Split(filename, "?")[0]

	localPath := filename
	for i := 1; i < 1000; i++ {
		_, err := os.Stat(localPath)
		if err != nil {
			if os.IsNotExist(err) {
				break // doesn't exist, ок
			}
			// permissions?
			return "", fmt.Errorf("can't check file %s: %w", localPath, err)
		}
		ext := filepath.Ext(filename)
		name := strings.TrimSuffix(filename, ext)
		localPath = fmt.Sprintf("%s_%d%s", name, i, ext)

	}

	fmt.Printf("Started: %s\n", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(localPath)
	if err != nil {
		return "", fmt.Errorf("can't create file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		os.Remove(localPath)
		return "", fmt.Errorf("can't save file: %w", err)
	}

	fmt.Printf("Completed: %s -> %s\n", url, localPath)
	return localPath, nil
}
