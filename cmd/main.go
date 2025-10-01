package main

import (
	"br-lesson-4/internal"
	"br-lesson-4/internal/repository/db"
	"br-lesson-4/internal/repository/inmemory"
	"br-lesson-4/internal/server"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func main() {
	var storage server.Storage
	config := internal.ReadConfig()

	storage, err := db.NewStorage(config.DNS)
	if err != nil {
		log.Println("Failed to connect to database")
		storage = inmemory.NewInMemoryStorage()
	}

	if err := db.Migrations(config.DNS, config.MigratePath); err != nil {
		log.Fatal(err)
	}

	srv := server.NewToDoServer(config, storage)

	if err := srv.Start(); err != nil {
		panic(err)
	}
}
