## Benchmarks

```go
make bench_04 
-------------
goos: linux
goarch: amd64
pkg: module07/internal/config
cpu: AMD Ryzen 3 3100 4-Core Processor              
BenchmarkReflectStructToMap-8            4034947       905.6 ns/op     416 B/op   3 allocs/op
BenchmarkGeneratedStructToMap-8         11807301       294.0 ns/op     64 B/op    4 allocs/op
PASS
ok      module07/internal/config        8.341s
```

## Task Conclusions

### Performance Comparison: Reflection vs Code Generation

| Method                                   | Time (ns/op) | Memory (B/op) | Allocations |
|------------------------------------------|--------------|---------------|-------------|
| Reflection (`convertor.StructToMap`)     | 905.6        | 416           | 3           |
| Code Generation (`Config.StructToMap()`) | 294.0        | 64            | 4           |

### Key Conclusions

1.  Performance
     - Code generation is ~3x faster than reflection
     - Generated code avoids runtime type inspection and field iteration
2.  Memory Efficiency
     - Code generation uses ~6.5x less memory (64 vs 416 bytes)
     - Fewer allocations during execution
3.  Trade-offs
     - Reflection: Flexible, works with any struct, but slower and more memory-heavy
     - Code Generation: Faster and more efficient, but requires known structs at build time
4.  Use Cases
     - Reflection for generic tools, unknown types, or flexible APIs
     - Code Generation for high-performance code, frequent calls, and known data structures

---

### Why Code Generation Wins

| Reflection                               | Code Generation               |
|------------------------------------------|-------------------------------|
| `reflect.ValueOf()` on each call         | Direct field access           |
| Type checking at runtime                 | Type checking at compile time |
| Dynamic field iteration                  | Hardcoded field mapping       |
| Memory allocation for reflection objects | Minimal heap allocations      |

---

### Summary

Code generation moves complexity from runtime to compile time. Reflection provides flexibility
at the cost of performance.  Code generation provides performance at the cost of flexibility.