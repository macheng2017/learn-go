package main

import (
	"fmt"
	"time"
)

// 定义channel只能收数据的修饰(告诉调用者只能收数据)
func createWorker1(id int) chan<- int {
	// 在worker内部建立channel,并返回出去给别人使用
	c := make(chan int)

	go func() {
		// 为什么openFile定义到goroutine外面就没有写入了呢?
		// 还有一个问题就是,当数据量很小的时候不能写入文件这是怎么回事?
		//file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0755)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//defer file.Close()
		//writer := bufio.NewWriter(file)
		//defer writer.Flush()
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
			//fmt.Fprintf(writer, "Worker %d received %c\n", id, <-c)
		}
	}()
	return c
}
func chanDemo() {
	// 定义channel只能收数据的修饰(告诉调用者只能收数据)
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker1(i)
	}

	for i := 0; i < 10; i++ {
		// 只能向上面定义的channel中发数据
		channels[i] <- 'a' + i
		// 这里如果向外发数据就会报错
		// n := <-channels[i] //invalid operation: <-channels[i] (receive from send-only type chan<- int)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	// 1. 为什么把生产者放到消费者前面会导致死锁问题
	// 2. 为什么将生产者放到goroutine当中会没有结果
	// 3. 用-race测试下,测试之后,可以正常打印,猜测问题是,时间太短了两个goroutine没时间启动就随者main goroutine结束了
	// 4. 解决方式是使用sleep函数,让main函数等一会
	//c <- 1
	//c <- 2
}

func main() {
	chanDemo()
	time.Sleep(time.Second * 2)
} //fatal error: all goroutines are asleep - deadlock!
