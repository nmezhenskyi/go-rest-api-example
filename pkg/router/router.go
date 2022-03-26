// Provides routing functionality for an http server.
package router

import (
	"context"
	"log"
	"net/http"
	"regexp"
)

type RouteEntry struct {
	Path    *regexp.Regexp
	Method  string
	Handler http.HandlerFunc
}

func (entry *RouteEntry) Match(req *http.Request) map[string]string {
	match := entry.Path.FindStringSubmatch(req.URL.Path)
	if match == nil {
		return nil
	}
	if entry.Method != req.Method {
		return nil
	}

	params := make(map[string]string)
	groupNames := entry.Path.SubexpNames()
	for i, group := range match {
		params[groupNames[i]] = group
	}

	return params
}

type Router struct {
	routes []RouteEntry
}

type contextKey string

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if req := recover(); req != nil {
			log.Println("Error:", req)
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}
	}()

	for _, entry := range router.routes {
		params := entry.Match(req)
		if params == nil {
			continue
		}

		ctx := context.WithValue(req.Context(), contextKey("params"), params)
		entry.Handler.ServeHTTP(w, req.WithContext(ctx))
		return
	}

	http.NotFound(w, req)
}

func (router *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	exactPath := regexp.MustCompile("^" + path + "$")
	entry := RouteEntry{
		Method:  method,
		Path:    exactPath,
		Handler: handlerFunc,
	}

	router.routes = append(router.routes, entry)
}

func URLParam(req *http.Request, name string) string {
	ctx := req.Context()
	params := ctx.Value(contextKey("params")).(map[string]string)
	return params[name]
}
