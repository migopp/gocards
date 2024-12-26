package main

import (
	"log"
	"os"
	"strconv"

	"github.com/migopp/gocards/internal/db"
	"github.com/migopp/gocards/internal/debug"
	"github.com/migopp/gocards/internal/server"
)

// Spawns server loop
func main() {
	// Get env
	host := os.Getenv("SERVER_HOST")
	port_raw := os.Getenv("SERVER_PORT")
	port, err := strconv.ParseUint(port_raw, 10, 16)
	if err != nil {
		log.Fatal("x Server port in environment could not be parsed")
	}
	debug.Printf("| Host @ %s, Port %d\n", host, port)

	// Create the DB
	if err = db.Init(); err != nil {
		log.Fatalf("x Failed to init db: %v", err)
	}

	// Create the server
	debug.Printf("| Initializing server loop\n|\tIP: \"%s\"\n|\tPort: %d\n", host, port)
	s := server.Server{
		IP:   host,
		Port: uint16(port),
	}
	if err = s.Spawn(); err != nil {
		// Maybe do something...
		db.Close()
		log.Fatal("x Server shut down")
	}
}
