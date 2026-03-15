- [links.md](../02_gomock/links.md)

```bash
mockgen -source=internal/app/services/post/client.go -destination=test/gomock/mocks/postmock/post_client_mock.go -package=postmock
```
**Result:** [post_client_mock.go](../01_task/test/gomock/mocks/postmock/post_client_mock.go)

---

```bash
go test ./...
?       module06/cmd/app        [no test files]
?       module06/internal/app/handlers/hello    [no test files]
ok      module06/internal/app/processors/counter        0.004s
?       module06/internal/app/services/post     [no test files]
ok      module06/internal/pkg/util      0.004s
?       module06/test/gomock/mocks/postmock     [no test files]
```

or

```bash
go test -v ./internal/app/processors/counter -run TestPostCount
```

Standard output in `t.Parallel()` mode

```text
=== RUN   TestPostCountTable
=== PAUSE TestPostCountTable
=== CONT  TestPostCountTable
=== RUN   TestPostCountTable/success_with_multiple_posts
=== PAUSE TestPostCountTable/success_with_multiple_posts
=== RUN   TestPostCountTable/success_with_empty_posts
=== PAUSE TestPostCountTable/success_with_empty_posts
=== RUN   TestPostCountTable/error_from_client
=== PAUSE TestPostCountTable/error_from_client
=== CONT  TestPostCountTable/success_with_multiple_posts
=== CONT  TestPostCountTable/error_from_client
=== CONT  TestPostCountTable/success_with_empty_posts
--- PASS: TestPostCountTable (0.00s)
    --- PASS: TestPostCountTable/success_with_multiple_posts (0.00s)
    --- PASS: TestPostCountTable/error_from_client (0.00s)
    --- PASS: TestPostCountTable/success_with_empty_posts (0.00s)
PASS
ok      module06/internal/app/processors/counter        0.004s
```