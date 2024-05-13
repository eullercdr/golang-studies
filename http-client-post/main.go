package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	jsonVar := []byte(`{"name":"Euller"}`)
	resp, err := client.Post("http://httpbin.org/post", "application/json", bytes.NewBuffer(jsonVar))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
