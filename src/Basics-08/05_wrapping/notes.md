## Wrapper

 - [gowrap](https://github.com/hexdigest/gowrap)
 - [Generating code](https://blog.golang.org/generate)

```go
gowrap gen -p worker -i Worker -t log -o ./worker_with_log.go
```

```go
package generated

//go:generate gowrap gen -p worker -i Worker -t log -o ./worker_with_log.go
```