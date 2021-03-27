package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go work(ctx, "node 1")
	go work(ctx, "node 2")
	go work(ctx, "node 3")

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("All Job Done.")
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name + " stop job.")
			return
		default:
			fmt.Println(name + " still work")
			time.Sleep(1 * time.Second)
		}
	}
}
