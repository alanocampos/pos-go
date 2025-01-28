package main

import (
	"errors"
	"fmt"
)

func main() {
	// validando se retornou um erro
	valor, err := sum(5, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

// retornando mais de um parametro na função
func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
