package main

import "fmt"

//Token := "1234567890"
// it will throw error
// define it as var Token or Const Token at list
// apt way is // Const Token := "1234567890" or // Const Token int = 1234567890

const LoginToken string = "hgfgfd"

// Put up the First Latter is as Capital for the public variable
// Put up the First Latter is as Small for the private variable
// Put up the First Latter is as Capital for the public function
// Put up the First Latter is as Small for the private function

func main() {
	var UserName string = "Somesh"
	fmt.Printf("Value of UserName is %s and type is %T \n", UserName, UserName)

	var isLogin bool = true
	fmt.Println(isLogin)
	fmt.Printf("Value of isLogin type is %T \n", isLogin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Value of smallVal type is %T \n", smallVal)

	var smallFloat float32 = 255.5
	fmt.Println(smallFloat)
	fmt.Printf("Value of smallFloat type is %T \n", smallFloat)

	// default value of aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Value of anotherVariable type is %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Value of anotherString type is %T \n", anotherString)

	// implicit type declaration

	var website = "https://www.someshkoli.com"
	fmt.Println(website)
	// you can't redeclare the variable type in goLang e.g: website = 1 // this will throw error

	// no var style

	numberOfUser := 100
	fmt.Println(numberOfUser)
	//numberOfUser := 200.20 // this will throw error

	fmt.Println(LoginToken)
	fmt.Printf("Value of LoginToken type is %T \n", LoginToken)
}
