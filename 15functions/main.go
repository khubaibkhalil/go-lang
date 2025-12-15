package main

import "fmt"

func main() {
	fmt.Println("Main")

	greetings()

	result := adder(3, 5)
	func() {
		fmt.Println("2")
	}()

	fmt.Println(result)

	fmt.Print(proAdder(10, 20, 30, 40, 50))
}
func proAdder(values ...int) int {
	total := 0
	// fmt.Printf("%T", values)
	for i := range values {
		fmt.Println(values[i])
		total += values[i]
	}
	return total
}

func greetings() {
	fmt.Println("Greetings bro")
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}
