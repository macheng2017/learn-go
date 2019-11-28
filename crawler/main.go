package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
)

func main() {
	// 包装成SimpleEngine{},将来如果改成ConcurrentEngine{}可以切换过去
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
