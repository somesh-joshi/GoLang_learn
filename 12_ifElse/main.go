package main

import "fmt"

func main() {

	loginCount := 10
	var message string

	if loginCount > 5 {
		message = "regular user"
	}else if loginCount < 10 {
		message = "power user"
	}else {
		message = "super user"
	}

	fmt.Println(message)
}
