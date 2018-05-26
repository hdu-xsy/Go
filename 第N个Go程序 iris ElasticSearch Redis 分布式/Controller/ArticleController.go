package Controller

import (
	"github.com/kataras/iris"
	"../Service"
	"../Entity"
)

var articlePageService = Service.ArticlePageService{}
type ArticleController struct {

}
func (c *ArticleController) BeginRequest(ctx iris.Context) {
	articlePageService.BeginRequest(ctx)
}
func (c *ArticleController) EndRequest(ctx iris.Context) {}
func (c *ArticleController) Get(ctx iris.Context) {
	articlePageService.Get(ctx)
}

var articlInsertService = Service.ArticleInsertService{}

func ArticleInsertController(ctx iris.Context) {
	articlInsertService.Get(ctx)
}

type ArticleModifyListController struct {

}
var articleModifyListService = Service.ArticleModifyListPage{}
func (c *ArticleModifyListController) BeginRequest(ctx iris.Context) {
	if auth, _ := Entity.Sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.Redirect("/adminlogin")
		return
	}
}
func (c *ArticleModifyListController) EndRequest(ctx iris.Context) {}
func (c *ArticleModifyListController) Get(ctx iris.Context) {
	articleModifyListService.Get(ctx)
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

var articleModifyPageService = Service.ArticleModifyPage{}
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
	articleModifyPageService.Update(ctx)
}

var articleModify = Service.ArticleModify{}
func ArticlemodifyController(ctx iris.Context) {
	articleModify.Update(ctx)
}

