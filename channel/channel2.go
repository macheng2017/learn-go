package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func worker1(id int, c chan int) {
	//注意 这里开了10个goroutine向文件中写入数据
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for {
		fmt.Printf("Worker %d received %d\n", id, <-c)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(writer, "Worker %d received %d\n", id, <-c)
	}
}
func chanDemo() {
	// 建立10个worker并为每个worker配置一个自己的channel
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker1(i, channels[i])
	}

	go func() {
		var i, count int
		for {
			i++
			count++
			if i%10 == 0 {
				i = 0
			}
			channels[i] <- count
		}
	}()

	// 1. 为什么把生产者放到消费者前面会导致死锁问题
	// 2. 为什么将生产者放到goroutine当中会没有结果
	// 3. 用-race测试下,测试之后,可以正常打印,猜测问题是,时间太短了两个goroutine没时间启动就随者main goroutine结束了
	// 4. 解决方式是使用sleep函数,让main函数等一会
	//c <- 1
	//c <- 2
}

func main() {
	chanDemo()
	time.Sleep(time.Second)
} //fatal error: all goroutines are asleep - deadlock!
