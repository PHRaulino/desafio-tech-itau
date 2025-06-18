run: sqlc wire
	clear && go run cmd/main.go
wire:
	wire ./...
sqlc:
	sqlc generate
reset_db:
	rm -rf cinetuber.db && sqlite3 cinetuber.db < shared/db/sqlc/schema.sql && sqlite3 cinetuber.db < shared/db/sqlc/seed.sql
lint:
	golangci-lint run
serve:
	go run cmd/main.go