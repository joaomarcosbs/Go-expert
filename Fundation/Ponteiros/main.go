package main

func main() {
	a := 10
	ponteiro := &a

	// valor de a
	println("Valor de a que recebeu 10 --> ", a)

	// endereço de memória onde está a
	println("Valor de &a --> ", &a)

	// com * mostra o VALOR dentro da variável
	println("Valor de *ponteiro que recebeu &a --> ", *ponteiro)

	// sem * mostra o ENDEREÇO DE MEMÓRIA
	println("Valor de ponteiro que recebeu &a --> ", ponteiro)

	// alterando o VALOR que tá em ponteiro
	*ponteiro = 20
	println("Valor de *ponteiro que foi alterado para 20 --> ", *ponteiro)

	// endereço da memória permanece o mesmo
	println("Endereço na memória de ponteiro permanece igual, mesmo alterando seu valor --> ", ponteiro)

	// criando variável b também com endereço de memória de a
	b := &a

	// endereço de memória de b
	println("Endereço de memória de b --> ", b)

	// valor de b
	println("Valor de b --> ", *b)

	//atribuíndo VALOR a b
	*b = 30
	println("Alterado o valor de *b para 30")

	// resultados da alteração de VALOR de b
	println("Valor de a -->", a)
	println("Valor de *b --> ", *b)
	println("Valor de *ponteiro -->", *ponteiro)
	println("Valor de &a --> ", &a)
	println("Valor de ponteiro --> ", ponteiro)
	println("Valor de b --> ", b)
}
