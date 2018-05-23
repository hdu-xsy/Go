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

//文章页面
type ArticleService struct {

}

func (s *ArticleService)Get(ctx iris.Context) {
	id,_ := strconv.ParseInt(ctx.Params().Get("id"),10,64)
	articleall := articledao.FindAllA()
	if int(id) > len(articleall) {
		ctx.Redirect("/404")
		return
	}
	_,_,article := articledao.Get(Entity.Article{Id:id})
	_,_,pre := articledao.Get(Entity.Article{Id:id-1})
	_,_,suc := articledao.Get(Entity.Article{Id:id+1})
	mid,_ := strconv.ParseInt(article.Menu,10,64)
	_,_,menu := menudao.Get(Entity.Menu{Id:mid})
	_,_,user := userdatadao.Get(Entity.UserData{Id:article.User})
	comment := commentdao.FindAll(ctx.Params().Get("id"))
	menulist := menudao.GetAll()
	entity := Entity.Entity{Article:article,UserData:user,Menu:menu,CommentList:comment,MenuList:menulist}
	Article.ContextWriter(entity,pre,suc,ctx,ctx)
}

//插入文章
type ArticleInsertService struct {

}
func (s *ArticleInsertService)Get(ctx iris.Context) {
	article := Entity.Article{}
	checkError(ctx.ReadForm(&article))
	article.User=1
	article.Time=time.Now()
	articledao.Insert(article)
	ctx.WriteString(" ")
	return
}

//文章修改列表
type ArticleModify struct {

}

func (s *ArticleModify)Get(ctx iris.Context) {
	article := articledao.FindAllA()
	var suc,max int
	page,_ := strconv.Atoi(ctx.Params().Get("page"))
	if len(article)%20 == 0 {
		max = len(article)/20
	} else {
		max = len(article)/20 + 1
	}
	if page>max || page == 0 {
		ctx.Redirect("/404")
		return
	}
	if page * 20 >= len(article) {
		suc = len(article)
	} else {
		suc = page * 20
	}
	for i,a := range article {
		id,_ := strconv.ParseInt(a.Menu,10,64)
		_,_,menu := menudao.Get(Entity.Menu{Id:id})
		article[i].Menu = menu.Name
	}
	userlist.ArticleListToWriter(article[(page-1)*20:suc],page,max,ctx)
}
func (s *ArticleModify)Update(ctx iris.Context) {
	id,_ := strconv.ParseInt(ctx.Params().Get("id"),10,64)
	_,_,article := articledao.Get(Entity.Article{Id:id})
	articleModify.ArticleToWriter(article,ctx)
}

//文章修改
type Articlemodify struct {

}
func (s *Articlemodify)Update(ctx iris.Context) {
	id,_ := strconv.ParseInt(ctx.PostValue("Id"),10,64)
	title := ctx.PostValue("Title")
	menu := ctx.PostValue("Menu")
	content := ctx.PostValue("Content")
	classify := ctx.PostValue("Classify")
	timenow := time.Now()
	article := Entity.Article{Id:id,User:1,Time:timenow,Title:title,Menu:menu,Classify:classify,Content:content}
	articledao.Update(article)
	ctx.Redirect("/articlemodifylist/1")
	return
}