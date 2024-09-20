package api

import "net/http"

type Handler interface {
	Register(http.ResponseWriter, *http.Request)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Route struct {
	Method string
	Path   string
}

type Router struct {
	routes map[Route]HandlerFunc
}

func NewRouter(handler Handler, r map[Route]HandlerFunc) *Router {
	return &Router{routes: r}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	route := Route{Method: req.Method, Path: req.URL.Path}
	handler, exists := r.routes[route]
	if !exists {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	handler(w, req)
}
