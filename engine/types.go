package engine

// Request 请求
type Request struct {
	URL        string
	ParserFunc func([]byte) ParserResult
}

// ParserResult 解析
type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

// NilPaser 无
func NilPaser([]byte) ParserResult {
	return ParserResult{}
}
