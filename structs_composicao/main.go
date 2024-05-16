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

func main() {
	savio := Cliente{
		Nome:  "Sávio",
		Idade: 42,
		Ativo: true,
	}

	savio.Ativo = false
	savio.Cidade = "Altamira"

	fmt.Printf("Cliente %v, mora na cidade de %v\n", savio.Nome, savio.Cidade)
}
