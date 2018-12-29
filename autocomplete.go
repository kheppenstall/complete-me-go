package autocomplete

import "sort"

type autocomplete struct {
	root node
}

func newAutocomplete(words []string) autocomplete {
	r := root()
	for _, w := range words {
		r.insert(w)
	}
	return autocomplete{root: r}
}

func (a autocomplete) getSuggestions(s string) []string {
	var suggs []string
	n := a.getNode(s)
	if n == nil {
		return suggs
	}

	suggs = a.addSuggestions(s, *n, suggs)
	sort.Strings(suggs)
	return suggs
}

func (a autocomplete) addSuggestions(s string, n node, suggs []string) []string {
	for l, c := range n.children {
		if c.isTerminus {
			suggs = append(suggs, s+l)
		}
		suggs = a.addSuggestions(s+l, c, suggs)
	}
	return suggs
}

func (a autocomplete) getNode(s string) *node {
	n := a.root
	for _, r := range s {
		c, ok := n.children[string(r)]
		if !ok {
			return nil
		}
		n = c
	}
	return &n
}
