package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("arquivos.txt")

	if err != nil {
		panic(err)
	}

	// escrevendo uma string
	//tamanho, err := f.WriteString("Hello, World!")

	// escrevendo bytes
	tamanho, err := f.Write([]byte("Hello, World! Escrevendo no arquivo"))

	if err != nil {
		panic(err)
	}
	fmt.Println("Arquivo criado com sucesso! Tamanho: %d bytes", tamanho)
	f.Close()

	// leitura
	arquivo, err := os.ReadFile("arquivos.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivos.txt")
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

	err = os.Remove("arquivos.txt")
	if err != nil {
		panic(err)
	}
}
