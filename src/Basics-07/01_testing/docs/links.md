```bash
go@linux:~/GolandProjects/rebrain-go/src/Basics-07/01_testing$ go test ./...
ok      01testing/cmd/01_testing        0.002s
ok      01testing/pkg/summy     0.002s
go@linux:~/GolandProjects/rebrain-go/src/Basics-07/01_testing$ go test -v ./pkg/summy
=== RUN   TestSum
--- PASS: TestSum (0.00s)
PASS
ok      01testing/pkg/summy     0.003s
go@linux:~/GolandProjects/rebrain-go/src/Basics-07/01_testing$ 
```

`testify`
```bash
go@linux:~/GolandProjects/rebrain-go/src/Basics-07/01_testing$ go test ./...
--- FAIL: TestGetOrder (0.00s)
    order_test.go:12: 
                Error Trace:    /home/go/GolandProjects/rebrain-go/src/Basics-07/01_testing/cmd/01_testing/order_test.go:12
                Error:          "[0xc0000129c8 0xc0000129d0 0xc0000129d8]" should have 5 item(s), but has 3
                Test:           TestGetOrder
FAIL
FAIL    01testing/cmd/01_testing        0.003s
ok      01testing/pkg/summy     (cached)
FAIL
go@linux:~/GolandProjects/rebrain-go/src/Basics-07/01_testing$ 
```

```go
go test [build/test flags] [packages] [build/test flags & test binary flags]
```

It's important to note that go test can be run in different modes:

**Local directory mode**. In this mode, tests are run only in the current directory, and 
test results are not cached. To run tests in this mode, simply omit the package arguments 
(i.e., don't specify which packages to run the tests from).
Example of running in local directory mode: `go test -v`

**Package list mode**. In this mode, packages are explicitly specified, and positive test 
results are cached (to avoid rerunning positive tests, as there can be many packages, and 
each test run without caching would take a long time). It's important to note that results 
will always be cached unless flags are specified that are not cacheable test flags: cpu, 
list, parallel, run, short, and v.
To run a specific package: go test mypackage. Run tests for all packages in the current 
directory: `go test ./...`

Run tests in package list mode, without caching: `go test ./... -count=1`.

## Links

* [Golang book testing](http://golang-book.ru/chapter-12-testing.html)
* [Godoc package testing](https://golang.org/pkg/testing/)
* [Testify package](https://github.com/stretchr/testify)
* [Команды go: go test, тестировать пакеты](https://golang-blog.blogspot.com/2019/06/go-commands-go-test.html)