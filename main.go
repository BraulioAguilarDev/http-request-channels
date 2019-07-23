package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func MakePetition(url string, ch chan<- string) {
	start := time.Now()
	res, _ := http.Get(url)

	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(res.Body)
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}

func main() {
	urls := []string{"https://reqres.in/api/users/1", "https://reqres.in/api/users/2", "https://reqres.in/api/users/3"}
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go MakePetition(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())
}
