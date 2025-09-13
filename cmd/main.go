package main

import (
	"br-lesson-4/internal/repository/inmemory"
	"br-lesson-4/internal/server"
)

func main() {
	storage := inmemory.NewInMemoryStorage()
	srv := server.NewToDoServer(storage)

	if err := srv.Start(); err != nil {
		panic(err)
	}
}
