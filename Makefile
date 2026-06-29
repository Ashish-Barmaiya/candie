.PHONY: build run test fmt lint clean

build:
	go build -o bin/candie ./cmd/candie

run:
	go run ./cmd/candie

test:
	go test ./...

fmt:
	gofmt -w .

lint:
	golangci-lint run

clean:
	rm -rf bin
