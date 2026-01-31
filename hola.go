package main

import (
	"fmt"

	"rsc.io/quote"
)

// Declaracion de constantes
const Pi float32 = 3.141592653589793

const (
	x = 100
	y = 0b100
	z = 0o12
	w = 0xFF
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Go())

	// Declaracion de variables
	var firstName, lastName, age = "Jane", "Doe", 23

	fmt.Println(firstName, lastName, age)

	fmt.Println(x, w, y, z)
}
