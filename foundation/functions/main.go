package main

import "fmt"

func main() {
	fmt.Println(sum2(5, 6))
}

func sum(a int, b int) int {
	return a + b
}

func sum2(a, b int) (int, bool) {
	if a+b > 10 {
		return a + b, true
	}
	return a + b, false
}
