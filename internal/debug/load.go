package debug

import (
	"github.com/migopp/gocards/internal/types"
)

// Print the loaded representation of a deck
func PrintLoadedDeck(ld types.LRepDeck) {
	// Base
	Printf(
		"Printing loaded deck:\n"+
			"\t`DeckName`: %s\n"+
			"\t`Cards`:\n",
		ld.Deck.DeckName,
	)

	// Each card
	for _, card := range ld.Cards {
		Printf("\t\t`front`: \"%s\", `back`: \"%s\"\n", card.Front, card.Back)
	}
}
