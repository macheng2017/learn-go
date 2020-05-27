package main

import (
	"fmt"
)

type work struct {
	c    chan int
	done chan bool
}

func (work *work) doWork(i int) {

	//file, err := os.OpenFile("1.txt", os.O_RDWR|os.O_CREATE, 0755)
	//if err != nil {
	//	fmt.Print(err)
	//}
	//defer file.Close()
	//writer := bufio.NewWriter(file)
	//defer writer.Flush()

	for {
		for n := range work.c {
			fmt.Printf("channel= %v %c\n", i, n)
			//fmt.Fprintf(writer, "channel=%v %c\n", i, <-c)
			work.done <- true
		}

	}
}

func createWork(i int) work {
	work := work{make(chan int), make(chan bool)}
	go work.doWork(i)
	return work
}

func chanDemo() {
	var works [10]work

	for i := 0; i < 10; i++ {
		works[i] = createWork(i)
	}
	for i := 0; i < 10; i++ {
		// debug了一下,发现这个结构很精巧
		// 1.先是消费者那边work.c阻塞并等待接收
		// 2.work.c收到消息以后,work.done发送一条消息
		// 3. 这边的work.done等待接收
		// 这两边互相阻塞
		works[i].c <- 'a' + i
		//<-works[i].done
	}

	for i := 0; i < 10; i++ {
		works[i].c <- 'A' + i
		//<-works[i].done
	}

}

func main() {
	chanDemo()
}
