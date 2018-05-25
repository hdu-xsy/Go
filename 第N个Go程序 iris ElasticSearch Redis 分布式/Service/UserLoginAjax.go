package Service

import (
	"github.com/kataras/iris"
	"../Entity"
	"time"
	"../UserIndex"
)

// 用户登录验证
type UserLoginAjax struct {

}

func (s *UserLoginAjax)Get(ctx iris.Context) int64{
	user := Entity.UserData{}
	checkError(ctx.ReadForm(&user))
	session := Entity.Sess.Start(ctx)
	if auth, _ := Entity.Sess.Start(ctx).GetBoolean("userauthenticated"); auth {
		ctx.WriteString("请勿重复登录帐号")
	} else if ok, _,userdata := userdatadao.Get(Entity.UserData{Username: user.Username}); ok {
		if userdata.Password == user.Password {
			uid = userdata.Id
			app.Logger().Println("user:"+user.Username+"  login")
			ctx.WriteString(" ")
			session.Set("userauthenticated", true)
			session.Set("Username",user.Username)
			userdatadao.Motify(Entity.UserData{Id:userdata.Id,Username:userdata.Username,Password:userdata.Password,SessionId:session.ID()})
		} else {
			ctx.WriteString("密码错误")
		}
	} else {
		ctx.WriteString("用户名不存在")
	}
	return uid
}

// 用户主页验证
 type UserLogin struct {

 }
 func (s *UserLogin)BeginRequest(ctx iris.Context){
	 if auth, _ := Entity.Sess.Start(ctx).GetBoolean("userauthenticated"); !auth {
		 ctx.Redirect("/login")
		 return
	 } else {
	 	id :=Entity.Sess.Start(ctx).ID()
	 	username := Entity.Sess.Start(ctx).GetString("Username")
	 	_,_,user := userdatadao.Get(Entity.UserData{Username:username})
	 	if id != user.SessionId {
	 		Entity.Sess.Start(ctx).Delete("Username")
	 		Entity.Sess.Start(ctx).Delete("userauthenticated")
			ctx.Redirect("/login")
			return
		}
	 }
 }

 func (s *UserLogin)Get(ctx iris.Context) {
 	menulist := menudao.GetAll()
 	entity := Entity.Entity{MenuList:menulist}
 	UserIndex.UserIndexWriter(ctx,ctx,entity)
 }
 //聊天页面验证
 type ChatForm struct {

 }
 func (s *ChatForm)BeginRequest(ctx iris.Context) {
	 username := Entity.Sess.Start(ctx).GetString("Username")
	 if ok, _, userdata := userdatadao.Get(Entity.UserData{Username: username}); ok {
		 uid = userdata.Id
		 if ok, _, _ := onlineuserdao.Get(Entity.OnlineUser{Username: userdata.Username}); ok {
			 ctx.Redirect("/error")
		 } else {
		 	checkError(onlineuserdao.Insert(Entity.OnlineUser{Uid: uid, Username: username, Logintime: time.Now()}))
		 }
	 }
 }
