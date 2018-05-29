package Service

import(
	"github.com/kataras/iris"
	"../Entity"
	"strconv"
)
type Register struct {

}

func (s *Register) Get(ctx iris.Context) {
	user := Entity.UserData{}
	checkError(ctx.ReadForm(&user))
	if ok,_,_ := userdatadao.Get(Entity.UserData{Username:user.Username});ok {
		ctx.WriteString("用户名已存在")
		return
	}else{
		userdatadao.Insert(user)
		ctx.WriteString(" ")
	}
	return
}

type Delete struct {

}

func (s *Delete) Get(ctx iris.Context) {
	id,_ := strconv.ParseInt(ctx.PostValue("Id"),10,64)
	username := ctx.PostValue("Username")
	password := ctx.PostValue("Password")
	user := Entity.UserData{Id:id,Username:username,Password:password}
	userdatadao.Delete(user)
	ctx.Redirect("/backend/1")
}

type Modity struct {

}

func (s *Modity) Get(ctx iris.Context)  {
	id,_ := strconv.ParseInt(ctx.PostValue("Id"),10,64)
	username := ctx.PostValue("Username")
	password := ctx.PostValue("Password")
	user := Entity.UserData{Id:id,Username:username,Password:password}
	userdatadao.Motify(user)
	ctx.Redirect("/backend/1")
}