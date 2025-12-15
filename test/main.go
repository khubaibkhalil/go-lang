package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string

	reader := bufio.NewReader(os.Stdin)

	name, _ = reader.ReadString(' ')

	fmt.Println("Welcome to go lang", name)

}
