package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	links := []string {
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.linkedin.com",
		"https://www.pinterest.com",
		"https://netflix.com",
	}

	c := make(chan bool)

	for num := 0; num < 5; num++ {
		for _, link := range links {
			go checkLink(link, c)
		}
	}

	for num := 0; num < 5; num++ {
		for i := 0; i < len(links); i++ {
			<-c
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

func checkLink(link string, c chan bool) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- false
		return
	}
	fmt.Println(link, "is up!")
	c <- true
}