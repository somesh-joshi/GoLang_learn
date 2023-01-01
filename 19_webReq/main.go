package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	defer res.Body.Close()
}

// Post Request

//payload

// reqbody := strings.NewReader(`{"name":"lco","age":23}`)
// res, err := http.Post(url, "application/json", reqbody)
// if err != nil {
// 	panic(err)
// }

// data, err := ioutil.ReadAll(res.Body)
// if err != nil {
// 	panic(err)
// }
// fmt.Println(string(data))
