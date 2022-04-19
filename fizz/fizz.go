package fizz

import (
	"net/http"
)

// HandlerFunc is customized based on http.HandlerFunc
type HandlerFunc func(c *Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	engine      *Engine
}

// Engine stores router info, etc.(TBD)
type Engine struct {
	*RouterGroup
	router       *router
	routerGroups []*RouterGroup
}

// New creates a new engine with router without routes
func New() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.routerGroups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// addRoute register handler into route,
// the key is concat by method, slash, and path as the key while handler as value
func (engine *Engine) addRoute(method string, path string, handler HandlerFunc) {
	engine.router.addRoute(method, path, handler)
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
	context := NewContext(w, req)
	engine.router.handle(context)
}

// Run starts an HTTP service
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}
	engine.routerGroups = append(engine.routerGroups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}
