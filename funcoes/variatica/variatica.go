package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Println("Media geral eh: ", media(7.8, 8.1, 9.0))
}
