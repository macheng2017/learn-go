package engine

// define request
// 将原来的url包装了一下,包含了url 和一个解析器
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

// 解析后的结果
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// 添加一个nilParser空函数
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
