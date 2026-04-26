package dataTypes

import "fmt"

func Arrays() {
	var intArr [5]uint16 = [5]uint16{1, 2, 3, 4, 5}
	fmt.Println(intArr)
	fmt.Println(intArr[2:4])

	intArr[1] = 1
	fmt.Println(&intArr[1])

}
