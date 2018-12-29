package main

import (
	"log"
	"os"
	"sort"
	"strings"
)

type autocomplete struct {
	root node
}

func newAutocompleteFromFileName(fileName string) autocomplete {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]byte, 10000000)
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	s := string(data)
	xs := strings.Split(s, "\n")
	return newAutocomplete(xs)
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
