package main

import (
	"fmt"
)

type work struct {
	c    chan int
	done chan bool
}

func (work *work) doWorker2(id int) {
	for n := range work.c {
		fmt.Printf("Worker %d received %c\n", id, n)
		// 通过channel通知外部的
		work.done <- true
	}
}
func createWorker1(id int) work {
	work := work{make(chan int), make(chan bool)}
	go work.doWorker2(id)
	return work
}

// 这样改好之后有个新问题: 打印是顺序执行的,也就是一个goroutine处理完之后通知你完成,
// 然后再换下一个goroutine,这样每一个goroutine只用了一次,并且会阻塞其他goroutine执行
func chanDemo() {
	var works [10]work
	for i := 0; i < 10; i++ {
		works[i] = createWorker1(i)
	}

	for i := 0; i < 10; i++ {
		works[i].c <- 'a' + i
		// 这里阻塞了多个goroutine执行,需要排队等待通过,这样就没有了多协程并发执行的优势了
		<-works[i].done
	}

	for i := 0; i < 10; i++ {
		works[i].c <- 'A' + i
		<-works[i].done
	}
}

func main() {
	chanDemo()
}
