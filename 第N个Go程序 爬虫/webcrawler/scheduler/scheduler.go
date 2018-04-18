package scheduler

import(
	"../analyzer"
	"../itempipeline"
	"net/http"
)
//调度器的接口类型
type Scheduler interface {
	//启动调度器
	//调用该方法会使调度器创建和初始化各个组件。在此之后，调度器会激活爬取流程的执行
	//参数channelLen被用来指定数据传输通道的长度
	//参数crawlDepth代表了需要被爬取的网页的最大深度值。深度大与此值的网页会被忽略
	//参数httpClientGenerator代表的是被用来生成HTTP客户端的函数
	//参数respParsers的值应为需要被置入条目处理管道中的条目处理器的序列
	//参数firstHttpReq即代表首次请求,调度器会以此为起始点开始执行爬取流程
	Start(channelLen uint,
		poolSize uint32,
		crawlDepth uint32,
		httpClientGenerator GenHttpClient,
		respParsers []analyzer.ParseResponse,
		itemProcessors	[]itempipeline.ProcessItem,
		firstHttpReq	*http.Request) (err error)
	//调用该方法会停止调度器的运行.或有处理模块执行过程中出现的所有错误都会被发送到该通道
	//若该方法的结果值为Nil,则说明错误通道不可用或已被停止
	ErrorChan() <-chan error
	//判断所有处理模块是否都处于空闲状态.
	Idle() bool
	//获得摘要信息
	Summary(prefix string) SchedSummary
}

//被用来生成HTTP客户端的函数类型
type GenHttpClient func() *http.Client

//调度器摘要信息的接口类型
type SchedSummary interface {
	String() string	//获得摘要信息的一般表示
	Detail() string	//获得摘要信息的详细表示
	Same(other SchedSummary) bool	//判断是否与另一份摘要信息相同
}