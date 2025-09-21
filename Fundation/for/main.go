package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	numeros := []string{"a", "b", "c", "d", "e", "f", "g"}

	for _, numero := range numeros {
		fmt.Println(numero)
	}

	a := 0

	for a < 10 {
		println(a)
		a++
	}

	// loopping infinito ex.: consumir uma mensagem
	//for {
	//	println(a)
	//}
}
