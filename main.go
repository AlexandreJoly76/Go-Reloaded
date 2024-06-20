package main

import (
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	if len(os.Args) < 3 {
		os.Exit(0)
	}
	file := os.Args[1]
	dat, err := os.ReadFile(file)
	check(err)
	words := strings.Split(string(dat), " ")
	words = TurnInAn(words)
	for i, word := range words {
		if word == "(hex)" {
			words = Hex(words, i)
		}
		if word == "(bin)" {
			words = Bin(words, i)
		}
		if word == "(up)" {
			words = Up(words, i, 1)
		} else if word == "(up," {
			nb, _ := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			words = Up(words, i, nb)
		}
		if word == "(low)" {
			words = Low(words, i, 1)
		} else if word == "(low," {
			nb, _ := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			words = Low(words, i, nb)
		}
		if word == "(cap)" {
			words = Cap(words, i, 1)
		} else if word == "(cap," {
			nb, _ := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			words = Cap(words, i, nb)
		}
	}
	result := strings.Join(words, " ")
	result = formatPunctuation(result)
	os.WriteFile(os.Args[2], []byte(result), 0777)
}
