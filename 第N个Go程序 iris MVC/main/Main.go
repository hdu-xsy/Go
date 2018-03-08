package main

import (
	"github.com/kataras/iris"
	_"github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/mvc"
	"../index"
)

var app = iris.New()
var uid int64
func checkError(err error) {
	if err != nil {
		app.Logger().Fatalf("err:",err)
	}
}
var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)
func main() {
	app.RegisterView(iris.HTML("html",".html").Reload(true))
	app.StaticWeb("/js", "./js") // serve our custom javascript code
	setupWebsocket(app)
	app.Get("/iii",func (ctx iris.Context) {
		ctx.View("index.html")
	})
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
		ctx.View("404.html")
	})
	app.Get("/login",func (ctx iris.Context) {
		ctx.View("userlogin.html")
	})
	mvc.New(app.Party("/backend")).Handle(new(AdminLoginController))
	app.Post("/AdminLoginAjax",AdminLoginAjax)
	app.Get("/adminlogout",adminlogout)
	mvc.New(app.Party("/chatform")).Handle(new(UserLoginController))
	app.Get("/logout",userlogout)
	app.Post("/UserLoginAjax",UserLoginAjax)
	app.Get("/", func(ctx iris.Context) {
		var data = []string{"data",}
		index.UserListToWriter(data, ctx)
	})
	app.Run(iris.Addr(":80"))
}

