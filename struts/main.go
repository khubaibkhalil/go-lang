package main

import (
	"fmt"
)

type Employee struct {
	name     string
	age      int
	isRemote bool
	Address
}
type Address struct {
	street string
	city   string
}

func (a Address) printAddress() {
	fmt.Println("Address")
}

func (e *Employee) updateName(newName string) {
	e.name = newName

}

func (e Employee) Print() {
	fmt.Printf("Employee Name : %v \n", e.name)
	fmt.Printf("Employee Age : %v \n", e.age)
	fmt.Printf("Employee Satus : ")

	if e.isRemote {
		fmt.Printf("Remote")
	} else {
		fmt.Printf("Onsite")
	}
	fmt.Printf("\nEmployee Address : %v %v\n", e.street, e.city)

}

func main() {

	emp1 := Employee{name: "Khubaib", age: 22, isRemote: false, Address: Address{street: "Avenue 9", city: "Lahore"}}

	emp2 := Employee{
		name:     "Khubaib",
		age:      22,
		isRemote: false,
	}

	emp1.Print()
	emp2.Print()
	fmt.Println("")
	job := struct {
		title  string
		salary int
	}{
		title:  "SE",
		salary: 1000,
	}

	fmt.Printf("%T", job)

	employeePtr := &emp1

	employeePtr.age = 90
	emp1.updateName("Mani")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	emp1.printAddress()
}
