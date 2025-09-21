package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool
	c int
	d string
	e float32
	f ID
)

func main() {
	shortDeclarationVar := "assim é mais rápido de declarar uma variável"
	println(shortDeclarationVar)
	fmt.Printf("O tipo de shortDeclarationVar é: %T\n", shortDeclarationVar)
}
