package main

import "fmt"

func main() {
	// it is alternative of classes
	// there is no inheritance, no super, no parent, no child

	somesh := Person{
		Name:   "Somesh",
		Age:    25,
		Email:  "abc@gmail.com",
		Status: true,
	}

	somesh.greet()
}

type Person struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (p *Person) greet() { // U can use pointer or value to ref the struct
	fmt.Println("Hello", p.Name)
}

// to make method exportable, we need to capitalize the first letter of the method name
// to make struct exportable, we need to capitalize the first letter of the struct name
// to make variable exportable, we need to capitalize the first letter of the variable name
