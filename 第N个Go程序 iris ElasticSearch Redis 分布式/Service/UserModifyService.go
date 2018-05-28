package Service

import (
	"github.com/kataras/iris"
	"strconv"
)
type UserModifyPage struct {

}

func (s *UserModifyPage)Get(ctx iris.Context) {
	userList := userdatadao.FindAll()
	var suc,max int
	page,_ := strconv.Atoi(ctx.Params().Get("page"))
	var c1,c2 string
	if page == 1 {
		c1 = "disabled"
	}else if page == max {
		c2 = "disabled"
	}
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
	nav := make([]Nav,5)
	for k := range nav {
		if max >5 {
			nav[k].Page = page + k
		} else {
			nav[k].Page = k + 1
		}
		flag := 0
		if (page == k+1 || page > 5) && flag == 0 {
			nav[k].Class = "active"
			flag = 1
		}
		if k+1 > max {
			nav[k].Class = "disabled"
		}
		nav[k].Num = k+1
	}
	//userlist.UserListToWriter(page,max,userList[(page-1)*20:suc], ctx)
	ctx.ViewData("userList",userList[(page-1)*20:suc])
	ctx.ViewData("pagedec",page-1)
	ctx.ViewData("pageplus",page+1)
	ctx.ViewData("c1",c1)
	ctx.ViewData("c2",c2)
	ctx.ViewData("nav",nav)
	ctx.View("userlist.html")
}
