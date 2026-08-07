package config

import (
	"module07/internal/convertor"
	"testing"
)

func BenchmarkReflectStructToMap(b *testing.B) {
	cfg := Config{
		Name: "test-service",
		Host: "localhost",
		Port: 8080,
		Environment: map[string]string{
			"key": "value",
		},
		Volumes: []string{"/tmp", "/var"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = convertor.StructToMap(cfg)
	}
}

func BenchmarkGeneratedStructToMap(b *testing.B) {
	cfg := Config{
		Name: "test-service",
		Host: "localhost",
		Port: 8080,
		Environment: map[string]string{
			"key": "value",
		},
		Volumes: []string{"/tmp", "/var"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// generated method for struct
		_ = cfg.StructToMap()
	}
}
