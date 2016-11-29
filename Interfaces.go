package main

import "fmt"

/*
-Interfaces are automatically implemented whenever a type satisfies that interface
-satisfaction is determined by whether or not that type has defined all the methods in the interface

*/

//Create our interface, anything with the defined "Speak() method will implement this"
type Animal interface {
	speak() string
}

//Create some types that satisfy "Animal"
type Dog struct{}

func (d *Dog) speak() string {
	return "Woof"
}

type Cat struct{}

func (c *Cat) speak() string {
	return "Meow"
}

type Programmer struct{}

func (p *Programmer) speak() string {
	return "Go's pretty cool"
}

//Make a slice of animals
func main() {
	animals := []Animal{&Dog{}, &Cat{}, &Programmer{}}
	for _, name := range animals {
		fmt.Println(name)
		fmt.Println(name.speak())
	}
}
