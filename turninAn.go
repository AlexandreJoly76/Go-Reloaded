package main

import (
	"strings"
)

func TurnInAn(words []string) []string {
	for i, word := range words {
		if word == "a" || word == "A" {
			firstLetter := strings.ToLower(string(words[i+1][0]))
			if firstLetter == "a" || firstLetter == "e" || firstLetter == "i" || firstLetter == "o" || firstLetter == "u" || firstLetter == "h" {
				words[i] += "n"
			}
		}
	}
	return words
}
