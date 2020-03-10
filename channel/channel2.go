package main

import (
	"fmt"
	"sync"
)

type work struct {
	id   int
	c    chan int
	done func()
}

func (work *work) doWorker2() {
	for n := range work.c {
		fmt.Printf("Worker %d received %c\n", work.id, n)
		work.done()
	}
}
func createWorker1(id int, wg *sync.WaitGroup) work {
	work := work{
		id: id,
		c:  make(chan int),
		// 仔细琢磨下这里这样写的好处
		// 通过函数式编程包装下之后将waitGroup抽象掉了放在createWorker中调用时配置

		done: func() {
			wg.Done()
		}}
	go work.doWorker2()
	return work
}

func chanDemo() {
	// 使用WaitGroup来解决循环等待问题
	var wg sync.WaitGroup
	var works [10]work
	for i := 0; i < 10; i++ {
		works[i] = createWorker1(i, &wg)
	}
	wg.Add(20)
	for i, work := range works {
		work.c <- 'a' + i
		//wg.Add(1) // 可以每次+1
	}

	for i, work := range works {
		work.c <- 'A' + i
	}
	wg.Wait()

}

func main() {
	chanDemo()
}
