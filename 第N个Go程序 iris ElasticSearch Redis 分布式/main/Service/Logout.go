package Service

import (
	"github.com/kataras/iris"
	"../Entity"
)
type Logout struct {

}

func (s *Logout) AdminLogout(ctx iris.Context) {
	session := Entity.Sess.Start(ctx)
	session.Set("authenticated",false)
	session.Set("Account","nil")
	ctx.Redirect("/")
}
func (s *Logout) Logout(ctx iris.Context) {
	session := Entity.Sess.Start(ctx)
	checkError(onlineuserdao.Delete(Entity.OnlineUser{Username:session.GetString("Username")}))
	session.Set("userauthenticated",false)
	session.Set("Username","nil")
	uid = -1
	ctx.Redirect("/")
}