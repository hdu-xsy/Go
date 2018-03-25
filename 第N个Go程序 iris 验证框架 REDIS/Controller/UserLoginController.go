package Controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"../Entity"
	"../Service"
	"../UserIndex"
)

//用户登录验证
type UserLoginController struct {
	Username string
}
var userloginservice = Service.UserLogin{}
var chatformservice = Service.ChatForm{}
func (c *UserLoginController) BeginRequest(ctx iris.Context) {
	userloginservice.BeginRequest(ctx)
	chatformservice.BeginRequest(ctx)
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
type UserIndexController struct {

}
func (c *UserIndexController) BeginRequest(ctx iris.Context) {
	userloginservice.BeginRequest(ctx)
}
func (c *UserIndexController) EndRequest(ctx iris.Context) {}
func (c *UserIndexController) Get(ctx iris.Context) {
	UserIndex.UserIndexWriter(ctx,ctx)
}