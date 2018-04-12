package Service

import (
	"github.com/kataras/iris"
	"strconv"
	"../Entity"
	"../Menu"
	"../Classify"
)

type MenuService struct {

}

func (s *MenuService)Get(ctx iris.Context) {
	id,_ := strconv.ParseInt(ctx.Params().Get("id"),10,64)
	_,_,themenu := menudao.Get(Entity.Menu{Id:id})
	articlelist := articledao.FindAll(strconv.FormatInt(id,10))
	entity := Entity.Entity{ArticleList:articlelist,Menu:themenu}
	m := articledao.Count()
	Menu.MenuWriter(m,entity,ctx,ctx)
}

type ClassifySercice struct {

}
func (s *ClassifySercice) BeginRequest(ctx iris.Context) {
	c := ctx.Params().Get("Classify")
	var articleList []Entity.Article
	articleList = articledao.GetClassify()
	for _,v := range articleList {
		if v.Classify == c {
			return
		}
	}
	ctx.Redirect("/404")
}
func (s *ClassifySercice) EndRequest(ctx iris.Context) {}
func (s *ClassifySercice) Get(ctx iris.Context) {
	c := ctx.Params().Get("Classify")
	var al []Entity.Article
	al = articledao.FindByClassify(c)
	m := articledao.Count()
	Classify.MenuWriter(m,Entity.Entity{ArticleList:al},ctx,ctx)
}