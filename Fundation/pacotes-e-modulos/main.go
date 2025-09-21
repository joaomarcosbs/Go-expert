package main

import (
	"fmt"
	"github.com/google/uuid"
	"pacotes-e-modulos/matematica"
)

func main() {
	s := matematica.Soma(85, 3)
	fmt.Printf("O Resultado da some é: %v \n", s)
	fmt.Printf("O uuid é: %v", uuid.New())
}
