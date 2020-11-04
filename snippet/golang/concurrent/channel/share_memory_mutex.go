package main

import (
	"fmt"
	"sync"
)

/*
 定义channel的时候，强烈建议明确channel是只读/只写/读写
 Write(chan<- int)
 Read(<-chan int)
 ReadWrite(chan int)
*/

// 用 共享内存（变量的方式） 来实现不同进程间（函数间）的数据流转
func addByShareMemWithMutex(n int) []int {
	var ints []int
	var wg sync.WaitGroup
	var mux sync.Mutex

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int){
			defer wg.Done()
			mux.Lock()
			ints = append(ints, i)
			mux.Unlock()
		}(i)
	}

	wg.Wait()
	return ints
}

func main() {

	output := addByShareMemWithMutex(10)
	fmt.Println(output)
}