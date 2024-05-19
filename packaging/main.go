package main

import (
	"fmt"

	"github.com/eullercdr/packaging/math"
)

func main() {
	m := math.Math{A: 10, B: 20}
	fmt.Println(m.Add())
}
