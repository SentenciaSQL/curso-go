package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Go())

	var name string
	var age int

	fmt.Println("What is your name?")
	fmt.Scanln(&name)

	fmt.Println("What is your age?")
	fmt.Scanln(&age)

	fmt.Printf("Hola me llamo, %s y tengo %d!\n", name, age)

	greeting := fmt.Sprintf("Hola me llamo, %s y tengo %d!", name, age)

	fmt.Println(greeting)

}
