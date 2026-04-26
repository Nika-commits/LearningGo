package dataTypes

import "fmt"

func Slices() {
	intSlice := []int8{1, 2, 3, 4, 5}
	fmt.Printf("%p", intSlice)
	fmt.Println(len(intSlice))
	fmt.Println(cap(intSlice))

	intSlice = append(intSlice, 8)
	fmt.Printf("%p", intSlice)
	fmt.Println(len(intSlice))
	fmt.Println(cap(intSlice))

}
