package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "hello"
	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", t, t)
}
