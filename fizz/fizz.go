package fizz

import (
	"fmt"
	"net/http"
)

// HandlerFunc is http.HandlerFunc
type HandlerFunc http.HandlerFunc

// Engine stores router info, etc.(TBD)
type Engine struct {
	router *Router
}

// New creates a new engine with router without routes
func New() *Engine {
	return &Engine{router: NewRouter()}
}

// addRoute register handler into route,
// the key is concat by method, slash, and path as the key while handler as value
func (engine *Engine) addRoute(method string, path string, handler HandlerFunc) {
	pattern := method + "-" + path
	engine.router.AddRoute(pattern, handler)
}

// GET defines HTTP GET method
func (engine *Engine) GET(path string, handler HandlerFunc) {
	engine.addRoute("GET", path, handler)
}

// POST defines HTTP POST method
func (engine *Engine) POST(path string, handler HandlerFunc) {
	engine.addRoute("POST", path, handler)
}

// ServeHTTP is implemented from http.Handler interface
// so that engine can be used in http.ListenAndServe
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	pattern := req.Method + "-" + req.URL.Path
	if handler, err := engine.router.GetRoute(pattern); err == nil {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "%v 404 NOT FOUND with error %v", req.URL, err)
	}
}

// Run starts an HTTP service
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}