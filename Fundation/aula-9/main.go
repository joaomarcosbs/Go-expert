package main

import "fmt"

func main() {
	fmt.Println(sum(5, 9, 75, 269, 96, 87))
}

// quando não se sabe o número de parâmetros que serão passados
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros { // blank identifier
		total += numero
	}
	return total
}
