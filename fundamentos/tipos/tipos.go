package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro eh", reflect.TypeOf(32000))

	// sem sinal (so positivo)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte eh ", reflect.TypeOf(b))

	// com sinal... uint8 uint16 uint32 uint64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int eh", i1)
	fmt.Println("O valor maximo do int eh", reflect.TypeOf(i1))

	var i2 rune = 'a' // Representa um mapeamento da tabela de Unicode (int32)
	fmt.Println("O rune eh", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais (float32,float64)

}
