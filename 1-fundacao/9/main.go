package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(1, 2, 3, 4, 5))
}

// retornando mais de um parametro na função
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
