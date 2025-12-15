package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// var ptr *int

	// fmt.Println("Value of Pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of Pointer is ", ptr)
	fmt.Println("Value of Pointer is ", *ptr)

	*ptr = *ptr * 2

	fmt.Println("Value", myNumber)
}
