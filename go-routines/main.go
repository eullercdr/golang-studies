package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Tarefa: %s - %d\n", name, i)
		time.Sleep(1 * time.Second)
	}
}

// thread 1
func main() {
	go task("Tarefa A")
	go task("Tarefa B")
	go func () {
		for i := 0; i < 10; i++ {
			fmt.Printf("Tarefa: %s - %d\n", "Tarefa C", i)
			time.Sleep(1 * time.Second)
		}
	}
	time.Sleep(15 * time.Second)
}
