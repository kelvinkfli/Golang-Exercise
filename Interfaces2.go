package main

import "fmt"

type triangle struct {
	length float64
	height float64
}

func (t *triangle) area() float64 {
	return t.length * t.height / 2
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return c.radius * c.radius
}

type isShape interface {
	area() float64
}

//now all things that have area methods are considered SHAPES
//create a function that takes a shape, and returns the area of that shape

func findArea(s isShape) {
	area := s.area()
	fmt.Println(area)
}

func main() {
	//instantiate some structs
	tri1 := triangle{length: 41, height: 12}
	circ1 := circle{radius: 17}
	findArea(&tri1)
	findArea(&circ1)

}
