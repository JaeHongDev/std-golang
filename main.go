package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request fail")

func main() {
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
		hitUrl(url)
	}

}

func hitUrl(url string) error {
	response, err := http.Get(url)
	if err != nil || response.StatusCode >= 400 {
		return errRequestFailed
	}
	fmt.Println(response)
	return nil

}
