PORT := 9090

run:
	@go run ./cmd/app

build:
	@go build ./cmd/app

profile:
	@echo "Starting service..."
	@go run ./cmd/app/main.go &
	@echo $$! > .server.pid
	@sleep 2
	@echo "Generating load..."
	@timeout 35 bash -c 'while true; do curl -s http://localhost:$(PORT)/hello > /dev/null; done' &
	@echo "Collecting CPU profile..."
	@curl -s -o cpu_before.pprof "http://localhost:$(PORT)/debug/pprof/profile?seconds=30"
	@echo "Collecting Memory profile..."
	@curl -s -o mem_before.pprof "http://localhost:$(PORT)/debug/pprof/heap"
	@echo "Stopping..."
	@kill $$(cat .server.pid) 2>/dev/null || true
	@rm .server.pid
	@echo "Done! Profiles saved:"
	@ls -la *.pprof

analyze-cpu:
	go tool pprof -http=:8081 cpu_before.pprof

analyze-mem:
	go tool pprof -http=:8082 mem_before.pprof

top-cpu:
	go tool pprof -top cpu_before.pprof | head -20

top-mem:
	go tool pprof -top mem_before.pprof | head -20

clean:
	rm -f *.pprof .server.pid