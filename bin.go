package main

import "strconv"

func Bin(words []string, i int) []string {
	decimalValue, err := strconv.ParseInt(words[i-1], 2, 64)
	check(err)
	words[i-1] = strconv.Itoa(int(decimalValue))
	words = append(words[:i], words[i+1:]...)
	return words
}
