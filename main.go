package goarea

import "math"

const Pi = 3.1416

//circ calculará a área da circunferencia.
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect responsável por calcular a área de triangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é visivel devido ao _
func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}

fmt.println("Fim do programa")
