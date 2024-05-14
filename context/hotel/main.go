package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()                            //will be run in thread main
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second) //context will be cancelled after 3 seconds
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout exceeded")
		return
	case <-time.After(1 * time.Second):
		fmt.Println("Hotel booking done")
	}
}
