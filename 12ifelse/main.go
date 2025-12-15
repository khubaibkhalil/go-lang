package main

import "fmt"

func main() {
	fmt.Println("IF Else")
	var result string

	loginCount := 23

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 25 {
		result = "if else"
	} else {
		result = "else"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is Not less then 10")
	}

	// if err!=  nill{

	// }
}
