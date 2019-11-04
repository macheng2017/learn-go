package main

import (
	"fmt"
	"learngo/retriever/mock"
	real2 "learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is fack google.com"}
	fmt.Printf("%T %v \n", r, r)
	r = real2.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	//# learngo/retriever
	//./main.go:22:4: cannot use real.Retriever literal (type real.Retriever) as type Retriever in assignment:
	//real.Retriever does not implement Retriever (Get method has pointer receiver)
	fmt.Printf("%T %v \n", r, r)
	//output:
	//mock.Retriever {this is fack google.com}
	//real.Retriever { 1m0s}

	//fmt.Println(download(mock.Retriever{"this is fack google.com"}))
	//fmt.Println(download(real2.Retriever{}))
}

//直接实现接口然后调用接口，使用者定义接口这个步骤会不会是多余的，存在肯定是有意义的，意义是什么？

//这个意义和其他语言的接口的意义是一样的。我们举了例子说我们调用了retriever.get(url)，那么retriever是啥呢？
//
//我们从解耦的方面来理解，如何让我们和retriever解耦？也就是对retriever是啥这个问题，规定的越少越好。如何才是最少呢？答案就是：retriever是一个可以调用get(string)方法的东西。
//
//这些就是背后的意义。retriever的使用者就是“我们”。“retriever是一个可以调用get(string)方法的东西”就是我们作为使用者定义出来的接口。而任何实现了get(string)方法的东西，不管是一个struct，还是一个其他类型别名，都可以作为retriever让我们来调用。
//
//这个意义可能的确需要一些项目经验才能真正理解。同学可以先打个问号，以后在项目中，尤其是开源项目中，看到了接口，而且也不一定是go语言的接口，我们再回过来看看这里retriever的例子，会有新的理解。
