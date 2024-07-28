package main

import (
	"fmt"
	"net/http"
)

func channelsPractice() {
	links := []string{

		"https://golang.org/",
		"https://www.google.co.in/",
		"https://stackoverflow.com/",
		"https://amazon.com/",
	}
	go fmt.Println("something like anything")
	c := make(chan string)
	for _, link := range links {
		fmt.Println("link")
		go checkStatusOfWebsite(link, c)

	}
	for l := range c {
		fmt.Println(l)
	}
	// for link := range c {
	// 	go func(link string) {
	// 		time.Sleep(5 * time.Second)

	// 		checkStatusOfWebsite(link, c)
	// 	}(link)
	// }
}

func checkStatusOfWebsite(link string, val chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " Might down ")
		val <- link
	}
	val <- link
	fmt.Println(link, "is up")
}
