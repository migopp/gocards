package state

import "github.com/migopp/gocards/internal/types"

func Init() {
	// Lazy init
	var ld types.LRepDeck
	GlobalState = State{
		LoadedDeck: ld,
	}
}
