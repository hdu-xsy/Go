package Controller

import (
	"github.com/kataras/iris"
	"../Entity"
	"../Service"
)
type UploadController struct{
}
var uploadpageservice = Service.UploadPage{}
func (c *UploadController) BeginRequest(ctx iris.Context) {
	if auth, _ := Entity.Sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.Redirect("/adminlogin")
		return
	}
}
func (c *UploadController) EndRequest(ctx iris.Context) {}
func (c *UploadController) Get(ctx iris.Context) {
	uploadpageservice.Get(ctx)
}

func Uploads(ctx iris.Context) {
	Service.Upload(ctx)
}