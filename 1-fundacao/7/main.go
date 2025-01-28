package main

import "fmt"

func main() {
	salarios := map[string]int{"Alano": 1000, "John": 2000, "Paul": 3001}

	fmt.Println(salarios["Paul"])

	delete(salarios, "Paul")

	salarios["Paula"] = 3000

	fmt.Println(salarios["Paul"])

	//sal := make(map[string]int)
	//sal1 := map[string]int{}

	for nome, salario := range salarios {
		fmt.Println(nome, salario)
	}

	// blank identifier
	for _, salario := range salarios {
		fmt.Println(salario)
	}
}
