package main

import (
	"fmt"
)

func main() {

	fmt.Println("Loop")
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])

	// }

	for index, day := range days {
		fmt.Printf("\nindex is %v and Value is %v ", index, day)
	}

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("LCO")
}
