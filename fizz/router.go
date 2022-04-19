package fizz

import (
	"net/http"
	"strings"
)

// Router stores routeMap that track kv relationship between pattern and handler
type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// NewRouter returns a Router with empty routeMap
func NewRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// Only one * is allowed
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// AddRoute adds route into handlers of Router
func (router *router) addRoute(method, path string, handler HandlerFunc) {
	parts := parsePattern(path)
	pattern := router.generatePattern(method, path)
	if _, ok := router.roots[method]; !ok {
		router.roots[method] = &node{}
	}
	router.roots[method].insert(path, parts, 0)
	router.handlers[pattern] = handler
}

// GetRoute gets corresponding handler of pattern if pattern exists in routeMap of Router else throw an error
func (router *router) getRoute(method, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := router.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.path)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (router *router) generatePattern(method, path string) string {
	return method + "-" + path
}

func (router *router) handle(c *Context) {
	n, params := router.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := router.generatePattern(c.Method, n.path)
		c.handlers = append(c.handlers, router.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
