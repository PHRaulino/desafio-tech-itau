package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

//go:embed sqlc/schema.sql
//go:embed sqlc/seed.sql
var sqlFS embed.FS

func NewSQLiteConnection(path string) *sql.DB {
	_, err := os.Stat(path)
	dbExists := err == nil

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("Erro ao abrir o banco: %v", err)
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		log.Fatal("Erro ao habilitar foreign_keys:", err)
	}

	if !dbExists {
		fmt.Println("Banco não encontrado, criando schema e seed...")
		if err := executeSQLFile(db, "sqlc/schema.sql"); err != nil {
			log.Fatalf("Erro ao executar schema.sql: %v", err)
		}
		if err := executeSQLFile(db, "sqlc/seed.sql"); err != nil {
			log.Fatalf("Erro ao executar seed.sql: %v", err)
		}
	}

	return db
}

func executeSQLFile(db *sql.DB, path string) error {
	content, err := sqlFS.ReadFile(path)
	if err != nil {
		return fmt.Errorf("erro ao ler %s: %w", path, err)
	}
	if _, err := db.Exec(string(content)); err != nil {
		return fmt.Errorf("erro ao executar %s: %w", path, err)
	}
	return nil
}