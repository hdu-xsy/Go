package Service

import(
	"github.com/kataras/iris"
	"../DELETED/userlist"
	"../Entity"
	"strconv"
	"time"
	"html/template"
)

//文章页面
type ArticlePageService struct {

}
type ArticleArticle struct {
	Title		string
	Classify	string
	Id			int64
	Time		string
	Content		template.HTML
	Writer		string
}
type ArticleComment struct {
	Content		template.HTML
	Time		string
	Username	string
	Floor		int64
}
func (s *ArticlePageService)BeginRequest(ctx iris.Context) {
	articleall := articledao.FindAllA()
	id,_ := strconv.ParseInt(ctx.Params().Get("id"),10,64)
	if int(id) > len(articleall) {
		ctx.Redirect("/404")
		return
	}
}
func (s *ArticlePageService)Get(ctx iris.Context) {
	h1 := redisdao.Get("h1")
	var auth string
	if userauth, _ := Entity.Sess.Start(ctx).GetBoolean("userauthenticated"); !userauth { auth = "false" } else { auth = "true" }
	username := Entity.Sess.Start(ctx).GetString("Username")
	ctx.ViewData("username",username)

	id,_ := strconv.ParseInt(ctx.Params().Get("id"),10,64)
	_,_,article := articledao.Get(Entity.Article{Id:id})
	_,_,user := userdatadao.Get(Entity.UserData{Id:article.User})
	var articleList [1]ArticleArticle
	articleList[0] = ArticleArticle{
		Title:article.Title,
		Classify:article.Classify,
		Id:article.Id,
		Time:article.Time.Format("2006-01-02 15:04:05"),
		Content:template.HTML(article.Content),
		Writer:user.Username,
	}

	_,_,pre := articledao.Get(Entity.Article{Id:id-1})
	_,_,suc := articledao.Get(Entity.Article{Id:id+1})
	preId := pre.Id
	preTitle := pre.Title
	sucId := suc.Id
	sucTitle := suc.Title

	mid,_ := strconv.ParseInt(article.Menu,10,64)
	_,_,menu := menudao.Get(Entity.Menu{Id:mid})
	var menuList [1]Entity.Menu
	menuList[0] = menu
	menulist := menudao.GetAll()

	comment := commentdao.FindAll(ctx.Params().Get("id"))
	commentList := make([]ArticleComment,len(comment))
	for i,v := range comment {
		_,_,commentuser := userdatadao.Get(Entity.UserData{Id:v.User})
		commentList[i].Content = template.HTML(v.Content)
		commentList[i].Username = commentuser.Username
		commentList[i].Time = v.Time.Format("2006-01-02 15:04:05")
		commentList[i].Floor = v.Floor
	}

	ctx.ViewData("h1",h1)
	ctx.ViewData("menuList",menulist)
	ctx.ViewData("auth",auth)
	ctx.ViewData("username",username)
	ctx.ViewData("menu",menuList)
	ctx.ViewData("article",articleList)
	ctx.ViewData("preId",preId)
	ctx.ViewData("preTitle",preTitle)
	ctx.ViewData("sucId",sucId)
	ctx.ViewData("sucTitle",sucTitle)
	ctx.ViewData("commentSum",len(comment))
	ctx.ViewData("commentList",commentList)
	ctx.ViewData("commentSumPlus",len(comment)+1)
	ctx.View("article.html")
}

//插入文章
type ArticleInsertService struct {

}
func (s *ArticleInsertService)Get(ctx iris.Context) {
	classify := ctx.PostValue("Classify")
	title := ctx.PostValue("Title")
	menu := ctx.PostValue("Menu")
	content := ctx.PostValue("Content")
	article := Entity.Article{User:1,Time:time.Now(),Classify:classify,Title:title,Menu:menu,Content:content}
	//checkError(ctx.ReadForm(&article))
	app.Logger().Println(article)
	//article.Classify=ctx.PostValue("Classify")
	//app.Logger().Println("Classify:"+article.Classify)
	articledao.Insert(article)
	ctx.Redirect("/backend/1")
	//ctx.WriteString(" ")
}

//文章修改列表
type ArticleModifyListPage struct {

}

func (s *ArticleModifyListPage)Get(ctx iris.Context) {
	article := articledao.FindAllA()
	var suc,max int
	page,_ := strconv.Atoi(ctx.Params().Get("page"))
	if len(article)%20 == 0 {
		max = len(article)/20
	} else {
		max = len(article)/20 + 1
	}
	if page>max || page == 0 {
		ctx.Redirect("/404")
		return
	}
	if page * 20 >= len(article) {
		suc = len(article)
	} else {
		suc = page * 20
	}
	for i,a := range article {
		id,_ := strconv.ParseInt(a.Menu,10,64)
		_,_,menu := menudao.Get(Entity.Menu{Id:id})
		article[i].Menu = menu.Name
	}
	userlist.ArticleListToWriter(article[(page-1)*20:suc],page,max,ctx)
}

//文章修改页面
type ArticleModifyPage struct {

}
func (s *ArticleModifyPage)Update(ctx iris.Context) {
	id,_ := strconv.ParseInt(ctx.Params().Get("id"),10,64)
	_,_,article := articledao.Get(Entity.Article{Id:id})
	var articleList [1]Entity.Article
	articleList[0].Content = article.Content
	articleList[0].Title = article.Title
	articleList[0].Id = article.Id
	articleList[0].Classify = article.Classify
	menulist := menudao.GetAll()
	menu,_ := strconv.Atoi(article.Menu)
	//articleModify.ArticleToWriter(menulist,article,ctx)
	ctx.ViewData("menuList",menulist)
	ctx.ViewData("menu",menu-1)
	ctx.ViewData("article",articleList)
	ctx.View("articlemodify.html")
}

//文章修改
type ArticleModify struct {

}
func (s *ArticleModify)Update(ctx iris.Context) {
	id,_ := strconv.ParseInt(ctx.PostValue("Id"),10,64)
	title := ctx.PostValue("Title")
	menu := ctx.PostValue("Menu")
	content := ctx.PostValue("Content")
	classify := ctx.PostValue("Classify")
	timenow := time.Now()
	article := Entity.Article{Id:id,User:1,Time:timenow,Title:title,Menu:menu,Classify:classify,Content:content}
	articledao.Update(article)
	ctx.Redirect("/articlemodifylist/1")
	return
}

//文章插入页面
type ArticleInsert struct {

}
func (s *ArticleInsert) Get(ctx iris.Context) {
	menuList := menudao.GetAll()
	ctx.ViewData("menuList",menuList)
	ctx.View("articleinsert.html")
}
