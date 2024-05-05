package main

import (
	"errors"
	"fmt"
	"time"
)

var (
	errRequestFailed = errors.New("Request Failed")
)

func main() {
	go coolCount("sashin")
	go coolCount("nico")
	time.Sleep(time.Second * 5)
}

func coolCount(preson string) {
	for i := 0; i < 10; i++ {
		fmt.Println(preson, "is cool!", i)
		time.Sleep(time.Second)
	}
}

// func main() {

// 	urls := []string{
// 		"https://www.airbnb.com",
// 		"https://www.google.com",
// 		"https://www.amazon.com",
// 		"https://www.reddit.com",
// 		"https://www.google.com",
// 		"https://soundcloud.com",
// 		"https://www.facebook.com",
// 		"https://www.instagram.com",
// 	}

// 	var results = make(map[string]string)
// 	for _, url := range urls {
// 		result := "OK"
// 		err := hitURL(url)
// 		if err != nil {
// 			result = "FAILED"
// 		}
// 		results[url] = result
// 	}
// 	for url, result := range results {
// 		fmt.Println(url, result)
// 	}
// }

// func hitURL(url string) error {
// 	fmt.Println("Checking: ", url)
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		fmt.Println(err)
// 		return errRequestFailed
// 	}
// 	return nil
// }
