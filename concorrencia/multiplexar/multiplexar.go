package main

import (
	"fmt"

	"github.com/brunoblazius/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.google.com", "https://www.youtube.com"),
		html.Titulo("https://www.amazon.com", "https://www.gmail.com"),
	)
	fmt.Println(<-c)
}
