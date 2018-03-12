package Controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"../Entity"
)

//用户登录验证
type UserLoginController struct {
	Username string
}
func (c *UserLoginController) BeginRequest(ctx iris.Context) {
	if auth, _ := Entity.Sess.Start(ctx).GetBoolean("userauthenticated"); !auth {
		ctx.Redirect("/login")
		return
	}
	c.Username = Entity.Sess.Start(ctx).GetString("Username")
}
func (c *UserLoginController) EndRequest(ctx iris.Context) {}
func (c *UserLoginController) Get() mvc.View {
	return mvc.View{
		Name: "chatform",
		Data: iris.Map{
			"Username": c.Username,
		},
	}
}
