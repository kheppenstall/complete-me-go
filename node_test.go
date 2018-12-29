package main

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	n := newNode("a", true)

	if n.letter != "a" {
		t.Errorf("Expected a, got %v", n.letter)
	}

	if len(n.children) != 0 {
		t.Errorf("Expected 0, got %v", len(n.children))
	}

	if n.isTerminus != true {
		t.Errorf("Expected true, got false")
	}
}
func TestNewNodeUppercaseLetter(t *testing.T) {
	n := newNode("A", false)
	if n.letter != "a" {
		t.Errorf("Expected a, got %v", n.letter)
	}
}

func TestRoot(t *testing.T) {
	r := root()

	if r.letter != "" {
		t.Errorf("Expected '', got %v", r.letter)
	}
}

func TestAddChild(t *testing.T) {
	r := root()
	a := r.addChild("a", false)

	if len(r.children) != 1 {
		t.Errorf("Expected 1, got %v", len(r.children))
	}
	if a.letter != "a" {
		t.Errorf("Expected a, got %v", a.letter)
	}
}

func TestChildNoDuplicate(t *testing.T) {
	r := root()
	r.addChild("a", false)
	r.addChild("a", false)

	if len(r.children) != 1 {
		t.Errorf("Expected 1, got %v", len(r.children))
	}
}

func TestInsertOneLetterWord(t *testing.T) {
	r := root()
	r.insert("a")

	if r.children["a"].isTerminus != true {
		t.Errorf("Expected true, got false")
	}
}
func TestInsertTwoLetterWord(t *testing.T) {
	r := root()
	r.insert("an")

	rChildren := r.children
	if rChildren["a"].isTerminus != false {
		t.Errorf("Expected false, got true")
	}

	aChildren := rChildren["a"].children
	if aChildren["n"].isTerminus != true {
		t.Errorf("Expected true, got false")
	}
}
func TestInsertThreeLetterWord(t *testing.T) {
	r := root()
	r.insert("and")

	rChildren := r.children
	if rChildren["a"].isTerminus != false {
		t.Errorf("Expected false, got true")
	}

	aChildren := rChildren["a"].children
	if aChildren["n"].isTerminus != false {
		t.Errorf("Expected false, got true")
	}

	nChildren := aChildren["n"].children
	if nChildren["d"].isTerminus != true {
		t.Errorf("Expected true, got false")
	}
}
