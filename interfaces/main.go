package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Measureable interface {
	Perimeter() float64
}

type Geometry interface {
	Shape
	Measureable
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func descibeShape(g Geometry) {
	fmt.Println("Area: ", g.Area())
	fmt.Println("Perimeter", g.Perimeter())
}

func main() {
	rect := Rectangle{width: 10, height: 20}

	descibeShape(rect)

	// circ := Circle{radius: 2}

	// fmt.Println("Rectangle Area", calculateArea(rect))

	// fmt.Println("Cicle Area", calculateArea(circ))

	// mysteryBox := interface{}(1)
	// descibeValue(mysteryBox)

	// retrievedInt, ok := mysteryBox.(int)
	// if ok {
	// 	fmt.Println("Retrieved int: ", retrievedInt)
	// } else {
	// 	fmt.Println("Failed")
	// }

}

func descibeValue(t interface{}) {
	fmt.Printf("Type: %T , Value: %v", t, t)

}

type CalcuteError struct {
	msg string
}

func (ce CalcuteError) Error() string {
	return ce.msg
}

func performCalculation(val float64) (float64, error) {
	if val < 0 {
		return 0, CalcuteError{msg: "Invalid"}
	}

	return math.Sqrt(val), nil
}
