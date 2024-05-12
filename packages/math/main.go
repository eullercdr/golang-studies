package math

func Sum[T int | float64](a, b T) T {
	return a + b
}

type Car struct {
	brand string // is private to the package
	Model string
	Year  int
}
