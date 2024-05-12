package main

import (
	"fmt"
	"io"
	"net/http"
)

// func main() {
// 	fmt.Println("This will be executed first")
// 	defer fmt.Println("This will be executed last")
// 	fmt.Println("This will be executed second")
// }

func main() {
	request, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	//defer is used to close the connection after the function is executed
	//defer is executed in LIFO order
	//LILO - Last In Last Out
	defer request.Body.Close()
	res, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
