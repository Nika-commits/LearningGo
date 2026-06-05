package loops

import (
	"fmt"
)

func Loop() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	array1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// fmt.Printf("Array1 Size: %d bytes\n", unsafe.Sizeof(array1))

	// array2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Printf("Array2 Size: %d bytes\n", unsafe.Sizeof(array2))

	for i := range len(array1) {
		fmt.Println(i)
	}
}
