package server

import (
	"fmt"
	"net/http"
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
	front, err := state.GlobalState.GetFront()
	if err != nil {
		errStr := fmt.Sprintf("ERROR GRABBING FRONT OF CURRENT CARD [%v]", err)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}
	dynContent := &DynContent{
		Word: front,
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

		// Prep page content
		exists := state.GlobalState.NextCard()
		if exists == false {
			tmpl, err := template.New("homeButton").Parse(homeButton)
			if err != nil {
				debug.Printf("x Could not parse `homeButton`\n")
				errStr := fmt.Sprintf("ERROR PARISNG `hometButton` [%v]", err)
				http.Error(w, errStr, http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/html")
			tmpl.Execute(w, nil)
			return
		}
		front, err := state.GlobalState.GetFront()
		if err != nil {
			debug.Printf("x Could not grab front of current card\n")
			errStr := fmt.Sprintf("ERROR GRABBING FRONT OF CURRENT CARD [%v]", err)
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}
		dynContent := &DynContent{
			Word: front,
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

		// Serve same page content
		front, err := state.GlobalState.GetFront()
		if err != nil {
			debug.Printf("x Could not grab front of current card\n")
			errStr := fmt.Sprintf("ERROR GRABBING FRONT OF CURRENT CARD [%v]", err)
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}
		dynContent := &DynContent{
			Word: front,
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
