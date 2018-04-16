package downloader

import "../base"

//网页下载器的接口类型
type PageDownloader interface {
	Id() uint32 //获得Id
	Downloader(req base.Request) (*base.Response,error)  //根据请求的下载网页并返回响应
}

//ID生成器的接口类型
type IdGenertor interface {
	GetUint32()	uint32	//获得一个uint32类型的ID
}

//网页下载器池的接口类型
type PageDownloaderPool interface {
	Take() (PageDownloader,error)	//从池中取出一个网页下载器
	Return(dl PageDownloader) error //把一个网页下载器归还给池
	Total() uint32					//获得池的总容量
	Used() uint32					//获得正在被使用的网页下载器的数量
}

