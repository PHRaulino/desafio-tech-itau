package main

import (
	apiEntrypoint "github.com/phraulino/cinetuber/interface/api"
	database "github.com/phraulino/cinetuber/shared/db"
)

func main() {
	db := database.NewSQLiteConnection("cinetuber.db")
	defer db.Close()
	apiEntrypoint.New(db)
}
