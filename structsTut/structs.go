package structsTut

import "fmt"

type Student struct {
	name    string
	age     uint8
	address string
	email   string
	course  string
}

func Computing() {
	Pranish := Student{
		name:    "Pranish",
		age:     20,
		address: "Kadaghari",
		email:   "pranish@gmail.com",
		course:  "Go",
	}
	fmt.Println(Pranish.name, Pranish.age, Pranish.address, Pranish.email, Pranish.course)
}
