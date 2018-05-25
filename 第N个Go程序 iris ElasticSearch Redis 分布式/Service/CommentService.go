package Service

import (
	"github.com/kataras/iris"
	"../Entity"
	"strconv"
	"time"
)
type Comment struct {

}
func (s *Comment)Get(ctx iris.Context) {
	if auth, _ := Entity.Sess.Start(ctx).GetBoolean("userauthenticated"); !auth {
		ctx.WriteString("请先登录")
	}else {
		article, _ := strconv.ParseInt(ctx.PostValue("Article"), 10, 64)
		comment := ctx.PostValue("Comment")
		floor, _ := strconv.ParseInt(ctx.PostValue("Floor"), 10, 64)
		_, _, id := userdatadao.Get(Entity.UserData{Username: Entity.Sess.Start(ctx).GetString("Username")})
		commentdao.Insert(Entity.Comment{Article: article, Floor: floor, User: id.Id, Time: time.Now(), Content: comment})
		ctx.WriteString(" ")
	}
	return
}
