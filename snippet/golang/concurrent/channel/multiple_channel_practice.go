package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"
)

/*
 任务要求：
	1. 通过goroutine跑20个job
	2. 通过outChan显示各个job的完成状况
	3. 通过errChan显示job发生的错误，并跳出main func
	4. 通过finishChan通知所有job已经完成
	5. 设定Timeout机制（1秒内要完成所有job，否则跳出）
*/
func main() {
	wg := sync.WaitGroup{}
	outchan := make(chan string, 100)
	errchan := make(chan error, 100)
	finishChan := make(chan struct{}) 										// 使用结构体的形式，不占内存

	wg.Add(20)

	for i := 0; i < 20; i++ { 												// 满足条件1
		go func(val int, wg *sync.WaitGroup, out chan<- string, exp chan<- error) { 			// 定义out为只写
			time.Sleep(1 * time.Second)
			out <- fmt.Sprintf("finished job id: %d", val)
			if val == 15 {
				exp <- errors.New("job 15 was wrong!")
			}
			wg.Done()
		}(i, &wg, outchan, errchan)
	}

	// 此线程在后台等待
	go func() {
		wg.Wait()
		close(finishChan)
	}()

	Loop:
		for {
			select {
				case out := <-outchan:       									// 满足条件2
					fmt.Println(out)
				case err := <-errchan:											// 满足条件3
					fmt.Println(err)
					break Loop
				case <-finishChan:
					break Loop 													// 注意这个跳出标签的写法（满足条件4）
				case <-time.After(10000 * time.Millisecond):
					fmt.Println("Timeout")								// 满足条件5
					break Loop
				}
		}
}