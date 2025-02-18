package main

import (
	"fmt"
	"github.com/alanocampos/pos-go/5-Packaging/1/math"
)

func main() {

	m := math.NewMath(1, 2)
	m.C = 3
	fmt.Println(m.Sum())
}
