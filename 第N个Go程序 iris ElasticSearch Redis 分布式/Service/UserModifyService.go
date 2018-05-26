package Service

import (
	"../userlist"
	"github.com/kataras/iris"
	"strconv"
)
type UserModifyPage struct {

}

func (s *UserModifyPage)Get(ctx iris.Context) {
	userList := userdatadao.FindAll()
	var suc,max int
	page,_ := strconv.Atoi(ctx.Params().Get("page"))
	if len(userList)%20 == 0 {
		max = len(userList)/20
	} else {
		max = len(userList)/20 + 1
	}
	if page>max || page == 0 {
		ctx.Redirect("/404")
		return
	}
	if page * 20 >= len(userList) {
		suc = len(userList)
	} else {
		suc = page * 20
	}
	userlist.UserListToWriter(page,max,userList[(page-1)*20:suc], ctx)
}
