package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"B": {
			"Bruno Fabiano": 11500.77,
			"Beto Carreiro": 12333.00,
		},
		"F": {
			"Fabiano Blazius": 49000.00,
		},
		"P": {
			"Pedro Junior": 33009.00,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
