start:
	go run ./cmd/main.go
test:
	go test -v -cover -count=1 ./...