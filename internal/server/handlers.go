package server

import (
	"fmt"
	"net/http"

	"github.com/migopp/gocards/internal/debug"
	"github.com/migopp/gocards/internal/load"
	"github.com/migopp/gocards/internal/state"
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
	http.HandleFunc("POST /cards/submit", cardsSubmit)
	http.HandleFunc("POST /decks/upload", decksUpload)
}

// Trigger to serve the home page
func home(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Serving home.html\n")

	// Serve `home.html`
	serveTmpl(w, "home.html", nil)
}

// Trigger to serve the cards page
func cards(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Serving cards.html\n")

	// Prep page content
	dynContent := &DynContent{
		Word: state.GlobalState.GetFront(),
	}

	// Serve `cards.html`
	serveTmpl(w, "cards.html", dynContent)
}

// Trigger to submit an answer
func cardsSubmit(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Hit `/cards/submit`\n")

	// Check if the answer is correct
	r.ParseForm()
	input := r.FormValue("ans")
	if input == state.GlobalState.GetBack() {
		debug.Printf("| Correct answer given\n")

		// TODO:
	} else {
		debug.Printf("| Wrong answer given\n")

		// TODO:
	}
}

// Trigger to upload a deck
func decksUpload(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Hit `/decks/upload`\n")

	// Parse the uploaded file
	const FileSize = 10 << 20 // ~10MB
	err := r.ParseMultipartForm(FileSize)
	if err != nil {
		http.Error(w, "ERROR PARSING FILE FORM", http.StatusBadRequest)
		return
	}

	// Open file
	//
	// r.FormFile(...) -> file, header, err
	file, header, err := r.FormFile("deck-name")
	if err != nil {
		errStr := fmt.Sprintf("ERROR RETRIEVING FORM DATA [%v]", err)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Parse file in usable form
	ld, err := load.ToDeck(file, header)
	if err != nil {
		errStr := fmt.Sprintf("ERROR RESOLIVING FILE TO DECK [%v]", err)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}
	debug.PrintLoadedDeck(ld)

	// Update the global state
	state.GlobalState.UpdateDeck(ld)
}
