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
	Menu.MenuWriter(entity,ctx,ctx)
}
