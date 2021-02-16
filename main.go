package main

import (
	"fmt"
	"net/http"
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
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

func checkStatus(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down")
		c <- "Might be down I think"
	}

	fmt.Println(url, "is ok!")
	c <- "Yep its up"
}
