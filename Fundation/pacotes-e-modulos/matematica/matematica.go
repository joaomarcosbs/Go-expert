package matematica

// As funções e variáveis e structs, inclusive dentro das structs, que começam
// com letra maiúscula, são visiveis em qualquer pacote
// se começa com letra minúscula, só é visível no pacote em que é criado
func Soma[T int | float64](a, b T) T {
	return a + b
}
