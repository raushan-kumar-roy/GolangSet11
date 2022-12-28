package main

import (
	"fmt"
)

func filter(slice []string, f func(string) bool) []string {
	filtered := make([]string, 0)
	for _, v := range slice {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	words := []string{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth"}

	startsWithB := filter(words, func(v string) bool {
		return v[0] == 'b'
	})
	fmt.Println("Words starting with 'b':")
	fmt.Println(startsWithB)

	threeLetters := filter(words, func(v string) bool {
		return len(v) == 3
	})
	fmt.Println("3 character words:")
	fmt.Println(threeLetters)
}
