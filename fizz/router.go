package fizz

import (
	"fmt"
	"net/http"
)

// Router stores routeMap that track kv relationship between pattern and handler
type Router struct {
	routeMap map[string]HandlerFunc
}

// NewRouter returns a Router with empty routeMap
func NewRouter() *Router {
	return &Router{routeMap: make(map[string]HandlerFunc)}
}

// AddRoute adds route into routeMap of Router
func (router Router) AddRoute(method, path string, handler HandlerFunc) {
	pattern := router.generatePattern(method, path)
	router.routeMap[pattern] = handler
}

// GetRoute gets corresponding handler of pattern if pattern exists in routeMap of Router else throw an error
func (router Router) GetRoute(req *http.Request) (HandlerFunc, error) {
	pattern := router.generatePattern(req.Method, req.URL.Path)
	if route, ok := router.routeMap[pattern]; ok {
		return route, nil
	} else {
		return nil, fmt.Errorf("route pattern %s not registered", pattern)
	}
}

func (router Router) generatePattern(method, path string) string {
	return method + "-" + path
}
