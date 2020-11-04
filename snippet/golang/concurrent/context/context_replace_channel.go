package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// context.Background() 基本就是ROOT节点
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
				case <- ctx.Done():
					fmt.Println("get cancel signal")
					return
				default:
					fmt.Println("still working")
					time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	cancel()
	// 比较一下无下面这句SLEEP的输出有何异同，"get cancel signal"是否打印？打印多少次呢？
	time.Sleep(1* time.Second)
	fmt.Print("All job has done")
}