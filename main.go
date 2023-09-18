package main

import (
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 3 {
		content, _ := ioutil.ReadFile(os.Args[1])
		cont := correct(string(content))
		cont = text_manip(cont)
		err := ioutil.WriteFile(os.Args[2], []byte(cont), 0644)
		if err != nil {
			return
		}
	}
}

// 1.TEXT FORMATAGE

// remove space
func remove_space(s string) string {
	corr0 := strings.Fields(s)
	corr1 := strings.Join(corr0, " ")
	return corr1
}

// handle parenthesis
func parenthese(s string) string {
	corr2 := strings.ReplaceAll(s, "(", " (")
	corr3 := strings.ReplaceAll(corr2, ")", ") ")
	corr4 := strings.ReplaceAll(corr3, "( ", "(")
	corr5 := strings.ReplaceAll(corr4, " )", ")")
	return corr5
}

// place pattern of ponctuation on a right way
func ponct(s string) string {
	str := remove_space(s)
	pattern := "[.,;:!?]+"
	regex, _ := regexp.Compile(pattern)
	matches := regex.FindAllString(str, -1)
	for _, motif := range matches {
		intxt := " " + motif
		replce := motif + " "
		str = strings.ReplaceAll(str, intxt, replce)
	}
	return str
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

// change a or A to an or An
func change(a string) string {
	if a == "a" {
		return "an"
	} else {
		return "An"
	}
}

// Print text in a good format
func Affichage(str []string) string {
	return strings.Join(str, " ")
}

// handle an or An in text
func Atoan(str string) string {
	str0 := strings.Fields(str)
	ind := []rune{'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H'}
	for index, word := range str0 {
		if word == "a" || word == "A" {
			if index+1 < len(str0)-1 {
				others := str0[index+1]
				if Inslice(rune(others[0]), ind) {
					str0[index] = change(word)
				}
			}
		}
	}
	return Affichage(str0)
}

// format text
func correct(s string) string {
	co1 := Atoan(s)
	co2 := ponct(co1)
	co3 := parenthese(co2)
	corrected := remove_space(co3)
	return corrected
}

//2.TEXT MANIPULATION

// particulars keys words functions
func hex(s string) string {
	a, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return s
	}
	s = strconv.Itoa(int(a))
	return s
}
func bin(s string) string {
	a, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return s
	}
	s = strconv.Itoa(int(a))
	return s
}
func up(s string) string {
	s = strings.ToUpper(s)
	return s
}
func low(s string) string {
	s = strings.ToLower(s)
	return s
}
func cap(s string) string {
	s = strings.ToLower(s)
	s = strings.Title(s)
	return s
}

// text's manipulation by simple keys words in particular index
func manip(slice []string, i int) []string {
	tabStruct := []struct {
		name string
		do   func(string) string
	}{
		{"(hex)", hex},
		{"(bin)", bin},
		{"(up)", up},
		{"(low)", low},
		{"(cap)", cap},
		{"(hex))", hex},
		{"(bin))", bin},
		{"(up))", up},
		{"(low))", low},
		{"(cap))", cap},
	}
	for _, mot := range tabStruct {
		if i < len(slice) {
			if slice[i] == mot.name {
				if i > 0 {
					slice[i-1] = mot.do(slice[i-1])
					slice = remove(slice, i)
				} else if i == 0 {
					slice = remove(slice, i)
				}
			}
		}
	}
	return slice
}

// text's manipulation by simple keys words....
func base_manip(cont []string) []string {
	i := 0
	lon := len(cont)
	for i < len(cont) {
		cont = manip(cont, i)
		if lon == len(cont) {
			i = i + 1
		} else if len(cont) < lon {
			lon = len(cont)
		}
	}
	return cont
}

// remove key word in text
func remove(slc []string, i int) []string {
	return append(slc[:i], slc[i+1:]...)
}

// for using different parameters on more than one string by example: (up, 3)
func iter_funct(str []string, faire func(string) string, num int, index int) []string {
	for i := index - 1; i >= index-num; i-- {
		if i >= 0 {
			str[i] = faire(str[i])
		}
	}
	str = remove(str, index)
	str = remove(str, index)
	return str
}

// find index of pattern in slice
func index_end(str []string, debut int) int {
	var result int
	for i := debut; i < len(str); i++ {
		match, _ := regexp.MatchString(`\w+\)\)`, str[i])
		if match {
			result = i + 1
			break
		}
	}
	return result
}

// text's manipulation in case of more complicated keys word like (cup, A (hex) (bin))
func manip_complex(s []string) []string {
	Tablestruct := []struct {
		name string
		do   func(string) string
	}{
		{"(low,", low},
		{"(up,", up},
		{"(cap,", cap},
	}
	for i, word := range s {
		for _, mot := range Tablestruct {
			if word == mot.name {
				if i < len(s)-1 {
					pattern := `^[0-9]\)$`
					match, _ := regexp.MatchString(pattern, s[i+1])
					if match {
						stop, _ := strconv.Atoi(s[i+1][:len(s[i+1])-1])
						s = iter_funct(s, mot.do, stop, i)
					} else {
						if i+1 < len(s)-1 {
							if i+1 < len(s)-1 {
								res := base_manip(s[i+1 : index_end(s, i)+1])
								s[i+1] = res[0] + ")"
								s = multi_remove(s, i+2, index_end(s, i)-i-1)
								s = manip_complex(s)
							}
						}
					}
				}
			}
		}
	}
	return s
}

// delete more than one string in string's slice
func multi_remove(str []string, index int, time int) []string {
	for i := 0; i < time; i++ {
		str = remove(str, index)
	}
	return str
}

// manipulation text
func text_manip(str string) string {
	str0 := strings.Fields(str)
	str0 = base_manip(str0)
	str0 = manip_complex(str0)
	return correct(Affichage(str0))
}
