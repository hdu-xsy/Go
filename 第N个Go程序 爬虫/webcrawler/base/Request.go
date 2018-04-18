package base

import "net/http"

//请求
type Request struct {
	httpReq *http.Request
	depth uint32
}

//创建新的请求
func NewRequest(httpReq *http.Request,depth uint32) *Request{
	return &Request{httpReq:httpReq,depth:depth}
}

//获取HTTP请求
func (req *Request) HttpReq() *http.Request {
	return req.httpReq
}

//获取深度值
func (req *Request) Depth() uint32 {
	return req.depth
}

