package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
)

func main() {
	// 包装成SimpleEngine{},将来如果改成ConcurrentEngine{}可以切换过去
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
