.PHONY: gendb

gendb:
	sqlc generate -f internal/infrastructure/database/store/sqlc/sqlc.yaml
