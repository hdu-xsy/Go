package base

import "net/http"

//响应
type Response struct {
	httpResp *http.Response
	depth uint32
}

//创建新的响应
func NewResponse(httpResp *http.Response,depth uint32) *Response {
	return &Response{httpResp:httpResp,depth:depth}
}

//获取HTTP响应
func (resp *Response) HttpResp() *http.Response {
	return resp.httpResp
}

//获取深度值
func (resp *Response) Depth() uint32 {
	return resp.depth
}