/*
Input: a string of ints and an integer n
Output: continuous substrings of length n in that string
Exclusions: the digits themselves do not need to numerically consecutive
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	getSubString("1337", 4)
	getSubString("Go is pretty cool", 3)
	getSubString("0983098234212353647", 4)
}

func getSubString(s string, n int) {
	stringSlice := strings.Split(s, "")
	if n > len(stringSlice) {
		fmt.Println("integer value larger than actual string, cannot make substrings")
	}
	for i, _ := range stringSlice {
		if i+n <= len(stringSlice) {
			subSlice := stringSlice[i : i+n]
			fmt.Println(subSlice)
		}
	}
}
