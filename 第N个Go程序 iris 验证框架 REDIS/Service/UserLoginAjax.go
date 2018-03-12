package Service

import (
	"github.com/kataras/iris"
	"time"
	"../Entity"
)

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
			if ok,_,olu:= onlineuserdao.Get(Entity.OnlineUser{Username:userdata.Username});ok {
				ctx.WriteString("该用户正在线上"+"("+olu.Username+")")
			}
			app.Logger().Println("user:"+user.Username+"  login")
			ctx.WriteString(" ")
			session.Set("userauthenticated", true)
			session.Set("Username",user.Username)
			uid = userdata.Id
			checkError(onlineuserdao.Insert(Entity.OnlineUser{Id:uid,Username:user.Username,Logintime:time.Now()}))
		} else {
			ctx.WriteString("密码错误")
		}
	} else {
		ctx.WriteString("用户名不存在")
	}
	return uid
}
