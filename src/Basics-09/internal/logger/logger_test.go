package logger

import (
	"bytes"
	"strings"
	"sync"
	"testing"
)

func TestParseLevel(t *testing.T) {
	cases := map[string]Level{
		"debug": LevelDebug,
		"INFO":  LevelInfo,
		"":      LevelInfo,
		"error": LevelError,
	}
	for in, want := range cases {
		got, err := ParseLevel(in)
		if err != nil {
			t.Fatalf("ParseLevel(%q): %v", in, err)
		}
		if got != want {
			t.Errorf("ParseLevel(%q) = %v, want %v", in, got, want)
		}
	}
	if _, err := ParseLevel("bogus"); err == nil {
		t.Errorf("expected error for bogus level")
	}
}

func TestLogFiltersByMinLevel(t *testing.T) {
	var buf bytes.Buffer
	log := NewWithWriter(&buf, LevelInfo)
	log.Debug("hidden")
	log.Info("shown", "k", "v")

	out := buf.String()
	if strings.Contains(out, "hidden") {
		t.Errorf("debug should be filtered: %q", out)
	}
	if !strings.Contains(out, "shown k=v") {
		t.Errorf("info line missing fields: %q", out)
	}
	if !strings.Contains(out, "INFO") {
		t.Errorf("missing level label: %q", out)
	}
}

func TestLogQuotesSpaces(t *testing.T) {
	var buf bytes.Buffer
	log := NewWithWriter(&buf, LevelDebug)
	log.Info("m", "path", "with space")
	if !strings.Contains(buf.String(), `path="with space"`) {
		t.Errorf("value with space not quoted: %q", buf.String())
	}
}

func TestLogConcurrent(t *testing.T) {
	var buf bytes.Buffer
	log := NewWithWriter(&buf, LevelDebug)

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Info("msg", "i", i)
		}(i)
	}
	wg.Wait()

	lines := strings.Count(buf.String(), "\n")
	if lines != 50 {
		t.Errorf("expected 50 lines, got %d", lines)
	}
}
