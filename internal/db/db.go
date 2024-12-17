package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/migopp/gocards/internal/debug"
)

var DB *sql.DB

func Init() error {
	debug.Printf("| DB Init\n")

	// Get env variables to spin up DB
	user := os.Getenv("DB_USER")
	passwd := os.Getenv("DB_PASSWD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// Spin up DB
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, passwd, host, port, name)
	debug.Printf("| Connection string: %s\n", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("Error opening database connection: %v", err)
	}
	DB = db

	// Test the connection
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("Could not ping database: %v", err)
	}

	debug.Printf("| DB connection success\n")

	return nil
}
