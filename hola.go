package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Go())

	fullName := "Andres Frias \t(alias \"afriasdev\")\n"
	fmt.Print(fullName)

	var a byte = 'A'
	fmt.Println(a)

	s := "Hola"
	fmt.Println(s[0])
}
