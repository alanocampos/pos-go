package main

import "fmt"

func main() {

	var a interface{} = 10
	var b interface{} = "Hello, World!"

	showType(a)
	showType(b)

}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v \n", t, t)
}
