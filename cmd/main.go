package main

import (
	"br-lesson-4/internal"
	"br-lesson-4/internal/repository/inmemory"
	"br-lesson-4/internal/server"
)

func main() {
	config := internal.ReadConfig()
	storage := inmemory.NewInMemoryStorage()
	srv := server.NewToDoServer(config, storage)

	if err := srv.Start(); err != nil {
		panic(err)
	}
}
