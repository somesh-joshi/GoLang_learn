package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	fmt.Println("World")

	fmt.Println("///////")

	func() {
		defer fmt.Println("Hello")
		defer fmt.Println("World")
		defer fmt.Println("!") 
		fmt.Println("Hello")
	}()
}

// defer is used to delay the execution of a function
// defer is used to execute a function after the current function is completed
// first in first out

