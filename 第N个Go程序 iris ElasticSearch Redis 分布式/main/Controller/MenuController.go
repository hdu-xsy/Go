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
func (c *ClassifyController) BeginRequest(ctx iris.Context) {
	classifyservice.BeginRequest(ctx)
}
func (c *ClassifyController) EndRequest(ctx iris.Context) {}
func (c *ClassifyController) Get(ctx iris.Context) {
	classifyservice.Get(ctx)
}

type MenuModifyPageController struct {

}
var menumodifypageservice = Service.MenuModifyPageSercice{}
func (c *MenuModifyPageController) BeginRequest(ctx iris.Context) {
	menumodifypageservice.BeginRequest(ctx)
}
func (c *MenuModifyPageController) EndRequest(ctx iris.Context) {}
func (c *MenuModifyPageController) Get(ctx iris.Context) {
	menumodifypageservice.Get(ctx)
}

var menuModify = Service.MenuModify{}
func MenuModifyController(ctx iris.Context) {
	menuModify.Update(ctx)
}

var h1Modify = Service.H1Modify{}
func H1ModifyController(ctx iris.Context) {
	h1Modify.Update(ctx)
}

var acModify = Service.ACModify{}
func ACModifyController(ctx iris.Context) {
	acModify.Update(ctx)
}
var menuInsert = Service.MenuInsert{}
func MenuInsertController(ctx iris.Context) {
	menuInsert.Update(ctx)
}
var menuDelete = Service.MenuDelete{}
func MenuDeleteController(ctx iris.Context) {
	menuDelete.Update(ctx)
}