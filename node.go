package autocomplete

import (
	"strings"
)

type node struct {
	letter     string
	children   map[string]node
	isTerminus bool
}

func root() node {
	return newNode("", false)
}

func newNode(l string, isTerminus bool) node {
	c := make(map[string]node)
	return node{
		letter:     strings.ToLower(l),
		children:   c,
		isTerminus: isTerminus,
	}
}

func (n *node) addChild(l string, isTerminus bool) *node {
	if c, ok := n.children[l]; ok {
		return &c
	}

	c := newNode(l, isTerminus)
	n.children[l] = c
	return &c
}

func (n *node) insert(w string) {
	isTerminus := len(w) == 1
	c := n.addChild(w[:1], isTerminus)

	if !isTerminus {
		c.insert(w[1:])
	}
}
