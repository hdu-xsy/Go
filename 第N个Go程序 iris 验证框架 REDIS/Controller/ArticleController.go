package Controller

import (
	"github.com/kataras/iris"
	"../Service"
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