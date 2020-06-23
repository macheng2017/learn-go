package main

import (
	"fmt"
	"learngo/interface1/infra"
)

func getRetriever() infra.Retrieve {
	return infra.Retrieve{}
}

func main() {
	// 把retriever 的类型现式的写出来
	var retrieve infra.Retrieve = getRetriever()
	fmt.Println(retrieve.Get("https://www.baidu.com"))
}
