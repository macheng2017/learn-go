package engine

type ConcurrentEngine struct {
}

func (ConcurrentEngine) Run(seeds ...Request) {
	// 这里有两个问题:
	// 1. scheduler从哪里来?
	// 2. scheduler 是什么东西?
	for _, r := range seeds {
		scheduler.Submit(r)
	}
}
