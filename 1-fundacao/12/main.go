package main

func main() {

	// Memoria -> Endereço -> Valor

	// variavel -> ponteiro que tem um endereco na memoria -> valor
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	println(*b) // imprimir o valor do ponteiro

}
