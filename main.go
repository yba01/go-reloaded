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
		result_cont := manip(string(content))
		final_reslt := affiche(result_cont)
		err := ioutil.WriteFile(os.Args[2], []byte(final_reslt), 0644)
		if err != nil {
			os.Exit(0)
		}
	}
}

// 2.les fonctions qui donnent les differentes changements voulus dans le texte:les parametres
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
	s = strings.Title(s)
	return s
}
func Inslice(a rune) bool {
	ind := []rune{'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H'}
	for _, char := range ind {
		if a == char {
			return true
		}
	}
	return false
}
func change(a string) string {
	if a == "a" {
		return "an"
	} else {
		return "An"
	}
}
func affiche(back []string) string {
	last := ""
	for _, mot := range back {
		last = last + " " + mot
	}
	return last
}

// une fonction qui trouvent l'index des parametres et change le texte
func manip(sent string) []string {
	newTab := strings.Fields(sent)
	for index, word := range newTab {
		if word == "(hex)" {
			newTab[index-1] = hex(newTab[index-1])
			newTab[index] = ""
		} else if word == "(bin)" {
			newTab[index-1] = bin(newTab[index-1])
			newTab[index] = ""
		} else if word == "(up)" {
			newTab[index-1] = up(newTab[index-1])
			newTab[index] = ""
		} else if word == "(low)" {
			newTab[index-1] = low(newTab[index-1])
			newTab[index] = ""
		} else if word == "(cap)" {
			newTab[index-1] = cap(newTab[index-1])
			newTab[index] = ""
		} else if word == "(up," {
			next := newTab[index+1]
			end, _ := strconv.Atoi(next[:len(next)-1])
			for i := 1; i <= end; i++ {
				newTab[index-i] = up(newTab[index-i])
			}
			newTab[index] = ""
			newTab[index+1] = ""
		} else if word == "(low," {
			next := newTab[index+1]
			end, _ := strconv.Atoi(next[:len(next)-1])
			for i := 1; i <= end; i++ {
				newTab[index-i] = low(newTab[index-i])
			}
			newTab[index] = ""
			newTab[index+1] = ""
		} else if word == "(cap," {
			next := newTab[index+1]
			end, _ := strconv.Atoi(next[:len(next)-1])
			for i := 1; i <= end; i++ {
				newTab[index-i] = cap(newTab[index-i])
			}
			newTab[index] = ""
			newTab[index+1] = ""
		} else if word == "a" || word == "A" {
			others := newTab[index+1]
			if Inslice(rune(others[0])) {
				newTab[index] = change(word)
			}
		}
	}
	return newTab
}

//les changements qui ne n'cessitent pas de paramètre
//Résultat écrit dans un autre fichier .txt
