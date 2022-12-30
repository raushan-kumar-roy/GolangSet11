package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth"}

	capitalizedWords := Map(words, func(word string) string {
		return strings.Title(word)
	})
	colonWords := Map(words, func(word string) string {
		return word + ":"
	})
	fmt.Println(capitalizedWords)
	fmt.Println(colonWords)
}

func Map(slice []string, fn func(string) string) []string {
	ModifiedStrings := make([]string, len(slice))
	for i, v := range slice {
		ModifiedStrings[i] = fn(v)
	}
	return ModifiedStrings
}
