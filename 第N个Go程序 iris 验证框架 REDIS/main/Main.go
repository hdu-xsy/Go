package main

import (
	"github.com/kataras/iris"
	_"github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/mvc"
	"../index"
	"../Controller"
)

var app = iris.New()
func main() {
	app.RegisterView(iris.HTML("html",".html").Reload(true))
	app.StaticWeb("/js", "./js") // serve our custom javascript code
	app.Get("/adminlogin",func (ctx iris.Context) {
		ctx.View("adminlogin.html")
	})
	app.Get("/404",func (ctx iris.Context) {
		ctx.View("404.html")
	})
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context){
		ctx.View("404.html")
	})
	app.OnErrorCode(500, func(ctx iris.Context){
		ctx.View("500.html")
	})
	app.Get("/login",func (ctx iris.Context) {
		ctx.View("userlogin.html")
	})
	app.Get("/register",func(ctx iris.Context) {
		ctx.View("register.html")
	})
	mvc.New(app.Party("/backend")).Handle(new(Controller.AdminLoginController))
	app.Post("/AdminLoginAjax",Controller.AdminLoginAjax)
	app.Post("/Register",Controller.Register)
	mvc.New(app.Party("/adminlogout")).Handle(new(Controller.AdminLogout))
	mvc.New(app.Party("/chatform")).Handle(new(Controller.UserLoginController))
	mvc.New(app.Party("/logout")).Handle(new(Controller.Logout))
	app.Post("/UserLoginAjax",Controller.UserLoginAjax)
	app.Get("/", func(ctx iris.Context) {
		var data = []string{"data",}
		index.UserListToWriter(data, ctx)
	})
	mvc.Configure(app.Party("/echo"), Controller.ConfigureMVC)
	app.Run(iris.Addr(":80"))
}

