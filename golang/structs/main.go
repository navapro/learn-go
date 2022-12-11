package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main () {
	jimsContact := contactInfo {
		email: "jim@gmail.com",
	 	zipCode: 95000,
	}
	jim := person {firstName: "jim", lastName: "Paryt", contact: jimsContact}
	jim.updateName("jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string){
	pointerToPerson.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}