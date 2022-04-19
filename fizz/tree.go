package fizz

import "strings"

type node struct {
	path     string
	part     string
	children []*node
	isWild   bool
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(path string, parts []string, level int) {
	if len(parts) == level {
		n.path = path
		return
	}
	part := parts[level]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(path, parts, level+1)
}

func (n *node) search(parts []string, level int) *node {
	if len(parts) == level || strings.HasPrefix(n.part, "*") {
		if n.path == "" {
			return nil
		}
		return n
	}

	part := parts[level]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, level+1)
		if result != nil {
			return result
		}
	}
	return nil
}
