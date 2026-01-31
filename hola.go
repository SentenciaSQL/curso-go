package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Go())

	var (
		defaultInt     int
		defaultUint    uint
		defaultFloat   float32
		defaultBoolean bool
		defaultString  string
	)

	fmt.Println(defaultInt, defaultUint, defaultFloat, defaultBoolean, defaultString)

}
