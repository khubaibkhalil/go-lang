package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Apple", "Orange", "Banana"}

	fmt.Printf("%T", fruitList)

	fruitList = append(fruitList, "Pear")

	fmt.Println(fruitList)

	// fruitList = append(fruitList[:2])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	// highScore[4] = 777

	highScore = append(highScore, 555, 666, 321)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Print(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))
	fmt.Println("")

	//how to remove a value as index
	var courses = []string{"react", "javascript", "python", "ruby"}
	// 							0,    , 1			2,			3

	fmt.Println(courses)
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
