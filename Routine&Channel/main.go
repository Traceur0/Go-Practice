package main

import (
	"fmt"
	"time"
)

func main() {
	channel_1 := make(chan string) // gochannel
	people := [2]string{"git", "github"}
	for _, person := range people {
		go isAwesome(person, channel_1) // goroutine
	}
	fmt.Println("Waiting for channel data")
	fmt.Println("Received data: ", <-channel_1)
	fmt.Println("Received data: ", <-channel_1)
}


func isAwesome(person string, channel_1 chan string) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	channel_1 <- person + " is Awesome!"
}