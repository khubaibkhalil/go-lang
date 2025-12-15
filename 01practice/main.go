package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func main() {
// 	reader := bufio.NewReader(os.Stdin)

// 	input, _ := reader.ReadString('\n')
// 	input = strings.TrimSpace(input)

// 	num, _ := strconv.ParseInt(input, 10, 64)

// 	if num%2 == 0 {
// 		fmt.Printf("%d is Even", num)
// 	} else {
// 		fmt.Printf("%d is Odd", num)
// 	}
// }

func main() {

	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter input:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// num, _ := strconv.Atoi(input)
	num, _ := strconv.ParseInt(input, 10, 64)

	// fmt.Printf("num %T", num)
	fmt.Println("User Input is ", num)

}
