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
	fmt.Println(isIsogram("golang is not an isogram"))
}

func isIsogram(s string) string {
	letters := strings.Split(s, "")
	myMap := make(map[string]int)
	for _, v := range letters {
		myMap[v] += 1
	}
	for _, v := range myMap {
		if v > 1 {
			return "This word/phrase is not an isogram"
		}
	}
	return "This word/phrase is an isogram"
}
