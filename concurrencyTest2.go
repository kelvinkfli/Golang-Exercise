package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	go sayHello(channel1)
	go sayHi(channel1)
	printMessage(channel1)
}

//write function that takes in a channel and puts values into it
func sayHi(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "Hi"
	}
}

func sayHello(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "Hello"
	}
}

func printMessage(ch chan string) {
	for i := 0; i < 20; i++ {
		message := <-ch
		fmt.Println(message)
		time.Sleep(time.Second * 1)
	}
}
