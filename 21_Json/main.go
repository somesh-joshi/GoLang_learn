package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name  string
	Price int32
	Tags  []string
}

func main() {
	EncodeJson()
}

func EncodeJson() {
	loc := []course{
		{Name: "Go", Price: 100, Tags: []string{"beginner", "intermediate"}},
		{Name: "Python", Price: 200, Tags: []string{"beginner", "intermediate"}},
		{Name: "Java", Price: 300, Tags: nil},
	}

	finalJson, err := json.MarshalIndent(loc, "", "\t")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(finalJson))

}
