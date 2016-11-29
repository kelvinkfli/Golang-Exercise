package main

import "fmt"
import "time"

func main() {
	myChannel := make(chan int)

	go printNum(myChannel)

	for nums := range myChannel {
		fmt.Println(nums)
	}
}

func printNum(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(250 * time.Millisecond)
	}
}
