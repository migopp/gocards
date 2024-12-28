package main

import (
	"log"
	"os"

	"github.com/migopp/gocards/internal/db"
	"github.com/migopp/gocards/internal/debug"
	"github.com/migopp/gocards/internal/server"
)

// Spawns server loop
func main() {
	// Open the DB
	if err := db.Init(); err != nil {
		log.Fatalf("Failed to init db: %v", err)
	}
	defer db.Close()

	// Get env variables
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	debug.Printf("Host @ %s, Port %s\n", host, port)

	// Create the server
	debug.Printf("Initializing server loop\n\tIP: \"%s\"\n\tPort: %s\n", host, port)
	s := server.Server{
		IP:   host,
		Port: port,
	}
	if err := s.Spawn(); err != nil {
		log.Fatal("Server shut down")
	}
}
