package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Phone   string
	Address string
}

func main() {

	p := Person{
		Name:    "",
		Age:     30,
		Phone:   "123-456-7890",
		Address: "123 Main St",
	}

	p1Name, p1Age, err := NameAndAge(p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("The person's name is %v with age of %v", p1Name, p1Age)
}

func NameAndAge(p Person) (string, int, error) {
	if p.Name == "" {
		err := errors.New("name is empty")
		return "", 0, err
	}
	return p.Name, p.Age, nil
}
