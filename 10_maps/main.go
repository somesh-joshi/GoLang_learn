package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["go"] = "Go"
	languages["js"] = "JavaScript"
	languages["py"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["go"])

	delete(languages, "py")
	fmt.Println(languages)

	// loop through map

	for key, value := range languages {
		fmt.Println(key, value)
	}
	
}