package templates

import (
	"github.com/migopp/gocards/internal/types"
)

type GameContent struct {
	Word  string
	Ratio float64
}

type HomeContent struct {
	Decks []types.LRepDeck
}
