package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AdminLoginController struct {
	Account string
}
func (c *AdminLoginController) BeginRequest(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.Redirect("/adminlogin")
		return
	}
	session := sess.Start(ctx)
	c.Account = session.GetString("Account")
}
func (c *AdminLoginController) EndRequest(ctx iris.Context) {}
func (c *AdminLoginController) Get() mvc.View {
	return mvc.View{
		Name: "backend",
		Data: iris.Map{
			"Account":  c.Account,
		},
	}
}

type UserLoginController struct {
	Username string
}
func (c *UserLoginController) BeginRequest(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("userauthenticated"); !auth {
		ctx.Redirect("/login")
		return
	}
	c.Username = sess.Start(ctx).GetString("Username")
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
