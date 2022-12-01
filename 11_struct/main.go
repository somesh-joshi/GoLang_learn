package main

import "fmt"

func main() {
	// it is alternative of classes
	// there is no inheritance, no super, no parent, no child

	somesh := Person{
		Name: "Somesh",
		Age:  25,
		Email: "abc@gmail.com",
		Status: true,
	}

	fmt.Println(somesh)
	fmt.Println(somesh.Name)
}


type Person struct {
	Name   string
	Age    int
	Email  string
	Status bool
}
