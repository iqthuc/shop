.PHONY: gendb run

gendb:
	sqlc generate -f internal/infrastructure/database/store/sqlc/sqlc.yaml
run:
	go run main.go
lint:
	golangci-lint run
