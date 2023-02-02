package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/somesh-joshi/MovieProject/db"
	"github.com/somesh-joshi/MovieProject/models/person"
)

var coll = db.Db.Collection("file")

func main() {
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	B := strings.Split(string(data), ",")
	datainterface := make([]interface{}, 0)
	for i := 0; i < len(B); i += 4 {
		date, _ := strconv.Atoi(B[i+2])
		person := person.Person{
			Title: B[i+0],
			Name:  B[i+1],
			Age:   date,
			DoB:   B[i+3],
		}
		datainterface = append(datainterface, person)
	}
	result, err := coll.InsertMany(context.TODO(), datainterface)
	if err != nil {
		panic(err)
	}
	for _, v := range result.InsertedIDs {
		fmt.Println(v)
	}
}
