package Controller

import (
	"github.com/kataras/iris"
	"../Service"
	"../Entity"
)

var articleservice = Service.ArticleService{}
type ArticleController struct {

}
func (c *ArticleController) BeginRequest(ctx iris.Context) {}
func (c *ArticleController) EndRequest(ctx iris.Context) {}
func (c *ArticleController) Get(ctx iris.Context) {
	articleservice.Get(ctx)
}

var articlinsert = Service.ArticleInsertService{}

func ArticleInsertController(ctx iris.Context) {
	articlinsert.Get(ctx)
}

type ArticleListController struct {

}
var articleService = Service.ArticleModify{}
func (c *ArticleListController) BeginRequest(ctx iris.Context) {
	if auth, _ := Entity.Sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.Redirect("/adminlogin")
		return
	}
}
func (c *ArticleListController) EndRequest(ctx iris.Context) {}
func (c *ArticleListController) Get(ctx iris.Context) {
	articleService.Get(ctx)
}

type ArticInsertController struct{
}
func (c *ArticInsertController) BeginRequest(ctx iris.Context) {
	if auth, _ := Entity.Sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.Redirect("/adminlogin")
		return
	}
}
func (c *ArticInsertController) EndRequest(ctx iris.Context) {}
func (c *ArticInsertController) Get(ctx iris.Context) {
	ctx.View("articleinsert.html")
}

type ArticleModifyController struct{
}
func (c *ArticleModifyController) BeginRequest(ctx iris.Context) {
	if auth, _ := Entity.Sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.Redirect("/adminlogin")
		return
	}
}
func (c *ArticleModifyController) EndRequest(ctx iris.Context) {}
func (c *ArticleModifyController) Get(ctx iris.Context) {
	articleService.Update(ctx)
}
var articlemodify = Service.Articlemodify{}
func ArticlemodifyController(ctx iris.Context) {
	articlemodify.Update(ctx)
}

