/*
Goal: write a function that determines the type of triangle
Input: triangle...dimensions?
Output: string, defining the type of triangle
Conditions: none? dimensions must be non-zero...
*/

package main

import "fmt"

type triangleDimensions struct {
	side1, side2, side3 int
}

func (t triangleDimensions) defineTriangle() string {
	measureList := []int{t.side1, t.side2, t.side3}
	measureTrack := make(map[int]int)

	for _, v := range measureList {
		measureTrack[v] = measureTrack[v] + 1
	}

	for i, v := range measureTrack {
		if i == 0 {
			return "triangle does not exist"
		} else if v == 3 {
			return "the triangle is equilateral"
		} else if v == 2 {
			return "the triangle is isosceles"
		}
	}
	return "The triangle is scalene"
}

func main() {
	triangleOne := triangleDimensions{3, 3, 3}
	fmt.Println(triangleOne.defineTriangle())
}
