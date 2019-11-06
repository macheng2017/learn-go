package mock

import "fmt"

type Retriever struct {
	Contents string
}

// 在一个struct中即实现了 Retriever接口中的方法也实现了 RetrieverPoster 接口中的方法,也没有说明实现了哪个接口,
// 这样可以在使用者是无感的情况下使用,感受下
func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}

// 实现一个官方定义的接口
func (r *Retriever) String() string {
	// 说实话很不喜欢node.js当中的一些写法, 比如这里可以直接访问 Contents golang中只能r.Contents,觉得这样做很好,
	// 这样可以显式的提示使用的是函数外部的变量,太灵活反而会带来麻烦
	return fmt.Sprintf("Retriever: {Contents=%s}", r.Contents)
}
