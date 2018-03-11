package Controller

import(
	"../Service"
	"github.com/kataras/iris"
	"../Entity"
)

var AdminLoginService = Service.AdminLogin{}
type AdminLoginController struct {
}
func (c *AdminLoginController) BeginRequest(ctx iris.Context) {
	if auth, _ := Entity.Sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.Redirect("/adminlogin")
		return
	}
}
func (c *AdminLoginController) EndRequest(ctx iris.Context) {}
func (c *AdminLoginController) Get(ctx iris.Context) {
	AdminLoginService.Get(ctx)
}