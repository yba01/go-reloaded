// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// handle all about ponctuation
func Ponc(s string) string {
	s = Single_Ponc(s)
	s = Rep_parenthese(s)
	s = Ahandled(s)
	return Quote_handle(s)
}

// handle articles
func Ahandled(s string) string {
	str := strings.Fields(s)
	ind := []rune{'a', 'o', 'i', 'u', 'e', 'h', 'A', 'O', 'I', 'U', 'E', 'H'}
	for i := 0; i < len(str)-1; i++ {
		if len(str[i+1]) >= 1 {
			if str[i] == "a" && Inslice(rune(str[i+1][0]), ind) {
				str[i] = "an"
			} else if str[i] == "A" && Inslice(rune(str[i+1][0]), ind) {
				str[i] = "An"
			}
		}
	}
	return write(str)
}

// handle one or more ponctuation
func Single_Ponc(s string) string {
	swear := regexp.MustCompile(`\s+([.,;:!?]+)`).ReplaceAllString(s, "$1")
	result := regexp.MustCompile(`([.,;:!?])([[:alnum:]])`).ReplaceAllString(swear, "$1 $2")
	return result
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
		{"( cap ,", "(cap,"},
		{"( up ,", "(up,"},
		{"( low ,", "(low,"},
	}
	for _, char := range TabPonc {
		s = strings.ReplaceAll(s, char.in, char.rep)
	}
	return s
}

// Replace all necessary parentheses
func Rep_parenthese(s string) string {
	Tabpara := []struct {
		in  string
		rep string
	}{
		{"( ", "("},
		{" )", ")"},
	}
	for _, char := range Tabpara {
		s = strings.ReplaceAll(s, char.in, char.rep)
	}
	return s
}

// delete something on slice
func remove(s []string, index int) []string {
	var str []string
	if index >= 0 && index < len(s) {
		for i := 0; i < len(s); i++ {
			if i != index {
				str = append(str, s[i])
			}
		}
		return str
	}
	return s
}

// simple manipulation mean simple key word
func simple_manip(s []string) []string {
	for i, word := range s {
		if word == "(hex)" {
			s = manip_hex(s, i)
			return simple_manip(s)
		} else if word == "(bin)" {
			s = manip_bin(s, i)
			return simple_manip(s)
		} else if word == "(up)" {
			s = manip_up(s, i)
			return simple_manip(s)
		} else if word == "(low)" {
			s = manip_low(s, i)
			return simple_manip(s)
		} else if word == "(cap)" {
			s = manip_cap(s, i)
			return simple_manip(s)
		} else if word == "(low," {
			s = manip_plow(s, i)
			return simple_manip(s)
		} else if word == "(up," {
			s = manip_pup(s, i)
			return simple_manip(s)
		} else if word == "(cap," {
			s = manip_pcap(s, i)
			return simple_manip(s)
		}
	}
	return s
}

