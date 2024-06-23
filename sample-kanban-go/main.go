package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app"
	"log"
)

// main - service entrypoint
func main() {
	server, err := app.NewKanbanServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	// check db
	err = server.CheckDb()
	if err != nil {
		log.Fatal(err)
	}
	// spin up server
	log.Fatal(server.Listen())
}
