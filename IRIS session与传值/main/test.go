package main

import (
	"fmt"
	"os"
	//"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	_"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/mvc"
)

var app = iris.New()

func checkError(err error) {
	if err != nil {
		fmt.Println("Error is ", err)
		os.Exit(-1)
	}
}
type AdminUser struct {
	Id       int
	Account  string
	Password string
}
var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)
func index(ctx iris.Context) {
	ctx.Redirect("/adminlogin")
}
func adminlogin(ctx iris.Context) {
	if err := ctx.View("adminlogin.html");err!=nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
}
func fourzerofour(ctx iris.Context) {
	if err := ctx.View("404.html");err!=nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
}
func backend(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.Redirect("/adminlogin")
		return
	}
	if err := ctx.View("backend.html");err!=nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
	session := sess.Start(ctx)
	if v,ok := session.Get("Account").(string);ok {
		ctx.WriteString(v)
	}
	ctx.JSON(iris.Map{"result": "Hello World!"})
}
func AdminLoginAjax(ctx iris.Context) {
	user := AdminUser{}
	err := ctx.ReadForm(&user)
	if err !=nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		ctx.Redirect("/404")
	}
	orm,err := xorm.NewEngine("mysql", "root:Xsydx886.@/javaweb?charset=utf8")
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
	}
	iris.RegisterOnInterrupt(func(){
		orm.Close()
	})
	err = orm.Sync2(new(AdminUser))
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized User table: %v", err)
	}
	adminuser := AdminUser{Account: user.Account}
	if ok, _ := orm.Get(&adminuser); ok {
		if adminuser.Password == user.Password {
			app.Logger().Println(user.Account+"   "+user.Password+"   login")
			ctx.WriteString("TRUE")
			session := sess.Start(ctx)
			session.Set("authenticated", true)
			session.Set("Account",user.Account)
		} else {
			ctx.WriteString("密码错误")
		}
	} else {
		ctx.WriteString("用户名不存在")
	}
	return
}
func adminlogout(ctx iris.Context) {
	session := sess.Start(ctx)
	session.Set("authenticated",false)
	session.Set("Account","nil")
}

func main() {
	app.RegisterView(iris.HTML("html",".html").Reload(true))
	app.Get("/",index)
	app.Get("/adminlogin",adminlogin)
	app.Get("/404",fourzerofour)
	mvc.New(app.Party("/backend")).Handle(new(AdminLoginController))
	//app.Get("/backend",backend)
	app.Post("/AdminLoginAjax",AdminLoginAjax)
	app.Get("/adminlogout",adminlogout)
	//app.Get("test",test)
	app.Run(iris.Addr(":4567"))
}

type AdminLoginController struct {
	Account string
}
func (c *AdminLoginController) BeginRequest(ctx iris.Context) {
	session := sess.Start(ctx)
	c.Account = session.Get("Account").(string)
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