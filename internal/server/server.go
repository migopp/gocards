package server

import (
	"fmt"
	"net/http"

	"github.com/migopp/gocards/internal/debug"
)

// Representation of the running server in memory
type Server struct {
	IP   string
	Port uint16
}

// Starts running the designated server
//
// If it fails, then we exit with an error,
// either on start-up or in the middle of handling --
// We can't know...
func (s *Server) Run() error {
	// Configure router
	initHandlers()

	// Decide where to serve
	serveAddr := fmt.Sprintf("%s:%d", s.IP, s.Port)

	// Serve
	//
	// `ListenAndServe` returns an error, so we will
	// just bubble it up when it happens
	debug.Printf("| Starting server at %s\n", serveAddr)
	return http.ListenAndServe(serveAddr, nil)
}
