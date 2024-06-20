package main

func formatPunctuation(s string) string {
	puncs := "!,.?;:"
	count := 0
	for i, char := range s {
		for j := 0; j < len(puncs); j++ {
			if char == rune(puncs[j]) {
				if s[i-1-count] == ' ' {
					s = s[:i-1-count] + s[i-count:]
					count++
				}
				if i-count+2 < len(s) && s[i-count+1] != '!' && s[i-count+1] != '?' && s[i-count+1] != '.' && s[i-count+1] != ' ' {
					s = s[:i-count+1] + " " + s[i-count+1:]
					count--
				}
			}
		}
	}
	countApostrophe := 0
	countDel := 0
	for i, char := range s {
		if string(char) == "'" {
			if countApostrophe%2 == 0 && string(s[i+1-countDel]) == " " {
				s = string(s[:i+1-countDel]) + string(s[i+2-countDel:])
				countDel++
			}
			if countApostrophe%2 != 0 && string(s[i-1-countDel]) == " " {
				s = string(s[:i-1-countDel]) + string(s[i-countDel:])
				countDel++
			}
			countApostrophe++
		}
	}
	return s
}
