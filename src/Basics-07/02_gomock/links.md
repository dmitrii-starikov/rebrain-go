# Install

```bash
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen@latest
```

```
go@linux:~/GolandProjects/rebrain-go/src/Basics-07/01_task$ ll $GOPATH/bin
total 10364
drwxrwxrwx 2 go go     4096 мар 15 12:43 ./
drwxr-xr-x 9 go go     4096 мар  5 09:03 ../
-rwxrwxr-x 1 go go 10600033 мар 15 12:43 mockgen*     <----


go@linux:~/GolandProjects/rebrain-go/src/Basics-07/01_task$ mockgen -version
v1.6.0
```

```bash
mockgen -source=internal/app/processors/counter/post_client.go -destination=test/gomock/mocks/postmock/post_client_mock.go -package=postmock
```