package main

import "fmt"

func main() {
	hello := make(chan string)
	go receive("Hello", hello)
	read(hello)
}

func read(data <-chan string) {
	fmt.Println(<-data)
}

func receive(name string, hello chan<- string) {
	hello <- name
}
