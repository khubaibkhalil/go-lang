package main

import "fmt"

type Animal interface {
	Speak()
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println("Meowwwww")
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println("Bowwwww")

}

func call(a Animal) {
	a.Speak()
}

func main() {
	cat := Cat{Name: "Meow"}

	dog := Dog{Name: "Bowww"}

	call(cat)

	call(dog)

}
