package Controller

import(
	"github.com/kataras/iris"
	"../Service"
)

var indexservice = Service.IndexService{}
type IndexController struct {

}
func (c *IndexController) BeginRequest(ctx iris.Context) {}
func (c *IndexController) EndRequest(ctx iris.Context) {}
func (c *IndexController) Get(ctx iris.Context) {
	indexservice.Get(ctx)
}