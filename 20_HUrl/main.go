package main

import (
	"fmt"
	"net/url"
)

var urls string = "https://lco.dev:3000/learn?coursename=reactjs&price=free&status=flase"

func main() {
	fmt.Println(urls)

	// parse url
	u, _ := url.Parse(urls)
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Port())
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)

	// get query params
	m, _ := url.ParseQuery(u.RawQuery)

	for k, v := range m {
		fmt.Println(k, ":", v)
	}

	partOfUrl := &url.URL{
		Scheme:   u.Scheme,
		Host:     u.Host,
		Path:     u.Path,
		RawQuery: u.RawQuery,
	}

	anotherUrl := partOfUrl.String()
	fmt.Println(anotherUrl)
}
