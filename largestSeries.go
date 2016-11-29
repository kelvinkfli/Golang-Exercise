/*

Input: a series of numbers (string of digits), and a length n
Output: the largest product of substrings n within the series of numbers
Exceptions/Assumptions:
    - series cannot be < 1
    - length n cannot be greater than the len(series)

*/

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	getLargestProduct("1027839564", 3)
}

func getLargestProduct(s string, n int) {

	numStrings := strings.Split(s, "")
	numList := []int{}

	//convert strings in slice to integers
	for _, v := range numStrings {
		j, _ := strconv.Atoi(v)
		numList = append(numList, j)
	}

	productList := []int{}

	//generate all the subsets
	for i, _ := range numList {
		if i+n <= len(numList) {
			numSlice := numList[i : i+n]
			product := 1
			for i = 0; i < n; i++ {
				product = product * numSlice[i]
			}
			productList = append(productList, product)
		}
	}
	sort.Ints(productList)
	largestProduct := productList[len(productList)-1]
	fmt.Println("The largest product of your number is:", largestProduct)
}
