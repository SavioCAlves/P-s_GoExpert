package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Gravando string no arquivo
	// tamanho, err := f.WriteString("Hello, world!")

	// Gravando dados no arquivo
	tamanho, err := f.Write([]byte("Hello, world!"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	f.Close()

	// Leitura de arquivos

	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	// Os dados sempre são armazenados em binário, por isso a necessidade de converter para strings.
	fmt.Println(string(arquivo))

	// Leitura de pouco em pouco, utilizada para preservar a memória.
	// Exemplo: um container com apenas 100 MB de memória precisar abrir um arquivo de 1GB ou maior.
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// Removendo o arquivo
	err = os.Remove("arquivo.txt")
}
