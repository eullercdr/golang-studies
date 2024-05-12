package main

import (
	"fmt"

	"github.com/eullercdr/math/math"
	"github.com/google/uuid"
)

func main() {
	sum := math.Sum(10, 20)
	fmt.Println("Result: 10 + 20 =", sum)

	car := math.Car{
		Model: "Corolla",
		Year:  2021,
	}
	fmt.Println("Car:", car)
	println(uuid.New().String())
}
