package main

import (
	"context"
	"fmt"
)

// passed context to bookHotel, token is passed as key-value pair
func main() {
	ctx := context.WithValue(context.Background(), "token", "password")
	bookHotel(ctx)
}

// context always should be passed as first argument
func bookHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)
}
