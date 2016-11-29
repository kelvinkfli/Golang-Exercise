/*
Input: a series of numbers (int)
Output: boolean determining whether or not given number passes luhn formula
Conditions:
    - counting from the rightmost digit, double the value of every second digit
    - for any value that becomes 10 or more, subtract 9 from the result
    - add all those digits together:
        - total must end in 10 = valid according to luhn formula

Exceptions/Assumptions:
    - number (int) must be greater than zero? maybe? 0 will technically pass

Approach: use modulo to iterate through every other number in the array
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(checkLuhn("2323200577663554"))
}

func checkLuhn(numbers string) string {
	numberStringList := strings.Split(numbers, "")

	//empty slice to store original values
	numberList := []int{}

	//convert string of numbers to integers for each one
	for _, v := range numberStringList {
		j, _ := strconv.Atoi(v)
		numberList = append(numberList, j)
	}

	//variable to store the sum of converted digits
	luhnSum := 0

	//for each int (starting from end), every other digit gets tested
	for i := len(numberList) - 1; i >= 0; i-- {
		if i%2 == 0 {
			evenValue := numberList[i] * 2
			if evenValue >= 10 {
				evenValue -= 9
			}
			//add the converted value to the sum
			luhnSum += evenValue
		} else {
			//add the original value to the sum
			luhnSum += numberList[i]
		}
	}

	if luhnSum%10 == 0 {
		return "The supplied number passed the luhn test"
	}

	return "The supplied number does not pass the luhn test"
}
