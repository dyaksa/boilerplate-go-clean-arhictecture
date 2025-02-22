run:
	go run ./cmd/main.go

build:
	go build -o ./bin/main ./cmd/main.go

clean:
	go mod tidy

.PHONY: run