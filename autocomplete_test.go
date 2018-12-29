package autocomplete

import (
	"testing"
)

func words() []string {
	return []string{
		"aardvark",
		"apple",
		"dinosaur",
		"potato",
	}
}

func TestNewAutocomplete(t *testing.T) {
	w := words()
	a := newAutocomplete(w)

	if len(a.root.children) != 3 {
		t.Errorf("Expected 3, got %v", len(a.root.children))
	}
}

func TestGetNodeWhenDoesNotExist(t *testing.T) {
	a := newAutocomplete([]string{})
	n := a.getNode("a")

	if n != nil {
		t.Errorf("Expected nil, got %v", a)
	}
}

func TestGetNodeWhenExists(t *testing.T) {
	w := words()
	a := newAutocomplete(w)
	n := a.getNode("potat")

	if n.letter != "t" {
		t.Errorf("Expected t, got %v", n.letter)
	}
}

func TestSuggestionsNoWordsPopulated(t *testing.T) {
	a := newAutocomplete([]string{})
	xs := a.getSuggestions("a")

	if len(xs) != 0 {
		t.Errorf("Expected 0, got %v", len(xs))
	}
}
func TestSuggestionsOneSuggestion(t *testing.T) {
	w := words()
	a := newAutocomplete(w)
	s := a.getSuggestions("po")

	if s[0] != "potato" {
		t.Errorf("Expected potato, got %v", s)
	}
}
func TestSuggestionsTwoSuggestions(t *testing.T) {
	w := words()
	a := newAutocomplete(w)
	s := a.getSuggestions("a")

	if s[0] != "aardvark" {
		t.Errorf("Expected aardvark, got %v", s[0])
	}
	if s[1] != "apple" {
		t.Errorf("Expected apple, got %v", s[1])
	}
}
