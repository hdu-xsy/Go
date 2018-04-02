package Controller

import (
	"github.com/kataras/iris"
	"../Service"
)
type DownloadPageController struct {

}
var downloadpage = Service.DownloadPage{}
func (c *DownloadPageController) Get(ctx iris.Context) {
	downloadpage.Get(ctx)
}
func DownloadFile(ctx iris.Context) {
	Service.DownloadFile(ctx)
}
