package main

import (
	"fmt"
	"strings"
)

func main() {
	s := findFirstStringInBracket(`find (this) string in bracket`)
	fmt.Println(s)
}
func findFirstStringInBracket(str string) string {
	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound >= 0 {
		indexClosingBracketFound := strings.Index(str[indexFirstBracketFound+1:], ")")
		if indexClosingBracketFound >= 0 {
			return str[indexFirstBracketFound+1 : indexFirstBracketFound+indexClosingBracketFound+1]
		}
	}
	return ""
}
