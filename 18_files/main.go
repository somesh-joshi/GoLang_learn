package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	constent := "Hello World"

	file, err := os.Create("test.txt")

	if err != nil {
		panic(err) //it stop the program and print the error
	}

	len, err := io.WriteString(file, constent)

	if err != nil {
		panic(err)
	}

	fmt.Println("Wrote bytes to file", len)

	readfile()
	defer file.Close()
}

func readfile() {
	data, err := ioutil.ReadFile("test.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("File contents:", string(data))
}
