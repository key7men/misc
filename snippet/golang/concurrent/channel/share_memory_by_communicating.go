package main

import "fmt"

// 用 通讯信道 (channel) 来实现不同进程间的数据流转
func addByShareCommunicate(n int) (ints []int) {
	channel := make(chan int, n)

	for i := 0; i < n; i++ {
		go func(channel chan<- int, value int) {
			channel <- value
		}(channel, i)
	}

	// for一直在执行（while），知道长度为n的时候跳出
	for i := range channel {
		ints = append(ints, i)

		if len(ints) == n {
			break
		}
	}

	// 使用完关掉
	close(channel)

	return
}

func main() {
	output := addByShareCommunicate(10)
	fmt.Println(output)
}

