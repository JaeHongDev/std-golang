package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request fail")

func main() {
	//results := make(map[string]string)
	c := make(chan result)

	url := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range url {
		fmt.Printf("Checking url:%s\n", url)
		go hitUrl(url, c)
	}

}

func hitUrl(url string, c chan<- result) {
	response, err := http.Get(url)
	status := "OK"
	if err != nil || response.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{url: url, status: status}
}
