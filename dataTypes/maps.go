package dataTypes

import "fmt"

func MapExample() {
	myMap := make(map[string]uint8)
	myMap["Pranish"] = 21
	myMap["Anushree"] = 19
	// fmt.Println(myMap["Pranish"])

	myMap["Pranish"] = 22

	for name, age := range myMap {
		fmt.Printf("Name: %s\nAge: %d\n", name, age)
	}
}
