## Links:
- https://medium.com/@yardenlaif/go-sync-or-go-home-errgroup-f91a0ee72d3f

---

```bash
go@linux: $ go get golang.org/x/sync@latest
go@linux: $ go mod download
go: downloading golang.org/x/sync v0.19.0
go: golang.org/x/sync@v0.19.0 requires go >= 1.24.0; switching to go1.25.7
go: downloading go1.25.7 (linux/amd64)
go: upgraded go 1.22.2 => 1.24.0
go: added golang.org/x/sync v0.19.0
go: downloading go1.24.0 (linux/amd64)
```

```bash
go@linux:~/GolandProjects/rebrain-go/src/Basics-06/08_errgroup$ go run errgroup.go 
runTasks: error in task 1
ok
```