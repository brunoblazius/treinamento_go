package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// e outros metodos.
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Ligar Turbo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Fazer baliza...")
}

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.fazerBaliza()
	b.ligarTurbo()
}
