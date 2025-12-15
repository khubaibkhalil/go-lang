package main

import "fmt"

func main() {
	fmt.Println("structs in golang")
	//no inheritance / no super or parent

	khubaib := User{"Khubaib", "mkhubaibkhalil@gmail.com", true, 16}
	fmt.Println(khubaib)
	fmt.Printf("Details: %+v\n", khubaib)
	khubaib.GetStatus()
	khubaib.NewMail("mkhubaibkhalil@hotmail.com")
	fmt.Printf("Name is %v and email is %v and status is %v and age is %v\n", khubaib.Name, khubaib.Email, khubaib.Status, khubaib.Age)

}

// first capital public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int
}

func (u User) GetStatus() {
	fmt.Println(u.Status)
}

func (u User) NewMail(mail string) {
	u.Email = mail
	fmt.Println(u.Email)
}
