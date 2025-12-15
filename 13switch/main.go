package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Switch")
	r := rand.New(rand.NewSource(time.Now().Unix()))

	num := r.Intn(6) + 1

	switch num {
	case 1:
		fmt.Println("Dice Value is 1")

	case 2:
		fmt.Println("Dice Value is 2")

	case 3:
		fmt.Println("Dice Value is 3")

	case 4:
		fmt.Println("Dice Value is 4")

	case 5:
		fmt.Println("Dice Value is 5")

	case 6:
		fmt.Println("Dice Value is 6")
	default:
		fmt.Print("Invalid")
	}

}
