package server

import (
	"net/http"

	"github.com/migopp/gocards/internal/debug"
)

func home(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Serving home.html\n")

	// Serve `home.html`
	servePage(w, "home.html")
}

func cards(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Serving cards.html\n")

	// Serve `cards.html`
	servePage(w, "cards.html")
}
