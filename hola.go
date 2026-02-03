package main

import (
	"fmt"
	"math"
)

func main() {
	var lado1, lado2 float64
	const precision = 2

	fmt.Print("Introduzca lado1: ")
	fmt.Scanln(&lado1)

	fmt.Print("Introduzca lado2: ")
	fmt.Scanln(&lado2)

	area := (lado1 * lado2) / precision
	hipotenusa := math.Sqrt(math.Pow(lado1, precision) + math.Pow(lado2, precision))
	perimetro := lado1 + lado2 + hipotenusa

	fmt.Printf("\nArea: %.*f", precision, area)
	fmt.Printf("\nPerimetro: %.*f", precision, perimetro)

}
