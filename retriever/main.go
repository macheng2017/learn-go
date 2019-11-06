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

// 演示组合的用法
// 定义一个接口,(接口由使用者定义)
type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.baidu.com"

// 使用一个接口
func post(poster Poster) {

	poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

// 组合接口
type RetrieverPoster interface {
	Retriever
	Poster
	// other method
}

// 使用组合后的接口
func session(rp RetrieverPoster) string {
	rp.Post(url, map[string]string{
		"contents": "another faked baidu.com",
	})
	return rp.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is fack google.com"}
	r = &retriever
	inspect(r)
	r = &real2.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	inspect(r)

	// Type assertion
	realRetriever := r.(*real2.Retriever)
	fmt.Println(realRetriever.TimeOut)

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))

	//mockRetriever := r.(mock.Retriever)
	//fmt.Println(mockRetriever.Contents)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

}

func inspect(r Retriever) {
	fmt.Printf("%T %v \n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

//直接实现接口然后调用接口，使用者定义接口这个步骤会不会是多余的，存在肯定是有意义的，意义是什么？

//这个意义和其他语言的接口的意义是一样的。我们举了例子说我们调用了retriever.get(url)，那么retriever是啥呢？
//
//我们从解耦的方面来理解，如何让我们和retriever解耦？也就是对retriever是啥这个问题，规定的越少越好。如何才是最少呢？答案就是：retriever是一个可以调用get(string)方法的东西。
//
//这些就是背后的意义。retriever的使用者就是“我们”。“retriever是一个可以调用get(string)方法的东西”就是我们作为使用者定义出来的接口。而任何实现了get(string)方法的东西，不管是一个struct，还是一个其他类型别名，都可以作为retriever让我们来调用。
//
//这个意义可能的确需要一些项目经验才能真正理解。同学可以先打个问号，以后在项目中，尤其是开源项目中，看到了接口，而且也不一定是go语言的接口，我们再回过来看看这里retriever的例子，会有新的理解。
