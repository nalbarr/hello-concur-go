package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum 
}

func main() {
	fmt.Println("Hello, concur ints and stdio.")

	xs := []int{7, 2, 8, -9, 4, 0}

	// 1. create channel of (native) type int
	c := make(chan int)

	// 2. push 2 sums of splitting xs into 2 halves
	go sum(xs[:len(xs)/2], c)
	go sum(xs[len(xs)/2:], c)

	// 3. pull 2 sums from channel
	x, y := <-c, <-c 

	// 4. output state
	fmt.Println(x, y, x+y)

	fmt.Println("Good bye.")
}
