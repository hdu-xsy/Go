package main

import (
	"github.com/kataras/iris"
	_"github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/mvc"
	"../index"
	"../Controller"
)

func main() {
	var app = iris.New()
/*	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: "Iris",
		DocPath:  "./apidoc/apidoc.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	app.Use(irisyaag.New()) // <- IMPORTANT, register the middleware.*/
	app.RegisterView(iris.HTML("html",".html").Reload(true))
	app.StaticWeb("/js", "./js") // serve our custom javascript code
	app.Get("/adminlogin",func (ctx iris.Context) {ctx.View("adminlogin.html")})
	app.Get("/404",func (ctx iris.Context) {ctx.View("404.html")})
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context){ctx.View("404.html")})
	app.OnErrorCode(500, func(ctx iris.Context){ctx.View("500.html")})
	app.Get("/login",func (ctx iris.Context) {ctx.View("userlogin.html")})
	app.Get("/register",func(ctx iris.Context) {ctx.View("register.html")})
	app.Get("/articleinsert",func (ctx iris.Context) {ctx.View("articleinsert.html")})
	mvc.New(app.Party("/backend")).Handle(new(Controller.AdminLoginController))
	app.Post("/AdminLoginAjax",Controller.AdminLoginAjax)
	app.Post("/Register",Controller.Register)
	mvc.New(app.Party("/adminlogout")).Handle(new(Controller.AdminLogout))
	mvc.New(app.Party("/chatform")).Handle(new(Controller.UserLoginController))
	mvc.New(app.Party("/logout")).Handle(new(Controller.Logout))
	app.Post("/UserLoginAjax",Controller.UserLoginAjax)
	app.Post("/artinsert",Controller.ArticleInsertController)
	app.Get("/", func(ctx iris.Context) {
		var data = "aaa"
		index.ContextWriter(data, ctx)
	})
	mvc.New(app.Party("/article/{id}")).Handle(new(Controller.ArticleController))
	mvc.New(app.Party("/menu/{id}")).Handle(new(Controller.MenuController))
	mvc.Configure(app.Party("/echo"), Controller.ConfigureMVC)
	app.Run(iris.Addr(":80"))
}

