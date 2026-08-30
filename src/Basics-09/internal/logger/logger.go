// Package logger is a small leveled logger that writes structured lines to a
// file. It is safe for concurrent use via an internal mutex.
package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelError
)

func ParseLevel(s string) (Level, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "debug":
		return LevelDebug, nil
	case "info", "":
		return LevelInfo, nil
	case "error", "err":
		return LevelError, nil
	}
	return 0, fmt.Errorf("unknown log level: %q", s)
}

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	}
	return "?"
}

type Logger struct {
	mu    sync.Mutex
	w     io.Writer
	close func() error
	min   Level
}

// New opens (or creates) the given file for append and returns a Logger writing
// to it. Passing an empty path returns a logger that writes to stderr.
func New(path string, min Level) (*Logger, error) {
	if path == "" {
		return &Logger{w: os.Stderr, min: min}, nil
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return nil, fmt.Errorf("open log file: %w", err)
	}
	return &Logger{w: f, close: f.Close, min: min}, nil
}

// NewWithWriter is used by tests to inject an in-memory writer.
func NewWithWriter(w io.Writer, min Level) *Logger {
	return &Logger{w: w, min: min}
}

func (l *Logger) Close() error {
	if l.close == nil {
		return nil
	}
	return l.close()
}

// Log writes a structured line. `fields` are appended as key=value pairs.
// Values are formatted with %v; strings containing spaces are quoted.
func (l *Logger) Log(level Level, msg string, fields ...any) {
	if level < l.min {
		return
	}
	if len(fields)%2 != 0 {
		fields = append(fields, "")
	}

	var b strings.Builder
	b.WriteString(time.Now().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(level.String())
	b.WriteByte(' ')
	b.WriteString(msg)
	for i := 0; i < len(fields); i += 2 {
		key := fmt.Sprintf("%v", fields[i])
		val := fmt.Sprintf("%v", fields[i+1])
		b.WriteByte(' ')
		b.WriteString(key)
		b.WriteByte('=')
		if strings.ContainsAny(val, " \t") {
			b.WriteByte('"')
			b.WriteString(val)
			b.WriteByte('"')
		} else {
			b.WriteString(val)
		}
	}
	b.WriteByte('\n')

	l.mu.Lock()
	_, _ = io.WriteString(l.w, b.String())
	l.mu.Unlock()
}

func (l *Logger) Debug(msg string, fields ...any) { l.Log(LevelDebug, msg, fields...) }
func (l *Logger) Info(msg string, fields ...any)  { l.Log(LevelInfo, msg, fields...) }
func (l *Logger) Error(msg string, fields ...any) { l.Log(LevelError, msg, fields...) }
