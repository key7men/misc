package main

import (
	"fmt"
	"time"
)

type email struct {
	from string
	to string
}

func (e *email) From(address string) {
	e.from = address
}

func (e *email) To(address string) {
	e.to = address
}

func (e *email) Send() {
	fmt.Printf("From: %s, To: %s\n", e.from, e.to)
}

func main() {
	//e := email{}                         // 在循坏外部声明，会导致gofunc复写
	for i := 0; i < 10; i++ {
		go func(val int) {
			e := email{}                   // 在内部声明会实例化多个email
			e.From(fmt.Sprintf("%d@gmail.com", val))
			e.To(fmt.Sprintf("%d@gmail.com", val+1))
			e.Send()
		}(i)
	}

	time.Sleep(1 * time.Second)
}
