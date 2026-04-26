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
	
	intSlice2 := []int8{10,11}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("%p\n", intSlice)
	fmt.Printf("%p\n", intSlice2)
	
	fmt.Println(intSlice)
	fmt.Println(len(intSlice))
	fmt.Println(cap(intSlice))

}
