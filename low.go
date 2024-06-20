package main

import (
	"strings"
)

func Low(words []string, i int, nb int) []string {
	for j := 1; j <= nb; j++ {
		words[i-j] = strings.ToLower(words[i-j])
	}
	if nb == 1 {
		words = append(words[:i], words[i+1:]...)
	} else {
		words = append(words[:i], words[i+2:]...)
	}
	return words
}
