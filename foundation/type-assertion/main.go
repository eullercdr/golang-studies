package main

import "fmt"

func main() {
	var myVar interface{} = "Welcome to Golang"
	println(myVar.(string))
	res, ok := myVar.(int)
	fmt.Printf("O valor de res é %v e ok é %v\n", res, ok)
	res2 := myVar.(int) //panic error
	fmt.Printf("O valor de res2 é %v\n", res2)
}
