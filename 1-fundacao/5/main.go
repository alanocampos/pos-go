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

	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	fmt.Println(len(meuArray))

	// percorrendo o indice e o valor do array
	for i, v := range meuArray {
		fmt.Printf("%d\t%d\n", i, v)
	}
}
