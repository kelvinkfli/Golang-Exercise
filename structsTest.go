package main

import "fmt"

//create a struct (a custom type with fields)
type dimensions struct {
	width, height int
}

//define a method on a struct
//syntax: func (receiver name_receiver type) funcName() {}
func (cm *dimensions) area() int {
	return cm.width * cm.height
}

func main() {
	//instantiate a our new dimensions type
	square1 := dimensions{width: 10, height: 10}
	//call our method using the instance
	foo := square1.area()
	fmt.Println(foo)

	//create a new instance of our type
	rectangleOne := dimensions{width: 4, height: 4}
	bar := rectangleOne.area()
	fmt.Println(bar)
}
