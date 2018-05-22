package Service

import (
	"github.com/kataras/iris"
	"../index"
	"../Entity"
)

//首页
type IndexService struct {

}

func (s *IndexService)Get(ctx iris.Context) {
	articleList := articledao.OrderByTime()
	comment := commentdao.OrderByTime()
	menulist := menudao.GetAll()
	articlelen := len(articledao.FindAllA())
	entity := Entity.Entity{ArticleList:articleList,CommentList:comment,MenuList:menulist}
	index.ListWriter(entity,ctx,ctx,articlelen)
}
