package Service

import (
	"../userlist"
	"github.com/kataras/iris"
)
type AdminLogin struct {

}

func (s *AdminLogin)Get(ctx iris.Context) {
	userList := userdatadao.FindAll()
	userlist.UserListToWriter(userList, ctx)
}