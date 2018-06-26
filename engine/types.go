package engine

import "Test/crawler_distributed/config"

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
	Requests   []Request
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

func (NilParser) Serialize() (name string, args interface{}) {
	return config.NilParser, nil
}

type FuncParser struct {
	parser  ParserFunc
	name	string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{ parser: p, name: name }
}


