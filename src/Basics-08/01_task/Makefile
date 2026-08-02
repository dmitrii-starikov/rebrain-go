test: test_01 test_02 bench_04

test_01:
	@cd ./internal/convertor && go test . -v

test_02:
	@cd ./internal/analys && go test . -v

bench_04:
	@cd ./internal/config && go test -bench . -benchmem -benchtime=3s

run:
	@go run ./cmd/app/