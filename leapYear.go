package main

import "fmt"

/*
Write a function that accepts a year, and report if its a leap year
*/

func isLeap(year int) bool {
	x := false
	if year%4 == 0 {
		x = true
	}
	return x
}

func main() {
	isIt := isLeap(2301)
	if isIt == true {
		fmt.Println("The year you chose was a leap year")
	} else {
		fmt.Println("The year you chose was not a leap year")
	}
}
