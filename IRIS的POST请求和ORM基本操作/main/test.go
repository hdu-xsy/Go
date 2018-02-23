package main

import (
	"fmt"
	"os"
	//"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	_"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
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
type AdminLogin struct {
	Account  string
	Password string
}
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
	if err := ctx.View("backend.html");err!=nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
}
func admincheck(ctx iris.Context) {
	user := AdminLogin{}
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
			ctx.Redirect("/backend")
		} else {
			ctx.Redirect("/adminlogin")
		}
	} else {
		ctx.Redirect("/adminlogin")
	}
}
func main() {
	app.RegisterView(iris.HTML("html",".html").Reload(true))
	app.Get("/",index)
	app.Get("/adminlogin",adminlogin)
	app.Post("/admincheck",admincheck)
	app.Get("/404",fourzerofour)
	app.Get("/backend",backend)
	app.Run(iris.Addr(":4567"))
}