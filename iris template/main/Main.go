package main

import (
	"github.com/kataras/iris"
	_"github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/mvc"
	"time"
	"../index"
	"../userlist"
	"../Util"
)

var app = iris.New()
var uuu int64
func checkError(err error) {
	if err != nil {
		app.Logger().Fatalf("err:",err)
	}
}
type AdminUser struct {
	Id       int64`pk`
	Account  string`unique`
	Password string
}
type UserData struct {
	Id       int64`pk`
	Username string`unique`
	Password string
}
type OnlineUser struct {
	Id        int64
	Username  string`unique`
	Logintime time.Time`created`
}
var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)
func AdminLoginAjax(ctx iris.Context) {
	user := AdminUser{}
	err := ctx.ReadForm(&user)
	checkError(err)
	orm := Util.GetAdminUser(*app)
	adminuser := AdminUser{Account: user.Account}
	if ok, _ := orm.Get(&adminuser); ok {
		if adminuser.Password == user.Password {
			app.Logger().Println(user.Account+"   "+user.Password+"   login")
			ctx.WriteString("TRUE")
			session :=sess.Start(ctx)
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
	ctx.Redirect("/")
}
func UserLoginAjax(ctx iris.Context) {
	user := UserData{}
	err := ctx.ReadForm(&user)
	checkError(err)
	orm := Util.GetUserData(*app)
	userdata := UserData{Username: user.Username}
	if ok, _ := orm.Get(&userdata); ok {
		if userdata.Password == user.Password {
			olu := OnlineUser{Username:userdata.Username}
			if ok,_ := orm.Get(&olu);ok {
				ctx.WriteString("该用户正在线上")
				return
			}
			app.Logger().Println("user    "+user.Username+"   "+user.Password+"   login")
			ctx.WriteString("TRUE")
			session := sess.Start(ctx)
			session.Set("userauthenticated", true)
			session.Set("Username",user.Username)
			uuu = userdata.Id
			var onlineuser = OnlineUser{Id:uuu,Username:user.Username,Logintime:time.Now()}
			orm.Insert(&onlineuser)
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
	orm := Util.GetOnlineUser(*app)
	onlineuser := OnlineUser{Username:session.GetString("Username")}
	orm.Delete(&onlineuser)
	session.Set("userauthenticated",false)
	session.Set("Username","nil")
	uuu = 0
	ctx.Redirect("/")
}
func main() {
	app.RegisterView(iris.HTML("html",".html").Reload(true))
	app.StaticWeb("/js", "./js") // serve our custom javascript code
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
	//mvc.New(app.Party("/backend")).Handle(new(AdminLoginController))
	app.Post("/AdminLoginAjax",AdminLoginAjax)
	app.Get("/adminlogout",adminlogout)
	mvc.New(app.Party("/chatform")).Handle(new(UserLoginController))
	app.Get("/logout",userlogout)
	app.Post("/UserLoginAjax",UserLoginAjax)
	setupWebsocket(app)
	app.Get("/", func(ctx iris.Context) {
		var data = []string{
			"data",
		}
		index.UserListToWriter(data, ctx)
	})
	app.Get("/backend", func(ctx iris.Context) {
		if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
			ctx.Redirect("/adminlogin")
			return
		}
		orm := Util.GetOnlineUser(*app)
		var userList []userlist.UserData
		orm.Find(&userList)
		userlist.UserListToWriter(userList, ctx)
	})
	app.Run(iris.Addr(":80"))
}

