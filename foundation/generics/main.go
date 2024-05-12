package main

type MyNumber int

// create a constraint for the type Number
type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compair[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	salaries := map[string]int{
		"José":  1000,
		"Maria": 5000,
		"Pedro": 2000,
	}
	println(Soma(salaries))
	salaries2 := map[string]float64{
		"José":  1000.50,
		"Maria": 5000.75,
		"Pedro": 2000.25,
	}
	println(Soma(salaries2))
	salaries3 := map[string]MyNumber{
		"José":  1000,
		"Maria": 5000,
		"Pedro": 2000,
	}
	println(Soma(salaries3))
	println(Compair(10, 10))
}
