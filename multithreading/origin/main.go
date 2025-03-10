package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	links := []string{
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

	for num := 0; num < 5; num++ {
		for _, link := range links {
			checkLink(link)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}