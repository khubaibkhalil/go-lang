package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	// fruitList[2] = "Banana"
	fruitList[3] = "Mango"

	fmt.Println("LIst is ", fruitList)

	fmt.Println(len(fruitList))

	var vegList = [4]string{" Potato", "Mashrooms", "beans"}

	fmt.Println("List", vegList)
}
