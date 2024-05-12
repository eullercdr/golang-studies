package main

func main() {
	a := 10
	b := 20
	if a > b {
		println("a is greater than b")
	} else {
		println("b is greater than a")
	}
	switch a {
	case 10:
		println("a is 10")
	case 20:
		println("a is 20")
	default:
		println("a is neither 10 nor 20")
	}
}
