package main

import (
	"fmt"
	"learngo/interface1/testing"
)

func getRetriever() testing.Retriever {
	return testing.Retriever{}
}

func main() {
	// 一个测试团队,需要测试这个方法,但是改动的地方太多了,很不满意
	var retrieve testing.Retriever = getRetriever()
	fmt.Println(retrieve.Get("https://www.baidu.com"))
}
