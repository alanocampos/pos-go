package main

import "fmt"

const a = "Hello, World!"

type ID int // criando um tipo próprio

var (
	b bool
	c int
	d string
	e float64
	f ID = 1
)

func main() {

	fmt.Printf("O tipo de E é %T", f)
}
