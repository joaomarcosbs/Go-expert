package main

import "fmt"

func main() {
	// map[tipo da chave]tipo do valor
	salarios := map[string]int{"João": 1000, "Pedro": 2000, "Augusto": 3000}
	fmt.Println(salarios["Augusto"])
	//delete(salarios, "Augusto")
	//fmt.Println(salarios["Augusto"])
	salarios["Serjão"] = 800
	fmt.Println(salarios["Serjão"])

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é: %d\n", nome, salario)
	}

	// blank identifier
	for _, salario := range salarios {
		fmt.Printf("O salário é: %d\n", salario)
	}
	// newSal := make(map[string]int)
}
