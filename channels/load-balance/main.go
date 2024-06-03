package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)
	qtdWork := 10000
	for i := 0; i < qtdWork; i++ {
		go worker(i, data)
	}
	for i := 0; i < 1000000; i++ {
		data <- i
	}
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
