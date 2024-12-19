package server

import (
	"net/http"

	"github.com/migopp/gocards/internal/debug"
)

// Configure the router
func initHandlers() {
	// For now, we just use the default `mux` in stdlib
	//
	// We could use something more advanced, but this
	// will do for now, since our API is not anything
	// revolutionary
	//
	// The enhancements from 1.22 are more than enough:
	// https://go.dev/blog/routing-enhancements
	http.HandleFunc("GET /", home)
	http.HandleFunc("GET /cards", cards)
}

// Trigger to serve the home page
func home(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Serving home.html\n")

	// Serve `home.html`
	serveTmpl(w, "home.html")
}

// Trigger to serve the cards page
func cards(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Serving cards.html\n")

	// Serve `cards.html`
	serveTmpl(w, "cards.html")
}
