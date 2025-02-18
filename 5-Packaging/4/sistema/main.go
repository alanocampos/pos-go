package main

import (
	"github.com/alanocampos/pos-go/5-Packaging/3/math"
	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Sum())
	println(uuid.New().String())
}
