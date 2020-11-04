package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, "worker 1")
	go worker(ctx, "worker 2")
	go worker(ctx, "worker 3")

	time.Sleep(5 * time.Second)
	cancel()
	fmt.Println("All workers has been canceled")

}

func worker(ctx context.Context, val string) {
	go func() {
		for {
			select {
			case <- ctx.Done():
				fmt.Println("get cancel signal by" + val)
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Println(val + " still working")
			}
		}
	}()
}