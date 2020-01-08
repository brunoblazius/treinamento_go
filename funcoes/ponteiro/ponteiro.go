package main

import "fmt"

func inc1(n int) {
	n++
}

// revisao: ponteiro e representando por um *
func inc2(n *int) {
	//revisao: * e usado para desreferenciar , ou seja
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	// revisao: & usando paraobter o endereco da variavel
	inc2(&n)
	fmt.Println(n)

}
