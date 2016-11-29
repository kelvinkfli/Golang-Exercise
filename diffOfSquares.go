/*
Goal: calculate the difference between the sum of the squares
      and the square of the sum

Input: n int
Output: final int

*/

package main

import "fmt"

func getNumList(n int) []int {
	numSlice := []int{}
	for i := 1; i <= n; i++ {
		numSlice = append(numSlice, i)
	}
	return numSlice
}

func diffOfSquares(nums []int) int {
	squareSum := 0
	sumSquare := 0
	for _, v := range nums {
		squareSum += v
		sumSquare += v * v
	}
	squareSum = squareSum * squareSum
	return squareSum - sumSquare
}

func main() {
	number := getNumList(5)
	fmt.Println(diffOfSquares(number))
}
