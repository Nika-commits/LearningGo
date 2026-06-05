package main

import (
	"errors"

	"github.com/Nika-commits/LearningGo/concurrencyTut"
)

type Person struct {
	Name    string
	Age     int
	Phone   string
	Address string
}

func main() {

	// p := Person{
	// 	Name:    "Pranish",
	// 	Age:     30,
	// 	Phone:   "123-456-7890",
	// 	Address: "123 Main St",
	// }

	// p1Name, p1Age, err := NameAndAge(p)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Printf("The person's name is %v with age of %v", p1Name, p1Age)
	// }
	// datatypes.Primitive()
	// datatypes.Arrays()
	// datatypes.Slices()
	// datatypes.MapExample()
	// loops.Loop()
	// dataTypes.Runes()
	// structsTut.Computing()
	// var a uint8 = 10
	// var b *uint8 = &a

	// fmt.Printf("Before Changine value: %v\n", a)
	// *b = 2
	// fmt.Printf("After Changine value: %v\n", a)
	// fmt.Printf("Value of b : %v\n", *b)
	concurrencyTut.SomeWork()
}

func NameAndAge(p Person) (string, int, error) {
	if p.Name == "" || p.Age == 0 {
		err := errors.New("name is empty or age is zero")
		return "", 0, err
	}
	return p.Name, p.Age, nil
}
