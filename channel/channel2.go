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
		work.done <- true
	}
}
func createWorker1(id int) work {
	work := work{make(chan int), make(chan bool)}
	go work.doWorker2(id)
	return work
}

func chanDemo() {
	var works [10]work
	for i := 0; i < 10; i++ {
		works[i] = createWorker1(i)
	}

	for i, work := range works {
		work.c <- 'a' + i
		//这个错误的原因是,所有的channel [发送] 都是阻塞式(block),当发送一个信息,必须有人去收,
		//的当这个循环完毕之后,消费者(doWorker2函数)对10个channel都发过了一次通信(发送了一次done),但是处于没人接受状态
		// 而这时程序成功执行了for循环小写部分,然后会进入下个循环,当消费者再次处理完第一个并又对channel发送一次通信,
		//这时候就报错了,出现deadlock error (这个错误是循环等待导致的)
	}
	for _, worker := range works {
		<-worker.done
	}
	// 这样是第二种方式,先并发的执行完第一组,然后这边收到通信结束
	// 在执行下一组
	for i, work := range works {
		work.c <- 'A' + i
	}

	for _, worker := range works {
		<-worker.done
	}
}

func main() {
	chanDemo()
}
