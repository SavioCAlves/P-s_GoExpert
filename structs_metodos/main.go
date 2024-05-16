package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) DesativarCliente() {
	c.Ativo = false
	fmt.Printf("O cliente %v foi desativado!\n", c.Nome)
}

func main() {
	savio := Cliente{
		Nome:  "Sávio",
		Idade: 42,
		Ativo: true,
	}

	// savio.Ativo = false
	savio.Cidade = "Altamira"

	savio.DesativarCliente()
}
