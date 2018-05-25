package main

import (
	"github.com/kataras/iris"
	_"github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/mvc"
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
	//app.Use(irisyaag.New()) // <- IMPORTANT, register the middleware.*/
	app.RegisterView(iris.HTML("html",".html").Reload(true))
	app.StaticWeb("/js", "./js") // serve our custom javascript code
	app.StaticWeb("/Picture","./Picture")
	app.Get("/adminlogin",func (ctx iris.Context) {ctx.View("adminlogin.html")})
	app.Get("/404",func (ctx iris.Context) {ctx.View("404.html")})
	app.Get("/login",func (ctx iris.Context) {ctx.View("userlogin.html")})
	app.Get("/error",func (ctx iris.Context) {ctx.View("error.html")})
	app.Get("/register",func(ctx iris.Context) {ctx.View("register.html")})
	app.Get("/File/{Name}",Controller.DownloadFile)
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context){ctx.View("404.html")})
	app.OnErrorCode(500, func(ctx iris.Context){ctx.View("500.html")})
	app.OnErrorCode(403, func(ctx iris.Context){ctx.View("403.html")})
	app.Post("/AdminLoginAjax",Controller.AdminLoginAjax)
	app.Post("/Register",Controller.Register)
	app.Post("/delete",Controller.Delete)
	app.Post("/modify",Controller.Modify)
	app.Post("/UserLoginAjax",Controller.UserLoginAjax)
	app.Post("/Comment",Controller.CommentController)
	app.Post("/artinsert",Controller.ArticleInsertController)
	app.Post("/articleModify",Controller.ArticlemodifyController)
	app.Post("/upload",Controller.Uploads)
	mvc.New(app.Party("/Classify/{Classify}/{page}")).Handle(new(Controller.ClassifyController))
	mvc.New(app.Party("/articleinsert")).Handle(new(Controller.ArticInsertController))
	mvc.New(app.Party("/backend/{page}")).Handle(new(Controller.AdminLoginController))
	mvc.New(app.Party("/adminlogout")).Handle(new(Controller.AdminLogout))
	mvc.New(app.Party("/chatform")).Handle(new(Controller.UserLoginController))
	mvc.New(app.Party("/logout")).Handle(new(Controller.Logout))
	mvc.New(app.Party("/")).Handle(new(Controller.IndexController))
	mvc.New(app.Party("/user")).Handle(new(Controller.UserIndexController))
	mvc.New(app.Party("/article/{id}")).Handle(new(Controller.ArticleController))
	mvc.New(app.Party("/articlemodifylist/{page}")).Handle(new(Controller.ArticleListController))
	mvc.New(app.Party("/menu/{id}/{page}")).Handle(new(Controller.MenuController))
	mvc.New(app.Party("/articlemodify/{id}")).Handle(new(Controller.ArticleModifyController))
	mvc.New(app.Party("/upload")).Handle(new(Controller.UploadController))
	mvc.New(app.Party("/download")).Handle(new(Controller.DownloadPageController))
	mvc.Configure(app.Party("/echo"), Controller.ConfigureMVC)
	app.Run(iris.Addr(":80"))
}

