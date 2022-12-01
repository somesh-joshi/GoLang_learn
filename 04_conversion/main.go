package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please enter a number: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numRating + 1)
	}
}

// IT WORKS!
/*
	fmt.Println("Please enter a number: ")

	var input float64
	fmt.Scanf("%f", &input)

	output := input + 2
	fmt.Println(output)
	fmt.Printf("%T", output)
*/
