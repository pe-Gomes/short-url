run:
	@go run main.go

create-migration:
	@migrate create -ext sql -dir infra/db/migration -seq $(d)

migrate-up:
	@migrate -path infra/db/migration -database "postgresql://admin:changeme@localhost:5432/short_url?sslmode=disable" -verbose up

migrate-down:
	@migrate -path infra/db/migration -database "postgresql://admin:changeme@localhost:5432/short_url?sslmode=disable" -verbose down
