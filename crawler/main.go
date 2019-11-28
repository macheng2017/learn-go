package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
)

func main() {
	// 包装成SimpleEngine{},将来如果改成ConcurrentEngine{}可以切换过去
	engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
	}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

//# learngo/crawler
//./main.go:13:3: cannot call pointer method on engine.ConcurrentEngine literal
//./main.go:13:3: cannot take the address of engine.ConcurrentEngine literal
