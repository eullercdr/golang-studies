package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	//length, err := f.WriteString("Hello, World!")
	length, err := f.Write([]byte("Escrevendo dados no arquivo!"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes\n", length)

	//Read the file
	file, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	//Read the file line a line using bufio and buffer
	// buffer is used to read the file in chunks
	// chunks are read in buffer and then printed
	f, err = os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	os.Remove("test.txt")
	f.Close()
}
