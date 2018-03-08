package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"../Util"
	"../userlist"
	"time"
)

//管理员登陆验证
type AdminLoginController struct {
}
func (c *AdminLoginController) BeginRequest(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.Redirect("/adminlogin")
		return
	}
}
func (c *AdminLoginController) EndRequest(ctx iris.Context) {}
func (c *AdminLoginController) Get(ctx iris.Context) {
	userList := FindAllUser()
	userlist.UserListToWriter(userList, ctx)
}

//用户登录验证
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
			uid = userdata.Id
			var onlineuser = OnlineUser{Id:uid,Username:user.Username,Logintime:time.Now()}
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
	uid = 0
	ctx.Redirect("/")
}