package types

// Idiomatic representations of data
//
// This is what makes most sense to us when we handle
// registration and such.

type IRepUser struct {
	UserName string `yaml:"user_name"`
}

type IRepDeck struct {
	DeckName string `yaml:"deck_name"`
}

type IRepCard struct {
	Front string `yaml:"front"`
	Back  string `yaml:"back"`
}

// Data registrations that make sense at load time
//
// While a `DeckName` field attatched to a deck
// may make sense on its own normally, in terms of
// loading, it makes more sense to associate with
// the cards that are stored inside of it

type LRepDeck struct {
	Deck  IRepDeck   `yaml:"deck"`
	Cards []IRepCard `yaml:"cards"`
}

// DB representations of data
//
// These are raw translations of the schema into struct
// form, so they don't make a ton of sense to work with.
//
// Perhaps they'll have use as an interface, though.

type DBRepUser struct {
	UserID    uint32
	UserName  string
	Timestamp uint64
}

type DBRepDeck struct {
	DeckID    uint32
	UserID    uint32
	DeckName  string
	Timestamp uint64
}

type DBRepCard struct {
	CardID    uint32
	DeckID    uint32
	Front     string
	Back      string
	Timestamp uint64
}
