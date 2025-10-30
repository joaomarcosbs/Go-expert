package main

import (
	"fmt"
	"packaging/math" // importando o pacote com o nome do pacote/import (sempre o endereço, preferencialmente url)
)

func main() {
	fmt.Println("Hello World!")
	m := math.Math{A: 10, B: 20}
	fmt.Println(m.Add())
}