// simple manipulation mean simple key word
func manip_hex(s []string, i int) []string {
	ind := []rune{'.', ',', ';', ':', '!', '?', '\'',')','('}
	if i-1 >= 0 && i < len(s) {
		if len(s[i-1]) > 0 {
			if Inslice(rune(s[i-1][0]), ind) {
				s[i], s[i-1] = s[i-1], s[i]
				return manip_hex(s, i-1)
			} else {
				a, err := strconv.ParseInt(s[i-1], 16, 64)
				if err != nil {
					return remove(s, i)
				}
				s[i-1] = strconv.Itoa(int(a))
				s = remove(s, i)
			}
		}
	} else if i == 0 {
		return remove(s, i)
	}
	return s
}
func manip_bin(s []string, i int) []string {
	ind := []rune{'.', ',', ';', ':', '!', '?', '\'',')','('}
	if i-1 >= 0 && i < len(s) {
		if len(s[i-1]) > 0 {
			if Inslice(rune(s[i-1][0]), ind) {
				s[i], s[i-1] = s[i-1], s[i]
				return manip_bin(s, i-1)
			} else {
				a, err := strconv.ParseInt(s[i-1], 2, 64)
				if err != nil {
					return remove(s, i)
				}
				s[i-1] = strconv.Itoa(int(a))
				s = remove(s, i)
			}
		}
	} else if i == 0 {
		return remove(s, i)
	}
	return s
}
func manip_up(s []string, i int) []string {
	ind := []rune{'.', ',', ';', ':', '!', '?', '\'',')','('}
	if i-1 >= 0 && i < len(s) {
		if len(s[i-1])>0 {
			if Inslice(rune(s[i-1][0]), ind) {
				s[i], s[i-1] = s[i-1], s[i]
				return manip_up(s, i-1)
			}else {
				s[i-1] = strings.ToUpper(s[i-1])
				return remove(s,i)
			}
		}
	}
	return remove(s, i)
}
func manip_low(s []string, i int) []string {
	ind := []rune{'.', ',', ';', ':', '!', '?', '\'',')','('}
	if i-1 >= 0 && i < len(s) {
		if len(s[i-1])>0 {
			if Inslice(rune(s[i-1][0]), ind) {
				s[i], s[i-1] = s[i-1], s[i]
				return manip_low(s, i-1)
			}else {
				s[i-1] = strings.ToLower(s[i-1])
				return remove(s,i)
			}
		}
	}
	return remove(s, i)
}
func manip_cap(s []string, i int) []string {
	ind := []rune{'.', ',', ';', ':', '!', '?', '\'',')','('}
	if i-1 >= 0 && i < len(s) {
		if len(s[i-1])>0 {
			if Inslice(rune(s[i-1][0]), ind) {
				s[i], s[i-1] = s[i-1], s[i]
				return manip_cap(s, i-1)
			}else {
				s[i-1] = strings.ToLower(s[i-1])
				s[i-1] = strings.Title(s[i-1])
				return remove(s,i)
			}
		}
	}
	return remove(s, i)
}
func manip_plow(s []string, i int) []string {
ind := []rune{'.', ',', ';', ':', '!', '?', '\'',')','('}
	if i >= 0 && i+1 < len(s)-1 {
		pattern := `[0-9]`
		match, _ := regexp.MatchString(pattern, s[i+1])
		if match && s[i+2] == ")" {
			stop, err := strconv.Atoi(s[i+1])
			if err != nil {
				fmt.Println("Incorrect text")
				os.Exit(0)
			}
			if stop > 0 {
				for a := 1; a <= stop; a++ {
					if i-a >= 0 && i-a < len(s) {
						if len(s[i-a])>0 {
							if Inslice(rune(s[i-a][0]),ind){
								stop++
							}
							s[i-a] = strings.ToLower(s[i-a])
						}
					}
				}
				return remove(remove(remove(s, i), i), i)
			} else if stop == 0 {
				return remove(remove(remove(s, i), i), i)
			} else {
				return remove(remove(remove(s, i), i), i)
			}
		} else if (!match) && s[i+2] == ")" {
			fmt.Println("Incorrect text")
			os.Exit(1)
		} else {
			index := find(s, i, ")")
			if index == i+1 {
				fmt.Println("Incorrect text")
				os.Exit(2)
			} else if index > i+1 {
				a := simple_manip(s[i+1 : index])
				if len(a) == 1 {
					stop, err := strconv.Atoi(a[0])
					if err != nil {
						fmt.Println("Incorrect text")
						os.Exit(0)
					}
					if stop > 0 {
						for a := 1; a <= stop; a++ {
							if i-a >= 0 && i-a < len(s) {
								if len(s[i-a])>0 {
									if Inslice(rune(s[i-a][0]),ind){
										stop++
									}
									s[i-a] = strings.ToLower(s[i-a])
								}
							}
						}
						return multi_remove(s, i, index)
					} else if stop == 0 {
						return multi_remove(s, i, index)
					} else {
						return multi_remove(s, i, index)
					}
				}
			} else {
				fmt.Println("Incorrect text")
				os.Exit(0)
			}
		}
	} else if i+1 < len(s) {
		if s[i+1] == ")" {
			fmt.Println("Incorrect text")
			os.Exit(0)
		}
	}
	fmt.Println("Incorrect text")
	os.Exit(0)
	return s
}
func manip_pup(s []string, i int) []string {
	ind := []rune{'.', ',', ';', ':', '!', '?', '\'',')','('}
	if i >= 0 && i+1 < len(s)-1 {
		pattern := `[0-9]`
		match, _ := regexp.MatchString(pattern, s[i+1])
		if match && s[i+2] == ")" {
			stop, err := strconv.Atoi(s[i+1])
			if err != nil {
				fmt.Println("Incorrect text")
				os.Exit(0)
			}
			if stop > 0 {
				for a := 1; a <= stop; a++ {
					if i-a >= 0 && i-a < len(s) {
						if len(s[i-a])>0 {
							if Inslice(rune(s[i-a][0]),ind){
								stop++
							}
							s[i-a] = strings.ToUpper(s[i-a])
						}
					}
				}
				return remove(remove(remove(s, i), i), i)
			} else if stop == 0 {
				return remove(remove(remove(s, i), i), i)
			} else {
				return remove(remove(remove(s, i), i), i)
			}
		} else if !(match) && s[i+2] == ")" {
			fmt.Println("Incorrect text")
			os.Exit(1)
		} else {
			index := find(s, i, ")")
			if index == i+1 {
				fmt.Println("Incorrect text")
				os.Exit(2)
			} else if index > i+1 {
				a := simple_manip(s[i+1 : index])
				if len(a) == 1 {
					stop, err := strconv.Atoi(a[0])
					if err != nil {
						fmt.Println("Incorrect text")
						os.Exit(0)
					}
					if stop > 0 {
						for a := 1; a <= stop; a++ {
							if i-a >= 0 && i-a < len(s) {
								if len(s[i-a])>0 {
									if Inslice(rune(s[i-a][0]),ind){
										stop++
									}
									s[i-a] = strings.ToUpper(s[i-a])
								}
							}
						}
						return multi_remove(s, i, index)
					} else if stop == 0 {
						return multi_remove(s, i, index)
					} else {
						return multi_remove(s, i, index)
					}
				}
			} else {
				fmt.Println("Incorrect text")
				os.Exit(0)
			}
		}
	} else if i+1 == len(s)-1 {
		if s[i+1] == ")" {
			fmt.Println("Incorrect text")
			os.Exit(0)
		}
	}
	fmt.Println("Incorrect text")
	os.Exit(0)
	return s
}
func manip_pcap(s []string, i int) []string {
	ind := []rune{'.', ',', ';', ':', '!', '?', '\'',')','('}
	if i >= 0 && i+1 < len(s)-1 {
		pattern := `[0-9]`
		match, _ := regexp.MatchString(pattern, s[i+1])
		if match && s[i+2] == ")" {
			stop, err := strconv.Atoi(s[i+1])
			if err != nil {
				fmt.Println("Incorrect text")
				os.Exit(0)
			}
			if stop > 0 {
				for a := 1; a <= stop; a++ {
					if i-a >= 0 && i-a < len(s) {
						if len(s[i-a])>0 {
							if Inslice(rune(s[i-a][0]),ind){
								stop++
							}
							s[i-a] = strings.Title(strings.ToLower(s[i-a]))
						}
					}
				}
				return remove(remove(remove(s, i), i), i)
			} else if stop == 0 {
				return remove(remove(remove(s, i), i), i)
			} else {
				return remove(remove(remove(s, i), i), i)
			}
		} else if !(match) && s[i+2] == ")" {
			fmt.Println("Incorrect text")
			os.Exit(1)
		} else {
			index := find(s, i, ")")
			if index == i+1 {
				fmt.Println("Incorrect text")
				os.Exit(2)
			} else if index > i+1 {
				a := simple_manip(s[i+1 : index])
				if len(a) == 1 {
					stop, err := strconv.Atoi(a[0])
					if err != nil {
						fmt.Println("Incorrect text")
						os.Exit(0)
					}
					if stop > 0 {
						for a := 1; a <= stop; a++ {
							if i-a >= 0 && i-a < len(s) {
								if len(s[i-a])>0 {
									if Inslice(rune(s[i-a][0]),ind){
										stop++
									}
									s[i-a] = strings.Title(strings.ToLower(s[i-a]))
								}
							}
						}
						return multi_remove(s, i, index)
					} else if stop == 0 {
						return multi_remove(s, i, index)
					} else {
						return multi_remove(s, i, index)
					}
				}
			} else {
				fmt.Println("Incorrect text")
				os.Exit(0)
			}
		}
	} else if i+1 < len(s) {
		if s[i+1] == ")" {
			fmt.Println("Incorrect text")
			os.Exit(0)
		}
	}
	fmt.Println("Incorrect text")
	os.Exit(0)
	return s
}
func find(s []string, debt int, str string) int {
	for j := debt; j < len(s); j++ {
		if j >= 0 && j < len(s) {
			if s[j] == str {
				return j
			}
		}
	}
	return -1
}
func multi_remove(s []string, indx int, time int) []string {
	for i := 0; i < time; i++ {
		if indx >= 0 && indx < len(s) {
			s = remove(s, indx)
		}
	}
	return s
}
func main() {
	if len(os.Args) == 3 {
		content, _ := os.ReadFile(os.Args[1])

		text := string(content)

		text = Go_reloaded(text)

		text = Go_reloaded(text)

		err := ioutil.WriteFile(os.Args[2], []byte(text), 0644)
		if err != nil {
			return
		}
	}
}
func Go_reloaded(s string) string {
	s = parenthese(s)
	str := strings.Fields(s)
	str = simple_manip(str)
	s = write(str)
	return Ponc(s)
}
