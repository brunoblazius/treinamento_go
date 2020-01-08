package main

import "math"

// Iniciando com a letra maiuscula e publico (visivel dentro e fora do pacote)
// Iniciando com letra minuscula e privafo (visivel somente dentro do pacote

// por exemplo....
// Ponto = gerara al publico
// ponto ou _Ponto - gerara algo privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)

	return
}

// Distancia e responsavel por calcular a distancia.
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
