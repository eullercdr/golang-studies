package main

import (
	"github.com/eullercdr/math"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New().String()
	println(id)
	math := math.Math{A: 3, B: 4}
	p := math.Add()
	println(p)
}
