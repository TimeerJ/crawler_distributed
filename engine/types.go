package engine

// ParserFunc解析函数
type ParserFunc func(contents []byte, url string) ParseResult

// Parser解析功能接口
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

// 请求体
type Request struct {
	Url    string
	Parser Parser
}

// 解析请求体
type ParseResult struct {
	Request   []Request
	Items     []Item
}

// 消息队列
type Item struct {
	Url		string
	Type	string
	Id		string
	Payload	interface{}
}

// 无解析
type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return  ParseResult{}
}
