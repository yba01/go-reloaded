package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 3 {
		content, _ := ioutil.ReadFile(os.Args[1])
		cont := correct(string(content))
		result_cont := manip(strings.Fields(cont))
		err := ioutil.WriteFile(os.Args[2], []byte(final_reslt), 0644)
		if err != nil {
			os.Exit(0)
		}
	}
}

// rewrite text
func correct(s string) string {
	a := strings.ReplaceAll(s, "(", " (")
	a = strings.ReplaceAll(a, ")", ") ")
	return a
}

// 5 particulars keys words functions
func hex(s string) string {
	a, _ := strconv.ParseInt(s, 16, 64)
	s = strconv.Itoa(int(a))
	return s
}
func bin(s string) string {
	a, _ := strconv.ParseInt(s, 2, 64)
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

// text's manipulation by simple keys words....
func manip(slice []string) []string {
	type App struct {
		name string
		do   func(string) string
	}
	man_hex := App{name: "(hex)", do: hex}
	man_bin := App{name: "(bin)", do: bin}
	man_up := App{name: "(up)", do: up}
	man_low := App{name: "(low)", do: low}
	man_cap := App{name: "(cap)", do: cap}

	list := []App{man_hex, man_bin, man_up, man_low, man_cap}

	for index, word := range slice {
		for _, mot := range list {
			if word == mot.name {
				if index > 0 {
					slice[index-1] = mot.do(slice[index-1])
					slice = remove(slice, index)
				} else if index == 0 {
					slice = remove(slice, index)
				}
			}
		}
	}
	return slice
}

// remove key word in text
func remove(slc []string, i int) []string {
	return append(slc[:i], slc[i+1:]...)
}

func gestion_ponc(do []string) []string {
	for index, mot := range do {
		char := rune(mot[0])
		id_ponc := []rune{'.', ',', '!', '?', ':', ';'}
		if Inslice(char, id_ponc) {
			do[index-1] = do[index-1] + ","
			do[index] = mot[1:]
		}
	}
	return do
}
