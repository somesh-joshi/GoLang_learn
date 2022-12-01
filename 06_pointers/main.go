package main

import "fmt"

func main() {
	var a int = 42
	var b *int = &a
	// ref means &

	fmt.Println(a, b)
	fmt.Printf("%T %T\n", a, b)

	c := 60
	var d = &c
	fmt.Println(c, d)
	fmt.Printf("%T %T\n", c, d)

	*d = *d + 2
	fmt.Println(c)
}
