package wrappers

//go:generate gowrap gen -p module07/internal/monitors -i Monitor -t log -o ./monitor_with_log.go
//go:generate gowrap gen -p module07/internal/monitors -i Monitor -t prometheus -o ./monitor_with_metrics.go
//go:generate gowrap gen -p module07/internal/monitors -i Monitor -t ../../templates/recovery.tmpl -o ./monitor_with_recovery.go
