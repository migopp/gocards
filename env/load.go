package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Idiomatic representation of environment variables
type V struct {
	ServerAddress string
	DBConn        string
	JWTSecret     string
}

// Debug environment variables
func (v *V) Debug() {
	fmt.Printf("V = %+v\n", v)
}

// Load environment variables into `V` struct instance
func Load() V {
	// Load raw environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("`env.Load()`: %v", err)
	}

	// Translate and -> instance
	return V{
		ServerAddress: fmt.Sprintf("%s:%s", os.Getenv("S_IP"), os.Getenv("S_PORT")),
		DBConn: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DBNAME"),
			os.Getenv("DB_PORT"),
		),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
