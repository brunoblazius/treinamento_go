package main

import "fmt"

func main() {
	fmt.Print("Mesma linha")

	fmt.Println(" Nova linha")
	fmt.Println("linha.")

	x := 3.1415

	xs := fmt.Sprint(x)
	fmt.Println("O Valor de x eh " + xs)
	fmt.Println("O valor de x eh", x)

	a := 1
	b := 1.9999
	c := false
	d := "Opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

}
