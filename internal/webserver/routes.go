package webserver

import (
	"fmt"
	"net/http"

	"github.com/nmezhenskyi/go-rest-api-example/internal/router"
)

func (s *Server) routes() {
	mux := &router.Router{}

	mux.Route("GET", "/", s.handleIndex())
	mux.Route("GET", "/api/wine", s.handleWineGetAll())
	mux.Route("POST", "/api/wine", s.handleWineCreate())
	mux.Route("GET", `/api/wine/(?P<id>\d+)`, s.handleWineGetOne())
	mux.Route("PUT", `/api/wine/(?P<id>\d+)`, s.handleWineUpdate())
	mux.Route("DELETE", `/api/wine/(?P<id>\d+)`, s.handleWineDelete())

	s.router = mux
}

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		s.Logger.Printf("Received %s %q", req.Method, req.URL.Path)
		fmt.Fprint(w, "Welcome to Winery!")
	}
}
