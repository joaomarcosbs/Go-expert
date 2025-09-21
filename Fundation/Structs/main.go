package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
	// Poderia também ser atribuído como tipo:
	// Address Endereco
}

func (c *Cliente) DesativarCliente() {
	c.Ativo = false
	c.Nome = "Birula"
	fmt.Printf("\n\nO status do cliente %s foi desativado\n\n", c.Nome)
}

func main() {
	joao := Cliente{
		Nome:  "João",
		Idade: 28,
		Ativo: true,
	}

	joao.Cidade = "Canoas"
	joao.Estado = "Rio Grande do Sul"
	joao.Rua = "Rua Jari"
	joao.Numero = 133

	fmt.Printf("Meu nome é %s\nTenho %d anos\nStatus: %t", joao.Nome, joao.Idade, joao.Ativo)

	joao.DesativarCliente()

	fmt.Println(joao.Ativo, "\n", joao.Nome)

}
