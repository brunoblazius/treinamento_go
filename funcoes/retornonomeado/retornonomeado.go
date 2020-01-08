package main

import "fmt"

func troca(p1, p2 int) (segundo int, primeiro int) {
	primeiro = p1
	segundo = p2

	return // retorno limpo
}

func main() {
	x, y := troca(2, 3)
	fmt.Println(x, y)
}
