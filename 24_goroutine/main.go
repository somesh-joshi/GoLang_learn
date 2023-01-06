package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var endpoints = []string{}
var mut sync.Mutex

func main() {
	var endpoints = []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.youtube.com",
		"https://www.yahoo.com",
	}

	for _, endpoint := range endpoints {
		go getStatusCode(endpoint)
		wg.Add(1)
	}

	wg.Wait()
}

//func greeting(s string) {
//	for i := 0; i < 6; i++ {
//		time.Sleep(2 * time.Millisecond)
//		fmt.println(s)
//	}
//}

func getStatusCode(endpoint string) {
	defer wg.Done()
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error")
	}
	mut.Lock()
	endpoints = append(endpoints, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for	%s\n", resp.StatusCode, endpoint)
}
