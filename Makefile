include dev.env
export $(shell sed 's/=.*//' dev.env)

.PHONY: sqlc lint\
		migrate-create  migrate-up migrate-up-1 migrate-down migrate-down-1

sqlc:
	sqlc generate -f internal/infrastructure/database/store/sqlc/sqlc.yaml

migrate-up:
	migrate -path migrations/postgres -database $(DB_SOURCE) -verbose up

migrate-up-1:
	migrate -path migrations/postgres -database $(DB_SOURCE) -verbose up 1


migrate-down:
	migrate -path migrations/postgres -database $(DB_SOURCE) -verbose down

migrate-down-1:
	migrate -path migrations/postgres -database $(DB_SOURCE) -verbose down 1

migrate-create:
	migrate create -ext sql -dir migrations/postgres -seq $(name)

lint:
	golangci-lint run
