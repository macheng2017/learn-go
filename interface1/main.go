package main

import (
	"fmt"
	"learngo/interface1/infra"
)

func getRetriever() retriever {
	return infra.Retriever{}
}

// 这里需要让我们的代码和逻辑一致
// 这里需要一个something ? 可以 Get

type retriever interface {
	Get(string) string
}

func main() {
	// 一个测试团队,需要测试这个方法,但是改动的地方太多了,很不满意
	// 这里对于动态语言解耦已经完成了,但对于静态语言(强类型语言)这里还需要其他操作

	var r retriever
	fmt.Println(r.Get("https://www.baidu.com"))
}
