package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("job 1 has done.")
		wg.Done()
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("job 2 has done")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("All job has done.")
}