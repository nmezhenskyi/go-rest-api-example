package webserver

import (
	"fmt"
	"net/http"
)

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		s.Logger.Printf("Received %s %q", req.Method, req.URL.Path)
		fmt.Fprint(w, "Welcome to Winery!")
	}
}
