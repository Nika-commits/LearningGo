package structsTut

import (
	"fmt"
)

type Student struct {
	name    string
	age     uint8
	address string
	email   string
	course  Course
}

type Course struct {
	name         string
	duration     uint8
	moduleLeader string
}

func Computing() {
	BSC_Computing := Course{
		name:         "BSc Hons in Computing",
		duration:     3,
		moduleLeader: "Anushree Paudyal",
	}
	Pranish := Student{
		name:    "Pranish",
		age:     20,
		address: "Kadaghari",
		email:   "pranish@gmail.com",
		course:  BSC_Computing,
	}
	fmt.Println(Pranish.name, Pranish.age, Pranish.address, Pranish.email, Pranish.course.name, Pranish.course.duration, Pranish.course.moduleLeader)
}
