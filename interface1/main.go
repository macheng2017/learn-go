package main

import (
	"fmt"
	"learngo/interface1/infra"
)

func main() {
	retrieve := infra.Retrieve{}
	fmt.Println(retrieve.Get("https://www.baidu.com"))
}
