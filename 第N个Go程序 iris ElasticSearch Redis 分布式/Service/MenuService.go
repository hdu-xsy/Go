package Service

import (
	"github.com/kataras/iris"
	"strconv"
	"../Entity"
	"../Menu"
	"../Classify"
)


//栏目页面
type MenuService struct {

}

func (s *MenuService)Get(ctx iris.Context) {
	id,_ := strconv.ParseInt(ctx.Params().Get("id"),10,64)
	_,_,themenu := menudao.Get(Entity.Menu{Id:id})
	articlelist := articledao.FindAll(strconv.FormatInt(id,10))
	menulist := menudao.GetAll()
	m := articledao.Count()
	var suc,max int
	page64,_ := strconv.ParseInt(ctx.Params().Get("page"),10,64)
	page := int(page64)
	if len(articlelist)%20 == 0 {
		max = len(articlelist)/20
	} else {
		max = len(articlelist)/20 + 1
	}
	if page>max || page == 0 {
		ctx.Redirect("/404")
		return
	}
	if page * 20 >= len(articlelist) {
		suc = len(articlelist)
	} else {
		suc = page * 20
	}
	for i,a := range articlelist {
		id,_ := strconv.ParseInt(a.Menu,10,64)
		_,_,menu := menudao.Get(Entity.Menu{Id:id})
		articlelist[i].Menu = menu.Name
	}
	entity := Entity.Entity{ArticleList:articlelist[(page-1)*20:suc],Menu:themenu,MenuList:menulist}
	Menu.MenuWriter(id,m,entity,page,max,ctx,ctx)
}


//分类页面
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
	menulist := menudao.GetAll()
	var suc,max int
	page,_ := strconv.Atoi(ctx.Params().Get("page"))
	if len(al)%20 == 0 {
		max = len(al)/20
	} else {
		max = len(al)/20 + 1
	}
	if page>max || page == 0 {
		ctx.Redirect("/404")
		return
	}
	if page * 20 >= len(al) {
		suc = len(al)
	} else {
		suc = page * 20
	}
	for i,a := range al {
		id,_ := strconv.ParseInt(a.Menu,10,64)
		_,_,menu := menudao.Get(Entity.Menu{Id:id})
		al[i].Menu = menu.Name
	}
	Classify.MenuWriter(c,page,max,m,Entity.Entity{ArticleList:al[(page-1)*20:suc],MenuList:menulist},ctx,ctx)
}