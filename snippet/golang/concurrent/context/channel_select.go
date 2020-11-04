package main

import (
	"fmt"
	"time"
)

/*
 以下示例中，通过goroutine
*/
func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
				case <-stop:
					fmt.Println("got the stop signal")
					return
				default:
					fmt.Println("sub job still working")
					time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("send stop signal")
	stop<- true
	time.Sleep(5 * time.Second)
	fmt.Println("Finished")

}