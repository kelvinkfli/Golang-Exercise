/*
Input: n int
Output: a string based on the factors of n

Conditions:
-if n has 3 as a factor: pling
-if n has 5 as a factor: plang
-if n has 7 as a factor: plong

Steps:
1. determine factors of n
2. iterate through factors to see if conditions are met
3. return string
*/
package main

import "fmt"

func findFactor(num int) []int {
	//create splice to store factors
	factorList := []int{}

	//loop to find factors
	for i := 1; i <= num; i++ {
		//if i evenly divides into num, its a factor
		if num%i == 0 {
			factorList = append(factorList, i)
		}
	}
	return factorList
}

func checkSound(nums []int) string {
	//declare message to be sent
	message := ""

	//modify message based on factor
	for _, v := range nums {
		if v == 3 {
			message += "Pling"
		} else if v == 5 {
			message += "Plang"
		} else if v == 7 {
			message += "Plong"
		}
	}

	//if no factors found
	if message == "" {
		message += "No rain, much sad"
	}

	return message
}

func main() {
	factors := findFactor(143)
	fmt.Println(checkSound(factors))
}
