.PHONY: play
play:
	go run cmd/main.go

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: test-cov
test-cov:
	go test -race -coverprofile=coverage.out ./...