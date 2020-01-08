package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"Jose Joao":      11326.45,
		"Maria Dolce":    15432.67,
		"Paulo Henrique": 1200.00,
	}

	funcESalarios["Rafael tese"] = 1300.0

	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}
}
