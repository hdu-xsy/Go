package middleware

import "../base"

//通道管理器的接口类型
type ChannelManager interface {
	//初始化通道管理器
	//参数channelLen代表通道管理器中的各类通道的初始长度
	//参数reset指明是否重新初始化通道管理器
	Init(channelLen uint,reset bool) bool
	//关闭通道管理器
	Close() bool
	//获取请求传输通道
	ReqChan() (chan base.Request,error)
	//获取条目传输通道
	ItemChan() (chan error,error)
	//获取通道长度值
	ChannelLen() uint
	//获取通道管理器的状态
	Status() ChannelManagerStatus
	//获取摘要信息
	Summary() string
}

//被用来表示通道管理器的状态的类型
type ChannelManagerStatus uint8

const(
	CHANNEL_MANAGER_STATUS_UNINTIALIZED ChannelManagerStatus = 0	//未初始化状态
	CHANNEL_MANAGER_STATUS_INITIALIZED	ChannelManagerStatus = 1	//已初始化状态
	CHANNEL_MANAGER_STATUS_CLOSED		ChannelManagerStatus = 2	//已关闭状态
)
