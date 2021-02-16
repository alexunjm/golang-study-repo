package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://azure.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://facebook.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkStatus(link, c)
	}
	for u := range c {
		// function literal
		go func(url string, c chan string) {
			duration := 6 * time.Second
			time.Sleep(duration)
			checkStatus(url, c)
		}(u, c)
	}

}

func checkStatus(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down")
	} else {
		fmt.Println(url, "is ok!")
	}

	c <- url
}
