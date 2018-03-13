package Service

import(
	"github.com/kataras/iris"
	"../Entity"
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