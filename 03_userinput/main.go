package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	WellcomeMessage := "Welcome to GoLang"
	fmt.Println(WellcomeMessage)

	
	fmt.Printf("Enter your number: ")

	// comma ok // err ok
	reader := bufio.NewReader(os.Stdin) // read from console // u can declare it outside the main function
	input, err := reader.ReadString('\n')  // it is a Promise in GoLang
	fmt.Println(input)
	fmt.Printf("Value of input type is %T \n", input)
	fmt.Println(err)





}