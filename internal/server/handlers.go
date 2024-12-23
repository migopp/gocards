package server

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

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
	http.HandleFunc("POST /decks/select", decksSelect)
	http.HandleFunc("POST /decks/upload", decksUpload)
}

// Trigger to serve the home page
func home(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Serving home.html\n")

	// Prep page content
	dynContent := HomeDynContent{
		Decks: state.GlobalState.UploadedDecks,
	}

	// Serve `home.html`
	serveTmpl(w, "home.html", dynContent)
}

// Trigger to serve the cards page
func cards(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Serving cards.html\n")

	// Prep page content
	state.GlobalState.Reset()
	front, err := state.GlobalState.GetFront()
	if err != nil {
		errStr := fmt.Sprintf("ERROR GRABBING FRONT OF CURRENT CARD [%v]", err)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}
	dynContent := GameDynContent{
		Word:  front,
		Ratio: state.GlobalState.Ratio(),
	}

	// Serve `cards.html`
	serveTmpl(w, "cards.html", dynContent)
}

// Trigger to submit an answer
func cardsSubmit(w http.ResponseWriter, r *http.Request) {
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
		debug.Printf("| Correct answer given\n")

		// Mark correct
		state.GlobalState.AddRight()

		// Prep page content
		exists := state.GlobalState.NextCard()
		if exists == false {
			dynContent := GameDynContent{
				Ratio: state.GlobalState.Ratio(),
			}
			tmpl, err := template.New("homeButton").Parse(homeButton)
			if err != nil {
				debug.Printf("x Could not parse `homeButton`\n")
				errStr := fmt.Sprintf("ERROR PARISNG `hometButton` [%v]", err)
				http.Error(w, errStr, http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/html")
			tmpl.Execute(w, dynContent)
			return
		}
		front, err := state.GlobalState.GetFront()
		if err != nil {
			debug.Printf("x Could not grab front of current card\n")
			errStr := fmt.Sprintf("ERROR GRABBING FRONT OF CURRENT CARD [%v]", err)
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}
		dynContent := GameDynContent{
			Word:  front,
			Ratio: state.GlobalState.Ratio(),
		}
		tmpl, err := template.New("ui").Parse(rightCardsUI)
		if err != nil {
			debug.Printf("x Could not parse `rightCardsUI`\n")
			errStr := fmt.Sprintf("ERROR PARISNG `rightCardsUI` [%v]", err)
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}

		// Serve updated content
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, dynContent)
	} else {
		// Wrong answer given
		//
		// Similar story to the correct case, but we don't want to
		// advnace the card until the user gets it right
		debug.Printf("| Wrong answer given\n")

		// Mark wrong
		state.GlobalState.AddWrong()

		// Serve same page content
		front, err := state.GlobalState.GetFront()
		if err != nil {
			debug.Printf("x Could not grab front of current card\n")
			errStr := fmt.Sprintf("ERROR GRABBING FRONT OF CURRENT CARD [%v]", err)
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}
		dynContent := GameDynContent{
			Word:  front,
			Ratio: state.GlobalState.Ratio(),
		}
		tmpl, err := template.New("ui").Parse(wrongCardsUI)
		if err != nil {
			debug.Printf("x Could not parse `wrongCardsUI`\n")
			errStr := fmt.Sprintf("ERROR PARISNG `wrongCardsUI` [%v]", err)
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}

		// Serve updated content
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, dynContent)
	}
}

// Trigger to select a deck
func decksSelect(w http.ResponseWriter, r *http.Request) {
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
	tmpl, err := template.New("startButton").Parse(startButton)
	if err != nil {
		debug.Printf("x Could not parse `startButton`\n")
		errStr := fmt.Sprintf("ERROR PARISNG `startButton` [%v]", err)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, nil)
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

	// Upload the deck to the state
	//
	// NOTE: This should be some db-fetching in the future.
	state.GlobalState.UploadDeck(ld)

	// Serve deck selection updates
	dynContent := HomeDynContent{
		Decks: state.GlobalState.UploadedDecks,
	}
	tmpl, err := template.New("deckSelect").Parse(deckSelect)
	if err != nil {
		debug.Printf("x Could not parse `deckSelect`\n")
		errStr := fmt.Sprintf("ERROR PARISNG `deckSelect` [%v]", err)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, dynContent)
}
