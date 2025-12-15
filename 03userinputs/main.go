package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the ratings of pizza: ")

	// comma OK || err ok

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for Rating", input)
	fmt.Printf("Type of this rating is %T", input)

}
