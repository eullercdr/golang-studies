package main

import "github.com/eullercdr/math"

func main() {
	math := math.Math{A: 3, B: 4}
	p := math.Add()
	println(p)
}
