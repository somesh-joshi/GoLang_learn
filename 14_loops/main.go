package main

import "fmt"

func main() {

	var days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	//for i := 0; i < 5; i++ {
	//	fmt.Println(i)
	//}

	//for i := range days {
	//	fmt.Println(days[i])
	//}

	for _, value := range days {
		fmt.Println(value)
	}

	rvalue := 0

	for rvalue < 5 {
		fmt.Println(rvalue)
		rvalue++
	}
}
