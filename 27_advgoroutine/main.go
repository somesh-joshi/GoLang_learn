package main

import "fmt"

func myfun(mychnl chan string) {

	var i int
	fmt.Println("add date")
	fmt.Scanln(&i)

	for v := 0; v < i; v++ {
		mychnl <- "some data"
	}
	close(mychnl)
}

func main() {

	c := make(chan string)

	go myfun(c)

	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
}
