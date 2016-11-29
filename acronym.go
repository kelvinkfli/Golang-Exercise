package main

import (
	"fmt"
	"strings"
)

/*
Goal: convert long phrases into acronyms
Input: a string of arbitrary length
Output: a string that is an acronym
Conditions: none?
*/

func makeAcro(phrase string) string {
	words := strings.Split(phrase, " ")
	acronym := ""
	for _, v := range words {
		firstLetter := v[0:1]
		acronym += firstLetter
	}

	return acronym
}

func main() {
	message := makeAcro("Hello World today")
	fmt.Println(strings.ToUpper(message))
}
