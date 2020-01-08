package main

import (
	"fmt"
	"time"
)

// Channel (canal) -  e a forma de cominicacao entra goroutines
// e um tipo de dado

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)

	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c // rebendo valores do canal

	fmt.Println(a, b)

	fmt.Println(<-c)
}
