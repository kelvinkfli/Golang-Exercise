/*
Input: a function that takes in two parameters: a max number / any number of numbers to find multiples for
Output: sum of all those multiples up to the max number
Exceptions: max number cannot be 0, variadic params cannot be 0
*/

package main

import "fmt"

func main() {
	fmt.Println(getNums(15, 3, 5))
}

func getNums(max int, factor1 int, factor2 int) int {
	// numList := []int{}
	total := 0
	for i := max - 1; i > 0; i-- {
		if i%factor1 == 0 || i%factor2 == 0 {
			total += i
			fmt.Println(i)
		}
	}
	return total
}
