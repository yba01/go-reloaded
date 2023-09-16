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
		result_cont := manip(strings.Fields(cont))
		final_reslt := Affichage(result_cont)
		err := ioutil.WriteFile(os.Args[2], []byte(final_reslt), 0644)
		if err != nil {
			os.Exit(0)
		}
	}
}

// 1.

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

// format text
func correct(s string) string {
	co1 := ponct(s)
	co2 := parenthese(co1)
	corrected := remove_space(co2)
	return corrected
}

//2.

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
	tabStruct := []struct {
		name string
		do   func(string) string
	}{
		{"(hex)", hex},
		{"(bin)", bin},
		{"(up)", up},
		{"(low)", low},
		{"(cap)", cap},
	}
	for index, word := range slice {
		for _, mot := range tabStruct {
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
func manip_complex(s []string) []string {
	Tablestruct:=[]struct{
		name string
		do func(string) string
	}{
		{"(low,",low},
		{"(up,",up},
		{"(cap,",cap},
	}
	for i,word:=s {
		for _,mot:=range Tablestruct {
			if word==mot.name {
				if i<len(s)-1 {
					pattern:=`^[0-9]\)$`
					match,_:=regexp.MatchString(pattern,s[i+1])
					if match {
						stop:=strings.Atoi(s[i+1][:len(s[i+1]-1)])
						s=iter_funct(s,mot.do,stop,i)
					}
				}
			}
		}
	}
}
func iter_funct(str []string,faire func(string) string,num int, index int)[]string {
	for i:=index-1;i>=index-num;i--{
		if i>=0 {
			str[i]=faire(s[i])
		}
	}
	s=remove(str,index)
	s=remove(str,index)
	return str
}
func index_end(str []string,debut int) int{
	var result int
	for i:=debut;i<len(str);i++ {
		match,_:=regexp.MatchString(`\)\)$`,str[i])
		if match {
			result:=i
			break
		}
	}
	return result
}