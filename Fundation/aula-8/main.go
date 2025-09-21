package main

import "fmt"

func main() {
	fmt.Println(sum3(5, 9))
}

// Se tiver um tipo depois dos parâmetros da função (), quer dizer que ela tem retorno. Se começar direto o corpo da função {}, não tem retorno
func sum(a int, b int) int {
	return a + b
}

// Se os parâmetros da função possuirem o mesmo tipo, não há necessidade de tipar cada um
func sum2(a, b int) int {
	return a + b
}

// As funções podem ter mais de um retorno, e os tipos devem ser colocados entre parênteses também depois dos parâmetros
func sum3(a, b int) (int, bool) {
	if a >= b {
		return a + b, true
	}
	return a + b, false
}
