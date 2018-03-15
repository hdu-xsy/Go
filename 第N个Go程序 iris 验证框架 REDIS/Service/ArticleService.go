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
	mid,_ := strconv.ParseInt(article.Menu,10,64)
	_,_,menu := menudao.Get(Entity.Menu{Id:mid})
	Article.ContextWriter(article,menu,ctx)
}
type ArticleInsertService struct {

}
func (s *ArticleInsertService)Get(ctx iris.Context) {
	article := Entity.Article{}
	checkError(ctx.ReadForm(&article))
	article.User=1
	article.Time=time.Now()
	article.Classify="学习笔记"
	articledao.Insert(article)
	ctx.WriteString(" ")
	return
}