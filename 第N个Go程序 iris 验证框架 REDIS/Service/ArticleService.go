package Service

import(
	"github.com/kataras/iris"
	"../Article"
	"../userlist"
	"../Entity"
	"../ArticleModify"
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

type ArticleModify struct {

}

func (s *ArticleModify)Get(ctx iris.Context) {
	article := articledao.FindAllA()
	for i,a := range article {
		id,_ := strconv.ParseInt(a.Menu,10,64)
		_,_,menu := menudao.Get(Entity.Menu{Id:id})
		article[i].Menu = menu.Name
	}
	userlist.ArticleListToWriter(article,ctx)
}
func (s *ArticleModify)Update(ctx iris.Context) {
	id,_ := strconv.ParseInt(ctx.Params().Get("id"),10,64)
	_,_,article := articledao.Get(Entity.Article{Id:id})
	articleModify.ArticleToWriter(article,ctx)
}

type Articlemodify struct {

}
func (s *Articlemodify)Update(ctx iris.Context) {
	id,_ := strconv.ParseInt(ctx.PostValue("Id"),10,64)
	title := ctx.PostValue("Title")
	menu := ctx.PostValue("Menu")
	content := ctx.PostValue("Content")
	article := Entity.Article{Id:id,User:1,Time:time.Now(),Title:title,Menu:menu,Classify:"nil",Content:Entity.CString(content)}
	articledao.Update(article)
	app.Logger().Println(strconv.FormatInt(id,10)+title+menu)
	ctx.Redirect("/articlemodify")
	return
}