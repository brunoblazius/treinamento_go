package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice

	fmt.Println(a1, s1)

	a2 := [5]int{1, 2, 3, 4, 5}

	// slice nao eh um array, e sim um pedaço de um array

	s2 := a2[1:3]
	fmt.Println(a2, s2)
}
