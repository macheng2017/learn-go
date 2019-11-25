package engine

// define request
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// 添加一个nilparser空函数
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
