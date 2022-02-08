package main

import (
	"fmt"
	"time"
)

func main() {
	channel_1 := make(chan string) // gochannel
	people := [4]string{"git", "github", "Visual Studio Code", "atom"}
	for _, person := range people {
		go isAwesome(person, channel_1) // goroutine
	}
	for i := 0; i < len(people); i++ {
		fmt.Print("Waiting for ", i)
		fmt.Println(<-channel_1) // blocking operation : receiving channel messages
	}
}


func isAwesome(person string, channel_1 chan string) {
	time.Sleep(time.Second * 5)
	channel_1 <- person + " is Awesome!" // sending channel message(info: int, string, boolean, whatever..)
}