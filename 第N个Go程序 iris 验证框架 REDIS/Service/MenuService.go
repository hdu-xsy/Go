package Service

import (
	"github.com/kataras/iris"
	"strconv"
	"../Entity"
	"../Menu"
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
func (s *ClassifySercice) Get(ctx iris.Context) {
	c := ctx.Params().Get("Classify")
	var al []Entity.Article
	al = articledao.FindByClassify(c)
	m := articledao.Count()
	Menu.MenuWriter(m,Entity.Entity{ArticleList:al},ctx,ctx)
}