package main

import "fmt"

func main() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("Hello")
	myDefer()

	// 3
	// 2
	// 1
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	// 0 1 2 3 4
}
