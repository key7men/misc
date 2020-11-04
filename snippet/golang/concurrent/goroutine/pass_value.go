package main

import (
	"fmt"
	"time"
)

type mail struct {
	from string
	to string
}

func (m mail) From(address string) mail {
	m.from = address
	return m
}

func (m mail) To(address string) mail {
	m.to = address
	return m
}

func (m mail) Send() {
	fmt.Printf("From: %s, To: %s\n", m.from, m.to)
}

func main() {
	m := mail{}
	for i := 0; i < 10; i++ {
		go func(val int){
			m.From(fmt.Sprintf("%d@gmail.com", val)).To(fmt.Sprintf("%d@gmail.com", val+1)).Send()
		}(i)
	}
	time.Sleep(1 * time.Second)
}