```bash
go run increment.go 
Result 1: Counter{A: 101, B: 102}
Result 2: Counter{A: 150, B: 250}
Result 3: Counter{A: 1123, B: 1321}
Result 4: Counter{A: 2000, B: 9888}
```

# Benchmarks and conclusions

Compare tests
```go
go test -bench=. -benchmem -benchtime=100x
```

| Benchmark                               | Iterations | ns/op        | B/op    | allocs/op   |
|-----------------------------------------|------------|--------------|---------|-------------|
| BenchmarkWithAllocationButWithoutPool-8 | 100        | 1020 ns/op   | 0 B/op  | 0 allocs/op |
| BenchmarkWithPool_SingleObject-8        | 100        | 892.9 ns/op  | 11 B/op | 0 allocs/op |
| BenchmarkWithPool-8                     | 100        | 469772 ns/op | 11 B/op | 0 allocs/op |
| BenchmarkWithoutPool-8                  | 100        | 439649 ns/op | 0 B/op  | 0 allocs/op |


```go
func BenchmarkWithAllocationButWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// new object every iteration
		c := &Counter{A: i, B: i * 2}

		for j := 0; j < 1000; j++ {
			incrementWithoutPool(c)
		}
	}
}
```

- one allocation in main loop
- then mutate one object (by pointer)
- 0 allocs/op - object is created once per iteration
- 1020 ns/op - very fast, just mutation

```go
func BenchmarkWithPool_SingleObject(b *testing.B) {
    for i := 0; i < b.N; i++ {
        c := counterPool.Get().(*Counter) // once get from Pool
        c.A = i
        c.B = i * 2
        
        for j := 0; j < 1000; j++ {
            c.A++   // mutate object
            c.B++
        }
        
        counterPool.Put(c) // put once to Pool
    }
}
```

- the fastest - 892.9 ns/op
- 11 B/op - sync.Pool overheads

```go
func BenchmarkWithPool(b *testing.B) {
	counterPool.Put(&Counter{A: 1, B: 2})

	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			b.StopTimer()
			incrementWithPool() // get+increment+put each iteration!
			b.StartTimer()
		}
	}
}
```

- the slowest - 469772 ns/op. Because of get/increment/put each iteration
- 11 B/op - sync.Pool overheads

```go
func BenchmarkWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			b.StopTimer()
			incrementWithoutPool(&Counter{A: 1, B: 2})
			b.StartTimer()
		}
	}
}
```

- 439649 ns/op

---

- Pool is effective when objects are long-lived: Get() once → many operations → Put() once
- sync.Pool makes sense for heavy objects (more than 11 bytes)