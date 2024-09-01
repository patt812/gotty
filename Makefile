.PHONY: play
play:
	go run cmd/main.go

.PHONY: test
test:
	go test -race -cover ./...