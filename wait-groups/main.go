package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Tarefa: %s - %d\n", name, i)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	go task("Tarefa A", &waitGroup)
	go task("Tarefa B", &waitGroup)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Tarefa: %s - %d\n", "Tarefa C", i)
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()
}
