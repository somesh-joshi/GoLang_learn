package main

import "fmt"

func main() {

	var fruitarray [4]string

	fruitarray[0] = "Apple"
	fruitarray[1] = "Orange"
	fruitarray[3] = "Grape"

	fmt.Println(fruitarray) // [Apple Orange  Grape] it add space for empty index
	fmt.Println(len(fruitarray))

	var vegarray = [3]string{"Potato", "Tomato", "Onion"}
	fmt.Println(vegarray) // [Potato Tomato Onion]
	fmt.Println(len(vegarray))	
}