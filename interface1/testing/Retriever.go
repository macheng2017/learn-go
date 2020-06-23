package testing

// 一个测试团队,写的测试消息
type Retriever struct {
}

func (Retriever) Get(url string) string {
	return "fake content"
}
