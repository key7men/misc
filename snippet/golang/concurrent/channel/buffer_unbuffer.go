package main

import (
	"fmt"
	"time"
)

/*
 假如step1~step3去掉，会打印吗？
 buffer channel 是一个异步状态
 unbuffer channel 是一个同步状态
*/

func main () {
	c := make(chan bool, 1)						// step 1 这是一个buffer channel
	//c := make(chan bool)						// 这是一个unbuffer channel
	go func() {
		fmt.Println("This is goroutine")
		time.Sleep(1 * time.Second)
		<- c									// step 2
	}()
c <- true										// step 3
}
