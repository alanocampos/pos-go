package main

import (
	"fmt"
)

func main() {

	// funcao dentro de uma funcao
	total := func() int {
		return sum(1, 2, 3, 4, 5) * 2
	}()

	fmt.Println(total)
}

// retornando mais de um parametro na função
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
