package main

import "fmt"

// jwtToken := 30000 error
// var jwtToken int = 300
// const LoginToken string = "gashlggashg" capital letter means public

func main() {
	// fmt.Println("variables")
	var username string = "Khubaib"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n ", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n ", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n ", smallVal)

	var smallFloatVal float64 = 255.4556575757474567
	fmt.Println(smallFloatVal)
	fmt.Printf("Variable is of type: %T \n ", smallFloatVal)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n ", anotherVariable)

	// var website = " code.com"
	// website = 5 error
	numberOfUser := 30000

	fmt.Print(numberOfUser)
}
