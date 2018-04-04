package Controller

import(
	"github.com/kataras/iris"
	"../Service"
)

var menuservice = Service.MenuService{}
type MenuController struct {

}
func (c *MenuController) BeginRequest(ctx iris.Context) {}
func (c *MenuController) EndRequest(ctx iris.Context) {}
func (c *MenuController) Get(ctx iris.Context) {
	menuservice.Get(ctx)
}
type ClassifyController struct {

}
var classifyservice = Service.ClassifySercice{}
func (c *ClassifyController) BeginRequest(ctx iris.Context) {}
func (c *ClassifyController) EndRequest(ctx iris.Context) {}
func (c *ClassifyController) Get(ctx iris.Context) {
	classifyservice.Get(ctx)
}