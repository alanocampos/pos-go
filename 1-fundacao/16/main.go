package main

import "fmt"

func main() {

	var minhaVar interface{} = "Alano Campos"

	println(minhaVar.(string))

	// type assertions, res sera o valor da conversar se der certo, ok vai ser o resultado da conversao se deu certo
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é: %v, e o resultado de ok é:  %v\n", res, ok)

	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v", res2)
}
