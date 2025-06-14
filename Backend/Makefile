run: sqlc wire
	clear && go run cmd/main.go
wire:
	wire ./...
sqlc:
	sqlc generate
reset_db:
	rm -rf meu_banco.db && sqlite3 meu_banco.db < adapters/db/sqlc/schema.sql && sqlite3 meu_banco.db < adapters/db/sqlc/seed.sql
lint:
	golangci-lint run