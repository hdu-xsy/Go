package Service

import(
	"github.com/kataras/iris"
	"../Article"
	"../Entity"
	"strconv"
	"time"
)
type ArticleService struct {

}

func (s *ArticleService)Get(ctx iris.Context) {
	id,_ := strconv.ParseInt(ctx.Params().Get("id"),10,64)
	_,_,article := articledao.Get(Entity.Article{Id:id})
	Article.ContextWriter(article,ctx)
}
type ArticleInsertService struct {

}
func (s *ArticleInsertService)Get(ctx iris.Context) {
	article := Entity.Article{}
	checkError(ctx.ReadForm(&article))
	articledao.Insert(Entity.Article{User:1,Time:time.Now(),Title:article.Title,Menu:article.Menu,Classify:"学习笔记",Content:article.Content})
	ctx.WriteString(" ")
}