package engine

type ConcurrentEngine struct {
	// scheduler在这个结构体中定义了
	Scheduler
}

// Scheduler 是一个接口
type Scheduler interface {
	Submit(r Request)
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	// 这里有两个问题:
	// 1. scheduler从哪里来?
	// 2. scheduler 是什么东西? scheduler在这个ConcurrentEngine结构体中定义了
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
}
