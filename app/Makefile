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

# Targets de build por sistema
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/cinetuber-mac cmd/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/cinetuber-linux cmd/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/cinetuber.exe cmd/main.go

# Build para todos
build-all: build-mac build-linux build-windows