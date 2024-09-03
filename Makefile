.PHONY: play
play:
	go run cmd/main.go

.PHONY: test
test:
	go test -cover ./...

.PHONY: test-cov
test-cov:
	go test -coverprofile=coverage.out ./...