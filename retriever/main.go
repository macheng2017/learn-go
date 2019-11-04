package main

import "fmt"

type Retiever interface {
	Get(url string) string
}

func download(r Retiever) string {
	return r.Get("www.baidu.com")
}

func main() {
	// 当前程序不知道Retriever是什么东西,单词故意拼错的
	var r Retiever
	fmt.Println(download(r))
}
