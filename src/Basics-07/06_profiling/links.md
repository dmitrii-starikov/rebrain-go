# Go Profiling with pprof

This chapter covers built‑in Go profiling tools for performance analysis:

- **pprof** – collects CPU, memory, goroutine, mutex, and blocking profiles
- **go test** – generates profiles during benchmarks (`-cpuprofile`, `-memprofile`)
- **net/http/pprof** – exposes live profiles from running services
- **go tool pprof** – analyzes profiles via terminal or web UI (`-http=:port`)
- **go tool trace** – captures detailed execution timelines for debugging latency

**Main use cases:**
- Find CPU bottlenecks
- Detect memory leaks
- Identify goroutine leaks
- Analyze mutex contention
- Trace blocking operations

## Commands & Explanations
### 1. Generate profiles from benchmarks

`go test -bench=. -benchtime=5s -cpuprofile=cpu.profile -memprofile=mem.profile`

- `-bench=.` – run all benchmarks
- `-benchtime=5s` – run each benchmark for 5 seconds
- `-cpuprofile` – save CPU profile to file
- `-memprofile` – save memory profile to file

### 2. View profile in web UI
`go tool pprof -http=:9090 cpu.profile`

- Opens interactive web interface at port 9090
- Shows flame graph, call graph, top functions

### 3. Profile live service – CPU
`go tool pprof -http=:9090 http://localhost:6060/debug/pprof/profile?seconds=5`

- Samples CPU usage for 5 seconds
- Requires `import _ "net/http/pprof"` in the service

### 4. Profile live service – Memory (heap)
`go tool pprof -http=:9090 http://localhost:6060/debug/pprof/heap`

- Shows current live memory allocations
- Good for finding memory leaks

### 5. Profile live service – All allocations
`go tool pprof -http=:9090 http://localhost:6060/debug/pprof/allocs`

- Shows all allocations since service start
- Good for finding what allocates frequently

### 6. Profile live service – Goroutines
`go tool pprof -http=:9090 http://localhost:6060/debug/pprof/goroutine`

- Shows all goroutines with stack traces
- Good for detecting goroutine leaks

### 7. Profile live service – Mutex contention
`go tool pprof -http=:9090 http://localhost:6060/debug/pprof/mutex`

- Shows mutex lock contention points
- Helps find locking bottlenecks

### 8. Profile live service – Blocking operations
`go tool pprof -http=:9090 http://localhost:6060/debug/pprof/block`

- Shows blocked goroutines (channels, mutexes, I/O)
- Helps find deadlocks and delays

### 9. Trace – detailed execution timeline
`curl http://localhost:6060/debug/pprof/trace?seconds=10 -o trace.out`

`go tool trace -http=:9191 trace.out`

- Captures detailed execution trace
- Shows goroutine scheduling, GC events, blocking

### 10. Compare two profiles (before/after)
`go tool pprof -http=:9090 -base=mem_before.profile mem_after.profile`

- Shows memory difference between two snapshots
- Great for detecting memory leaks over time

---

## Useful links:
* [Profiling and Optimizing Go Programs](https://habr.com/ru/company/badoo/blog/301990/)
* [Profiling and Optimizing Go Web Applications](https://habr.com/ru/company/badoo/blog/324682/)
* [PPROF Github](https://github.com/google/pprof)
* [An Introduction to Go Tool Trace](https://about.sourcegraph.com/go/an-introduction-to-go-tool-trace-rhys-hiltner)