package piscine

import (
	"fmt"
)

func PrintWordsTables(a []string) {
	for _, word := range a {
		fmt.Println(word)
	}
}

func SplitWhiteSpaces(s string) []string {
	var words []string
	var word []rune

	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			if len(word) > 0 {
				words = append(words, string(word))
				word = nil
			}
		} else {
			word = append(word, r)
		}
	}
	if len(word) > 0 {
		words = append(words, string(word))
	}
	return words
}
