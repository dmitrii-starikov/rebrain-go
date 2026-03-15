```text
go@linux:~/GolandProjects/rebrain-go/src/Basics-07/01_task$ go test -cover ./internal/pkg/util
ok      module06/internal/pkg/util      0.006s  coverage: 70.6% of statements
```

```text
go test -coverprofile=profile.out
go tool cover -html=profile.out
go tool cover -func=profile.out
```

## Links
[The cover story](https://blog.golang.org/cover)