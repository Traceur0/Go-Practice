package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url string
	status string
}

var errRequestFailed = errors.New("request failed")

func main() {
	results := make(map[string]string)
	channel := make(chan requestResult)
	urls := []string{
		"https://www.jsafklajweioj.com",
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://soundcloud.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://academy.nomadcoders.co",
	}
	for _, url := range urls {
		go hitURL(url, channel)
	}
	for i:=0; i < len(urls); i++ {
		result := <-channel
		results[result.url] = result.status 
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, channel chan<- requestResult) {
	resp, err := http.Get(url)
	status := "Status : Online"
	if err != nil || resp.StatusCode >= 400 {
		status = "Status : Offline"
	}
	channel <- requestResult{url: url, status: status}
}