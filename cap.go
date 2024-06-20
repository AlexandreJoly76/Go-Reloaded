package main

import "strings"

func Cap(words []string, i int, nb int) []string {
	for j := 1; j <= nb; j++ {
		words[i-j] = strings.ToUpper(string(words[i-j][0])) + words[i-j][1:]
	}
	if nb == 1 {
		words = append(words[:i], words[i+1:]...)
	} else {
		words = append(words[:i], words[i+2:]...)
	}
	return words
}
