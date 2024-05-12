package main

import "fmt"

type ID int

var (
	a ID
)

func main() {
	fmt.Printf("O tipo de a é %T", a)
}
