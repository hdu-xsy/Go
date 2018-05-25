package Service

import (
	"github.com/kataras/iris"
	"../Entity"
)

//首页
type IndexService struct {
}

type IndexArticle struct {
	Id	int64
	Time string
	Title string
}
type IndexComment struct {
	Article int64
	Content string
	Time 	string
}
func (s *IndexService)Get(ctx iris.Context) {
	articlelist := articledao.OrderByTime()
	n := len(articlelist)
	articleList := make([]IndexArticle,n)
	for i,v := range articlelist {
		articleList[i].Id = v.Id
		articleList[i].Title = v.Title
		articleList[i].Time = v.Time.Format("2006-01-02 15:04:05")
	}
	comment := commentdao.OrderByTime()
	n = len(comment)
	commentList := make([]IndexComment,n)
	for i,v := range comment {
		commentList[i].Article = v.Article
		commentList[i].Time = v.Time.Format("2006-01-02 15:04:05")
		if len(v.Content)>30 {
			commentList[i].Content = v.Content[0:30]+"..."
		} else {
			commentList[i].Content = v.Content
		}
	}
	menulist := menudao.GetAll()
	var auth string
	if userauth, _ := Entity.Sess.Start(ctx).GetBoolean("userauthenticated"); !userauth { auth = "false" } else { auth = "true" }
	username := Entity.Sess.Start(ctx).GetString("Username")
	ctx.ViewData("menuList",menulist)
	ctx.ViewData("auth",auth)
	ctx.ViewData("username",username)
	ctx.ViewData("articlesum",len(articleList))
	ctx.ViewData("articleList",articleList)
	ctx.ViewData("commentList",commentList)
	ctx.View("index.html")
}
