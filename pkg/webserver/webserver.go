// Package webserver contains http webserver implementation
// and associated route handlers.
package webserver

import (
	"log"
	"net/http"

	"github.com/nmezhenskyi/go-rest-api-example/pkg/storage"
)

type Server struct {
	router  http.Handler
	storage *storage.Storage
	Logger  *log.Logger
}

func NewServer() *Server {
	s := &Server{
		storage: storage.NewStorage(),
		Logger:  log.Default(),
	}
	s.routes()
	return s
}

func (s *Server) ListenAndServe(addr string) error {
	httpServer := http.Server{
		Addr:    addr,
		Handler: s.router,
	}
	s.Logger.Printf("Server is starting on %s\n", addr)
	return httpServer.ListenAndServe()
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
