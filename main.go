package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls []string

func init() {
	for i := 0; i <= 30; i++ {
		urls = append(urls, fmt.Sprintf("https://reqres.in/api/users/%v", i))
	}

}

type HttpResponse struct {
	Url      string
	Response *http.Response
	Err      error
}

func AsyncHttpGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse)
	responses := []*HttpResponse{}
	client := http.Client{}

	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching ==> %s \n", url)

			resp, err := client.Get(url)
			ch <- &HttpResponse{url, resp, err}

			if err != nil && resp != nil && resp.StatusCode == http.StatusOK {
				resp.Body.Close()
			}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched \n", r.Url)
			if r.Err != nil {
				fmt.Println("xxx With an error xxx", r.Err)
			}

			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Println(".")
		}
	}

	return responses
}

func main() {
	results := AsyncHttpGets(urls)

	for key, result := range results {
		if result != nil && result.Response != nil {
			fmt.Printf("%v ==> %s\n", key, result.Response.Status)
		}
	}
}
