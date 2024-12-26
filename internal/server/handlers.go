package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/migopp/gocards/internal/debug"
	"github.com/migopp/gocards/internal/load"
	"github.com/migopp/gocards/internal/state"
	"github.com/migopp/gocards/internal/templates"
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
	http.HandleFunc("GET /", getHome)
	http.HandleFunc("GET /cards", getCards)
	http.HandleFunc("POST /cards/submit", postCardsSubmit)
	http.HandleFunc("POST /decks/select", postDecksSelect)
	http.HandleFunc("POST /decks/upload", postDecksUpload)
}

// Trigger to serve the home page
func getHome(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Serving home.html\n")
	dc := templates.HomeContent{
		Decks: state.GlobalState.UploadedDecks,
	}
	templates.ServeTemplate(w, templates.Home, dc)
}

// Trigger to serve the cards page
func getCards(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Serving cards.html\n")
	state.GlobalState.Reset()
	front, err := state.GlobalState.GetFront()
	if err != nil {
		errStr := fmt.Sprintf("ERROR GRABBING FRONT OF CURRENT CARD [%v]", err)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}
	dc := templates.GameContent{
		Word:  front,
		Ratio: state.GlobalState.Ratio(),
	}
	templates.ServeTemplate(w, templates.Cards, dc)
}

// Trigger to submit an answer
func postCardsSubmit(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Hit `/cards/submit`\n")

	// Check if the answer is correct
	err := r.ParseForm()
	if err != nil {
		errStr := fmt.Sprintf("ERROR PARSING RAW ANSWER DATA [%v]", err)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}
	input := r.FormValue("ans")
	back, err := state.GlobalState.GetBack()
	if err != nil {
		errStr := fmt.Sprintf("ERROR GRABBING BACK OF CURRENT CARD [%v]", err)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}
	if input == back {
		// Correct answer given
		//
		// Resend the `ui` DOM subtree, but update the dynamic content
		// to the next card contents
		state.GlobalState.AddRight()

		// Prep page content
		exists := state.GlobalState.NextCard()
		if exists == false {
			// We are at the end of the deck
			//
			// Send some report and allow the user to return home
			dc := templates.GameContent{
				Ratio: state.GlobalState.Ratio(),
			}
			templates.ServeTemplate(w, templates.End, dc)
			return
		}
		front, err := state.GlobalState.GetFront()
		if err != nil {
			debug.Printf("x Could not grab front of current card\n")
			errStr := fmt.Sprintf("ERROR GRABBING FRONT OF CURRENT CARD [%v]", err)
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}
		dc := templates.GameContent{
			Word:  front,
			Ratio: state.GlobalState.Ratio(),
		}
		templates.ServeTemplate(w, templates.CardRight, dc)
	} else {
		// Wrong answer given
		//
		// Similar story to the correct case, but we don't want to
		// advnace the card until the user gets it right
		state.GlobalState.AddWrong()

		// Serve same page content
		front, err := state.GlobalState.GetFront()
		if err != nil {
			debug.Printf("x Could not grab front of current card\n")
			errStr := fmt.Sprintf("ERROR GRABBING FRONT OF CURRENT CARD [%v]", err)
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}
		dc := templates.GameContent{
			Word:  front,
			Ratio: state.GlobalState.Ratio(),
		}
		templates.ServeTemplate(w, templates.CardWrong, dc)
	}
}

// Trigger to select a deck
func postDecksSelect(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Hit `/decks/select`\n")

	// Parse the form and get the selected value
	//
	// We've set this up to be the index of the deck that we want
	// to serve in the future
	err := r.ParseForm()
	if err != nil {
		errStr := fmt.Sprintf("ERROR PARSING DECK SELECTION FORM [%v]", err)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}
	sdn := r.FormValue("decks")
	if sdn == "" {
		http.Error(w, "PLEASE SELECT A DECK...", http.StatusBadRequest)
		return
	}
	sdi, err := strconv.Atoi(sdn)
	if err != nil {
		errStr := fmt.Sprintf("INVALID DECK ID: %s [%v]", sdn, err)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	// Actually index and get the deck
	deck := state.GlobalState.UploadedDecks[sdi]

	// Update the global state
	state.GlobalState.UpdateDeck(deck)

	// Serve the start button
	templates.ServeTemplate(w, templates.Start, nil)
}

// Trigger to upload a deck
func postDecksUpload(w http.ResponseWriter, r *http.Request) {
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

	// Upload the deck to the state
	//
	// NOTE: This should be some db-fetching in the future.
	state.GlobalState.UploadDeck(ld)

	// Serve deck selection updates
	dc := templates.HomeContent{
		Decks: state.GlobalState.UploadedDecks,
	}
	templates.ServeTemplate(w, templates.DeckSelect, dc)
}
