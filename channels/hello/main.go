package main

func main() {
	channel := make(chan string) // Channel is empty
	go func() {
		channel <- "Hello World!" // Channel has a message is full
	}()
	message := <-channel //Channel is empty
	println(message)
}
