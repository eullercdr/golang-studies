package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	//rabbit
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{id: i, Msg: "Hello RabbitMQ"}
			time.Sleep(time.Second)
			c1 <- msg
		}
	}()

	//kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{id: i, Msg: "Hello Kafka"}
			time.Sleep(2 * time.Second)
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("received from RabbitMQ Id %d - %+v\n", msg.id, msg.Msg)
		case msg := <-c2:
			fmt.Printf("received from Kafka Id %d - %+v\n", msg.id, msg.Msg)
		case <-time.After(3 * time.Second):
			println("timeout")
			// default:
			// 	println("no value ready, moving on")
		}
	}
}
