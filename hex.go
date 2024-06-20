package main

import "strconv"

func Hex(words []string, i int) []string {
	DecimalValue, err := strconv.ParseInt(words[i-1], 16, 64)
	check(err)
	words[i-1] = strconv.Itoa(int(DecimalValue))
	words = append(words[:i], words[i+1:]...)
	return words
}
