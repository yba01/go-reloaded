// You can edit this code!
// Click here and start typing.
package main

import (
	"strconv"
	"strings"
)

// handle all about ponctuation
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

// handle all paranthesis
func parenthese(s string) string {
	TabPonc := []struct {
		in  string
		rep string
	}{
		{".", " . "},
		{",", " , "},
		{";", " ; "},
		{":", " : "},
		{"!", " ! "},
		{"?", " ? "},
		{"(", " ( "},
		{")", " ) "},
		{"( hex )", "(hex)"},
		{"( bin )", "(bin)"},
		{"( up )", "(up)"},
		{"( low )", "(low)"},
		{"( cap )", "(cap)"},
		{" ) ", ") "},
	}
	for _, char := range TabPonc {
		s = strings.ReplaceAll(s, char.in, char.rep)
	}
	return s
}

// delete something on slice
func remove(s []string, index int) []string {
	var str []string
	if index > 0 && index < len(s) {
		for i := 0; i < len(s); i++ {
			if i != index {
				str = append(str, s[i])
			}
		}
	}
	return str
}

// simple manipulation mean simple key word
func manip_hex(s []string) []string {
	ind := []rune{'.', ',', ';', ':', '!', '?', '\''}
	for i := 0; i < len(s); i++ {
		if s[i] == "(hex)" {
			if i-1 >= 0 {
				if Inslice(rune(s[i-1][0]), ind) {
					a := s[i-1]
					s[i-1] = s[i]
					s[i] = a
					return manip_hex(s)
				} else {
					a, err := strconv.ParseInt(s[i-1], 16, 64)
					if err != nil {
						return remove(s, i)
					}
					s[i-1] = strconv.Itoa(int(a))
					s = remove(s, i)
				}
			} else {
				return remove(s, i)
			}
		}
	}
	return s
}
func manip_bin(s []string) []string {
	ind := []rune{'.', ',', ';', ':', '!', '?', '\''}
	for i := 0; i < len(s); i++ {
		if s[i] == "(bin)" {
			if i-1 >= 0 {
				if Inslice(rune(s[i-1][0]), ind) {
					a := s[i-1]
					s[i-1] = s[i]
					s[i] = a
					return manip_bin(s)
				} else {
					a, err := strconv.ParseInt(s[i-1], 2, 64)
					if err != nil {
						return remove(s, i)
					}
					s[i-1] = strconv.Itoa(int(a))
					s = remove(s, i)
				}
			} else {
				return remove(s, i)
			}
		}
	}
	return s
}
