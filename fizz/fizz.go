package fizz

import (
	"net/http"
)

// HandlerFunc is customized based on http.HandlerFunc
type HandlerFunc func(c *Context)

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
	engine.router.AddRoute(method, path, handler)
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
	if handler, err := engine.router.GetRoute(req); err == nil {
		context := NewContext(w, req)
		handler(context)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Run starts an HTTP service
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
