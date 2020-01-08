package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[3234234234] = "Maria"
	aprovados[3221212121] = "Paulo"
	aprovados[4545454545] = "Luiz"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[4545454545])
	delete(aprovados, 4545454545)
	fmt.Println(aprovados)
}
