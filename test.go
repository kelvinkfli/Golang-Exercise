package main

import (
	"fmt"
	"time"
)

func getMessagesChannel(msg string, delay time.Duration) <-chan string {
	c := make(chan string)
	go func() {
		for i := 1; i <= 3; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			// Wait before sending next message
			time.Sleep(time.Millisecond * delay)
		}
	}()
	return c
}

func main() {
	c1 := getMessagesChannel("first", 1000)
	c2 := getMessagesChannel("second", 500)
	c3 := getMessagesChannel("third", 10)

	for i := 1; i <= 9; i++ {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		case msg := <-c3:
			fmt.Println(msg)
		}
	}
}
