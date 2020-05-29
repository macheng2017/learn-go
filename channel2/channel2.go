package main

import (
	"fmt"
	"sync"
)

type work struct {
	c  chan int
	wg *sync.WaitGroup
}

func (work *work) doWork(i int, wg *sync.WaitGroup) {

	for c := range work.c {
		fmt.Printf("channel= %v %c\n", i, c)
		wg.Done()
	}

}

func createWork(i int, wg *sync.WaitGroup) work {
	work := work{make(chan int), wg}
	go work.doWork(i, wg)
	return work
}

func chanDemo() {
	var works [10]work
	wg := sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 10; i++ {
		works[i] = createWork(i, &wg)
	}
	for i := 0; i < 10; i++ {
		// debug了一下,发现这个结构很精巧
		// 1.先是消费者那边work.c阻塞并等待接收
		// 2.work.c收到消息以后,work.done发送一条消息
		// 3. 这边的work.done等待接收
		// 这两边互相阻塞
		works[i].c <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		works[i].c <- 'A' + i
	}
	wg.Wait()

}

func main() {
	chanDemo()
}
