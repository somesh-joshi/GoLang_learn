package main

import "fmt"

func main() {
	geeting()
	fmt.Println(add(2, 5))
	fmt.Println(proadd(1, 3, 4))
	a, b := evencheck(1)
	fmt.Println(a, b)
	func() {
		fmt.Println("I am a anonymous function")
	}()
}

func geeting() {
	fmt.Println("Hello World")
}

func add(x int, y int) int {
	return x + y
}

func proadd(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func evencheck(x int) (int, string) {
	if x%2 == 0 {
		return x, "even"
	}
	return x, "odd"
}
