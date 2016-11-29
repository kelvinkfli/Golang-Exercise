package main

import "fmt"

//create a struct
type laptop struct {
	size, weight int
}

//create another struct, that has another struct embedded within it
type macBook struct {
	laptop
	brand  string
	screen int
}

//create a method for the laptop struct
func (property *laptop) randomCpuCalc() int {
	total := property.size + property.weight
	return total
}

func main() {
	//instantiate a struct
	laptopOne := laptop{15, 300}
	//use the struct method
	fmt.Println(laptopOne.randomCpuCalc())
	//instantiate our new struct
	MacBookOne := macBook{laptop: laptopOne, brand: "Apple", screen: 13}
	fmt.Println(MacBookOne)

	fmt.Println(MacBookOne.randomCpuCalc())
}
