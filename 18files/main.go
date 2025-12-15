package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This is a file content"

	file, err := os.Create("./file.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("Lenght is ", length)

	defer file.Close()

	readFile("./file.txt")
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)

	checkNilErr(err)
	fmt.Println("Text data is :", string(data))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
