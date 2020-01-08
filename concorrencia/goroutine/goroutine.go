package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interacao %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// chamada em serial
	// fale("Maria", "Pq vc nao fale comigo?", 3)
	// fale("Joao", "So posso falar depois de vc!", 3)

	// go fale("Maria", "Pq vc nao fale comigo?", 500)
	// go fale("Joao", "Pq vc nao fale comigo?", 500)

	go fale("Maria", "Entendi...", 10)
	fale("Joao", "Parabens!", 5)
}
