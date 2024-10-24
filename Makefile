build: test
	go build -o build/numbers ./cmd/main.go

test:
	go test -v ./pkg/...

run:
	go run ./cmd/main.go
