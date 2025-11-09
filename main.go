package main

import (
	"fmt"
	"time"
)

func sum(id int, xs []int, c chan int) {
	sum := 0

	fmt.Printf("Goroutine %d start.\n", id)

	for _, v := range xs {
		sum += v
		time.Sleep(time.Second)
	}

	fmt.Printf("Goroutine %d stop.\n", id)

	c <- sum
}

func main() {
	fmt.Println("Hello, concur ints and stdio.")

	xs := []int{7, 2, 8, -9, 4, 0}

	// 1. create channel of (native) type int
	c := make(chan int)

	// 2. push 2 sums of splitting xs into 2 halves
	leftxs := xs[:len(xs)/2]
	rightxs := xs[len(xs)/2:]
	go sum(1, leftxs, c)
	go sum(2, rightxs, c)

	// 3. pull 2 sums from channel
	x, y := <-c, <-c

	// 4. output state
	fmt.Println(x, y, x+y)

	fmt.Println("Good bye.")
}
