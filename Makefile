run:
	go run ./cmd/main.go

build:
	go build -o ./bin/main ./cmd/main.go

clean:
	go mod tidy && rm -rf bin/*

test:
	go test -v ./...

.PHONY: run