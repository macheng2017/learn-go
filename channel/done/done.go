package main

import (
	"fmt"
	"sync"
)

// 定义一个struct
type worker struct {
	in chan int // 向内部发送数据
	// 使用函数式编程将wg抽象一下
	done func()
}

// 需要从外部传入一个 wg 只能有一个(共用一个)
var createWorker = func(id int, wg *sync.WaitGroup) worker {
	// 初始化struct包含两个channel,这个很像javabean
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}

	go doWork(id, w)

	return w
}

var doWork = func(id int, w worker) {

	for n := range w.in {
		fmt.Printf("receive value via channel id : %d value %c\n", id, n)
		// 这里有个相似的地方,当时开了个goroutine解决了这个问题
		w.done()
	}
}

func channelDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	// 加入等待20个任务
	wg.Add(20)

	// 向channel发数据
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	channelDemo()
}

//receive value via channel id : 5 value f
//receive value via channel id : 4 value e
//receive value via channel id : 7 value h
//receive value via channel id : 0 value a
//receive value via channel id : 2 value c
//receive value via channel id : 1 value b
//receive value via channel id : 2 value C
//receive value via channel id : 9 value j
//receive value via channel id : 0 value A
//receive value via channel id : 1 value B
//receive value via channel id : 3 value d
//receive value via channel id : 3 value D
//receive value via channel id : 4 value E
//receive value via channel id : 8 value i
//receive value via channel id : 5 value F
//receive value via channel id : 6 value g
//receive value via channel id : 6 value G
//receive value via channel id : 9 value J
//receive value via channel id : 8 value I
//receive value via channel id : 7 value H
