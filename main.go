package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	autocomplete := newAutocompleteFromFileName("words.txt")
	fmt.Println("Ready! Type a string and press enter.")

	for {
		fmt.Print("\n=>")
		s := getInput()
		suggs := autocomplete.getSuggestions(s)
		fmt.Println("Your suggestions: \n", suggs)
	}
}

func getInput() string {
	input := make([]byte, 20)
	file := *os.Stdin
	_, err := file.Read(input)
	if err != nil {
		log.Fatal(err)
	}

	s := string(input)
	s = strings.Trim(s, "\x00")
	s = strings.Trim(s, "\n")
	return s
}
