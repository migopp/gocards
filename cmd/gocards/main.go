package main

import (
	"log"

	"github.com/migopp/gocards/db"
	"github.com/migopp/gocards/env"
	"github.com/migopp/gocards/server"
)

func main() {
	// Load environment variables
	env.GCV = env.Load()

	// DB Connection
	//
	// Use SQLite for now...
	db.GCDB = db.New(db.SQLite, "gocards.db")
	if err := db.GCDB.Connect(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	if err := db.GCDB.Migrate(); err != nil {
		log.Fatalf("Failed to migrate DB: %v", err)
	}

	// Start server
	server.GCS = server.New(env.GCV.ServerAddress)
	server.GCS.Config()
	if err := server.GCS.Up(); err != nil {
		log.Fatalf("Server shut down: %v", err)
	}
}
