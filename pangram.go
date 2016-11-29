/*
Goal: write a function that determines if a string is a pangram
Input: s string
Output: bool? "Is a pangram" vs. "Isn't a pangram"
Conditions: none
*/

package main

import (
	"fmt"
	"strings"
)

func isPangram(s string) string {
	alphabet := "abcdefghijklmopqrstuvwxyz"
	alphalist := strings.Split(alphabet, "")
	alphamap := make(map[string]int)

	phrase := strings.Split(s, "")

	for _, v := range alphalist {
		alphamap[v] = 0
	}

	for _, c := range phrase {
		alphamap[c] = 1
	}

	for _, v := range alphamap {
		if v == 0 {
			return "This sentence is not a pangram"
		}
	}
	return "This sentence is a pangram"
}

func main() {
	fmt.Println(isPangram("The quick brown fox jumps over the lazy dog"))
	fmt.Println(isPangram("This is likely not a pangram"))
}
