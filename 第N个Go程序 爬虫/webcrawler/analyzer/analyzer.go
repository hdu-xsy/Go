package analyzer

import (
	"../base"
	"net/http"
)

//分析器的接口类型
type Analyzer interface {
	Id() uint32 //获得ID
	Analyzer(
		resParsers []ParseResponse,
		resp base.Response)([]base.Data,[]error)
}

//被用于解析HTTP响应的函数类型
type ParseResponse func(httpResp *http.Response,respDepth uint32) ([]base.Data,[]error)

//分析器池的接口类型
type AnalyzerPool interface {
	Take() (Analyzer,error)	//从池中取出一个分析器
	Return (analyzer Analyzer) error	//把一个分析器归还给池
	Total() uint32	//获得池的总容量
	Used()	uint32	//获得正在被使用的分析器的数量
}