package load

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/migopp/gocards/internal/types"
)

// Parse a file into some internal deck representation
func ToDeck(f multipart.File, h *multipart.FileHeader) (types.LRepDeck, error) {
	var ld types.LRepDeck
	var err error

	// Cut and store the file name
	base := filepath.Base(h.Filename)
	ext := filepath.Ext(base)
	ld.Deck.DeckName = strings.TrimSuffix(base, ext)

	// Parse the yaml file -> irep
	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&ld); err != nil {
		return ld, fmt.Errorf("Cannot decode file: %s", h.Filename)
	}
	return ld, nil
}
