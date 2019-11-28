package engine

import "fmt"

type ConcurrentEngine struct {
	// scheduler在这个结构体中定义了
	Scheduler   Scheduler
	WorkerCount int
}

// Scheduler 是一个接口
type Scheduler interface {
	Submit(r Request)
	// 配置WorkerChan接口
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// 这里有两个问题:
	// 1. scheduler从哪里来?
	// 2. scheduler 是什么东西? scheduler在这个ConcurrentEngine结构体中定义了

	// 使用 channel进行通信
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	// 接收out并打印
	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("Got item: %v", item)
		}
		// 将所有的Request 再放进Scheduler
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, e := worker(request)
			if e != nil {
				continue
			}
			out <- result
		}
	}()
}
