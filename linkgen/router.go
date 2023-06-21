package linkgen

import (
	"context"
	"net/http"
	"strings"
)

type key string

type contextKey key

func (c contextKey) String() string {
	return "linkgen." + string(c)
}

const (
	paramsKey = contextKey("params")
)

// route - describes a route with pattern and handler
type route struct {
	pattern string
	handler http.HandlerFunc
}

// Router - Stores route mapping by method
type Router struct {
	routes map[string][]route
}

// NewRouter - Creates a new instance of Router
func NewRouter() *Router {
	return &Router{routes: map[string][]route{}}
}

// addRoute - Add a route for the given method to the router
func (r *Router) addRoute(method string, pattern string, handler http.HandlerFunc) {
	r.routes[method] = append(r.routes[method], route{pattern: pattern, handler: handler})
}

// Serve - returns a handler for the built router
func (r *Router) Serve() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		ctx := context.WithValue(req.Context(), paramsKey, map[string]string{})
		for _, route := range r.routes[req.Method] {
			if match, newContext := matchPath(ctx, path, route.pattern); match {
				route.handler.ServeHTTP(w, req.WithContext(newContext))
				return
			}
		}
		http.NotFound(w, req)
	})
}

// matchPath - check if the path matches a route pattern
func matchPath(ctx context.Context, path, pattern string) (bool, context.Context) {
	pathChunks := strings.Split(path, "/")
	patternChunks := strings.Split(pattern, "/")

	if len(pathChunks) != len(patternChunks) {
		return false, ctx
	}

	for i, chunk := range patternChunks {
		// This is a parameter
		if len(chunk) > 0 && chunk[0] == ':' {
			ctxValues := ctx.Value(paramsKey).(map[string]string)
			ctxValues[chunk[1:]] = pathChunks[i]
			ctx = context.WithValue(ctx, paramsKey, ctxValues)
		} else if chunk != pathChunks[i] {
			return false, ctx
		}
	}

	return true, ctx
}
