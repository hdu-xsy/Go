package Service

import (
	"github.com/kataras/iris"
	"../index"
	"../Entity"
)

type IndexService struct {

}

func (s *IndexService)Get(ctx iris.Context) {
	articleList := articledao.OrderByTime()
	comment := commentdao.OrderByTime()
	entity := Entity.Entity{ArticleList:articleList,CommentList:comment}
	index.ListWriter(entity,ctx,ctx)
}
