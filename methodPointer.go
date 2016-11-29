package main

import "fmt"

//Why use a pointer receiver???
//1. modify the value that the receiver points to
//2. avoid copying the value on each method call

//create a struct type
type myStruct struct {
	x, y int
}

//define methods for struct
func (c *myStruct) foo(f int) {
	c.x = c.x * f
	c.y = c.y * f
}

func main() {
	//instantiate the struct
	first := myStruct{5, 10}
	fmt.Println(first)
	first.foo(5)
	fmt.Println(first)
}
