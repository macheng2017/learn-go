package main

import (
	"fmt"
	"learngo/interface1/infra"
)

func getRetriever() infra.Retrieve {
	return infra.Retrieve{}
}

func main() {
	// 这样看着比较像函数
	retrieve := getRetriever()
	fmt.Println(retrieve.Get("https://www.baidu.com"))
}
