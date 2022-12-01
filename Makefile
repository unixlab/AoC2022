all: clean fmt test lint build
clean:
	rm -f AoC2022
fmt:
	find . -type f -name '*.go' | xargs -L 1 go fmt
test:
	go test -cover ./...
lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.50.1 golangci-lint run ./... -v -E revive -E govet -E misspell -E goimports -E gofmt -D errcheck --exclude-use-default=false
build:
	CGO_ENABLED=0 \
	go build \
	-o AoC2022
