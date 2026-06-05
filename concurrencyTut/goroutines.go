package concurrencyTut

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var mockData = []string{"Pranish", "Anushree", "Ram", "Shyam", "John", "Smith"}

func SomeWork() {
	time0 := time.Now()
	for i := range len(mockData) {
		wg.Add(1)
		go mockDbCall(i)
	}
	wg.Wait()
	fmt.Printf("\n Total execution time %v\n", time.Since(time0))

}

func mockDbCall(i int) {
	var delay float32 = rand.Float32() * 2000

	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the db is: ", mockData[i])
	wg.Done()
}
