package Service

import (
	"github.com/kataras/iris"
	"../Entity"
)

type AdminLoginAjax struct {

}

func (s *AdminLoginAjax) Get(ctx iris.Context) {
	user := Entity.AdminUser{}
	checkError(ctx.ReadForm(&user))
	if ok,_,adminuser := adminuserdao.Get(Entity.AdminUser{Account: user.Account}); ok {
		if adminuser.Password == user.Password {
			app.Logger().Println("admin:"+user.Account+"  login")
			ctx.WriteString(" ")
			session :=Entity.Sess.Start(ctx)
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