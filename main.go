// You can edit this code!
// Click here and start typing.
package main

import (
	"strings"
)

func Ponc(s string) string {
	for i := 0; i < len(strings.Fields(s)); i++ {
		s = Single_Ponc(s)
	}
	return Quote_handle(s)
}

// handle one or more ponctuation
func Single_Ponc(s string) string {
	str := strings.Fields(s)
	ind := []rune{'.', ',', ';', ':', '!', '?'}
	for i, word := range str {
		if i > 0 {
			a := 0
			if Inslice(rune(word[0]), ind) {
				for j := 0; j < len(word); j++ {
					if Inslice(rune(word[j]), ind) {
						str[i-1] = str[i-1] + string(word[j])
						a++
					}
				}
			}
			if len(word) > 0 {
				if Inslice(rune(word[0]), ind) {

					str[i] = str[i][a:]
				}
			}
		}
	}

	return write(str)
}

// presence of character in slice
func Inslice(a rune, ind []rune) bool {
	for _, char := range ind {
		if a == char {
			return true
		}
	}
	return false
}

// rewrite function in right way
func write(s []string) string {
	var s0 []string
	for _, word := range s {
		if word != "" {
			s0 = append(s0, word)
		}
	}
	return strings.Join(s0, " ")
}

// handle quote or double quotes
func Quote_handle(s string) string {
	str := strings.Fields(s)
	new := make(map[int]int)
	d := 0
	for i, word := range str {
		if word == "'" {
			d = d + 1
			new[i] = d
		}
	}
	for index, score := range new {
		if score%2 != 0 {
			if index < len(str)-1 {
				str[index+1] = "'" + str[index+1]
				str[index] = ""
			}
		} else {
			if index > 0 {
				str[index-1] = str[index-1] + "'"
				str[index] = ""
			}
		}
	}
	return write(str)
}
