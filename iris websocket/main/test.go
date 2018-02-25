package main

import (
	"fmt"
	"os"
	"github.com/kataras/iris"
	_"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/websocket"
)

var app = iris.New()
func checkError(err error) {
	if err != nil {
		fmt.Println("Error is ", err)
		os.Exit(-1)
	}
}
type AdminUser struct {
	Id       int64`pk`
	Account  string
	Password string
}
type UserData struct {
	Id       int64`pk`
	Username string
	Password string
}
var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)
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
func UserLoginAjax(ctx iris.Context) {
	user := UserData{}
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
	err = orm.Sync2(new(UserData))
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized User table: %v", err)
	}
	userdata := UserData{Username: user.Username}
	if ok, _ := orm.Get(&userdata); ok {
		if userdata.Password == user.Password {
			app.Logger().Println("user    "+user.Username+"   "+user.Password+"   login")
			ctx.WriteString("TRUE")
			session := sess.Start(ctx)
			session.Set("userauthenticated", true)
			session.Set("Username",user.Username)
		} else {
			ctx.WriteString("密码错误")
		}
	} else {
		ctx.WriteString("用户名不存在")
	}
	return
}
func userlogout(ctx iris.Context) {
	session := sess.Start(ctx)
	session.Set("userauthenticated",false)
	session.Set("Username","nil")
}
func main() {
	app.RegisterView(iris.HTML("html",".html").Reload(true))
	app.Get("/",func (ctx iris.Context) {
		ctx.View("index.html")
	})
	app.Get("/adminlogin",func (ctx iris.Context) {
		ctx.View("adminlogin.html")
	})
	app.Get("/404",func (ctx iris.Context) {
		ctx.View("404.html")
	})
	// when 404 then render the template $templatedir/errors/404.html
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
	setupWebsocket(app)
	//app.Get("test",test)
	app.Run(iris.Addr(":4567"))
}

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
	session := sess.Start(ctx)
	c.Username = session.GetString("Username")
}
func (c *UserLoginController) EndRequest(ctx iris.Context) {
	userlogout(ctx)
}
func (c *UserLoginController) Get() mvc.View {
	return mvc.View{
		Name: "chatform",
		Data: iris.Map{
			"Username":  c.Username,
		},
	}
}
func setupWebsocket(app *iris.Application) {
	// create our echo websocket server
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})
	ws.OnConnection(handleConnection)
	// register the server on an endpoint.
	app.Get("/echo", ws.Handler())
	// see html script tags, this path is used.
	app.Any("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(websocket.ClientSource)
	})
}

func handleConnection(c websocket.Connection) {
	c.On("chat", func(msg string) {
		// fmt.Printf("%s sent: %s\n", c.Context().RemoteAddr(), msg)
		// Write message back to the client message owner with: c.Emit("chat", msg)
		// Write message to all except this client with:
		c.To(websocket.Broadcast).Emit("chat",  msg)
	})
}