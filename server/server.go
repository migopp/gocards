package server

import (
	"github.com/gin-gonic/gin"
	"github.com/migopp/gocards/db"
)

type Server struct {
	address    string
	engine     *gin.Engine
	deckStates map[uint]*deckState
}

func New(a string) Server {
	return Server{
		address:    a,
		engine:     gin.Default(),
		deckStates: make(map[uint]*deckState),
	}
}

func (s *Server) Config() {
	// Proxies
	s.engine.SetTrustedProxies(nil)

	// Templates and static assets
	s.engine.LoadHTMLGlob("./web/templates/*")
	s.engine.Static("/static", "./web/static")

	// Routes
	s.engine.GET("/", getIndex)
	s.engine.GET("/signup", getSignup)
	s.engine.POST("/signup", postSignup)
	s.engine.GET("/login", getLogin)
	s.engine.POST("/login", postLogin)
	s.engine.GET("/cards", getCards)
	s.engine.POST("/cards", postCards)
	s.engine.GET("/decks", getDecks)
	s.engine.POST("/decks", postDecks)
	s.engine.POST("/decks/select", postDecksSelect)
}

func (s *Server) Up() error {
	return s.engine.Run(s.address)
}

func (s *Server) deckStateForUserID(id uint) (*deckState, bool) {
	ds, ok := s.deckStates[id]
	return ds, ok
}

func (s *Server) deckStateForUser(u db.User) (*deckState, bool) {
	ds, ok := s.deckStates[u.ID]
	return ds, ok
}
