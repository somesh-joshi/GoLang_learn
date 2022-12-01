package main

import (
	"fmt"
	"sort"
)

func main() {
	
	var fruitarray = []string{"apple", "grape", "banana", "melon"}
	fmt.Printf("%T\n", fruitarray) // []string

	fruitarray = append(fruitarray, "orange")
	fmt.Println(fruitarray) // [apple grape banana melon orange]

	fruitarray = fruitarray[1:3]
	fmt.Println(fruitarray) // [grape banana]

	HighScore := make([]int, 4)
	HighScore[0] = 100
	HighScore[1] = 200
	HighScore[2] = 300
	HighScore[3] = 400

	HighScore = append(HighScore, 500, 600, 700, 800)

	fmt.Println(HighScore) 

	sort.Ints(HighScore)
	fmt.Println(HighScore)
	fmt.Println(sort.IntsAreSorted(HighScore)) // true

	// remove a value from slice based on index\

	HighScore = append(HighScore[:2], HighScore[3:]...)
	fmt.Println(HighScore) // [100 200 400 500 600 700 800]

}