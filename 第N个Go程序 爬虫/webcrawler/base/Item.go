package base

//条目
type Item map[string]interface{}

//数据的接口
type Data interface {
	Valid() bool //数据是否有效
}

//数据是否有效
func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

func (resp *Response) Valid() bool {
	return resp.httpResp != nil && resp.httpResp.Body != nil
}

func (item Item) Valid() bool {
	return item != nil
}
