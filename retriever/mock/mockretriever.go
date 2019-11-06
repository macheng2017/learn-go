package mock

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
