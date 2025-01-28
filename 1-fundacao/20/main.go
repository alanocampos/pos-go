package main

func main() {
	// condicionais
	a := 1
	b := 2
	c := 3

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")

	}

	if a > b && c > a {
		println(a)
	} else {
		println(b)
	}
}
