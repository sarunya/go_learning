package main

import (
	"fmt"
	"net/http"
)

func main() {
	naiveStatusChecker()
}

func naiveStatusChecker() {
	links := []string{"http://www.google.com", "http://www.yahoo.com", "http://www.golang.org", "http://www.facebook.com"}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	//infinite
	for resp := range c {
		fmt.Println(resp)
	}

	//repeating routines
	// for resp := range c {
	// 	// resp <- c
	// 	fmt.Println(resp)
	// 	resp = strings.Split(resp, " ")[0]
	// 	// go checkLink(resp, c)
	// 	// function literal
	// 	go func(link string) {
	// 		time.Sleep(time.Second * 5)
	// 		checkLink(link, c)
	// 	}(resp)
	// }
}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil || resp.StatusCode != 200 {
		// fmt.Println(link, "is not available")
		c <- link + " is not available"
	} else {
		// fmt.Println(link, "is available")
		c <- link + " is available"
	}
}
