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

type Empresa struct {
	Nome string
	Cnpj int
}

func (e Empresa) Desativar() {
	fmt.Println("Empresa desativada")
}

type Pessoa interface {
	Desativar()
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %v foi desativado!\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	savio := Cliente{
		Nome:  "Sávio",
		Idade: 42,
		Ativo: true,
	}

	// savio.Ativo = false
	savio.Cidade = "Altamira"

	savio.Desativar()

	Desativacao(savio)

	minhaEmpresa := Empresa{
		Nome: "Fatasia123",
		Cnpj: 12321231231111,
	}

	Desativacao(minhaEmpresa)
}
