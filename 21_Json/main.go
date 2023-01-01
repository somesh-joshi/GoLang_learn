package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name  string
	price int32
	tags  []string
}

func main() {
	EncodeJson()
}

func EncodeJson() {
	loc := []course{
		{Name: "Go", price: 100, tags: []string{"Go", "Golang"}},
		{Name: "Python", price: 200, tags: []string{"Python", "Python3"}},
		{Name: "Java", price: 300, tags: nil},
	}

	finalJson, err := json.MarshalIndent(loc, "", "\t")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(finalJson))

}
