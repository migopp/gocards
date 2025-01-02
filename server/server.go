package server

import "github.com/gin-gonic/gin"

type Server struct {
	address string
	engine  *gin.Engine
}

func New(a string) Server {
	return Server{
		address: a,
		engine:  gin.Default(),
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
}

func (s *Server) Up() error {
	return s.engine.Run(s.address)
}
