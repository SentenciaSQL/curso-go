package main

import (
	"fmt"
	"strconv"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Go())

	s := "100"
	i, _ := strconv.Atoi(s)

	fmt.Println(i + 1)

	n := 42
	s = strconv.Itoa(n)

	fmt.Println(s)

}
