package Service

import (
	"github.com/kataras/iris"
	"../index"
)

type IndexService struct {

}

func (s *IndexService)Get(ctx iris.Context) {
	articleList := articledao.OrderByTime()
	comment := commentdao.OrderByTime()
	index.ListWriter(articleList,comment,ctx,ctx)
}
