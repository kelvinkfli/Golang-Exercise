/*
Input: a string "s"
Output: boolean? whether or not "s" is an isogram
Exceptions/Assumptions:  string is of length > 0
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	isIsogram("Hello")
}

func isIsogram(s string) {
	letters := strings.Split(s, "")
	myMap := make(map[string]int)
	for _, v := range letters {
		myMap[v] += 1
	}
	fmt.Println(myMap)
}
