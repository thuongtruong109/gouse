install:
	go get -v ./...
	go mod download
	go install golang.org/x/tools/cmd/goimports@latest
	go get golang.org/x/perf/cmd/benchstat
	go install golang.org/x/perf/cmd/benchstat
	go install honnef.co/go/tools/cmd/staticcheck@latest

dev:
	go run samples/main/index.go -isDev=true

build:
	go build ./...

doc:
	@echo "Generating docs..."
	go run cmd/doc_gen.go ./samples
	@echo "Done!"

test:
	@echo "Running tests..."
	go clean -testcache
	go test -v -count=1 -cover -coverprofile=coverage.out ./*.go
	go tool cover -func=coverage.out
	@echo "Done!"

BENCH_CMD = go test -count 5 -run=^\# -bench=. ./number/... ./regex/...

bench:
	@echo "Running benchmarks..."
	$(BENCH_CMD)
	@echo "Done!"

bench_filter:
	BENCH_CMD -benchmem > $(f)

bench_stat:
	@echo "Running stat filter benchmarks..."
	make bench_filter f=old.txt
	make bench_filter f=new.txt
	benchstat old.txt new.txt
	@echo "Done!"

format:
	@echo "Running format..."
	gofmt -w -s . && goimports -w . && go fmt ./...
	@echo "Done!"

lint:
	@echo "Running lint..."
	export PATH=$PATH:$(go env GOPATH)/bin
	staticcheck ./...
	@echo "Done!"

count:
	@echo "Counting lines..."
	bash count.sh public/count.svg ./*.go
	@echo "Done!"

pre:
	go mod tidy
	make format && make lint && make test && make build && make doc && make count
	git add .

clean:
	go clean -i -x -cache -testcache -modcache
	rm -rf coverage.out
	rm -rf ./tmp