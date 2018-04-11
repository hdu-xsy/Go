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