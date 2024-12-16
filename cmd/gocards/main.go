package main

import (
	"github.com/migopp/gocards/internal/debug"
	"github.com/migopp/gocards/internal/server"
)

// Spawns server loop
func main() {
	// Hardcode for `localhost:8080` for now
	//
	// We will likely set this up programmatically later
	ip := "localhost"
	port := 8080

	// Create the server
	debug.Printf("| Initializing server loop\n|\tIP: \"%s\"\n|\tPort: %d\n", ip, port)
	s := server.Server{
		IP:   ip,
		Port: uint16(port),
	}
	if err := s.Run(); err != nil {
		// Maybe do something...
	}
}
