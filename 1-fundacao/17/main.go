package main

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

type MyNumber int

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	m := map[string]int{"Alan": 1000, "John": 2000, "Maria": 3000}
	m2 := map[string]float64{"Alan": 1000.10, "John": 2000.30, "Maria": 3000.0}
	m3 := map[string]MyNumber{"Alan": 1000, "John": 2000, "Maria": 3000}

	println(SomaInteiro(m))
	println(SomaFloat(m2))

	// utilização de generics

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))

}
