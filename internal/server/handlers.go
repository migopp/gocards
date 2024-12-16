package server

import (
	"net/http"

	"github.com/migopp/gocards/internal/debug"
)

func home(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Serving home page")
}
