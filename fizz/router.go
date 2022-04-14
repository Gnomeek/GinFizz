package fizz

import "fmt"

// Router stores routeMap that track kv relationship between pattern and handler
type Router struct {
	routeMap map[string]HandlerFunc
}

// NewRouter returns a Router with empty routeMap
func NewRouter() *Router {
	return &Router{routeMap: make(map[string]HandlerFunc)}
}

// AddRoute adds route into routeMap of Router
func (router Router) AddRoute(pattern string, handler HandlerFunc) {
	router.routeMap[pattern] = handler
}

// GetRoute gets corresponding handler of pattern if pattern exists in routeMap of Router else throw an error
func (router Router) GetRoute(pattern string) (HandlerFunc, error) {
	if route, ok := router.routeMap[pattern]; ok {
		return route, nil
	} else {
		return nil, fmt.Errorf("route pattern %s not registered", pattern)
	}
}
