// Package webserver contains http webserver implementation
// and associated route handlers.
package webserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/nmezhenskyi/go-rest-api-example/internal/model"
	"github.com/nmezhenskyi/go-rest-api-example/internal/storage"
)

type Server struct {
	httpServer *http.Server
	router     http.Handler
	storage    *storage.Storage

	Logger *log.Logger
}

func NewServer() *Server {
	s := &Server{
		httpServer: &http.Server{},
		storage:    storage.NewStorage(),
		Logger:     log.Default(),
	}
	s.routes()
	return s
}

func (s *Server) ListenAndServe(addr string) error {
	s.httpServer.Addr = addr
	s.httpServer.Handler = s.router
	s.Logger.Printf("Server is starting on %s\n", addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Helpers:

func (s *Server) PopulateWithData(file string) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	var records = []model.Wine{}
	err = json.Unmarshal(bytes, &records)
	if err != nil {
		return err
	}

	for _, rec := range records {
		s.storage.Save(rec.ID, rec)
	}

	return nil
}

func (s *Server) RemoveData() {
	s.storage.SetEmpty()
}
